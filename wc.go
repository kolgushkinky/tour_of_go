package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	words := strings.Fields(s)
	m := make(map[string]int)

	for _, w := range words {
		m[string(w)] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
