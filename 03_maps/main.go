package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var resultMap map[string]int = make(map[string]int)

	for _, v := range strings.Fields(s) {
		count, isPresent := resultMap[v];
		if isPresent {
			resultMap[v] = count + 1;
		} else {
			resultMap[v] = 1;
		}
	}

	return resultMap
}

func main() {
	wc.Test(WordCount)
}
