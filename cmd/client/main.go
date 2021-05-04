package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"remote-system-monitor/pkg/api/monitorApiv1"
	"remote-system-monitor/pkg/logging"
	"strconv"
	"time"
)

var (
	clientId     string
	meanPeriod   int
	reportPeriod int
	serverPort   int
	serverHost   string
)

func init() {
	flag.StringVar(&clientId, "i", "", "client id to distinguish their output in stdout")
	flag.IntVar(&reportPeriod, "n", 5, "streaming period")
	flag.IntVar(&meanPeriod, "m", 8, "average period")
	flag.IntVar(&serverPort, "p", 3000, "port search for gRPC server")
	flag.StringVar(&serverHost, "h", "localhost", "host to search for gRPC server")
}

func main() {
	flag.Parse()
	log := logging.GetLogger("DEBUG")
	client, cc, err := StartClient(serverHost, serverPort)
	if err != nil {
		log.Fatal("err starting client: ",err)
	}
	defer func() {
		err := cc.Close()
		if err != nil {
			log.Fatalf("err closing client: %s", err)
		}
	}()
	ctx := context.Background()
	stream, err := client.SignUp(ctx, &monitorApiv1.SignUpRequest{
		ReportPeriod: int32(reportPeriod),
		MeanPeriod:   int32(meanPeriod),
	})
	if err != nil {
		log.Fatalf("err connecting to server: %s", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Infof("reached EOF, stopping client")
			return
		}
		if err != nil {
			log.Warnf("err receiving streaming data: %s", err)
			return
		}
		fmt.Printf("[%s][%s]: %s avged by %d seconds\n", clientId, time.Now().Format("2006-01-02 15:04:05"), res.GetState().String(), meanPeriod)
	}
}

func StartClient(host string, port int) (monitorApiv1.SignUpHandlerClient, *grpc.ClientConn, error) {
	cc, err := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	client := monitorApiv1.NewSignUpHandlerClient(cc)
	return client, cc, nil
}
