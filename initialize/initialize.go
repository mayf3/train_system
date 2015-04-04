package initialize

import (
//	"fmt"
	"time"
	"html/template"

	"github.com/astaxie/beego"
)

func Attr(s string) template.HTMLAttr{
	return template.HTMLAttr(s)
}

func Safe(s string) template.HTML{
	return template.HTML(s)
}

func EqThenSelected(a, b int) string{
	if a == b{
		return "selected=\"selected\""
	}
	return ""
}

func EqThenChecked(a, b int) string{
	if a == b{
		return "checked=\"true\""
	}
	return ""
}

func FormatTime(t time.Time) string{
	return t.Format("2006-01-02")
}

func init(){
	beego.AddFuncMap("Attr", Attr)
	beego.AddFuncMap("Safe", Safe)

	beego.AddFuncMap("EqThenSelected", EqThenSelected)
	beego.AddFuncMap("EqThenChecked", EqThenChecked)

	beego.AddFuncMap("FormatTime", FormatTime)
}
