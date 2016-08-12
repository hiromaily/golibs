package mysql_test

import (
	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/mysql"
	"github.com/hiromaily/golibs/db/redis"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	u "github.com/hiromaily/golibs/utils"
	"os"
	"testing"
)

type MySQL struct {
	Db *MS
}

//For embeded type
type MySQL2 struct {
	*MS
}

var (
	benchFlg bool = false
	db       MySQL
	db2      MySQL2
	//cahce
	cacheData map[string][]map[string]interface{}
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[MySQL_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		benchFlg = true
	}
}

func setup() {
	if !benchFlg {
		//New("localhost", "hiromaily", "root", "", 3306)
		NewMySQL()

		//Redis
		redis.New("localhost", 6379, "")
		//redis.GetRedisInstance().Connection(0)
	}
}

func teardown() {
	if !benchFlg {
		GetMySQLInstance().Db.Close()
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
	conf.SetTomlPath("../../settings.toml")
	c := conf.GetConfInstance().MySQL

	New(c.Host, c.DbName, c.User, c.Pass, c.Port)

	db = MySQL{}
	db.Db = GetDBInstance()
}

//For embeded type
func NewMySQL2() {
	conf.SetTomlPath("../../settings.toml")
	c := conf.GetConfInstance().MySQL

	New(c.Host, c.DbName, c.User, c.Pass, c.Port)

	db2 = MySQL2{}
	db2.MS = GetDBInstance()
}

//using singleton design pattern
func GetMySQLInstance() *MySQL {
	if db.Db == nil {
		NewMySQL()
	}
	return &db
}

func GetMySQL2Instance() *MySQL2 {
	if db2.MS == nil {
		//db2.DB, err = db2.Connection()
		NewMySQL2()
	}
	return &db2
}

// Get User List
func (ms *MySQL) GetUserList() ([]map[string]interface{}, error) {
	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?"
	data, _, err := ms.Db.SelectSQLAllField(sql, 0)
	if err != nil {
		//lg.Errorf("SQL may be wrong. : %s\n", err.Error())
		return nil, err
	} else if len(data) == 0 {
		//lg.Info("No data.")
		return nil, u.Stoe("No data.")
	}
	return data, nil
}

func (ms *MySQL) GetUserListOnCache() ([]map[string]interface{}, error) {
	var cacheKey string = "sql001"

	//check cache data
	if value, ok := cacheData[cacheKey]; ok {
		return value, nil
	}

	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?"
	data, _, err := ms.Db.SelectSQLAllField(sql, 0)
	if err != nil {
		//lg.Errorf("SQL may be wrong. : %s\n", err.Error())
		return nil, err
	} else if len(data) == 0 {
		//lg.Info("No data.")
		return nil, u.Stoe("No data.")
	}
	//set chache
	cacheData[cacheKey] = data

	return data, nil
}

// Get User List(Using Cache)
// TODO: not yet finished
func (ms *MySQL) GetUserListUsingCache() ([]map[string]interface{}, error) {
	//--------------------------------------
	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?"
	data, _, err := ms.Db.SelectSQLAllField(sql, 0)
	if err != nil {
		//lg.Errorf("SQL may be wrong. : %s\n", err.Error())
		return nil, err
	} else if len(data) == 0 {
		//lg.Info("No data.")
		return nil, u.Stoe("No data.")
	}
	return data, nil
}

//
func (ms *MySQL) GetComplicatedSQL() ([]map[string]interface{}, error) {
	sql := `SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?
`

	data, _, err := ms.Db.SelectSQLAllField(sql, 1)
	if err != nil {
		//lg.Errorf("SQL may be wrong. : %s\n", err.Error())
		return nil, err
	} else if len(data) == 0 {
		//lg.Info("No data.")
		return nil, u.Stoe("No data.")
	}
	return data, nil
}

func (ms *MySQL) GetSimpleSQL1() ([]map[string]interface{}, error) {
	sql := `SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?
`

	data, _, err := ms.Db.SelectSQLAllField(sql, 1)
	if err != nil {
		//lg.Errorf("SQL may be wrong. : %s\n", err.Error())
		return nil, err
	} else if len(data) == 0 {
		//lg.Info("No data.")
		return nil, u.Stoe("No data.")
	}
	return data, nil

}

func (ms *MySQL) GetSimpleSQL2(id int) ([]map[string]interface{}, error) {
	sql := `SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?
`

	data, _, err := ms.Db.SelectSQLAllField(sql, 1)
	if err != nil {
		//lg.Errorf("SQL may be wrong. : %s\n", err.Error())
		return nil, err
	} else if len(data) == 0 {
		//lg.Info("No data.")
		return nil, u.Stoe("No data.")
	}
	return data, nil

}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestGetUserList(t *testing.T) {
	//t.Skip("skipping TestGetUserList")

	data, err := GetMySQLInstance().GetUserList()
	if err != nil {
		t.Fatalf("TestGetUserList: %s", err)
	}

	if u.Itos(data[0]["first_name"]) != "taro" {
		t.Errorf("TestGetUserList result: %#v", data[0])
	}
	t.Logf("data %+v", data[0])
	//TODO:format
	//create_datetime:2016-04-29 21:43:15 +0900 JST
}

func TestSelectSQLAllField(t *testing.T) {
	//t.Skip("skipping TestSelectSQLAllField")

	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?"
	//data, _, err := GetDBInstance().SelectSQLAllField(sql, 0)
	data, _, err := GetMySQLInstance().Db.SelectSQLAllField(sql, 0)

	if err != nil {
		t.Fatalf("SelectSQLAllField: %s", err)
	}

	if u.Itos(data[0]["first_name"]) != "taro" {
		t.Errorf("TestSelectSQLAllField result: %#v", data[0])
	}
	t.Logf("data %+v", data[0])
}

//TODO:work in progress
func TestSelectSQLAllFieldIns(t *testing.T) {
	//t.Skip("skipping TestSelectSQLAllFieldIns")

	type Person struct {
		UserId    int    `db:"id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		//DateTime  time.Time `db:"create_datetime"`
		DateTime string `db:"create_datetime"`
	}

	//db := GetDBInstance()
	db := GetMySQLInstance().Db

	//single data
	var person Person
	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?"
	db.SelectSQLAllFieldIns(sql, "0")
	if db.Err != nil {
		t.Fatalf("TestSelectSQLAllFieldIns: %s", db.Err)
	}
	for {
		//Next
		db.ScanOne(&person)
		if db.Err != nil {
			t.Fatalf("TestSelectSQLAllFieldIns: %s", db.Err)
		} else {
			t.Log(person)
			t.Logf("person.UserId: %d", person.UserId)
			t.Logf("person.FirstName: %s", person.FirstName)
			t.Logf("person.LastName: %s", person.LastName)
			t.Logf("person.DateTime: %s", person.DateTime)
		}
	}
	/*
		db.SelectSQLAllFieldIns(sql, "0").ScanOne(&person)

		if db.Err != nil {
			t.Fatalf("TestSelectSQLAllFieldIns: %s", db.Err)
		} else {
			t.Log(person)
			t.Logf("person.UserId: %d", person.UserId)
			t.Logf("person.FirstName: %s", person.FirstName)
			t.Logf("person.LastName: %s", person.LastName)
		}
		//Next
		db.ScanOne(&person)
		if db.Err != nil {
			t.Fatalf("TestSelectSQLAllFieldIns: %s", db.Err)
		} else {
			t.Log(person)
			t.Logf("person.UserId: %d", person.UserId)
			t.Logf("person.FirstName: %s", person.FirstName)
			t.Logf("person.LastName: %s", person.LastName)
		}
	*/

	//slice
	/*
		var persons []Person
		sql = "SELECT user_id, first_name, last_name FROM t_users WHERE delete_flg=?"
		db.SelectSQLAllFieldIns(sql, 0).Scan(&persons)

		if db.Err != nil {
			t.Fatalf("TestSelectSQLAllFieldIns: %s", db.Err)
		}
	*/

}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
// ConnectionPool VS Not use it
//-----------------------------------------------------------------------------
func BenchmarkConnectionPool(b *testing.B) {
	b.Skip("skipping BenchmarkConnectionPool")

	//BenchmarkConnectionPool-4
	b.ResetTimer()
	NewMySQL()
	GetMySQLInstance().Db.SetMaxIdleConns(100)
	GetMySQLInstance().Db.SetMaxOpenConns(10000)
	for i := 0; i < b.N; i++ {
		//
		_, _ = GetMySQLInstance().GetUserList()
		//
	}
	GetMySQLInstance().Db.Close()
	b.StopTimer()

	//20000000	        87.1 ns/op (93.8 ns/op)
	//ok  	github.com/hiromaily/golibs/db/mysql	1.885s
}

func BenchmarkOpenClose(b *testing.B) {
	b.Skip("skipping BenchmarkOpenClose")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewMySQL()
		//performance for select query
		_, _ = GetMySQLInstance().GetUserList()
		//
		GetMySQLInstance().Db.Close()
	}
	b.StopTimer()
	//100000	     24609 ns/op (19178 ns/op)
	//ok  	github.com/hiromaily/golibs/db/mysql	2.730s
}

//-----------------------------------------------------------------------------
// ComplicatedSQL VS MultiSimpleSQL(Don't use it)
//-----------------------------------------------------------------------------
func BenchmarkComplicatedSQL(b *testing.B) {
	b.Skip("skipping BenchmarkComplicatedSQL")

	b.ResetTimer()
	NewMySQL()
	for i := 0; i < b.N; i++ {
		//
		_, _ = GetMySQLInstance().GetComplicatedSQL()
		//
	}
	GetMySQLInstance().Db.Close()
	b.StopTimer()
	//94.2 ns/op
}

func BenchmarkMultiSimpleSQL(b *testing.B) {
	b.Skip("skipping BenchmarkMultiSimpleSQL")

	b.ResetTimer()
	NewMySQL()
	for i := 0; i < b.N; i++ {
		//
		_, _ = GetMySQLInstance().GetSimpleSQL1()
		_, _ = GetMySQLInstance().GetSimpleSQL2(9)
		_, _ = GetMySQLInstance().GetSimpleSQL2(13)
		_, _ = GetMySQLInstance().GetSimpleSQL2(21)
		//
	}
	GetMySQLInstance().Db.Close()
	b.StopTimer()
	//375 ns/op
}

//-----------------------------------------------------------------------------
// Set map VS Set struct
//-----------------------------------------------------------------------------
//TODO:work in progress
func BenchmarkSetStruct(b *testing.B) {
	b.Skip("skipping BenchmarkSetStruct")

	type Person struct {
		UserId    int    `db:"id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"first_name"`
	}

	b.ResetTimer()
	NewMySQL()
	//db := GetMySQLInstance().Db

	//single data
	var person Person
	sql := "SELECT user_id, first_name, last_name FROM t_users WHERE delete_flg=?"

	for i := 0; i < b.N; i++ {
		//
		GetMySQLInstance().Db.SelectSQLAllFieldIns(sql, "0").ScanOne(&person)
		//db.SelectSQLAllFieldIns(sql, "0")
		//
	}
	GetMySQLInstance().Db.Close()
	b.StopTimer()

	//20000000	        87.1 ns/op (93.8 ns/op)
}

//-----------------------------------------------------------------------------
// NoCache VS Using Redis
//   Use for only heavy query
//-----------------------------------------------------------------------------
func BenchmarkCacheResponse(b *testing.B) {
	b.Skip("skipping BenchmarkCacheResponse")

	b.ResetTimer()
	NewMySQL()
	for i := 0; i < b.N; i++ {
		//
		_, _ = GetMySQLInstance().GetUserListOnCache()
		//
	}
	GetMySQLInstance().Db.Close()
	b.StopTimer()

	//20000000	        87.1 ns/op (93.8 ns/op)
	//ok  	github.com/hiromaily/golibs/db/mysql	1.885s
}

/*
func BenchmarkCacheResponseToRedis(b *testing.B) {
	b.ResetTimer()
	NewMySQL()
	redis.New("localhost", 6379)
	//GetRedisInstance().Connection(0)

	//1. exec sql
	//2. save result using serialized query sentense as key
	//3. check redis using that key
	for i := 0; i < b.N; i++ {
		//
		_, _ = GetMySQLInstance().GetUserList()
		//
	}
	GetMySQLInstance().Db.Close()
	b.StopTimer()
}
*/
