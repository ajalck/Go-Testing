package main

import "fmt"

func Calculate(x int) int {
	return x * x
}
func main() {
	result := Calculate(2)
	fmt.Printf("Result is :%v", result)
}
