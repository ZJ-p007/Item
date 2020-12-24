package btcService

import (
	"Item/entity"
	"Item/utils"
	"encoding/json"
	"fmt"
)

//该包用来封装bitcoin命令

//用于查询区块链中的区块数量
func GetBlockCount() (interface{}, error) {
	jsonMes, err := utils.PrepareJsonStr("getblockcount", nil)
	if err != nil {
		fmt.Println("准备json传输数据失败，错误原因：", err.Error())
		return nil, err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		fmt.Println("请求RPC服务失败，错误原因：", err.Error())
		return nil, err
	}
	Rpcresult := entity.RPCResult{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	return Rpcresult.Result, nil
}

//用于查询区块链中最新区块的hash

func GetBestBlockHash() (interface{}, error) {
	jsonMes, err := utils.PrepareJsonStr("getbestblockhash", nil)
	if err != nil {
		return  nil,err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil, err
	}
	Rpcresult := entity.RPCResult{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	return Rpcresult.Result, nil
}

//根据区块高度查询区块hash
//func GetBlockHashByHeight() (interface{}, error) {
//	jsonMes, err := utils.PrepareJsonStr("getblockhash",nil)
//	if err != nil {
//		return nil, err
//	}
//	resBytes, err := utils.SendRPCPost(jsonMes)
//	if err != nil {
//		return nil, err
//	}
//	Rpcresult := entity.RPCResult{}
//	err = json.Unmarshal(resBytes,&Rpcresult)
//	if err != nil {
//		return nil,err
//	}
//	return Rpcresult.Result, nil
//}

//用于获取当前区块的难度
func GetDifficulty() (interface{}, error) {
	jsonMes, err := utils.PrepareJsonStr("getdifficulty", nil)
	if err != nil {
		return nil, err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil, err
	}
	Rpcresult := entity.RPCResult{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	return Rpcresult.Result, nil
}

//用于获取当前区块链信息
func GetBlockChainInfo() (*entity.BlockChainInfo,error) {
	jsonMes, err := utils.PrepareJsonStr("getblockchaininfo",nil)
	if err != nil {
		return nil,err
	}
	resBytes,err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil,err
	}
	Rpcresult := entity.BlockChainInfo{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	return &Rpcresult,nil
}

//生成一个新地址
func GetNewAddress()(interface{},error)  {
	jsonMes, err := utils.PrepareJsonStr("getnewaddress", nil)
	if err != nil {
		return nil, err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil, err
	}
	Rpcresult := entity.RPCResult{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	return Rpcresult.Result, nil

}

//
/*func GetBlockHeader(hash interface{})(interface{},error)  {
	jsonMes, err := utils.PrepareJsonStr("getblockheader", nil)
	if err != nil {
		return nil, err
	}
	resBytes, err := utils.SendRPCPost(jsonMes)
	if err != nil {
		return nil, err
	}
	Rpcresult := entity.RPCResult{}
	err = json.Unmarshal(resBytes,&Rpcresult)
	if err != nil {
		return nil,err
	}
	return Rpcresult.Result, nil*/

