package main

import (
	"context"
	"fmt"
	"sync"
)

func slicetochannel(my_slice []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range my_slice {
			out <- v
		}
	}()

	return out
}

func squareWorker(ctx context.Context, id int, in <-chan int, out chan<- int) {
	for {

		select {
		case <-ctx.Done():
			return
		case v, ok := <-in:
			if !ok {
				return
			}
			fmt.Println("squareWorker", id, "got v: ", v)
			out <- v * v

		}
	}
}

func main() {
	var ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	my_slice := []int{2, 4, 1, 8, 4, 7}

	in := slicetochannel(my_slice)
	out := make(chan int)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			squareWorker(ctx, id, in, out)
		}(i)
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()
	count := 0
	for result := range out {
		fmt.Println(result)
		count++
		if count >= 3 {
			cancel()
			break
		}

	}

}
