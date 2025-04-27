package crontab

import (
	"github.com/spf13/cobra"
)

// 命令类型
const (
	_ = iota
	ArgTypeString
	ArgTypeInt
	ArgTypeSlice
)

type Runner struct {
	Cmd    string                                  `desc:"命令的名称"`
	Remark string                                  `desc:"命令的介绍说明"`
	Run    func(cmd *cobra.Command, args []string) `desc:"命令执行的业务逻辑"`
	Flags  []Flag
}

type Flag struct {
	Name      string `desc:"标志名称"`
	ShortName byte   `desc:"标志短名,只能是一个字符"`
	Type      int    `desc:"标志类型"`
	Remark    string `desc:"命令介绍"`
}
