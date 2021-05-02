package monitors

import "math"

func calcAvg(states []*State) *State {
	state := State{}
	state.LoadAverage.One = avgFloat64(states, func(state *State) float64 { return state.LoadAverage.One })
	state.LoadAverage.Five = avgFloat64(states, func(state *State) float64 { return state.LoadAverage.Five })
	state.LoadAverage.Fifteen = avgFloat64(states, func(state *State) float64 { return state.LoadAverage.Fifteen })
	state.CPULoad.User = avgFloat64(states, func(state *State) float64 { return state.CPULoad.User })
	state.CPULoad.System = avgFloat64(states, func(state *State) float64 { return state.CPULoad.System })
	state.CPULoad.Idle = avgFloat64(states, func(state *State) float64 { return state.CPULoad.Idle })
	state.Mem.Total = avgFloat64(states, func(state *State) float64 { return state.Mem.Total })
	state.Mem.Used = avgFloat64(states, func(state *State) float64 { return state.Mem.Used })
	state.Mem.Free = avgFloat64(states, func(state *State) float64 { return state.Mem.Free })
	return &state
}

func calcIncrementAvg(states []*State, avg *State) *State {
	state := State{}
	state.LoadAverage.One = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.LoadAverage.One })
	state.LoadAverage.Five = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.LoadAverage.Five })
	state.LoadAverage.Fifteen = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.LoadAverage.Fifteen })
	state.CPULoad.User = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.CPULoad.User })
	state.CPULoad.System = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.CPULoad.System })
	state.CPULoad.Idle = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.CPULoad.Idle })
	state.Mem.Total = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.Mem.Total })
	state.Mem.Used = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.Mem.Used })
	state.Mem.Free = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.Mem.Free })
	return &state
}

func avgIncrementFloat64(states []*State, avgState *State, fn func(state *State) float64) float64 {
	avg := fn(avgState) - fn(states[0])/float64(len(states)) + fn(states[len(states)-1])/float64(len(states))
	return math.Round(avg*1000) / 1000
}

func avgFloat64(states []*State, fn func(state *State) float64) float64 {
	var avg float64
	for _, state := range states {
		avg += fn(state)
	}
	avg /= float64(len(states))
	return math.Round(avg*1000) / 1000
}
