package utils

import (
	"Item/Constants"
	"Item/entity"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)
/**
1、准备一个json数据
2、使用http链接的post请求，发送json数据
3、接受http响应
4、为post请求添加请求头
5、处理返回的结果
 */


//1、准备一个json数据用于向
func PrepareJsonStr(method string,params []interface{}) ([]byte,error) {
	Mes := entity.RPCRequest{}
	Mes.Method = method
	Mes.Jsonrpc = "2.0"
	Mes.Id = time.Now().Unix()
	Mes.Params = params
	jsonBytes,err := json.Marshal(Mes)
	if err != nil {
		return nil,err
	}
	return jsonBytes,nil
}


func SendRPCPost(jsonBytes []byte) ([]byte,error) {
	//准备一个客户端
	client := http.Client{}
	//实例化一个请求
	req ,err := http.NewRequest("POST",Constants.RPCURL,bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil,err
	}
	//为请求头添加数据
	req.Header.Add("Encoding","UTF-8")
	req.Header.Add("Content-Type","application/json")
	req.Header.Add("Authorization","Basic "+Base64Str(Constants.RPCUSER+":"+Constants.RPCPASSWORD))
	//使用客户端发起请求
	resp,err := client.Do(req)
	if err != nil {
		return nil,err
	}
	//处理响应结果
	code := resp.StatusCode
	if code == 200 {
		resBytes,err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil,err
		}

		return resBytes,nil
	}else {
		return nil,err
	}
}
