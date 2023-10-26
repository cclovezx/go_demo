package main

import "fmt"

func main() {
	min := func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	fmt.Print(min(1, 3))
}
