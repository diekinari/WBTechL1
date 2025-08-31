package taskSix

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func EndGoroutines() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	// Обыкновенное завершение кода
	go func() {
		fmt.Println("Goroutine 1 ended")
		wg.Done()
	}()

	wg.Add(1)
	// Завершение по каналу
	stop := make(chan bool)
	defer close(stop)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Goroutine 2 ended")
				wg.Done()
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()
	stop <- true

	wg.Add(1)
	// Завершение по контексту(может быть таймаут, сигнал, и проч.)
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 3 ended")
				wg.Done()
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	cancel()

	wg.Add(1)
	dataChan := make(chan int)
	// Завершение по закрытию канала
	go func() {
		for data := range dataChan {
			fmt.Println("Goroutine 4 payload: ", data)
		}
		fmt.Println("Goroutine 4 ended")
		wg.Done()
	}()
	dataChan <- 1
	close(dataChan)

	wg.Add(1)
	// Немеделенное завершение
	go func() {
		defer fmt.Println("Goroutine 4 ended")
		defer wg.Done()
		runtime.Goexit()
	}()

	wg.Wait()
}
