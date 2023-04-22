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
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xbd, 0x6e, 0xd4, 0x4c,
	0x14, 0xf5, 0xe8, 0xcb, 0x97, 0x9f, 0xd1, 0x26, 0x84, 0x0b, 0x49, 0xb1, 0x48, 0x43, 0x43, 0x11,
	0x8a, 0xd8, 0x4b, 0x7e, 0x00, 0x09, 0x09, 0x29, 0x98, 0xf0, 0x23, 0x51, 0x58, 0x9b, 0x44, 0x94,
	0xc8, 0xeb, 0x9d, 0x38, 0x06, 0xdb, 0x33, 0xf1, 0xcc, 0x2a, 0x4a, 0xc7, 0x23, 0xf0, 0x18, 0x3c,
	0x02, 0x15, 0x4a, 0x99, 0x32, 0x65, 0x4a, 0xd6, 0xdb, 0x50, 0xa6, 0xa4, 0x44, 0x33, 0xb6, 0xd7,
	0x76, 0x92, 0x45, 0xde, 0x66, 0x77, 0xee, 0x99, 0x73, 0xe6, 0x9e, 0x19, 0xdf, 0x83, 0x3b, 0x7e,
	0x20, 0x8f, 0x06, 0x3d, 0xd3, 0x63, 0x91, 0x45, 0x43, 0xf6, 0xd9, 0x3d, 0xb2, 0x7c, 0x37, 0xa2,
	0x9f, 0x3a, 0x9b, 0x96, 0x17, 0xf5, 0x2d, 0x97, 0x07, 0x96, 0x9f, 0x70, 0x4f, 0x2d, 0x4c, 0x9e,
	0x30, 0xc9, 0x60, 0x46, 0xd5, 0xed, 0xf5, 0x8a, 0xce, 0x67, 0x3e, 0xb3, 0xf4, 0x66, 0x6f, 0x70,
	0xa8, 0x2b, 0x5d, 0xe8, 0x55, 0x26, 0x6a, 0x3f, 0xf0, 0x19, 0xf3, 0x43, 0x5a, 0xb2, 0x68, 0xc4,
	0xe5, 0x69, 0xbe, 0xf9, 0x64, 0xb2, 0x07, 0xfe, 0xc5, 0xb7, 0x68, 0x2c, 0x03, 0x79, 0x9a, 0xff,
	0xe5, 0x92, 0xf5, 0x46, 0x12, 0xee, 0xe5, 0xf4, 0xed, 0x46, 0xf4, 0xbe, 0x64, 0xf5, 0x2e, 0xcf,
	0x1b, 0xcb, 0xdc, 0x38, 0x88, 0x5c, 0x19, 0xb0, 0x78, 0xaa, 0x2b, 0x29, 0xe5, 0xd8, 0xe3, 0xb3,
	0xc6, 0x12, 0x49, 0x23, 0x1e, 0xba, 0x92, 0x36, 0x7b, 0x8b, 0x84, 0xb1, 0x48, 0xff, 0x4c, 0x41,
	0x1f, 0x08, 0x9a, 0x34, 0xbb, 0x89, 0xa6, 0x2b, 0x53, 0x1e, 0x0d, 0xc3, 0x29, 0x25, 0x15, 0x53,
	0x4d, 0x25, 0x15, 0x63, 0x1b, 0x0d, 0x25, 0x27, 0x2c, 0x09, 0xfb, 0xcd, 0xe6, 0x40, 0x9e, 0x04,
	0xd2, 0x3b, 0xd2, 0xaa, 0x43, 0x16, 0x86, 0xec, 0x24, 0x97, 0x6d, 0xfd, 0x5b, 0xa6, 0x4c, 0x69,
	0x91, 0xa0, 0x42, 0x8c, 0x67, 0x60, 0xe3, 0xe7, 0x1c, 0xfe, 0x6f, 0xc7, 0x79, 0x0f, 0x26, 0x5e,
	0xb0, 0x59, 0x1c, 0x53, 0x4f, 0x3a, 0x36, 0x60, 0x33, 0x9f, 0x30, 0xc7, 0x6e, 0xdf, 0x33, 0xfb,
	0x92, 0x99, 0x1f, 0x02, 0x21, 0x77, 0x35, 0xd6, 0xa5, 0x82, 0x77, 0x10, 0x6c, 0xe2, 0xd6, 0x01,
	0xef, 0xbb, 0x92, 0x66, 0x28, 0x2c, 0x14, 0x92, 0xdd, 0xf6, 0xaa, 0x99, 0xe5, 0xc8, 0x2c, 0x72,
	0x64, 0xee, 0xaa, 0x1c, 0xad, 0x21, 0xe8, 0xe0, 0x56, 0x77, 0xdf, 0x2e, 0xfb, 0xcc, 0xeb, 0xb3,
	0xf7, 0x5e, 0x3b, 0x93, 0x34, 0xf0, 0x12, 0x2f, 0xda, 0x09, 0x75, 0x25, 0xdd, 0xcb, 0x5c, 0xc3,
	0x8a, 0x96, 0xd4, 0xb0, 0x2e, 0x3d, 0x6e, 0xaf, 0xde, 0x06, 0x0b, 0x0e, 0xdb, 0x18, 0x97, 0xd6,
	0x01, 0x6e, 0xdc, 0xe5, 0xf8, 0xd6, 0xfb, 0xa9, 0xb6, 0x0a, 0xd9, 0x29, 0x02, 0x93, 0xb7, 0xad,
	0x61, 0x65, 0xdb, 0x6b, 0xb0, 0xe0, 0xb0, 0x86, 0xe7, 0x33, 0x2f, 0x8e, 0x0d, 0xcb, 0x15, 0x6b,
	0x8e, 0xad, 0x54, 0x95, 0xe7, 0x85, 0xc7, 0x78, 0x56, 0xc9, 0x1d, 0x1b, 0x96, 0xc6, 0x67, 0x65,
	0xac, 0x3b, 0xb5, 0x5a, 0x70, 0x78, 0x88, 0xff, 0x7f, 0x4b, 0x15, 0x73, 0x51, 0xef, 0xe8, 0xb5,
	0x22, 0xce, 0xe9, 0xd2, 0xb1, 0xe1, 0x05, 0x6e, 0x29, 0xfa, 0x7e, 0x9e, 0x3c, 0xb8, 0x3f, 0x3e,
	0xa1, 0x80, 0x14, 0x7d, 0xe5, 0x16, 0x54, 0x70, 0x78, 0x84, 0x5b, 0x99, 0xc7, 0x9b, 0x1f, 0xb4,
	0x5c, 0x02, 0xc1, 0x38, 0x63, 0x75, 0x19, 0x8b, 0x60, 0xce, 0xd4, 0xe9, 0xe8, 0xb6, 0x8b, 0x05,
	0x58, 0x78, 0x5e, 0x9d, 0xac, 0x77, 0x97, 0xc7, 0x8d, 0x54, 0xa9, 0x5a, 0xdf, 0xbd, 0x86, 0xe8,
	0x0f, 0xb4, 0x54, 0xd4, 0xce, 0xa0, 0x17, 0x06, 0x5e, 0x33, 0xd9, 0x16, 0x5e, 0x2a, 0x7d, 0x1c,
	0x08, 0x9a, 0x40, 0x75, 0x02, 0x0a, 0x30, 0x7b, 0x6c, 0x6d, 0x4d, 0x73, 0x72, 0x77, 0x36, 0x0d,
	0xc3, 0x4a, 0x1b, 0x55, 0xd6, 0xdb, 0x64, 0x88, 0xe0, 0xb0, 0x81, 0x17, 0x54, 0xfd, 0x51, 0xa5,
	0x13, 0xca, 0x7d, 0x5d, 0x2b, 0x09, 0x5c, 0x87, 0xca, 0x91, 0x7b, 0xa3, 0xb3, 0x59, 0x19, 0xb9,
	0x0c, 0xa8, 0x8f, 0x5c, 0x81, 0x09, 0x0e, 0x4f, 0xf1, 0x8c, 0x13, 0xc4, 0x3e, 0x4c, 0x48, 0xc2,
	0xa4, 0x84, 0xbc, 0x7a, 0x77, 0x3e, 0x24, 0xc6, 0xc5, 0x90, 0x18, 0x97, 0x43, 0x62, 0x5c, 0x0d,
	0x09, 0xfa, 0x33, 0x24, 0xe8, 0x6b, 0x4a, 0xd0, 0xf7, 0x94, 0xa0, 0x1f, 0x29, 0x41, 0x67, 0x29,
	0x41, 0xe7, 0x29, 0x41, 0x17, 0x29, 0x41, 0xbf, 0x52, 0x82, 0x7e, 0xa7, 0xc4, 0xb8, 0x4a, 0x09,
	0xfa, 0x36, 0x22, 0xc6, 0xd9, 0x88, 0xa0, 0x8b, 0x11, 0x31, 0x2e, 0x47, 0xc4, 0xe8, 0xcd, 0xea,
	0x93, 0x37, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xf5, 0x60, 0x7a, 0x6e, 0x07, 0x00, 0x00,
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
	// WIP
	RTCConnectPC(ctx context.Context, in *dto1.SDP, opts ...grpc.CallOption) (*types.Empty, error)
	// Session
	CreateSession(ctx context.Context, in *dto1.CreateSessionReq, opts ...grpc.CallOption) (*dto1.CreateSessionResp, error)
	// Entity
	ListEntity(ctx context.Context, in *dto.ListEntityReq, opts ...grpc.CallOption) (*dto.ListEntityResp, error)
	// Animation
	ListAnimation(ctx context.Context, in *dto.ListAnimationReq, opts ...grpc.CallOption) (*dto.ListAnimationResp, error)
	// PC
	CreatePC(ctx context.Context, in *dto.CreatePCReq, opts ...grpc.CallOption) (*entity.PC, error)
	ListPC(ctx context.Context, in *dto.ListPCReq, opts ...grpc.CallOption) (*dto.ListPCResp, error)
	GetPC(ctx context.Context, in *dto.GetPCReq, opts ...grpc.CallOption) (*dto.PC, error)
	// Template
	ListTemplate(ctx context.Context, in *dto.ListTemplateReq, opts ...grpc.CallOption) (*dto.ListTemplateResp, error)
	// Entity
	CreateEntity(ctx context.Context, in *entity.E, opts ...grpc.CallOption) (*entity.E, error)
	// Room
	CreateRoom(ctx context.Context, in *room.R, opts ...grpc.CallOption) (*room.R, error)
	ListRoom(ctx context.Context, in *dto2.ListRoomReq, opts ...grpc.CallOption) (*dto2.ListRoomResp, error)
	ListRoomPublic(ctx context.Context, in *dto2.ListRoomReq, opts ...grpc.CallOption) (*dto2.ListRoomResp, error)
	CreateRoomUser(ctx context.Context, in *dto2.CreateRoomUserReq, opts ...grpc.CallOption) (*room.User, error)
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

