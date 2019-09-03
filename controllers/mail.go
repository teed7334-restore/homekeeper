package controllers

import (
	"log"
	"net/http"

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
	mail.SetHeader("From", cfg.Mail.From)
	mail.SetHeader("To", params.GetTo())
	if "" != params.GetCc() {
		mail.SetHeader("Cc", params.GetCc())
	}
	mail.SetHeader("Subject", params.GetSubject())
	mail.SetBody("text/html", params.GetContent())
	send := gomail.NewPlainDialer(cfg.Mail.Host, cfg.Mail.Port, cfg.Mail.User, cfg.Mail.Password)
	if err := send.DialAndSend(mail); err != nil {
		log.Panicln(err)
	}
}
