package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	. "train_system/models"
)

type EditTableController struct {
	beego.Controller
}

type TotalStatus struct{
	Id int
	Name string
	Status []string
}

func (c *EditTableController) Get() {
	var (
		_ error
		tmp string
		problem_list []string
		status []string
		person Person
		all_person []Person
		problem Problem
		all_problem []Problem
		member Member
		all_member []Member
		information Information
		all_information []Information
		table Tables
		total_status []TotalStatus
	)
	tmp = c.Input().Get("table_id")
	c.Data["table_id"] = tmp
	table.Id, _ = strconv.Atoi(tmp)
	_ = table.GetTableById(table.Id)
	for i := 0; i < table.ProblemNumber; i++ {
		problem_list = append(problem_list, string(rune(i + 65)))
	}
	c.Data["problem_list"] = problem_list
	all_information, _ = information.GetInformationByTable(table)
	for _, val := range all_information{
		status = []string{}
		all_member, _ = member.GetMemberByInformation(val)
		all_problem, _ = problem.GetProblemByInformation(val)
		for i := 0; i < len(all_member); i++{
			person.GetPersonById(all_member[i].Person.Id)
			all_person = append(all_person, person)
		}
		fmt.Println(all_problem)
		fmt.Println(all_person)
		for i := 0; i < table.ProblemNumber; i++{
			if all_problem[i].Participant == 0{
				status = append(status, "")
				continue;
			}
			tmp = "";
			if (all_problem[i].Participant & 1) == 1{
				tmp = fmt.Sprintf("%s%s", tmp, all_person[0].Name[0:3])
			}
			if (all_problem[i].Participant & 2) == 2{
				tmp = fmt.Sprintf("%s%s", tmp, all_person[1].Name[0:3])
			}
			if (all_problem[i].Participant & 4) == 4{
				tmp = fmt.Sprintf("%s%s", tmp, all_person[2].Name[0:3])
			}
			if all_problem[i].Status == 2{
				tmp = fmt.Sprintf("%sAC", tmp)
			}
			if all_problem[i].Status != 1{
				tmp = fmt.Sprintf("(%s)", tmp)
			}
			fmt.Println(i)
			fmt.Println(tmp)
			status = append(status, tmp)
		}
		total_status = append(total_status, TotalStatus{val.Id, val.Name, status})
		c.Data["total_status"] = total_status
	}
	c.TplNames = "edit_table.tpl"
}
