package main

import (
	"fmt"
)

func main() {
	source := "Hello Bing"
	target := "o Bin"

	ans := Index(source, target)
	fmt.Print(ans) // we expect the answer is 4
}

func Index(source, target string) int {
	// please implement your code here
	ans := -1
	for i := 0; i < len(source); i++ {
		s1 := source[i]
		t1 := target[0]

		if s1 == t1 {
			match := false
			for j := 1; j < len(target); j++ {
				s := source[i+j]
				t := target[j]
				if s != t {
					match = false
					break
				} else {
					match = true
				}
			}
			if match {
				return i
			}
		}
	}
	return ans
}
