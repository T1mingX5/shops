package controllers

import (
	"shops/models"
	"strconv"
	"time"
)

// ShGoodsController operations for ShGoods
type Goods struct {
	Info
}

// 删除
func (g *Goods) GoodsDelete() {
	if g.IsPost() {
		id, _ := g.GetInt("goodsid")
		if id == 0 {
			g.Rsp(false, "获取不到ID")
			return
		}
		err := models.GoodsDelete(id)
		if err != nil {
			g.Rsp(false, err.Error())
			return
		}
		g.Rsp(true, "删除成功")
	}
}

// 添加
func (g *Goods) GoodsAdd() {
	if g.IsPost() {
		name := g.GetString("name")
		content := g.GetString("content")
		price, _ := g.GetFloat("price")
		number, _ := g.GetInt("number")
		status, _ := g.GetInt8("status")
		if status != 1 && status != 2 {
			g.Rsp(false, "状态异常")
			return
		}
		if name == "" || content == "" {
			g.Rsp(false, "输入内容不能为空")
		}
		if price == 0.0 {
			g.Rsp(false, "输入价格不能为空")
		}
		err := models.GoodsAdd(&models.ShGoods{
			Name:       name,
			Price:      price,
			Number:     number,
			Content:    content,
			Status:     status,
			CreateTime: time.Now(),
		})
		if err != nil {
			g.Rsp(false, err.Error())
			return
		}
		g.Rsp(true, "添加成功")
	}
}

// 编辑
func (g *Goods) GoodsEdit() {
	if !g.CheckAdminRole() {
		g.Rsp(false, "您没有权限")
		return
	}
	if g.IsPost() {
		id, _ := g.GetInt("goodsid")
		if id == 0 {
			g.Rsp(false, "获取不到id")
		}
		name := g.GetString("name")
		str := g.GetString("price")
		number, _ := g.GetInt("number")
		content := g.GetString("content")
		status, _ := g.GetInt("status")
		price, _ := strconv.ParseFloat(str, 64)
		table := &models.ShGoods{
			Id:      id,
			Name:    name,
			Price:   price,
			Number:  number,
			Content: content,
			Status:  int8(status),
		}
		err := models.GoodsEdit(table)
		if err != nil {
			g.Rsp(false, err.Error())
		}
		g.Rsp(true, "更改成功")
	} else {
		g.TplName = "admin/goodsedit.html"
	}
}

// 查询一个商品信息
