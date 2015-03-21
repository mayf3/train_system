package tools

import (
	"fmt"
	. "train_system/models"
)

type TotalStatus struct{
	Id int
	Name string
	Status []string
	MemberName []string
}

func GenerateProblemList(table_id int) []string{
	var(
		_ error
		table Tables
		problem_list []string
	)
	_ = table.GetTableById(table_id)
	for i := 0; i < table.ProblemNumber; i++ {
		problem_list = append(problem_list, string(rune(i + 65)))
	}
	return problem_list
}

func GenerateTable(table_id int) []TotalStatus{
	var(
		_ error
		tmp string
		status []string
		member_name []string
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
	_ = table.GetTableById(table_id)
	all_information, _ = information.GetInformationByTable(table)
	for _, val := range all_information{
		status = []string{}
		member_name = []string{}
		all_person = []Person{}
		all_member, _ = member.GetMemberByInformation(val)
		all_problem, _ = problem.GetProblemByInformation(val)
		for i := 0; i < len(all_member); i++{
			person.GetPersonById(all_member[i].Person.Id)
			all_person = append(all_person, person)
			member_name = append(member_name, person.Name)
		}
		for i := 0; i < table.ProblemNumber; i++{
			if all_problem[i].Participant == 0{
				status = append(status, "")
				continue;
			}
			tmp = "";
			if (all_problem[i].Participant & 1) == 1{
				tmp = fmt.Sprintf("%s%s ", tmp, all_person[0].Name)
			}
			if (all_problem[i].Participant & 2) == 2{
				tmp = fmt.Sprintf("%s%s ", tmp, all_person[1].Name)
			}
			if (all_problem[i].Participant & 4) == 4{
				tmp = fmt.Sprintf("%s%s ", tmp, all_person[2].Name)
			}
			if all_problem[i].Status == 2{
				tmp = fmt.Sprintf("%sAC", tmp)
			}
			if all_problem[i].Status != 1{
				tmp = fmt.Sprintf("(%s)", tmp)
			}
			status = append(status, tmp)
		}
		total_status = append(total_status, TotalStatus{val.Id, val.Name, status, member_name})
	}
	return total_status
}
