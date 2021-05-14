// build +integration

package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"os"
	"remote-system-monitor/pkg/api/monitorApiv1"
	"strconv"
	"testing"
	"time"
)

func StartClient(host string, port int) (monitorApiv1.SignUpHandlerClient, *grpc.ClientConn, error) {
	cc, err := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	client := monitorApiv1.NewSignUpHandlerClient(cc)
	return client, cc, nil
}

func TestMonitor(t *testing.T) {
	tests := []struct {
		name   string
		period int32
		report int
	}{
		{"period 2, 14 states", 2, 14},
		{"period 3, 9 states", 3, 9},
		{"period 1, 29 states", 1, 29},
		{"period 4, 7 states", 4, 7},
		{"period 2, 14 states", 5, 5},
	}
	var (
		port int
		err error
	)
	portStr := os.Getenv("PORT")
	if portStr == "" {
		port = 3000
	} else {
		port, err = strconv.Atoi(portStr)
	}
	require.NoError(t, err)
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			client, cc, err := StartClient(os.Getenv("HOST_DOCKER_INTERNAL"), port)
			require.NoError(t, err)
			defer func() {
				err = cc.Close()
				require.NoError(t, err)
			}()
			ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
			timer := time.NewTimer(30 * time.Second)
			defer cancel()
			stream, err := client.SignUp(ctx, &monitorApiv1.SignUpRequest{
				ReportPeriod: test.period,
				MeanPeriod:   2,
			})
			require.NoError(t, err)
			respCnt := 0
		LOOP:
			for {
				select {
				case <-timer.C:
					break LOOP
				default:
				}
				res, err := stream.Recv()
				require.NoError(t, err)
				require.IsType(t, &monitorApiv1.State{}, res.GetState())
				respCnt++
			}
			require.True(t, respCnt >= test.report)
			require.True(t, respCnt <= test.report+2)
		})
	}
}
