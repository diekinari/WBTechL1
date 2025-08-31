package taskFive

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Timeout(n time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), n)
	defer cancel()
	c := make(chan int)
	defer close(c)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Closing writer")
				return
			case <-ticker.C:
				c <- 1
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Closing reader")
				return
			case val := <-c:
				fmt.Printf("val: %v\n", val)
			}
		}
	}()
	wg.Wait()
}
