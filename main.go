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

	// globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	// go globalSessions.GC()
}

func main() {
	u := models.User{}
	u.Init()

	beego.Run()
}
