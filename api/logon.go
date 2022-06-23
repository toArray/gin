package api

type AAA struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ReturnData struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data []interface{} `json:"data"`
}
