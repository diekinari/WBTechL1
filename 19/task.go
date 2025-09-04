package taskNineteen

func RevertString(s string) string {
	buf := []rune(s)
	for i := len(buf)/2 - 1; i >= 0; i-- {
		opp := len(buf) - 1 - i
		buf[i], buf[opp] = buf[opp], buf[i]
	}
	return string(buf)
}
