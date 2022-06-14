package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ShAdmin struct {
	Account    string    `orm:"column(account);size(10)" description:"管理员账号"`
	CreateTime time.Time `orm:"column(create_time);type(datetime)"`
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(10)" description:"管理员名称"`
	Password   string    `orm:"column(password);size(100)" description:"管理员密码"`
	Status     int       `orm:"column(status)" description:"管理员状态"`
}

func (t *ShAdmin) TableName() string {
	return "sh_admin"
}

func init() {
	orm.RegisterModel(new(ShAdmin))
}

func (s *ShAdmin) QueryTable() orm.QuerySeter {
	return orm.NewOrm().QueryTable(s)
}

// 添加
func AdminAdd(table *ShAdmin) error {
	o := orm.NewOrm()
	create, _, err := o.ReadOrCreate(table, "account")
	if err != nil {
		return err
	}
	if !create {
		return errors.New("用户名已存在")
	}
	return nil
}

// 管理员列表
func AdminList(id int) ([]*ShAdmin, error) {
	var table ShAdmin
	var list []*ShAdmin
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

//删除
func AdminDelete(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&ShAdmin{Id: id})
	if err != nil {
		return err
	}
	return nil
}

//编辑(启用/禁用)
func AdminEdit(id, status, flag int, account, name, password string) error {
	var table ShAdmin
	qs := table.QueryTable()
	q := table.QueryTable()
	if flag == 0 {
		cond := orm.NewCondition().And("name", name).Or("account", account)
		num, _ := q.SetCond(cond).Count()
		if num > 0 {
			return errors.New("用户名重复")
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

//登录验证
func AdminLogin(account string, password string) (int, int, error) {
	o := orm.NewOrm() // 注册句柄
	table := &ShAdmin{
		Account:  account,
		Password: password,
	}
	if err := o.Read(table, "Account", "Password"); err != nil {
		return 0, 0, err
	}
	return table.Id, table.Status, nil
}

// 查询一个管理员信息
func AdminInfo(id int) (*ShAdmin, error) {
	var table ShAdmin
	qs := table.QueryTable()
	err := qs.Filter("id", id).One(&table)
	if err != nil {
		return nil, err
	}
	return &table, nil
}
