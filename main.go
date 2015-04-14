package main

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/go-sql-driver/mysql"

	_ "train_system/initialize"
	_ "train_system/models"
	_ "train_system/routers"
)

const (
	kDataBaseUser     = "iladmin"
	kDataBasePassword = ""
	kDataBaseName     = "train"
)

func init() {
	orm.Debug = true
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", kDataBaseUser+":"+kDataBasePassword+"@/"+kDataBaseName+"?charset=utf8&loc=Asia%2FShanghai")
}

func main() {
	orm.RunCommand()
	beego.Run()
}
