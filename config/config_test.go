package config_test

import (
	"flag"
	. "github.com/hiromaily/golibs/config"
	//lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"testing"
)

type User struct {
	Id   int
	Name string
}

var (
	confFile = flag.String("fp", "", "Config File Path")
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Config]")

	if *confFile == "" {
		*confFile = os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/config/settings.toml"
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
	//tu.SkipLog(t)

	tomlPath := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/config/settings.toml"

	//New(*confFile, false)
	SetTOMLPath(tomlPath)
	conf := GetConf()

	t.Logf("conf.Environment: %v", conf.Environment)
	t.Logf("conf.Aws.AccessKey: %v", conf.Aws.AccessKey)
	t.Logf("conf.Mail.Address: %v", conf.Mail.Address)
	t.Logf("conf.Mail.Smtp.Server: %v", conf.Mail.SMTP.Server)
	t.Logf("conf.Mail.Content[0].Subject: %v", conf.Mail.Content[0].Subject)
	t.Logf("conf.MySQL.Host: %v", conf.MySQL.Host)

	ResetConf()
}

func TestConfig2(t *testing.T) {
	tomlPath := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/config/settings.default.toml"

	SetTOMLPath(tomlPath)
	conf := GetConf()

	t.Logf("conf.Environment: %v", conf.Environment)
	t.Logf("conf.Aws.AccessKey: %v", conf.Aws.AccessKey)
	t.Logf("conf.Mail.Address: %v", conf.Mail.Address)
	t.Logf("conf.Mail.Smtp.Server: %v", conf.Mail.SMTP.Server)
	t.Logf("conf.Mail.Content[0].Subject: %v", conf.Mail.Content[0].Subject)
	t.Logf("conf.MySQL.Host: %v", conf.MySQL.Host)

	ResetConf()
}

func TestConfig3(t *testing.T) {
	tomlPath := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/config/tavis.toml"

	SetTOMLPath(tomlPath)
	conf := GetConf()

	t.Logf("conf.Environment: %v", conf.Environment)
	t.Logf("conf.Aws.AccessKey: %v", conf.Aws.AccessKey)
	t.Logf("conf.Mail.Address: %v", conf.Mail.Address)
	t.Logf("conf.Mail.Smtp.Server: %v", conf.Mail.SMTP.Server)
	t.Logf("conf.Mail.Content[0].Subject: %v", conf.Mail.Content[0].Subject)
	t.Logf("conf.MySQL.Host: %v", conf.MySQL.Host)

	ResetConf()
}
