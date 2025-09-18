package main

import "golang.org/x/tour/tree"
import (
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	var current = t
	stack := make([]*tree.Tree, 0)

	for len(stack) != 0 || current != nil {

		if current != nil {
			stack = append(stack, current)
			current = current.Left
			continue
		}

		parent := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ch <- parent.Value

		current = parent.Right
	}

	close(ch)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	channel1 := make(chan int)
	channel2 := make(chan int)

	go Walk(t1, channel1)
	go Walk(t2, channel2)

	for {
		val1, ok1 := <-channel1
		val2, ok2 := <-channel2

		fmt.Println(val1, val2)

		if val1 != val2 {
			return false
		}

		if !ok1 || !ok2 {
			return true
		}

	}

}

func main() {

	sameTree := Same(tree.New(1), tree.New(1))
	fmt.Println("Same tree", sameTree)

}
