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
	users := models.User{}
	users.Init()

	beego.Run()
}
