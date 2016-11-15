package main

import "fmt"

type Vertex struct {
	X float64
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	q  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)

	//ポインタを通してのアクセス
	p := &v
	p.X = 1e-9  // (*p).X と書くこともできる
	fmt.Println(v)

	fmt.Println(v1, q, v2, v3)

	
}
