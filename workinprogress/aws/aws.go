package aws

import (
	"os"
)

// InitAwsEnv is to set access key and secret key to environment variable
func InitAwsEnv(accessKey string, secretkey string) {
	os.Setenv("AWS_ACCESS_KEY_ID", accessKey)
	os.Setenv("AWS_SECRET_ACCESS_KEY", secretkey)
}
