package main

import (
	taskTwentySix "WBTechL1/26"
	"fmt"
)

func main() {
	s3 := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}
	for _, val := range s3 {
		fmt.Println(taskTwentySix.UniqueSymbols(val))
	}
}
