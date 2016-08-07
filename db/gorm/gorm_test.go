package gorm_test

import (
	"flag"
	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/gorm"
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
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[GORM_TEST]", "/var/log/go/test.log")
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
		Id        int
		FirstName string
		LastName  string
	}

	var users []User
	GetMySQLInstance().DB.Raw("SELECT user_id, first_name, last_name FROM t_users WHERE delete_flg=?", "0").Scan(&users)

	t.Logf("len(users): %v", len(users))
	t.Logf("users[0].FirstName: %v", users[0].FirstName)
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
