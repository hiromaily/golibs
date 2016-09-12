package mail_test

import (
	"flag"
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
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Mail]")

	if *confFile == "" {
		os.Exit(1)
		return
	}

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
// Test
//-----------------------------------------------------------------------------
func TestMail(t *testing.T) {

	type BodyInfo struct {
		ToName   string
		FromName string
	}

	conf.New(*confFile)
	conf := conf.GetConf().Mail

	//subject
	subject := conf.Content[0].Subject
	lg.Debugf("mail conf: %s", subject)

	//body
	path := os.Getenv("GOPATH") + conf.Content[0].Tplfile
	bodyParam := BodyInfo{ToName: "Hiroki", FromName: "Harry"}
	body, err := tpl.FilePathParser(path, bodyParam)
	if err != nil {
		t.Fatalf("TestMail[01] error: %s", err)
	}

	//mails
	smtp := SMTP{Address: conf.Address, Pass: conf.Password, Server: conf.SMTP.Server, Port: conf.SMTP.Port}
	ml := &Info{ToAddress: []string{"hiromaily@gmail.com"}, FromAddress: conf.Address,
		Subject: subject, Body: body, SMTP: smtp}
	ml.SendMail(conf.Timeout)
}
