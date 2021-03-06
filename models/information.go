package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Information struct {
	Id		int			`orm:"pk;auto"`
	Rank	int
	Name	string
	Table	*Tables		`orm:"rel(fk)"`
	Member	[]*Member	`orm:"reverse(many)"`
	Problem []*Problem	`orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Information))
}

// All
func (this *Information) GetInformationByTable(relation_table Tables) ([]Information, error) {
	o := orm.NewOrm()
	var all []Information
	_, err := o.QueryTable("information").Filter("Table", relation_table).OrderBy("Rank").All(&all)
	return all, err
}

// Single
func (this *Information) GetInformationById(id int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("information").Filter("Id", id).All(this)
	return err
}

func (this *Information) Insert() error{
	o := orm.NewOrm()
	_, err := o.Insert(this)
	fmt.Println(err)
	return err
}

func (this *Information) Update() error{
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}

func (this *Information) Delete() error{
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}
