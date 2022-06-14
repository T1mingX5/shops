package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ShUser struct {
	Account    string    `orm:"column(account);size(10)" description:"账号"`
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"创建时间"`
	Id         int       `orm:"column(id);auto" description:"用户id"`
	Money      float64   `orm:"column(money);digits(10);decimals(2)" description:"用户账户余额"`
	Name       string    `orm:"column(name);size(10)" description:"用户名称"`
	Password   string    `orm:"column(password);size(100)" description:"用户密码"`
	Status     int       `orm:"column(status)" description:"用户状态"`
}

func (t *ShUser) TableName() string {
	return "sh_user"
}

func (s *ShUser) QueryTable() orm.QuerySeter {
	return orm.NewOrm().QueryTable(s)
}

func init() {
	orm.RegisterModel(new(ShUser))
}

// 用户列表
func UserList(id int) ([]*ShUser, error) {
	var table ShUser
	var list []*ShUser
	qs := table.QueryTable()
	if id != 0 {
		qs = qs.Filter("id__in", id)
		qs = qs.OrderBy("-id")
	}
	_, err := qs.All(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 添加
func UserAdd(table *ShUser) error {
	o := orm.NewOrm()
	err1 := o.Read(table, "account")
	if err1 == nil {
		return errors.New("账号已存在")
	}
	err2 := o.Read(table, "name")
	if err2 == nil {
		return errors.New("用户名已存在")
	}
	_, err3 := o.Insert(table)
	if err3 != nil {
		return err3
	}
	return nil
}

// 登录验证
func UserLogin(account string, password string) (int, int, error) {
	o := orm.NewOrm() // 注册句柄
	table := &ShUser{
		Account:  account,
		Password: password,
	}
	if err := o.Read(table, "Account", "Password"); err != nil {
		return 0, 0, err
	}
	return table.Id, table.Status, nil
}

// 删除
func UserDelete(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&ShUser{Id: id})
	if err != nil {
		return err
	}
	return nil
}

// 显示一个用户信息
func UserInfo(id int) (*ShUser, error) {
	o := orm.NewOrm()
	var tabel ShUser
	err := o.QueryTable("sh_user").Filter("Id", id).One(&tabel)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
	}
	return &tabel, nil
}

// 编辑用户信息
func UserEdit(id, status, flag int, account, name, password string) error {
	var table ShUser
	qs := table.QueryTable()
	q := table.QueryTable()
	if flag == 0 {
		cond := orm.NewCondition().And("name", name).Or("account", account)
		num, _ := q.SetCond(cond).Count()
		if num > 0 {
			return errors.New("用户名或账号重复")
		}
	}

	_, err1 := qs.Filter("id", id).Update(orm.Params{
		"account":  account,
		"name":     name,
		"password": password,
		"status":   status,
	})

	if err1 != nil {
		return err1
	}
	return nil
}

type Note struct {
	Money    float64
	EndMoney float64
}

// 更改用户金额
// 开启事务 o orm.TxOrmer,
func ChangeMoney(id int, money float64, symbol string) (*Note, error) {
	// var o interface{}
	// if o == nil {
	// 	o = orm.NewOrm()
	// }
	o := orm.NewOrm()
	user := ShUser{
		Id: id,
	}
	err := o.Read(&user)
	if err != nil {
		return nil, err
	}
	if symbol == "-" {
		if user.Money < money {
			return nil, errors.New("账户余额不足")
		}
		user.Money = user.Money - money
	}
	if symbol == "+" {
		user.Money = user.Money + money
	}
	_, err1 := o.Update(&user, "Money")
	if err1 != nil {
		return nil, err1
	}
	note := &Note{
		Money:    money,
		EndMoney: user.Money,
	}
	return note, nil
}
