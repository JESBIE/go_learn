package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

/*
小项目：并发网站健康检查器（Concurrent Checker）

你会练到的能力：
1) Worker Pool：固定 worker 数并发检查 URL。
2) Context：为整批任务设置总超时，为单个请求设置请求超时。
3) Channel：jobs 分发任务，results 汇总结果（典型 fan-in）。
4) WaitGroup：保证 worker 全部收尾后再关闭 results。

运行方式：
- 直接运行默认目标：go run .
- 传入自定义目标：go run . https://example.com https://golang.org
*/

type Job struct {
	ID  int
	URL string
}

type Result struct {
	JobID      int
	URL        string
	StatusCode int
	Duration   time.Duration
	Err        error
}

func worker(
	ctx context.Context,
	id int,
	client *http.Client,
	jobs <-chan Job,
	results chan<- Result,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[worker-%d] 收到取消信号，退出\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("[worker-%d] jobs 已关闭，退出\n", id)
				return
			}

			res := checkOne(ctx, client, job)

			// 发送结果时也监听 ctx，避免上游已取消时卡住。
			select {
			case <-ctx.Done():
				fmt.Printf("[worker-%d] 发送结果前被取消，退出\n", id)
				return
			case results <- res:
			}
		}
	}
}

func checkOne(parent context.Context, client *http.Client, job Job) Result {
	start := time.Now()

	// 单个请求超时：即使整批超时时间较长，也避免某个 URL 卡太久。
	reqCtx, cancel := context.WithTimeout(parent, 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, job.URL, nil)
	if err != nil {
		return Result{JobID: job.ID, URL: job.URL, Duration: time.Since(start), Err: err}
	}

	resp, err := client.Do(req)
	if err != nil {
		return Result{JobID: job.ID, URL: job.URL, Duration: time.Since(start), Err: err}
	}
	defer resp.Body.Close()

	return Result{
		JobID:      job.ID,
		URL:        job.URL,
		StatusCode: resp.StatusCode,
		Duration:   time.Since(start),
	}
}

func defaultTargets() []string {
	return []string{
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/503",
		"https://this-domain-should-fail.invalid",
		"https://httpbin.org/delay/2",
	}
}

func main() {
	fmt.Println("=== 11 mini project: concurrent checker ===")

	targets := os.Args[1:]
	if len(targets) == 0 {
		targets = defaultTargets()
		fmt.Println("未传入 URL，使用默认目标：", targets)
	}

	// 整批任务总超时：超过这个时间，所有 worker 都会收到取消信号。
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	workerCount := 3
	jobs := make(chan Job, 4)
	results := make(chan Result)
	client := &http.Client{}

	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i := 1; i <= workerCount; i++ {
		go worker(ctx, i, client, jobs, results, &wg)
	}

	// 生产者：发送任务并关闭 jobs。
	go func() {
		defer close(jobs)
		for i, url := range targets {
			select {
			case <-ctx.Done():
				fmt.Println("[producer] 收到取消信号，停止投递任务")
				return
			case jobs <- Job{ID: i + 1, URL: url}:
			}
		}
	}()

	// 所有 worker 完成后关闭 results，主协程可安全 range 结束。
	go func() {
		wg.Wait()
		close(results)
	}()

	var (
		successCount int
		failCount    int
		totalCost    time.Duration
	)

	for r := range results {
		totalCost += r.Duration

		if r.Err != nil {
			failCount++
			fmt.Printf("[FAIL] #%d %s | cost=%v | err=%v\n", r.JobID, r.URL, r.Duration, r.Err)
			continue
		}

		// 把 2xx 和 3xx 视作“可达”。
		if r.StatusCode >= 200 && r.StatusCode < 400 {
			successCount++
			fmt.Printf("[ OK ] #%d %s | status=%d | cost=%v\n", r.JobID, r.URL, r.StatusCode, r.Duration)
		} else {
			failCount++
			fmt.Printf("[BAD ] #%d %s | status=%d | cost=%v\n", r.JobID, r.URL, r.StatusCode, r.Duration)
		}
	}

	total := successCount + failCount
	avg := time.Duration(0)
	if total > 0 {
		avg = totalCost / time.Duration(total)
	}

	fmt.Println("---------- 汇总 ----------")
	fmt.Printf("总数=%d 成功=%d 失败=%d 平均耗时=%v\n", total, successCount, failCount, avg)
	if err := ctx.Err(); err != nil {
		fmt.Println("批任务结束原因：", err)
	}
}

