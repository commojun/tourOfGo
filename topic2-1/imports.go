package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println ("円周率表示")
	fmt.Println (math.Pi)

	fmt.Println("小文字で始まる名前は外部からアクセスできない。")
	//fmt.Println(math.pi)
}
