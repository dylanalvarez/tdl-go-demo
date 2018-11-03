package main

import "fmt"

type Result struct {
    value int
    done  bool
}

type Response struct {
    fibonacci, factorial []int
}

func calculateFibonacci(values chan Result) {
    x, y := 0, 1
    for i := 0; i < 20; i++ {
        x, y = y, x+y
        values <- Result{value: x, done: false}
    }
    values <- Result{done: true}
}

func calculateFactorial(values chan Result) {
    accumulator := 1
    values <- Result{value: 1, done: false}
    for i := 1; i < 10; i++ {
        accumulator *= i
        values <- Result{value: accumulator, done: false}
    }
    values <- Result{done: true}
}

func tryReply(response Response, calculatedCount int) (replied bool) {
    if calculatedCount < 2 {
        return false
    }
		series := map[string] []int {
			"fibonacci": response.fibonacci,
      "factorial": response.factorial,
		}
		fmt.Println(series)
		
    replied = true
    return
}

func main() {
    fibonacciChannel := make(chan Result, 5) // Size of buffer: 5
    factorialChannel := make(chan Result, 5)
    var response Response
    calculatedCount := 0

    go calculateFibonacci(fibonacciChannel)
    go calculateFactorial(factorialChannel)

    for {
        select {
        case result := <-fibonacciChannel:
            if result.done {
                calculatedCount++
                if tryReply(response, calculatedCount) {
                    return
                }
            } else {
                response.fibonacci = append(response.fibonacci, result.value)
            }
        case result := <-factorialChannel:
            if result.done {
                calculatedCount++
                if tryReply(response, calculatedCount) {
                    return
                }
            } else {
                response.factorial = append(response.factorial, result.value)
            }
        }
    }
}
