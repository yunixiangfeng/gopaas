package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	new2 "github.com/yunixiangfeng/gopaas/yu-tool/module"
)

//自动生成 service 目录
var new = &cobra.Command{
	Use:   "new",
	Short: "定制服务，自动生成Service目录",
	Long:  `该命令能够自动生成项目目录，方便快速创建基础项目代码`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("提醒您，请输入项目名称！")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		new2.NewServiceProject(cmd, args)
	},
}

//自动生成 service 目录
var newService = &cobra.Command{
	Use:   "newService",
	Short: "定制服务，自动生成Service目录",
	Long:  `该命令能够自动生成项目目录，方便快速创建基础项目代码`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("提醒您，请输入项目名称！")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		new2.NewServiceProject(cmd, args)
	},
}

//自动生成基础接口程序

var createApi = &cobra.Command{
	Use:   "createApi",
	Short: "定制服务，自动生成 API 目录",
	Long:  `该命令能够自动生成项目目录，方便快速创建基础项目代码`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("提醒您，请输入项目名称！")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		new2.NewApiProject(cmd, args)
	},
}

func init() {

	//添加命令
	rootCmd.AddCommand(new)
	rootCmd.AddCommand(newService)
	rootCmd.AddCommand(createApi)
	//设置flags
	//newCmd.PersistentFlags().String("service", "s", "自动生成 service 项目代码")

}
