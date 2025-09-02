package taskTwelve

// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func MakeSet(seq []string) []string {
	set := make(map[string]interface{})
	var result []string
	for _, val := range seq {
		set[val] = struct{}{}
	}

	for key, _ := range set {
		result = append(result, key)
	}
	return result
}
