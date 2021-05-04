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

func (dsc *StateCollector) GetCurrentState(ctx context.Context) *State {
	var (
		la  LoadAverage
		cpu CPULoad
		mem Mem
		wg  sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		la, cpu, mem = dsc.GetLoadAverageCPULoadMem(ctx)
	}()
	wg.Wait()
	return &State{LoadAverage: la, CPULoad: cpu, Mem: mem}
}

func (dsc *StateCollector) GetLoadAverageCPULoadMem(ctx context.Context) (LoadAverage, CPULoad, Mem) {
	la := LoadAverage{}
	cpu := CPULoad{}
	mem := Mem{}
	out, err := exec.CommandContext(ctx, "top", "-n", "0", "-l", "1").Output()
	if err != nil {
		dsc.log.Warn("err executing top: ", err)
		return la, cpu, mem
	}
	lines := strings.Split(string(out), "\n")
	laElems := strings.Fields(strings.TrimSpace(lines[2]))
	cpuElems := strings.Fields(strings.TrimSpace(lines[3]))
	memElems := strings.Fields(strings.TrimSpace(lines[6]))
	la.One, err = strconv.ParseFloat(strings.TrimRight(laElems[2], ","), 64)
	if err != nil {
		dsc.log.Warn("err processing la: ", err)
		return la, cpu, mem
	}
	la.Five, err = strconv.ParseFloat(strings.TrimRight(laElems[3], ","), 64)
	if err != nil {
		dsc.log.Warn("err processing la: ", err)
		return la, cpu, mem
	}
	la.Fifteen, err = strconv.ParseFloat(strings.TrimRight(laElems[4], ","), 64)
	if err != nil {
		dsc.log.Warn("err processing la: ", err)
		return la, cpu, mem
	}
	if cpu.User, err = strconv.ParseFloat(strings.TrimRight(cpuElems[2], "%"), 64); err != nil {
		dsc.log.Warn("err collecting cpu: ", err)
		return la, cpu, mem
	}
	if cpu.System, err = strconv.ParseFloat(strings.TrimRight(cpuElems[4], "%"), 64); err != nil {
		dsc.log.Warn("err collecting cpu: ", err)
		return la, cpu, mem
	}
	if cpu.Idle, err = strconv.ParseFloat(strings.TrimRight(cpuElems[6], "%"), 64); err != nil {
		dsc.log.Warn("err collecting cpu: ", err)
		return la, cpu, mem
	}
	if mem.Free, err = parseMemValue(memElems[5], 64); err != nil {
		dsc.log.Warn("err collecting mem: ", err)
		return la, cpu, mem
	}
	if mem.Used, err = parseMemValue(memElems[1], 64); err != nil {
		dsc.log.Warn("err collecting mem: ", err)
		return la, cpu, mem
	}
	mem.Total = mem.Used + mem.Free
	return la, cpu, mem
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
