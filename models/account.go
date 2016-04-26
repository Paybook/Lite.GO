package models

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"paybook.com/lite/services"
)

// Account ...
type Account struct {
	IDAccount  string  `orm:"column(id_account);pk;unique"`
	IDUser     string  `orm:"column(id_user)"`
	IDExternal string  `orm:"column(id_external)"`
	IDSite     string  `orm:"column(id_site)"`
	ISDisable  bool    `orm:"column(is_disable);bool"`
	Name       string  `orm:"column(name)"`
	Number     string  `orm:"column(number)"`
	Balance    float64 `orm:"column(balance);real"`
	SiteName   string  `orm:"column(site_name))"`
	SiteAvatar string  `orm:"column(site_avatar))"`
	DTCreate   string  `orm:"column(dt_create);datetime)"`
	DTModify   string  `orm:"column(dt_modify);datetime)"`
}

// TableName ...
func (a *Account) TableName() string {
	return "accounts"
}

// GetAPI ...
func (a *Account) GetAPI(token string) {

	url := beego.AppConfig.String("pbsync_base_url") + "/accounts?token=" + token

	services := services.Services{}
	res := services.Get(url)
	api := API{}
	err := json.Unmarshal([]byte(res), &api)
	if err != nil {
		beego.Error("Error parsing API post output: ", err, res)
	}
	beego.Debug(string(api.Response))

	accountsAPI := []AccountsAPI{}
	err = json.Unmarshal([]byte(api.Response), &accountsAPI)
	if err != nil {
		beego.Error("Error parsing AccountsAPI post output: ", err, api.Response)
	}
	beego.Debug(accountsAPI)

	for _, value := range accountsAPI {

		account := new(Account)
		account.IDAccount = value.IDAccount
		account.IDUser = value.IDUser
		account.IDExternal = value.IDExternal
		account.IDSite = value.IDSite
		account.Name = value.Name
		account.Number = value.Number
		account.Balance = value.Balance
		account.SiteName = value.Site.Name
		account.SiteAvatar = value.Site.Avatar

		o := orm.NewOrm()
		id, err := o.Insert(account)
		if err != nil {
			beego.Error("Error on inserting account: ", err, id, account)
		}
	}
}

// Get ...
func (a *Account) Get() []Account {
	o := orm.NewOrm()
	accounts := []Account{}
	qs, err := o.QueryTable("accounts").All(&accounts)
	if err == nil {
		beego.Error("Error reading accounts: ", err, qs)
	}

	return accounts
}

// GetOne ...
func (a *Account) GetOne(IdAccount string) Account {
	o := orm.NewOrm()
	account := Account{IDAccount: IdAccount}

	err := o.Read(&account)
	if err == nil {
		beego.Error("Account not found: ", err, IdAccount)
	}

	return account
}
