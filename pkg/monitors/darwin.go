// +build darwin

package monitors

import (
	"context"
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type StateCollector struct {
	log *logrus.Logger
}

func (dsc *StateCollector) GetCurrentState(ctx context.Context, metrics *MetricsPresent) *State {
	var (
		la  *LoadAverage
		cpu *CPULoad
		mem *Mem
		wg  sync.WaitGroup
	)
	output := dsc.getTopOutput(ctx)
	if metrics.la {
		wg.Add(1)
		go func() {
			defer wg.Done()
			la = dsc.getLoadAverage(output)
		}()
	}
	if metrics.cpu {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cpu = dsc.getCPULoad(output)
		}()
	}
	if metrics.mem {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mem = dsc.getMem(output)
		}()
	}
	wg.Wait()
	return &State{LoadAverage: la, CPULoad: cpu, Mem: mem}
}

func (dsc *StateCollector) getTopOutput(ctx context.Context) []string {
	out, err := exec.CommandContext(ctx, "top", "-n", "0", "-l", "1").Output()
	if err != nil {
		dsc.log.Warn("err executing top: ", err)
		return nil
	}
	lines := strings.Split(string(out), "\n")
	return lines
}

func (dsc *StateCollector) getLoadAverage(output []string) *LoadAverage {
	var err error
	la := &LoadAverage{}
	laElems := strings.Fields(strings.TrimSpace(output[2]))
	la.One, err = strconv.ParseFloat(strings.TrimRight(laElems[2], ","), 64)
	if err != nil {
		dsc.log.Warn("err processing la: ", err)
		return la
	}
	la.Five, err = strconv.ParseFloat(strings.TrimRight(laElems[3], ","), 64)
	if err != nil {
		dsc.log.Warn("err processing la: ", err)
		return la
	}
	la.Fifteen, err = strconv.ParseFloat(strings.TrimRight(laElems[4], ","), 64)
	if err != nil {
		dsc.log.Warn("err processing la: ", err)
		return la
	}
	return la
}

func (dsc *StateCollector) getCPULoad(output []string) *CPULoad {
	var err error
	cpu := &CPULoad{}
	cpuElems := strings.Fields(strings.TrimSpace(output[3]))
	if cpu.User, err = strconv.ParseFloat(strings.TrimRight(cpuElems[2], "%"), 64); err != nil {
		dsc.log.Warn("err collecting cpu: ", err)
		return cpu
	}
	if cpu.System, err = strconv.ParseFloat(strings.TrimRight(cpuElems[4], "%"), 64); err != nil {
		dsc.log.Warn("err collecting cpu: ", err)
		return cpu
	}
	if cpu.Idle, err = strconv.ParseFloat(strings.TrimRight(cpuElems[6], "%"), 64); err != nil {
		dsc.log.Warn("err collecting cpu: ", err)
		return cpu
	}
	return cpu
}

func (dsc *StateCollector) getMem(output []string) *Mem {
	var err error
	mem := &Mem{}
	memElems := strings.Fields(strings.TrimSpace(output[6]))
	if mem.Free, err = parseMemValue(memElems[5], 64); err != nil {
		dsc.log.Warn("err collecting mem: ", err)
		return mem
	}
	if mem.Used, err = parseMemValue(memElems[1], 64); err != nil {
		dsc.log.Warn("err collecting mem: ", err)
		return mem
	}
	mem.Total = mem.Used + mem.Free
	return mem
}

func parseMemValue(s string, bitSize int) (float64, error) {
	val, scale := s[:len(s)-1], s[len(s)-1]
	result, err := strconv.ParseFloat(val, bitSize)
	if err != nil {
		return 0, err
	}
	switch scale {
	case 'M':
		return result, nil
	case 'G':
		return result * 1024, nil
	case 'K':
		return math.Round(10*result/1024) / 10, nil
	default:
		return 0, fmt.Errorf("unknown symbols %s in %s", string(scale), s)
	}
}
