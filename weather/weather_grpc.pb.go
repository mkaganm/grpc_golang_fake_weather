// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: weather/weather.proto

package weather

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

// WeatherClient is the client API for Weather service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherClient interface {
	GetWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherResponse, error)
	StreamWeatherUpdates(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (Weather_StreamWeatherUpdatesClient, error)
}

type weatherClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherClient(cc grpc.ClientConnInterface) WeatherClient {
	return &weatherClient{cc}
}

func (c *weatherClient) GetWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherResponse, error) {
	out := new(WeatherResponse)
	err := c.cc.Invoke(ctx, "/weather.Weather/GetWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherClient) StreamWeatherUpdates(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (Weather_StreamWeatherUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Weather_ServiceDesc.Streams[0], "/weather.Weather/StreamWeatherUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &weatherStreamWeatherUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Weather_StreamWeatherUpdatesClient interface {
	Recv() (*WeatherResponse, error)
	grpc.ClientStream
}

type weatherStreamWeatherUpdatesClient struct {
	grpc.ClientStream
}

func (x *weatherStreamWeatherUpdatesClient) Recv() (*WeatherResponse, error) {
	m := new(WeatherResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WeatherServer is the server API for Weather service.
// All implementations must embed UnimplementedWeatherServer
// for forward compatibility
type WeatherServer interface {
	GetWeather(context.Context, *WeatherRequest) (*WeatherResponse, error)
	StreamWeatherUpdates(*WeatherRequest, Weather_StreamWeatherUpdatesServer) error
	mustEmbedUnimplementedWeatherServer()
}

// UnimplementedWeatherServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServer struct {
}

func (UnimplementedWeatherServer) GetWeather(context.Context, *WeatherRequest) (*WeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeather not implemented")
}
func (UnimplementedWeatherServer) StreamWeatherUpdates(*WeatherRequest, Weather_StreamWeatherUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamWeatherUpdates not implemented")
}
func (UnimplementedWeatherServer) mustEmbedUnimplementedWeatherServer() {}

// UnsafeWeatherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServer will
// result in compilation errors.
type UnsafeWeatherServer interface {
	mustEmbedUnimplementedWeatherServer()
}

func RegisterWeatherServer(s grpc.ServiceRegistrar, srv WeatherServer) {
	s.RegisterService(&Weather_ServiceDesc, srv)
}

func _Weather_GetWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServer).GetWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weather.Weather/GetWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServer).GetWeather(ctx, req.(*WeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weather_StreamWeatherUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WeatherRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WeatherServer).StreamWeatherUpdates(m, &weatherStreamWeatherUpdatesServer{stream})
}

type Weather_StreamWeatherUpdatesServer interface {
	Send(*WeatherResponse) error
	grpc.ServerStream
}

type weatherStreamWeatherUpdatesServer struct {
	grpc.ServerStream
}

func (x *weatherStreamWeatherUpdatesServer) Send(m *WeatherResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Weather_ServiceDesc is the grpc.ServiceDesc for Weather service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Weather_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weather.Weather",
	HandlerType: (*WeatherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWeather",
			Handler:    _Weather_GetWeather_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamWeatherUpdates",
			Handler:       _Weather_StreamWeatherUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "weather/weather.proto",
}
