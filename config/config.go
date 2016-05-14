package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

var tomlFileName string = "./settings.toml"

var conf *Config

type Config struct {
	Environment int
	Aws         AwsConfig
	Database    DatabaseConfig
	Mail        MailConfig
}

type AwsConfig struct {
	AccessKey string    `toml:"access_key"`
	SecretKey string    `toml:"secret_key"`
	Region    string    `toml:"region"`
	Sqs       SqsConfig `toml:"sqs"`
}

type SqsConfig struct {
	Endpoint      string        `toml:"endpoint"`
	QueueName     string        `toml:"queue_name"`
	DeadQueueName string        `toml:"deadque_name"`
	MsgAttr       MsgAttrConfig `toml:"msgattr"`
}

type MsgAttrConfig struct {
	OpType      string `toml:"operation_type"`
	ContentType string `toml:"content_type"`
}

type DatabaseConfig struct {
	User   string `toml:"user"`
	Pass   string `toml:"pass"`
	DbName string `toml:"dbname"`
}

type MailConfig struct {
	Address   string              `toml:"address"`
	Password  string              `toml:"password"`
	Timeout   string              `toml:"timeout"`
	Smtp      SmtpConfig          `toml:"smtp"`
	Content   []MailContentConfig `toml:"content"`
}

type SmtpConfig struct {
	Server  string  `toml:"server"`
	Port    int     `toml:"port"`
}

type MailContentConfig struct {
	Subject  string  `toml:"subject"`
	Tplfile  string  `toml:"tplfile"`
}


//check validation of config
func validateConfig(conf *Config, md *toml.MetaData) error {
	//for protection when debugging on non production environment
	var errStrings []string

	//Check added new items on toml
	if !md.IsDefined("environment") {
		errStrings = append(errStrings, "environment")
	}

	if !md.IsDefined("database", "user") {
		errStrings = append(errStrings, "[database] user")
	}

	if len(errStrings) != 0 {
		return fmt.Errorf("Error  There are lack of keys : %#v \n", errStrings)
	}

	return nil
}

// load configfile
func loadConfig() (*Config, error) {
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

// singleton architecture
func GetConfInstance() *Config {
	var err error
	if conf == nil {
		conf, err = loadConfig()
	}
	if err != nil {
		panic(err)
	}

	return conf
}

func SetTomlPath(path string) {
	tomlFileName = path
}
