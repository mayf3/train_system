package models

import (
	"github.com/astaxie/beego/orm"

	"train_system/modules/utils"
	"train_system/setting"
)

type User struct {
	Id       int    `orm:"pk;auto"`
	Username string `orm:"unique"`
	Email    string `orm:"unique"`
	Score    int
	Verify   bool
	Nickname string
	Password string
	Power    string
	Rands    string
}

func init() {
	orm.RegisterModel(new(User))
}

// All

// Single
func (this *User) GetUserByUsername(username string) error {
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("Username", username).One(this)
	return err
}

// Tool
func UpdatePassword(username, password string) error {
	o := orm.NewOrm()
	var single User
	err := o.QueryTable("user").Filter("Username", username).One(&single)
	if err != nil {
		return err
	}
	single.Password = password
	err = single.Update()
	return err
}

func GetRandsByUsername(username string) string {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.QueryTable("user").Filter("Username", username).Values(&maps, "Rands")
	if err != nil || num == 0 {
		return ""
	}
	return maps[0]["Rands"].(string)
}

func GetPasswordByUsername(username string) string {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.QueryTable("user").Filter("Username", username).Values(&maps, "Password")
	if err != nil || num == 0 {
		return ""
	}
	return maps[0]["Password"].(string)
}

func GetPowerByUsername(username string) string {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.QueryTable("user").Filter("Username", username).Values(&maps, "Power")
	if err != nil || num == 0 {
		return ""
	}
	return maps[0]["Power"].(string)
}

func HasUserByUsername(username string) bool {
	o := orm.NewOrm()
	return o.QueryTable("user").Filter("Username", username).Exist()
}

func HasUserByEmail(email string) bool {
	o := orm.NewOrm()
	return o.QueryTable("user").Filter("Email", email).Exist()
}

// Other
func (this *User) SetVerify() error {
	this.Verify = true
	return this.Update()
}

// Base
func (this *User) Insert() error {
	this.Rands = utils.GetRandomString(setting.SaltLen)
	this.Verify = false
	this.Score = 0
	o := orm.NewOrm()
	_, err := o.Insert(this)
	return err
}

func (this *User) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}

func (this *User) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}
