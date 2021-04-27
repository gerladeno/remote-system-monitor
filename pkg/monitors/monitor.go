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
	GetCurrentState(ctx context.Context) (*State, error)
	GetLoadAverage(ctx context.Context) (LoadAverage, error)
}

type OsMonitor struct {
	log            *logrus.Entry
	stateCollector StateCollector
	states         []*State
	averages       map[int]*State
	avgRequired    map[int]struct{}
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
		stateContainer = &LinuxStateCollector{}
	case "darwin":
		stateContainer = &DarwinStateCollector{}
	case "windows":
		stateContainer = &WindowsStateCollector{}
	default:
		return nil, fmt.Errorf("unsupported os: %s, only linux darwin and windows are supported", goos)
	}
	monitor := OsMonitor{
		log:            log.WithField("system", "monitor"),
		stateCollector: stateContainer,
		maxM:           initialWindowLengthSeconds,
		averages:       make(map[int]*State),
		avgRequired:    make(map[int]struct{}),
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
	om.mxAR.RLock()
	_, ok := om.avgRequired[m]
	om.mxAR.RUnlock()
	if !ok {
		om.mxAR.Lock()
		om.avgRequired[m] = struct{}{}
		om.mxAR.Unlock()

		om.mxMaxM.Lock()
		if m > om.maxM {
			om.maxM = m
		}
		om.mxMaxM.Unlock()
	}
}

func (om *OsMonitor) GetMAverage(m int) (*State, error) {
	om.mxAvg.RLock()
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
			om.mxAR.Lock()
			for k, _ := range om.avgRequired {
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
			om.mxAR.Unlock()

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
