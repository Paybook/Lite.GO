package controllers

import "github.com/astaxie/beego"

// SessionController ...
type SessionController struct {
	beego.Controller
}

// Session ...
type Session struct {
	Status bool `json:"status"`
}

// Get ...
func (c *SessionController) Get() {
	var session Session
	session.Status = true

	pbSession := c.GetSession("paybook-lite")
	if pbSession == nil {
		session.Status = false
	}

	c.Data["json"] = &session
	c.ServeJSON()
}
