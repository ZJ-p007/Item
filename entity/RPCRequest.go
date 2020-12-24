package entity

/**
	"id" :编号
	"method":功能方法和命令
	"jsonrpc":rpc版本
	"params":[携带的参数，数组格式]
 */
type RPCRequest struct {
	Id int64 `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method string `json:"method"`
	Params []interface{} `json:"params"`
}
