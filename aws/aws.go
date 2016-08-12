package aws

import (
	"os"
)

func InitAwsEnv(accessKey string, secretkey string) {
	os.Setenv("AWS_ACCESS_KEY_ID", accessKey)
	os.Setenv("AWS_SECRET_ACCESS_KEY", secretkey)
}
