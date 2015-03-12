package main

import (
	_ "train_system/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)


const (
	kDataBaseUser		=  "myf"
	kDataBasePassword	=	"myfadmin"
	kDataBaseName		=	"test"
)

func init(){
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", kDataBaseUser + ":" + kDataBasePassword + "@/" + kDataBaseName + "?charset=utf8")
}

func main() {
	orm.RunCommand()
	beego.Run()
}
