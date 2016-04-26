package controllers

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"paybook.com/lite/models"
)

// TransactionController struct
type TransactionController struct {
	beego.Controller
}

// Get ...
func (t *TransactionController) Get() {
	var intLimit, intSkip int
	var strSort string
	var err error

	IDAccount := t.Input().Get("id")
	Limit := t.Input().Get("limit")
	Skip := t.Input().Get("skip")
	Sort := t.Input().Get("sort")

	if len(Sort) > 0 {
		beego.Info(Sort)
		SortSplitted := strings.Split(Sort, " ")
		beego.Info(SortSplitted[0])
		beego.Info(SortSplitted[1])

		if SortSplitted[1] == "DESC" {
			SortSplitted[1] = "-"
		} else {
			SortSplitted[1] = ""
		}

		strSort = SortSplitted[1] + SortSplitted[0]

	}

	if len(Limit) > 0 {
		intLimit, err = strconv.Atoi(Limit)
		if err != nil {
			beego.Error("Error converting Limit to int: ", err)
		}
	}

	if len(Skip) > 0 {
		intSkip, err = strconv.Atoi(Skip)
		if err != nil {
			beego.Error("Error converting Skip to int: ", err)
		}
	}

	if len(IDAccount) == 0 {
		t.Redirect("/", 302)
	}

	if t.GetSession("token") == nil {
		t.Redirect("/", 302)
	}

	transactions := models.Transaction{}
	t.Data["json"] = transactions.Get(IDAccount, strSort, intLimit, intSkip)

	t.ServeJSON()
}

// Count ...
func (t *TransactionController) Count() {
	t.Data["Token"] = t.GetSession("token")
	t.Data["IDAccount"] = t.Ctx.Input.Param(":id_account")
	beego.Info(t.Data["IDAccount"])

	if t.Data["IDAccount"] == nil {
		t.Redirect("/", 302)
	}

	if t.Data["Token"] == nil {
		t.Redirect("/", 302)
	}

	transaction := models.Transaction{}
	t.Data["json"] = transaction.Count(t.Data["IDAccount"].(string))

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

	account := models.Account{}
	t.Data["Account"] = account.GetOne(t.Data["IDAccount"].(string))

	t.Data["Host"] = beego.AppConfig.String("pbsync_files_url")

	t.Layout = "inc/layout.tpl"
	t.TplName = "transactions.tpl"

}
