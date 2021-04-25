package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"remote-system-monitor/pkg/monitors"
	"syscall"
)

var (
	serverPort int
	version = "0.0.0"
	goos = "linux"
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

	osMonitor.Run(ctx)
}
