package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ShGoods struct {
	Content    string    `orm:"column(content)" description:"内容"`
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"创建时间"`
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(20)" description:"名字"`
	Number     int       `orm:"column(number)" description:"库存"`
	Price      float64   `orm:"column(price);digits(10);decimals(2)" description:"价格"`
	Status     int8      `orm:"column(status)" description:"状态,0为下架,1为上架"`
}

func (s *ShGoods) TableName() string {
	return "sh_goods"
}

func init() {
	orm.RegisterModel(new(ShGoods))
}

func (s *ShGoods) QueryTable() orm.QuerySeter {
	return orm.NewOrm().QueryTable(s)
}

// 删除
func GoodsDelete(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&ShGoods{Id: id})
	if err != nil {
		return err
	}
	return nil
}

// 添加
func GoodsAdd(table *ShGoods) error {
	o := orm.NewOrm()
	_, err := o.Insert(table)
	if err != nil {
		return err
	}
	return nil
}

// 编辑
func GoodsEdit(table *ShGoods) error {
	goods, err := GoodsInfo(table.Id)
	if err != nil {
		return err
	}
	if table.Name == "" {
		table.Name = goods.Name
	}
	if table.Price == 0 {
		table.Price = goods.Price
	}
	if table.Number == 0 {
		table.Number = goods.Number
	}
	if table.Content == "" {
		table.Content = goods.Content
	}
	if table.Status == 0 {
		table.Status = goods.Status
	}
	qs := table.QueryTable()
	_, err1 := qs.Filter("id", table.Id).Update(orm.Params{
		"name":    table.Name,
		"price":   table.Price,
		"number":  table.Number,
		"content": table.Content,
		"status":  table.Status,
	})
	if err1 != nil {
		return err1
	}
	return nil
}

// 列表
func GoodsList(p, id int, name string) (int64, []*ShGoods, error) {
	// o := orm.NewOrm()
	// qs := o.QueryTable("sh_goods")
	var table ShGoods
	var list []*ShGoods
	qs := table.QueryTable()

	if name != "" {
		qs = qs.Filter("name__icontains", name)
	}
	if id != 0 {
		qs = qs.Filter("id__in", id)
	}
	num1, _ := qs.Count()
	qs = qs.OrderBy("-id")
	if p <= 0 {
		p = 1
	}
	num := (p - 1) * 20 // 这里的数字20，是根据当前页面的显示条数改变的，当前页面要求显示20条，这里就*20
	qs = qs.Limit(20, num)
	_, err := qs.All(&list)
	if err != nil {
		return 0, nil, err
	}
	return num1, list, nil
}

// 查询一个商品
func GoodsInfo(goodsid int) (*ShGoods, error) {
	var table ShGoods
	qs := table.QueryTable()
	err := qs.Filter("id", goodsid).One(&table)
	if err != nil {
		return nil, err
	}
	return &table, nil
}

// 购买
func GoodsBuy(goodsid, num int, user *ShUser) (string, float64, error) {
	o := orm.NewOrm()
	table := ShGoods{Id: goodsid}
	if err := o.Read(&table); err != nil { // 通过查询数据库的表单的ID，判断是否有这一个商品，如果没有则不进行编辑
		return "", 0, err
	}
	num1 := table.Number - num
	if num1 < 0 {
		return "", 0, errors.New("库存不足")
	}
	if table.Status != 1 {
		return "", 0, errors.New("商品已下架")
	}
	sumprice := table.Price * float64(num)
	fmt.Printf("table.number:%v\nnum1%v\ntable.status:%v\n:商品单价是:%v\n购买的总价是:%v\n购买的数量是:%v\n用户的钱:%v\n", table.Number, num1, table.Status, table.Price, sumprice, num, user.Money)
	if user.Money < sumprice {
		return "", 0, errors.New("余额不足")
	}
	table.Number = num1
	if _, err := o.Update(&table, "Number"); err != nil { // update方法更新指定表的指定字段的内容
		return "", 0, err
	}

	return table.Name, sumprice, nil
}
