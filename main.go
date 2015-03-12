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
	db_password := "myfadmin"
	db_name		:= "test"
	orm.RegisterDataBase("default", "mysql", db_user + ":" + db_password + "@/" + db_name + "?charset=utf8")
}

func main() {
	o := orm.NewOrm()
	o.Raw("insert into tables(Contest_name,Problem_number,Source) values(?,?,?)", "table1", "2", "ss").Exec()
	beego.Run()
}
