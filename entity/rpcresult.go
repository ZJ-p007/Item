package entity

type RpcResult struct {
	Id     int `json:"id"`
	Error  string `json:"error"`
	//go语言中万能类型：interface
	Result interface{} `json:"result"`
}
