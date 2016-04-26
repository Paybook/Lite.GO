package models

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"paybook.com/lite/services"
)

// Transaction ...
type Transaction struct {
	IDTransaction          string  `orm:"column(id_transaction);pk;unique"`
	IDUser                 string  `orm:"column(id_user)"`
	IDExternal             string  `orm:"column(id_external)"`
	IDSite                 string  `orm:"column(id_site)"`
	IDSiteOrganization     string  `orm:"column(id_site_organiaztion)"`
	IDSiteOrganizationType string  `orm:"column(id_site_organiaztion_type)"`
	IDAccount              string  `orm:"column(id_account)"`
	IDAccountType          string  `orm:"column(id_account_type)"`
	Description            string  `orm:"column(description)"`
	DTTransaction          int64   `orm:"column(dt_transaction);bigint"`
	Period                 int64   `orm:"column(period);integer"`
	Amount                 float64 `orm:"column(amount);real"`
	DTCreate               string  `orm:"column(dt_create);bigint"`
	DTModify               string  `orm:"column(dt_modify);bigint"`
}

// TableName ...
func (t *Transaction) TableName() string {
	return "transactions"
}

// GetAPI ...
func (t *Transaction) GetAPI(token string) {

	url := beego.AppConfig.String("pbsync_base_url") + "/transactions?token=" + token

	services := services.Services{}
	res := services.Get(url)
	api := API{}
	err := json.Unmarshal([]byte(res), &api)
	if err != nil {
		beego.Error("Error parsing API post output: ", err, res)
	}
	beego.Debug(string(api.Response))

	transactionsAPI := []TransactionsAPI{}
	err = json.Unmarshal([]byte(api.Response), &transactionsAPI)
	if err != nil {
		beego.Error("Error parsing TransactionsAPI post output: ", err, api.Response)
	}
	beego.Debug(transactionsAPI)

	for _, value := range transactionsAPI {

		transaction := new(Transaction)
		transaction.IDTransaction = value.IDTransaction
		transaction.IDAccount = value.IDAccount
		transaction.IDUser = value.IDUser
		transaction.IDExternal = value.IDExternal
		transaction.IDSite = value.IDSite
		transaction.IDSiteOrganization = value.IDSiteOrganization
		transaction.IDSiteOrganizationType = value.IDSiteOrganizationType
		transaction.IDAccountType = value.IDAccountType
		transaction.Description = value.Description
		transaction.Period = value.Period
		transaction.Amount = value.Amount

		o := orm.NewOrm()
		id, err := o.Insert(transaction)
		if err != nil {
			beego.Error("Error on inserting transaction: ", err, id, transaction)
		}
	}
}

// Get ...
func (t *Transaction) Get(IDAccount string) []Transaction {
	o := orm.NewOrm()
	transactions := []Transaction{}

	qs, err := o.QueryTable("transactions").Filter("id_account", IDAccount).All(&transactions)
	if err == nil {
		beego.Error("Error reading transactions: ", err, qs)
	}

	return transactions
}
