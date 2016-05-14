package mysql

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hiromaily/golibs/config"
	"github.com/hiromaily/golibs/dblab/models"
	"github.com/jinzhu/gorm"
	"github.com/naoina/genmai"
)

//TODO:トランザクションの機能もあるので、どこかに追加しておく
//TODO:異なるlibraryを使っているが、各funcのInterfaceを統一すればよいのでは？
//http://qiita.com/tenntenn/items/dddb13c15643454a7c3b
//http://go-database-sql.org/

type DBInfo struct {
	db        *sql.DB
	gdb       *genmai.DB
	gorpDbMap *gorp.DbMap
	gormDb    *gorm.DB
}

var dbInfo DBInfo

// singleton architecture
func GetDBInstance() *DBInfo {
	var err error
	if dbInfo.db == nil {
		dbInfo.db, err = connection()
	}
	if err != nil {
		panic(err.Error())
	}

	return &dbInfo
}

func getDsn() string {
	conf := config.GetConfInstance()
	//TODO:日付をperseするときには注意が必要
	//?charset=utf8&parseTime=True&loc=Local
	param := "?parseTime=true"
	return fmt.Sprintf("%s:%s@/%s%s", conf.Database.User, conf.Database.Pass, conf.Database.DbName, param)
}

// Connection
func connection() (*sql.DB, error) {

	//db, err := sql.Open("mysql", "user:password@/dbname")
	//fmt.Println(fmt.Sprintf("%s:%s@/%s", conf.Database.User, conf.Database.Pass, conf.Database.DbName))
	return sql.Open("mysql", getDsn())
	//defer db.Close() // 関数がリターンする直前に呼び出される
}

// Close
func (self *DBInfo) Close() {
	self.db.Close()
}

// 単純に実行
func (self *DBInfo) ExecSQL(sqlString string) {
	//result, err := self.db.Exec("INSERT t_users SET first_name=?, last_name=?", "Mika", "Haruda")
	result, err := self.db.Exec(sqlString)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
}

// BULKの実行(未完)
func (self *DBInfo) BulkExecSQL(sqlString string, args ...interface{}) {
	//utils.convertToInterface()
	//result, err := self.db.Exec("INSERT t_users SET first_name=?, last_name=?", "mitsuo", "fujita")
	result, err := self.db.Exec(sqlString, args...)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
}

// INSERT
func (self *DBInfo) InesrtSQL(insertSQL string, args ...interface{}) {
	fmt.Println("\nInesrtSQL()--------")

	//1.creates a prepared statement (placeholder)
	//insertSQL := "INSERT t_users SET first_name=?, last_name=?"
	stmt, err := self.db.Prepare(insertSQL)
	if err != nil {
		panic(err.Error())
	}

	//2.set parameter to prepared statement
	//res, err := stmt.Exec("mitsuo", "fujita")
	res, err := stmt.Exec(args...)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close() //statementもcloseする必要がある

	//3.レスポンスからidを取得
	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Inserted User : %d\n", id)
}

// UPDATE
func (self *DBInfo) UpdateSQL(updateSQL string, args ...interface{}) {
	fmt.Println("\nUpdateSQL()--------")

	//qu1.creates a prepared statement (placeholder)
	//updateSQL := "UPDATE t_users SET first_name=? WHERE user_id=?"
	stmt, err := self.db.Prepare(updateSQL)
	if err != nil {
		panic(err.Error())
	}

	//2.set parameter to prepared statement
	//res, err := stmt.Exec("genjiro", 3)
	res, err := stmt.Exec(args...)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close() //statementもcloseする必要がある

	//3.変化のあった行を取得
	rows, err := res.RowsAffected()

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("changed rows number : %d\n", rows)
}

// SELECT 1 : 部分フィールド取得 [未完]
func (self *DBInfo) SelectSQL() {
	fmt.Println("\nSelectSQL()--------")

	//1. create sql and exec
	rows, err := self.db.Query("SELECT user_id, first_name, last_name FROM t_users WHERE delete_flg=?", "0")

	if err != nil {
		panic(err.Error())
	}

	//field on table
	var (
		userId    uint16
		firstName string
		lastName  string
	)
	defer rows.Close() //rowsをcloseする必要がある
	for rows.Next() {
		//レコードから情報を取得
		err := rows.Scan(&userId, &firstName, &lastName)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("got user name is %s %s\n", firstName, lastName)
	}
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

}

