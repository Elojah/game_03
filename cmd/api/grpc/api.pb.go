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
	// 716 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x3b, 0x4f, 0xdc, 0x4c,
	0x14, 0xb5, 0xf5, 0xf1, 0xf1, 0x98, 0x00, 0x21, 0x97, 0x40, 0x61, 0xa2, 0x49, 0x93, 0x22, 0x29,
	0xb0, 0x09, 0x8f, 0x40, 0x14, 0x29, 0x12, 0xb1, 0xc8, 0x43, 0x4a, 0x61, 0x91, 0xa0, 0x94, 0xc8,
	0xeb, 0x1d, 0x8c, 0x13, 0x7b, 0x67, 0xe2, 0x99, 0x15, 0xa2, 0xcb, 0x4f, 0xc8, 0xcf, 0x48, 0x97,
	0x36, 0x25, 0x25, 0x25, 0x25, 0x65, 0xd6, 0xdb, 0xa4, 0xa4, 0x4c, 0x19, 0xcd, 0xd8, 0x5e, 0x8f,
	0x61, 0x17, 0x79, 0x1b, 0x98, 0x7b, 0xee, 0x39, 0x77, 0xce, 0xdc, 0x9d, 0x3b, 0x46, 0x6b, 0x61,
	0x24, 0x8e, 0xbb, 0x2d, 0x3b, 0xa0, 0x89, 0x43, 0x62, 0xfa, 0xd9, 0x3f, 0x76, 0x42, 0x3f, 0x21,
	0x87, 0x6b, 0x1b, 0x4e, 0x90, 0xb4, 0x1d, 0x9f, 0x45, 0x4e, 0x98, 0xb2, 0x40, 0x2e, 0x6c, 0x96,
	0x52, 0x41, 0x61, 0x42, 0xc6, 0xd6, 0xaa, 0xa6, 0x0b, 0x69, 0x48, 0x1d, 0x95, 0x6c, 0x75, 0x8f,
	0x54, 0xa4, 0x02, 0xb5, 0xca, 0x45, 0xd6, 0x4a, 0x48, 0x69, 0x18, 0x93, 0x8a, 0x45, 0x12, 0x26,
	0x4e, 0x8b, 0xe4, 0xf6, 0x68, 0x0f, 0xec, 0x4b, 0xe8, 0xf8, 0xad, 0x28, 0x8e, 0xc4, 0xa9, 0xd3,
	0x16, 0xb4, 0x5c, 0x17, 0xc2, 0xa7, 0xb7, 0x0b, 0x49, 0x47, 0x48, 0x5d, 0xfe, 0xaf, 0x90, 0xac,
	0x36, 0x92, 0xb0, 0xa0, 0xa0, 0x3f, 0x6f, 0x48, 0x3f, 0x64, 0x29, 0x39, 0x22, 0x29, 0xe9, 0x04,
	0x84, 0x17, 0xd2, 0xad, 0x46, 0x52, 0x79, 0xa8, 0x9a, 0xc1, 0x9d, 0xc6, 0x32, 0xbf, 0x13, 0x25,
	0xbe, 0x88, 0x68, 0x67, 0xac, 0x6e, 0x48, 0xe5, 0xe0, 0x78, 0xdb, 0x8d, 0x25, 0x82, 0x24, 0x2c,
	0xf6, 0x05, 0x69, 0xd6, 0xc6, 0x94, 0xd2, 0x44, 0xfd, 0x19, 0x83, 0xde, 0xe5, 0x24, 0x6d, 0x76,
	0x12, 0x45, 0x97, 0xa6, 0x02, 0x12, 0xc7, 0x63, 0x4a, 0x34, 0x53, 0x4d, 0x25, 0x9a, 0xb1, 0xf5,
	0x86, 0x92, 0x13, 0x9a, 0xc6, 0xed, 0x66, 0xf7, 0x40, 0x9c, 0x44, 0x22, 0x38, 0x56, 0xaa, 0x23,
	0x1a, 0xc7, 0xf4, 0xa4, 0x90, 0x6d, 0xde, 0x2e, 0x93, 0xa6, 0x94, 0x88, 0x13, 0xce, 0x07, 0x77,
	0x60, 0xfd, 0xe7, 0x34, 0xfa, 0x6f, 0xd7, 0x7b, 0x07, 0x2f, 0xd1, 0x9c, 0x9b, 0x12, 0x5f, 0x90,
	0x0f, 0x79, 0x1a, 0x96, 0xec, 0xb6, 0xa0, 0x76, 0x0d, 0xdb, 0x27, 0x5f, 0xad, 0xe5, 0x61, 0x30,
	0x67, 0xb0, 0x85, 0xd0, 0xfb, 0x88, 0x8b, 0x3d, 0x75, 0x01, 0x00, 0x14, 0xab, 0x02, 0xa4, 0x72,
	0xf1, 0x06, 0xc6, 0x19, 0x3c, 0x42, 0xb3, 0x79, 0xad, 0x42, 0x38, 0x63, 0x17, 0x77, 0x7b, 0xcf,
	0xaa, 0x96, 0xe0, 0xa1, 0x45, 0x9d, 0xb5, 0x9b, 0xcf, 0x34, 0xac, 0x68, 0x5e, 0x6a, 0x19, 0xb9,
	0xdd, 0x83, 0xd1, 0x49, 0xce, 0x60, 0x07, 0xdd, 0x91, 0x4e, 0xca, 0x4a, 0x95, 0x37, 0xad, 0xc2,
	0xfd, 0x9b, 0x20, 0x67, 0xb2, 0x51, 0x0a, 0x2a, 0x67, 0xa9, 0x68, 0x54, 0x0d, 0xab, 0x1a, 0x75,
	0x0d, 0xe6, 0x0c, 0x1e, 0xa3, 0xe9, 0xdc, 0x94, 0xe7, 0xc2, 0x82, 0xe6, 0xd1, 0x73, 0xa5, 0x0a,
	0x95, 0x87, 0xf6, 0x5c, 0x78, 0x82, 0x26, 0xa5, 0xdc, 0x73, 0x61, 0x7e, 0x50, 0x2b, 0x67, 0xdd,
	0xad, 0xc5, 0x9c, 0xc1, 0x43, 0xf4, 0xff, 0x1b, 0x22, 0x99, 0x73, 0x2a, 0xa3, 0xd6, 0x92, 0x38,
	0xa5, 0x42, 0xcf, 0x85, 0x2d, 0xb4, 0xa0, 0x40, 0xaf, 0x7a, 0x75, 0x40, 0xdb, 0xcb, 0x5a, 0xaa,
	0xd6, 0x3a, 0xc5, 0x45, 0x8b, 0x07, 0xac, 0xad, 0xdc, 0xe9, 0xf0, 0x70, 0xf6, 0xa8, 0x22, 0x2f,
	0xd0, 0xac, 0xb4, 0xfa, 0xb1, 0x78, 0x10, 0xa0, 0xea, 0x6b, 0x09, 0x49, 0xab, 0x4b, 0x43, 0x50,
	0xce, 0x00, 0x23, 0x94, 0xf7, 0x67, 0x9f, 0xd2, 0x04, 0xa6, 0x6c, 0x35, 0x8e, 0xfb, 0x56, 0xb9,
	0x00, 0x07, 0x4d, 0x4b, 0x8d, 0xca, 0x2e, 0x0c, 0x4a, 0xc8, 0x50, 0x16, 0xbd, 0x77, 0x0d, 0x51,
	0x17, 0x75, 0xbe, 0x8c, 0xbd, 0x6e, 0x2b, 0x8e, 0x82, 0x66, 0xb2, 0x4d, 0x34, 0x5f, 0xf9, 0x38,
	0xe0, 0x24, 0x05, 0x7d, 0x12, 0x4a, 0x30, 0xff, 0x09, 0x95, 0x35, 0xc5, 0x29, 0xdc, 0xb9, 0x24,
	0x8e, 0xb5, 0x6d, 0x64, 0x58, 0xdf, 0x26, 0x47, 0x38, 0x83, 0x75, 0x34, 0x23, 0xe3, 0x4f, 0xf2,
	0x39, 0x80, 0x2a, 0xaf, 0x62, 0x29, 0x81, 0xeb, 0x50, 0x35, 0x7a, 0xaf, 0xd5, 0x63, 0xa0, 0x8d,
	0x5e, 0x0e, 0xd4, 0x47, 0xaf, 0xc4, 0x38, 0x83, 0x67, 0x68, 0xc2, 0x8b, 0x3a, 0x21, 0x2c, 0xdb,
	0xf9, 0xa7, 0xd6, 0x2e, 0x3f, 0xb5, 0xf6, 0x9e, 0xfc, 0xd4, 0x5a, 0x23, 0xf0, 0x57, 0x6f, 0xcf,
	0x7b, 0xd8, 0xb8, 0xe8, 0x61, 0xe3, 0xb2, 0x87, 0x8d, 0xab, 0x1e, 0x36, 0xff, 0xf6, 0xb0, 0xf9,
	0x2d, 0xc3, 0xe6, 0x8f, 0x0c, 0x9b, 0xbf, 0x32, 0x6c, 0x9e, 0x65, 0xd8, 0x3c, 0xcf, 0xb0, 0x79,
	0x91, 0x61, 0xf3, 0x77, 0x86, 0xcd, 0x3f, 0x19, 0x36, 0xae, 0x32, 0x6c, 0x7e, 0xef, 0x63, 0xe3,
	0xac, 0x8f, 0xcd, 0x8b, 0x3e, 0x36, 0x2e, 0xfb, 0xd8, 0x68, 0x4d, 0xaa, 0xca, 0x1b, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xfa, 0x07, 0x91, 0x3a, 0x53, 0x08, 0x00, 0x00,
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
	GetPCPreferences(ctx context.Context, in *entity.PC, opts ...grpc.CallOption) (*entity.PCPreferences, error)
	UpdatePCPreferences(ctx context.Context, in *entity.PCPreferences, opts ...grpc.CallOption) (*entity.PCPreferences, error)
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

