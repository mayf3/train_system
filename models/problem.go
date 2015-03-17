package models

import (
	"github.com/astaxie/beego/orm"
)

type Problem struct {
	Id    		int     `orm:"pk;auto"`
	Participant	int
	Status		int
	Information Information	`orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Problem))
}
