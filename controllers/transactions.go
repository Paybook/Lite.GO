package controllers

import (
	"github.com/astaxie/beego"
	"paybook.com/lite/models"
)

// TransactionController struct
type TransactionController struct {
	beego.Controller
}

// Get ...
func (t *TransactionController) Get() {
	var IDAccount string
	// t.Data["IDAccount"] = t.Ctx.Input.Param(":id_account")
	// beego.Info("id_account", t.Data["IDAccount"])

	// if t.Data["IDAccount"] == nil {
	// 	t.Redirect("/", 302)
	// }

	if t.GetSession("token") == nil {
		t.Redirect("/", 302)
	}

	request := t.Ctx.Request
	length := request.ContentLength
	p := make([]byte, length)
	bytesRead, err := t.Ctx.Request.Body.Read(p)
	if err == nil {
		beego.Error("Error on getting info: ", err)
	} else {
		beego.Info("Ginfo: ", bytesRead, length)
	}

	transactions := models.Transaction{}
	t.Ctx.Input.Bind(&IDAccount, ":id_account")
	t.Data["json"] = transactions.Get(IDAccount)

	t.ServeJSON()
}

// View ...
func (t *TransactionController) View() {
	t.Data["Token"] = t.GetSession("token")
	t.Data["IDAccount"] = t.Ctx.Input.Param(":id_account")
	beego.Info(t.Data["IDAccount"])

	if t.Data["IDAccount"] == nil {
		t.Redirect("/", 302)
	}

	if t.Data["Token"] == nil {
		t.Redirect("/", 302)
	}

	t.Layout = "inc/layout.tpl"
	t.TplName = "transactions.tpl"

}
