package api

import (
	"context"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"remote-system-monitor/pkg/api/monitorApiv1"
	"remote-system-monitor/pkg/monitors"
)

func TestRPCServer_SignUp(t *testing.T) {
	port := 3002
	if testing.Short() {
		t.Skip()
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	log := logrus.New()
	monitor, err := monitors.GetOsMonitor(log, "linux")
	require.NoError(t, err)
	r := NewRPCServer(log, monitor, port, "tcp", "version")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = r.Start(ctx)
		require.NoError(t, err)
	}()

	time.Sleep(2 * time.Second)
	cc, err := grpc.Dial("localhost:"+strconv.Itoa(port), grpc.WithInsecure())
	require.NoError(t, err)
	client := monitorApiv1.NewSignUpHandlerClient(cc)
	_, err = client.SignUp(ctx, &monitorApiv1.SignUpRequest{
		ReportPeriod: int32(1),
		MeanPeriod:   int32(1),
	})
	require.NoError(t, err)
	cancel()
}
