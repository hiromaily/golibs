package gorp_test

import (
	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/gorp"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	"os"
	"testing"
)

//
type MySQL struct {
	*GR
}

var (
	benchFlg bool = false
	db       MySQL
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[GORP_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		benchFlg = true
	}
}

func setup() {
	if !benchFlg {
		//New("localhost", "hiromaily", "root", "", 3306)
		NewMySQL()
	}
}

func teardown() {
	if !benchFlg {
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
	conf.SetTOMLPath("../../settings.toml")
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
