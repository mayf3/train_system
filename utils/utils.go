package utils

import (
	"fmt"
	. "train_system/models"
)

type ShowStatus struct {
	Context string
	Colour  string
}

type TotalStatus struct {
	Id         int
	Name       string
	Status     []ShowStatus
	MemberName []string
}

func GenerateProblemList(table_id int) []string {
	var (
		_            error
		table        Tables
		problem_list []string
	)
	_ = table.GetTableById(table_id)
	for i := 0; i < table.ProblemNumber; i++ {
		problem_list = append(problem_list, string(rune(i+65)))
	}
	return problem_list
}

func GenerateTable(table_id int) []TotalStatus {
	var (
		_                  error
		tmp                string
		member_name        []string
		single_show_status ShowStatus
		show_status        []ShowStatus
		person             Person
		all_person         []Person
		problem            Problem
		all_problem        []Problem
		member             Member
		all_member         []Member
		information        Information
		all_information    []Information
		table              Tables
		total_status       []TotalStatus
	)
	_ = table.GetTableById(table_id)
	all_information, _ = information.GetInformationByTable(table)
	for _, val := range all_information {
		show_status = []ShowStatus{}
		member_name = []string{}
		all_person = []Person{}
		all_member, _ = member.GetMemberByInformation(val)
		all_problem, _ = problem.GetProblemByInformation(val)
		for i := 0; i < len(all_member); i++ {
			person.GetPersonById(all_member[i].Person.Id)
			all_person = append(all_person, person)
			member_name = append(member_name, person.Name)
		}
		for i := 0; i < table.ProblemNumber; i++ {
			switch all_problem[i].Status {
			case 0:
				single_show_status.Colour = "text-muted"
			case 1:
				single_show_status.Colour = "text-success"
			case 2:
				single_show_status.Colour = "text-danger"
			}
			if all_problem[i].Participant == 0 {
				single_show_status.Context = ""
				show_status = append(show_status, single_show_status)
				continue
			}
			tmp = ""
			if (all_problem[i].Participant & 1) == 1 {
				tmp = fmt.Sprintf("%s%s ", tmp, all_person[0].Name)
			}
			if (all_problem[i].Participant & 2) == 2 {
				tmp = fmt.Sprintf("%s%s ", tmp, all_person[1].Name)
			}
			if (all_problem[i].Participant & 4) == 4 {
				tmp = fmt.Sprintf("%s%s ", tmp, all_person[2].Name)
			}
			if all_problem[i].Status == 2 {
				tmp = fmt.Sprintf("%sAC", tmp)
			}
			if all_problem[i].Status != 1 {
				tmp = fmt.Sprintf("(%s)", tmp)
			}
			single_show_status.Context = tmp
			show_status = append(show_status, single_show_status)
		}
		total_status = append(total_status, TotalStatus{val.Id, val.Name, show_status, member_name})
	}
	return total_status
}
