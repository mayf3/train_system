package information

import (
	"fmt"
	"strconv"
	//	"errors"

	"github.com/mayf3/train_system/models"
)

// getPreInformation
func (this *InformationController) getPreInformation() models.Information {
	var (
		err         error
		information models.Information
	)
	information = this.getInformation()
	if err != nil {
		return models.Information{}
	}
	err = information.GetInformationById(information.Id)
	if err != nil {
		//TODO should we abort ?
		this.Abort("Error information id")
	}
	return information
}

// getAllPerson
func (this *InformationController) getAllPerson() []models.Person {
	var (
		err        error
		person     models.Person
		all_person []models.Person
	)
	all_person, err = person.GetAllPersonOrderByGrade()
	if err != nil {
		return nil
	}
	return all_person
}

// getMember
func (this *InformationController) getMember() []int {
	var (
		_            error
		show_members []int
		information  models.Information
		member       models.Member
		members      []models.Member
	)
	information = this.getInformation()
	members, _ = member.GetMemberByInformation(information)
	for _, val := range members {
		show_members = append(show_members, val.Person.Id)
	}
	for i := len(members); i < 3; i++ {
		show_members = append(show_members, 0)
	}
	return show_members
}

func (this *InformationController) getProblemStatus() []ProblemStatus {
	var (
		_                  error
		err                  error
		table              models.Tables
		information        models.Information
		problem            models.Problem
		all_problem        []models.Problem
		problem_status     ProblemStatus
		all_problem_status []ProblemStatus
	)
	table.Id, _ = this.getTableId()
	_ = table.GetTableById(table.Id)
	information.Id, err = this.getInformationId()
	if err != nil{
		for i := 0; i < table.ProblemNumber; i++ {
			problem_status.Name = string(rune(i + 65))
			problem_status.Member1 = 0
			problem_status.Member2 = 0
			problem_status.Member3 = 0
			problem_status.Status = 0
			all_problem_status = append(all_problem_status, problem_status)
		}
		return all_problem_status
	}
	information = this.getInformation()
	for i := 0; i < table.ProblemNumber; i++ {
		problem_status.Name = string(rune(i + 65))
		all_problem_status = append(all_problem_status, problem_status)
	}
	all_problem, _ = problem.GetProblemByInformation(information)
	for _, val := range all_problem {
		all_problem_status[val.Number].Status = val.Status
		if (val.Participant & 1) == 1 {
			all_problem_status[val.Number].Member1 = 1
		}
		if (val.Participant & 2) == 2 {
			all_problem_status[val.Number].Member2 = 1
		}
		if (val.Participant & 4) == 4 {
			all_problem_status[val.Number].Member3 = 1
		}
	}
	return all_problem_status
}

func (this *InformationController) submitInformation() {
	var (
		err            error
		table_id       int
		information_id int
		person_id      int
		information    models.Information
		table          models.Tables
		member         models.Member
		person         models.Person
		problem        models.Problem
	)
	table_id, err = this.getTableId()
	err = table.GetTableById(table_id)
	information_id, err = this.getInformationId()
	if err == nil && information_id > 0 {
		err = information.GetInformationById(information_id)
	}
	information.Table = &table
	information.Rank, err = strconv.Atoi(this.GetString("rank"))
	information.Name = this.GetString("name")
	if information_id > 0 {
		err = information.Update()
	} else {
		err = information.Insert()
	}
	for i := 0; i < 3; i++ {
		person_id, err = strconv.Atoi(this.GetString(fmt.Sprintf("%s%d", "member", i)))
		if err == nil && person_id > 0 {
			err = person.GetPersonById(person_id)
			err = member.GetMemberByInformationAndOrder(information, i)
			if err == nil {
				member.Person = &person
				member.Update()
			} else {
				member = models.Member{}
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
		strings = this.GetStrings(fmt.Sprintf("problem_participant%d[]", i))
		for _, val := range strings {
			num, err = strconv.Atoi(val)
			sum = sum + num
		}
		status, err = strconv.Atoi(this.GetString(fmt.Sprintf("problem_status%d", i)))
		if information_id > 0 {
			err = problem.GetProblemByInformationAndNumber(information, i)
			problem.Participant = sum
			problem.Status = status
			problem.Update()
		} else {
			problem = models.Problem{}
			problem.Number = i
			problem.Participant = sum
			problem.Status = status
			problem.Information = &information
			problem.Insert()
		}
	}
	this.Redirect(fmt.Sprintf("/table/%d/edit", table.Id), 302)
}
