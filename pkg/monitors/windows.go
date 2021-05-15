// +build windows

package monitors

import (
	"context"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type StateCollector struct {
	log *logrus.Logger
}

func (wsc *StateCollector) GetCurrentState(ctx context.Context, metrics *MetricsPresent) *State {
	var (
		la  *LoadAverage
		cpu *CPULoad
		mem *Mem
		wg sync.WaitGroup
	)
	if metrics.la {
		wg.Add(1)
		go func() {
			defer wg.Done()
			la = wsc.getLoadAverage(ctx)
		}()
	}
	if metrics.cpu {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cpu = &CPULoad{}
		}()
	}
	if metrics.mem {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mem = &Mem{}
		}()
	}
	wg.Wait()
	return &State{LoadAverage: la, CPULoad: cpu, Mem: mem}
}

func (wsc *StateCollector) getLoadAverage(ctx context.Context) *LoadAverage {
	LOADAVG_FACTOR_1F := 0.9200444146293232478931553241
	LOADAVG_FACTOR_5F := 0.6592406302004437462547604110
	LOADAVG_FACTOR_15F := 0.2865047968601901003248854266

	var load_avg_1m float64 = 0
	var load_avg_5m float64 = 0
	var load_avg_15m float64 = 0

	la := &LoadAverage{}
	out, err := exec.CommandContext(ctx, "cmd", "/k", "wmic cpu get LoadPercentage -value").Output()
	if err != nil {
		wsc.log.Warn("err getting loadAverage: ", err)
		return la
	}

	currentLoad, err := strconv.ParseFloat(strings.Trim(string(out[21:23]), " "), 64)
	if err != nil {
		wsc.log.Warn("err getting loadAverage: ", err)
		return la
	}

	load_avg_1m = load_avg_1m*LOADAVG_FACTOR_1F + currentLoad*(1.0-LOADAVG_FACTOR_1F)
	load_avg_5m = load_avg_5m*LOADAVG_FACTOR_5F + currentLoad*(1.0-LOADAVG_FACTOR_5F)
	load_avg_15m = load_avg_15m*LOADAVG_FACTOR_15F + currentLoad*(1.0-LOADAVG_FACTOR_15F)

	la.One = load_avg_1m
	la.Five = load_avg_5m
	la.Fifteen = load_avg_15m
	return la
}
