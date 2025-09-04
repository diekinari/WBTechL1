package taskEighteen

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type incrementor struct {
	count int32
}

func ConcurrentIncrement() {
	incr := &incrementor{0}
	wg := sync.WaitGroup{}
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func(obj *incrementor) {
			defer wg.Done()
			for j := 1; j < 6; j++ {
				time.Sleep(time.Millisecond * 2) // для более явного data race
				atomic.AddInt32(&obj.count, 1)

			}
		}(incr)
	}
	wg.Wait()
	// expecting 250
	fmt.Printf("final count: %d\n", incr.count)
}