func (c *aPIClient) GetPCPreferences(ctx context.Context, in *entity.PC, opts ...grpc.CallOption) (*entity.PCPreferences, error) {
	out := new(entity.PCPreferences)
	err := c.cc.Invoke(ctx, "/grpc.API/GetPCPreferences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) UpdatePCPreferences(ctx context.Context, in *entity.PCPreferences, opts ...grpc.CallOption) (*entity.PCPreferences, error) {
	out := new(entity.PCPreferences)
	err := c.cc.Invoke(ctx, "/grpc.API/UpdatePCPreferences", in, out, opts...)
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
	GetPCPreferences(context.Context, *entity.PC) (*entity.PCPreferences, error)
	UpdatePCPreferences(context.Context, *entity.PCPreferences) (*entity.PCPreferences, error)
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
func (*UnimplementedAPIServer) GetPCPreferences(ctx context.Context, req *entity.PC) (*entity.PCPreferences, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPCPreferences not implemented")
}
func (*UnimplementedAPIServer) UpdatePCPreferences(ctx context.Context, req *entity.PCPreferences) (*entity.PCPreferences, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePCPreferences not implemented")
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

func _API_GetPCPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(entity.PC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetPCPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/GetPCPreferences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetPCPreferences(ctx, req.(*entity.PC))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_UpdatePCPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(entity.PCPreferences)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).UpdatePCPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.API/UpdatePCPreferences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).UpdatePCPreferences(ctx, req.(*entity.PCPreferences))
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
			MethodName: "GetPCPreferences",
			Handler:    _API_GetPCPreferences_Handler,
		},
		{
			MethodName: "UpdatePCPreferences",
			Handler:    _API_UpdatePCPreferences_Handler,
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
