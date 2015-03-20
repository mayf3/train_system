package controllers

import (
	"fmt"
	"strconv"
	"github.com/astaxie/beego"
	. "train_system/models"
)

type EditInformationController struct {
	beego.Controller
}

//TODO move this struct 
type ProblemStatus struct{
	Name string
	Member1 int
	Member2 int
	Member3 int
	Status int
}

func (c *EditInformationController) Get() {
	var(
		id int
		err error
		table Tables
		information Information
		member Member
		all_member []Member
		problem Problem
		all_problem []Problem
		problem_status ProblemStatus
		problem_list []ProblemStatus
		person Person
		all_person []Person
	)
	all_person, err = person.GetAllPerson()
	fmt.Println(all_person)
	c.Data["all_person"] = all_person

	table_id := c.Input().Get("table_id")
	id, err = strconv.Atoi(table_id)
	if err == nil && id != 0{
		err = table.GetTableById(id)
		for i := 0; i < table.ProblemNumber; i++ {
			problem_status.Name = string(rune(i + 65))
			problem_list = append(problem_list, problem_status)
		}
	}

	information_id := c.Input().Get("information_id")
	id, err = strconv.Atoi(information_id)
	if err == nil && id != 0{
		err = information.GetInformationById(id)
		c.Data["Init"] = information
		all_member, err = member.GetMemberByInformation(information)
		for key, val := range all_member{
			c.Data["Member" + string(rune(key + 49))] = val.Person.Id
		}
		all_problem, err = problem.GetProblemByInformation(information)
		for _, val := range all_problem{
			problem_list[val.Number].Status = val.Status
			if (val.Participant & 1) == 1{
				problem_list[val.Number].Member1 = 1
			}
			if (val.Participant & 2) == 2{
				problem_list[val.Number].Member2 = 1
			}
			if (val.Participant & 4) == 4{
				problem_list[val.Number].Member3 = 1
			}
		}
	}
	for i := len(all_member); i < 3; i++{
		c.Data["Member" + string(rune(i + 49))] = 0
	}
	fmt.Println(c.Data)
	c.Data["problem_list"] = problem_list
	c.TplNames = "edit_information.tpl"
}

func (c *EditInformationController) Post() {
	var(
		err error
		num int
		table_id int
		information_id int
		person_id int
		information Information
		table Tables
		member Member
		person Person
	)
	information_id, err = strconv.Atoi(c.GetString("information"))
	if err == nil && information_id > 0{
		err = information.GetInformationById(information_id)
	}
	table_id, err = strconv.Atoi(c.GetString("table_id"))
	err = table.GetTableById(table_id)
	information.Table = &table
	information.Rank, err = strconv.Atoi(c.GetString("rank"))
	information.Name = c.GetString("name")
	if err == nil && information_id > 0{
		err = information.Update()
	} else{
		err = information.Insert()
	}
	for i = 1; i <= 3; i++{
		person_id, err = strconv.Atoi(c.GetString(fmt.Printf("%s%d", "member", i)))
		if err == nil && person > 0{
			err = person.GetPersonById(person_id)
			err = member.GetMemberByInformationAndOrder(information, i)
			if err == nil{
				member.Person = person
				member.Update()
			} else{
				member = Member{}
				member.Order = i
				member.Person = person
				member.Information = information
				member.Insert()
			}
		}
	}
	for i = 0; i < table.ProblemNumber; i++{
		num, err = strconv.Atoi(c.GetString(fmt.Sprintf("%s%d", "problem_participant" , i)))
		if err == nil{
			if information_id > 0{
				problem = problem.Get
			}
			else{
			}
		}
	}
	c.Redirect("/index", 302)
}
