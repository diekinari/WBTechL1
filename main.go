package main

import (
	taskTwentyTwo "WBTechL1/22"
	"fmt"
	"math"
	"math/big"
)

func main() {
	a := big.NewInt(int64(math.Pow(2, 20)))
	b := big.NewInt(int64(math.Pow(2, 25)))
	res := taskTwentyTwo.MathBigNumbers(a, b, "*")
	fmt.Println(res)
}
