package cassandra_test

import (
	"os"
	"testing"
	"time"

	"github.com/gocql/gocql"

	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/cassandra"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[Cassandra]")

	connection()
}

func teardown() {
	if GetCass() != nil {
		GetCass().Close()
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

func connection() {
	c := conf.GetConf().Cassa

	hosts := []string{c.Host}
	//port := 9042
	err := New(hosts, int(c.Port), c.KeySpace)
	if err != nil {
		lg.Errorf("New() error: %s", err)
	}
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
// Insert
//-----------------------------------------------------------------------------
func TestInsert(t *testing.T) {
	tu.SkipLog(t)

	db := GetCass()
	if db == nil {
		t.Fatal("db object should not be nil")
	}

	//INSERT
	sql := `INSERT INTO t_users
	(id, first_name, last_name, email, password, age, created_at, updated_at)
	VALUES
	(?, ?, ?, ?, ?, ?, ?, ?)`

	err := db.Session.Query(sql, gocql.TimeUUID(), "taro", "yamada", "bb@test.jp", "aaaa", 24, time.Now(), time.Now()).Exec()
	if err != nil {
		t.Errorf("INSERT error: %s", err)
	}
}

//-----------------------------------------------------------------------------
// Select
//-----------------------------------------------------------------------------
func TestSelectOne(t *testing.T) {
	tu.SkipLog(t)

	db := GetCass()
	//SELECT
	sql := `SELECT id, first_name, last_name FROM t_users LIMIT 1`

	var id gocql.UUID
	var firstName string
	var lastName string

	//err := db.Session.Query(sql, nil).Consistency(gocql.One).Scan(&id, &firstName, &lastName)
	// gocql: expected 0 values send got 1
	err := db.Session.Query(sql).Consistency(gocql.One).Scan(&id, &firstName, &lastName)
	if err != nil {
		t.Errorf("TestSelectOne error: %s", err)
	} else {
		lg.Debugf("%v, %s, %s", id, firstName, lastName)
	}
}

func TestSelectAll(t *testing.T) {
	tu.SkipLog(t)

	db := GetCass()
	//SELECT
	sql := `SELECT id, first_name, last_name FROM t_users`

	var id gocql.UUID
	var firstName string
	var lastName string

	iter := db.Session.Query(sql).Iter()
	//return value of Scan is bool
	for iter.Scan(&id, &firstName, &lastName) {
		lg.Debugf("%v, %s, %s", id, firstName, lastName)
	}

	//Close
	if err := iter.Close(); err != nil {
		t.Fatal(err)
	}
}

//-----------------------------------------------------------------------------
// Update
//-----------------------------------------------------------------------------
func TestUpdate(t *testing.T) {
	tu.SkipLog(t)

	db := GetCass()
	//UPDATE
	sql := `UPDATE t_users SET email = ? WHERE id = ? IF EXISTS`

	err := db.Session.Query(sql, "ccc@test.jp", "ac9321f5-5089-11e6-ac5d-acbc32b5de29").Exec()
	if err != nil {
		//Some partition key parts are missing: id
		t.Errorf("UPDATE error: %s", err)
	}

	//check
	sql = `SELECT first_name, last_name, email FROM t_users LIMIT 1`

	var firstName string
	var lastName string
	var email string

	err = db.Session.Query(sql).Consistency(gocql.One).Scan(&firstName, &lastName, &email)
	if err != nil {
		t.Errorf("check is invalid after updated data,  error: %s", err)
	} else {
		lg.Debugf("%s, %s, %s", email, firstName, lastName)
	}

}

//-----------------------------------------------------------------------------
// Delete Row
//-----------------------------------------------------------------------------
func TestDeleteRow(t *testing.T) {
	tu.SkipLog(t)

	db := GetCass()
	//DELETE
	sql := `DELETE FROM t_users WHERE id=?`

	err := db.Session.Query(sql, "ac9321f5-5089-11e6-ac5d-acbc32b5de29").Exec()
	if err != nil {
		//DELETE error: Some partition key parts are missing: id
		t.Errorf("DELETE Row error: %s", err)
	}
}

//-----------------------------------------------------------------------------
// Delete Data
//-----------------------------------------------------------------------------
func TestDeleteData(t *testing.T) {
	tu.SkipLog(t)

	db := GetCass()
	//DELETE
	sql := `DELETE last_name FROM t_users WHERE id=?`
	//->This sql is set null to filed.

	err := db.Session.Query(sql, "ac9321f5-5089-11e6-ac5d-acbc32b5de29").Exec()
	if err != nil {
		//DELETE error: Some partition key parts are missing: id
		t.Errorf("DELETE Data error: %s", err)
	}
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkCassandra(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
