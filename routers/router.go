package routers

import (
	"github.com/astaxie/beego"
	"paybook.com/lite/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UsersController{}, "post:Login")
	beego.SetStaticPath("/static", "static")
	beego.SetStaticPath("/img", "img")
	beego.SetStaticPath("/css", "css")
	beego.SetStaticPath("/js", "js")
	beego.Router("/session", &controllers.SessionController{})
	beego.Router("/signup", &controllers.MainController{}, "get:SignUp")
	beego.Router("/users", &controllers.UsersController{})
	beego.Router("/logout", &controllers.UsersController{}, "get:Logout")
}
