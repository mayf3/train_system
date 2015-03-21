package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	. "train_system/models"
)

type PersonController struct {
	beego.Controller
}

func (this *PersonController) Get() {
	operation := this.GetString(":action")
	fmt.Println(operation)
	switch operation {
	case "add":
		this.add()
		this.Redirect("/person", 302)
	case "edit":
		this.edit()
		this.Redirect("/person", 302)
	case "delete":
		this.del()
		this.Redirect("/person", 302)
	default:
		this.show();
	}
}

func (this *PersonController) Post(){
	operation := this.GetString(":action")
	switch operation {
	case "add":
		this.add()
		this.Redirect("/person", 302)
	default:
		this.Redirect("/person", 302)
	}
}

func (this *PersonController) show(){
	var(
		_ error
		tmp string
		person Person
		all_person []Person
	)
	tmp = this.Input().Get("admin")
	if tmp == "true"{
		this.Data["show"] = 1
	}
	all_person, _ = person.GetAllPerson()
	this.Data["all_person"] = all_person
	this.TplNames = "person.tpl"
}

func (this *PersonController) add(){
	var(
		err error
		grade int
		name string
		person Person
	)
	name = this.GetString("name")
	grade, err = this.GetInt("grade")
	if err == nil && person.IsExistByNameAndGrade(name, grade) == false{
		person.Name = name;
		person.Grade = grade;
		person.Insert()
	}
}

func (this *PersonController) del(){
	var(
		_ error
		person Person
	)
	person.Id, _ = this.GetInt("person_id")
	_ = person.GetPersonById(person.Id)
	person.Delete()
}

func (this *PersonController) edit(){
}
