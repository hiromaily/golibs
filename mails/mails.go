package mails

import (
	"fmt"
	"log"
	"net/smtp"
	"time"
	//lg "github.com/hiromaily/golibs/log"
	conf "github.com/hiromaily/goweb/configs"
	"encoding/base64"
)

//https://gist.github.com/andelf/5004821

type MailInfo struct {
	ToAddress []string
	Subject   string
	Body      string
}

func (self *MailInfo) makeMailBody(){
	//TODO:when toaddress have more one address
	header := make(map[string]string)
	header["From"] = conf.GetConfInstance().Mail.Address
	header["To"] = self.ToAddress[0]  //TODO:暫定
	header["Subject"] = self.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	var message string = ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(self.Body))

	//update self.Body
	//TODO:????
	//*self.Body = message
	self.Body = message
}

func (self *MailInfo) sendMail(c chan bool) {
	co := conf.GetConfInstance().Mail

	// Set up authentication information.
	// func PlainAuth(identity, username, password, host string) Auth {
	auth := smtp.PlainAuth(
		"",
		co.Address,      // info@gmail.com
		co.Password,     // password
		co.Smtp.Server, // smtp.xxxx.com
	)

	//body
	self.makeMailBody()

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error {
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", co.Smtp.Server, co.Smtp.Port), // smtp.xxxxx.com:255
		auth,
		co.Address,         //(from)info@gmail.com
		self.ToAddress,     //(to)address
		[]byte(self.Body),  //body
	)
	if err != nil {
		log.Fatal(err)
		c <- false
		return
	}
	c <- true
	return
}

//mails.SendMail([]string{"hiromaily@gmail.com"}, "subject", "body")
func SendMail(toAddr []string, subject string, body string) {
	//MailInfo
	mi := &MailInfo{ToAddress:toAddr, Subject:subject, Body:body}

	emailTimeout, _ := time.ParseDuration(conf.GetConfInstance().Mail.Timeout) //10s
	c := make(chan bool)

	//send mail
	//for handling timeout using goroutine.
	go mi.sendMail(c)

	fmt.Println("start to wait")
	select {
	case bRet := <-c:
	//it may be ok
		if bRet {
			fmt.Println("succeeded")
		}else{
			fmt.Println("failed")
		}
	case <-time.After(emailTimeout):
	//timeout code
		fmt.Println("timeout")
	}

}
