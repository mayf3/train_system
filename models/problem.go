package models

import (
	"github.com/astaxie/beego/orm"
)

type Problem struct {
	Id			int				`orm:"pk;auto"`
	Number		int
	Participant	int
	Status		int
	Information	*Information	`orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Problem))
}

// All
func (this *Problem) GetProblemByInformation(relation_information Information) ([]Problem, error) {
	o := orm.NewOrm()
	var all []Problem
	_, err := o.QueryTable("problem").Filter("Information", relation_information).All(&all)
	return all, err
}

// Single
func (this *Problem) Insert() error{
	o := orm.NewOrm()
	_, err := o.Insert(this)
	return err
}

func (this *Problem) Update() error{
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}

func (this *Problem) Delete() error{
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}
