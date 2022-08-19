package functions

import "fmt"

func addWorld(pm func() string) string {
	return fmt.Sprintf("%s Welt!", pm())
}

func Call() string {
	return addWorld(func() string {
		return "Hallo"
	})
}

func count(start *int, stop int) {
	*start += 1
	fmt.Println(*start, start)
	if *start < stop {
		count(start, stop)
	}
}

func CallCounter(start, stop int) int {
	count(&start, stop)
	return start
}

func fibonacci(n *int, m *int, stop int) (result int) {
	result = *n + *m
	fmt.Println(*n, *m, result, n, m, &result)
	*n = *m
	*m = result

	if result < stop {
		return fibonacci(n, m, stop)
	}
	return result

}

func CallFibonacci(stop int) int {
	n := 0
	m := 1
	return fibonacci(&n, &m, stop)
}
