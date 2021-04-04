// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package consignment

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

// SheppingServiceClient is the client API for SheppingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SheppingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error)
}

type sheppingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSheppingServiceClient(cc grpc.ClientConnInterface) SheppingServiceClient {
	return &sheppingServiceClient{cc}
}

func (c *sheppingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/consignment.SheppingService/CreateConsignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SheppingServiceServer is the server API for SheppingService service.
// All implementations should embed UnimplementedSheppingServiceServer
// for forward compatibility
type SheppingServiceServer interface {
	CreateConsignment(context.Context, *Consignment) (*Response, error)
}

// UnimplementedSheppingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSheppingServiceServer struct {
}

func (UnimplementedSheppingServiceServer) CreateConsignment(context.Context, *Consignment) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsignment not implemented")
}

// UnsafeSheppingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SheppingServiceServer will
// result in compilation errors.
type UnsafeSheppingServiceServer interface {
	mustEmbedUnimplementedSheppingServiceServer()
}

func RegisterSheppingServiceServer(s grpc.ServiceRegistrar, srv SheppingServiceServer) {
	s.RegisterService(&SheppingService_ServiceDesc, srv)
}

func _SheppingService_CreateConsignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consignment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SheppingServiceServer).CreateConsignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.SheppingService/CreateConsignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SheppingServiceServer).CreateConsignment(ctx, req.(*Consignment))
	}
	return interceptor(ctx, in, info, handler)
}

// SheppingService_ServiceDesc is the grpc.ServiceDesc for SheppingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SheppingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "consignment.SheppingService",
	HandlerType: (*SheppingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsignment",
			Handler:    _SheppingService_CreateConsignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consignment/consignment.proto",
}
