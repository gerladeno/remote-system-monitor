package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"remote-system-monitor/api/monitorApiv1"
	"strconv"
)

func main() {
	client, cc, err := StartClient(3000)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := cc.Close()
		if err != nil {
			panic(err)
		}
	}()
	ctx := context.Background()
	stream, err := client.SignUp(ctx, &monitorApiv1.SignUpRequest{
		Version:      1,
		ReportPeriod: 3,
		MeanPeriod:   9,
	})
	if err != nil {
		panic(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(res.GetState().String())
	}
}

func StartClient(port int) (monitorApiv1.SignUpHandlerClient, *grpc.ClientConn, error) {
	cc, err := grpc.Dial("localhost:"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	client := monitorApiv1.NewSignUpHandlerClient(cc)
	return client, cc, nil
}
