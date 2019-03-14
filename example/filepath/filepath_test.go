package filepath_test

import (
	"os"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	//join
	t.Logf("filepath.Join(): %s", filepath.Join(os.TempDir(), "sample.txt"))

	//split
	dir, name := filepath.Split(os.Getenv("GOPATH"))
	t.Logf("parent dir: %s, last dir: %s", dir, name)

	//separator
	t.Logf("filepath.Separator: %s", string(filepath.Separator))

	//current dir
	cur, _ := os.Getwd()
	t.Logf("current dir: %s", cur)

	//abs
	//実行パスを追加してるだけ。。。
	absPath, _ := filepath.Abs("./filepath_test.go")
	t.Logf("abs path: %s", absPath)

	// absPath, _ = filepath.Abs("./config/settings.toml")
	// t.Logf("abs path for toml: %s", absPath)

	//base
	base := filepath.Base(absPath)
	t.Logf("filepath.Base(%s): %s", absPath, base)

	//dir
	dir = filepath.Dir(absPath)
	t.Logf("filepath.Dir(%s): %s", absPath, dir)

	//ext
	ext := filepath.Ext(absPath)
	t.Logf("filepath.Ext(%s): %s", absPath, ext)

	//clean
	t.Logf("filepath.Clean(./path/something1/../file.go): %s", filepath.Clean("./path/something1/../file.go"))
}
