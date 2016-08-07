package config_test

import (
	"flag"
	. "github.com/hiromaily/golibs/config"
	lg "github.com/hiromaily/golibs/log"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
	confFile = flag.String("fp", "", "Config File Path")
)

type User struct {
	Id   int
	Name string
}

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[CONFIG_TEST]", "/var/log/go/test.log")
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
// Compress
//-----------------------------------------------------------------------------
func TestConfig(t *testing.T) {
	//t.Skip("skipping TestEncodeStruct")
	//*

	New(*confFile)
	conf := GetConfInstance()

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
	conf := GetConfInstance()

	t.Logf("conf.Environment: %v", conf.Environment)
	t.Logf("conf.Aws.AccessKey: %v", conf.Aws.AccessKey)
	t.Logf("conf.Mail.Address: %v", conf.Mail.Address)
	t.Logf("conf.Mail.Smtp.Server: %v", conf.Mail.Smtp.Server)
	t.Logf("conf.Mail.Content[0].Subject: %v", conf.Mail.Content[0].Subject)
	t.Logf("conf.MySQL.Host: %v", conf.MySQL.Host)
}
