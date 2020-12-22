package utils

import (
	"BcRPC/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Post(method string) (*entity.RpcResult,error) {
	/**
	1、准备要进行rpc通信的json数据
	2、使用http连接的post请求，发送json数据
	3、接受http响应
	4、根据响应的结果，进行判断和处理
		状态码200正常
		非200不正常
	*/

	//1、准备要发送的json数据
	/**
	{
		"id":编号
		"method":"更能方法或者命令"
		"jsonrpc":"rpc版本2.0"
		"params":[携带的参数，数组形式]
	}
	*/
	//json操作：序列化、反序列化
	rpcReq := entity.RPCRequest{
		Id:      time.Now().Unix(),
		Jsonrpc: "2.0",
		Method:  method,

	}
	//对结构体类型进行序列化
	rpcBytes,err := json.Marshal(&rpcReq)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	//fmt.Println(string(rpcBytes))

	//2、发送post请求
	//client：客户端；客户端用于发起请求
	client := http.Client{}//实例化一个客户端

	//实例化一个请求
	request,err := http.NewRequest("POST",RPCURL,bytes.NewBuffer(rpcBytes))
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	//给post请求添加请求头设置信息
	request.Header.Add("Encoding","UTF-8")
	request.Header.Add("Content-Type","application/json")
	request.Header.Add("Authorization","Basic " + Base64Str(RPCUSER+":"+RPCPASSWORD))

	//使用客户端发送请求
	response,err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}

	//通过response获取响应的数据
	code := response.StatusCode
	resultBytes,err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	//fmt.Println(resultBytes)

	if code == http.StatusOK {

		//fmt.Println("这",string(resultBytes))
		var result *entity.RpcResult
		err = json.Unmarshal(resultBytes,&result)
		if err != nil {
			fmt.Println(err.Error())
			return nil,err
		}

		//反序列化正常，没有出现错误
		//fmt.Println("功能调用结果：",result.Result)
		return result,err
	}else {
		fmt.Println("请求失败，错误状态码是：",code)
		return nil,err
	}

}
