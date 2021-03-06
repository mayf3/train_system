package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"

	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/orm"
	. "github.com/smartystreets/goconvey/convey"

	_ "train_system/routers"
	_ "train_system/initialize"
//	. "train_system/models"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}


//TestMain is a sample to run an endpoint test
func Tmp(t *testing.T) {
	r, _ := http.NewRequest("GET", "/table/4/information/create", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	        Convey("The Result Should Not Be Empty", func() {
	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
	        })
	})
}

func TestRouter(t *testing.T){
	fmt.Println(beego.UrlFor("InformationController.Edit", ":information_id", 3, ":table_id", 4))
	fmt.Println(beego.UrlFor("TableController.Show", ":table_id", 4))
	fmt.Println(beego.UrlFor("RootController.Index"))
}
