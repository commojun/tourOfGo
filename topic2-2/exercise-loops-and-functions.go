package main

import (
	"fmt"
	"math"
)

// 近似計算を１０回繰り返す
func Sqrt10(x float64) (z float64) {
	z = 1.0
	for n := 0; n < 10; n++ {
		z = (z - (math.Pow(z, 2) - x) / (2 * z))
	}
	return
}

// 収束するまで計算を繰り返す
func Sqrt(x float64) (z float64) {
	z = 1.0
	diff := 1.0
	for diff > 0.000000000001 {
		z_ := float64(z)
		z = (z - (math.Pow(z, 2) - x) / (2 * z))
		diff = math.Abs(z_ - z)
		//fmt.Println(diff)
	}
	return
}

func main() {
	fmt.Println("Sqrt10(2) =  ", Sqrt10(2))
	fmt.Println("Sqrt  (2) =  ",   Sqrt(2))
	fmt.Println("Sqrt10(3) =  ", Sqrt10(3))
	fmt.Println("Sqrt  (3) =  ",   Sqrt(3))
	fmt.Println("Sqrt10(4) =  ", Sqrt10(4))
	fmt.Println("Sqrt  (4) =  ",   Sqrt(4))
	fmt.Println("Sqrt10(5) =  ", Sqrt10(5))
	fmt.Println("Sqrt  (5) =  ",   Sqrt(5))
	fmt.Println("Sqrt10(6) =  ", Sqrt10(6))
	fmt.Println("Sqrt  (6) =  ",   Sqrt(6))

}
