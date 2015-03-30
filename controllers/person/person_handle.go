package person

import (
	//"fmt"

	"train_system/models"
)

func (this *PersonController) getAllPerson() []models.Person {
	var (
		err        error
		person     models.Person
		all_person []models.Person
	)
	all_person, err = person.GetAllPersonOrderById()
	if err != nil {
		this.Abort("Can't get all person information")
	}
	return all_person
}

func (this *PersonController) add() {
	var (
		err    error
		grade  int
		name   string
		person models.Person
	)
	name = this.GetString("name")
	if name == "" {
		this.Abort("Name can't be empty")
	}
	grade, err = this.GetInt("grade")
	if err != nil {
		this.Abort("Grade rrror")
	}
	if person.IsExistByNameAndGrade(name, grade) == false {
		this.Abort("Aleady exists person")
	}
	person.Name = name
	person.Grade = grade
	person.Insert()
}

func (this *PersonController) del() {
	var (
		err    error
		person models.Person
	)
	person.Id, err = this.GetInt("person_id")
	if err != nil {
		this.Abort("Error person id")
	}
	err = person.GetPersonById(person.Id)
	if err != nil {
		this.Abort("No such person")
	}
	person.Delete()
}

func (this *PersonController) edit() {
}
