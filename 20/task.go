package taskTwenty

import "strings"

// RevertSentence принимает на вход строку и возвращает
// ее с обратным порядком слов. Тот же алгоритм, что и в 19 задании
func RevertSentence(s string) string {
	buf := strings.Split(s, " ")
	for i := len(buf)/2 - 1; i >= 0; i-- {
		opp := len(buf) - 1 - i
		buf[i], buf[opp] = buf[opp], buf[i]
	}
	return strings.Join(buf, " ")
}
