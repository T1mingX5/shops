package controllers

import (
	"fmt"
	"shops/models"
	"strconv"
	"time"
)

// ShAdminController operations for ShAdmin
type Admin struct {
	Info
}

func (a *Admin) Index() {
	if !a.CheckAdminRole() {
		a.Rsp(false, "您没有权限")
		return
	}
	a.TplName = "admin/controller.html"
}

func (a *Admin) AddIndex() {
	if !a.CheckAdminRole() {
		a.Rsp(false, "您没有权限")
		return
	}
	a.TplName = "admin/add.html"
}

// 管理员管理
func (a *Admin) AdminManage() {
	if !a.CheckAdminRole() {
		a.Rsp(false, "您没有权限")
		return
	}
	if a.IsPost() {
		id, _ := a.GetInt("adminid")
		if id == 0 {
			a.Rsp(false, "获取不到id")
			return
		}
		table, err := models.AdminInfo(id)
		if err != nil {
			a.Rsp(false, err.Error())
		}
		fmt.Println("table.name:", table.Name)
		a.Rsp(true, table)

	} else {
		a.TplName = "admin/manage.html"
	}
}

// 管理员列表
func (a *Admin) AdminList() {
	if a.IsPost() {
		adminid := a.GetString("adminid")
		id, _ := strconv.Atoi(adminid)
		table, err := models.AdminList(id)
		if err != nil {
			a.Rsp(false, err.Error())
			return
		}
		a.Rsp(true, map[string]interface{}{
			"list": table,
		})
	}
}

// 添加
func (a *Admin) AdminAdd() {
	if a.IsPost() {
		account := a.GetString("account")
		name := a.GetString("name")
		password := a.GetString("password")
		repassword := a.GetString("repassword")
		status, _ := a.GetInt("status")
		if status != 1 && status != 2 {
			a.Rsp(false, "状态异常")
			return
		}
		switch "" {
		case account:
			a.Rsp(false, "账号不能为空")
			return
		case name:
			a.Rsp(false, "用户名不能为空")
			return
		case password:
			a.Rsp(false, "密码不能为空")
			return
		}
		if password != repassword {
			a.Rsp(false, "密码不匹配")
			return
		}
		err := models.AdminAdd(&models.ShAdmin{
			Account:    account,
			Name:       name,
			Password:   password,
			Status:     status,
			CreateTime: time.Now(),
		})
		if err != nil {
			a.Rsp(false, err.Error())
			return
		}
		a.Rsp(true, "添加成功")
	}
}

// 查询一个管理员的信息
func (a *Admin) AdminInfo() {
	if a.IsPost() {
		id := a.GetInfoID()
		if id == 0 {
			a.Rsp(false, "获取不到id")
			return
		}
		table, err := models.AdminInfo(id)
		if err != nil {
			a.Rsp(false, err.Error())
		}
		fmt.Println("table.name:", table.Name)
		a.Rsp(true, table)
	} else {
		a.TplName = "admin/personal.html"
	}
}

// 删除
func (a *Admin) AdminDelete() {
	if a.IsPost() {
		id, _ := a.GetInt("goodsid")
		if id == 0 {
			a.Rsp(false, "获取不到ID")
			return
		}
		err := models.AdminDelete(id)
		if err != nil {
			a.Rsp(false, err.Error())
			return
		}
		a.Rsp(true, "删除成功")
	}
}

// 编辑(启用/禁用)
func (a *Admin) AdminEdit() {
	if a.IsPost() {
		flag := 0
		id := a.GetInfoID()
		adminid, _ := a.GetInt("adminid")
		if id == 0 && adminid == 0 {
			a.Rsp(false, "获取不到ID")
			return
		}
		newid := id
		if adminid != 0 {
			newid = adminid
		}
		table, err := models.AdminInfo(newid)
		if err != nil {
			a.Rsp(false, err.Error())
			return
		}
		account := a.GetString("account")
		name := a.GetString("name")
		password := a.GetString("password")
		repassword := a.GetString("repassword")
		status, _ := a.GetInt("status")
		if password != repassword {
			a.Rsp(false, "密码不匹配")
			return
		}
		if account == "" {
			account = table.Account
			flag = 1
		}
		if name == "" {
			name = table.Name
			flag = 1
		}
		if password == "" {
			password = table.Password
			flag = 1
		}
		if status != 1 && status != 2 {
			a.Rsp(false, "状态异常")
			return
		}
		err1 := models.AdminEdit(newid, status, flag, account, name, password)
		if err1 != nil {
			a.Rsp(false, err1.Error()) // err.Error()将返回的错误转换为字符串类型
		}
		fmt.Printf("id:%d\n,adminid:%d\n,newid:%d\n", id, adminid, newid)
		if id == adminid {
			a.Rsp(true, "编辑成功,即将退出")
			return
		}
		a.Rsp(true, "编辑成功")
	} else {
		a.TplName = "admin/personal.html"
	}
}

// 用户管理
func (a *Admin) UserManage() {
	if !a.CheckAdminRole() {
		a.Rsp(false, "您没有权限")
		return
	}
	if a.IsPost() {
		id, _ := a.GetInt("userid")
		if id == 0 {
			a.Rsp(false, "获取不到ID")
		}
		table, err := models.UserInfo(id)
		if err != nil {
			a.Rsp(false, err.Error())
			return
		}
		a.Rsp(true, table)
	} else {
		a.TplName = "admin/usermanage.html"
	}

}

// 登录
func (a *Admin) AdminLogin() {
	if a.IsPost() { // 后端是不能主动去请求前端的，这句话的意思是前端“请求(request)”以POST方式给后端传送数据
		account := a.GetString("account") // 后端接受所需数据
		password := a.GetString("password")
		if account == "" || password == "" {
			a.Rsp(false, "输入内容不能为空")
			return
		}
		id, status, _ := models.AdminLogin(account, password)
		if id == 0 {
			a.Rsp(false, "用户名或密码不正确")
			return
		}
		if status == 2 {
			a.Rsp(false, "用户账户被禁用")
			return
		} else if status > 2 || status < 1 {
			a.Rsp(false, "账号异常")
			return
		}
		// token, err := generateToken(id, account)
		// if err != nil {
		// 	u.Rsp(false, err.Error())
		// 	return
		// }
		// a.SetSession("adminid", id)
		a.SetSession("info", Info{
			Role: "admin",
			Id:   id,
		}) // admin-1 admin-2 ...
		a.Rsp(true, "登录成功，等待跳转")
	} else {
		a.TplName = "admin/login.html" // 如果前端传回来的不是post请求，那么就使用get方法，访问这个页面，这个的地址是左边列表目录的地址
	}
}

// 退出登录
func (a *Admin) Loginout() {
	if a.IsPost() {
		a.DelSession("info")
		a.Rsp(true, "退出成功，等待跳转")
	}
}
