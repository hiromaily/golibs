package i18n

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/i18n"
	"os"
)

//https://github.com/nicksnyder/go-i18n

// Translation is translated data
type Translation struct {
	msgID string
	args  []interface{}
}

// Languages is en, nl, de, fr and so on.
var (
	Languages []string
	files     = []string{"en-us.all", "nl-nl.all", "de-de.all", "fr-fr.all"}
)

func init() {
	path := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/i18n/locale/"
	//wd, _ := os.Getwd()
	//if strings.HasSuffix(wd, "test") {
	//	wd = ".."
	//}
	for _, v := range files {
		i18n.MustLoadTranslationFile(fmt.Sprintf("%s%s.json", path, v))
	}
	Languages = i18n.LanguageTags()
}

// MustTfunc is
func MustTfunc(msgID string, args ...interface{}) func(lang string) string {
	return func(lang string) string {
		return i18n.MustTfunc(lang)(msgID, args...)
	}
}

// GetTranslations is to get translated message by ID
func GetTranslations(msgID string, args ...interface{}) map[string]string {
	res := make(map[string]string, len(Languages))
	for _, tag := range Languages {
		res[tag] = MustTfunc(msgID, args...)(tag)
	}
	return res
}

// T is to return Translation struct
func T(msgID string, args ...interface{}) Translation {
	return Translation{msgID, args}
}

// Map is to get map data
func (t Translation) Map() map[string]string {
	return GetTranslations(t.msgID)
}

// String is to get string data
func (t Translation) String(lang string) string {
	return MustTfunc(t.msgID, t.args...)(lang)
}
