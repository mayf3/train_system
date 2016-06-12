package root

import (
	"github.com/mayf3/train_system/models"
)

func (this *RootController) getAllTable() []models.Tables {
	var (
		err       error
		table     models.Tables
		all_table []models.Tables
	)
	all_table, err = table.GetAllTable()
	if err != nil {
		return nil
	}
	return all_table
}
