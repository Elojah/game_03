// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/cmd/api/grpc/api.proto

package grpc

import (
	context "context"
	fmt "fmt"
	entity "github.com/elojah/game_03/pkg/entity"
	dto "github.com/elojah/game_03/pkg/entity/dto"
	room "github.com/elojah/game_03/pkg/room"
	dto2 "github.com/elojah/game_03/pkg/room/dto"
	dto3 "github.com/elojah/game_03/pkg/twitch/dto"
	user "github.com/elojah/game_03/pkg/user"
	dto1 "github.com/elojah/game_03/pkg/user/dto"
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
	proto.RegisterFile("github.com/elojah/game_03/cmd/api/grpc/api.proto", fileDescriptor_6dc96cba8df2a9f1)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/cmd/api/grpc/api.proto", fileDescriptor_6dc96cba8df2a9f1)
}

var fileDescriptor_6dc96cba8df2a9f1 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0x63, 0x51, 0x4a, 0x6b, 0xb5, 0x50, 0x1e, 0xb4, 0x43, 0x90, 0x3c, 0x31, 0x94, 0xa1,
	0xce, 0xd1, 0x52, 0x40, 0x42, 0x42, 0x2a, 0x51, 0x11, 0x48, 0x0c, 0xd1, 0x01, 0x62, 0x44, 0x69,
	0xce, 0x4d, 0x03, 0x49, 0xec, 0xc6, 0x3e, 0x9d, 0xba, 0xf1, 0x11, 0xf8, 0x18, 0x7c, 0x04, 0xc6,
	0x8e, 0x1d, 0x6f, 0xec, 0xc8, 0xe5, 0x16, 0xc6, 0x2e, 0x48, 0x8c, 0xc8, 0x4e, 0x72, 0x97, 0xb4,
	0x77, 0x28, 0x5d, 0xee, 0xfc, 0xfe, 0x7e, 0x3f, 0xbf, 0x7f, 0xec, 0xf7, 0x70, 0x27, 0x8c, 0xd4,
	0x51, 0xff, 0x80, 0x06, 0x3c, 0x71, 0x58, 0xcc, 0xbf, 0xf8, 0x47, 0x4e, 0xe8, 0x27, 0xec, 0x73,
	0x67, 0xc7, 0x09, 0x92, 0x9e, 0xe3, 0x8b, 0xc8, 0x09, 0x33, 0x11, 0xe8, 0x05, 0x15, 0x19, 0x57,
	0x1c, 0x16, 0x74, 0x6c, 0x6f, 0xd5, 0xb8, 0x90, 0x87, 0xdc, 0x31, 0x9b, 0x07, 0xfd, 0x43, 0x13,
	0x99, 0xc0, 0xac, 0x0a, 0xc8, 0x7e, 0x10, 0x72, 0x1e, 0xc6, 0x6c, 0x9a, 0xc5, 0x12, 0xa1, 0x4e,
	0xca, 0xcd, 0xc7, 0xf3, 0x3d, 0x88, 0xaf, 0xa1, 0xc3, 0x52, 0x15, 0xa9, 0x93, 0xf2, 0xaf, 0x44,
	0xb6, 0x5a, 0x21, 0x22, 0x28, 0xd3, 0x77, 0x5b, 0xa5, 0xf7, 0x14, 0x6f, 0x56, 0x79, 0xde, 0x1a,
	0xf3, 0xd3, 0x28, 0xf1, 0x55, 0xc4, 0xd3, 0x6b, 0x7d, 0x92, 0x26, 0x27, 0x1e, 0x9f, 0xb5, 0x46,
	0x14, 0x4b, 0x44, 0xec, 0x2b, 0xd6, 0xee, 0x2e, 0x32, 0xce, 0x13, 0xf3, 0xd3, 0xce, 0x9a, 0x49,
	0xd7, 0x55, 0x02, 0x16, 0xc7, 0xd7, 0x44, 0x6a, 0x55, 0xb6, 0x5b, 0x22, 0x03, 0x9e, 0xc5, 0xbd,
	0x76, 0xaf, 0xa4, 0x06, 0x91, 0x0a, 0x8e, 0x0c, 0x75, 0xc8, 0xe3, 0x98, 0x0f, 0x4a, 0xec, 0xc9,
	0xff, 0xb1, 0xbe, 0x64, 0x99, 0x81, 0x24, 0x93, 0x72, 0xfa, 0x42, 0x9d, 0x16, 0x54, 0x83, 0xd8,
	0xfe, 0x73, 0x13, 0xdf, 0xd8, 0xf3, 0xde, 0x02, 0xc5, 0xcb, 0x2e, 0x4f, 0x53, 0x16, 0x28, 0xcf,
	0x05, 0x4c, 0xcb, 0x8e, 0xf1, 0x5c, 0xfb, 0x1e, 0xed, 0x29, 0x4e, 0xdf, 0x45, 0x52, 0xed, 0x1b,
	0xad, 0xcb, 0xa4, 0xe8, 0x20, 0xd8, 0xc1, 0x2b, 0x1f, 0x45, 0xcf, 0x57, 0xac, 0x50, 0x61, 0xb9,
	0x42, 0xf6, 0xed, 0x0d, 0x5a, 0xcc, 0x05, 0xad, 0xe6, 0x82, 0xee, 0xeb, 0xb9, 0xd8, 0x44, 0xb0,
	0x8b, 0x57, 0xdd, 0x8c, 0xf9, 0x8a, 0xbd, 0x2f, 0x3c, 0xc0, 0xba, 0x39, 0xbc, 0xa1, 0x75, 0xd9,
	0xb1, 0xbd, 0x4a, 0xb5, 0x53, 0x5a, 0x65, 0xed, 0x62, 0x3c, 0xad, 0x0f, 0x70, 0xc5, 0xd0, 0xf1,
	0x4c, 0x93, 0xf0, 0x12, 0xaf, 0x6a, 0x65, 0xaf, 0xea, 0xe2, 0xb2, 0x5a, 0x43, 0xd3, 0xf0, 0xc6,
	0x2c, 0x59, 0x0a, 0xd8, 0xc4, 0x4b, 0x85, 0x33, 0xcf, 0x85, 0xb5, 0x9a, 0x51, 0xcf, 0xd5, 0x54,
	0xed, 0x8e, 0xe0, 0x11, 0x5e, 0xd4, 0xb8, 0xe7, 0xc2, 0xed, 0xc9, 0x59, 0x45, 0xd6, 0x9d, 0x46,
	0x2c, 0x05, 0xbc, 0xc0, 0x2b, 0x3a, 0xfa, 0x50, 0x76, 0x3b, 0xdc, 0x9f, 0x24, 0x54, 0x92, 0xc6,
	0xd6, 0x67, 0xa8, 0x52, 0xc0, 0x43, 0xbc, 0x52, 0x58, 0xb8, 0x7a, 0xe9, 0xd3, 0x25, 0x10, 0x8c,
	0x8b, 0xac, 0x2e, 0xe7, 0x09, 0xdc, 0xa2, 0xa6, 0x81, 0xbb, 0x76, 0xb5, 0x00, 0x07, 0x2f, 0xe9,
	0x93, 0xcd, 0xee, 0xda, 0xa4, 0x90, 0x0e, 0x75, 0xe9, 0xbb, 0x97, 0x14, 0x29, 0x2a, 0xc0, 0x65,
	0x71, 0x5c, 0x03, 0x74, 0xd8, 0x04, 0x0a, 0x45, 0x0a, 0xd8, 0xc6, 0xcb, 0x3a, 0xfe, 0xa4, 0xc7,
	0x00, 0xa6, 0xfb, 0x26, 0xd6, 0x08, 0x5c, 0x96, 0xa4, 0xa8, 0x1e, 0xf9, 0xb5, 0x19, 0x82, 0xda,
	0x23, 0x17, 0x42, 0xf3, 0x91, 0x2b, 0x4d, 0x0a, 0x78, 0x8a, 0x17, 0xbc, 0x28, 0x0d, 0x61, 0x4e,
	0xd3, 0xcd, 0x6b, 0xc6, 0x57, 0x6f, 0xce, 0x46, 0xc4, 0x1a, 0x8e, 0x88, 0x75, 0x3e, 0x22, 0xd6,
	0xc5, 0x88, 0xa0, 0xbf, 0x23, 0x82, 0xbe, 0xe5, 0x04, 0xfd, 0xc8, 0x09, 0xfa, 0x99, 0x13, 0x74,
	0x9a, 0x13, 0x74, 0x96, 0x13, 0x34, 0xcc, 0x09, 0xfa, 0x95, 0x13, 0xf4, 0x3b, 0x27, 0xd6, 0x45,
	0x4e, 0xd0, 0xf7, 0x31, 0xb1, 0x4e, 0xc7, 0x04, 0x0d, 0xc7, 0xc4, 0x3a, 0x1f, 0x13, 0xeb, 0x60,
	0xd1, 0x9c, 0xbc, 0xf3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xbb, 0xf7, 0xc4, 0x75, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	// STREAM LOOPS
	ConnectPC(ctx context.Context, in *entity.PC, opts ...grpc.CallOption) (API_ConnectPCClient, error)
	UpdateEntity(ctx context.Context, opts ...grpc.CallOption) (API_UpdateEntityClient, error)
	// Session
	CreateSession(ctx context.Context, in *dto1.CreateSessionReq, opts ...grpc.CallOption) (*user.Session, error)
	// Entity
	ListEntity(ctx context.Context, in *dto.ListEntityReq, opts ...grpc.CallOption) (*dto.ListEntityResp, error)
	// Animation
	ListAnimation(ctx context.Context, in *dto.ListAnimationReq, opts ...grpc.CallOption) (*dto.ListAnimationResp, error)
	// PC
	CreatePC(ctx context.Context, in *dto.CreatePCReq, opts ...grpc.CallOption) (*entity.PC, error)
	ListPC(ctx context.Context, in *dto.ListPCReq, opts ...grpc.CallOption) (*dto.ListPCResp, error)
	// Template
	ListTemplate(ctx context.Context, in *dto.ListTemplateReq, opts ...grpc.CallOption) (*dto.ListTemplateResp, error)
	// Entity
	CreateEntity(ctx context.Context, in *entity.E, opts ...grpc.CallOption) (*entity.E, error)
	// Room
	CreateRoom(ctx context.Context, in *room.R, opts ...grpc.CallOption) (*room.R, error)
	ListRoom(ctx context.Context, in *dto2.ListRoomReq, opts ...grpc.CallOption) (*dto2.ListRoomResp, error)
	// Cell
	ListCell(ctx context.Context, in *dto2.ListCellReq, opts ...grpc.CallOption) (*dto2.ListCellResp, error)
	// World
	ListWorld(ctx context.Context, in *dto2.ListWorldReq, opts ...grpc.CallOption) (*dto2.ListWorldResp, error)
	// Twitch
	ListFollow(ctx context.Context, in *dto3.ListFollowReq, opts ...grpc.CallOption) (*dto3.ListFollowResp, error)
	// Ping
	Ping(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) ConnectPC(ctx context.Context, in *entity.PC, opts ...grpc.CallOption) (API_ConnectPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[0], "/grpc.API/ConnectPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIConnectPCClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_ConnectPCClient interface {
	Recv() (*dto.ListEntityResp, error)
	grpc.ClientStream
}

