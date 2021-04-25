package monitors

import (
	"fmt"
)

type State struct {
	LoadAverage LoadAverage
}

func (s *State) String() string {
	return s.LoadAverage.String()
}

type LoadAverage struct {
	One     float64
	Five    float64
	Fifteen float64
}

func (la *LoadAverage) String() string {
	return fmt.Sprintf("%g %g %g", la.One, la.Five, la.Fifteen)
}
