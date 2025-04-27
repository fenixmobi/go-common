package crontab

import (
	"sync"

	"github.com/spf13/cobra"
)

var root *cobra.Command

var once sync.Once

func New() *cobra.Command {
	once.Do(func() {
		root = &cobra.Command{
			Use:   "crontab",
			Short: "定时任务系统",
		}

		// 添加全局参数
		root.PersistentFlags().StringP("date", "d", "", "日期,格式为YYYY-MM-DD")
		root.PersistentFlags().Uint8P("date-flag", "f", 0, "日期标志, -1 昨天, 0 当天, 1 指定日期")
		root.PersistentFlags().StringSliceP("app-list", "a", []string{}, "应用列表")

		// 配置参数
		root.PersistentFlags().StringP("conf", "c", "config.yaml", "配置文件")
	})

	return root
}

func Execute() error {
	return root.Execute()
}

func AddCommand(runner *Runner) {
	cmd := cobra.Command{
		Use:   runner.Cmd,
		Short: runner.Remark,
		Run:   runner.Run,
	}

	// 注册局部命令
	for _, flag := range runner.Flags {
		switch flag.Type {
		case ArgTypeInt:
			cmd.Flags().IntP(flag.Name, string(flag.ShortName), 0, flag.Remark)
		case ArgTypeString:
			cmd.Flags().StringP(flag.Name, string(flag.ShortName), "", flag.Remark)
		case ArgTypeSlice:
			cmd.Flags().StringSliceP(flag.Name, string(flag.ShortName), []string{}, flag.Remark)
		}
	}

	root.AddCommand(&cmd)
}
