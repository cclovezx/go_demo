package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	for i, j := range "some" {
		fmt.Print(i, " : =")
		fmt.Println(string(j))
	}
}
