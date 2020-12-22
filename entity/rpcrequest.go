package entity

/**
  {
	"id":编号
	"method":"更能方法或者命令"
	"jsonrpc":"rpc版本2.0"
	"params":[携带的参数，数组形式]
  }
*/
type RPCRequest struct {
	Id      int64         `json:"id"`
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}
