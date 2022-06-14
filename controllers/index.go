package controllers

import (
	"shops/models"
)

type Index struct {
	Info
}

// 主页
func (i *Index) Index() {
	i.TplName = "app/index/index.html"
}

func (i *Index) List() {
	if i.IsPost() {
		p, _ := i.GetInt("p")
		name := i.GetString("name")
		id, _ := i.GetInt("goodsid")
		num, messages, err := models.GoodsList(p, id, name)
		if err != nil {
			i.Rsp(false, err.Error())
		}
		i.Rsp(true, map[string]interface{}{
			"total": num,
			"list":  messages,
		})
	}
}
