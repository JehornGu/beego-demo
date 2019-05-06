package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 链接数据库
	orm.RegisterDataBase("default", "mysql", "root:111111@tcp(127.0.0.1:3306)/civ?charset=utf8")
	// 映射model数据(new(Type)...)
	orm.RegisterModel(new(User))
	// 生成表(别名, 是否强制更新, 是否可见)
	orm.RunSyncdb("default", false, true)
}
