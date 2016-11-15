package main

import (
	"./tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
//func Walk(t *tree.Tree, ch chan int)
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	t := tree.New(4)
	fmt.Println("---Walking test---")
	fmt.Println("tree: ", t)
	ch := make(chan int, 10)
	fmt.Print("Walking: ")
	go Walk(t, ch)
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println("\n")

	fmt.Println("---Same test---")
	t1 := tree.New(2)
	t2 := tree.New(2)
	fmt.Println("t1: ", t1)
	fmt.Println("t2: ", t2)
	fmt.Println("Same(t1, t2): ", Same(t1, t2))

	fmt.Println("---Same test---")
	t1 = tree.New(5)
	t2 = tree.New(6)
	fmt.Println("t1: ", t1)
	fmt.Println("t2: ", t2)
	fmt.Println("Same(t1, t2): ", Same(t1, t2))
}
