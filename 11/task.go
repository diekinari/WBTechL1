package taskEleven

// Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов)
// — т.е. вывести элементы, присутствующие и в первом, и во втором.

// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}

func FindIntersection(a []int, b []int) []int {
	mA := map[int]interface{}{}

	for _, val := range a {
		mA[val] = struct{}{}
	}
	result := []int{}
	for _, val := range b {
		if _, exists := mA[val]; exists {
			result = append(result, val)
		}
	}

	return result

}
