package main

import (
	"fmt"
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

func main() {
	o := orm.NewOrm()

	user := UserOn{Name: "slene"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := UserOn{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	/*
		// delete
		num, err = o.Delete(&u)
		fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	*/
}
