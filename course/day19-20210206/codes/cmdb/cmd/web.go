package cmd

import (
	_ "cmdb/routers"
	"cmdb/services"
	"log"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"

	"github.com/astaxie/beego/orm"

	_ "github.com/astaxie/beego/session/redis"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "cmdb web",
	Long:  "cmdb program web",
	RunE: func(cmd *cobra.Command, args []string) error {
		beego.LoadAppConfig("ini", "conf/web.conf")
		if beego.AppConfig.String("RunMode") == "dev" {
			orm.Debug = true
		}
		dsn := beego.AppConfig.String("db::dsn")

		orm.RegisterDriver("mysql", orm.DRMySQL)
		if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
			return err
		}

		orm.RunSyncdb("default", false, true)

		if user := services.GetUserByName("admin"); user == nil {
			log.Print("创建管理员账号")
			services.AddUser("admin", "123@abc", "西安", true)
		}

		beego.BConfig.Log.AccessLogs = true
		beego.BConfig.Log.FileLineNum = true
		beego.BeeLogger.DelLogger("console")
		beego.SetLogger("file", `{
			"level": 7,
			"filename": "logs/web.log",
			"daily": true,
			"maxdays": 15
		}`)
		beego.AddFuncMap("upper", strings.ToUpper)
		beego.AddFuncMap("kk", func(t string) string {
			return "kk:" + t
		})
		beego.Run()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
