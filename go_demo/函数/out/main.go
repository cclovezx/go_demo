package main

import "fmt"

func main() {
	fmt.Print(min(1, -5))
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return b
}
