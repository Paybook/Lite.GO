package models

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"paybook.com/lite/services"
)

// User of models
type User struct {
	ID       int    `form:"-"`
	IDUser   string `form:"-"`
	Email    string `form:"email",orm:"size(100);unique"`
	Password string `form:"password",orm:"size(255)"`
}

// TableName creation
func (u *User) TableName() string {
	return "users"
}

// Init models
func (u *User) Init() {
	beego.Info("Init models")

	// orm.Debug = true
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	// set default database
	orm.RegisterDataBase("default", "sqlite3", "data.db")

	orm.RegisterModel(new(User), new(Account), new(Transaction))

	orm.RunCommand()

	name := "default"
	force := false
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

// Create user
func (u *User) Create() {

	o := orm.NewOrm()

	// insert
	u.Password = getPass(u.Password)

	if u.find() {
		id, err := o.Insert(u)

		url := beego.AppConfig.String("pbsync_base_url") +
			"users?api_key=" + beego.AppConfig.String("pbsync_api_key") +
			"&id_external=" + string(id) +
			"&name=" + u.Email

		services := services.Services{}
		res := services.Post(url)

		api := API{}
		err = json.Unmarshal([]byte(res), &api)
		if err != nil {
			beego.Error("Error parsing API post output: %s - %s", err, res)
		}

		if api.Code == 200 {
			beego.Info(api.Response)
			userAPI := UserAPI{}
			err = json.Unmarshal([]byte(api.Response), &userAPI)
			if err != nil {
				beego.Error("Error parsing userAPI response: %s - %s", err, res)
			}

			u.IDUser = userAPI.IDUser
			num, err := o.Update(u)
			if err != nil {
				beego.Error("Error updating user.id_user ", err)
			}
			beego.Info("IDUser updated: ", num)
		} else {
			beego.Error("Error: on create user on pbsync")
		}

		fmt.Printf("ID: %d, Email: %s, Password: %s, ERR: %v\n", id, u.Email, u.Password, err)
	} else {
		fmt.Printf("User alredy exists Email: %s\n", u.Email)
	}
}

func (u *User) getUsers(token string) {

	url := beego.AppConfig.String("pbsync_base_url") + "/users?token=" + token

	// api := API{}

	services := services.Services{}
	res := services.Get(url)
	beego.Debug(res)
}

func (u *User) find() bool {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable("users").Filter("email", u.Email).One(&user)
	if err == nil {
		beego.Error("ERR: ", err)
		return false
	}
	beego.Info("User info: ", user.ID, user.Email)

	return true
}

// Auth users
func (u *User) Auth() (bool, string) {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable("users").Filter("email", u.Email).One(&user)
	if err == nil {
		beego.Info("User info: ", user.ID, user.Email)

		if comparePass(user.Password, u.Password) {
			return true, user.IDUser
		}

		beego.Debug("Wrong password ", u.Password)
		return false, ""
	}

	beego.Debug("No user found: ", err)

	return false, ""
}

func getPass(password string) string {
	passwordByte := []byte(password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)

}

func comparePass(hashedPassword, password string) bool {
	passwordByte := []byte(password)
	hashedPasswordByte := []byte(hashedPassword)

	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword(hashedPasswordByte, passwordByte)
	if err == nil {
		return true
	}
	return false
}
