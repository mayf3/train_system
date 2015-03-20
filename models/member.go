package models

import (
	"github.com/astaxie/beego/orm"
)

type Member struct{
	Id			int				`orm:"pk;auto"`
	Order		int
	Person		*Person			`orm:"rel(fk)"`
	Information	*Information	`orm:"rel(fk)"`
}

func init(){
	orm.RegisterModel(new(Member))
}

// All
func (this *Member) GetMemberByInformation(relation_information Information) ([]Member, error) {
	o := orm.NewOrm()
	var all []Member
	_, err := o.QueryTable("member").Filter("Information", relation_information).All(&all)
	return all, err
}

// Single
func (this *Member) Insert() error{
	o := orm.NewOrm()
	_, err := o.Insert(this)
	return err
}

func (this *Member) Update() error{
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}

func (this *Member) Delete() error{
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}
