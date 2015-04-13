package user

import (
	//	"fmt"
	"testing"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	kDataBaseUser     = "iladmin"
	kDataBasePassword = ""
	kDataBaseName     = "test"
)

func init() {
	orm.Debug = true
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", kDataBaseUser+":"+kDataBasePassword+"@/"+kDataBaseName+"?charset=utf8&loc=Asia%2FShanghai")
}

func TestCanRegistered(t *testing.T) {
	username := "abc"
	email := "a@b.com"
	b1, b2 := CanRegistered(username, email)
	username = "bc"
	email = "a@b.com"
	b3, b4 := CanRegistered(username, email)
	username = "abc"
	email = "aa@b.com"
	b5, b6 := CanRegistered(username, email)
	username = "bc"
	email = "aa@b.com"
	b7, b8 := CanRegistered(username, email)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("b1 must be true", func() {
			So(b1, ShouldEqual, true)
		})
		Convey("b2 must be true", func() {
			So(b2, ShouldEqual, true)
		})
		Convey("b3 must be false", func() {
			So(b3, ShouldEqual, false)
		})
		Convey("b4 must be true", func() {
			So(b4, ShouldEqual, true)
		})
		Convey("b5 must be true", func() {
			So(b5, ShouldEqual, true)
		})
		Convey("b6 must be false", func() {
			So(b6, ShouldEqual, false)
		})
		Convey("b7 must be false", func() {
			So(b7, ShouldEqual, false)
		})
		Convey("b8 must be false", func() {
			So(b8, ShouldEqual, false)
		})
	})
}

func TestVerifyPassword(t *testing.T) {
	pwd := "111111"
	salt := "0123456789"
	encoded := GeneratePassword(pwd, salt)
	coded := GeneratePassword(pwd, salt)
	b1 := VerifyPassword(pwd, encoded)
	b2 := VerifyPassword("fake", encoded)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("code must be right", func() {
			So(encoded, ShouldEqual, coded)
		})
		Convey("b1 be true", func() {
			So(b1, ShouldEqual, true)
		})
		Convey("b2 be false", func() {
			So(b2, ShouldEqual, false)
		})
	})
}
