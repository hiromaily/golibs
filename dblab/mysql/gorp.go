package mysql

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hiromaily/golibs/dblab/models"
)

// singleton architecture
func GetDBGorpInstance() *DBInfo {
	var err error
	if dbInfo.gorpDbMap == nil {
		dbInfo.gorpDbMap, err = connectionGorp()
	}
	if err != nil {
		panic(err.Error())
	}

	return &dbInfo
}

func connectionGorp() (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", getDsn())
	if err != nil {
		// construct a gorp DbMap
		return nil, err
	} else {
		//construct a gorp DbMap
		dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
		return dbmap, nil
	}
}

// Select (複数フィールド、全フィールド共に対応可能)
func (self *DBInfo) SelectSQLGorp(selectSQL string) {
	var users []models.Users
	_, err := self.gorpDbMap.Select(&users, selectSQL)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("%#v \n", users)
	//fmt.Printf("%#v \n", users[0])
	fmt.Printf("%d : %s %s \n", users[0].UserId, users[0].FirstName, users[0].LastName)
}

// Insert
func (self *DBInfo) InesrtSQLFromStructGorp() {
	//現在の最新のidを取得
	//SELECT MAX(user_id) FROM t_users;

	//idフィールドのauto increment機能がこのやり方だと不可能なので、
	// 利用価値があるかどうかは、テーブルによる
	currentMaxId := self.GetMaxId()
	users := models.NewUser(currentMaxId+1, "key", "Lacyl")
	//fmt.Println(users)

	//No table found for type: Users -> 実在するテーブル名との関連付けが必要
	self.gorpDbMap.AddTableWithName(models.Users{}, "t_users").SetKeys(true, "user_id")

	//gorpで使用する引用符は「"」のため、修正せねばならない。
	err := self.gorpDbMap.Insert(&users)
	if err != nil {
		panic(err.Error())
	}
}

// 戻り値がintのSQL用
// 最大値を取得
func (self *DBInfo) GetMaxId() uint16 {
	//SelectIntはint64を返すのでキャストが必要
	maxId, err := self.gorpDbMap.SelectInt("SELECT MAX(user_id) FROM t_users")
	if err != nil {
		panic(err.Error())
	}
	return uint16(maxId)
}

// 件数を取得
func (self *DBInfo) GetCount() uint16 {
	count, err := self.gorpDbMap.SelectInt("SELECT count(user_id) FROM t_users")
	if err != nil {
		panic(err.Error())
	}
	return uint16(count)
}

// Insert
func (self *DBInfo) InesrtSQLGorp(insertSQL string) {
	_, err := self.gorpDbMap.Exec(insertSQL)
	if err != nil {
		panic(err.Error())
	}
}

func (self *DBInfo) GetRecord() {
	//t := self.gorpdbMap.AddTableWithName(models.Users{}, "t_users").SetKeys(true, "Id")

	//users := make([]models.Users, 0, count)
	users := new(models.Users)

	_, err := self.gorpDbMap.Get(users, 1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(users)
}
