package controllers

import (
	"github.com/astaxie/beego"
	"paybook.com/lite/models"
)

// UsersController definition
type UsersController struct {
	beego.Controller
}

// Post Users
func (c *UsersController) Post() {
	u := models.User{}

	if err := c.ParseForm(&u); err != nil {
		beego.Debug("User: ", u.Email, u.Password)
	}

	u.Create()
	c.Redirect("/", 302)
}

// Login Users
func (c *UsersController) Login() {
	u := models.User{}

	if err := c.ParseForm(&u); err != nil {
		beego.Debug("User: ", u.Email, u.Password)
	}

	if u.Auth() {
		c.Layout = "inc/layout.tpl"
		c.TplName = "dashboard.tpl"
		c.SetSession("paybook-lite", true)
	} else {
		c.Redirect("/", 302)
	}
}

// Logout ...
func (c *UsersController) Logout() {
	c.DelSession("paybook-lite")

	c.Redirect("/", 302)
}
