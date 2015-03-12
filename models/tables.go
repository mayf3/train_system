package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tables struct{
	Id				int			`orm:"pk;auto"`
	Contest_name	string
	Problem_number	int			`orm:"null"`
	Source			string		`orm:"null"`
	Create_time		time.Time	`orm:"auto_now_add;type(datetime)"`
	Information		[]*Information	`orm:"reverse(many)"`
}

func init(){
	orm.RegisterModel(new(Tables))
}

func GetAllTable() (maps []orm.Params, err error){
	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM tables").Values(&maps)
	if num > 0 && err == nil {
		return maps, err
	}
	return maps, err
}
