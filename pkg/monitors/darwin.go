package monitors

import (
	"context"
	"os/exec"
	"strconv"
	"strings"
)

type DarwinStateCollector struct {
}

func (dsc *DarwinStateCollector) GetCurrentState(ctx context.Context) (*State, error) {
	la, err := dsc.GetLoadAverage(ctx)
	return &State{LoadAverage: la}, err
}

func (dsc *DarwinStateCollector) GetLoadAverage(ctx context.Context) (LoadAverage, error) {
	la := LoadAverage{}
	out, err := exec.CommandContext(ctx, "uptime").Output()
	if err != nil {
		return la, err
	}
	elems := strings.Split(strings.Replace(string(out), ",", ".", -1), " ")
	laOneIdx := len(elems) - 3
	la.One, err = strconv.ParseFloat(elems[laOneIdx], 64)
	if err != nil {
		return la, err
	}
	la.Five, err = strconv.ParseFloat(elems[laOneIdx+1], 64)
	if err != nil {
		return la, err
	}
	la.Fifteen, err = strconv.ParseFloat(elems[laOneIdx+2][:len(elems[laOneIdx+2])-1], 64)
	if err != nil {
		return la, err
	}
	return la, nil
}
