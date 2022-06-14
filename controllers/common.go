package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type CommonController struct {
	web.Controller
}

func (c *CommonController) Rsp(status bool, message interface{}) {
	statusCode := 200
	if !status {
		statusCode = 400
	}
	c.Data["json"] = map[string]interface{}{"status": statusCode, "message": message}
	c.ServeJSON()
}

func (c *CommonController) IsPost() bool {
	return c.Ctx.Request.Method == "POST"
}
