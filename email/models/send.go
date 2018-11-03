package models

import (
	"log"
	"net/smtp"
	"strings"
)

//func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
//func PlainAuth(identity, username, password, host string) Auth

type SmtpClient struct {
	Addr string `json:"addr" binding:"required"` // 账号/邮箱地址
	Pass string `json:"pass"  binding:"required"`
	Host string `json:"host" binding:"required"` // smtp服务器地址
	Auth smtp.Auth
}
type Message struct {
	Subject  string `json:"subject" binding:"required"`
	Body     string `json:"body" binding:"required"`
	To       string `json:"to" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}



func GenerateAuth(addr, pass, host string) *SmtpClient {
	client := &SmtpClient{
		Addr: addr,
		Pass: pass,
		Host: host,
	}
	client.Auth = smtp.PlainAuth("", client.Addr, client.Pass, client.Host)
	return client
}

func (this *SmtpClient) AsyncSend(msg Message, isMass bool, handle func(err error)) error {
	go func() {
		err := this.Send(msg, isMass)
		handle(err)
	}()
	return nil
}
func (this *SmtpClient) Send(msg Message, isMass bool) error {
	switch isMass {
	case false:
		this.OneSend(msg)
	case true:
		this.MassSend(msg)
	}
	return nil
}
func (this *SmtpClient) OneSend(msg Message) {
	this.SendMail(msg)

}
func (this *SmtpClient) MassSend(msg Message) {

}
func (this *SmtpClient) SendMail(mess Message) bool {

	content_type := "Content-Type: text/plain; charset=UTF-8"

	to := []string{mess.To}
	nickname := mess.Nickname
	subject := mess.Subject
	body := mess.Body

	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user.Addr + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	print("msg", mess.To)

	err := smtp.SendMail("smtp.163.com:25", this.Auth, this.Addr, to, msg)
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		return true
	}

}
