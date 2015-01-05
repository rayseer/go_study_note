package user

import (
	//TODO: 不确定为什么还要加上usermodel，不然会报错
	//使用_ 并不是引入包时候执行init
	//而是调用方式时，不用加上包名
	"../../model/usermodel"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"net/http"
)

/*
TODO: DBMap
struct和数据库结构同步,使用会同步过去
可以使用缓存,例如Memcache
可以制定规则，不符合规则的不能insert or update or delete
selete 也可以制定规则，但是可选项，因为类似Like等
简化操作, 比如FindAll() 等
查询屏蔽字段
*/

func PostLogin(w http.ResponseWriter, r *http.Request) {
	db := orm.NewOrm()
	user := usermodel.UserOn{Name: "slene"}
	// insert
	id, err := db.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	params := r.URL.Query()
	lastName := params.Get(":last")
	firstName := params.Get(":first")
	fmt.Fprintf(w, "you are %s %s", firstName, lastName)
}
