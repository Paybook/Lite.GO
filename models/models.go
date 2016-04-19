package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// User of models
type User struct {
	ID       int    `form:"-"`
	Email    string `form:"email",orm:"size(100);unique"`
	Password string `form:"password",orm:"size(255)"`
}

// TableName creation
func (u *User) TableName() string {
	return "users"
}

// Init models
func (u *User) Init() {
	fmt.Println("Init models")

	// orm.Debug = true
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	// set default database
	orm.RegisterDataBase("default", "sqlite3", "data.db")

	orm.RegisterModel(new(User))

	orm.RunCommand()

	name := "default"
	force := true
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
		fmt.Printf("ID: %d, Email: %s, Password: %s, ERR: %v\n", id, u.Email, u.Password, err)
	} else {
		fmt.Printf("User alredy exists Email: %s\n", u.Email)
	}
}

func (u *User) find() bool {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable("users").Filter("email", u.Email).One(&user)
	if err == nil {
		return false
	}
	fmt.Println("User name", user.ID, user.Email)
	fmt.Printf("ERR: %v\n", err)

	return true
}

// Auth users
func (u *User) Auth() bool {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable("users").Filter("email", u.Email).One(&user)
	if err == nil {
		fmt.Println("User name", user.ID, user.Email)

		if comparePass(user.Password, u.Password) {
			return true
		}

		beego.Debug("Wrong password ", u.Password)
		return false
	}

	beego.Debug("No user found: ", err)

	return false
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
