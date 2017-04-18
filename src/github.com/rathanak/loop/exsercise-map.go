package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, value := range strings.Fields(s) {
		v, ok := m[value]
		if ok == true {
			 m[value] =strings v + 1
		 }
		 else {
			  m[value] = 1
			}
		}
		return m
	}

	func main() {
		wc.okTest(WordCount)
	}

