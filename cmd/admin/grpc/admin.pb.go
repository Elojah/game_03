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
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbd, 0x6e, 0x14, 0x31,
	0x14, 0x85, 0xc7, 0x22, 0x20, 0x61, 0xc4, 0x8f, 0x2c, 0x11, 0xa4, 0x01, 0xdd, 0x37, 0xc0, 0x8e,
	0x58, 0x04, 0x88, 0x02, 0x69, 0x13, 0x28, 0x91, 0x10, 0xbf, 0x25, 0xf2, 0xee, 0x5c, 0x1c, 0xc3,
	0xcc, 0xd8, 0x78, 0xef, 0x0a, 0xa5, 0xa3, 0xa3, 0xe5, 0x31, 0x78, 0x04, 0xca, 0x94, 0x94, 0x5b,
	0xa6, 0x64, 0x3d, 0x0d, 0x65, 0x4a, 0x4a, 0x34, 0x33, 0x4e, 0xc8, 0x42, 0x26, 0xda, 0xce, 0x3e,
	0xe7, 0x1e, 0xdf, 0xe3, 0x8f, 0xdf, 0x35, 0x96, 0x76, 0xe7, 0x13, 0x39, 0x75, 0x95, 0xc2, 0xd2,
	0xbd, 0xd7, 0xbb, 0xca, 0xe8, 0x0a, 0xdf, 0x6e, 0x8d, 0xd4, 0xb4, 0x2a, 0x94, 0x2e, 0x2a, 0x5b,
	0x2b, 0x13, 0xfc, 0xb4, 0x3f, 0x4a, 0x1f, 0x1c, 0x39, 0xb1, 0xd1, 0x2a, 0xf9, 0xed, 0x13, 0x59,
	0xe3, 0x8c, 0x53, 0x9d, 0x39, 0x99, 0xbf, 0xeb, 0x6e, 0xdd, 0xa5, 0x3b, 0xf5, 0xa1, 0xfc, 0xa6,
	0x71, 0xce, 0x94, 0xf8, 0x77, 0x0a, 0x2b, 0x4f, 0x7b, 0xc9, 0x84, 0x7f, 0xcd, 0x4f, 0x41, 0x7b,
	0x8f, 0x61, 0x96, 0xfc, 0xad, 0xe1, 0x9e, 0xfe, 0x83, 0x51, 0x64, 0x4b, 0x54, 0x05, 0x39, 0x35,
	0x43, 0x4a, 0x89, 0xd1, 0xd9, 0x09, 0xac, 0xc9, 0xd2, 0x9e, 0x22, 0xac, 0x7c, 0xa9, 0x09, 0x53,
	0xe8, 0xc1, 0x5a, 0xa1, 0x76, 0x91, 0xae, 0x6d, 0xa5, 0xc9, 0xba, 0x84, 0x24, 0xbf, 0xbf, 0x76,
	0x72, 0x75, 0xe5, 0x9d, 0x2f, 0xe7, 0xf8, 0xf9, 0x71, 0xcb, 0x56, 0x8c, 0xf9, 0xc5, 0xa7, 0xd6,
	0x04, 0x4d, 0xf8, 0xca, 0x8b, 0x5b, 0xb2, 0x27, 0x22, 0x8f, 0x88, 0xc8, 0x17, 0x14, 0x6c, 0x6d,
	0x5e, 0xeb, 0x72, 0x8e, 0xf9, 0xe6, 0x7f, 0xee, 0x93, 0x16, 0xa6, 0x78, 0xc4, 0x2f, 0xef, 0x04,
	0xd4, 0x84, 0x2f, 0x6d, 0x89, 0x33, 0x24, 0x71, 0x5d, 0x16, 0xe4, 0xe4, 0x8a, 0xf6, 0x1c, 0x3f,
	0xe6, 0x9b, 0xa7, 0xc9, 0x33, 0x2f, 0x76, 0xf8, 0xa5, 0x5e, 0x7c, 0xe3, 0x42, 0x59, 0x88, 0x81,
	0x35, 0xf9, 0x99, 0xe5, 0xc4, 0x43, 0x7e, 0x25, 0xbd, 0x9c, 0x7e, 0x2a, 0x56, 0xd6, 0x25, 0xb1,
	0xad, 0x71, 0x4d, 0xf6, 0x5c, 0xe4, 0xf1, 0xe4, 0x36, 0xbf, 0xda, 0x8f, 0x8d, 0x8f, 0xf8, 0x8a,
	0x1b, 0x27, 0xc2, 0xc7, 0x6a, 0xff, 0x89, 0xd3, 0x21, 0xdc, 0xe3, 0x1b, 0xcf, 0x6c, 0x6d, 0x06,
	0xdb, 0x0f, 0xe8, 0xdb, 0x8f, 0x17, 0x4b, 0xc8, 0x0e, 0x96, 0x90, 0x1d, 0x2e, 0x81, 0xfd, 0x5e,
	0x02, 0xfb, 0x1c, 0x81, 0x7d, 0x8b, 0xc0, 0xbe, 0x47, 0x60, 0xfb, 0x11, 0xd8, 0x8f, 0x08, 0x6c,
	0x11, 0x81, 0xfd, 0x8c, 0xc0, 0x7e, 0x45, 0xc8, 0x0e, 0x23, 0xb0, 0xaf, 0x0d, 0x64, 0xfb, 0x0d,
	0xb0, 0x45, 0x03, 0xd9, 0x41, 0x03, 0xd9, 0xe4, 0x42, 0xf7, 0xea, 0xe8, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xde, 0xa0, 0xfb, 0xec, 0x5a, 0x03, 0x00, 0x00,
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
	// DB migrations
	MigrateUp(ctx context.Context, in *types.StringValue, opts ...grpc.CallOption) (*types.Empty, error)
	// Map creation
	CreateTileset(ctx context.Context, in *dto.CreateTilesetReq, opts ...grpc.CallOption) (*dto.CreateTilesetResp, error)
	CreateWorld(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.StringValue, error)
	// Entity
	CreateTemplate(ctx context.Context, in *dto1.CreateTemplateReq, opts ...grpc.CallOption) (*entity.Template, error)
	CreateAnimation(ctx context.Context, in *dto1.CreateAnimationReq, opts ...grpc.CallOption) (*types.Empty, error)
	// Ping
	Ping(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
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

func (c *adminClient) Ping(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Admin/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
type AdminServer interface {
	// DB migrations
	MigrateUp(context.Context, *types.StringValue) (*types.Empty, error)
	// Map creation
	CreateTileset(context.Context, *dto.CreateTilesetReq) (*dto.CreateTilesetResp, error)
	CreateWorld(context.Context, *types.Empty) (*types.StringValue, error)
	// Entity
	CreateTemplate(context.Context, *dto1.CreateTemplateReq) (*entity.Template, error)
	CreateAnimation(context.Context, *dto1.CreateAnimationReq) (*types.Empty, error)
	// Ping
	Ping(context.Context, *types.Empty) (*types.Empty, error)
}

// UnimplementedAdminServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (*UnimplementedAdminServer) MigrateUp(ctx context.Context, req *types.StringValue) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MigrateUp not implemented")
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
func (*UnimplementedAdminServer) Ping(ctx context.Context, req *types.Empty) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
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

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MigrateUp",
			Handler:    _Admin_MigrateUp_Handler,
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
		{
			MethodName: "Ping",
			Handler:    _Admin_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/elojah/game_03/cmd/admin/grpc/admin.proto",
}
