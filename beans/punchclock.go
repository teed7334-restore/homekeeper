package beans

//TimeStruct 時間參數
type TimeStruct struct {
	Year   string //`json:"year"`
	Month  string //`json:"month"`
	Day    string //`json:"day"`
	Hour   string //`json:"hour"`
	Minute string //`json:"minute"`
	Second string //`json:"second"`
}

//EmployeeOnChain 上鏈用員工資料
type EmployeeOnChain struct {
	Class     string //`json:"$class" form:"$class"`
	Identify  string //`json:"identify" form:"identify"`
	FirstName string //`json:"firstName" form:"firstName"`
	LastName  string //`json:"lastName" form:"lastName"`
}

//PunchclockOnChain 上鏈用打卡記錄
type PunchclockOnChain struct {
	Class       string //`json:"$class" form:"$class"`
	ID          string //`json:"id" form:"id"`
	OnWorkDate  string //`json:"onWorkDate" form:"onWorkDate"`
	OnWorkTime  string //`json:"onWorkTime" form:"onWorkTime"`
	OffWorkTime string //`json:"offWorkTime" form:"offWorkTime"`
	WorkTimes   string //`json:"workTimes" form:"workTimes"`
	Employee    string //`json:"employee" form:"employee"`
}

//Punchclock 上下班時間參數
type Punchclock struct {
	Begin TimeStruct //`json:"begin"`
	End   TimeStruct //`json:"end"`
}

//DailyPunchclockData 員工打卡記錄
type DailyPunchclockData struct {
	Employee   EmployeeOnChain //`json:"employee" form:"employee"`
	Punchclock Punchclock      //`json:"punchclock" form:"punchclock"`
}
