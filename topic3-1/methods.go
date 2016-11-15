package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// これだと、メソッド呼び出しのたびに変数がコピーされる
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 構造体の中身の変数を変更するには、
// ポインタレシーバでメソッドを宣言しなければならない
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(3)
	fmt.Println(v.Abs())

	p :=  &Vertex{4, 3}
	//このような書き方でもOK
	p.Scale(5)
	//これもOK
	fmt.Println(p.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

