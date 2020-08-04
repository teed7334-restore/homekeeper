package beans

//Redis Query用參數
type Redis struct {
	Key   string //`json:"key,omitempty"`
	Value string //`json:"value,omitempty"`
	Hkey  string //`json:"hkey,omitempty"`
}
