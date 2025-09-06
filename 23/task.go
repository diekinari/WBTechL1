package taskTwentyThree

func DeleteItemFromSlice(srcSlice []int, i int) []int {
	slice := make([]int, len(srcSlice))
	copy(slice, srcSlice)
	if i < len(slice)-1 {
		copy(slice[i:], slice[i+1:])
	}
	// для других типов здесь стоило бы обнулять последний элемент, чтобы предотвратить утечку памяти
	slice[len(slice)-1] = 0
	slice = slice[:len(slice)-1]
	return slice
}
