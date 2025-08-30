package taskThree

import (
	"context"
	"fmt"
	"time"
)

// StartWorkers запускает workersAmount горутин, которые конкурентно выводят в stdout значения.
// После чего основная горутина функции пишет в канал значения с интервалом во времени.
// При получении сигнала от ctx горутины должны корректно завершить работу.
func StartWorkers(workersAmount int, ctx context.Context) {
	c := make(chan int)
	for i := 0; i < workersAmount; i++ {
		go func(num int) {
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-c:
					if !ok {
						return
					}
					fmt.Printf("goroutine %v: val = %v\n", num, val)
				}
			}

		}(i)
	}
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			close(c)
			return
		case <-ticker.C:
			c <- 1
		}
	}
}
