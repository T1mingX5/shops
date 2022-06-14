package controllers

type Info struct {
	Role string
	Id   int
	CommonController
}

// 检查是否是管理员
func (i *Info) CheckAdminRole() bool {
	info, ok := i.GetSession("info").(Info)
	if !ok || info.Id == 0 {
		return false
	}
	if info.Role == "admin" {
		return true
	}
	return false
}

// 获取id
func (i *Info) GetInfoID() int {
	info, ok := i.GetSession("info").(Info)
	if !ok || info.Id == 0 {
		return 0
	}
	return info.Id
}
