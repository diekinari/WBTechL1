package taskFifteen

import (
	"fmt"
	"strings"
)

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

// someFunc is a correct function.
// Вместо глобальной переменной – возвращаем значение из функции
// Проверяем на ошибку
// Используем strings.Builder
func someFunc() (string, error) {
	v := createHugeString(1 << 10)
	if len(v) < 100 {
		return "", fmt.Errorf("string too short: expected 100, got %d", len(v))
	}
	var builder strings.Builder
	builder.WriteString(v[:100])
	return builder.String(), nil
}
