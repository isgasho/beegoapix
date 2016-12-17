package goxapi

import (
	"github.com/luffyke/goxapi/controllers"

	"github.com/astaxie/beego"
)

const (
	VERSION = "0.1.0"
)

func Router() {
	beego.Router("/v1/:controller/:method", &controllers.BaseController{}, "*:Post")
}
