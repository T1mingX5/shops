package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ShFlow struct {
	Cate       string    `orm:"column(cate);size(10)" description:"分类"`
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"创建时间"`
	EndMoney   float64   `orm:"column(end_money);digits(10);decimals(2)" description:"剩余金额"`
	Id         int       `orm:"column(id);pk"`
	Money      float64   `orm:"column(money);digits(10);decimals(2)" description:"操作金额"`
	UserId     *ShUser   `orm:"column(user_id);rel(fk)" description:"用户"`
}

func (s *ShFlow) TableName() string {
	return "sh_flow"
}

func init() {
	orm.RegisterModel(new(ShFlow))
}

func (s *ShFlow) QueryTable() orm.QuerySeter {
	return orm.NewOrm().QueryTable(s)
}

// 添加
func FlowAdd(table *ShFlow) error {
	o := orm.NewOrm()
	_, err := o.Insert(table)
	if err != nil {
		return err
	}
	return nil
}

// 删除

// 列表
func FlowList(userid int) ([]*ShFlow, error) {
	var table ShFlow
	var list []*ShFlow
	qs := table.QueryTable().Filter("user_id", userid).OrderBy("-id")
	// 关联查询
	_, err := qs.All(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
