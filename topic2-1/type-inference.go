package main

import "fmt"

func main() {
	// 空気を読んで型を設定してくれる話
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
	x := 42.195 // change me!
	fmt.Printf("x is of type %T\n", x)
	y := 123456789087 // change me!
	fmt.Printf("y is of type %T\n", y)
	z := 0.777 + 736i // change me!
	fmt.Printf("z is of type %T\n", z)
}
