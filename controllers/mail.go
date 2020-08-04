package controllers

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/teed7334-restore/homekeeper/beans"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

//SendMail 寄信用API
func SendMail(c *gin.Context) {
	params := &beans.SendMail{}
	getParams(c, params)
	doSendMail(params)
	c.JSON(http.StatusOK, gin.H{"status": "true"})
}

//doSendMail 寄發通知郵件
func doSendMail(params *beans.SendMail) {
	mail := gomail.NewMessage()
	from := os.Getenv("mail.from")
	mail.SetHeader("From", from)
	mail.SetHeader("To", params.To)
	if "" != params.Cc {
		mail.SetHeader("Cc", params.Cc)
	}
	mail.SetHeader("Subject", params.Subject)
	mail.SetBody("text/html", params.Content)

	host := os.Getenv("mail.host")
	port, _ := strconv.Atoi(os.Getenv("mail.port"))
	user := os.Getenv("mail.user")
	passwd := os.Getenv("mail.password")
	send := gomail.NewPlainDialer(host, port, user, passwd)

	//取消TLS連線
	send.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := send.DialAndSend(mail); err != nil {
		log.Println(err)
	}
}
