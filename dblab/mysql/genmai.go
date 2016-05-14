package mysql

//genmaiなんだけど、genmaiの機能を使わない実装方法
//github.com/naoina/genmai
//Star数: たったの96 ...
//https://github.com/naoina/genmai
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/naoina/genmai"
)

//TODO: genmaiを使いながら普通のsqsを書く方法

// singleton architecture
func GetDBGenmaiInstance() *DBInfo {
	/*
		type DB struct {
			db      *sql.DB
			dialect Dialect
			tx      *sql.Tx
			m       sync.Mutex
			logger  logger
		}
	*/
	var err error
	if dbInfo.gdb == nil {
		dbInfo.gdb, err = connectionGenmai()
	}
	if err != nil {
		panic(err.Error())
	}

	return &dbInfo
}

func connectionGenmai() (*genmai.DB, error) {
	return genmai.New(&genmai.MySQLDialect{}, getDsn())
}

// Close
func (self *DBInfo) CloseGenmai() {
	self.gdb.Close()
}

// INSERT
func (self *DBInfo) InesrtSQLGenmai(insertSQL string, args ...interface{}) {

	//1.creates a prepared statement (placeholder)
	//insertSQL := "INSERT t_sight SET t_account_id=?, sight_name=?, latitude=?, longitude=?"
	stmt, err := self.gdb.DB().Prepare(insertSQL)
	if err != nil {
		panic(err.Error())
	}

	//2.set parameter to prepared statement
	//res, err := stmt.Exec(1, "テスト見所(INSERT)", 35.640353, 139.712969)
	res, err := stmt.Exec(args...)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	//3.レスポンスからidを取得
	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(id)
}
