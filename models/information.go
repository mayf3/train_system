package models

import (
	"github.com/astaxie/beego/orm"
)

type Information struct {
	Id    int     `orm:"pk;auto"`
	Table *Tables `orm:"rel(fk)"`
	Rank  int     `orm:"null"`
	Info  string  `orm:"null"`
}

func init() {
	orm.RegisterModel(new(Information))
}
