package main

import (
	"path/filepath"

	"github.com/hiromaily/golibs/files"
)

func main() {
	//filelist
	file, _ := filepath.Abs("./memo.txt")
	fileList := []string{file}

	err := files.WatchFile(fileList)
	if err != nil {
		panic(err)
	}
}
