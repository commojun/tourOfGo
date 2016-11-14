package main

import "fmt"

func main() {
	sum := 0
	for i :=0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 初期化と後処理のステートメントはなくてもOK
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	// セミコロンも省略可能！whileと同じように使える
	sum = 1
	for  sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// こうすると無限ループ
	//for {
	//}
}
