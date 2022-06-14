// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"shops/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// 首页
	web.Router("/", &controllers.Index{}, "get,post:Index")
	web.Router("/list", &controllers.Index{}, "post:List")
	// 公开方法
	web.Router("/login", &controllers.Pub{}, "get,post:Login")       // 登录
	web.Router("/register", &controllers.Pub{}, "get,post:Register") // 注册
	web.Router("/loginout", &controllers.Pub{}, "post:Loginout")     // 退出登录
	// 用户
	web.Router("/user", &controllers.User{}, "get:Index")
	web.Router("/user/info", &controllers.User{}, "post:UserInfo")     // 个人中心
	web.Router("/user/list", &controllers.User{}, "post:UserList")     // 用户列表
	web.Router("/user/delete", &controllers.User{}, "post:UserDelete") // 用户删除
	web.Router("/edit", &controllers.User{}, "get,post:UserEdit")      // 用户编辑
	web.Router("/user/recharge", &controllers.User{}, "post:Recharge") // 充值

	// 订单
	web.Router("/order", &controllers.Order{}, "post:OrderList")          // 订单列表
	web.Router("/user/order", &controllers.Order{}, "get,post:OrderInfo") // 订单信息
	web.Router("/order/status", &controllers.Order{}, "post:OrderStatus")

	// 管理员
	web.Router("/adminlogin", &controllers.Admin{}, "get,post:AdminLogin")    // 登录
	web.Router("/controller", &controllers.Admin{}, "get:Index")              // 加载管理员后台页面
	web.Router("/admin/personal", &controllers.Admin{}, "get,post:AdminInfo") // 管理员信息查询
	web.Router("/admin/delete", &controllers.Admin{}, "post:AdminDelete")
	web.Router("/admin/personal/edit", &controllers.Admin{}, "post:AdminEdit")   // 管理员个人信息编辑
	web.Router("/admin/manage", &controllers.Admin{}, "get,post:AdminManage")    // 管理员管理页面
	web.Router("/admin/manage/edit", &controllers.Admin{}, "post:AdminEdit")     // 管理员管理编辑
	web.Router("/admin/list", &controllers.Admin{}, "post:AdminList")            // 管理员列表
	web.Router("/admin/usermanage", &controllers.Admin{}, "get,post:UserManage") // 管理用户界面
	web.Router("/admin/add", &controllers.Admin{}, "get:AddIndex")               // 管理员添加界面
	web.Router("/admin/useradd", &controllers.User{}, "post:UserAdd")            // 添加新用户
	web.Router("/admin/adminadd", &controllers.Admin{}, "post:AdminAdd")         // 添加新管理员
	web.Router("/admin/goodsadd", &controllers.Goods{}, "post:GoodsAdd")         // 添加新商品
	web.Router("/controller/loginout", &controllers.Admin{}, "post:Loginout")    // 退出
	// 商品
	web.Router("/admin/goodsedit", &controllers.Goods{}, "get,post:GoodsEdit") // 修改商品信息
	web.Router("/goods/delete", &controllers.Goods{}, "post:GoodsDelete")      // 删除商品
	// 购买
	web.Router("/goods/buy", &controllers.Buy{}, "post:GoodsBuy")

	// 流水·
	web.Router("/flows", &controllers.Flows{}, "post:FlowList")
}
