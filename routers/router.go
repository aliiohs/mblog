package routers

import (
	"github.com/astaxie/beego"
	"github.com/mblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{})
}
