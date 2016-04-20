package services

import (
	"crypto/tls"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

// Services ...
type Services struct {
}

// Post ...
func (s *Services) Post(url string) string {
	beego.Info(url)

	req := httplib.Post(url)
	req.Debug(true)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	str, err := req.String()
	if err != nil {
		beego.Warning(err)
	}

	return str
}

// Get ...
func (s *Services) Get(url string) string {
	beego.Info(url)

	req := httplib.Get(url)
	req.Debug(true)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	str, err := req.String()
	if err != nil {
		beego.Warning(err)
	}

	return str
}
