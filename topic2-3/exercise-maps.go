package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (m map[string]int) {
	fields := strings.Fields(s)
	m = make(map[string]int)
	for _, element := range fields {
		m[element]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
