package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"

	_ "github.com/mattn/go-sqlite3"
	"paybook.com/lite/models"
	_ "paybook.com/lite/routers"
)

var globalSessions *session.Manager

func init() {

}

func main() {
	u := models.User{}
	u.Init()

	beego.Run()
}
