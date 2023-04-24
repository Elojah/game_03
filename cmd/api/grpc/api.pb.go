// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/cmd/api/grpc/api.proto

package grpc

import (
	context "context"
	fmt "fmt"
	entity "github.com/elojah/game_03/pkg/entity"
	dto1 "github.com/elojah/game_03/pkg/entity/dto"
	room "github.com/elojah/game_03/pkg/room"
	dto2 "github.com/elojah/game_03/pkg/room/dto"
	dto3 "github.com/elojah/game_03/pkg/twitch/dto"
	dto "github.com/elojah/game_03/pkg/user/dto"
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
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xbf, 0x6f, 0xd3, 0x40,
	0x18, 0xb5, 0x45, 0xe9, 0x8f, 0x53, 0x5b, 0xca, 0x07, 0xed, 0x60, 0xa4, 0x63, 0x61, 0x80, 0xa1,
	0x76, 0x49, 0x5b, 0x40, 0x42, 0x42, 0x2a, 0x56, 0xf8, 0x21, 0x31, 0x58, 0x01, 0xc4, 0x88, 0x1c,
	0xe7, 0xea, 0x18, 0xce, 0xb9, 0xc3, 0x77, 0x51, 0xd4, 0x8d, 0x3f, 0x81, 0xff, 0x02, 0xfe, 0x04,
	0xc6, 0x8e, 0x1d, 0x33, 0x76, 0x24, 0xce, 0xc2, 0xd8, 0x91, 0x11, 0xdd, 0xd9, 0x8e, 0xed, 0xd0,
	0x22, 0x67, 0x69, 0xef, 0x7b, 0xdf, 0x7b, 0xf7, 0x3d, 0x5f, 0xde, 0x1d, 0xda, 0x0b, 0x23, 0xd9,
	0x1f, 0x76, 0xed, 0x80, 0xc5, 0x0e, 0xa1, 0xec, 0x93, 0xdf, 0x77, 0x42, 0x3f, 0x26, 0x1f, 0xf7,
	0xf6, 0x9d, 0x20, 0xee, 0x39, 0x3e, 0x8f, 0x9c, 0x30, 0xe1, 0x81, 0x5a, 0xd8, 0x3c, 0x61, 0x92,
	0xc1, 0x92, 0xaa, 0xad, 0xdd, 0x8a, 0x2e, 0x64, 0x21, 0x73, 0x74, 0xb3, 0x3b, 0x3c, 0xd6, 0x95,
	0x2e, 0xf4, 0x2a, 0x13, 0x59, 0x77, 0x42, 0xc6, 0x42, 0x4a, 0x4a, 0x16, 0x89, 0xb9, 0x3c, 0xc9,
	0x9b, 0x0f, 0xaf, 0xf6, 0xc0, 0x3f, 0x87, 0x0e, 0x19, 0xc8, 0x48, 0x9e, 0xe4, 0xff, 0x72, 0xc9,
	0x6e, 0x23, 0x09, 0x0f, 0x72, 0xfa, 0x61, 0x23, 0x7a, 0x4f, 0xb2, 0xfa, 0x94, 0x27, 0x8d, 0x65,
	0xfe, 0x20, 0x8a, 0x7d, 0x19, 0xb1, 0xc1, 0x42, 0x9f, 0xa4, 0x94, 0x33, 0x8f, 0x8f, 0x1b, 0x4b,
	0x24, 0x89, 0x39, 0xf5, 0x25, 0x69, 0x76, 0x16, 0x09, 0x63, 0xb1, 0xfe, 0xb3, 0x00, 0x7d, 0x28,
	0x48, 0xd2, 0xec, 0x4b, 0x34, 0x5d, 0x99, 0x0a, 0x08, 0xa5, 0x0b, 0x4a, 0x2a, 0xa6, 0x9a, 0x4a,
	0x2a, 0xc6, 0x5a, 0x0d, 0x25, 0x23, 0x96, 0xd0, 0x5e, 0xb3, 0x1c, 0xc8, 0x51, 0x24, 0x83, 0xbe,
	0x56, 0x1d, 0x33, 0x4a, 0xd9, 0x28, 0x97, 0x1d, 0xfc, 0x5f, 0xa6, 0x4c, 0x69, 0x91, 0x20, 0x42,
	0xcc, 0x32, 0xd0, 0xfa, 0xbe, 0x8c, 0xae, 0x1d, 0x79, 0xaf, 0xe1, 0x19, 0xda, 0x70, 0x13, 0xe2,
	0x4b, 0xf2, 0x36, 0x6b, 0xc3, 0xb6, 0xdd, 0x93, 0xcc, 0xae, 0x61, 0x1d, 0xf2, 0xc5, 0xda, 0xb9,
	0x0c, 0x16, 0x1c, 0x0e, 0x11, 0x7a, 0x13, 0x09, 0xd9, 0xd6, 0x01, 0x00, 0xd0, 0xac, 0x12, 0x50,
	0xca, 0x5b, 0xff, 0x60, 0x82, 0xab, 0xb1, 0x0a, 0x39, 0x2a, 0x92, 0x99, 0x8f, 0xad, 0x61, 0xe5,
	0xd8, 0x39, 0x58, 0x70, 0xb8, 0x8f, 0x56, 0x33, 0x2f, 0x9e, 0x0b, 0x5b, 0x15, 0x6b, 0x9e, 0xab,
	0x54, 0xc8, 0xce, 0x6f, 0x8a, 0xe7, 0xc2, 0x03, 0xb4, 0xac, 0xe4, 0x9e, 0x0b, 0x9b, 0xb3, 0xbd,
	0x32, 0xd6, 0x8d, 0x5a, 0x2d, 0x38, 0xdc, 0x45, 0xd7, 0x5f, 0x12, 0xc5, 0xdc, 0xd0, 0x1d, 0xbd,
	0x56, 0xc4, 0x15, 0x5d, 0x7a, 0x2e, 0x3c, 0x45, 0xeb, 0x8a, 0xfe, 0x2e, 0x8f, 0x38, 0xdc, 0x9e,
	0xed, 0x50, 0x40, 0x8a, 0xbe, 0x7d, 0x09, 0x2a, 0x38, 0xdc, 0x43, 0xeb, 0x99, 0xc7, 0xfc, 0xac,
	0xd6, 0x0a, 0x93, 0x6d, 0xab, 0x5c, 0x02, 0x46, 0x28, 0x63, 0x75, 0x18, 0x8b, 0x61, 0xc5, 0xd6,
	0x31, 0xec, 0x58, 0xc5, 0x02, 0x1c, 0xb4, 0xaa, 0x76, 0xd6, 0xdd, 0xad, 0xd9, 0x20, 0x55, 0xaa,
	0xd1, 0x37, 0xe7, 0x10, 0xfd, 0x03, 0x6d, 0x16, 0xb5, 0x37, 0xec, 0xd2, 0x28, 0x68, 0x26, 0x3b,
	0x40, 0x9b, 0xa5, 0x8f, 0xf7, 0x82, 0x24, 0x50, 0x4d, 0x40, 0x01, 0x66, 0x87, 0xad, 0xad, 0x69,
	0x4e, 0xee, 0xce, 0x25, 0x94, 0x56, 0xc6, 0xa8, 0xb2, 0x3e, 0x26, 0x43, 0x04, 0x87, 0x16, 0x5a,
	0x53, 0xf5, 0x07, 0x75, 0x0d, 0xa0, 0xec, 0xeb, 0x5a, 0x49, 0x60, 0x1e, 0x2a, 0x23, 0xf7, 0x42,
	0x5f, 0x82, 0x4a, 0xe4, 0x32, 0xa0, 0x1e, 0xb9, 0x02, 0x13, 0x1c, 0x1e, 0xa1, 0x25, 0x2f, 0x1a,
	0x84, 0xb0, 0x63, 0x67, 0xcf, 0xbd, 0x5d, 0x3c, 0xf7, 0x76, 0x5b, 0x3d, 0xf7, 0xd6, 0x15, 0xf8,
	0xf3, 0x57, 0x67, 0x13, 0x6c, 0x8c, 0x27, 0xd8, 0x38, 0x9f, 0x60, 0xe3, 0x62, 0x82, 0xcd, 0x3f,
	0x13, 0x6c, 0x7e, 0x4d, 0xb1, 0xf9, 0x23, 0xc5, 0xe6, 0xcf, 0x14, 0x9b, 0xa7, 0x29, 0x36, 0xcf,
	0x52, 0x6c, 0x8e, 0x53, 0x6c, 0xfe, 0x4a, 0xb1, 0xf9, 0x3b, 0xc5, 0xc6, 0x45, 0x8a, 0xcd, 0x6f,
	0x53, 0x6c, 0x9c, 0x4e, 0xb1, 0x39, 0x9e, 0x62, 0xe3, 0x7c, 0x8a, 0x8d, 0xee, 0xb2, 0xde, 0x79,
	0xff, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x67, 0xb9, 0x88, 0xd7, 0x06, 0x00, 0x00,
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
	// Session
	CreateSession(ctx context.Context, in *dto.CreateSessionReq, opts ...grpc.CallOption) (*dto.CreateSessionResp, error)
	// Entity
	ListEntity(ctx context.Context, in *dto1.ListEntityReq, opts ...grpc.CallOption) (*dto1.ListEntityResp, error)
	// Animation
	ListAnimation(ctx context.Context, in *dto1.ListAnimationReq, opts ...grpc.CallOption) (*dto1.ListAnimationResp, error)
	// PC
	CreatePC(ctx context.Context, in *dto1.CreatePCReq, opts ...grpc.CallOption) (*entity.PC, error)
	ListPC(ctx context.Context, in *dto1.ListPCReq, opts ...grpc.CallOption) (*dto1.ListPCResp, error)
	GetPC(ctx context.Context, in *dto1.GetPCReq, opts ...grpc.CallOption) (*dto1.PC, error)
	// Template
	ListTemplate(ctx context.Context, in *dto1.ListTemplateReq, opts ...grpc.CallOption) (*dto1.ListTemplateResp, error)
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

func (c *aPIClient) CreateSession(ctx context.Context, in *dto.CreateSessionReq, opts ...grpc.CallOption) (*dto.CreateSessionResp, error) {
	out := new(dto.CreateSessionResp)
	err := c.cc.Invoke(ctx, "/grpc.API/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListEntity(ctx context.Context, in *dto1.ListEntityReq, opts ...grpc.CallOption) (*dto1.ListEntityResp, error) {
	out := new(dto1.ListEntityResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListAnimation(ctx context.Context, in *dto1.ListAnimationReq, opts ...grpc.CallOption) (*dto1.ListAnimationResp, error) {
	out := new(dto1.ListAnimationResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListAnimation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreatePC(ctx context.Context, in *dto1.CreatePCReq, opts ...grpc.CallOption) (*entity.PC, error) {
	out := new(entity.PC)
	err := c.cc.Invoke(ctx, "/grpc.API/CreatePC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListPC(ctx context.Context, in *dto1.ListPCReq, opts ...grpc.CallOption) (*dto1.ListPCResp, error) {
	out := new(dto1.ListPCResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetPC(ctx context.Context, in *dto1.GetPCReq, opts ...grpc.CallOption) (*dto1.PC, error) {
	out := new(dto1.PC)
	err := c.cc.Invoke(ctx, "/grpc.API/GetPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListTemplate(ctx context.Context, in *dto1.ListTemplateReq, opts ...grpc.CallOption) (*dto1.ListTemplateResp, error) {
	out := new(dto1.ListTemplateResp)
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
	// Session
	CreateSession(context.Context, *dto.CreateSessionReq) (*dto.CreateSessionResp, error)
	// Entity
	ListEntity(context.Context, *dto1.ListEntityReq) (*dto1.ListEntityResp, error)
	// Animation
	ListAnimation(context.Context, *dto1.ListAnimationReq) (*dto1.ListAnimationResp, error)
	// PC
	CreatePC(context.Context, *dto1.CreatePCReq) (*entity.PC, error)
	ListPC(context.Context, *dto1.ListPCReq) (*dto1.ListPCResp, error)
	GetPC(context.Context, *dto1.GetPCReq) (*dto1.PC, error)
	// Template
	ListTemplate(context.Context, *dto1.ListTemplateReq) (*dto1.ListTemplateResp, error)
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

func (*UnimplementedAPIServer) CreateSession(ctx context.Context, req *dto.CreateSessionReq) (*dto.CreateSessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (*UnimplementedAPIServer) ListEntity(ctx context.Context, req *dto1.ListEntityReq) (*dto1.ListEntityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntity not implemented")
}
func (*UnimplementedAPIServer) ListAnimation(ctx context.Context, req *dto1.ListAnimationReq) (*dto1.ListAnimationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnimation not implemented")
}
func (*UnimplementedAPIServer) CreatePC(ctx context.Context, req *dto1.CreatePCReq) (*entity.PC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePC not implemented")
}
func (*UnimplementedAPIServer) ListPC(ctx context.Context, req *dto1.ListPCReq) (*dto1.ListPCResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPC not implemented")
}
func (*UnimplementedAPIServer) GetPC(ctx context.Context, req *dto1.GetPCReq) (*dto1.PC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPC not implemented")
}
func (*UnimplementedAPIServer) ListTemplate(ctx context.Context, req *dto1.ListTemplateReq) (*dto1.ListTemplateResp, error) {
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

func _API_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto.CreateSessionReq)
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
		return srv.(APIServer).CreateSession(ctx, req.(*dto.CreateSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto1.ListEntityReq)
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
		return srv.(APIServer).ListEntity(ctx, req.(*dto1.ListEntityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListAnimation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto1.ListAnimationReq)
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
		return srv.(APIServer).ListAnimation(ctx, req.(*dto1.ListAnimationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreatePC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto1.CreatePCReq)
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
		return srv.(APIServer).CreatePC(ctx, req.(*dto1.CreatePCReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto1.ListPCReq)
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
		return srv.(APIServer).ListPC(ctx, req.(*dto1.ListPCReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto1.GetPCReq)
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
		return srv.(APIServer).GetPC(ctx, req.(*dto1.GetPCReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto1.ListTemplateReq)
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
		return srv.(APIServer).ListTemplate(ctx, req.(*dto1.ListTemplateReq))
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
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/elojah/game_03/cmd/api/grpc/api.proto",
}
