package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"remote-system-monitor/pkg/api"
	"remote-system-monitor/pkg/logging"
	"remote-system-monitor/pkg/monitors"
	"sync"
	"syscall"
)

var (
	serverPort int
	logLevel   string
	network    string
	version    = "0.0.0"
	goos       = "linux"
)

func init() {
	flag.IntVar(&serverPort, "p", 3000, "port to start gRPC server")
	flag.StringVar(&logLevel, "l", "TRACE", "log level. Accepted values: panic, fatal, error, warn, warning, info, debug, trace. Case insensitive.")
	flag.StringVar(&network, "n", "tcp", "specify protocol for server to run: tcp (default) or udp")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		fmt.Println(version)
		return
	}

	log := logging.GetLogger(logLevel)
	osMonitor, err := monitors.GetOsMonitor(log, goos)
	if err != nil {
		log.Fatalf("err initing monitor: %s", err)
	}

	server := api.NewRPCServer(log, osMonitor, serverPort, network, version)

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
