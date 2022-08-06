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
