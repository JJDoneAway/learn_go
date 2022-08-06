package arrays

import "fmt"

func Sum(numbers [5]int) (result int) {
	for index, number := range numbers {
		fmt.Println(index)
		result += number
	}
	return
}
