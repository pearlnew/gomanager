package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pearlnew/gomanager/pkg/utils/option"
	"xorm.io/core"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func InitDB() error {
	var err error
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&allowNativePasswords=true",
		option.GlobalOption.Db.User,
		option.GlobalOption.Db.Passwd,
		option.GlobalOption.Db.Host,
		option.GlobalOption.Db.Port,
		option.GlobalOption.Db.Database)
	engine, err = xorm.NewEngine("mysql", connectString)
	if err != nil {
		return fmt.Errorf("创建数据库链接失败: %s", err.Error())
	}
	engine.SetMaxIdleConns(50)
	engine.SetMaxOpenConns(1000)
	engine.SetConnMaxLifetime(60*5)
	engine.SetMapper(core.SameMapper{})
	//engine.ShowSQL(true)
	return nil
}