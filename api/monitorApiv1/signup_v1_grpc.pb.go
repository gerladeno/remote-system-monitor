// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.6.1
// source: signup_v1.proto

package monitorApiv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SignUpHandlerClient is the client API for SignUpHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignUpHandlerClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (SignUpHandler_SignUpClient, error)
}

type signUpHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewSignUpHandlerClient(cc grpc.ClientConnInterface) SignUpHandlerClient {
	return &signUpHandlerClient{cc}
}

func (c *signUpHandlerClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (SignUpHandler_SignUpClient, error) {
	stream, err := c.cc.NewStream(ctx, &SignUpHandler_ServiceDesc.Streams[0], "/monitorApiv1.SignUpHandler/SignUp", opts...)
	if err != nil {
		return nil, err
	}
	x := &signUpHandlerSignUpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SignUpHandler_SignUpClient interface {
	Recv() (*SignUpResponse, error)
	grpc.ClientStream
}

type signUpHandlerSignUpClient struct {
	grpc.ClientStream
}

func (x *signUpHandlerSignUpClient) Recv() (*SignUpResponse, error) {
	m := new(SignUpResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SignUpHandlerServer is the server API for SignUpHandler service.
// All implementations should embed UnimplementedSignUpHandlerServer
// for forward compatibility
type SignUpHandlerServer interface {
	SignUp(*SignUpRequest, SignUpHandler_SignUpServer) error
}

// UnimplementedSignUpHandlerServer should be embedded to have forward compatible implementations.
type UnimplementedSignUpHandlerServer struct {
}

func (UnimplementedSignUpHandlerServer) SignUp(*SignUpRequest, SignUpHandler_SignUpServer) error {
	return status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}

// UnsafeSignUpHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignUpHandlerServer will
// result in compilation errors.
type UnsafeSignUpHandlerServer interface {
	mustEmbedUnimplementedSignUpHandlerServer()
}

func RegisterSignUpHandlerServer(s grpc.ServiceRegistrar, srv SignUpHandlerServer) {
	s.RegisterService(&SignUpHandler_ServiceDesc, srv)
}

func _SignUpHandler_SignUp_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SignUpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SignUpHandlerServer).SignUp(m, &signUpHandlerSignUpServer{stream})
}

type SignUpHandler_SignUpServer interface {
	Send(*SignUpResponse) error
	grpc.ServerStream
}

type signUpHandlerSignUpServer struct {
	grpc.ServerStream
}

func (x *signUpHandlerSignUpServer) Send(m *SignUpResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SignUpHandler_ServiceDesc is the grpc.ServiceDesc for SignUpHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignUpHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monitorApiv1.SignUpHandler",
	HandlerType: (*SignUpHandlerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SignUp",
			Handler:       _SignUpHandler_SignUp_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "signup_v1.proto",
}