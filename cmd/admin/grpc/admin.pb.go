// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/cmd/admin/grpc/admin.proto

package grpc

import (
	context "context"
	fmt "fmt"
	entity "github.com/elojah/game_03/pkg/entity"
	dto1 "github.com/elojah/game_03/pkg/entity/dto"
	dto "github.com/elojah/game_03/pkg/tile/dto"
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
	proto.RegisterFile("github.com/elojah/game_03/cmd/admin/grpc/admin.proto", fileDescriptor_c466ab583a6679eb)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/cmd/admin/grpc/admin.proto", fileDescriptor_c466ab583a6679eb)
}

var fileDescriptor_c466ab583a6679eb = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x6a, 0x14, 0x41,
	0x10, 0xc6, 0xa7, 0x31, 0x0a, 0xb6, 0xa8, 0xa1, 0xc1, 0x08, 0xa3, 0xd4, 0x1b, 0xd8, 0x1d, 0x5c,
	0x51, 0xf1, 0x20, 0xec, 0x2e, 0x82, 0x20, 0x82, 0xc4, 0x7f, 0x47, 0xe9, 0xdd, 0x29, 0x3b, 0x6d,
	0x66, 0xa6, 0xdb, 0x99, 0x5a, 0x64, 0x6f, 0x3e, 0x82, 0x8f, 0xe1, 0x23, 0x78, 0xcc, 0x31, 0xe0,
	0x65, 0x8f, 0x39, 0x3a, 0xbd, 0x17, 0x8f, 0x39, 0x7a, 0x94, 0xf9, 0x93, 0x98, 0xd5, 0x4c, 0x58,
	0x6f, 0xdd, 0xdf, 0x57, 0xbf, 0xaa, 0xea, 0xaf, 0xf9, 0x3d, 0x63, 0x69, 0x77, 0x36, 0x91, 0x53,
	0x97, 0x29, 0x4c, 0xdd, 0x07, 0xbd, 0xab, 0x8c, 0xce, 0xf0, 0xdd, 0xf6, 0x40, 0x4d, 0xb3, 0x44,
	0xe9, 0x24, 0xb3, 0xb9, 0x32, 0x85, 0x9f, 0xb6, 0x47, 0xe9, 0x0b, 0x47, 0x4e, 0x6c, 0xd4, 0x4a,
	0x7c, 0xe7, 0x14, 0x6b, 0x9c, 0x71, 0xaa, 0x31, 0x27, 0xb3, 0xf7, 0xcd, 0xad, 0xb9, 0x34, 0xa7,
	0x16, 0x8a, 0x6f, 0x19, 0xe7, 0x4c, 0x8a, 0x7f, 0xaa, 0x30, 0xf3, 0x34, 0xef, 0x4c, 0xf8, 0xdb,
	0xfc, 0x54, 0x68, 0xef, 0xb1, 0x28, 0x3b, 0x7f, 0xbb, 0x7f, 0x4f, 0xbf, 0x67, 0x14, 0xd9, 0x14,
	0x55, 0x42, 0x4e, 0x95, 0x48, 0x1d, 0x31, 0x38, 0x9f, 0xc0, 0x9c, 0x2c, 0xcd, 0x15, 0x61, 0xe6,
	0x53, 0x4d, 0xd8, 0x41, 0x0f, 0xd7, 0x82, 0xea, 0x41, 0x3a, 0xb7, 0x99, 0x26, 0xeb, 0xba, 0x48,
	0xe2, 0x07, 0x6b, 0x93, 0xab, 0x23, 0xef, 0x7e, 0xbf, 0xc0, 0x2f, 0x0e, 0xeb, 0x6c, 0xc5, 0x90,
	0x5f, 0x7e, 0x6e, 0x4d, 0xa1, 0x09, 0x5f, 0x7b, 0x71, 0x5b, 0xb6, 0x89, 0xc8, 0xe3, 0x44, 0xe4,
	0x4b, 0x2a, 0x6c, 0x6e, 0xde, 0xe8, 0x74, 0x86, 0xf1, 0xd6, 0x3f, 0xee, 0x93, 0x3a, 0x4c, 0x31,
	0xe2, 0x9b, 0x3b, 0x8e, 0x34, 0xe1, 0xd8, 0xb9, 0x3d, 0x8b, 0xcf, 0x70, 0x5e, 0x8a, 0x9e, 0xda,
	0xde, 0x1e, 0xf7, 0xf9, 0xc6, 0x0b, 0x9b, 0x9b, 0xff, 0xe6, 0x1e, 0xf3, 0xab, 0xe3, 0x02, 0x35,
	0xe1, 0x2b, 0x9b, 0x62, 0x89, 0x24, 0x6e, 0xc8, 0x84, 0x9c, 0x5c, 0xd1, 0x76, 0xf0, 0x63, 0xbc,
	0x75, 0x96, 0x5c, 0x7a, 0x31, 0xe6, 0x57, 0x5a, 0xf1, 0xad, 0x2b, 0xd2, 0xa4, 0x77, 0xfc, 0xb9,
	0xc1, 0x88, 0x47, 0xfc, 0x5a, 0xd7, 0xb9, 0x4b, 0x59, 0xac, 0x8c, 0xeb, 0xc4, 0x7a, 0x8d, 0x4d,
	0xd9, 0xfe, 0x89, 0x3c, 0xa9, 0x1c, 0xf1, 0xeb, 0x6d, 0xd9, 0xf0, 0xf8, 0x6f, 0xc5, 0xcd, 0x53,
	0xf0, 0x89, 0xda, 0x3e, 0xe2, 0xcc, 0xed, 0x46, 0x4f, 0x0f, 0x2a, 0x88, 0x16, 0x15, 0x44, 0x87,
	0x15, 0x44, 0x47, 0x15, 0xb0, 0x5f, 0x15, 0xb0, 0xcf, 0x01, 0xd8, 0xd7, 0x00, 0xec, 0x5b, 0x00,
	0xb6, 0x1f, 0x80, 0x1d, 0x04, 0x60, 0x8b, 0x00, 0xec, 0x47, 0x00, 0xf6, 0x33, 0x40, 0x74, 0x14,
	0x80, 0x7d, 0x59, 0x42, 0xb4, 0xbf, 0x04, 0xb6, 0x58, 0x42, 0x74, 0xb8, 0x84, 0x68, 0x72, 0xa9,
	0xe9, 0x3c, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xec, 0xfe, 0xe2, 0xa2, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminClient interface {
	// ### TECHNICAL ADMIN ###
	// DB migrations
	MigrateUp(ctx context.Context, in *types.StringValue, opts ...grpc.CallOption) (*types.Empty, error)
	// Cookie secure management
	RotateCookieKeys(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
	// Ping
	Ping(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
	// Map creation
	CreateTileset(ctx context.Context, in *dto.CreateTilesetReq, opts ...grpc.CallOption) (*dto.CreateTilesetResp, error)
	CreateWorld(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.StringValue, error)
	// Entity
	CreateTemplate(ctx context.Context, in *dto1.CreateTemplateReq, opts ...grpc.CallOption) (*entity.Template, error)
	CreateAnimation(ctx context.Context, in *dto1.CreateAnimationReq, opts ...grpc.CallOption) (*types.Empty, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) MigrateUp(ctx context.Context, in *types.StringValue, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Admin/MigrateUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) RotateCookieKeys(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Admin/RotateCookieKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Ping(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Admin/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CreateTileset(ctx context.Context, in *dto.CreateTilesetReq, opts ...grpc.CallOption) (*dto.CreateTilesetResp, error) {
	out := new(dto.CreateTilesetResp)
	err := c.cc.Invoke(ctx, "/grpc.Admin/CreateTileset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CreateWorld(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.StringValue, error) {
	out := new(types.StringValue)
	err := c.cc.Invoke(ctx, "/grpc.Admin/CreateWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CreateTemplate(ctx context.Context, in *dto1.CreateTemplateReq, opts ...grpc.CallOption) (*entity.Template, error) {
	out := new(entity.Template)
	err := c.cc.Invoke(ctx, "/grpc.Admin/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CreateAnimation(ctx context.Context, in *dto1.CreateAnimationReq, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Admin/CreateAnimation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
type AdminServer interface {
	// ### TECHNICAL ADMIN ###
	// DB migrations
	MigrateUp(context.Context, *types.StringValue) (*types.Empty, error)
	// Cookie secure management
	RotateCookieKeys(context.Context, *types.Empty) (*types.Empty, error)
	// Ping
	Ping(context.Context, *types.Empty) (*types.Empty, error)
	// Map creation
	CreateTileset(context.Context, *dto.CreateTilesetReq) (*dto.CreateTilesetResp, error)
	CreateWorld(context.Context, *types.Empty) (*types.StringValue, error)
	// Entity
	CreateTemplate(context.Context, *dto1.CreateTemplateReq) (*entity.Template, error)
	CreateAnimation(context.Context, *dto1.CreateAnimationReq) (*types.Empty, error)
}

// UnimplementedAdminServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (*UnimplementedAdminServer) MigrateUp(ctx context.Context, req *types.StringValue) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MigrateUp not implemented")
}
func (*UnimplementedAdminServer) RotateCookieKeys(ctx context.Context, req *types.Empty) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateCookieKeys not implemented")
}
func (*UnimplementedAdminServer) Ping(ctx context.Context, req *types.Empty) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedAdminServer) CreateTileset(ctx context.Context, req *dto.CreateTilesetReq) (*dto.CreateTilesetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTileset not implemented")
}
func (*UnimplementedAdminServer) CreateWorld(ctx context.Context, req *types.Empty) (*types.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorld not implemented")
}
func (*UnimplementedAdminServer) CreateTemplate(ctx context.Context, req *dto1.CreateTemplateReq) (*entity.Template, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}
func (*UnimplementedAdminServer) CreateAnimation(ctx context.Context, req *dto1.CreateAnimationReq) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnimation not implemented")
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_MigrateUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).MigrateUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Admin/MigrateUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).MigrateUp(ctx, req.(*types.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_RotateCookieKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).RotateCookieKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Admin/RotateCookieKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).RotateCookieKeys(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Admin/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Ping(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CreateTileset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto.CreateTilesetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateTileset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Admin/CreateTileset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateTileset(ctx, req.(*dto.CreateTilesetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CreateWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Admin/CreateWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateWorld(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto1.CreateTemplateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Admin/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateTemplate(ctx, req.(*dto1.CreateTemplateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CreateAnimation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto1.CreateAnimationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateAnimation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Admin/CreateAnimation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateAnimation(ctx, req.(*dto1.CreateAnimationReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MigrateUp",
			Handler:    _Admin_MigrateUp_Handler,
		},
		{
			MethodName: "RotateCookieKeys",
			Handler:    _Admin_RotateCookieKeys_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Admin_Ping_Handler,
		},
		{
			MethodName: "CreateTileset",
			Handler:    _Admin_CreateTileset_Handler,
		},
		{
			MethodName: "CreateWorld",
			Handler:    _Admin_CreateWorld_Handler,
		},
		{
			MethodName: "CreateTemplate",
			Handler:    _Admin_CreateTemplate_Handler,
		},
		{
			MethodName: "CreateAnimation",
			Handler:    _Admin_CreateAnimation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/elojah/game_03/cmd/admin/grpc/admin.proto",
}
