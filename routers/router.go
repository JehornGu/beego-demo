// @APIVersion 1.0.0
// @Title CIV API
// @Description CIV APIs.
// @Contact jehornguu@gmail.com
package routers

import (
	"civ/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
