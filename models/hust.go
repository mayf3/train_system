package models

import (
	"github.com/astaxie/beego/orm"
)

type Hust struct {
	Id      int     `orm:"pk;auto"`
	Context string  `orm:"type(text)"`
	Table   *Tables `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Hust))
}

// All

// Single
func (this *Hust) GetHustById(id int) error {
	o := orm.NewOrm()
	err := o.QueryTable("hust").Filter("Id", id).One(this)
	return err
}

func (this *Hust) GetHustByTabel(table Tables) error {
	o := orm.NewOrm()
	err := o.QueryTable("hust").Filter("table", table).One(this)
	return err
}

func (this *Hust) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(this)
	return err
}

func (this *Hust) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}

func (this *Hust) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}
