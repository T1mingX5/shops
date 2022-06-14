package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ShOrder struct {
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"创建时间"`
	GoodsId    *ShGoods  `orm:"column(goods_id);rel(fk)" description:"商品"`
	Id         int       `orm:"column(id);pk"`
	Status     int8      `orm:"column(status)" description:"状态,0为未发货,1为已发货,2为拒绝,3为已完成"`
	UserId     *ShUser   `orm:"column(user_id);rel(fk)" description:"用户"`
	GoodsCount int
}

func (t *ShOrder) TableName() string {
	return "sh_order"
}

func (s *ShOrder) QueryTable() orm.QuerySeter {
	return orm.NewOrm().QueryTable(s)
}

func init() {
	orm.RegisterModel(new(ShOrder))
}

// 添加
func OrderAdd(table ShOrder) error {
	o := orm.NewOrm()
	_, err := o.Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

// 删除

// 列表
func OrderList(p, id int) (int64, []*ShOrder, error) {
	var list []*ShOrder
	o := orm.NewOrm()
	qs := o.QueryTable("sh_order").Filter("user_id", id).OrderBy("-id") // 查询数据库里的字段为"user_id"=id的信息，与关联查询无关
	num1, _ := qs.Count()
	if p <= 0 {
		p = 1
	}
	num := (p - 1) * 10 // 这里的数字10，是根据当前页面的显示条数改变的，当前页面要求显示10条，这里就*10
	qs = qs.Limit(10, num)
	// 关联查询
	_, err := qs.RelatedSel().All(&list) // 这里用到关联查询是因为前端需要与这些信息相关的另一张表的一些信息，如商品名称，只有在商品表里有
	if err != nil {
		return 0, nil, err
	}
	return num1, list, nil
}

// 更改状态
func OrderStatus(id int, symbol string) (int, error) {
	o := orm.NewOrm()
	var table ShOrder
	table.Id = id
	fmt.Printf("数据库接受到的symbol值是:%v\nid:%v\n", symbol, table.Id)
	err1 := o.Read(&table)
	fmt.Printf("数据库接受到的symbol值是:%v\nstatus:%v\n", symbol, table.Status)
	if err1 != nil {
		return 0, err1
	}
	if table.Status == 2 {
		return 0, errors.New("商品已退货")
	}
	if table.Status == 3 {
		return 0, errors.New("商品已签收")
	}
	if table.Status == 1 {
		return 0, errors.New("商品已发货,如退货需自己承担运费 ")
	}
	flag := 0
	if symbol == "收货" {
		table = ShOrder{
			Id:     id,
			Status: 3,
		}
		flag = 1
	}
	if symbol == "退货" {
		table = ShOrder{
			Id:     id,
			Status: 2,
		}
		flag = 2
	}
	_, err := o.Update(&table, "status")
	fmt.Println("table.status:", table.Status)
	if err != nil {
		return 0, err
	}
	return flag, nil
}

// 查询一条订单
func OrderInfo(orderid int) (*ShOrder, error) {
	var table ShOrder
	err := table.QueryTable().RelatedSel().Filter("id", orderid).One(&table)
	fmt.Printf("table.goodsid.content:%v\n", table.GoodsId.Content)
	if err != nil {
		return nil, err
	}
	return &table, nil
}
