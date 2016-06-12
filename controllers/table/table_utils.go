package table
import(
	"fmt"
	"errors"
	"strings"

	"github.com/mayf3/train_system/models"
)

// utils

func (this *TableController) getTableId() (table_id int, err error) {
	table_id, err = this.GetInt(":table_id")
	if err != nil || table_id == 0{
		return 0, errors.New("no table id")
	}
	return table_id, nil
}

func (this *TableController) getTable() models.Tables{
	var(
		_ error
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

func (this *TableController) getProblemList() (ret []string){
	ret = append(ret, "")
	for i := 0; i < 26; i++{
		ret = append(ret, fmt.Sprintf("%s-%d", string(rune(i + 65)), i + 1))
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
