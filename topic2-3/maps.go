package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// こういった省略も可能
var m3 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(m2)
	fmt.Println(m3)

	answer := make(map[string]int)

	answer["Answer"] = 42
	fmt.Println("The value:", answer["Answer"])

	answer["Answer"] = 48
	fmt.Println("The value:", answer["Answer"])

	delete(answer, "Answer")
	fmt.Println("The value:", answer["Answer"])

	v, ok := answer["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}