// SELECT :  全フィールド取得
// パフォーマンスが悪いのであまり使わないが、例のみ
func (self *DBInfo) SelectSQLAllField(selectSQL string, args ...interface{}) {
	fmt.Println("\nSelectSQLAllField()--------")

	//1. create sql and exec
	//rows, err := self.db.Query("SELECT * FROM t_users WHERE delete_flg=?", "0")
	rows, err := self.db.Query(selectSQL, args...)
	if err != nil {
		panic(err.Error())
	}

	// カラム名を取得
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() { //true or false
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
	}

}

// SELECT 3 : 1レコード取得 / パターン1 [未完]
func (self *DBInfo) SelectOneRowSQL() {
	fmt.Println("\nSelectOneRowSQL()--------")

	//field on table
	var (
		firstName string
		lastName  string
	)
	userId := 1

	//1. create sql and exec
	err := self.db.QueryRow("SELECT first_name, last_name FROM t_users WHERE user_id=?", userId).Scan(&firstName, &lastName)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("user_id = %d : got user name is %s %s\n", userId, firstName, lastName)
}

// SELECT 4 : 1レコード取得 / パターン2 [未完]
func (self *DBInfo) SelectOneRowSQL2() {
	fmt.Println("\nSelectOneRowSQL2()--------")

	//field on table
	var (
		firstName string
		lastName  string
	)
	userId := 2

	//1. create sql and exec
	stmt, err := self.db.Prepare("SELECT first_name, last_name FROM t_users WHERE user_id=?")
	if err != nil {
		panic(err.Error())
	}

	//2. QueryRow
	//QueryRowのパラメータは引数
	err = stmt.QueryRow(userId).Scan(&firstName, &lastName)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("user_id = %d : got user name is %s %s\n", userId, firstName, lastName)
}

// SELECT Count(id) 5 : 行数を取得
func (self *DBInfo) SelectCount(countSql string, args ...interface{}) int {
	fmt.Println("\nSelectCount()--------")

	//field on table
	var count int

	//1. create sql and exec
	//err := self.db.QueryRow("SELECT count(user_id) FROM t_users WHERE delete_flg=?", "0").Scan(&count)
	err := self.db.QueryRow(countSql, args...).Scan(&count)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(count)

	return count
}

// TODO:取得したフィールドをまとめて変数にセットしたい。。。[未完]
func (self *DBInfo) SelectGetFieldData() {
	fmt.Println("\nSelectGetFieldData()--------")

	//1. count
	countSql := "SELECT count(user_id) FROM t_users WHERE delete_flg=?"
	count := self.SelectCount(countSql, "0")

	//2.領域を確保
	userIds := make([]uint16, 0, count)
	users := make([]models.Users, 0, count)

	//3.. create sql and exec
	rows, err := self.db.Query("SELECT user_id, first_name, last_name FROM t_users WHERE delete_flg=?", "0")
	//rows, err := self.db.Query("SELECT user_id, first_name, last_name, (SELECT count(user_id) FROM t_users WHERE delete_flg=?) AS cnt FROM t_users WHERE delete_flg=?", "0", "0")

	if err != nil {
		panic(err.Error())
	}

	//field on table
	var (
		userId    uint16
		firstName string
		lastName  string
	)
	defer rows.Close() //rowsをcloseする必要がある
	idx := 0
	for rows.Next() {
		//レコードから情報を取得
		err := rows.Scan(&userId, &firstName, &lastName)
		if err != nil {
			panic(err.Error())
		}

		//TODO:まとめて取得した値を変数にセットしたい
		//Only UserId
		userIds = append(userIds, userId)

		//Users Table
		users = append(users, models.NewUser(userId, firstName, lastName))
		//users = append(users, models.NewUser(firstName, lastName))

		fmt.Printf("user_id is %d, %d\n", users[idx].UserId, userIds[idx])
		idx++

	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("userIds %#v\n", userIds)
	fmt.Printf("userIds %+v\n", userIds)

	fmt.Printf("users %#v\n", users)

}

//-----------------------------------------------------------------------------
// 構造体で、DBのフィールドが定義されている場合
//-----------------------------------------------------------------------------
