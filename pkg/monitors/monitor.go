package monitors

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

const initialWindowLengthSeconds = 5

var ErrNotCalculated = errors.New("err this average is not yet calculated")

type StateCollector interface {
	GetCurrentState(ctx context.Context) *State
	//GetLoadAverage(ctx context.Context) (LoadAverage, error)
	//GetCPULoadAndMem(ctx context.Context) (CPULoad, Mem, error)
}

type OsMonitor struct {
	log            *logrus.Entry
	stateCollector StateCollector
	states         []*State
	averages       map[int]*State
	avgRequired    map[int]int
	maxM           int
	mxAR           sync.RWMutex
	mxAvg          sync.RWMutex
	mxMaxM         sync.RWMutex
	mxStates       sync.RWMutex
}

func GetOsMonitor(log *logrus.Logger, goos string) (*OsMonitor, error) {
	var stateContainer StateCollector
	switch goos {
	case "linux":
		stateContainer = &LinuxStateCollector{log: log}
	case "darwin":
		stateContainer = &DarwinStateCollector{log: log}
	case "windows":
		stateContainer = &WindowsStateCollector{log: log}
	default:
		return nil, fmt.Errorf("unsupported os: %s, only linux darwin and windows are supported", goos)
	}
	monitor := OsMonitor{
		log:            log.WithField("system", "monitor"),
		stateCollector: stateContainer,
		maxM:           initialWindowLengthSeconds,
		averages:       make(map[int]*State),
		avgRequired:    make(map[int]int),
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
	om.mxAR.Lock()
	_, ok := om.avgRequired[m]
	if ok {
		om.avgRequired[m] += 1
	} else {
		om.avgRequired[m] = 1
	}
	om.mxAR.Unlock()

	om.mxMaxM.Lock()
	if m > om.maxM {
		om.maxM = m
	}
	om.mxMaxM.Unlock()
}

func (om *OsMonitor) RemoveMAverage(m int) {
	om.mxAR.Lock()
	_, ok := om.avgRequired[m]
	if ok {
		if om.avgRequired[m] > 1 {
			om.avgRequired[m] -= 1
		} else {
			delete(om.avgRequired, m)
			om.mxAvg.Lock()
			delete(om.averages, m)
			om.mxAvg.Unlock()

			// not sure if it's necessary
			om.mxMaxM.Lock()
			max := 0
			for k := range om.avgRequired {
				if k > max {
					max = k
				}
			}
			if max > initialWindowLengthSeconds {
				om.maxM = max
			} else {
				om.maxM = initialWindowLengthSeconds
			}
			om.mxMaxM.Unlock()
		}
	}
	om.mxAR.Unlock()
}

func (om *OsMonitor) GetMAverage(m int) (*State, error) {
	om.mxAvg.RLock()
	defer om.mxAvg.RUnlock()
	var avg State
	_, ok := om.averages[m]
	if !ok {
		return nil, ErrNotCalculated
	}
	avg = *om.averages[m]
	return &avg, nil
}

func (om *OsMonitor) startCollector(ctx context.Context) {
	scheduler := time.NewTicker(time.Second)
	defer scheduler.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-scheduler.C:
			started := time.Now()
			state := om.stateCollector.GetCurrentState(ctx)
			om.mxStates.Lock()
			om.mxMaxM.RLock()
			if len(om.states) < om.maxM {
				om.states = append(om.states, state)
			} else {
				// if maxM decreased
				idx := len(om.states) - om.maxM + 1
				om.states = om.states[idx:]
				om.states = append(om.states, state)
			}
			om.mxStates.Unlock()
			om.mxMaxM.RUnlock()

			om.mxStates.RLock()
			om.log.Tracef("collect: %d microseconds", time.Since(started).Microseconds())
			om.log.Trace("_____________________________________")
			for _, state := range om.states {
				om.log.Tracef(state.String())
			}
			om.log.Trace("len(states) == ", len(om.states))
			om.log.Trace("_____________________________________")
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
			started := time.Now()
			om.mxAR.RLock()
			for k := range om.avgRequired {
				switch {
				case k > len(om.states):
					continue
				case k == len(om.states):
					om.mxAvg.Lock()
					om.averages[k] = calcAvg(om.states)
					om.mxAvg.Unlock()
				default:
					om.mxAvg.Lock()
					if _, ok := om.averages[k]; !ok {
						om.averages[k] = calcAvg(om.states)
					} else {
						startIdx := len(om.states) - k
						om.averages[k] = calcIncrementAvg(om.states[startIdx-1:], om.averages[k])
					}
					om.mxAvg.Unlock()
				}
			}
			om.mxAR.RUnlock()

			om.mxAvg.RLock()
			om.log.Tracef("calculate: %d microseconds", time.Since(started).Microseconds())
			om.log.Trace("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			for k, average := range om.averages {
				om.log.Tracef("%d: %s", k, average.String())
			}
			om.log.Trace("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			om.mxAvg.RUnlock()
		}
	}
}
