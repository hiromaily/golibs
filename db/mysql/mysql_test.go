package mysql_test

import (
	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/mysql"
	"github.com/hiromaily/golibs/db/redis"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	u "github.com/hiromaily/golibs/utils"
	"os"
	"testing"
)

type MySQL struct {
	Db *MS
}

//For embedded type
type MySQL2 struct {
	*MS
}

var (
	db  MySQL
	db2 MySQL2
	//cahce
	cacheData map[string][]map[string]interface{}
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[MySQL]")
}

func setup() {
	if !tu.BenchFlg {
		newMySQL()

		//Redis
		newRedis()
	}
}

func teardown() {
	if !tu.BenchFlg {
		getMySQL().Db.Close()
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
func newMySQL() {
	c := conf.GetConf().MySQL

	New(c.Host, c.DbName, c.User, c.Pass, c.Port)

	db = MySQL{}
	db.Db = GetDB()
}

//For embedded type
func newMySQL2() {
	conf.SetTOMLPath("../../settings.toml")
	c := conf.GetConf().MySQL

	New(c.Host, c.DbName, c.User, c.Pass, c.Port)

	db2 = MySQL2{}
	db2.MS = GetDB()
}

func newRedis() {
	r := conf.GetConf().Redis

	redis.New(r.Host, r.Port, r.Pass, 0)
	//redis.GetRedisInstance().Connection(0)
}

//using singleton design pattern
func getMySQL() *MySQL {
	if db.Db == nil {
		newMySQL()
	}
	return &db
}

func getMySQL2() *MySQL2 {
	if db2.MS == nil {
		//db2.DB, err = db2.Connection()
		newMySQL2()
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
	//tu.SkipLog(t)

	data, err := getMySQL().getUserList()
	if err != nil {
		t.Fatalf("getMySQL().getUserList() error: %s", err)
	}

	if u.Itos(data[0]["first_name"]) != "harry" {
		t.Errorf(" getMySQL().getUserList() result: %#v", data[0])
	} else {
		lg.Debugf("result: %+v", data[0])
	}
}

func TestSelect(t *testing.T) {
	//tu.SkipLog(t)

	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?"
	data, _, err := getMySQL().Db.Select(sql, 0)
	if err != nil {
		t.Fatalf("getMySQL().Db.Select(sql, 0) error: %s", err)
	}

	if u.Itos(data[0]["first_name"]) != "harry" {
		t.Errorf("getMySQL().Db.Select(sql, 0) result: %#v", data[0])
	} else {
		lg.Debugf("result: %+v", data[0])
	}
}

func TestSelectInsScanOne(t *testing.T) {
	//tu.SkipLog(t)

	type Person struct {
		UserId    int    `db:"id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		//DateTime  time.Time `db:"create_datetime"`
		DateTime string `db:"create_datetime"`
	}

	db := getMySQL().Db

	//1.Single data
	var person Person
	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?"
	db.SelectIns(sql, "0")
	if db.Err != nil {
		t.Fatalf("[1]db.SelectIns(): %s", db.Err)
	}

	for db.ScanOne(&person) {
		//lg.Debugf(person)
		//lg.Debugf("person.UserId: %d", person.UserId)
		//lg.Debugf("person.FirstName: %s", person.FirstName)
		//lg.Debugf("person.LastName: %s", person.LastName)
		//lg.Debugf("person.DateTime: %s", person.DateTime)
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
	//tu.SkipLog(t)

	type Person struct {
		UserId    int    `db:"id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		//DateTime  time.Time `db:"create_datetime"`
		DateTime string `db:"create_datetime"`
	}

	//db := GetDBInstance()
	db := getMySQL().Db

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
	tu.SkipBLog(b)

	//BenchmarkConnectionPool-4
	b.ResetTimer()
	newMySQL()
	getMySQL().Db.SetMaxIdleConns(100)
	getMySQL().Db.SetMaxOpenConns(10000)
	for i := 0; i < b.N; i++ {
		//
		_, _ = getMySQL().getUserList()
		//
	}
	getMySQL().Db.Close()
	b.StopTimer()

	//20000000	        87.1 ns/op (93.8 ns/op)
	//ok  	github.com/hiromaily/golibs/db/mysql	1.885s
}

func BenchmarkOpenClose(b *testing.B) {
	tu.SkipBLog(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newMySQL()
		//performance for select query
		_, _ = getMySQL().getUserList()
		//
		getMySQL().Db.Close()
	}
	b.StopTimer()
	//100000	     24609 ns/op (19178 ns/op)
	//ok  	github.com/hiromaily/golibs/db/mysql	2.730s
}

//-----------------------------------------------------------------------------
// ComplicatedSQL VS MultiSimpleSQL(Don't use it)
//-----------------------------------------------------------------------------
func BenchmarkComplicatedSQL(b *testing.B) {
	tu.SkipBLog(b)

	b.ResetTimer()
	newMySQL()
	for i := 0; i < b.N; i++ {
		//
		_, _ = getMySQL().getComplicatedSQL()
		//
	}
	getMySQL().Db.Close()
	b.StopTimer()
	//94.2 ns/op
}

func BenchmarkMultiSimpleSQL(b *testing.B) {
	tu.SkipBLog(b)

	b.ResetTimer()
	newMySQL()
	for i := 0; i < b.N; i++ {
		//
		_, _ = getMySQL().getSimpleSQL1()
		_, _ = getMySQL().getSimpleSQL2(9)
		_, _ = getMySQL().getSimpleSQL2(13)
		_, _ = getMySQL().getSimpleSQL2(21)
		//
	}
	getMySQL().Db.Close()
	b.StopTimer()
	//375 ns/op
}

//-----------------------------------------------------------------------------
// Set map VS Set struct
//-----------------------------------------------------------------------------
//TODO:work in progress
func BenchmarkSetStruct(b *testing.B) {
	tu.SkipBLog(b)

	type Person struct {
		UserId    int    `db:"id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"first_name"`
	}

	b.ResetTimer()
	newMySQL()
	//db := getMySQL().Db

	//single data
	var person Person
	sql := "SELECT user_id, first_name, last_name FROM t_users WHERE delete_flg=?"

	for i := 0; i < b.N; i++ {
		//
		getMySQL().Db.SelectIns(sql, "0").ScanOne(&person)
		//db.SelectSQLAllFieldIns(sql, "0")
		//
	}
	getMySQL().Db.Close()
	b.StopTimer()

	//20000000	        87.1 ns/op (93.8 ns/op)
}

//-----------------------------------------------------------------------------
// NoCache VS Using Redis
//   Use for only heavy query
//-----------------------------------------------------------------------------
func BenchmarkCacheResponse(b *testing.B) {
	tu.SkipBLog(b)

	b.ResetTimer()
	newMySQL()
	for i := 0; i < b.N; i++ {
		//
		_, _ = getMySQL().getUserListOnCache()
		//
	}
	getMySQL().Db.Close()
	b.StopTimer()

	//20000000	        87.1 ns/op (93.8 ns/op)
	//ok  	github.com/hiromaily/golibs/db/mysql	1.885s
}
