package utils

import (
	"BcRPC/entity"
	"encoding/json"
	"fmt"
)

/**
获取钱包中的所有余额
返回Mine结构体
*/
func Balances() (*entity.Mine,error) {

	result,err := Post(GETBALANCES)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}

	balance,err := json.Marshal(result)
	if err != nil {
		return nil,err
	}

	var balances *entity.Mine
	err = json.Unmarshal(balance,&balances)
	if err != nil {
		return nil,err
	}
	return balances,err

}

/**
获取服务器运行的时间
返回运行的秒数
*/
func UpTime() interface{} {
	result,_ := Post(UPTIME)
	//单位：秒
	return result.Result
}

/**
获取所有服务端命令
返回字符串
*/
func Help() interface{} {
	result,_ := Post(HELP)
	return result.Result

}

/**
获取最新区块的hash值
返回字符串类型
*/
func BestBlockHash(data string) interface{} {
	result,_ := Post(data)
	return result.Result
}

/**
获取比特币节点的区块数
返回int类型数字
*/
func BlockCount() interface{} {
	result,_ := Post(GETBLOCKCOUNT)
	return result.Result
}

/**
获取当前区块的信息
返回ChainInfo结构体
*/
func ChainInfo(data string) (*entity.ChainInfo,error) {
	result,err := Post(data)

	chainByte,err := json.Marshal(result.Result)
	if err != nil {
		return nil,err
	}
	//fmt.Println(chainByte)
	var chaininfo *entity.ChainInfo
	err = json.Unmarshal(chainByte,&chaininfo)
	if err != nil {
		return nil,err
	}

	return chaininfo,nil

}
