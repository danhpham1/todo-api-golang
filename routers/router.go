// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/getsentry/sentry-go"
	"todo-api/controllers"
)

func init() {
	sentryDNS, _ := beego.AppConfig.String("sentryDNS")
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:   sentryDNS,
		Debug: true,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	ns := beego.NewNamespace("/v1",
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
		beego.NSNamespace("/todo",
			beego.NSInclude(
				&controllers.TodoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
