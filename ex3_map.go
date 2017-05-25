package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	strs := strings.Fields(s)
	for _, str := range strs {
		ret[str]++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}

