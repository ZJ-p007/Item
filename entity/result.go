package entity

type Result struct {
	Code int       `json:"code"`//返回的状态码
	Msg  string    `json:"msg"`//状态码对应的描述信息
	Data RpcResult `json:"data"`
}
