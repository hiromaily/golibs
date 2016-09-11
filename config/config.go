package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	u "github.com/hiromaily/golibs/utils"
	"io/ioutil"
)

var tomlFileName = "./settings.toml"

var conf *Config

// Config is of root
type Config struct {
	Environment int
	Aws         AwsConfig
	MySQL       MySQLConfig
	Redis       RedisConfig
	Mongo       MongoConfig `toml:"mongodb"`
	Mail        MailConfig
}

// AwsConfig is for Aamazon Web Service
type AwsConfig struct {
	AccessKey string    `toml:"access_key"`
	SecretKey string    `toml:"secret_key"`
	Region    string    `toml:"region"`
	Sqs       SqsConfig `toml:"sqs"`
}

// SqsConfig is for SQS of AWS
type SqsConfig struct {
	Endpoint      string        `toml:"endpoint"`
	QueueName     string        `toml:"queue_name"`
	DeadQueueName string        `toml:"deadque_name"`
	MsgAttr       MsgAttrConfig `toml:"msgattr"`
}

// MsgAttrConfig is for part of SQS
type MsgAttrConfig struct {
	OpType      string `toml:"operation_type"`
	ContentType string `toml:"content_type"`
}

// MySQLConfig is for MySQL server
type MySQLConfig struct {
	Host   string `toml:"host"`
	Port   uint16 `toml:"port"`
	DbName string `toml:"dbname"`
	User   string `toml:"user"`
	Pass   string `toml:"pass"`
}

// RedisConfig is for Redis server
type RedisConfig struct {
	Host string `toml:"host"`
	Port uint16 `toml:"port"`
}

// MongoConfig is for MongoDB server
type MongoConfig struct {
	Host   string `toml:"host"`
	Port   uint16 `toml:"port"`
	DbName string `toml:"dbname"`
	User   string `toml:"user"`
	Pass   string `toml:"pass"`
}

// MailConfig is for mail
type MailConfig struct {
	Address  string              `toml:"address"`
	Password string              `toml:"password"`
	Timeout  string              `toml:"timeout"`
	SMTP     SMTPConfig          `toml:"smtp"`
	Content  []MailContentConfig `toml:"content"`
}

// SMTPConfig is for SMTP server of mail
type SMTPConfig struct {
	Server string `toml:"server"`
	Port   int    `toml:"port"`
}

// MailContentConfig is for mail contents
type MailContentConfig struct {
	Subject string `toml:"subject"`
	Tplfile string `toml:"tplfile"`
}

var checkTomlKeys = [][]string{
	{"environment"},
	{"aws", "access_key"},
	{"aws", "secret_key"},
	{"aws", "region"},
	{"aws", "sqs", "endpoint"},
	{"aws", "sqs", "queue_name"},
	{"aws", "sqs", "deadque_name"},
	{"aws", "sqs", "msgattr", "operation_type"},
	{"aws", "sqs", "msgattr", "content_type"},
	{"mysql", "host"},
	{"mysql", "port"},
	{"mysql", "dbname"},
	{"mysql", "user"},
	{"mysql", "pass"},
	{"redis", "host"},
	{"redis", "port"},
	{"redis", "pass"},
	{"mongodb", "host"},
	{"mongodb", "port"},
	{"mongodb", "dbname"},
	{"mongodb", "user"},
	{"mongodb", "pass"},
	{"mail", "address"},
	{"mail", "password"},
	{"mail", "timeout"},
	{"mail", "smtp", "server"},
	{"mail", "smtp", "port"},
	//{"mail", "content", "subject"},
	//{"mail", "content", "tplfile"},
}

//check validation of config
func validateConfig(conf *Config, md *toml.MetaData) error {
	//for protection when debugging on non production environment
	var errStrings []string

	//Check added new items on toml
	// environment
	//if !md.IsDefined("environment") {
	//	errStrings = append(errStrings, "environment")
	//}

	format := "[%s]"
	inValid := false
	for _, keys := range checkTomlKeys {
		if !md.IsDefined(keys...) {
			switch len(keys) {
			case 1:
				format = "[%s]"
			case 2:
				format = "[%s] %s"
			case 3:
				format = "[%s.%s] %s"
			default:
				//invalid check string
				inValid = true
				break
			}
			keysIfc := u.SliceStrToInterface(keys)
			errStrings = append(errStrings, fmt.Sprintf(format, keysIfc...))
		}
	}

	// Error
	if inValid {
		return errors.New("Error: Check Text has wrong number of parameter")
	}
	if len(errStrings) != 0 {
		return fmt.Errorf("Error: There are lacks of keys : %#v \n", errStrings)
	}

	return nil
}

// load configfile
func loadConfig(path string) (*Config, error) {
	if path != "" {
		tomlFileName = path
	}

	d, err := ioutil.ReadFile(tomlFileName)
	if err != nil {
		return nil, fmt.Errorf(
			"Error reading %s: %s", tomlFileName, err)
	}

	var config Config
	md, err := toml.Decode(string(d), &config)
	if err != nil {
		return nil, fmt.Errorf(
			"Error parsing %s: %s(%v)", tomlFileName, err, md)
	}

	//check validation of config
	err = validateConfig(&config, &md)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// New is to create config instance
func New(file string) {
	var err error
	conf, err = loadConfig(file)
	if err != nil {
		panic(err)
	}
}

// GetConf is to get config instance. singleton architecture
func GetConf() *Config {
	var err error
	if conf == nil {
		conf, err = loadConfig("")
	}
	if err != nil {
		panic(err)
	}

	return conf
}

// SetTOMLPath is to set TOML file path
func SetTOMLPath(path string) {
	tomlFileName = path
}

// ResetConf is to clear config instance
func ResetConf() {
	conf = nil
}
