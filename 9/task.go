package taskNine

import (
	"fmt"
	"sync"
)

func StartConveyer() {
	nums := []int{2, 4, 6, 8, 10}
	firstCh := make(chan int)
	secCh := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func(ch chan<- int) {
		for _, val := range nums {
			firstCh <- val
		}
		close(firstCh)
		wg.Done()
	}(firstCh)

	go func(ch <-chan int) {
		for val := range ch {
			secCh <- val * 2
		}
		close(secCh)
		wg.Done()
	}(firstCh)

	go func(ch <-chan int) {
		for val := range ch {
			fmt.Println(val)
		}
		wg.Done()
	}(secCh)

	wg.Wait()
}
