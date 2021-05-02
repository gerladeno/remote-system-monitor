package monitors

import (
	"context"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type LinuxStateCollector struct {
	log *logrus.Logger
}

func (lsc *LinuxStateCollector) GetCurrentState(ctx context.Context) *State {
	var (
		la  LoadAverage
		cpu CPULoad
		mem Mem
		wg  sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		if la, err = lsc.GetLoadAverage(ctx); err != nil {
			lsc.log.Warn("err getting loadAverage: ", err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		if cpu, mem, err = lsc.GetCPULoadAndMem(ctx); err != nil {
			lsc.log.Warn("err getting cpu and mem: ", err)
		}
	}()
	wg.Wait()
	return &State{LoadAverage: la, CPULoad: cpu, Mem: mem}
}

func (lsc *LinuxStateCollector) GetLoadAverage(ctx context.Context) (LoadAverage, error) {
	la := LoadAverage{}
	out, err := exec.CommandContext(ctx, "cat", "/proc/loadavg").Output()
	if err != nil {
		return la, err
	}
	elems := strings.Split(string(out), " ")
	la.One, err = strconv.ParseFloat(elems[0], 64)
	if err != nil {
		return la, err
	}
	la.Five, err = strconv.ParseFloat(elems[1], 64)
	if err != nil {
		return la, err
	}
	la.Fifteen, err = strconv.ParseFloat(elems[2], 64)
	if err != nil {
		return la, err
	}
	return la, nil
}

func (lsc *LinuxStateCollector) GetCPULoadAndMem(ctx context.Context) (CPULoad, Mem, error) {
	cpu := CPULoad{}
	mem := Mem{}
	out, err := exec.CommandContext(ctx, "top", "-bn1").Output()
	if err != nil {
		return CPULoad{}, Mem{}, err
	}
	lines := strings.SplitN(string(out), "\n", 5)
	cpuElems := strings.Split(strings.Replace(lines[2], ",", ".", -1), " ")
	memElems := strings.Split(strings.Replace(lines[3], ",", ".", -1), " ")
	if cpu.User, err = strconv.ParseFloat(cpuElems[1], 64); err != nil {
		return CPULoad{}, Mem{}, err
	}
	if cpu.System, err = strconv.ParseFloat(cpuElems[4], 64); err != nil {
		return CPULoad{}, Mem{}, err
	}
	if cpu.Idle, err = strconv.ParseFloat(cpuElems[9], 64); err != nil {
		return CPULoad{}, Mem{}, err
	}
	if mem.Total, err = strconv.ParseFloat(memElems[4], 64); err != nil {
		return cpu, Mem{}, err
	}
	if mem.Free, err = strconv.ParseFloat(memElems[8], 64); err != nil {
		return cpu, Mem{}, err
	}
	if mem.Used, err = strconv.ParseFloat(memElems[11], 64); err != nil {
		return cpu, Mem{}, err
	}
	return cpu, mem, nil
}
