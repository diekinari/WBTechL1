package taskTen

// Дана последовательность температурных колебаний. Объединить эти значения в группы с шагом 10 градусов.
//
// Пример: -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20:{24.5}, 30:{32.5}.
//
// Пояснение: диапазон -20 включает значения от -20 до -29.9, диапазон 10 – от 10 до 19.9, и т.д.
// Порядок в подмножествах не важен.

var tens = []int{-50, -40, -30, -20, -10, 0, 10, 20, 30, 40, 50}

func GroupTemperature() map[int][]float64 {
	tmps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	result := map[int][]float64{}
	for _, currNum := range tmps {
		for i := 0; i < len(tens)-1; i++ {
			left := tens[i]
			right := tens[i+1]
			if currNum >= float64(left) && currNum < float64(right) {
				if currNum < 0 {
					result[right] = append(result[right], currNum)
				} else {
					result[left] = append(result[left], currNum)
				}
				break
			}
		}
	}
	return result

}
