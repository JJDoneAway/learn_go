package main

import (
	"fmt"
	"learn_go/arrays"
	"learn_go/hello_world"
)

func main() {
	fmt.Println(hello_world.Hello("Johannes", ""))
	fmt.Println(arrays.Sum([]int{1, 2, 3, 4, 5}))
}
