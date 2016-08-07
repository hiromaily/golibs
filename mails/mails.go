package mails

import (
	"fmt"
	"log"
	"net/smtp"
	"time"
	//lg "github.com/hiromaily/golibs/log"
	"encoding/base64"
	//conf "github.com/hiromaily/golibs/config"
)

//https://gist.github.com/andelf/5004821

type MailInfo struct {
	ToAddress   []string
	FromAddress string
	Subject     string
	Body        string
	Smtp
}

type Smtp struct {
	Address string
	Pass    string
	Server  string
	Port    int
}

func (ml *MailInfo) makeMailBody() {
	//TODO:when toaddress have more one address
	header := make(map[string]string)
	header["From"] = ml.FromAddress
	header["To"] = ml.ToAddress[0] //TODO:temporary
	header["Subject"] = ml.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	var message string = ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(ml.Body))

	//update self.Body
	//*self.Body = message
	ml.Body = message
}

func (ml *MailInfo) sendMail(c chan bool) {
	//co := conf.GetConfInstance().Mail

	// Set up authentication information.
	// func PlainAuth(identity, username, password, host string) Auth {
	auth := smtp.PlainAuth(
		"",
		ml.Address, // info@gmail.com
		ml.Pass,    // password
		ml.Server,  // smtp.xxxx.com
	)

	//body
	ml.makeMailBody()

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error {
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", ml.Server, ml.Port), // smtp.xxxxx.com:255
		auth,
		ml.Address,      //(from)info@gmail.com
		ml.ToAddress,    //(to)address
		[]byte(ml.Body), //body
	)
	if err != nil {
		log.Fatal(err)
		c <- false
		return
	}
	c <- true
	return
}

func (ml *MailInfo) SendMail(timeOut string) {

	emailTimeout, _ := time.ParseDuration(timeOut) //10s
	c := make(chan bool)

	//send mail
	//for handling timeout using goroutine.
	go ml.sendMail(c)

	fmt.Println("start to wait")
	select {
	case bRet := <-c:
		//it may be ok
		if bRet {
			fmt.Println("succeeded")
		} else {
			fmt.Println("failed")
		}
	case <-time.After(emailTimeout):
		//timeout code
		fmt.Println("timeout")
	}

}
