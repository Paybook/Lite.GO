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
