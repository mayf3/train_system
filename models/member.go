package models

import (
	"github.com/astaxie/beego/orm"
)

type Member struct{
	Id			int			`orm:"pk:auto"`
	Number		int
	Person		Person		`orm:"rel(fk)"`
	Information Information	`orm:"rel(fk)"`
}

func init(){
	orm.RegisterModel(new(Member))
}
