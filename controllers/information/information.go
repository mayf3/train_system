package information

import (
	"github.com/mayf3/train_system/controllers/base"
)

type InformationController struct {
	base.BaseController
}

//TODO move this struct
type ProblemStatus struct {
	Name    string
	Member1 int
	Member2 int
	Member3 int
	Status  int
}
