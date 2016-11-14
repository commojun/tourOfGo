package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// ifの条件文のところにかんたんなステートメントを入れることが
	// 可能！（ためしに計算してみる、とかが可能なんだろう）
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// 一連の条件文の中では、vが使える。
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// ここでvは使えない
	return lim
}


func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
