package main

import "fmt"

func calculateFibonacci(n int) {
	var x int
	y := 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Println(x)
	}
}

func calculateFactorials(n int) {
	accumulator := 1
	fmt.Println(accumulator)
	for i := 1; i < n; i++ {
		accumulator *= i
		fmt.Println(accumulator)
	}
}

func main() {
	calculateFactorials(10)
	calculateFibonacci(10)
}
