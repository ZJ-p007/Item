package main

import (
	"Item/btcService"
	"Item/dbmysql"
	_ "Item/routers"
	"fmt"
	"github.com/astaxie/beego"
)
//https://github.com/GongJiangHua/BcConnectWeb.git

func main() {
	blockount,err :=btcService.GetBlockCount()
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("区块总数:",blockount)

	//最新区块的hash
	hash,err:=btcService.GetBestBlockHash()
    if err!= nil{
    	fmt.Println(err.Error())
		return
	}
	fmt.Println("最新区块的hash:",hash)

	//获取区块的难度
	difficulty,err :=btcService.GetDifficulty()
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("获取区块的难度:",difficulty)

	address,err :=btcService.GetNewAddress()
	if err!= nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("获取新地址:",address)

	
	dbmysql.Connect()
	//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
