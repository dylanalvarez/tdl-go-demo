package response

import . "example/series"

type Response struct {
	Fibonacci, Factorial Series
}

func (response *Response) Init() () {
	response.Factorial.Init()
	response.Fibonacci.Init()
return
}

func (response Response) Print() (series map[string] []int) {
series = map[string] []int {
	"Fibonacci": response.Fibonacci.GetSeries(),
	"Factorial": response.Factorial.GetSeries(),
}
return
}

func (response Response) IsDone() (replied bool) {
replied = response.Factorial.Done && response.Fibonacci.Done
return
}
