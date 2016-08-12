package mails_test

import (
	"flag"
	conf "github.com/hiromaily/golibs/config"
	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/mails"
	o "github.com/hiromaily/golibs/os"
	tpl "github.com/hiromaily/golibs/tmpl"
	"os"
	"testing"
)

var (
	confFile      = flag.String("fp", "", "Config File Path")
	benchFlg bool = false
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	flag.Parse()

	if *confFile == "" {
		os.Exit(1)
		return
	}

	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[MAIL_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		benchFlg = true
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
	conf := conf.GetConfInstance().Mail

	//subject
	subject := conf.Content[0].Subject
	t.Logf("mail conf: %s", subject)

	//body
	path := os.Getenv("GOPATH") + conf.Content[0].Tplfile
	bodyParam := BodyInfo{ToName: "Hiroki", FromName: "Harry"}
	body, err := tpl.FilePathParser(path, bodyParam)
	if err != nil {
		t.Fatalf("TestMail[01] error: %s", err)
	}

	//mails
	smtp := Smtp{Address: conf.Address, Pass: conf.Password, Server: conf.Smtp.Server, Port: conf.Smtp.Port}
	ml := &MailInfo{ToAddress: []string{"hiromaily@gmail.com"}, FromAddress: conf.Address,
		Subject: subject, Body: body, Smtp: smtp}
	ml.SendMail(conf.Timeout)
}
