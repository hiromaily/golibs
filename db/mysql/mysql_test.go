package mysql_test

import (
	"fmt"
	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/mysql"
	"github.com/hiromaily/golibs/db/redis"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	r "github.com/hiromaily/golibs/runtimes"
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
	conf.SetTOMLPath("../../settings.toml")
	c := conf.GetConf().MySQL

	New(c.Host, c.DbName, c.User, c.Pass, c.Port)

	db = MySQL{}
	db.Db = GetDB()
}

//For embeded type
func NewMySQL2() {
	conf.SetTOMLPath("../../settings.toml")
	c := conf.GetConf().MySQL

	New(c.Host, c.DbName, c.User, c.Pass, c.Port)

	db2 = MySQL2{}
	db2.MS = GetDB()
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
func (ms *MySQL) getUserList() ([]map[string]interface{}, error) {
	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?"
	data, _, err := ms.Db.Select(sql, 0)
	if err != nil {
		return nil, err
	} else if len(data) == 0 {
		//lg.Info("No data.")
		return nil, u.Stoe("No data.")
	}
	return data, nil
}

// Get User List(Using Cache)
func (ms *MySQL) getUserListOnCache() ([]map[string]interface{}, error) {
	var cacheKey string = "sql001"

	//check cache data
	if value, ok := cacheData[cacheKey]; ok {
		return value, nil
	}

	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?"
	data, _, err := ms.Db.Select(sql, 0)
	if err != nil {
		return nil, err
	} else if len(data) == 0 {
		return nil, u.Stoe("No data.")
	}
	//set chache
	cacheData[cacheKey] = data

	return data, nil
}

// For comparison of multiple simple queries and one complicated heavy query in performance
func (ms *MySQL) getComplicatedSQL() ([]map[string]interface{}, error) {
	sql := `SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?`

	data, _, err := ms.Db.Select(sql, 1)
	if err != nil {
		return nil, err
	} else if len(data) == 0 {
		return nil, u.Stoe("No data.")
	}
	return data, nil
}

func (ms *MySQL) getSimpleSQL1() ([]map[string]interface{}, error) {
	sql := `SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?`

	data, _, err := ms.Db.Select(sql, 1)
	if err != nil {
		return nil, err
	} else if len(data) == 0 {
		return nil, u.Stoe("No data.")
	}
	return data, nil
}

func (ms *MySQL) getSimpleSQL2(id int) ([]map[string]interface{}, error) {
	sql := `SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?`

	data, _, err := ms.Db.Select(sql, 1)
	if err != nil {
		return nil, err
	} else if len(data) == 0 {
		return nil, u.Stoe("No data.")
	}
	return data, nil
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestGetUserList(t *testing.T) {
	//t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))

	data, err := GetMySQLInstance().getUserList()
	if err != nil {
		t.Fatalf("GetMySQLInstance().getUserList() error: %s", err)
	}

	if u.Itos(data[0]["first_name"]) != "harry" {
		t.Errorf(" GetMySQLInstance().getUserList() result: %#v", data[0])
	} else {
		t.Logf("result: %+v", data[0])
	}
}

func TestSelect(t *testing.T) {
	//t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))

	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?"
	data, _, err := GetMySQLInstance().Db.Select(sql, 0)
	if err != nil {
		t.Fatalf("GetMySQLInstance().Db.Select(sql, 0) error: %s", err)
	}

	if u.Itos(data[0]["first_name"]) != "harry" {
		t.Errorf("GetMySQLInstance().Db.Select(sql, 0) result: %#v", data[0])
	} else {
		t.Logf("result: %+v", data[0])
	}
}

func TestSelectInsScanOne(t *testing.T) {
	//t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))

	type Person struct {
		UserId    int    `db:"id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		//DateTime  time.Time `db:"create_datetime"`
		DateTime string `db:"create_datetime"`
	}

	db := GetMySQLInstance().Db

	//1.Single data
	var person Person
	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?"
	db.SelectIns(sql, "0")
	if db.Err != nil {
		t.Fatalf("[1]db.SelectIns(): %s", db.Err)
	}

	for db.ScanOne(&person) {
		//t.Log(person)
		//t.Logf("person.UserId: %d", person.UserId)
		//t.Logf("person.FirstName: %s", person.FirstName)
		//t.Logf("person.LastName: %s", person.LastName)
		//t.Logf("person.DateTime: %s", person.DateTime)
	}
	if db.Err != nil {
		t.Fatalf("[1]db.ScanOne(): %s", db.Err)
	}

	//When result is nothing -> nodata
	var person2 Person
	b := db.SelectIns(sql, "1").ScanOne(&person2)
	if db.Err != nil {
		t.Fatalf("[2]db.SelectIns(): %s", db.Err)
	}
	if b {
		t.Error("[2]db.ScanOne(): return bool may be wrong.")
	}
}

func TestSelectInsScan(t *testing.T) {
	//t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))

	type Person struct {
		UserId    int    `db:"id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		//DateTime  time.Time `db:"create_datetime"`
		DateTime string `db:"create_datetime"`
	}

	//db := GetDBInstance()
	db := GetMySQLInstance().Db

	//2. Get All Data
	var persons []Person
	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?"
	db.SelectIns(sql, 0).Scan(&persons)

	if db.Err != nil {
		t.Errorf("[1]db.SelectIns(sql, 0).Scan(): %s", db.Err)
	}

	//When result is nothing -> len(xxx)==0
	var persons2 []Person
	b := db.SelectIns(sql, 1).Scan(&persons2)
	if db.Err != nil {
		t.Errorf("[2]db.SelectIns(sql, 1).Scan(): %s", db.Err)
	}
	//if len(persons2) != 0 {
	if b {
		t.Errorf("[2]db.SelectIns(sql, 1).Scan(): number of result is %d", len(persons2))
	}
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
// ConnectionPool VS Not use it
//-----------------------------------------------------------------------------
func BenchmarkConnectionPool(b *testing.B) {
	b.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))

	//BenchmarkConnectionPool-4
	b.ResetTimer()
	NewMySQL()
	GetMySQLInstance().Db.SetMaxIdleConns(100)
	GetMySQLInstance().Db.SetMaxOpenConns(10000)
	for i := 0; i < b.N; i++ {
		//
		_, _ = GetMySQLInstance().getUserList()
		//
	}
	GetMySQLInstance().Db.Close()
	b.StopTimer()

	//20000000	        87.1 ns/op (93.8 ns/op)
	//ok  	github.com/hiromaily/golibs/db/mysql	1.885s
}

func BenchmarkOpenClose(b *testing.B) {
	b.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewMySQL()
		//performance for select query
		_, _ = GetMySQLInstance().getUserList()
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
	b.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))

	b.ResetTimer()
	NewMySQL()
	for i := 0; i < b.N; i++ {
		//
		_, _ = GetMySQLInstance().getComplicatedSQL()
		//
	}
	GetMySQLInstance().Db.Close()
	b.StopTimer()
	//94.2 ns/op
}

func BenchmarkMultiSimpleSQL(b *testing.B) {
	b.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))

	b.ResetTimer()
	NewMySQL()
	for i := 0; i < b.N; i++ {
		//
		_, _ = GetMySQLInstance().getSimpleSQL1()
		_, _ = GetMySQLInstance().getSimpleSQL2(9)
		_, _ = GetMySQLInstance().getSimpleSQL2(13)
		_, _ = GetMySQLInstance().getSimpleSQL2(21)
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
	b.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))

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
		GetMySQLInstance().Db.SelectIns(sql, "0").ScanOne(&person)
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
	b.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))

	b.ResetTimer()
	NewMySQL()
	for i := 0; i < b.N; i++ {
		//
		_, _ = GetMySQLInstance().getUserListOnCache()
		//
	}
	GetMySQLInstance().Db.Close()
	b.StopTimer()

	//20000000	        87.1 ns/op (93.8 ns/op)
	//ok  	github.com/hiromaily/golibs/db/mysql	1.885s
}
