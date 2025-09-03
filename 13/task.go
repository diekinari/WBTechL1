package taskThirteen

// ChangePositions changes positions of two variables without using additional memory
func ChangePositions(a int, b int) (int, int) {
	// a = a + b; b = b
	a = a + b
	// a = a + b; b = (a+b) - b = a
	b = a - b
	// a = (a+b) - a; b = a
	a = a - b
	return a, b
}
