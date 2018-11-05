package handlers

import "github.com/gin-gonic/gin"

type Result struct {
    value int
    done  bool
}

type Series struct {
	sequence []Result
	done bool
}

type Response struct {
    fibonacci, factorial Series
}

func (s *Series) append(elem Result) () {
	if (elem.done) {
		s.done = true
		return
	}
	s.sequence = append(s.sequence, elem)
	return
}

func (s Series) getSeries() (seriesValues []int) {
	for _, aResult := range s.sequence {
		seriesValues = append(seriesValues, aResult.value)
	}
	return
}

func (s *Series) init() () {
	s.done = false
	return
}

func (response *Response) init() () {
	response.factorial.init()
	response.fibonacci.init()
	return
}

func (response Response) print() (series map[string] []int) {
	series = map[string] []int {
		"fibonacci": response.fibonacci.getSeries(),
		"factorial": response.factorial.getSeries(),
	}
	return
}

func (response Response) isDone() (replied bool) {
	replied = response.factorial.done && response.fibonacci.done
  return
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

func tryReply(context *gin.Context, response Response) (replied bool) {
	if !response.isDone() {
		replied = false	
		return
	}
	context.JSON(200, gin.H{
		"fibonacci": response.fibonacci.getSeries(),
		"factorial": response.factorial.getSeries(),
	})
	replied = true
	return
}

func SeriesHandler(context *gin.Context) {
    fibonacciChannel := make(chan Result, 5) // Size of buffer: 5
    factorialChannel := make(chan Result, 5)
		var response Response
		response.init()

    go calculateFibonacci(fibonacciChannel)
    go calculateFactorial(factorialChannel)

    for {
        select {
				case result := <-fibonacciChannel:
					response.fibonacci.append(result)
				case result := <-factorialChannel:
					response.factorial.append(result)
        }
		if tryReply(context, response) {
			return
		}
	}
}