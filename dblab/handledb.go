package main

import (
	"flag"
	"fmt"
	"github.com/hiromaily/golibs/config"
	"github.com/hiromaily/golibs/dblab/mysql"
	logex "github.com/hiromaily/golibs/log"
	"log"
	"os"
	"runtime"
)

//command line
var (
	mode  = flag.Int("mode", 1, "")
	toml  = flag.String("toml", "", "toml file path")
	debug = flag.Int("debug", 1, "")
)
var usage = `Usage: boom [options...] <url>

Options:
  -mode       1:create, 2:purge
  -toml       tomlファイルのパス
  -debug      1:show something logs.

e.g. for how to use
  // mode1 is for specific sqs data
  $ ./handledb -mode 1 -debug 1
`

// handling command line
func handleCmdline() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, runtime.NumCPU()))
	}

	flag.Parse()

	logex.Debug("flag.NArg(): ", flag.NArg())
	logex.Debug("mode: ", *mode)
	logex.Debug("toml: ", *toml)
	logex.Debug("debug: ", *debug)
}

//init
func init() {
	logex.InitLog(2, log.Ltime, "") //From debug log

	//handle command line
	handleCmdline()
}

//basic hot to use mysql default packages
func basicDB() {
	dbInfo := mysql.GetDBInstance()

	//dbInfo.InesrtSQL()
	//dbInfo.UpdateSQL()
	//dbInfo.SelectSQL()
	//dbInfo.SelectSQLAllField()
	//dbInfo.SelectOneRowSQL()
	//dbInfo.SelectOneRowSQL2()
	//dbInfo.SelectCount()

	dbInfo.SelectGetFieldData()

	dbInfo.Close()
}

func genmaiDB() {
	dbInfo := mysql.GetDBGenmaiInstance()

	insertSQL := "INSERT t_users SET first_name=?, last_name=?"
	args := []interface{}{"hita", "asakura"}
	dbInfo.InesrtSQLGenmai(insertSQL, args...)

	dbInfo.CloseGenmai()
}

func gorpDB() {
	dbInfo := mysql.GetDBGorpInstance()

	//insertSQL := "INSERT t_users SET first_name='ren', last_name='yamaha'"
	//dbInfo.InesrtSQLGorp(insertSQL)

	dbInfo.InesrtSQLFromStructGorp()

	//dbInfo.GetRecord()

	//selectSQL := "SELECT * FROM t_users"
	selectSQL := "SELECT user_id, first_name, last_name FROM t_users"
	dbInfo.SelectSQLGorp(selectSQL)

}

// main
func main() {

	//get toml path
	//fmt.Println("toml file is ", *toml)
	if *toml != "" {
		config.SetTomlPath(*toml)
	}
	//config.GetConfInstance()

	if *mode == 1 {
		basicDB()
	} else if *mode == 2 {
		genmaiDB()
	} else if *mode == 3 {
		gorpDB()
	}
}
