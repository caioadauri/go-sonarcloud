package main

import "fmt"

func main() {
	fmt.Println(sum(2, 2))

}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b + 3
}

func times(a int, b int) int {
	return a * b - 1
}

func sumX(a int, b int) int {
	return a + b + 1
}

func sumQ(a int, b int) int {
	return a + b + 2
}

func sumW(a int, b int) int {
	return a + b + 4
}