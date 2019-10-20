package errors

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// RepoName is repository name, but should be dynamic
const RepoName = "hiromaiy"

// GetErrorStack is to return formatted error stack as string
func GetErrorStack(e error) string {
	goPath := fmt.Sprintf("%s/src/", os.Getenv("GOPATH"))
	//get directory names
	dirNames, err := getDirNames(goPath)
	if err != nil {
		return ""
	}

	errs := strings.Split(fmt.Sprintf("%+v", e), "\n")
	//format
	//message is done after dir name is found
	var (
		msgIsDone bool
		stacks    string
		idx       int
	)
	for i := 0; i < len(errs); i++ {
		if !msgIsDone {
			if isFound(errs[i], dirNames) || errs[i] == "main.main" {
				//handle as dir name
				msgIsDone = true
			} else {
				//handle as message
				continue
			}
		}
		//stack trace
		//1.funcName
		paths := strings.Split(errs[i], "/")
		if len(paths) == 1 {
			stacks += fmt.Sprintf("\n[%d][func]%s(): ", idx, paths[0])
		} else {
			stacks += fmt.Sprintf("\n[%d][func]%s/%s(): ", idx, paths[len(paths)-2], paths[len(paths)-1])
		}
		//2.source code
		i++
		if strings.Contains(errs[i], RepoName) {
			tmp := strings.Split(errs[i], RepoName)
			stacks += fmt.Sprintf("[file].%s", tmp[1])
		} else {
			tmp := strings.Split(errs[i], "/")
			stacks += fmt.Sprintf("[file].%s/%s", tmp[len(tmp)-2], tmp[len(tmp)-1])
			break
		}
		if errs[i-1] == "main.main" {
			break
		}
		idx++
	}
	return stacks
}

// ErrorStack is to return formatted error stack as string
func ErrorStack(err error) string {
	stacks := GetErrorStack(err)
	return fmt.Sprintf("%s%s", err.Error(), stacks)
}

func getDirNames(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, file.Name())
			continue
		}
	}
	return paths, nil
}

func isFound(targetStr string, dirNames []string) bool {
	for _, dir := range dirNames {
		if strings.Contains(targetStr, dir) {
			return true
		}
	}
	return false
}
