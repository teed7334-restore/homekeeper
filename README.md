# homekeeper
服務管理員，日後會將工作上常用到的服務放在裡面，好比變動資料庫、區塊鏈上鏈、爬蟲抓資料等放進來，讓裡面可以有許多現成的微服務，帶了就走

## 資料夾結構
beans 用來裝Call API後之ResultObject

controllers Restful API呼叫用控制器

.env 系統設定

route 系統路由設定

main.go 主程式

## 程式運行原理
本系統會啟動一個Restful API

目前它己有一個現成的寄信服務，其他可以依自己做擴充

本系統可以搭配佇列管理員管理員使用，這樣一來你就能安排一些非同步任務行程

https://github.com/teed7334-restore/counter

## 程式操作流程
1. 將.env.swp檔名改成.env
2. 修改.env並設定您的SMTP Server
3. go run main.go

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

取得Redis資料 http://[Your Host Name]:8806/Redis/Get
```
//HTTP Header需設定成Content-Type: application/json
{
    "key": "score"
}
```

設定Redis資料 http://[Your Host Name]:8806/Redis/Set
```
//HTTP Header需設定成Content-Type: application/json
{
    "key": "score",
    "value": "0"
}
```

對Redis資料進行遞增 http://[Your Host Name]:8806/Redis/Incr
```
//HTTP Header需設定成Content-Type: application/json
{
    "key": "score"
}
```

對Redis資料進行遞減 http://[Your Host Name]:8806/Redis/Decr
```
//HTTP Header需設定成Content-Type: application/json
{
    "key": "score"
}
```

對Redis資料建立Hashmap http://[Your Host Name]:8806/Redis/HSet
```
//HTTP Header需設定成Content-Type: application/json
{
    "key": "profile",
    "hkey": "name",
    "value": "Peter"
}
```

取得建立Hashmap的Redis資料 http://[Your Host Name]:8806/Redis/HGet
```
//HTTP Header需設定成Content-Type: application/json
{
    "key": "profile",
    "hkey": "name"
}
```

對Redis進行資料添加，並排除重複項目 http://[Your Host Name]:8806/Redis/SAdd
```
//HTTP Header需設定成Content-Type: application/json
{
    "key": "box",
    "value": "Joe"
}
```

取得Key中的資料集合總數 http://[Your Host Name]:8806/Redis/SCard
```
//HTTP Header需設定成Content-Type: application/json
{
    "key": "box"
}
```

對Redis中List資料的後面做資料添加 http://[Your Host Name]:8806/Redis/RPush
```
//HTTP Header需設定成Content-Type: application/json
{
    "key": "menu",
    "value": "apple"
}
```

對Redis中List資料的前面做資料添加 http://[Your Host Name]:8806/Redis/LPush
```
//HTTP Header需設定成Content-Type: application/json
{
    "key": "menu",
    "value": "banana"
}
```

從Redis中List資料的第...列做資料修改 http://[Your Host Name]:8806/Redis/LSet
```
//HTTP Header需設定成Content-Type: application/json
{
    "key": "menu",
    "hkey": "1",
    "value": "cherry"
}
```

取得Redis中List特定範圍裡面的資料 http://[Your Host Name]:8806/Redis/LRange
```
//HTTP Header需設定成Content-Type: application/json
{
    "key": "menu",
    "value": "0:5"
}
```

計算休假區間，中間會扣除午休九十分鐘 http://[Your Host Name]:8806/PunchClock/CalcTime
```
//HTTP Header需設定成Content-Type: application/json
{
  "begin":{"year":"2019","month":"6","day":"19","hour":"10","minute":"46"},
  "end":{"year":"2019","month":"6","day":"25","hour":"18","minute":"0"}
}
```

重新設定所有員工休假區間 http://[Your Host Name]:8806/PunchClock/ResetAllUseMinute
```
//HTTP Header需設定成Content-Type: application/json
{}
```