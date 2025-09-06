package taskTwentyTwo

import "math/big"

// MathBigNumbers выполняет арифметическую операцию над двумя числами
func MathBigNumbers(a, b *big.Int, operation string) *big.Int {
	res := &big.Int{}
	switch operation {
	case "+":
		return res.Add(a, b)
	case "-":
		return res.Sub(a, b)
	case "*":
		return res.Mul(a, b)
	case "/":
		return res.Div(a, b)
	default:
		return res
	}
}
