# corntab

定时任务

# 使用方式
    go get git.qdreads.com/gotools/corntab


    1. 创建定时任务
    corntab.New()
    2. 注册任务
    corntab.Register(...)
    3. 执行任务
    corntab.Execute()

    定义任务
    runner := &corntab.Runner{
        Cmd:    "report-list",
        Remark: "报表集合"
        Run    func(cmd *cobra.Command, args []string) {
            // 任务逻辑
        }
        Flags  []corntab.Flag{
            ....
        }
    }