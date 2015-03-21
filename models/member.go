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

func (this *Member) GetMemberByInformation(information Information) ([]Member, error) {
	o := orm.NewOrm()
	var all []Member
//	_, err := o.QueryTable("member").RelatedSel("person").Filter("Information", relation_information).All(&all)
	_, err := o.QueryTable("member").Filter("Information", information).OrderBy("Order").All(&all)
	return all, err
}

// Single
func (this *Member) GetMemberByInformationAndOrder(information Information, order int) error{
	o := orm.NewOrm()
	err := o.QueryTable("member").Filter("Information", information).Filter("Order", order).One(this)
	return err
}

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
