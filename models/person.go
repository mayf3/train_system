package models

import (
	"github.com/astaxie/beego/orm"
)

type Person struct{
	Id		int		`orm:"pk;auto"`
	Name	string
	Grade	int
	Member	[]*Member	`orm:"reverse(many)"`
}

func init(){
	orm.RegisterModel(new(Person))
}

// All
func (this *Person) GetAllPerson() ([]Person, error) {
	o := orm.NewOrm()
	var all []Person
	_, err := o.QueryTable("person").All(&all)
	return all, err
}

// Single
func (this *Person) GetPersonById(id int) error {
	o := orm.NewOrm()
	err := o.QueryTable("person").Filter("Id", id).One(this)
	return err
}

func (this *Person) GetPersonByName(name string) error {
	o := orm.NewOrm()
	err := o.QueryTable("person").Filter("Name", name).One(this)
	return err
}

func (this *Person) Insert() error{
	o := orm.NewOrm()
	_, err := o.Insert(this)
	return err
}

func (this *Person) Update() error{
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}

func (this *Person) Delete() error{
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}
