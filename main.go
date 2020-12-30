package main

import (
	"Item/dbmysql"
	_ "Item/routers"
	"github.com/astaxie/beego"
)
//code=9ced7f4713ff616ac333
func main() {
	/*	blockount,err :=btcService.GetBlockCount()
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
			fmt.Println("获取新地址:",address)*/


	dbmysql.Connect()
	//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")

/*	http.HandleFunc("/", utils.Hello)
	http.HandleFunc("/oauth/redirect", utils.Oauth) // 这个和 Authorization callback URL 有关*/
	/*	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println("监听失败，错误信息为:", err)
		//return
	*/
	beego.Run()
}
