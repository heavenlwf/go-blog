package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"log"
	"github.com/heavenlwf/go-blog/pkg/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	)

var db *gorm.DB

type Model struct {
	gorm.Model
}

func init() {

	sec := config.Conf.MySQL
	conf := sec["Blog"]
	initConn(conf)
}

func initConn(conf *config.Mysql) {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	dbType = "mysql"
	dbName = conf.NAME
	user = conf.USER
	password = conf.PASSWORD
	host = conf.HOST
	tablePrefix = conf.TablePrefix

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Fatal(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
