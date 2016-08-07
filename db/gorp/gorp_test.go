package gorp_test

import (
	"flag"
	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/gorp"
	lg "github.com/hiromaily/golibs/log"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

type MySQL struct {
	*GR
}

var db MySQL

func NewMySQL() {
	conf.SetTomlPath("../../settings.toml")
	c := conf.GetConfInstance().MySQL

	New(c.Host, c.DbName, c.User, c.Pass, c.Port)

	db = MySQL{}
	db.GR = GetDBInstance()
}

func GetMySQLInstance() *MySQL {
	if db.GR == nil {
		NewMySQL()
	}
	return &db
}

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[GORP_TEST]", "/var/log/go/test.log")
	if *benchFlg == 0 {
		//New("localhost", "hiromaily", "root", "", 3306)
		NewMySQL()
	}
}

func teardown() {
	if *benchFlg == 0 {
		GetMySQLInstance().Close()
	}
}

// Initialize
func TestMain(m *testing.M) {
	flag.Parse()

	//TODO: According to argument, it switch to user or not.
	//TODO: For bench or not bench
	setup()

	code := m.Run()

	teardown()

	// 終了
	os.Exit(code)
}

func TestSelectUser(t *testing.T) {

	type User struct {
		Id        int    `db:"user_id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
	}

	rows, _ := GetMySQLInstance().DB.Select(&User{}, "SELECT user_id, first_name, last_name FROM t_users WHERE delete_flg=?", "0")
	for _, row := range rows {
		user := *row.(*User)
		t.Logf("%d, %s %s\n", user.Id, user.FirstName, user.LastName)
	}
}

//TODO:work in progress
func TestInsertUser(t *testing.T) {

	type User struct {
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		Email     string `db:"email"`
		Password  string `db:"password"`
	}
	db := GetMySQLInstance().DB
	_ = db.AddTableWithName(User{}, "t_users")
	GetMySQLInstance().DB.Insert(&User{FirstName: "gorp", LastName: "mysql", Email: "123@gg.com", Password: "zzzzz"})
}
