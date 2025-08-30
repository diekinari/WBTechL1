package taskTwo

import (
	"fmt"
	"sync"
)

// ConcurrentSquare – Написать программу, которая конкурентно рассчитает значения квадратов чисел,
// взятых из массива [2,4,6,8,10], и выведет результаты в stdout.
func ConcurrentSquare() {
	nums := [...]int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	wg.Add(len(nums))
	for _, num := range nums {
		go func(num int) {
			fmt.Println(num * num)
			wg.Done()
		}(num)
	}
	wg.Wait()
}
