package main

import (
	taskTwentyThree "WBTechL1/23"
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4}
	newS := taskTwentyThree.DeleteItemFromSlice(s, 3)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("newS: %v \n", newS)
}
