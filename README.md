# homekeeper
服務管理員，日後會將工作上常用到的服務放在裡面，好比變動資料庫、區塊鏈上鏈、爬蟲抓資料等放進來，讓裡面可以有許多現成的微服務，帶了就走

## 資料夾結構
beans 用來裝Call API後之ResultObject

controllers Restful API呼叫用控制器

env 系統設定

route 系統路由設定

main.go 主程式

## 程式運行原理
本系統會啟動一個Restful API

目前它己有一個現成的寄信服務，其他可以依自己做擴充

本系統可以搭配佇列管理員管理員使用，這樣一來你就能安排一些非同步任務行程

https://github.com/teed7334-restore/counter

## 必須套件
本程式透過Google Protobuf 3產生所需之ResultObject，然Proto 3之後官方不支持Custom Tags，所以還需要多安裝一個寫入retags的套件

git clone https://github.com/qianlnk/protobuf.git $GOPATH/src/github.com/golang/protobuf

go install $GOPATH/src/github.com/golang/protobuf/protoc-gen-go

及Restful Framework

go get -u -v github.com/gin-gonic/gin

## 程式操作流程
1. 將./env/env.swp檔名改成env.go
2. 修改./env/env.go並設定您的SMTP Server
3. 到./beans底下，運行protoc --go_out=plugins=grpc+retag:. *.proto
4. go run main.go

## API呼叫網址與參數
寄信服務 http://[Your Host Name]:8806/Mail/SendMail
```
//HTTP Header需設定成Content-Type: application/json
{
    "to": "admin@admin.com",
    "subject": "這是一封測試信",
    "content": "這是一封測試信<br />這是一封測試信<br />這是一封測試信<br />這是一封測試信<br />這是一封測試信<br />"
}
```