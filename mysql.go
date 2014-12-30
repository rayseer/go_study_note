package main

import (
	//类似于PHP里的PDO，多数都用这个来连接数据库
	"database/sql"
	"fmt"
	"log"
	//需要引入mysql package，不然无法驱动mysql，相当于PHP里的Mysql扩展，主要功能是对接Mysql与GO
	//引用mysql不会报错，因为使用了_
	//相当于调用了 mysql.Init() 方法
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//通过Socket来连接Mysql
	//但是应该存在Bug，如果密码里同时存在 ':' '@' '/' 可能会发生错误
	//开始连接报错如下：
	//2014/12/24 15:05:37 dial tcp 127.0.0.1:3306: connection refused
	//exit status 1
	//解决方案如下：在MAMP里Mysql开启Allow network access to Mysql
	//主要是允许Mysql通过Socket连接

	//TODO: 测试PHP里是否允许这样的连接
	db, err := sql.Open("mysql", "root:000111@tcp(localhost:3306)/golangtest")
	if err != nil {
		log.Fatal(err)
	}

	id := 1

	//如果使用"select * from TABLE"
	//必须枚举全部的列，例如：rows.Scan(&id, &name)
	//如果使用“select name from TABLE”
	//只需要 rows.Scan(&name)
	rows, err := db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
	}

}
