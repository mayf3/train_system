package models

import (
	"github.com/astaxie/beego/orm"
)

type Person struct{
	Id	int		`orm:"pk;auto"`
	Name	string
	Grade	int
	Member	[]*Member	`orm:"reverse(many)"`
}

func init(){
	orm.RegisterModel(new(Person))
}
