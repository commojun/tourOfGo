package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a_next1 := 0
	a_next2 := 1
	return func() int {
		a := a_next1
		a_next1, a_next2 = a_next2, (a_next1 + a_next2)
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
