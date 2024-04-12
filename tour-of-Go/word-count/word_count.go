package main

import (
	str "strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	dict := make(map[string]int)

	for _, ele := range str.Split(s, " ") {
		dict[ele]++
	}
	return dict
}

func main() {
	wc.Test(WordCount)
}
