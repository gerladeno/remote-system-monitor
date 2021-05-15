package monitors

import (
	"fmt"
	"reflect"
)

type State struct {
	LoadAverage *LoadAverage
	CPULoad     *CPULoad
	Mem         *Mem
}

func (s *State) String() string {
	return fmt.Sprintf("%s %s %s", SafeString(s.LoadAverage), SafeString(s.CPULoad), SafeString(s.Mem))
}

type LoadAverage struct {
	One     float64
	Five    float64
	Fifteen float64
}

func (la *LoadAverage) String() string {
	return fmt.Sprintf("LoadAvg [One: %g Five: %g Fifteen: %g]", la.One, la.Five, la.Fifteen)
}

type CPULoad struct {
	User   float64
	System float64
	Idle   float64
}

func (cpu *CPULoad) String() string {
	return fmt.Sprintf("CPU [User: %g%%%% System: %g%%%% Idle: %g%%%%]", cpu.User, cpu.System, cpu.Idle)
}

type Mem struct {
	Total float64
	Free  float64
	Used  float64
}

func (m *Mem) String() string {
	return fmt.Sprintf("Mem [Total MB: %g Free MB: %g Used MB: %g]", m.Total, m.Free, m.Used)
}

func SafeString(stringer fmt.Stringer) string {
	if stringer == nil || (reflect.ValueOf(stringer).Kind() == reflect.Ptr && reflect.ValueOf(stringer).IsNil()) {
		return ""
	}
	return stringer.String()
}
