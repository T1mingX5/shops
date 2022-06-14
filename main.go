package main

import (
	"shops/controllers"
	_ "shops/initial" // _ 后面调用包，指的是自动执行这些包里的init方法
	_ "shops/routers"

	"github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

func main() {
	web.InsertFilter("*", web.BeforeRouter, FilterUser)
	web.Run()
}

// 过滤器，验证用户是否登录
var FilterUser = func(ctx *beecontext.Context) {
	// token := ctx.Request.Header.Get("token")
	// _, _, err := controllers.ParseToken(token)
	// url := ctx.Request.RequestURI
	// if err != nil && url != "/login" && url != "/register" {
	// 	ctx.Redirect(302, "/login")
	// }
	info, ok := ctx.Input.Session("info").(controllers.Info)

	url := ctx.Request.RequestURI
	if !ok && url != "/login" && url != "/register" && url != "/adminlogin" {
		ctx.Redirect(302, "/login")
	}
	if ok && (url == "/login" || url == "/adminlogin") {
		if info.Role == "admin" {
			ctx.Redirect(302, "/controller")
		}
		if info.Role == "user" {
			ctx.Redirect(302, "/")
		}
	}
}
