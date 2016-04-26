package models

import "encoding/json"

// API of models
type API struct {
	Code     int             `json:"code"`
	Status   bool            `json:"status"`
	Message  string          `json:"message"`
	Response json.RawMessage `json:"response"`
}

// UserAPI ...
type UserAPI struct {
	IDUser string `json:"id_user"`
	Name   string `json:"name"`
}

// TokenAPI ...
type TokenAPI struct {
	Token string `json:"token"`
}

// SiteAPI ...
type SiteAPI struct {
	IDSite    string `json:"id_site""`
	Avatar    string `json:"avatar"`
	Cover     string `json:"cover"`
	Name      string `json:"name"`
	SmalCover string `json:"small_cover"`
}

// AccountsAPI ...
type AccountsAPI struct {
	IDAccount  string  `json:"id_account"`
	IDUser     string  `json:"id_user"`
	IDExternal string  `json:"id_external"`
	IDSite     string  `json:"id_site"`
	Site       SiteAPI `json:"site"`
	ISDisable  bool    `json:"is_disable"`
	Name       string  `json:"name"`
	Number     string  `json:"number"`
	Balance    float64 `json:"balance"`
	DTCreate   string  `json:"dt_create"`
	DTModify   string  `json:"dt_modify"`
}

// TransactionsAPI ...
type TransactionsAPI struct {
	IDTransaction          string  `json:"id_transaction"`
	IDUser                 string  `json:"id_user"`
	IDExternal             string  `json:"id_external"`
	IDSite                 string  `json:"id_site"`
	IDSiteOrganization     string  `json:"id_site_organization"`
	IDSiteOrganizationType string  `json:"id_site_organization_type"`
	IDAccount              string  `json:"id_account"`
	IDAccountType          string  `json:"id_account_type"`
	Description            string  `json:"description"`
	DTTransaction          int64   `json:"dt_transaction"`
	Period                 int64   `json:"period"`
	Amount                 float64 `json:"amount"`
	DTCreate               string  `json:"dt_create"`
	DTModify               string  `json:"dt_modify"`
}
