package mails_test

import (
	"flag"
	conf "github.com/hiromaily/golibs/config"
	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/mails"
	tpl "github.com/hiromaily/golibs/tmpl"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
	confFile = flag.String("fp", "", "Config File Path")
)

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[MAIL_TEST]", "/var/log/go/test.log")
	if *benchFlg == 0 {
	}
}

func teardown() {
	if *benchFlg == 0 {
	}
}

// Initialize
func TestMain(m *testing.M) {
	flag.Parse()

	if *confFile == "" {
		os.Exit(1)
		return
	}

	//TODO: According to argument, it switch to user or not.
	//TODO: For bench or not bench
	setup()

	code := m.Run()

	teardown()

	// 終了
	os.Exit(code)
}

//-----------------------------------------------------------------------------
// Mail
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
