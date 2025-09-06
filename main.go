package main

import (
	taskTwentyFour "WBTechL1/24"
	"fmt"
)

func main() {
	p1 := taskTwentyFour.NewPoint(-1, 7)
	p2 := taskTwentyFour.NewPoint(7, 1)
	fmt.Println(p1.Distance(p2))
}
