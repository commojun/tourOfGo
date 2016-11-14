package main

import "fmt"

func main() {
	defer fmt.Println("world")

	// defer は「スタック」の挙動をする
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}


	fmt.Println("hello")
}
