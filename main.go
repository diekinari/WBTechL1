package main

import (
	taskSixteen "WBTechL1/16"
	taskSeventeen "WBTechL1/17"
	"fmt"
)

func main() {
	slice := taskSixteen.QuickSort([]int{3, 7, 8, 5, 2, 1, 9, 5, 4})
	result := taskSeventeen.BinarySearch(slice, 8, 0, len(slice)-1)
	fmt.Println(result)
}
