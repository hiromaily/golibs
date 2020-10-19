package gorp_test

import (
	"os"
	"testing"

	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/gorp"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
)

//
type MySQL struct {
	*GR
}

var (
	db MySQL
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[Gorp]")
	if !tu.BenchFlg {
		//New("localhost", "hiromaily", "root", "", 3306)
		NewMySQL()
	}
}

func teardown() {
	if !tu.BenchFlg {
		GetMySQLInstance().Close()
	}
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// functions
//-----------------------------------------------------------------------------
func NewMySQL() {
	c := conf.GetConf().MySQL

	New(c.Host, c.DbName, c.User, c.Pass, c.Port)

	db = MySQL{}
	db.GR = GetDB()
}

func GetMySQLInstance() *MySQL {
	if db.GR == nil {
		NewMySQL()
	}
	return &db
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestSelectUser(t *testing.T) {

	type User struct {
		ID        int    `db:"user_id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
	}

	rows, _ := GetMySQLInstance().DB.Select(&User{}, "SELECT user_id, first_name, last_name FROM t_users WHERE delete_flg=?", "0")
	for _, row := range rows {
		user := *row.(*User)
		lg.Debugf("%d, %s %s\n", user.ID, user.FirstName, user.LastName)
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
