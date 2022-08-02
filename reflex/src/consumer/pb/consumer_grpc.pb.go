// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: consumer.proto

package pb

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

// ProducerClient is the client API for Producer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProducerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type producerClient struct {
	cc grpc.ClientConnInterface
}

func NewProducerClient(cc grpc.ClientConnInterface) ProducerClient {
	return &producerClient{cc}
}

func (c *producerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/Producer/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProducerServer is the server API for Producer service.
// All implementations must embed UnimplementedProducerServer
// for forward compatibility
type ProducerServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedProducerServer()
}

// UnimplementedProducerServer must be embedded to have forward compatible implementations.
type UnimplementedProducerServer struct {
}

func (UnimplementedProducerServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedProducerServer) mustEmbedUnimplementedProducerServer() {}

// UnsafeProducerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProducerServer will
// result in compilation errors.
type UnsafeProducerServer interface {
	mustEmbedUnimplementedProducerServer()
}

func RegisterProducerServer(s grpc.ServiceRegistrar, srv ProducerServer) {
	s.RegisterService(&Producer_ServiceDesc, srv)
}

func _Producer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Producer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Producer_ServiceDesc is the grpc.ServiceDesc for Producer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Producer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Producer",
	HandlerType: (*ProducerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Producer_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consumer.proto",
}
