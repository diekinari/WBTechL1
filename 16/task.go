package taskSixteen

func QuickSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	pivot := slice[len(slice)/2]
	var left []int
	var middle []int
	var right []int
	for _, x := range slice {
		if x < pivot {
			left = append(left, x)
		} else if x > pivot {
			right = append(right, x)
		} else {
			middle = append(middle, x)
		}
	}
	return append(append(QuickSort(left), middle...), QuickSort(right)...)
}
