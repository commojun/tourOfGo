package main

import "fmt"

// 変数宣言を一度にできる（最後が型名）
var c, python, java bool

// 中身も同時に宣言するときはこんな感じ
var k, j int = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)

    var	c, python, javascript = true, false, "no!"
	fmt.Println(k, j, c, python, javascript)

	//関数の中でのみかんたんな宣言が可能 [:=]
	l := 3
	hoge, huga := "foo", "bar"
	fmt.Println(l, hoge, huga)
}
