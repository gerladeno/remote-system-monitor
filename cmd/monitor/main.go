package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"remote-system-monitor/api"
	"remote-system-monitor/pkg/monitors"
	"sync"
	"syscall"
)

var (
	serverPort int
	version    = "0.0.0"
	goos       = "linux"
)

func init() {
	flag.IntVar(&serverPort, "p", 3000, "port to start gRPC server")
}

func main() {
	log := getLogger(os.Getenv("LOG_LEVEL"))
	osMonitor, err := monitors.GetOsMonitor(log, goos)
	if err != nil {
		log.Fatalf("err initing monitor: %s", err)
	}

	server := api.NewRPCServer(log, osMonitor, serverPort, "tcp", version)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP)

		select {
		case <-ctx.Done():
			return
		case <-sigCh:
		}

		signal.Stop(sigCh)
		cancel()
	}()

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		osMonitor.Run(ctx)
	}()
	go func() {
		defer wg.Done()
		if err = server.Start(ctx); err != nil {
			log.Fatalf("err starting grpc server, %s", err)
		}
	}()
	wg.Wait()
}
