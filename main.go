package main

import (
	taskTwentyFive "WBTechL1/25"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	taskTwentyFive.Sleep(2 * time.Second)
	elapsed := time.Since(start)

	// Проверяем, что прошло примерно 2 секунды (допускаем небольшую погрешность)
	if elapsed < 2*time.Second || elapsed > 2100*time.Millisecond {
		fmt.Printf("Ожидалась задержка ~2 секунды, получили: %v", elapsed)
	} else {
		fmt.Printf("Прошло ровно 2 секунды\n")
	}

	// Тест 2: Проверяем работу с меньшим интервалом
	start = time.Now()
	taskTwentyFive.Sleep(100 * time.Millisecond)
	elapsed = time.Since(start)

	if elapsed < 100*time.Millisecond || elapsed > 150*time.Millisecond {
		fmt.Printf("Ожидалась задержка ~100ms, получили: %v", elapsed)
	} else {
		fmt.Printf("Прошло ровно 100 мс")
	}

}
