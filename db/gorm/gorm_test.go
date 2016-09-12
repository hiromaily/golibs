package gorm_test

import (
	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/gorm"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"testing"
)

type MySQL struct {
	*GR
}

var (
	db MySQL
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Gorm]")
}

func setup() {
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
		Id        int
		FirstName string
		LastName  string
	}

	var users []User
	GetMySQLInstance().DB.Raw("SELECT user_id, first_name, last_name FROM t_users WHERE delete_flg=?", "0").Scan(&users)

	lg.Debugf("len(users): %v", len(users))
	lg.Debugf("users[0].FirstName: %v", users[0].FirstName)
}

func TestInsertUser(t *testing.T) {

	type User struct {
		FirstName string `gorm:"column:first_name"`
		LastName  string `gorm:"column:last_name"`
		Email     string `gorm:"column:email"`
		Password  string `gorm:"column:password"`
	}

	GetMySQLInstance().DB.Table("t_users").Save(&User{FirstName: "gormkun", LastName: "sasaki", Email: "abc@db.com", Password: "xxxxx"})
}
