package main

import (
	"fmt"
	"strings"
)

func main() {
	names := [4]string{
		"John",
		"Paul",
		"Geoge",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// スライスは配列へのリファレンスなので、
	// スライス側を変更すると、もとの方も変更されてしまう。
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var sl []int = primes[1:4]
	fmt.Println(sl)

	//配列の個数を言わないとスライスになる
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// どう解釈すればいいんだ？
	h := primes[1:4]
	printSlice(h)

	// スライス位置を省略すると最大・最小になる
	j := primes[:2]
	printSlice(j)

	// 手前は切れる、後ろは広げられる
	g := j[1:4]
	printSlice(g)
	
	k := primes[3:]
	printSlice(k)

	// 空スライス nil
	var y []int
	printSlice(y)
	if y == nil {
		fmt.Println("nil!")
	}

	// making slice
	c := make([]int, 5)
	printSlice2("c",c)
	d := make([]int, 0, 5)
	printSlice2("d",d)
	e := c[:2]
	printSlice2("e",e)
	f := d[2:5]
	printSlice2("f",f)

	// slices of slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// スライスへの追加
	var o []int
	printSlice(o)

	o = append(o, 0)
	printSlice(o)

	o = append(o, 1)
	printSlice(o)

	o = append(o, 2, 3, 4)
	printSlice(o)
}


func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
