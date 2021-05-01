package api

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"remote-system-monitor/api/monitorApiv1"
	"remote-system-monitor/pkg/monitors"
	"strconv"
	"time"
)

//go:generate protoc -I=proto/ proto/signup_v1.proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:.

type Monitor interface {
	AddMAverage(m int)
	GetMAverage(m int) (*monitors.State, error)
}

type RPCServer struct {
	log     *logrus.Logger
	monitor Monitor
	network string
	port    int
	server  *grpc.Server
	version string
}

func NewRPCServer(log *logrus.Logger, monitor Monitor, port int, network, version string) *RPCServer {
	return &RPCServer{
		log:     log,
		monitor: monitor,
		port:    port,
		network: network,
		server:  grpc.NewServer(),
		version: version,
	}
}

func (r *RPCServer) Start(ctx context.Context) error {
	l, err := net.Listen(r.network, ":"+strconv.Itoa(r.port))
	if err != nil {
		return err
	}
	reflection.Register(r.server)
	monitorApiv1.RegisterSignUpHandlerServer(r.server, r)
	go func() {
		<-ctx.Done()
		r.Stop()
	}()
	if err = r.server.Serve(l); err != nil {
		return err
	}
	return nil
}

func (r *RPCServer) Stop() {
	r.server.Stop()
}

func (r *RPCServer) SignUp(request *monitorApiv1.SignUpRequest, stream monitorApiv1.SignUpHandler_SignUpServer) error {
	m := int(request.GetMeanPeriod())
	r.monitor.AddMAverage(m)
	timer := time.NewTicker(time.Duration(request.GetReportPeriod()) * time.Second)
LOOP:
	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-timer.C:
			avg, err := r.monitor.GetMAverage(m)
			if err != nil {
				r.log.Warn("err receiving data from monitor: ", err)
				continue
			}
			err = stream.Send(&monitorApiv1.SignUpResponse{State: state2Pb(avg)})
			if err != nil {
				break LOOP
			}
		}
	}
	return nil
}

func state2Pb(state *monitors.State) *monitorApiv1.State {
	pbState := monitorApiv1.State{
		LoadAverage: &monitorApiv1.LoadAverage{
			One:     state.LoadAverage.One,
			Five:    state.LoadAverage.Five,
			Fifteen: state.LoadAverage.Fifteen,
		},
	}
	return &pbState
}
