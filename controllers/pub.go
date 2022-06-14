package controllers

import (
	"shops/models"
	"time"
)

type Pub struct {
	CommonController
}

func (p *Pub) Login() {
	if p.IsPost() { // 后端是不能主动去请求前端的，这句话的意思是前端“请求(request)”以POST方式给后端传送数据
		account := p.GetString("account") // 后端接受所需数据
		password := p.GetString("password")
		if account == "" || password == "" {
			p.Rsp(false, "输入内容不能为空")
			return
		}
		id, status, _ := models.UserLogin(account, password)
		if id == 0 {
			p.Rsp(false, "用户名或密码不正确")
			return
		}
		if status == 2 {
			p.Rsp(false, "用户账户被禁用")
			return
		} else if status > 2 || status < 1 {
			p.Rsp(false, "账号异常")
			return
		}
		// token, err := generateToken(id, account)
		// if err != nil {
		// 	u.Rsp(false, err.Error())
		// 	return
		// }
		p.SetSession("info", Info{
			Role: "user",
			Id:   id,
		}) // user-1 user-2
		p.Rsp(true, "登录成功，等待跳转")
	} else {
		p.TplName = "app/pub/login.html" // 如果前端传回来的不是post请求，那么就使用get方法，访问这个页面
	}
}

func (p *Pub) Register() {
	if p.IsPost() {
		account := p.GetString("account")
		name := p.GetString("name")
		password := p.GetString("password")
		repassword := p.GetString("repassword")
		switch "" {
		case account:
			p.Rsp(false, "账号不能为空")
			return
		case name:
			p.Rsp(false, "用户名不能为空")
			return
		case password:
			p.Rsp(false, "密码不能为空")
			return
		}
		if password != repassword {
			p.Rsp(false, "密码不匹配")
			return
		}
		err := models.UserAdd(&models.ShUser{
			Account:    account,
			Name:       name,
			Password:   password,
			CreateTime: time.Now(),
		})
		if err != nil {
			p.Rsp(false, err.Error())
			return
		}
		p.Rsp(true, "注册成功")
	} else {
		p.TplName = "app/pub/register.html"
	}
}

// 退出登录
func (p *Pub) Loginout() {
	if p.IsPost() {
		p.DelSession("info")
		p.Rsp(true, "退出成功，等待跳转")
	}
}
