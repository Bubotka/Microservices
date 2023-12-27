// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/protoc/geo_provider.proto

package geo

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

// GeoProviderClient is the client API for GeoProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeoProviderClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	GeoCode(ctx context.Context, in *GeoRequest, opts ...grpc.CallOption) (*GeoResponse, error)
}

type geoProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoProviderClient(cc grpc.ClientConnInterface) GeoProviderClient {
	return &geoProviderClient{cc}
}

func (c *geoProviderClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/geo_provider.GeoProvider/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoProviderClient) GeoCode(ctx context.Context, in *GeoRequest, opts ...grpc.CallOption) (*GeoResponse, error) {
	out := new(GeoResponse)
	err := c.cc.Invoke(ctx, "/geo_provider.GeoProvider/GeoCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoProviderServer is the server API for GeoProvider service.
// All implementations must embed UnimplementedGeoProviderServer
// for forward compatibility
type GeoProviderServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	GeoCode(context.Context, *GeoRequest) (*GeoResponse, error)
	mustEmbedUnimplementedGeoProviderServer()
}

// UnimplementedGeoProviderServer must be embedded to have forward compatible implementations.
type UnimplementedGeoProviderServer struct {
}

func (UnimplementedGeoProviderServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedGeoProviderServer) GeoCode(context.Context, *GeoRequest) (*GeoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoCode not implemented")
}
func (UnimplementedGeoProviderServer) mustEmbedUnimplementedGeoProviderServer() {}

// UnsafeGeoProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeoProviderServer will
// result in compilation errors.
type UnsafeGeoProviderServer interface {
	mustEmbedUnimplementedGeoProviderServer()
}

func RegisterGeoProviderServer(s grpc.ServiceRegistrar, srv GeoProviderServer) {
	s.RegisterService(&GeoProvider_ServiceDesc, srv)
}

func _GeoProvider_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoProviderServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo_provider.GeoProvider/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoProviderServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoProvider_GeoCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoProviderServer).GeoCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo_provider.GeoProvider/GeoCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoProviderServer).GeoCode(ctx, req.(*GeoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GeoProvider_ServiceDesc is the grpc.ServiceDesc for GeoProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeoProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geo_provider.GeoProvider",
	HandlerType: (*GeoProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _GeoProvider_Search_Handler,
		},
		{
			MethodName: "GeoCode",
			Handler:    _GeoProvider_GeoCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protoc/geo_provider.proto",
}
