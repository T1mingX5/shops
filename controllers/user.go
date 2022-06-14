package controllers

import (
	"errors"
	"fmt"
	"shops/models"
	"strconv"
	"time"
)

type User struct {
	Info
}

func (u *User) Index() {
	u.TplName = "user/user.html"
}

// 用户列表
func (u *User) UserList() {
	if u.IsPost() {
		usersid := u.GetString("usersid")
		id, _ := strconv.Atoi(usersid)
		table, err := models.UserList(id)
		if err != nil {
			u.Rsp(false, err.Error())
			return
		}
		u.Rsp(true, map[string]interface{}{
			"list": table,
		})
	}
}

// 添加
func (u *User) UserAdd() {
	if u.IsPost() {
		account := u.GetString("account")
		name := u.GetString("name")
		password := u.GetString("password")
		repassword := u.GetString("repassword")
		status, _ := u.GetInt("status")
		if status != 1 && status != 2 {
			u.Rsp(false, "状态异常")
			return
		}
		switch "" {
		case account:
			u.Rsp(false, "账号不能为空")
			return
		case name:
			u.Rsp(false, "用户名不能为空")
			return
		case password:
			u.Rsp(false, "密码不能为空")
			return
		}
		if password != repassword {
			u.Rsp(false, "密码不匹配")
			return
		}
		err := models.UserAdd(&models.ShUser{
			Account:    account,
			Name:       name,
			Password:   password,
			Status:     status,
			CreateTime: time.Now(),
		})
		if err != nil {
			u.Rsp(false, err.Error())
			return
		}
		u.Rsp(true, "添加成功")
	}
}

// 删除
func (u *User) UserDelete() {
	if u.IsPost() {
		id, _ := u.GetInt("goodsid")
		if id == 0 {
			u.Rsp(false, "获取不到ID")
			return
		}
		err := models.UserDelete(id)
		if err != nil {
			u.Rsp(false, err.Error())
			return
		}
		u.Rsp(true, "删除成功")
	}
}

// 编辑
func (u *User) UserEdit() {
	if u.IsPost() {
		flag := 0
		id := u.GetInfoID()
		userid, _ := u.GetInt("userid")
		if id == 0 && userid == 0 {
			u.Rsp(false, "获取不到id")
			return
		}
		newid := id
		if userid != 0 {
			newid = userid
		}
		table, err := models.UserInfo(newid)
		if err != nil {
			u.Rsp(false, err.Error())
			return
		}
		account := u.GetString("account")
		name := u.GetString("name")
		password := u.GetString("password")
		repassword := u.GetString("repassword")
		status, _ := u.GetInt("status")
		if password != repassword {
			u.Rsp(false, "密码不匹配")
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
			u.Rsp(false, "状态异常")
			return
		}
		err1 := models.UserEdit(newid, status, flag, account, name, password)
		if err1 != nil {
			u.Rsp(false, err1.Error()) // err.Error()将返回的错误转换为字符串类型
		}
		u.Rsp(true, "编辑成功")
	} else {
		u.TplName = "user/edit.html"
	}
}

// 用户信息
func (u *User) UserInfo() {
	if u.IsPost() {
		id := u.GetInfoID()
		if id == 0 {
			u.Rsp(false, "访问失败")
			return
		}
		table, err := models.UserInfo(id)
		if err != nil {
			u.Rsp(false, err.Error())
			return
		}
		u.Rsp(true, table)
	}
}

// 充值
func (u *User) Recharge() {
	if u.IsPost() {
		id := u.GetInfoID()
		if id == 0 {
			u.Rsp(false, "访问失败")
			return
		}
		str := u.GetString("money")
		if str == "" {
			u.Rsp(false, errors.New("出现错误"))
		}
		fmt.Printf("传过来的字符%v:\n", str)
		money, _ := strconv.ParseFloat(str, 64)
		fmt.Printf("金额是%f:\n", money)
		if money < 0 {
			u.Rsp(false, "金额充值错误")
		}
		note, err := models.ChangeMoney(id, money, "+")
		if err != nil {
			u.Rsp(false, err.Error())
		}
		// 插入一条流水
		time := time.Now()
		flow := models.ShFlow{
			CreateTime: time,
			UserId:     &models.ShUser{Id: id},
			Money:      note.Money,
			EndMoney:   note.EndMoney,
			Cate:       "充值",
		}
		err3 := models.FlowAdd(&flow)
		if err3 != nil {
			u.Rsp(false, err3.Error())
			return
		}
		u.Rsp(true, "充值成功")
	}
}
