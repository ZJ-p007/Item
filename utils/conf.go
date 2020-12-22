package utils

const (
	//rpc服务协议
	RPCURL = "http://127.0.0.1:7001"//rpc的链接
	RPCUSER = "user"//rpc用户名
	RPCPASSWORD = "pwd"//rpc用户密码
	//获取当前区块的信息
	GETBLOCKCHAININFO = "getblockchaininfo"
	//查询钱包的所有余额
	GETBALANCES = "getbalances"
	//查询服务端的所有命令
	HELP = "help"
	//获取比特币节点的区块总数
	GETBLOCKCOUNT = "getblockcount"
	//获取服务器的运行时间
	UPTIME = "uptime"
	//获取最新区块的hash值
	GETBESTBLOCKHASH = "getbestblockhash"

)
