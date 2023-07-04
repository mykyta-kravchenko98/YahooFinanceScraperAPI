// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: api/yahooScraper_v1/yahooStocks.proto

package yahooscraper_v1

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

const (
	YahooStocksService_GetAllValidStocks_FullMethodName = "/yahooscraper_v1.YahooStocksService/GetAllValidStocks"
)

// YahooStocksServiceClient is the client API for YahooStocksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YahooStocksServiceClient interface {
	GetAllValidStocks(ctx context.Context, in *GetAllValidStocksRequest, opts ...grpc.CallOption) (*GetAllValidStocksResponse, error)
}

type yahooStocksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYahooStocksServiceClient(cc grpc.ClientConnInterface) YahooStocksServiceClient {
	return &yahooStocksServiceClient{cc}
}

func (c *yahooStocksServiceClient) GetAllValidStocks(ctx context.Context, in *GetAllValidStocksRequest, opts ...grpc.CallOption) (*GetAllValidStocksResponse, error) {
	out := new(GetAllValidStocksResponse)
	err := c.cc.Invoke(ctx, YahooStocksService_GetAllValidStocks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YahooStocksServiceServer is the server API for YahooStocksService service.
// All implementations must embed UnimplementedYahooStocksServiceServer
// for forward compatibility
type YahooStocksServiceServer interface {
	GetAllValidStocks(context.Context, *GetAllValidStocksRequest) (*GetAllValidStocksResponse, error)
	mustEmbedUnimplementedYahooStocksServiceServer()
}

// UnimplementedYahooStocksServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYahooStocksServiceServer struct {
}

func (UnimplementedYahooStocksServiceServer) GetAllValidStocks(context.Context, *GetAllValidStocksRequest) (*GetAllValidStocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllValidStocks not implemented")
}
func (UnimplementedYahooStocksServiceServer) mustEmbedUnimplementedYahooStocksServiceServer() {}

// UnsafeYahooStocksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YahooStocksServiceServer will
// result in compilation errors.
type UnsafeYahooStocksServiceServer interface {
	mustEmbedUnimplementedYahooStocksServiceServer()
}

func RegisterYahooStocksServiceServer(s grpc.ServiceRegistrar, srv YahooStocksServiceServer) {
	s.RegisterService(&YahooStocksService_ServiceDesc, srv)
}

func _YahooStocksService_GetAllValidStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllValidStocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YahooStocksServiceServer).GetAllValidStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YahooStocksService_GetAllValidStocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YahooStocksServiceServer).GetAllValidStocks(ctx, req.(*GetAllValidStocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YahooStocksService_ServiceDesc is the grpc.ServiceDesc for YahooStocksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YahooStocksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yahooscraper_v1.YahooStocksService",
	HandlerType: (*YahooStocksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllValidStocks",
			Handler:    _YahooStocksService_GetAllValidStocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/yahooScraper_v1/yahooStocks.proto",
}