package os

import (
	"os"

	reg "github.com/hiromaily/golibs/regexp"
)

// GetArgs is to get args
func GetArgs(i int) string {
	return os.Args[i]
}

// AddParam is to add value as arguments
func AddParam(val string) {
	os.Args = append(os.Args, val)
}

// FindParam is to find specific args
func FindParam(key string) (bRet bool) {
	//fmt.Println(os.Args)
	bRet = false
	for _, v := range os.Args {
		if reg.CheckRegexp(`^`+key, v) {
			bRet = true
			break
		}
	}
	return
}

// IsFileExisted check if file exists
func IsFileExisted(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == os.ErrNotExist {
		return false
	} else if err != nil {
		// error may be better to return
		return false
	}
	return true
}
