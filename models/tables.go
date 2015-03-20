package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// TODO make a tool.go file for other function like change_string_to_int

type Tables struct {
	Id				int `orm:"pk;auto"`
	ContestName		string
	ProblemNumber	int
	Source			string
	CreateTime		time.Time      `orm:"auto_now_add;type(datetime)"`
	Information		[]*Information `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Tables))
}

// All
func (this *Tables) GetAllTable() ([]Tables, error) {
	o := orm.NewOrm()
	var all []Tables
	_, err := o.QueryTable("tables").All(&all)
	return all, err
}

// Single
func (this *Tables) GetTableById(id int) error {
	o := orm.NewOrm()
	err := o.QueryTable("tables").Filter("Id", id).One(this)
	return err
}

func (this *Tables) Insert() error{
	o := orm.NewOrm()
	_, err := o.Insert(this)
	return err
}

func (this *Tables) Update() error{
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}

func (this *Tables) Delete() error{
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}
