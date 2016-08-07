package cassandra_test

import (
	"flag"
	"fmt"
	"github.com/gocql/gocql"
	. "github.com/hiromaily/golibs/db/cassandra"
	lg "github.com/hiromaily/golibs/log"
	r "github.com/hiromaily/golibs/runtimes"
	"os"
	"testing"
	"time"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[Cassandra_TEST]", "/var/log/go/test.log")
	if *benchFlg == 0 {
	}
	//
	keyspace := "hiromaily"
	hosts := []string{"localhost"}
	New(hosts, keyspace)
}

func teardown() {
	if *benchFlg == 0 {
	}
	GetCass().Close()
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

//-----------------------------------------------------------------------------
// Insert
//-----------------------------------------------------------------------------
func TestInsert(t *testing.T) {
	db := GetCass()
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
		t.Logf("%v, %s, %s", id, firstName, lastName)
	}
}

func TestSelectAll(t *testing.T) {
	db := GetCass()
	//SELECT
	sql := `SELECT id, first_name, last_name FROM t_users`

	var id gocql.UUID
	var firstName string
	var lastName string

	iter := db.Session.Query(sql).Iter()
	//return value of Scan is bool
	for iter.Scan(&id, &firstName, &lastName) {
		t.Logf("%v, %s, %s", id, firstName, lastName)
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
		t.Logf("%s, %s, %s", email, firstName, lastName)
	}

}

//-----------------------------------------------------------------------------
// Delete Row
//-----------------------------------------------------------------------------
func TestDeleteRow(t *testing.T) {
	t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))
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
//Benchmark
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
