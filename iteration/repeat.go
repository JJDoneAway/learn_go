package iterations

import (
	"bytes"
	"strings"
)

func Repeat(char string, length int) string {
	var b bytes.Buffer
	for i := 0; i < length; i++ {
		b.WriteString(char)
	}
	return b.String()
}

func Repeat2(char string, length int) string {

	return strings.Repeat(char, length)
}

func SimpleLoop() (i int) {
	for i < 10 {
		println(i)
		i++
	}
	println("My i is ", i)
	return i
}

func NormalLoop() int {
	count := 0

	for i := 0; i < 10; i++ {
		print(i)
		count = i
	}

	return count

}

func InfiniteLoop() int {
	count := 0

	for {
		count++
		if count > 10 {
			break
		}
	}
	return count
}

func CollectionLoop(values []int) int {
	ret := 0
	for i, v := range values {
		println(i, v)
		ret = v
	}

	return ret
}

func MapLoop(values map[string]int) (string, int) {
	var key string
	var value int
	for k, v := range values {
		println(k, v)
		key = k
		value = v
	}

	return key, value
}

func IfCondition() bool {
	i := 10

	if i == 10 {
		println("Is 10")
	}
	if i%2 == 0 {
		println("Is even")
	} else {
		println("Strange")
	}

	return true
}
