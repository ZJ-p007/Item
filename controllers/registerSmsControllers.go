package controllers

import (
	"Item/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterSmsController struct {
	beego.Controller
}

/*func (c *RegisterSmsController) Get() {
	c.Ctx.WriteString("aaaaaa")
}*/

//处理注册请求
func (c * RegisterSmsController) Post() {
	//解析
	var user models.User
	err :=c.ParseForm(&user)
	if err!= nil{
		c.Ctx.WriteString("数据解析失败")
	}
	fmt.Println("电话",user.Phone)

	//将解析的数据保存带数据库
	_,err =user.AddUser()
	if err != nil{
		fmt.Println(err.Error())
		c.Ctx.WriteString("用户注册失败")
		return
	}
	//fmt.Println(res)

	c.TplName = "login.html"
}
