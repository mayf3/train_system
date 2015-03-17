package models

import (
	"github.com/astaxie/beego/orm"
)

type Problem struct {
	Number		int		`orm:"pk"`
	Participant	int
	Status		int
	Information 	*Information	`orm:"pk;rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Problem))
}
