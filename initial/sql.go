package initial

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	user, _ := web.AppConfig.String("db_user") // db_user这些是在conf文件中的配置文件中配置的
	passwd, _ := web.AppConfig.String("db_pass")
	host, _ := web.AppConfig.String("db_host")
	port, err := web.AppConfig.Int("db_port")
	if err != nil {
		port = 3306
	}
	dbname, _ := web.AppConfig.String("db_name")
	runmode, _ := web.AppConfig.String("runmode")
	if runmode == "dev" {
		orm.Debug = true
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)                                                                                              // 注册一个mysql驱动
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", user, passwd, host, port, dbname)) // 注册一个数据库驱动
	fmt.Println("连接数据库成功")
}
