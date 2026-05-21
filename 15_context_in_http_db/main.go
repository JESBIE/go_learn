package main

import (
	"context"
	"fmt"
	"time"
)

func simulateDB(ctx context.Context) error {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return ctx.Err()
	case <-time.After(2 * time.Second):
		return nil
	}
}

func runQuery(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return simulateDB(ctx)
}

func runWithCancel() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		errCh <- simulateDB(ctx)
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("cancel query after 500ms")
	cancel()

	err := <-errCh
	if err != nil {
		fmt.Println("DB operation aborted:", err)
	}
	return err
}

func main() {
	if err := runWithCancel(); err != nil {
		fmt.Println("main received:", err)
	}
}
