package table

import (
	"errors"
	"fmt"
	"strings"

	"train_system/models"
)

// utils

func (this *TableController) getTableId() (table_id int, err error) {
	table_id, err = this.GetInt(":table_id")
	if err != nil || table_id == 0 {
		return 0, errors.New("no table id")
	}
	return table_id, nil
}

func (this *TableController) getTable() models.Tables {
	var (
		_     error
		table models.Tables
	)
	table.Id, _ = this.getTableId()
	_ = table.GetTableById(table.Id)
	return table
}

func (this *TableController) mustTable() {
	var (
		err   error
		table models.Tables
	)
	table.Id, err = this.getTableId()
	if err != nil {
		this.Abort("Error table id")
	}
	err = table.GetTableById(table.Id)
	if err != nil {
		this.Abort("Error table id")
	}
}

func (this *TableController) checkTable() {
	var (
		err   error
		table models.Tables
	)
	table.Id, err = this.getTableId()
	if err != nil {
		return
	}
	err = table.GetTableById(table.Id)
	if err != nil {
		this.Abort("Error table id")
	}
}

func (this *TableController) getProblemList() (ret []string) {
	ret = append(ret, "")
	for i := 0; i < 26; i++ {
		ret = append(ret, fmt.Sprintf("%s-%d", string(rune(i+65)), i+1))
	}
	return ret
}

func (this *TableController) splitHustTable(hust models.Hust) [][]string {
	var (
		hust_table     []string
		all_hust_table [][]string
	)
	hust_table = strings.Split(hust.Context, string(rune(32)))
	for _, val := range hust_table {
		all_hust_table = append(all_hust_table, strings.Split(val, string(rune(9))))
	}
	for i := 1; i < len(all_hust_table); i++ {
		all_hust_table[i] = all_hust_table[i][0 : len(all_hust_table[i])-1]
	}
	return all_hust_table
}

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
		table        models.Tables
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
		person             models.Person
		all_person         []models.Person
		problem            models.Problem
		all_problem        []models.Problem
		member             models.Member
		all_member         []models.Member
		information        models.Information
		all_information    []models.Information
		table              models.Tables
		total_status       []TotalStatus
	)
	_ = table.GetTableById(table_id)
	all_information, _ = information.GetInformationByTable(table)
	for _, val := range all_information {
		show_status = []ShowStatus{}
		member_name = []string{}
		all_person = []models.Person{}
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
