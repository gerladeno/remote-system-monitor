// +build linux

package monitors

import (
	"context"
	"github.com/sirupsen/logrus"
	"math"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type StateCollector struct {
	log *logrus.Logger
}

func (lsc *StateCollector) GetCurrentState(ctx context.Context) *State {
	var (
		la  LoadAverage
		cpu CPULoad
		mem Mem
		wg  sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		la = lsc.GetLoadAverage(ctx)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		cpu = lsc.GetCPULoad(ctx)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		mem = lsc.GetMem(ctx)
	}()
	wg.Wait()
	return &State{LoadAverage: la, CPULoad: cpu, Mem: mem}
}

func (lsc *StateCollector) GetLoadAverage(ctx context.Context) LoadAverage {
	la := LoadAverage{}
	out, err := exec.CommandContext(ctx, "cat", "/proc/loadavg").Output()
	if err != nil {
		lsc.log.Warn("err processing la: ", err)
		return la
	}
	elems := strings.Split(string(out), " ")
	la.One, err = strconv.ParseFloat(elems[0], 64)
	if err != nil {
		lsc.log.Warn("err processing la: ", err)
		return la
	}
	la.Five, err = strconv.ParseFloat(elems[1], 64)
	if err != nil {
		lsc.log.Warn("err processing la: ", err)
		return la
	}
	la.Fifteen, err = strconv.ParseFloat(elems[2], 64)
	if err != nil {
		lsc.log.Warn("err processing la: ", err)
		return la
	}
	return la
}

func (lsc *StateCollector) GetCPULoad(ctx context.Context) CPULoad {
	cpu := CPULoad{}
	out, err := exec.CommandContext(ctx, "head", "-n1", "/proc/stat").Output()
	if err != nil {
		lsc.log.Warn("err processing cpu: ", err)
		return cpu
	}
	values := strings.Fields(strings.TrimSpace(string(out)))
	if len(values) != 11 {
		lsc.log.Warnf("err processing cpu: expected 11 elements in a string, got %d: %s", len(values), string(out))
		return cpu
	}
	user, err := strconv.Atoi(values[1])
	if err != nil {
		lsc.log.Warn("err processing cpu: ", err)
		return cpu
	}
	nice, err := strconv.Atoi(values[2])
	if err != nil {
		lsc.log.Warn("err processing cpu: ", err)
		return cpu
	}
	system, err := strconv.Atoi(values[3])
	if err != nil {
		lsc.log.Warn("err processing cpu: ", err)
		return cpu
	}
	idle, err := strconv.Atoi(values[4])
	if err != nil {
		lsc.log.Warn("err processing cpu: ", err)
		return cpu
	}
	iowait, err := strconv.Atoi(values[5])
	if err != nil {
		lsc.log.Warn("err processing cpu: ", err)
		return cpu
	}
	irq, err := strconv.Atoi(values[6])
	if err != nil {
		lsc.log.Warn("err processing cpu: ", err)
		return cpu
	}
	softirq, err := strconv.Atoi(values[7])
	if err != nil {
		lsc.log.Warn("err processing cpu: ", err)
		return cpu
	}
	sum := float64(user + nice + system + idle + iowait + irq + softirq)
	cpu.User = math.Round(float64(user) / sum * 1000) / 10
	cpu.System = math.Round(float64(system) / sum * 1000) / 10
	cpu.Idle = math.Round(float64(idle) / sum * 1000) / 10
	return cpu
}

func (lsc *StateCollector) GetMem(ctx context.Context) Mem {
	mem := Mem{}
	out, err := exec.CommandContext(ctx, "free").Output()
	if err != nil {
		lsc.log.Warn("err processing mem: ", err)
		return mem
	}
	lines := strings.Split(string(out), "\n")
	if len(lines) < 3 {
		lsc.log.Warn("err unexpected free output")
		return mem
	}
	memElems := strings.Fields(strings.TrimSpace(lines[1]))
	var tmp float64
	if tmp, err = strconv.ParseFloat(memElems[1], 64); err != nil {
		lsc.log.Warn("err processing mem: ", err)
		return mem
	}
	mem.Total = math.Round(10 * tmp / 1024) / 10
	if tmp, err = strconv.ParseFloat(memElems[2], 64); err != nil {
		lsc.log.Warn("err processing mem: ", err)
		return mem
	}
	mem.Used = math.Round(10 * tmp / 1024) / 10
	if tmp, err = strconv.ParseFloat(memElems[3], 64); err != nil {
		lsc.log.Warn("err processing mem: ", err)
		return mem
	}
	mem.Free = math.Round(10 * tmp / 1024) / 10
	return mem
}
