package beans

//SendMail 寄信用參數
type SendMail struct {
	To      string //`form:"to"`
	Cc      string //`form:"cc"`
	Subject string //`form:"subject"`
	Content string //`form:"content"`
}
