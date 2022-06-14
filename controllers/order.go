package controllers

import (
	"fmt"
	"shops/models"
	"time"
)

type Order struct {
	Info
}

// 添加

// 删除

// 列表
func (o *Order) OrderList() {
	if o.IsPost() {
		id := o.GetInfoID()
		if id == 0 {
			o.Rsp(false, "验证失败")
			return
		}

		p, _ := o.GetInt("p")
		num, table, err := models.OrderList(p, id)
		if err != nil {
			o.Rsp(false, err.Error())
			return
		}
		o.Rsp(true, map[string]interface{}{
			"total": num,
			"list":  table,
		})
	}
}

// 显示一个订单信息
func (o *Order) OrderInfo() {
	if o.IsPost() {
		id := o.GetInfoID() // 获取用户id
		if id == 0 {
			o.Rsp(false, "验证失败")
			return
		}
		goodsid, _ := o.GetInt("goodsid")
		// if err != nil {
		// 	g.Rsp(false, err)
		// }
		if goodsid == 0 {
			o.Rsp(false, "获取不到商品")
			return
		}
		table, err := models.GoodsInfo(goodsid)
		if err != nil {
			o.Rsp(false, err.Error())
			return
		}
		o.Rsp(true, table)
	} else {
		o.TplName = "user/order.html"
	}
}

// 更改状态
func (o *Order) OrderStatus() {
	if o.IsPost() {
		orderid, _ := o.GetInt("orderid")
		symbol := o.GetString("symbol")
		userid := o.GetInfoID()
		if orderid == 0 {
			o.Rsp(false, "获取不到订单id")
			return
		}
		if symbol == "" {
			o.Rsp(false, "状态获取错误")
		}
		fmt.Printf("symbol是:%v\n", symbol)
		flag, err := models.OrderStatus(orderid, symbol)
		if err != nil {
			o.Rsp(false, err.Error())
			return
		}
		var str string
		if flag == 1 {
			str = "已确认收货"
		}
		// 事务，隔离级别
		if flag == 2 {
			str = "已退货"
			table, err1 := models.OrderInfo(orderid) // 通过订单id，查询商品金额
			if err1 != nil {
				o.Rsp(false, err1.Error())
				return
			}
			// oo := orm.NewOrm()
			// tx, err := oo.Begin()
			// if err != nil {
			// 	o.Rsp(false, err)
			// 	return
			// }
			money := table.GoodsId.Price
			note, err2 := models.ChangeMoney(userid, money, "+") // 更改账户余额
			if err2 != nil {
				o.Rsp(false, err2.Error())
				return
			}
			// 添加商品数量
			num := table.GoodsId.Number + table.GoodsCount
			goodstable := &models.ShGoods{
				Id:     table.GoodsId.Id,
				Number: num,
			}
			err3 := models.GoodsEdit(goodstable)
			if err3 != nil {
				// tx.Rollback()
				o.Rsp(false, err3.Error())
				return
			}
			// 插入一条流水
			time := time.Now()
			flow := models.ShFlow{
				CreateTime: time,
				UserId:     &models.ShUser{Id: userid},
				Money:      note.Money,
				EndMoney:   note.EndMoney,
				Cate:       "退还" + table.GoodsId.Name + "金额",
			}
			err4 := models.FlowAdd(&flow)
			if err4 != nil {
				o.Rsp(false, err4.Error())
				return
			}
		}
		o.Rsp(true, str)
	}
}
