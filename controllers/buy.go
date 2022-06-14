package controllers

import (
	"fmt"
	"shops/models"
	"time"
)

type Buy struct {
	Info
}

func (b *Buy) GoodsBuy() {
	if b.IsPost() {
		userid := b.GetInfoID()
		num, _ := b.GetInt("num")
		goodsid, _ := b.GetInt("goodsid")
		if num == 0 {
			num = 1
		}
		if goodsid == 0 {
			b.Rsp(false, "订单状态异常")
			return
		}
		// 调用shuser数据库
		user, err := models.UserInfo(userid)
		if err != nil {
			b.Rsp(false, err.Error())
			return
		}
		// 调用shgoods数据库
		goodsname, sumprice, err := models.GoodsBuy(goodsid, num, user)
		if err != nil {
			b.Rsp(false, err.Error())
			return
		}
		// 调用shuser数据库
		note, err1 := models.ChangeMoney(userid, sumprice, "-")
		if err1 != nil {
			b.Rsp(false, err1.Error())
			return
		}
		// 向订单表插入一条订单
		time := time.Now()
		fmt.Printf("goodsid:%v\nstauts:%v\nuserid:%v\n,goodscount:%v\n", &models.ShGoods{Id: goodsid}, 0, &models.ShUser{Id: userid}, num)
		order := models.ShOrder{
			CreateTime: time,
			GoodsId:    &models.ShGoods{Id: goodsid},
			Status:     0,
			UserId:     &models.ShUser{Id: userid},
			GoodsCount: num,
		}
		err2 := models.OrderAdd(order)
		if err2 != nil {
			b.Rsp(false, err2.Error())
			return
		}
		// 插入一条流水
		flow := models.ShFlow{
			CreateTime: time,
			UserId:     &models.ShUser{Id: userid},
			Money:      -note.Money,
			EndMoney:   note.EndMoney,
			Cate:       "购买" + goodsname,
		}
		err3 := models.FlowAdd(&flow)
		if err3 != nil {
			b.Rsp(false, err3.Error())
			return
		}
		b.Rsp(true, "购买成功")
	}
}