type aPIConnectPCClient struct {
	grpc.ClientStream
}

func (x *aPIConnectPCClient) Recv() (*dto.ListEntityResp, error) {
	m := new(dto.ListEntityResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) UpdateEntity(ctx context.Context, opts ...grpc.CallOption) (API_UpdateEntityClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[1], "/grpc.API/UpdateEntity", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIUpdateEntityClient{stream}
	return x, nil
}

type API_UpdateEntityClient interface {
	Send(*entity.E) error
	CloseAndRecv() (*types.Empty, error)
	grpc.ClientStream
}

type aPIUpdateEntityClient struct {
	grpc.ClientStream
}

func (x *aPIUpdateEntityClient) Send(m *entity.E) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPIUpdateEntityClient) CloseAndRecv() (*types.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(types.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) CreateSession(ctx context.Context, in *dto1.CreateSessionReq, opts ...grpc.CallOption) (*user.Session, error) {
	out := new(user.Session)
	err := c.cc.Invoke(ctx, "/grpc.API/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListEntity(ctx context.Context, in *dto.ListEntityReq, opts ...grpc.CallOption) (*dto.ListEntityResp, error) {
	out := new(dto.ListEntityResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListAnimation(ctx context.Context, in *dto.ListAnimationReq, opts ...grpc.CallOption) (*dto.ListAnimationResp, error) {
	out := new(dto.ListAnimationResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListAnimation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreatePC(ctx context.Context, in *dto.CreatePCReq, opts ...grpc.CallOption) (*entity.PC, error) {
	out := new(entity.PC)
	err := c.cc.Invoke(ctx, "/grpc.API/CreatePC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListPC(ctx context.Context, in *dto.ListPCReq, opts ...grpc.CallOption) (*dto.ListPCResp, error) {
	out := new(dto.ListPCResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListTemplate(ctx context.Context, in *dto.ListTemplateReq, opts ...grpc.CallOption) (*dto.ListTemplateResp, error) {
	out := new(dto.ListTemplateResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateEntity(ctx context.Context, in *entity.E, opts ...grpc.CallOption) (*entity.E, error) {
	out := new(entity.E)
	err := c.cc.Invoke(ctx, "/grpc.API/CreateEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateRoom(ctx context.Context, in *room.R, opts ...grpc.CallOption) (*room.R, error) {
	out := new(room.R)
	err := c.cc.Invoke(ctx, "/grpc.API/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListRoom(ctx context.Context, in *dto2.ListRoomReq, opts ...grpc.CallOption) (*dto2.ListRoomResp, error) {
	out := new(dto2.ListRoomResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListCell(ctx context.Context, in *dto2.ListCellReq, opts ...grpc.CallOption) (*dto2.ListCellResp, error) {
	out := new(dto2.ListCellResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListCell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListWorld(ctx context.Context, in *dto2.ListWorldReq, opts ...grpc.CallOption) (*dto2.ListWorldResp, error) {
	out := new(dto2.ListWorldResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListFollow(ctx context.Context, in *dto3.ListFollowReq, opts ...grpc.CallOption) (*dto3.ListFollowResp, error) {
	out := new(dto3.ListFollowResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Ping(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/grpc.API/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	// STREAM LOOPS
	ConnectPC(*entity.PC, API_ConnectPCServer) error
	UpdateEntity(API_UpdateEntityServer) error
	// Session
	CreateSession(context.Context, *dto1.CreateSessionReq) (*user.Session, error)
	// Entity
	ListEntity(context.Context, *dto.ListEntityReq) (*dto.ListEntityResp, error)
	// Animation
	ListAnimation(context.Context, *dto.ListAnimationReq) (*dto.ListAnimationResp, error)
	// PC
	CreatePC(context.Context, *dto.CreatePCReq) (*entity.PC, error)
	ListPC(context.Context, *dto.ListPCReq) (*dto.ListPCResp, error)
	// Template
	ListTemplate(context.Context, *dto.ListTemplateReq) (*dto.ListTemplateResp, error)
	// Entity
	CreateEntity(context.Context, *entity.E) (*entity.E, error)
	// Room
	CreateRoom(context.Context, *room.R) (*room.R, error)
	ListRoom(context.Context, *dto2.ListRoomReq) (*dto2.ListRoomResp, error)
	// Cell
	ListCell(context.Context, *dto2.ListCellReq) (*dto2.ListCellResp, error)
	// World
	ListWorld(context.Context, *dto2.ListWorldReq) (*dto2.ListWorldResp, error)
	// Twitch
	ListFollow(context.Context, *dto3.ListFollowReq) (*dto3.ListFollowResp, error)
	// Ping
	Ping(context.Context, *types.Empty) (*types.Empty, error)
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) ConnectPC(req *entity.PC, srv API_ConnectPCServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectPC not implemented")
}
func (*UnimplementedAPIServer) UpdateEntity(srv API_UpdateEntityServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateEntity not implemented")
}
func (*UnimplementedAPIServer) CreateSession(ctx context.Context, req *dto1.CreateSessionReq) (*user.Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (*UnimplementedAPIServer) ListEntity(ctx context.Context, req *dto.ListEntityReq) (*dto.ListEntityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntity not implemented")
}
func (*UnimplementedAPIServer) ListAnimation(ctx context.Context, req *dto.ListAnimationReq) (*dto.ListAnimationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnimation not implemented")
}
func (*UnimplementedAPIServer) CreatePC(ctx context.Context, req *dto.CreatePCReq) (*entity.PC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePC not implemented")
}
func (*UnimplementedAPIServer) ListPC(ctx context.Context, req *dto.ListPCReq) (*dto.ListPCResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPC not implemented")
}
func (*UnimplementedAPIServer) ListTemplate(ctx context.Context, req *dto.ListTemplateReq) (*dto.ListTemplateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTemplate not implemented")
}
func (*UnimplementedAPIServer) CreateEntity(ctx context.Context, req *entity.E) (*entity.E, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntity not implemented")
}
func (*UnimplementedAPIServer) CreateRoom(ctx context.Context, req *room.R) (*room.R, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (*UnimplementedAPIServer) ListRoom(ctx context.Context, req *dto2.ListRoomReq) (*dto2.ListRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoom not implemented")
}
func (*UnimplementedAPIServer) ListCell(ctx context.Context, req *dto2.ListCellReq) (*dto2.ListCellResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCell not implemented")
}
func (*UnimplementedAPIServer) ListWorld(ctx context.Context, req *dto2.ListWorldReq) (*dto2.ListWorldResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorld not implemented")
}
func (*UnimplementedAPIServer) ListFollow(ctx context.Context, req *dto3.ListFollowReq) (*dto3.ListFollowResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFollow not implemented")
}
func (*UnimplementedAPIServer) Ping(ctx context.Context, req *types.Empty) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_ConnectPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(entity.PC)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).ConnectPC(m, &aPIConnectPCServer{stream})
}

type API_ConnectPCServer interface {
	Send(*dto.ListEntityResp) error
	grpc.ServerStream
}

type aPIConnectPCServer struct {
	grpc.ServerStream
}

func (x *aPIConnectPCServer) Send(m *dto.ListEntityResp) error {
	return x.ServerStream.SendMsg(m)
}

func _API_UpdateEntity_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APIServer).UpdateEntity(&aPIUpdateEntityServer{stream})
}

type API_UpdateEntityServer interface {
	SendAndClose(*types.Empty) error
	Recv() (*entity.E, error)
	grpc.ServerStream
}

type aPIUpdateEntityServer struct {
	grpc.ServerStream
}

func (x *aPIUpdateEntityServer) SendAndClose(m *types.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPIUpdateEntityServer) Recv() (*entity.E, error) {
	m := new(entity.E)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _API_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto1.CreateSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateSession(ctx, req.(*dto1.CreateSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto.ListEntityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/ListEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListEntity(ctx, req.(*dto.ListEntityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListAnimation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto.ListAnimationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListAnimation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/ListAnimation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListAnimation(ctx, req.(*dto.ListAnimationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreatePC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto.CreatePCReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreatePC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/CreatePC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreatePC(ctx, req.(*dto.CreatePCReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto.ListPCReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/ListPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListPC(ctx, req.(*dto.ListPCReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto.ListTemplateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/ListTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListTemplate(ctx, req.(*dto.ListTemplateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(entity.E)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/CreateEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateEntity(ctx, req.(*entity.E))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(room.R)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateRoom(ctx, req.(*room.R))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto2.ListRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/ListRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListRoom(ctx, req.(*dto2.ListRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListCell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto2.ListCellReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListCell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/ListCell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListCell(ctx, req.(*dto2.ListCellReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto2.ListWorldReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/ListWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListWorld(ctx, req.(*dto2.ListWorldReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto3.ListFollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/ListFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListFollow(ctx, req.(*dto3.ListFollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Ping(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _API_CreateSession_Handler,
		},
		{
			MethodName: "ListEntity",
			Handler:    _API_ListEntity_Handler,
		},
		{
			MethodName: "ListAnimation",
			Handler:    _API_ListAnimation_Handler,
		},
		{
			MethodName: "CreatePC",
			Handler:    _API_CreatePC_Handler,
		},
		{
			MethodName: "ListPC",
			Handler:    _API_ListPC_Handler,
		},
		{
			MethodName: "ListTemplate",
			Handler:    _API_ListTemplate_Handler,
		},
		{
			MethodName: "CreateEntity",
			Handler:    _API_CreateEntity_Handler,
		},
		{
			MethodName: "CreateRoom",
			Handler:    _API_CreateRoom_Handler,
		},
		{
			MethodName: "ListRoom",
			Handler:    _API_ListRoom_Handler,
		},
		{
			MethodName: "ListCell",
			Handler:    _API_ListCell_Handler,
		},
		{
			MethodName: "ListWorld",
			Handler:    _API_ListWorld_Handler,
		},
		{
			MethodName: "ListFollow",
			Handler:    _API_ListFollow_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _API_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectPC",
			Handler:       _API_ConnectPC_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateEntity",
			Handler:       _API_UpdateEntity_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/elojah/game_03/cmd/api/grpc/api.proto",
}
