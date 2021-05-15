package monitors

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestOsMonitor_AddRemoveMAverage(t *testing.T) {
	metrics := MetricsPresent{true, true, true}
	monitor, err := GetOsMonitor(logrus.New(), "linux", &metrics)
	require.NoError(t, err)
	require.Len(t, monitor.avgRequired, 0)
	monitor.AddMAverage(4)
	require.Len(t, monitor.avgRequired, 1)
	monitor.AddMAverage(5)
	require.Len(t, monitor.avgRequired, 2)
	monitor.AddMAverage(5)
	require.Len(t, monitor.avgRequired, 2)
	monitor.RemoveMAverage(4)
	require.Len(t, monitor.avgRequired, 1)
	monitor.RemoveMAverage(5)
	require.Len(t, monitor.avgRequired, 1)
	monitor.RemoveMAverage(5)
	require.Len(t, monitor.avgRequired, 0)
}

func TestOsMonitor_Concurrency(t *testing.T) {
	metrics := MetricsPresent{true, true, true}
	monitor, err := GetOsMonitor(logrus.New(), "linux", &metrics)
	require.NoError(t, err)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			monitor.AddMAverage(i)
		}(i)
	}
	wg.Wait()
	require.Len(t, monitor.avgRequired, 100)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			monitor.RemoveMAverage(i)
		}(i)
	}
	for i := 100; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			monitor.AddMAverage(i)
		}(i)
	}
	wg.Wait()
	require.Len(t, monitor.avgRequired, 100)
}

func TestOsMonitor_MaxM(t *testing.T) {
	metrics := MetricsPresent{true, true, true}
	monitor, err := GetOsMonitor(logrus.New(), "linux", &metrics)
	require.NoError(t, err)
	require.Equal(t, monitor.maxM, defaultWindowLengthSeconds)
	monitor.AddMAverage(defaultWindowLengthSeconds)
	require.Equal(t, monitor.maxM, defaultWindowLengthSeconds)
	monitor.AddMAverage(defaultWindowLengthSeconds + 10)
	require.Equal(t, monitor.maxM, defaultWindowLengthSeconds+10)
}

func TestOsMonitor_Run(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	log := logrus.New()
	metrics := MetricsPresent{true, true, true}
	monitor, err := GetOsMonitor(log, "linux", &metrics)
	require.NoError(t, err)
	go func() {
		monitor.Run(ctx)
	}()
	time.Sleep(time.Second)
	monitor.mx.Lock()
	require.Len(t, monitor.avgRequired, 0)
	require.Len(t, monitor.averages, 0)
	monitor.mx.Unlock()

	monitor.AddMAverage(4)
	time.Sleep(4 * time.Second)
	monitor.mx.Lock()
	require.True(t, len(monitor.states) > 0)
	require.Len(t, monitor.avgRequired, 1)
	require.Len(t, monitor.averages, 1)
	monitor.mx.Unlock()
}
