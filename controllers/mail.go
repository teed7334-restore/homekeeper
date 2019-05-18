package controllers

import (
	"log"
	"net/http"

	"../beans"
	"../env"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

//SendMail 寄信用API
func SendMail(c *gin.Context) {
	params := &beans.SendMail{}
	err := c.BindJSON(params)
	if err != nil {
		log.Println(err)
	}
	doSendMail(params)
	c.JSON(http.StatusOK, gin.H{"status": "true"})
}

//doSendMail 寄發通知郵件
func doSendMail(params *beans.SendMail) {
	cfg := env.GetEnv()
	mail := gomail.NewMessage()
	mail.SetHeader("From", cfg.Mail.From)
	mail.SetHeader("To", params.GetTo())
	mail.SetHeader("Subject", params.GetSubject())
	mail.SetBody("text/html", params.GetContent())
	send := gomail.NewPlainDialer(cfg.Mail.Host, cfg.Mail.Port, cfg.Mail.User, cfg.Mail.Password)
	if err := send.DialAndSend(mail); err != nil {
		log.Panic(err)
	}
}
