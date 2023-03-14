package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
	if t.Left != nil {
		walk(t.Left, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func toSlice(ch chan int) []int {
	var a []int
	for a_i := range ch {
		a = append(a, a_i)
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	return a
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch_a, ch_b := make(chan int), make(chan int)
	go Walk(t1, ch_a)
	go Walk(t2, ch_b)

	a := toSlice(ch_a)
	b := toSlice(ch_b)

	if len(a) != len(b) {
		return false
	}

	for pos, v := range a {
		if b[pos] != v {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
