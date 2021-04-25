package monitors

import (
	"context"
	"os/exec"
	"strconv"
	"strings"
)

func (la *LoadAverage) GetWindowsLoadAverage(ctx context.Context) error {
	LOADAVG_FACTOR_1F := 0.9200444146293232478931553241
	LOADAVG_FACTOR_5F := 0.6592406302004437462547604110
	LOADAVG_FACTOR_15F := 0.2865047968601901003248854266

	var load_avg_1m float64 = 0
	var load_avg_5m float64 = 0
	var load_avg_15m float64 = 0

	out, err := exec.CommandContext(ctx, "wmic cpu get LoadPercentage /value").Output()
	if err != nil {
		return err
	}

	currentLoad, err := strconv.ParseFloat(strings.Split(string(out), "=")[1], 64)
	if err != nil {
		return err
	}

	load_avg_1m = load_avg_1m*LOADAVG_FACTOR_1F + currentLoad*(1.0-LOADAVG_FACTOR_1F)
	load_avg_5m = load_avg_5m*LOADAVG_FACTOR_5F + currentLoad*(1.0-LOADAVG_FACTOR_5F)
	load_avg_15m = load_avg_15m*LOADAVG_FACTOR_15F + currentLoad*(1.0-LOADAVG_FACTOR_15F)

	la.One = load_avg_1m
	la.Five = load_avg_5m
	la.Fifteen = load_avg_15m
	return nil
}
