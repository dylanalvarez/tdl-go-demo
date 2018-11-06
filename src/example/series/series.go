package series

import . "example/result"

type Series struct {
	sequence []Result
	Done bool
}

func (s *Series) Append(elem Result) () {
	if elem.Done {
		s.Done = true
		return
	}
	s.sequence = append(s.sequence, elem)
	return
}

func (s Series) GetSeries() (seriesValues []int) {
	for _, aResult := range s.sequence {
		seriesValues = append(seriesValues, aResult.Value)
	}
	return
}

func (s *Series) Init() () {
	s.Done = false
	return
}
