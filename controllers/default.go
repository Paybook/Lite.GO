package controllers

import (
	"github.com/astaxie/beego"
)

// MainController struct
type MainController struct {
	beego.Controller
}

// Get ...
func (c *MainController) Get() {
	c.Layout = "inc/layout.tpl"
	c.TplName = "login.tpl"
}

// SignUp ...
func (c *MainController) SignUp() {
	c.Layout = "inc/layout.tpl"
	c.TplName = "signup.tpl"
}
