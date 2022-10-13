// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/cmd/auth/grpc/auth.proto

package grpc

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("github.com/elojah/game_03/cmd/auth/grpc/auth.proto", fileDescriptor_d94f16348cd64e2c)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/cmd/auth/grpc/auth.proto", fileDescriptor_d94f16348cd64e2c)
}

var fileDescriptor_d94f16348cd64e2c = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x8f, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0x77, 0x20, 0x58, 0x04, 0xab, 0x14, 0x16, 0x51, 0xe6, 0x0d, 0xcc, 0xca, 0x1d, 0xd8,
	0x2b, 0x88, 0x62, 0x25, 0x9c, 0xd8, 0x4a, 0x12, 0xd7, 0xc9, 0x4a, 0x92, 0x0d, 0x71, 0x97, 0xc3,
	0xce, 0x47, 0xf0, 0x31, 0x7c, 0x04, 0xcb, 0x2b, 0xaf, 0x4c, 0x25, 0x57, 0xba, 0x9b, 0xc6, 0xf2,
	0x4a, 0x4b, 0xb9, 0x8d, 0x72, 0xa2, 0x68, 0x63, 0x37, 0x33, 0xdf, 0xff, 0xfd, 0x30, 0xe1, 0x88,
	0xa4, 0x2e, 0x4c, 0x96, 0xe4, 0xaa, 0xe2, 0xa2, 0x54, 0x37, 0x69, 0xc1, 0x29, 0xad, 0xc4, 0xe5,
	0xde, 0x98, 0xe7, 0xd5, 0x15, 0x4f, 0x8d, 0x2e, 0x38, 0xb5, 0x4d, 0xee, 0xa7, 0xa4, 0x69, 0x95,
	0x56, 0x51, 0xb0, 0x3a, 0xc4, 0xbb, 0x5f, 0x4c, 0x52, 0xa4, 0xb8, 0x87, 0x99, 0xb9, 0xf6, 0x9b,
	0x5f, 0xfc, 0x34, 0x48, 0xf1, 0x36, 0x29, 0x45, 0xa5, 0x58, 0xa7, 0x44, 0xd5, 0xe8, 0xbb, 0x0f,
	0x88, 0xdf, 0xe1, 0xb4, 0x4d, 0x9b, 0x46, 0xb4, 0xb7, 0x03, 0x1f, 0x3d, 0x43, 0x18, 0x1c, 0x18,
	0x5d, 0x44, 0xa7, 0xe1, 0xe6, 0x44, 0x52, 0x2d, 0xeb, 0xf3, 0xa9, 0xd4, 0x79, 0x11, 0xed, 0x24,
	0x83, 0x99, 0x7c, 0x9a, 0xc9, 0x44, 0xb7, 0xb2, 0xa6, 0x8b, 0xb4, 0x34, 0x22, 0xfe, 0x93, 0xae,
	0xbb, 0x8e, 0x7d, 0xe8, 0x5f, 0x5d, 0xfb, 0x61, 0x70, 0x26, 0x6b, 0x8a, 0xb6, 0x7e, 0xa4, 0x8e,
	0x56, 0x6f, 0xc6, 0xbf, 0xdc, 0x0f, 0x4f, 0xe6, 0x16, 0x59, 0x67, 0x91, 0x2d, 0x2c, 0xb2, 0xa5,
	0x45, 0x78, 0xb3, 0x08, 0xf7, 0x0e, 0xe1, 0xd1, 0x21, 0x3c, 0x39, 0x84, 0x99, 0x43, 0x98, 0x3b,
	0x84, 0xce, 0x21, 0xbc, 0x38, 0x84, 0x57, 0x87, 0x6c, 0xe9, 0x10, 0x1e, 0x7a, 0x64, 0xb3, 0x1e,
	0xa1, 0xeb, 0x91, 0x2d, 0x7a, 0x64, 0xd9, 0x86, 0x6f, 0x1e, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x57, 0xb0, 0x1a, 0xf1, 0xd1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	// Signin Twitch
	SigninTwitch(ctx context.Context, in *types.StringValue, opts ...grpc.CallOption) (*types.StringValue, error)
	// Signin Google
	SigninGoogle(ctx context.Context, in *types.StringValue, opts ...grpc.CallOption) (*types.StringValue, error)
	// Ping
	Ping(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) SigninTwitch(ctx context.Context, in *types.StringValue, opts ...grpc.CallOption) (*types.StringValue, error) {
	out := new(types.StringValue)
	err := c.cc.Invoke(ctx, "/grpc.Auth/SigninTwitch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SigninGoogle(ctx context.Context, in *types.StringValue, opts ...grpc.CallOption) (*types.StringValue, error) {
	out := new(types.StringValue)
	err := c.cc.Invoke(ctx, "/grpc.Auth/SigninGoogle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Ping(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Auth/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	// Signin Twitch
	SigninTwitch(context.Context, *types.StringValue) (*types.StringValue, error)
	// Signin Google
	SigninGoogle(context.Context, *types.StringValue) (*types.StringValue, error)
	// Ping
	Ping(context.Context, *types.Empty) (*types.Empty, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) SigninTwitch(ctx context.Context, req *types.StringValue) (*types.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigninTwitch not implemented")
}
func (*UnimplementedAuthServer) SigninGoogle(ctx context.Context, req *types.StringValue) (*types.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigninGoogle not implemented")
}
func (*UnimplementedAuthServer) Ping(ctx context.Context, req *types.Empty) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_SigninTwitch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SigninTwitch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Auth/SigninTwitch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SigninTwitch(ctx, req.(*types.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SigninGoogle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SigninGoogle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Auth/SigninGoogle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SigninGoogle(ctx, req.(*types.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Auth/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Ping(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SigninTwitch",
			Handler:    _Auth_SigninTwitch_Handler,
		},
		{
			MethodName: "SigninGoogle",
			Handler:    _Auth_SigninGoogle_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Auth_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/elojah/game_03/cmd/auth/grpc/auth.proto",
}
