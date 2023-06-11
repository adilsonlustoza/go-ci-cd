package main

import "fmt"

func main() {
	fmt.Println(soma(10, 10))

	fmt.Println(sub(20, 10))
}

func soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
