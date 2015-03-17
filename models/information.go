package models

import (
	"github.com/astaxie/beego/orm"
)

type Information struct {
	Id    	int     	`orm:"pk;auto"`
	Rank  	int     	`orm:"null"`
	Table 	*Tables 	`orm:"rel(fk)"`
	Member	[]*Member	`orm:"reverse(many)"`
	Problem []*Problem	`orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Information))
}
