package monitors

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type StateCollector interface {
	GetCurrentState(ctx context.Context) (*State, error)
	GetLoadAverage(ctx context.Context) (LoadAverage, error)
}

type OsMonitor struct {
	log            *logrus.Entry
	stateCollector StateCollector
	states         []*State
	averages       map[int]*State
	maxM           int
	mxAvg          sync.RWMutex
	mxMaxM         sync.RWMutex
	mxStates       sync.RWMutex
}

func GetOsMonitor(log *logrus.Logger, goos string) (*OsMonitor, error) {
	var stateContainer StateCollector
	switch goos {
	case "linux":
		stateContainer = &LinuxStateCollector{}
	case "darwin":
	case "windows":
	default:
		return nil, fmt.Errorf("unsupported os: %s, only linux darwin and windows are supported", goos)
	}
	tmp := make(map[int]*State)
	tmp[2] = nil
	tmp[5] = nil
	tmp[11] = nil
	monitor := OsMonitor{
		log:            log.WithField("system", "monitor"),
		stateCollector: stateContainer,
		maxM:           12,
		//averages:       make(map[int]*State),
		averages:       tmp,
	}
	return &monitor, nil
}

func (om *OsMonitor) Run(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		om.startCollector(ctx)
	}()
	go func() {
		defer wg.Done()
		om.startCalculator(ctx)
	}()
	wg.Wait()
}

func (om *OsMonitor) AddMAverage(m int) {
	om.mxMaxM.Lock()
	if m > om.maxM {
		om.maxM = m
	}
	om.mxMaxM.Unlock()
}

func (om *OsMonitor) startCollector(ctx context.Context) {
	scheduler := time.NewTicker(time.Second)
	defer scheduler.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-scheduler.C:
			state, err := om.stateCollector.GetCurrentState(ctx)
			if err != nil {
				om.log.Warn("err receiving stats: ", err)
				continue
			}
			om.mxStates.Lock()
			om.mxMaxM.RLock()
			if len(om.states) < om.maxM {
				om.states = append(om.states, state)
			} else {
				om.states = om.states[1:]
				om.states = append(om.states, state)
			}
			om.mxStates.Unlock()
			om.mxMaxM.RUnlock()
			om.mxStates.RLock()

			om.log.Debug("_____________________________________")
			for _, state := range om.states {
				om.log.Debugf(state.String())
			}
			om.log.Debug("len(states) == ", len(om.states))
			om.log.Debug("_____________________________________")
			om.mxStates.RUnlock()
		}
	}
}

func (om *OsMonitor) startCalculator(ctx context.Context) {
	scheduler := time.NewTicker(time.Second)
	defer scheduler.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-scheduler.C:
			om.mxAvg.Lock()
			for k, _ := range om.averages {
				switch {
				case k < len(om.states):
					continue
				case k == len(om.states):
					om.averages[k] = calcAvg(om.states)
				default:
					om.averages[k] = calcAvg(om.states)
				}
			}
			om.mxAvg.Unlock()

			om.mxAvg.RLock()
			om.log.Debug("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			for k, average := range om.averages {
				om.log.Debugf("%d: %s", k, average.String())
			}
			om.log.Debug("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			om.mxAvg.RUnlock()
		}
	}
}

func calcAvg(states []*State) *State {
	state := State{}
	state.LoadAverage.One = avgFloat64(states, func(state *State) float64 { return state.LoadAverage.One })
	state.LoadAverage.Five = avgFloat64(states, func(state *State) float64 { return state.LoadAverage.Five })
	state.LoadAverage.Fifteen = avgFloat64(states, func(state *State) float64 { return state.LoadAverage.Fifteen })
	return &state
}

func avgFloat64(states []*State, fn func(state *State) float64) float64 {
	var avg float64
	for _, state := range states {
		avg += fn(state)
	}
	avg /= float64(len(states))
	return avg
}
