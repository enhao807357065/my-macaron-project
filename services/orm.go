package main

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/go-xorm/xorm"
	"adwall/util"
	"adwall/models"
	"fmt"
)

var engine *xorm.Engine

func init() {
	fmt.Println("----------------------------------")
	var err error
	engine, err = xorm.NewEngine("mysql", util.GetMysqlUrl())

	if err != nil {
		panic(err)
	}

	err = engine.Ping()
	if err != nil {
		panic(err)
	}

	engine.ShowSQL(true)
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(30)

	tables := []interface{}{
		new(models.TestLh),
	}
	engine.Sync2(tables...)
}

func Ormer() *xorm.Engine {
	return engine
}