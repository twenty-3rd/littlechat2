package routers

import (
	"littlechat2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.AppController{})
	beego.Router("/join", &controllers.AppController{}, "post:Join")
}
