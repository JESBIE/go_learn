package main

import (
	"fmt"
	"sync"
	"time"
)

/*
第 10 章（高级并发）- 02：Worker Pool（固定工人池）

核心思想：
1) jobs channel 里放“任务”。
2) 启动固定数量 worker 并发消费任务。
3) 用 WaitGroup 等待 worker 全部结束。
4) 关闭 jobs channel 表示“没有新任务了”。

适用场景：
- 批量图片处理、批量发请求、批量写库。
- 通过“固定 worker 数”限制并发度，避免把 CPU/DB 打满。
*/

type Job struct {
	ID      int
	Payload int
}

type Result struct {
	JobID int
	Value int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// 模拟计算耗时
		time.Sleep(1200 * time.Millisecond)
		out := job.Payload * job.Payload
		fmt.Printf("[worker-%d] 完成 job=%d, payload=%d, result=%d\n", id, job.ID, job.Payload, out)
		results <- Result{JobID: job.ID, Value: out}
	}
	fmt.Printf("[worker-%d] jobs 已关闭，结束\n", id)
}

func main() {
	fmt.Println("=== 02 worker pool ===")

	workerCount := 3
	jobCount := 10

	jobs := make(chan Job, 4)    // 带缓冲：降低生产者阻塞概率
	results := make(chan Result) // 结果也可设缓冲，这里为了演示先用无缓冲

	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i := 1; i <= workerCount; i++ {
		go worker(i, jobs, results, &wg)
	}

	// 生产者：发送任务
	go func() {
		for j := 1; j <= jobCount; j++ {
			jobs <- Job{ID: j, Payload: j + 1}
		}
		close(jobs) // 关键：告诉 worker 不会再有新任务
	}()

	// 结果收集：为了简单，已知任务数就收 jobCount 次
	for i := 0; i < jobCount; i++ {
		r := <-results
		fmt.Printf("[collector] 收到结果: job=%d value=%d\n", r.JobID, r.Value)
	}

	// 等全部 worker 收尾
	wg.Wait()
	fmt.Println("所有任务处理完成")
}
