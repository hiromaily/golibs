package mysql

import (
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//https://github.com/jinzhu/gorm

// singleton architecture
func GetDBGormInstance() *DBInfo {
	var err error
	if dbInfo.gormDb == nil {
		dbInfo.gormDb, err = connectionGorm()
	}
	if err != nil {
		panic(err.Error())
	}

	return &dbInfo
}

func connectionGorm() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", getDsn())
	if err != nil {
		return nil, err
	} else {
		return &db, nil
	}
}
