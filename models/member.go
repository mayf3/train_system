package models

import (
	"github.com/astaxie/beego/orm"
)

type Member struct{
	Id		int		`orm:"pk;auto"`
	Name	string
	Grade	int
}

func init(){
	orm.RegisterModel(new(Member))
}
