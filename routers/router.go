package routers

import (
	"Item/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //登录接口
    beego.Router("/home",&controllers.LoginController{})

    //注册接口
    beego.Router("/register",&controllers.RegisterController{})


   beego.Router("/register_sms",&controllers.RegisterSmsController{})


}
