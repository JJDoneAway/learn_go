package arrays

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(numbers ...[]int) (result []int) {
	elements := len(numbers)

	for i := 0; i < elements; i++ {
		result = append(result, Sum(numbers[i]))
	}

	return result
}

func SumAllTails(numbers ...[]int) (result []int) {
	elements := len(numbers)

	for i := 0; i < elements; i++ {
		if len(numbers[i]) > 1 {
			result = append(result, Sum(numbers[i][1:]))
		} else {
			result = append(result, 0)
		}
	}

	return result
}
