# 11_mini_project_concurrent_checker

并发小项目：网站健康检查器。

## 目标

- 并发检查一组 URL 是否可达
- 统计成功/失败与平均耗时
- 练习 `worker pool + context + channel + WaitGroup`

## 运行

进入目录后：

```bash
go run .
```

传入自定义 URL：

```bash
go run . https://example.com https://golang.org
```

## 你可以尝试改造

1. 增加 `-workers` 参数，支持动态 worker 数
2. 增加重试（失败后重试 1~2 次）
3. 把结果写入 CSV 文件
4. 加一个总超时参数（目前代码里是固定 8 秒）

