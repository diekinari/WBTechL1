package taskSeven

import (
	"fmt"
	"sync"
	"time"
)

func WriteToMap() {
	mu := sync.RWMutex{}
	m := make(map[int]int)
	m[1] = 0
	// Без data race последняя горутина запишет финальное значение.
	for i := 10; i < 13; i++ {
		go func(num int) {
			mu.Lock()
			m[1] = num
			fmt.Println(m[1])
			mu.Unlock()
			fmt.Println(num, "was written")
		}(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(m[1])
}
