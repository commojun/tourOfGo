package main

import (
	"fmt"
	"time"
)

func main() {
	// こいつらもチャネルなのか。
	// func After(d Duration) <-chan Time
	// func Tick(d Duration) <-chan Time
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
