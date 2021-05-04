package monitors

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

const defaultWindowLengthSeconds = 5

var ErrNotCalculated = errors.New("err this average is not yet calculated")

type OsMonitor struct {
	log            *logrus.Entry
	stateCollector StateCollector
	states         []*State
	averages       map[int]*State
	avgRequired    map[int]int
	maxM           int
	mx             sync.RWMutex
}

func GetOsMonitor(log *logrus.Logger, goos string) (*OsMonitor, error) {
	var stateContainer StateCollector
	switch goos {
	case "linux", "darwin", "windows":
		stateContainer = StateCollector{log: log}
	default:
		return nil, fmt.Errorf("unsupported os: %s, only linux darwin and windows are supported", goos)
	}
	monitor := OsMonitor{
		log:            log.WithField("system", "monitor"),
		stateCollector: stateContainer,
		maxM:           defaultWindowLengthSeconds,
		averages:       make(map[int]*State),
		avgRequired:    make(map[int]int),
	}
	return &monitor, nil
}

func (om *OsMonitor) Run(ctx context.Context) {
	scheduler := time.NewTicker(time.Second)
	defer scheduler.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-scheduler.C:
			startedCollect := time.Now()
			state := om.stateCollector.GetCurrentState(ctx)
			om.mx.Lock()
			if len(om.states) < om.maxM {
				om.states = append(om.states, state)
			} else {
				// if maxM decreased
				idx := len(om.states) - om.maxM + 1
				om.states = om.states[idx:]
				om.states = append(om.states, state)
			}

			startedCalculate := time.Now()
			for k := range om.avgRequired {
				switch {
				case k > len(om.states):
					continue
				case k == len(om.states):
					om.averages[k] = calcAvg(om.states)
				default:
					if _, ok := om.averages[k]; !ok {
						om.averages[k] = calcAvg(om.states)
					} else {
						startIdx := len(om.states) - k
						om.averages[k] = calcIncrementAvg(om.states[startIdx-1:], om.averages[k])
					}
				}
			}
			om.traceLog(startedCollect, startedCalculate)
			om.mx.Unlock()
		}
	}
}

func (om *OsMonitor) traceLog(startedCollect, startedCalculate time.Time) {
	om.log.Tracef("collect: %d microseconds", time.Since(startedCollect).Microseconds())
	om.log.Trace("_____________________________________")
	for _, state := range om.states {
		om.log.Tracef(state.String())
	}
	om.log.Trace("len(states) == ", len(om.states))
	om.log.Trace("_____________________________________")
	om.log.Tracef("calculate: %d microseconds", time.Since(startedCalculate).Microseconds())
	om.log.Trace("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	for k, average := range om.averages {
		om.log.Tracef("%d: %s", k, average.String())
	}
	om.log.Trace("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
}

func (om *OsMonitor) AddMAverage(m int) {
	om.mx.Lock()
	defer om.mx.Unlock()
	_, ok := om.avgRequired[m]
	if ok {
		om.avgRequired[m] += 1
	} else {
		om.avgRequired[m] = 1
	}

	if m > om.maxM {
		om.maxM = m
	}
}

func (om *OsMonitor) RemoveMAverage(m int) {
	om.mx.Lock()
	defer om.mx.Unlock()
	_, ok := om.avgRequired[m]
	if ok {
		if om.avgRequired[m] > 1 {
			om.avgRequired[m] -= 1
		} else {
			delete(om.avgRequired, m)
			delete(om.averages, m)

			max := 0
			for k := range om.avgRequired {
				if k > max {
					max = k
				}
			}
			if max > defaultWindowLengthSeconds {
				om.maxM = max
			} else {
				om.maxM = defaultWindowLengthSeconds
			}
		}
	}
}

func (om *OsMonitor) GetMAverage(m int) (*State, error) {
	om.mx.RLock()
	defer om.mx.RUnlock()
	var avg State
	_, ok := om.averages[m]
	if !ok {
		return nil, ErrNotCalculated
	}
	avg = *om.averages[m]
	return &avg, nil
}
