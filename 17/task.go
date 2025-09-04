package taskSeventeen

// BinarySearch реализует алгоритм бинарного поиска встроенными методами языка.
// Функция принимает отсортированный слайс и искомый элемент,
// возвращает индекс элемента или -1, если элемент не найден.
func BinarySearch(slice []int, value int, first int, last int) int {
	var middle int
	if last >= first {
		middle = first + (last-first)/2
		if slice[middle] == value {
			return middle
		} else if slice[middle] < value {
			return BinarySearch(slice, value, middle+1, last)
		} else {
			return BinarySearch(slice, value, first, middle-1)
		}
	}
	return -1
}
