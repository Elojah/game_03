// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/cmd/api/grpc/api.proto

package grpc

import (
	context "context"
	fmt "fmt"
	dto2 "github.com/elojah/game_03/pkg/ability/dto"
	entity "github.com/elojah/game_03/pkg/entity"
	dto1 "github.com/elojah/game_03/pkg/entity/dto"
	room "github.com/elojah/game_03/pkg/room"
	dto3 "github.com/elojah/game_03/pkg/room/dto"
	dto4 "github.com/elojah/game_03/pkg/twitch/dto"
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
	// 669 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x3b, 0x6f, 0xd4, 0x40,
	0x10, 0xb6, 0x45, 0xc8, 0x63, 0x49, 0x42, 0x98, 0x90, 0x14, 0x0e, 0x5a, 0x1a, 0x0a, 0x28, 0x62,
	0x87, 0x3c, 0x48, 0x24, 0x24, 0xa4, 0x60, 0x85, 0x87, 0x44, 0x61, 0x05, 0x10, 0x25, 0xf2, 0xf9,
	0x36, 0x8e, 0xc1, 0xbe, 0x5d, 0xbc, 0x7b, 0x8a, 0xd2, 0xf1, 0x13, 0xf8, 0x19, 0xf4, 0x34, 0x94,
	0x29, 0x53, 0xa6, 0x4c, 0xc9, 0xf9, 0x1a, 0xca, 0x94, 0x94, 0x68, 0xd6, 0xf6, 0xd9, 0xce, 0x4b,
	0x4e, 0x73, 0x37, 0xf3, 0xcd, 0xf7, 0xcd, 0x7e, 0xbb, 0x9a, 0x91, 0xc9, 0x4a, 0x18, 0xa9, 0xfd,
	0x7e, 0xc7, 0x0e, 0x78, 0xe2, 0xb0, 0x98, 0x7f, 0xf1, 0xf7, 0x9d, 0xd0, 0x4f, 0xd8, 0xe7, 0x95,
	0x35, 0x27, 0x48, 0xba, 0x8e, 0x2f, 0x22, 0x27, 0x4c, 0x45, 0x80, 0x81, 0x2d, 0x52, 0xae, 0x38,
	0x8c, 0x61, 0x6e, 0x2d, 0xd7, 0x74, 0x21, 0x0f, 0xb9, 0xa3, 0x8b, 0x9d, 0xfe, 0x9e, 0xce, 0x74,
	0xa2, 0xa3, 0x5c, 0x64, 0x2d, 0x85, 0x9c, 0x87, 0x31, 0xab, 0x58, 0x2c, 0x11, 0xea, 0xb0, 0x28,
	0x6e, 0x5e, 0xed, 0x41, 0x7c, 0x0d, 0x1d, 0xbf, 0x13, 0xc5, 0x91, 0x3a, 0x74, 0xba, 0x8a, 0x97,
	0x71, 0x21, 0x7c, 0x7a, 0xbd, 0x90, 0xf5, 0x14, 0xea, 0xf2, 0xbf, 0x42, 0xb2, 0xdc, 0x4a, 0x22,
	0x82, 0x82, 0xbe, 0xd1, 0x8a, 0x8e, 0xce, 0x1a, 0xa7, 0x6c, 0xb5, 0x96, 0xf9, 0xbd, 0x28, 0xf1,
	0x55, 0xc4, 0x7b, 0x37, 0xba, 0x12, 0x2a, 0x47, 0x1e, 0x37, 0x5b, 0x4b, 0x14, 0x4b, 0x44, 0xec,
	0x2b, 0xd6, 0xee, 0x2d, 0x52, 0xce, 0x13, 0xfd, 0x73, 0x03, 0x7a, 0x5f, 0xb2, 0xb4, 0xdd, 0x4d,
	0x34, 0x1d, 0x4d, 0x05, 0x2c, 0x8e, 0x6f, 0x28, 0xa9, 0x99, 0x6a, 0x2b, 0xa9, 0x19, 0x5b, 0x6d,
	0x29, 0x39, 0xe0, 0x69, 0xdc, 0x6d, 0x37, 0x07, 0xea, 0x20, 0x52, 0xc1, 0xbe, 0x56, 0xed, 0xf1,
	0x38, 0xe6, 0x07, 0x85, 0x6c, 0xfd, 0x7a, 0x19, 0x9a, 0xd2, 0x22, 0xc9, 0xa4, 0x1c, 0xcd, 0xc0,
	0xea, 0xaf, 0x09, 0x72, 0x6b, 0xdb, 0x7b, 0x0b, 0x2f, 0xc8, 0x8c, 0x9b, 0x32, 0x5f, 0xb1, 0xf7,
	0x79, 0x19, 0x16, 0xec, 0xae, 0xe2, 0x76, 0x03, 0xdb, 0x65, 0xdf, 0xac, 0xc5, 0xcb, 0x60, 0x29,
	0x60, 0x83, 0x90, 0x77, 0x91, 0x54, 0x3b, 0x7a, 0x00, 0x00, 0x34, 0xab, 0x02, 0x50, 0x39, 0x7f,
	0x01, 0x93, 0x02, 0x1e, 0x91, 0xe9, 0xbc, 0x57, 0x21, 0x9c, 0xb2, 0x8b, 0xd9, 0xde, 0xb1, 0xaa,
	0x10, 0x3c, 0x32, 0x5f, 0x67, 0x6d, 0xe7, 0x8b, 0x09, 0x4b, 0x35, 0x2f, 0x8d, 0x0a, 0x1e, 0xf7,
	0xe0, 0xea, 0xa2, 0x14, 0xb0, 0x45, 0xee, 0xa0, 0x93, 0xb2, 0x53, 0xe5, 0xad, 0xd6, 0xe1, 0xfe,
	0x45, 0x50, 0x0a, 0x7c, 0x28, 0x0d, 0x95, 0xbb, 0x54, 0x3c, 0x54, 0x03, 0xab, 0x1e, 0xea, 0x1c,
	0x2c, 0x05, 0x3c, 0x26, 0x93, 0xb9, 0x29, 0xcf, 0x85, 0xb9, 0x9a, 0x47, 0xcf, 0x45, 0x15, 0x29,
	0x2f, 0xed, 0xb9, 0xf0, 0x84, 0x8c, 0xa3, 0xdc, 0x73, 0x61, 0x76, 0xd4, 0x2b, 0x67, 0xdd, 0x6d,
	0xe4, 0x52, 0xc0, 0x43, 0x72, 0xfb, 0x35, 0x43, 0xe6, 0x8c, 0xae, 0xe8, 0x18, 0x89, 0x13, 0x3a,
	0xf5, 0x5c, 0x78, 0x4e, 0xa6, 0x91, 0xfe, 0xa1, 0x58, 0x4a, 0xa8, 0xee, 0x56, 0x42, 0x48, 0x5f,
	0xb8, 0x04, 0x95, 0x02, 0x28, 0x21, 0xb9, 0xc7, 0x5d, 0xce, 0x13, 0x98, 0xb0, 0xf5, 0x4a, 0xec,
	0x5a, 0x65, 0x00, 0x0e, 0x99, 0x44, 0x8d, 0xae, 0xce, 0x8d, 0x5a, 0x60, 0x8a, 0x4d, 0xef, 0x9d,
	0x43, 0xf4, 0xb0, 0xcc, 0x96, 0xb9, 0xd7, 0xef, 0xc4, 0x51, 0xd0, 0x4e, 0xb6, 0x4e, 0x66, 0x2b,
	0x1f, 0x1f, 0x25, 0x4b, 0xa1, 0x3e, 0x8d, 0x25, 0x98, 0x3f, 0xa3, 0xb6, 0xa6, 0x39, 0x85, 0x3b,
	0x97, 0xc5, 0x71, 0xed, 0x18, 0x4c, 0x9b, 0xc7, 0xe4, 0x88, 0x14, 0xb0, 0x4a, 0xa6, 0x30, 0xff,
	0x84, 0x2b, 0x09, 0x55, 0x5d, 0xe7, 0x28, 0x81, 0xf3, 0x50, 0x35, 0xfe, 0xaf, 0xf4, 0x42, 0xd6,
	0xc6, 0x3f, 0x07, 0x9a, 0xe3, 0x5f, 0x62, 0x52, 0xc0, 0x33, 0x32, 0xe6, 0x45, 0xbd, 0x10, 0x16,
	0xed, 0xfc, 0x9b, 0x65, 0x97, 0xdf, 0x2c, 0x7b, 0x07, 0xbf, 0x59, 0xd6, 0x15, 0xf8, 0xcb, 0x37,
	0xc7, 0x03, 0x6a, 0x9c, 0x0c, 0xa8, 0x71, 0x3a, 0xa0, 0xc6, 0xd9, 0x80, 0x9a, 0xff, 0x06, 0xd4,
	0xfc, 0x9e, 0x51, 0xf3, 0x67, 0x46, 0xcd, 0xdf, 0x19, 0x35, 0x8f, 0x32, 0x6a, 0x1e, 0x67, 0xd4,
	0x3c, 0xc9, 0xa8, 0xf9, 0x27, 0xa3, 0xe6, 0xdf, 0x8c, 0x1a, 0x67, 0x19, 0x35, 0x7f, 0x0c, 0xa9,
	0x71, 0x34, 0xa4, 0xe6, 0xc9, 0x90, 0x1a, 0xa7, 0x43, 0x6a, 0x74, 0xc6, 0x75, 0xe7, 0xb5, 0xff,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x75, 0x62, 0xb5, 0xab, 0x9c, 0x07, 0x00, 0x00,
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
	CreateEntity(ctx context.Context, in *entity.E, opts ...grpc.CallOption) (*entity.E, error)
	CreateEntityAbility(ctx context.Context, in *dto1.CreateEntityAbilityReq, opts ...grpc.CallOption) (*dto1.CreateEntityAbilityResp, error)
	// Ability
	ListAbility(ctx context.Context, in *dto2.ListAbilityReq, opts ...grpc.CallOption) (*dto2.ListAbilityResp, error)
	// Animation
	ListAnimation(ctx context.Context, in *dto1.ListAnimationReq, opts ...grpc.CallOption) (*dto1.ListAnimationResp, error)
	// PC
	CreatePC(ctx context.Context, in *dto1.CreatePCReq, opts ...grpc.CallOption) (*entity.PC, error)
	ListPC(ctx context.Context, in *dto1.ListPCReq, opts ...grpc.CallOption) (*dto1.ListPCResp, error)
	GetPC(ctx context.Context, in *dto1.GetPCReq, opts ...grpc.CallOption) (*dto1.PC, error)
	// Template
	ListTemplate(ctx context.Context, in *dto1.ListTemplateReq, opts ...grpc.CallOption) (*dto1.ListTemplateResp, error)
	// Room
	CreateRoom(ctx context.Context, in *room.R, opts ...grpc.CallOption) (*room.R, error)
	ListRoom(ctx context.Context, in *dto3.ListRoomReq, opts ...grpc.CallOption) (*dto3.ListRoomResp, error)
	ListRoomPublic(ctx context.Context, in *dto3.ListRoomReq, opts ...grpc.CallOption) (*dto3.ListRoomResp, error)
	CreateRoomUser(ctx context.Context, in *dto3.CreateRoomUserReq, opts ...grpc.CallOption) (*room.User, error)
	// Cell
	ListCell(ctx context.Context, in *dto3.ListCellReq, opts ...grpc.CallOption) (*dto3.ListCellResp, error)
	// World
	ListWorld(ctx context.Context, in *dto3.ListWorldReq, opts ...grpc.CallOption) (*dto3.ListWorldResp, error)
	// Twitch
	ListFollow(ctx context.Context, in *dto4.ListFollowReq, opts ...grpc.CallOption) (*dto4.ListFollowResp, error)
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

func (c *aPIClient) CreateEntity(ctx context.Context, in *entity.E, opts ...grpc.CallOption) (*entity.E, error) {
	out := new(entity.E)
	err := c.cc.Invoke(ctx, "/grpc.API/CreateEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateEntityAbility(ctx context.Context, in *dto1.CreateEntityAbilityReq, opts ...grpc.CallOption) (*dto1.CreateEntityAbilityResp, error) {
	out := new(dto1.CreateEntityAbilityResp)
	err := c.cc.Invoke(ctx, "/grpc.API/CreateEntityAbility", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListAbility(ctx context.Context, in *dto2.ListAbilityReq, opts ...grpc.CallOption) (*dto2.ListAbilityResp, error) {
	out := new(dto2.ListAbilityResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListAbility", in, out, opts...)
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

func (c *aPIClient) CreateRoom(ctx context.Context, in *room.R, opts ...grpc.CallOption) (*room.R, error) {
	out := new(room.R)
	err := c.cc.Invoke(ctx, "/grpc.API/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListRoom(ctx context.Context, in *dto3.ListRoomReq, opts ...grpc.CallOption) (*dto3.ListRoomResp, error) {
	out := new(dto3.ListRoomResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListRoomPublic(ctx context.Context, in *dto3.ListRoomReq, opts ...grpc.CallOption) (*dto3.ListRoomResp, error) {
	out := new(dto3.ListRoomResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListRoomPublic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateRoomUser(ctx context.Context, in *dto3.CreateRoomUserReq, opts ...grpc.CallOption) (*room.User, error) {
	out := new(room.User)
	err := c.cc.Invoke(ctx, "/grpc.API/CreateRoomUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListCell(ctx context.Context, in *dto3.ListCellReq, opts ...grpc.CallOption) (*dto3.ListCellResp, error) {
	out := new(dto3.ListCellResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListCell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListWorld(ctx context.Context, in *dto3.ListWorldReq, opts ...grpc.CallOption) (*dto3.ListWorldResp, error) {
	out := new(dto3.ListWorldResp)
	err := c.cc.Invoke(ctx, "/grpc.API/ListWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListFollow(ctx context.Context, in *dto4.ListFollowReq, opts ...grpc.CallOption) (*dto4.ListFollowResp, error) {
	out := new(dto4.ListFollowResp)
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
	CreateEntity(context.Context, *entity.E) (*entity.E, error)
	CreateEntityAbility(context.Context, *dto1.CreateEntityAbilityReq) (*dto1.CreateEntityAbilityResp, error)
	// Ability
	ListAbility(context.Context, *dto2.ListAbilityReq) (*dto2.ListAbilityResp, error)
	// Animation
	ListAnimation(context.Context, *dto1.ListAnimationReq) (*dto1.ListAnimationResp, error)
	// PC
	CreatePC(context.Context, *dto1.CreatePCReq) (*entity.PC, error)
	ListPC(context.Context, *dto1.ListPCReq) (*dto1.ListPCResp, error)
	GetPC(context.Context, *dto1.GetPCReq) (*dto1.PC, error)
	// Template
	ListTemplate(context.Context, *dto1.ListTemplateReq) (*dto1.ListTemplateResp, error)
	// Room
	CreateRoom(context.Context, *room.R) (*room.R, error)
	ListRoom(context.Context, *dto3.ListRoomReq) (*dto3.ListRoomResp, error)
	ListRoomPublic(context.Context, *dto3.ListRoomReq) (*dto3.ListRoomResp, error)
	CreateRoomUser(context.Context, *dto3.CreateRoomUserReq) (*room.User, error)
	// Cell
	ListCell(context.Context, *dto3.ListCellReq) (*dto3.ListCellResp, error)
	// World
	ListWorld(context.Context, *dto3.ListWorldReq) (*dto3.ListWorldResp, error)
	// Twitch
	ListFollow(context.Context, *dto4.ListFollowReq) (*dto4.ListFollowResp, error)
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
func (*UnimplementedAPIServer) CreateEntity(ctx context.Context, req *entity.E) (*entity.E, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntity not implemented")
}
func (*UnimplementedAPIServer) CreateEntityAbility(ctx context.Context, req *dto1.CreateEntityAbilityReq) (*dto1.CreateEntityAbilityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntityAbility not implemented")
}
func (*UnimplementedAPIServer) ListAbility(ctx context.Context, req *dto2.ListAbilityReq) (*dto2.ListAbilityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAbility not implemented")
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
func (*UnimplementedAPIServer) CreateRoom(ctx context.Context, req *room.R) (*room.R, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (*UnimplementedAPIServer) ListRoom(ctx context.Context, req *dto3.ListRoomReq) (*dto3.ListRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoom not implemented")
}
func (*UnimplementedAPIServer) ListRoomPublic(ctx context.Context, req *dto3.ListRoomReq) (*dto3.ListRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoomPublic not implemented")
}
func (*UnimplementedAPIServer) CreateRoomUser(ctx context.Context, req *dto3.CreateRoomUserReq) (*room.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoomUser not implemented")
}
func (*UnimplementedAPIServer) ListCell(ctx context.Context, req *dto3.ListCellReq) (*dto3.ListCellResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCell not implemented")
}
func (*UnimplementedAPIServer) ListWorld(ctx context.Context, req *dto3.ListWorldReq) (*dto3.ListWorldResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorld not implemented")
}
func (*UnimplementedAPIServer) ListFollow(ctx context.Context, req *dto4.ListFollowReq) (*dto4.ListFollowResp, error) {
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

func _API_CreateEntityAbility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto1.CreateEntityAbilityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateEntityAbility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/CreateEntityAbility",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateEntityAbility(ctx, req.(*dto1.CreateEntityAbilityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListAbility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto2.ListAbilityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListAbility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/ListAbility",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListAbility(ctx, req.(*dto2.ListAbilityReq))
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
	in := new(dto3.ListRoomReq)
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
		return srv.(APIServer).ListRoom(ctx, req.(*dto3.ListRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListRoomPublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto3.ListRoomReq)
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
		return srv.(APIServer).ListRoomPublic(ctx, req.(*dto3.ListRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateRoomUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto3.CreateRoomUserReq)
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
		return srv.(APIServer).CreateRoomUser(ctx, req.(*dto3.CreateRoomUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListCell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto3.ListCellReq)
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
		return srv.(APIServer).ListCell(ctx, req.(*dto3.ListCellReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto3.ListWorldReq)
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
		return srv.(APIServer).ListWorld(ctx, req.(*dto3.ListWorldReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dto4.ListFollowReq)
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
		return srv.(APIServer).ListFollow(ctx, req.(*dto4.ListFollowReq))
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
			MethodName: "CreateEntity",
			Handler:    _API_CreateEntity_Handler,
		},
		{
			MethodName: "CreateEntityAbility",
			Handler:    _API_CreateEntityAbility_Handler,
		},
		{
			MethodName: "ListAbility",
			Handler:    _API_ListAbility_Handler,
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
