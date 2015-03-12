package main

import (
	_ "train_system/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func init(){
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	db_user		:= "myf"
	db_password := ""
	db_name		:= "test"
	orm.RegisterDataBase("default", "mysql", db_user + ":" + db_password + "@/" + db_name + "?charset=utf8")
}

func main() {
	beego.Run()
}
