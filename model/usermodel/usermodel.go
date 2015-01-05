package usermodel

import (
	"github.com/astaxie/beego/orm"
	//exec init
	_ "../../model"
)

// Model Struct
type UserOn struct {
	Id   int
	Name string `orm:"size(1)"`
}

func init() {
	// register model
	orm.RegisterModel(new(UserOn))
}
