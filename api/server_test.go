package api

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"remote-system-monitor/pkg/monitors"
	"sync"
	"testing"
	"time"
)

//Not really used atm
type TestMonitor struct {
}

func (t TestMonitor) AddMAverage(m int) {
	panic("implement me")
}

func (t TestMonitor) GetMAverage(m int) (*monitors.State, error) {
	panic("implement me")
}

//Not really a unit test atm
func TestRPCServer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	log := logrus.New()
	monitor, err := monitors.GetOsMonitor(log, "linux")
	require.NoError(t, err)
	r := NewRPCServer(log, monitor, 3002, "tcp", "version")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = r.Start(ctx)
		require.NoError(t, err)
	}()

	time.Sleep(6 * time.Second)
	r.monitor.AddMAverage(4)
	time.Sleep(2 * time.Second)
	r.monitor.AddMAverage(5)
	time.Sleep(2 * time.Second)
	r.monitor.AddMAverage(10)

	wg.Wait()
}
