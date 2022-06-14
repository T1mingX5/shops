package controllers

import "shops/models"

type Flows struct {
	Info
}

func (f *Flows) FlowList() {
	if f.IsPost() {
		userid := f.GetInfoID()
		id, _ := f.GetInt("userid")
		if userid == 0 && id == 0 {
			f.Rsp(false, "验证错误")
		}
		newid := userid
		if id != 0 {
			newid = id
		}

		flows, err := models.FlowList(newid)
		if err != nil {
			f.Rsp(false, err.Error())
			return
		}
		f.Rsp(true, map[string]interface{}{
			"list": flows,
		})
	}
}
