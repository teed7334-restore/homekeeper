package beans

//GetEmployee 取得員工資料
type GetEmployee struct {
	Class     string   //`json:"$class" form:"$class"`
	Identify  string   //`json:"identify" form:"identify"`
	FirstName string   //`json:"firstName" form:"firstName"`
	LastName  string   //`json:"lastName" form:"lastName"`
	Error     APIError //`json:"error" form:"error"`
}

//GetPunchclock 取得打卡記錄
type GetPunchclock struct {
	Class       string   //`json:"$class" form:"$class"`
	ID          string   //`json:"id" form:"id"`
	OnWorkDate  string   //`json:"onWorkDate" form:"onWorkDate"`
	OffWorkTime string   //`json:"offWorkTime" form:"offWorkTime"`
	WorkTimes   string   //`json:"workTimes" form:"workTimes"`
	Employee    string   //`json:"employee" form:"employee"`
	Error       APIError //`json:"error" form:"error"`
}

//APIError API錯誤訊息
type APIError struct {
	StatusCode string //`json:"status_code" form:"status_code"`
	Name       string //`json:"name" form:"name"`
	Message    string //`json:"message" form:"message"`
	Status     int    //`json:"status" form:"status"`
	Code       string //`json:"code" form:"code"`
	Stack      string //`json:"stack" form:"stack"`
}
