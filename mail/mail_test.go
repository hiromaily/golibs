package mail_test

import (
	"flag"
	enc "github.com/hiromaily/golibs/cipher/encryption"
	conf "github.com/hiromaily/golibs/config"
	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/mail"
	tu "github.com/hiromaily/golibs/testutil"
	tpl "github.com/hiromaily/golibs/tmpl"
	"os"
	"testing"
)

var (
	confFile = flag.String("fp", "", "Config File Path")
	mailTo   = "FcQjhb5ErsSBlh6EDwe69eLdcW/eJPKWtnTWmDPoAAM=" //encrypted mail address
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Mail]")

	//crypt
	enc.NewCryptDefault()

	crypt := enc.GetCrypt()
	mailTo, _ = crypt.DecryptBase64(mailTo)

	//conf
	if *confFile == "" {
		*confFile = os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/settings.toml"
	}
	conf.New(*confFile, true)
}

func setup() {
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// functions
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestMail(t *testing.T) {

	type BodyInfo struct {
		ToName   string
		FromName string
	}

	conf := conf.GetConf().Mail

	//subject
	subject := conf.Content[0].Subject
	lg.Debugf("mail conf: %s", subject)

	//body
	path := os.Getenv("GOPATH") + conf.Content[0].Tplfile
	bodyParam := BodyInfo{ToName: "Hiroki", FromName: "Harry"}
	body, err := tpl.FilePathParser(path, bodyParam)
	if err != nil {
		t.Fatalf("FilePathParser error: %s", err)
	}

	//mails
	smtp := SMTP{Address: conf.Address, Pass: conf.Password, Server: conf.SMTP.Server, Port: conf.SMTP.Port}
	ml := &Info{ToAddress: []string{mailTo}, FromAddress: conf.Address,
		Subject: subject, Body: body, SMTP: smtp}
	ml.SendMail(conf.Timeout)
}
