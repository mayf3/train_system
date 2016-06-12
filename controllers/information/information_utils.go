package information

import (
	"errors"

	"github.com/mayf3/train_system/models"
)

func (this *InformationController) getTableId() (table_id int, err error) {
	table_id, err = this.GetInt(":table_id")
	if err != nil {
		return 0, errors.New("no table id")
	}
	return table_id, nil
}

func (this *InformationController) getInformationId() (information_id int, err error) {
	information_id, err = this.GetInt(":information_id")
	if err != nil || information_id == 0{
		return 0, errors.New("no information id")
	}
	return information_id, nil
}

func (this *InformationController) getInformation() models.Information {
	var (
		information models.Information
	)
	information.Id, _ = this.getInformationId()
	_ = information.GetInformationById(information.Id)
	return information
}

func (this *InformationController) checkTable() {
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

func (this *InformationController) checkInformation() bool {
	var (
		err         error
		information models.Information
	)
	information.Id, err = this.getInformationId()
	if err != nil  || information.Id == 0{
		return false
	}
	err = information.GetInformationById(information.Id)
	if err != nil {
		this.Abort("Error information id")
	}
	return true
}
