package main

import (
	. "example/response"
	. "example/result"
	"fmt"
)

func calculateFibonacci(values chan Result) {
    x, y := 0, 1
    for i := 0; i < 20; i++ {
        x, y = y, x+y
        values <- Result{Value: x, Done: false}
    }
    values <- Result{Done: true}
}

func calculateFactorial(values chan Result) {
    accumulator := 1
    values <- Result{Value: 1, Done: false}
    for i := 1; i < 10; i++ {
        accumulator *= i
        values <- Result{Value: accumulator, Done: false}
    }
    values <- Result{Done: true}
}

func tryReply(response Response) (replied bool) {
	if !response.IsDone() {
		replied = false	
		return
	}
	fmt.Println(response.Print())
	replied = true
	return
}

func main() {
    fibonacciChannel := make(chan Result, 1) // Size of buffer: 5
    factorialChannel := make(chan Result, 1)
	var response Response
	response.Init()

    go calculateFibonacci(fibonacciChannel)
    go calculateFactorial(factorialChannel)

    for {
        select {
			case result := <-fibonacciChannel:
				response.Fibonacci.Append(result)
			case result := <-factorialChannel:
				response.Factorial.Append(result)
			}
		if tryReply(response) {
			return
		}
    }
}
