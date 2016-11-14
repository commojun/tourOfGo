package main

import "fmt"

// 戻り値の方、個数、名前が宣言の部分で全部できちゃう
func split(sum int)(x, y int){
	x = sum * 4 / 9
	y = sum - x
	return  // naked returnというらしい。
}

func main() {
	fmt.Println(split(17))
}
