package mysql_test

import (
	"encoding/json"
	"os"
	"testing"

	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/mysql"
	"github.com/hiromaily/golibs/db/redis"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	u "github.com/hiromaily/golibs/utils"
)

type MySQL struct {
	Db *MS
}

//For embedded type
//type MySQL2 struct {
//	*MS
//}

var (
	db MySQL
	//db2 MySQL2
	//cahce
	cacheData map[string][]map[string]interface{}
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[MySQL]")
	if !tu.BenchFlg {
		newMySQL()

		if tu.BenchFlg {
			//Redis
			newRedis()
		}
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
//func newMySQL2() {
//	conf.SetTOMLPath("../../settings.toml")
//	c := conf.GetConf().MySQL
//
//	New(c.Host, c.DbName, c.User, c.Pass, c.Port)
//
//	db2 = MySQL2{}
//	db2.MS = GetDB()
//}

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

//func getMySQL2() *MySQL2 {
//	if db2.MS == nil {
//		//db2.DB, err = db2.Connection()
//		newMySQL2()
//	}
//	return &db2
//}

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
	var cacheKey = "sql001"

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

// nolint: unparam
func (ms *MySQL) getSimpleSQL2(id int) ([]map[string]interface{}, error) {
	sql := `SELECT user_id, first_name, last_name, create_datetime FROM t_users WHERE delete_flg=?`

	data, _, err := ms.Db.Select(sql, id)
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
// to get map data of string
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

// to get map data of string
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

// using default QueryRow func
func TestSelectOne(t *testing.T) {
	db := getMySQL().Db

	//1
	var userID int
	err := db.DB.QueryRow("SELECT user_id FROM t_users WHERE delete_flg=?", 0).Scan(&userID)
	if err != nil {
		t.Fatalf("[1]db.DB.QueryRow(): %s", db.Err)
	}
	lg.Debugf("user_id is %d", userID)

	//2
	err = db.DB.QueryRow("SELECT user_id FROM t_users").Scan(&userID)
	if err != nil {
		t.Fatalf("[2]db.DB.QueryRow(): %s", db.Err)
	}
	lg.Debugf("user_id is %d", userID)

}

//without placeholder
func TestSelectInsScanOne1(t *testing.T) {
	db := getMySQL().Db

	//1.int
	var userID int
	sql := "SELECT user_id FROM t_users"
	b := db.SelectIns(sql).ScanOne(&userID)
	if db.Err != nil {
		t.Fatalf("[1]db.SelectIns(): %s", db.Err)
	}
	lg.Debugf("user_id is %d, result is %v", userID, b)

	//2.string
	var firstName string
	sql = "SELECT first_name FROM t_users"
	b = db.SelectIns(sql).ScanOne(&firstName)
	if db.Err != nil {
		t.Fatalf("[2]db.SelectIns(): %s", db.Err)
	}
	lg.Debugf("firstName is %s, result is %v", firstName, b)

}

//plus placeholder
func TestSelectInsScanOne2(t *testing.T) {
	db := getMySQL().Db

	//1.int
	var userID int
	sql := "SELECT user_id FROM t_users WHERE delete_flg=?"
	b := db.SelectIns(sql, 0).ScanOne(&userID)
	if db.Err != nil {
		t.Fatalf("[1]db.SelectIns(): %s", db.Err)
	}
	lg.Debugf("user_id is %d, result is %v", userID, b)

	//2.string
	var firstName string
	sql = "SELECT first_name FROM t_users WHERE delete_flg=?"
	b = db.SelectIns(sql, 0).ScanOne(&firstName)
	if db.Err != nil {
		t.Fatalf("[2]db.SelectIns(): %s", db.Err)
	}
	lg.Debugf("firstName is %s, result is %v", firstName, b)

	//3.string
	var updated string
	sql = "SELECT update_datetime FROM t_users WHERE delete_flg=?"
	b = db.SelectIns(sql, 0).ScanOne(&updated)
	if db.Err != nil {
		t.Fatalf("[3]db.SelectIns(): %s", db.Err)
	}
	lg.Debugf("update_datetime is %s, result is %v", updated, b)
}

//pass struct and without placeholder
//TODO:this pattern doesn't work yet
//TODO:SelectInsにパラメータがないと動かん・・・
//panic: reflect.Set: value of type string is not assignable to type int [recovered]
//panic: reflect.Set: value of type string is not assignable to type int
func TestSelectInsScanOne3(t *testing.T) {
	tu.SkipLog(t)

	type Person struct {
		UserID    int    `db:"id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		//DateTime  time.Time `db:"create_datetime"`
		DateTime string `db:"create_datetime"`
	}

	db := getMySQL().Db

	//1.Single data
	var person Person
	sql := "SELECT user_id, first_name, last_name, create_datetime FROM t_users"
	db.SelectIns(sql)
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

//pass struct and plus placeholder
func TestSelectInsScanOne4(t *testing.T) {
	//tu.SkipLog(t)

	type Person struct {
		UserID    int    `db:"id"`
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

//without placeholder
//TODO:this pattern doesn't work yet
//TODO:SelectInsにパラメータがないと動かん・・・
//panic: reflect.Set: value of type string is not assignable to type int [recovered]
func TestSelectInsScan1(t *testing.T) {
	tu.SkipLog(t)

	db := getMySQL().Db

	var userIDs []int

	sql := "SELECT user_id FROM t_users"
	db.SelectIns(sql).Scan(&userIDs)
	if db.Err != nil {
		t.Errorf("[1]db.SelectIns(sql, 0).Scan(): %s", db.Err)
	}
	lg.Debugf("userIDs is %v", userIDs)
}

//plus placeholder
func TestSelectInsScan2(t *testing.T) {

	db := getMySQL().Db

	var userIDs []int

	sql := "SELECT user_id FROM t_users WHERE delete_flg=?"
	db.SelectIns(sql, 0).Scan(&userIDs)
	if db.Err != nil {
		t.Errorf("[1]db.SelectIns(sql, 0).Scan(): %s", db.Err)
	}
	lg.Debugf("userIDs is %v", userIDs)
}

//struct plus placeholder
func TestSelectInsScan3(t *testing.T) {
	//tu.SkipLog(t)

	type Person struct {
		UserID    int    `db:"id"`
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

//JSON Type without placeholder
func TestSelectInsScanOneJSON(t *testing.T) {
	tu.SkipLog(t)

	db := getMySQL().Db

	//1.Json Data as string
	var memo string
	sql := "SELECT memo2 FROM t_invoices"
	b := db.SelectIns(sql).ScanOne(&memo)
	if db.Err != nil {
		t.Fatalf("[1]db.SelectIns(): %s", db.Err)
	}
	lg.Debugf("memo2 is %s, result is %v", memo, b)

	//convert string to slice
	var converted []int
	json.Unmarshal([]byte(memo), &converted)
	lg.Debugf("converted memo2 is %v", converted)

	//2.Json Data as Array (This is not correct)
	var memo2 []int
	sql = "SELECT memo2 FROM t_invoices"
	b = db.SelectIns(sql).ScanOne(&memo2)
	if db.Err != nil {
		t.Fatalf("[2]db.SelectIns(): %s", db.Err)
	}
	lg.Debugf("memo2 is %s, result is %v", memo2, b)

}

//Insert JSON Type
func TestInsertJSON(t *testing.T) {
	tu.SkipLog(t)

	db := getMySQL().Db

	jsonBase := []int{10, 20, 30, 40, 50}
	retByte, _ := json.Marshal(jsonBase)

	// Insert
	sql := "INSERT INTO t_invoices (user_id, memo2) VALUES (1, ?)"
	newID, err := db.Insert(sql, string(retByte))
	if err != nil {
		t.Fatalf("[1]db.Insert(): %s", db.Err)
	}
	lg.Debugf("newID is %d", newID)
}

//Insert JSON Type
//func TestCallStored(t *testing.T) {
//	tu.SkipLog(t)
//
//	var result string
//
//	//var outArg string
//	//_, err := db.ExecContext(ctx, "ProcName", sql.Named("Arg1", Out{Dest: &outArg}))
//
//	sql := `CALL proc_search_slug('%s', '%s', @result);`
//	sql = fmt.Sprintf(sql, path, lastPath)
//
//	_, err := f.db.Exec(sql)
//	if err != nil {
//		return "", err
//	}
//
//	sql = `SELECT @result;`
//	err = f.db.Get(&result, sql)
//	if err != nil {
//		return "", err
//	}
//
//	return result, nil
//}
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
		UserID    int    `db:"id"`
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
