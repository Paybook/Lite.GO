package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"paybook.com/lite/models"
	"paybook.com/lite/services"
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

	err, IDUser := u.Auth()

	if err {

		api := models.API{}
		services := services.Services{}
		url := beego.AppConfig.String("pbsync_base_url") +
			"sessions?api_key=" + beego.AppConfig.String("pbsync_api_key") +
			"&id_user=" + string(IDUser)
		res := services.Post(url)
		err := json.Unmarshal([]byte(res), &api)
		if err != nil {
			beego.Info("Error parsing API post output: %s - %s", err, res)
		}

		if api.Code == 200 {
			tokenAPI := models.TokenAPI{}
			err = json.Unmarshal([]byte(api.Response), &tokenAPI)
			if err != nil {
				beego.Info("Error parsing tokenAPI response: %s - %s", err, res)
			}

			beego.Info("Setting token: ", tokenAPI.Token)
			c.SetSession("token", tokenAPI.Token)
			c.SetSession("paybook-lite", true)

			account := models.Account{}
			account.GetAPI(tokenAPI.Token)

			transaction := models.Transaction{}
			transaction.GetAPI(tokenAPI.Token)

			c.Redirect("/dashboard", 302)

		} else {
			beego.Error("Error: on create session on pbsync")
		}

	} else {
		c.Redirect("/", 302)
	}
}

// Dashboard ...
func (c *UsersController) Dashboard() {
	c.Data["Token"] = c.GetSession("token")

	if c.Data["Token"] == nil {
		c.Redirect("/", 302)
	}

	accounts := models.Account{}
	c.Data["Accounts"] = accounts.Get()

	c.Data["Host"] = beego.AppConfig.String("pbsync_files_url")

	c.Layout = "inc/layout.tpl"
	c.TplName = "dashboard.tpl"
}

// Logout ...
func (c *UsersController) Logout() {
	c.DelSession("paybook-lite")

	c.Redirect("/", 302)
}
