package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hiromaily/golibs/utils"
)

//TODO:トランザクションの機能もあるので、どこかに追加しておく
//TODO:異なるlibraryを使っているが、各funcのInterfaceを統一すればよいのでは？
//http://qiita.com/tenntenn/items/dddb13c15643454a7c3b
//http://go-database-sql.org/

type DBInfo struct {
	db     *sql.DB
	host   string
	port   uint16
	dbname string
	user   string
	pass   string
}

var dbInfo DBInfo

func New(host, dbname, user, pass string, port uint16) {
	var err error
	if dbInfo.db == nil {
		dbInfo.host = host
		dbInfo.port = port
		dbInfo.dbname = dbname
		dbInfo.user = user
		dbInfo.pass = pass

		dbInfo.db, err = connection()
	}
	fmt.Printf("dbInfo.db %#v\n", *dbInfo.db)
	if err != nil {
		panic(err.Error())
	}
}

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
	//If use nil on Date column, set *time.Time
	//Be careful when parsing is required on Date type
	// e.g. db, err := sql.Open("mysql", "root:@/?parseTime=true")
	param := "?charset=utf8&parseTime=True&loc=Local"
	//user:password@tcp(localhost:3306)/dbname?tls=skip-verify&autocommit=true
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s", dbInfo.user, dbInfo.pass, dbInfo.host, dbInfo.port, dbInfo.dbname, param)
}

// Connection
// Be careful, sql.Open() doesn't return err. Use db.Ping() to check DB condition.
func connection() (*sql.DB, error) {
	//return sql.Open("mysql", getDsn())
	db, _ := sql.Open("mysql", getDsn())
	return db, db.Ping()
}

// Close
func (self *DBInfo) Close() {
	self.db.Close()
}

// SELECT Count: Get number of rows
func (self *DBInfo) SelectCount(countSql string, args ...interface{}) (int, error) {
	//field on table
	var count int

	//1. create sql and exec
	//err := self.db.QueryRow("SELECT count(user_id) FROM t_users WHERE delete_flg=?", "0").Scan(&count)
	err := self.db.QueryRow(countSql, args...).Scan(&count)
	if err != nil {
		//panic(err.Error())
		return 0, err
	}

	return count, nil
}

// Convert result of select into Map[] type. Return multiple array map and interface(plural lines)
func (this *DBInfo) convertRowsToMaps(rows *sql.Rows) ([]map[string]interface{}, []string, error) {
	// Get column name
	columns, err := rows.Columns()
	if err != nil {
		//panic(err.Error())
		return nil, nil, err
	}

	//variable for stored each field data
	//values := make([]sql.RawBytes, len(columns)) //it cause error
	values := make([]interface{}, len(columns))

	// rows.Scan は引数に `[]interface{}`が必要.
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		//I don't know why set address to another variable
		//set address of value to variable for scan
		scanArgs[i] = &values[i]
	}

	//retMaps := []map[string]string{}
	//rowdata := map[string]string{}
	retMaps := []map[string]interface{}{}
	//
	for rows.Next() { //true or false
		//Get data into scanArgs
		err = rows.Scan(scanArgs...)

		if err != nil {
			//panic(err.Error())
			return nil, columns, err
		}

		rowdata := map[string]interface{}{}

		//var v string
		for i, value := range values {
			//Check type
			//val := reflect.ValueOf(value) // ValueOfでreflect.Value型のオブジェクトを取得
			//fmt.Println("val.Type()", val.Type(), "val.Kind()", val.Kind()) // Typeで変数の型を取得

			if utils.CheckInterface(value) == "[]uint8" {
				//[]uint8 to []byte to string
				if tmp, ok := value.([]byte); ok {
					//value = strconv.Itoa(int(tmp))
					value = string(tmp)
				}
			}

			// Here we can check if the value is nil (NULL value)
			//if value == nil {
			//	v = "NULL"
			//} else {
			//	v = string(value)
			//}

			//if b, ok := value.([]byte); ok{
			//	v = string(b)
			//} else {
			//	v = "NULL"
			//}

			//rowdata[columns[i]] = v
			rowdata[columns[i]] = value
			//fmt.Println(columns[i], ": ", v)
		}
		retMaps = append(retMaps, rowdata)
	}
	return retMaps, columns, nil
}

// SELECT : Get All field you set(Though you get only record, use it.)
func (self *DBInfo) SelectSQLAllField(selectSQL string, args ...interface{}) ([]map[string]interface{}, []string, error) {

	//1. create sql and exec
	//rows, err := self.db.Query("SELECT * FROM t_users WHERE delete_flg=?", "0")
	rows, err := self.db.Query(selectSQL, args...)
	if err != nil {
		//panic(err.Error())
		return nil, nil, err
	}

	return self.convertRowsToMaps(rows)
}

// Execution simply
func (self *DBInfo) ExecSQL(sqlString string, args ...interface{}) error {
	//result, err := self.db.Exec("INSERT t_users SET first_name=?, last_name=?", "Mika", "Haruda")
	_, err := self.db.Exec(sqlString, args...)
	return err
}

// INSERT
func (self *DBInfo) InesrtSQL(insertSQL string, args ...interface{}) (int64, error) {
	//1.creates a prepared statement (placeholder)
	//insertSQL := "INSERT t_users SET first_name=?, last_name=?"
	stmt, err := self.db.Prepare(insertSQL)
	if err != nil {
		//panic(err.Error())
		return 0, err
	}

	//2.set parameter to prepared statement
	//res, err := stmt.Exec("mitsuo", "fujita")
	res, err := stmt.Exec(args...)
	if err != nil {
		//panic(err.Error())
		return 0, err
	}

	defer stmt.Close() //statementもcloseする必要がある

	//3.Get id from response
	//id, err := res.LastInsertId()
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Printf("Inserted User : %d\n", id)
	return res.LastInsertId()
}

// UPDATE
func (self *DBInfo) UpdateSQL(updateSQL string, args ...interface{}) (int64, error) {

	//1.creates a prepared statement (placeholder)
	//updateSQL := "UPDATE t_users SET first_name=? WHERE user_id=?"
	stmt, err := self.db.Prepare(updateSQL)
	if err != nil {
		//panic(err.Error())
		return 0, err
	}

	//2.set parameter to prepared statement
	//res, err := stmt.Exec("genjiro", 3)
	res, err := stmt.Exec(args...)
	if err != nil {
		//panic(err.Error())
		return 0, err
	}

	defer stmt.Close() //statementもcloseする必要がある

	//3.Get number of changed rows
	//rows, err := res.RowsAffected()
	//if err != nil {
	//	panic(err.Error())
	//}
	return res.RowsAffected()
}