func (c *aPIClient) RTCConnectPC(ctx context.Context, in *dto1.SDP, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/grpc.API/RTCConnectPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateSession(ctx context.Context, in *dto1.CreateSessionReq, opts ...grpc.CallOption) (*dto1.CreateSessionResp, error) {
	out := new(dto1.CreateSessionResp)
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

func (c *aPIClient) GetPC(ctx context.Context, in *dto.GetPCReq, opts ...grpc.CallOption) (*dto.PC, error) {
	out := new(dto.PC)
	err := c.cc.Invoke(ctx, "/grpc.API/GetPC", in, out, opts...)
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

func (c *aPIClient) ListRoomPublic(ctx context.Context, in *dto2.ListRoomReq, opts ...grpc.CallOption) (*dto2.ListRoomResp, error) {
	out := new(dto2.ListRoomResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListRoomPublic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateRoomUser(ctx context.Context, in *dto2.CreateRoomUserReq, opts ...grpc.CallOption) (*room.User, error) {
	out := new(room.User)
	err := c.cc.Invoke(ctx, "/grpc.API/CreateRoomUser", in, out, opts...)
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
	// WIP
	RTCConnectPC(context.Context, *dto1.SDP) (*types.Empty, error)
	// Session
	CreateSession(context.Context, *dto1.CreateSessionReq) (*dto1.CreateSessionResp, error)
	// Entity
	ListEntity(context.Context, *dto.ListEntityReq) (*dto.ListEntityResp, error)
	// Animation
	ListAnimation(context.Context, *dto.ListAnimationReq) (*dto.ListAnimationResp, error)
	// PC
	CreatePC(context.Context, *dto.CreatePCReq) (*entity.PC, error)
	ListPC(context.Context, *dto.ListPCReq) (*dto.ListPCResp, error)
	GetPC(context.Context, *dto.GetPCReq) (*dto.PC, error)
	// Template
	ListTemplate(context.Context, *dto.ListTemplateReq) (*dto.ListTemplateResp, error)
	// Entity
	CreateEntity(context.Context, *entity.E) (*entity.E, error)
	// Room
	CreateRoom(context.Context, *room.R) (*room.R, error)
	ListRoom(context.Context, *dto2.ListRoomReq) (*dto2.ListRoomResp, error)
	ListRoomPublic(context.Context, *dto2.ListRoomReq) (*dto2.ListRoomResp, error)
	CreateRoomUser(context.Context, *dto2.CreateRoomUserReq) (*room.User, error)
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
func (*UnimplementedAPIServer) RTCConnectPC(ctx context.Context, req *dto1.SDP) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RTCConnectPC not implemented")
}
func (*UnimplementedAPIServer) CreateSession(ctx context.Context, req *dto1.CreateSessionReq) (*dto1.CreateSessionResp, error) {
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
func (*UnimplementedAPIServer) GetPC(ctx context.Context, req *dto.GetPCReq) (*dto.PC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPC not implemented")
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
func (*UnimplementedAPIServer) ListRoomPublic(ctx context.Context, req *dto2.ListRoomReq) (*dto2.ListRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoomPublic not implemented")
}
func (*UnimplementedAPIServer) CreateRoomUser(ctx context.Context, req *dto2.CreateRoomUserReq) (*room.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoomUser not implemented")
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

func _API_RTCConnectPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto1.SDP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).RTCConnectPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/RTCConnectPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).RTCConnectPC(ctx, req.(*dto1.SDP))
	}
	return interceptor(ctx, in, info, handler)
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

func _API_GetPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto.GetPCReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/GetPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetPC(ctx, req.(*dto.GetPCReq))
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

func _API_ListRoomPublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto2.ListRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListRoomPublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/ListRoomPublic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListRoomPublic(ctx, req.(*dto2.ListRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateRoomUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto2.CreateRoomUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateRoomUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/CreateRoomUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateRoomUser(ctx, req.(*dto2.CreateRoomUserReq))
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
			MethodName: "RTCConnectPC",
			Handler:    _API_RTCConnectPC_Handler,
		},
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
			MethodName: "GetPC",
			Handler:    _API_GetPC_Handler,
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
			MethodName: "ListRoomPublic",
			Handler:    _API_ListRoomPublic_Handler,
		},
		{
			MethodName: "CreateRoomUser",
			Handler:    _API_CreateRoomUser_Handler,
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
