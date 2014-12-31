package model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type UserOn struct {
	Id   int
	Name string `orm:"size(101)"`
}

func init() {
	// register model
	orm.RegisterModel(new(UserOn))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:000111@/golangtest?charset=utf8", 30)
	//orm.RunCommand()
}
