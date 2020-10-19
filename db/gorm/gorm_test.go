package gorm_test

import (
	"os"
	"testing"

	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/gorm"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
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

func setup() {
	tu.InitializeTest("[Gorm]")
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
func NewMySQL() error {
	c := conf.GetConf().MySQL

	err := New(c.Host, c.DbName, c.User, c.Pass, c.Port)
	if err != nil {
		return err
	}

	db = MySQL{}
	db.GR = GetDB()

	return nil
}

func GetMySQLInstance() *MySQL {
	if db.GR == nil {
		err := NewMySQL()
		if err != nil {
			return nil
		}
	}
	return &db
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestSelectUser(t *testing.T) {

	type User struct {
		ID        int
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
