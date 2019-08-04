package cookie

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"runtime"
	"strings"

	//nolint:golint
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"golang.org/x/crypto/pbkdf2"
)

// Inspiration
// http://n8henrie.com/2013/11/use-chromes-cookies-for-easier-downloading-with-python-requests/
// https://gist.github.com/dacort/bd6a5116224c594b14db
// https://stackoverflow.com/questions/23153159/decrypting-chromium-cookies/23727331#23727331

// This code works on only Mac OS

const BlockSize = 16

// This path would be changed by environment even same OS
//var cookieBaseDir = map[string]string{
//	"darwin": "%s/Library/Application Support/Google/Chrome/Default/Cookies", //mac
//	"linux":  "%s/.config/google-chrome/Default/Cookies",
//}
var cookieBaseDir = map[string]string{
	"darwin": "%s/Library/Application Support/Google/Chrome/%s/Cookies", //mac
	"linux":  "%s/.config/google-chrome/%s/Cookies",
}

// Chromium Mac os_crypt:  http://dacort.me/1ynPMgx
var (
	profileName = "Default"
	salt        = "saltysalt"
	iv          = "                "
	password    = ""
	iterations  = 1003
)

// Cookie - Items for a cookie
type Cookie struct {
	Domain         string
	Key            string
	Value          string
	EncryptedValue []byte
}

func init() {
	var err error
	switch runtime.GOOS {
	case "darwin":
		password, err = getPasswordMac()
		if err != nil {
			log.Printf("failed to call getPassword: %s", err)
		}
	case "linux":
		iterations = 1
		password, err = getPasswordLinux()
		if err != nil {
			log.Printf("failed to call getPassword: %s", err)
		}
	default:
		//not supported
	}
}

//func callerSample() {
//	domain := "localhost"
//	PrintCookies(domain)
//
//	_ = GetValue(domain, "key")
//}

func SetProfile(name string) {
	profileName = name
}

func PrintCookies(url string) error {
	cookies, err := getCookies(url)
	if err != nil {
		return err
	}

	for _, cookie := range cookies {
		decrypted, err := cookie.DecryptedValue()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("%s/%s: %s\n", cookie.Domain, cookie.Key, decrypted)
	}
	//localhost/cookiename: xxxxxx

	return nil
}

func GetValue(url, key string) (string, error) {
	cookies, err := getCookies(url)
	if err != nil {
		return "", err
	}

	for _, cookie := range cookies {
		if cookie.Domain == url && cookie.Key == key {
			return cookie.DecryptedValue()
		}
	}
	return "", nil
}

func GetAllValue(url string) (map[string]string, error) {
	decryptedCookies := make(map[string]string)

	cookies, err := getCookies(url)
	if err != nil {
		return nil, err
	}

	for _, cookie := range cookies {
		decrypted, err := cookie.DecryptedValue()
		if err != nil {
			log.Println(err)
			continue
		}
		decryptedCookies[cookie.Key] = decrypted
	}
	return decryptedCookies, nil
}

// DecryptedValue - Get the unencrypted value of a Chrome cookie
func (c *Cookie) DecryptedValue() (string, error) {
	if c.Value > "" {
		return c.Value, nil
	}

	if len(c.EncryptedValue) > 0 {
		encryptedValue := c.EncryptedValue[3:]
		return decryptValue(encryptedValue)
	}

	return "", nil
}

func decryptValue(encryptedValue []byte) (string, error) {
	key := pbkdf2.Key([]byte(password), []byte(salt), iterations, BlockSize, sha1.New)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	decrypted := make([]byte, len(encryptedValue))
	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	cbc.CryptBlocks(decrypted, encryptedValue)

	plainText, err := aesStripPadding(decrypted)
	if err != nil {
		return "", errors.Errorf("Error decrypting: %s", err)
	}
	return string(plainText), nil
}

// In the padding scheme the last <padding length> bytes
// have a value equal to the padding length, always in (1,16]
func aesStripPadding(data []byte) ([]byte, error) {
	if len(data)%BlockSize != 0 {
		//log.Printf("decrypted data block length is not a multiple of %d", BlockSize)
		return nil, errors.Errorf("decrypted data block length is not a multiple of %d", BlockSize)
	}
	paddingLen := int(data[len(data)-1])
	if paddingLen > BlockSize {
		//log.Printf("invalid last block padding length: %d", paddingLen)
		return nil, errors.Errorf("invalid last block padding length: %d", paddingLen)
	}
	return data[:len(data)-paddingLen], nil
}

func getPasswordMac() (string, error) {
	//this command is for only MacOS
	parts := strings.Fields("security find-generic-password -wga Chrome")

	cmd := parts[0]
	parts = parts[1:]

	out, err := exec.Command(cmd, parts...).Output()
	if err != nil {
		return "", errors.Errorf("failed to call security command to find password: %s", err)
	}

	return strings.Trim(string(out), "\n"), nil
}

func getPasswordLinux() (string, error) {
	//this command is for only Linux and `libsecret-tools` is required
	//`sudo apt install libsecret-tools`
	parts := strings.Fields("secret-tool search application chrome")

	cmd := parts[0]
	parts = parts[1:]

	out, err := exec.Command(cmd, parts...).Output()
	if err != nil {
		return "", errors.Errorf("failed to call secret-tool command to find password: %s\n `sudo apt install libsecret-tools`", err)
	}

	//retrieve
	//label = Chrome Safe Storage
	//secret = xxxxxxxxxxxxx
	//created = 2019-01-08 03:32:22
	//modified = 2019-01-08 03:32:22
	//schema = chrome_libsecret_os_crypt_password_v2
	//attribute.application = chrome
	ret := strings.Split(string(out), "\n")
	for _, val := range ret {
		tmp := strings.Split(val, " = ")
		if len(tmp) != 2 {
			continue
		}
		if tmp[0] == "secret" {
			return tmp[1], nil
		}
	}
	return "", errors.New("password is not found")
}

func getCookies(domain string) ([]Cookie, error) {
	var cookies []Cookie
	usr, _ := user.Current()

	var cookiesFile string
	if val, ok := cookieBaseDir[runtime.GOOS]; ok {
		cookiesFile = fmt.Sprintf(val, usr.HomeDir, profileName)
	} else {
		return nil, errors.Errorf("os[%s] is not supported ", runtime.GOOS)
	}

	db, err := sql.Open("sqlite3", cookiesFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, value, host_key, encrypted_value FROM cookies WHERE host_key like ?", fmt.Sprintf("%%%s%%", domain))
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var name, value, hostKey string
		var encryptedValue []byte
		rows.Scan(&name, &value, &hostKey, &encryptedValue)
		cookies = append(cookies, Cookie{hostKey, name, value, encryptedValue})
	}

	return cookies, nil
}
