package monitors

import (
	"context"
	"os/exec"
	"strconv"
	"strings"
)

type WindowsStateCollector struct {
}

func (wsc *WindowsStateCollector) GetCurrentState(ctx context.Context) (*State, error) {
	la, err := wsc.GetLoadAverage(ctx)
	return &State{LoadAverage: la}, err
}

func (wcs *WindowsStateCollector) GetLoadAverage(ctx context.Context) (LoadAverage, error) {
	LOADAVG_FACTOR_1F := 0.9200444146293232478931553241
	LOADAVG_FACTOR_5F := 0.6592406302004437462547604110
	LOADAVG_FACTOR_15F := 0.2865047968601901003248854266

	var load_avg_1m float64 = 0
	var load_avg_5m float64 = 0
	var load_avg_15m float64 = 0

	la := LoadAverage{}
	out, err := exec.Command("cmd", "/k", "wmic cpu get LoadPercentage -value").Output()
	if err != nil {
		return la, err
	}

	currentLoad, err := strconv.ParseFloat(strings.Trim(string(out[21:23]), " "), 64)
	if err != nil {
		return la, err
	}

	load_avg_1m = load_avg_1m*LOADAVG_FACTOR_1F + currentLoad*(1.0-LOADAVG_FACTOR_1F)
	load_avg_5m = load_avg_5m*LOADAVG_FACTOR_5F + currentLoad*(1.0-LOADAVG_FACTOR_5F)
	load_avg_15m = load_avg_15m*LOADAVG_FACTOR_15F + currentLoad*(1.0-LOADAVG_FACTOR_15F)

	la.One = load_avg_1m
	la.Five = load_avg_5m
	la.Fifteen = load_avg_15m
	return la, nil
}
