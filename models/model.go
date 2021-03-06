package models

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var DB *xorm.Engine

func SetUp(dsn string) error {

	var err error
	DB, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return err
	}

	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)
	return nil
}
