// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"passapp-engine-api/controllers"
	"github.com/astaxie/beego"
	"passapp-engine-api/handlers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/jwt",
			beego.NSInclude(
				&controllers.JWTController{},
			),
		),
		beego.NSBefore(handlers.Jwt),
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/access",
			beego.NSInclude(
				&controllers.AccessController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
