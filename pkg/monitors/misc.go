package monitors

import "math"

func calcAvg(states []*State, metrics *MetricsPresent) *State {
	state := State{}
	if metrics.la {
		state.LoadAverage = &LoadAverage{}
		state.LoadAverage.One = avgFloat64(states, func(state *State) float64 { return state.LoadAverage.One })
		state.LoadAverage.Five = avgFloat64(states, func(state *State) float64 { return state.LoadAverage.Five })
		state.LoadAverage.Fifteen = avgFloat64(states, func(state *State) float64 { return state.LoadAverage.Fifteen })
	}
	if metrics.cpu {
		state.CPULoad = &CPULoad{}
		state.CPULoad.User = avgFloat64(states, func(state *State) float64 { return state.CPULoad.User })
		state.CPULoad.System = avgFloat64(states, func(state *State) float64 { return state.CPULoad.System })
		state.CPULoad.Idle = avgFloat64(states, func(state *State) float64 { return state.CPULoad.Idle })
	}
	if metrics.mem {
		state.Mem = &Mem{}
		state.Mem.Total = avgFloat64(states, func(state *State) float64 { return state.Mem.Total })
		state.Mem.Used = avgFloat64(states, func(state *State) float64 { return state.Mem.Used })
		state.Mem.Free = avgFloat64(states, func(state *State) float64 { return state.Mem.Free })
	}
	return &state
}

func calcIncrementAvg(states []*State, avg *State, metrics *MetricsPresent) *State {
	state := State{}
	if metrics.la {
		state.LoadAverage = &LoadAverage{}
		state.LoadAverage.One = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.LoadAverage.One })
		state.LoadAverage.Five = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.LoadAverage.Five })
		state.LoadAverage.Fifteen = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.LoadAverage.Fifteen })
	}
	if metrics.cpu {
		state.CPULoad = &CPULoad{}
		state.CPULoad.User = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.CPULoad.User })
		state.CPULoad.System = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.CPULoad.System })
		state.CPULoad.Idle = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.CPULoad.Idle })
	}
	if metrics.mem {
		state.Mem = &Mem{}
		state.Mem.Total = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.Mem.Total })
		state.Mem.Used = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.Mem.Used })
		state.Mem.Free = avgIncrementFloat64(states, avg, func(state *State) float64 { return state.Mem.Free })
	}
	return &state
}

func avgIncrementFloat64(states []*State, avgState *State, fn func(state *State) float64) float64 {
	avg := fn(avgState) - fn(states[0])/float64(len(states)) + fn(states[len(states)-1])/float64(len(states))
	return math.Round(avg*100) / 100
}

func avgFloat64(states []*State, fn func(state *State) float64) float64 {
	var avg float64
	for _, state := range states {
		avg += fn(state)
	}
	avg /= float64(len(states))
	return math.Round(avg*100) / 100
}

func findString(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func InitMetricPresent(suppressed []string) *MetricsPresent {
	metrics := MetricsPresent{}
	if !findString(suppressed, "la") {
		metrics.la = true
	}
	if !findString(suppressed, "mem") {
		metrics.mem = true
	}
	if !findString(suppressed, "cpu") {
		metrics.cpu = true
	}
	return &metrics
}
