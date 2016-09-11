package gorp

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/coopernurse/gorp"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// GR is for database information and gorp instance
type GR struct {
	DB         *gorp.DbMap
	ServerInfo //embeded
}

// ServerInfo is for database information
type ServerInfo struct {
	host   string
	port   uint16
	dbname string
	user   string
	pass   string
}

var dbInfo GR

// New is for create instance
func New(host, dbname, user, pass string, port uint16) {

	if dbInfo.DB == nil {
		dbInfo.host = host
		dbInfo.port = port
		dbInfo.dbname = dbname
		dbInfo.user = user
		dbInfo.pass = pass

		db, err := dbInfo.Connection()

		if err != nil {
			panic(err.Error())
		}

		dbInfo.DB = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	}

	return
}

// GetDB is to get instance. singleton architecture
func GetDB() *GR {
	if dbInfo.DB == nil {
		panic(errors.New("DB instance is nil"))
	}
	return &dbInfo
}

func (gr *GR) getDsn() string {
	//If use nil on Date column, set *time.Time
	//Be careful when parsing is required on Date type
	// e.g. db, err := sql.Open("mysql", "root:@/?parseTime=true")
	param := "?charset=utf8&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		gr.user, gr.pass, gr.host, gr.port, gr.dbname, param)
}

// Connection is to connect to MySQL server
func (gr *GR) Connection() (*sql.DB, error) {
	//return sql.Open("mysql", getDsn())
	db, _ := sql.Open("mysql", gr.getDsn())
	return db, db.Ping()
}

// Close is to close connection
func (gr *GR) Close() {
	gr.DB.Db.Close()
}
