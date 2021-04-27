package monitors

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

//Not really a unit test atm
func TestRPCServer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	log := logrus.New()
	monitor, err := GetOsMonitor(log, "linux")
	go func() {
		monitor.Run(ctx)
	}()
	require.NoError(t, err)

	time.Sleep(6 * time.Second)
	monitor.AddMAverage(4)
	time.Sleep(2 * time.Second)
	monitor.AddMAverage(5)
	time.Sleep(2 * time.Second)
	monitor.AddMAverage(10)
	time.Sleep(10 * time.Second)

	<-ctx.Done()
}