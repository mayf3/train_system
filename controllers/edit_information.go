package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	. "train_system/models"
)

type EditInformationController struct {
	beego.Controller
}

//TODO move this struct
type ProblemStatus struct {
	Name    string
	Member1 int
	Member2 int
	Member3 int
	Status  int
}

func (c *EditInformationController) Get() {
	var (
		id             int
		err            error
		table          Tables
		information    Information
		show_members   []int
		member         Member
		all_member     []Member
		problem        Problem
		all_problem    []Problem
		problem_status ProblemStatus
		problem_list   []ProblemStatus
		person         Person
		all_person     []Person
	)
	all_person, err = person.GetAllPersonOrderByGrade()
	fmt.Println(all_person)
	c.Data["all_person"] = all_person

	table_id := c.Input().Get("table_id")
	id, err = strconv.Atoi(table_id)
	if err == nil && id != 0 {
		err = table.GetTableById(id)
		for i := 0; i < table.ProblemNumber; i++ {
			problem_status.Name = string(rune(i + 65))
			problem_list = append(problem_list, problem_status)
		}
	}

	information_id := c.Input().Get("information_id")
	id, err = strconv.Atoi(information_id)
	if err == nil && id != 0 {
		err = information.GetInformationById(id)
		c.Data["Init"] = information
		all_member, err = member.GetMemberByInformation(information)
		for _, val := range all_member {
			show_members = append(show_members, val.Person.Id)
		}
		all_problem, err = problem.GetProblemByInformation(information)
		for _, val := range all_problem {
			problem_list[val.Number].Status = val.Status
			if (val.Participant & 1) == 1 {
				problem_list[val.Number].Member1 = 1
			}
			if (val.Participant & 2) == 2 {
				problem_list[val.Number].Member2 = 1
			}
			if (val.Participant & 4) == 4 {
				problem_list[val.Number].Member3 = 1
			}
		}
	}
	for i := len(all_member); i < 3; i++ {
		show_members = append(show_members, 0)
	}
	c.Data["show_members"] = show_members
	fmt.Println(c.Data)
	c.Data["problem_list"] = problem_list
	c.TplNames = "edit_information.tpl"
}

func (c *EditInformationController) Post() {
	var (
		err            error
		table_id       int
		information_id int
		person_id      int
		information    Information
		table          Tables
		member         Member
		person         Person
		problem        Problem
	)
	information_id, err = strconv.Atoi(c.GetString("information_id"))
	if err == nil && information_id > 0 {
		err = information.GetInformationById(information_id)
	}
	table_id, err = strconv.Atoi(c.GetString("table_id"))
	err = table.GetTableById(table_id)
	information.Table = &table
	information.Rank, err = strconv.Atoi(c.GetString("rank"))
	information.Name = c.GetString("name")
	if err == nil && information_id > 0 {
		err = information.Update()
	} else {
		err = information.Insert()
	}
	for i := 0; i < 3; i++ {
		person_id, err = strconv.Atoi(c.GetString(fmt.Sprintf("%s%d", "member", i)))
		if err == nil && person_id > 0 {
			err = person.GetPersonById(person_id)
			err = member.GetMemberByInformationAndOrder(information, i)
			if err == nil {
				member.Person = &person
				member.Update()
			} else {
				member = Member{}
				member.Order = i
				member.Person = &person
				member.Information = &information
				member.Insert()
			}
		}
	}
	for i := 0; i < table.ProblemNumber; i++ {
		var (
			sum     int
			num     int
			status  int
			strings []string
		)
		sum = 0
		strings = c.GetStrings(fmt.Sprintf("problem_participant%d[]", i))
		for _, val := range strings {
			num, err = strconv.Atoi(val)
			sum = sum + num
		}
		status, err = strconv.Atoi(c.GetString(fmt.Sprintf("problem_status%d", i)))
		if information_id > 0 {
			err = problem.GetProblemByInformationAndNumber(information, i)
			problem.Participant = sum
			problem.Status = status
			problem.Update()
		} else {
			problem = Problem{}
			problem.Number = i
			problem.Participant = sum
			problem.Status = status
			problem.Information = &information
			problem.Insert()
		}
	}
	c.Redirect(fmt.Sprintf("/edit_table?table_id=%d", table.Id), 302)
}
