package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"strings"
)

var FromEmail string
var PasswordEmail string
var HostEmail string

func SendMail(to, message, externalUrlToPic string) {

	m := gomail.NewMessage()
	m.SetHeader("From", FromEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Служба поддержки")
	body := "<a href=" + message + ">Ссылка для сброса пароля</a>" + "<img src=" + externalUrlToPic + "\">"
	m.SetBody("text/html", body)

	d := gomail.NewDialer(HostEmail, 465, FromEmail, PasswordEmail)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

//check valid mail
func CheckMail(email string) bool {
	contain := strings.Contains(email, "@")
	if contain {
		return true
	} else {
		fmt.Printf("'%s' is not a valid email\n", email)
		return false
	}
}