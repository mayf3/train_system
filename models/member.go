package models

import (
	"github.com/astaxie/beego/orm"
)

type Member struct{
	Order		int
	Person		*Person		`orm:"pk:rel(fk)"`
	Information 	*Information	`orm:"pk:rel(fk)"`
}

func init(){
	orm.RegisterModel(new(Member))
}
