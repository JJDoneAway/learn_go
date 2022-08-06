package main

import (
	"fmt"
	"learn_go/arrays"
	"learn_go/hello_world"
)

func main() {
	fmt.Println(hello_world.Hello("Johannes", ""))
	fmt.Println(arrays.Sum([5]int{1, 2, 3, 4, 5}))
}
