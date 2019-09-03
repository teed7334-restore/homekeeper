package controllers

import (
	"testing"

	"github.com/teed7334-restore/homekeeper/beans"
)

func Test_doSendMail(t *testing.T) {
	type args struct {
		params *beans.SendMail
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "一般測試",
			args: args{
				params: &beans.SendMail{
					To:      "teed7334@gmail.com",
					Cc:      "teed7334@163.com",
					Subject: "這是一封測試信",
					Content: "這是一封測試信",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doSendMail(tt.args.params)
		})
	}
}
