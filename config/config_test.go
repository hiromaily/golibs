package config_test

import (
	"flag"
	. "github.com/hiromaily/golibs/config"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	"os"
	"testing"
)

type User struct {
	Id   int
	Name string
}

var (
	confFile      = flag.String("fp", "", "Config File Path")
	benchFlg bool = false
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	//Here is [slower] than included file's init()
	flag.Parse()

	if *confFile == "" {
		os.Exit(1)
		return
	}

	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[CONFIG_TEST]", "/var/log/go/test.log")
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
func TestConfig(t *testing.T) {
	//t.Skip("skipping TestEncodeStruct")
	//*

	New(*confFile)
	conf := GetConf()

	t.Logf("conf.Environment: %v", conf.Environment)
	t.Logf("conf.Aws.AccessKey: %v", conf.Aws.AccessKey)
	t.Logf("conf.Mail.Address: %v", conf.Mail.Address)
	t.Logf("conf.Mail.Smtp.Server: %v", conf.Mail.Smtp.Server)
	t.Logf("conf.Mail.Content[0].Subject: %v", conf.Mail.Content[0].Subject)
	t.Logf("conf.MySQL.Host: %v", conf.MySQL.Host)

	ResetConf()
}

func TestConfig2(t *testing.T) {
	//${GOPATH}/src/github.com/hiromaily/golibs/settings.toml
	tomlPath := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/settings.default.toml"

	SetTomlPath(tomlPath)
	conf := GetConf()

	t.Logf("conf.Environment: %v", conf.Environment)
	t.Logf("conf.Aws.AccessKey: %v", conf.Aws.AccessKey)
	t.Logf("conf.Mail.Address: %v", conf.Mail.Address)
	t.Logf("conf.Mail.Smtp.Server: %v", conf.Mail.Smtp.Server)
	t.Logf("conf.Mail.Content[0].Subject: %v", conf.Mail.Content[0].Subject)
	t.Logf("conf.MySQL.Host: %v", conf.MySQL.Host)
}
