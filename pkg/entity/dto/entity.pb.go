// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/entity/dto/entity.proto

package dto

import (
	bytes "bytes"
	fmt "fmt"
	ability "github.com/elojah/game_03/pkg/ability"
	entity "github.com/elojah/game_03/pkg/entity"
	_ "github.com/elojah/game_03/pkg/gogoproto"
	github_com_elojah_game_03_pkg_ulid "github.com/elojah/game_03/pkg/ulid"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type ListEntityReq struct {
	IDs     []github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,rep,name=IDs,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"IDs"`
	CellIDs []github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,rep,name=CellIDs,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"CellIDs"`
	Size_   int64                                   `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	State   []byte                                  `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *ListEntityReq) Reset()      { *m = ListEntityReq{} }
func (*ListEntityReq) ProtoMessage() {}
func (*ListEntityReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0596b60aed7bf7a, []int{0}
}
func (m *ListEntityReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListEntityReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListEntityReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListEntityReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEntityReq.Merge(m, src)
}
func (m *ListEntityReq) XXX_Size() int {
	return m.Size()
}
func (m *ListEntityReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEntityReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListEntityReq proto.InternalMessageInfo

type ListEntityResp struct {
	Entities []entity.E `protobuf:"bytes,1,rep,name=Entities,proto3" json:"Entities"`
	State    []byte     `protobuf:"bytes,2,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *ListEntityResp) Reset()      { *m = ListEntityResp{} }
func (*ListEntityResp) ProtoMessage() {}
func (*ListEntityResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0596b60aed7bf7a, []int{1}
}
func (m *ListEntityResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListEntityResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListEntityResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListEntityResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEntityResp.Merge(m, src)
}
func (m *ListEntityResp) XXX_Size() int {
	return m.Size()
}
func (m *ListEntityResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEntityResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListEntityResp proto.InternalMessageInfo

type CreateEntityAbilityReq struct {
	EntityID github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=EntityID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"EntityID"`
	Ability  ability.A                             `protobuf:"bytes,2,opt,name=Ability,proto3" json:"Ability"`
}

func (m *CreateEntityAbilityReq) Reset()      { *m = CreateEntityAbilityReq{} }
func (*CreateEntityAbilityReq) ProtoMessage() {}
func (*CreateEntityAbilityReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0596b60aed7bf7a, []int{2}
}
func (m *CreateEntityAbilityReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateEntityAbilityReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateEntityAbilityReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateEntityAbilityReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntityAbilityReq.Merge(m, src)
}
func (m *CreateEntityAbilityReq) XXX_Size() int {
	return m.Size()
}
func (m *CreateEntityAbilityReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntityAbilityReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntityAbilityReq proto.InternalMessageInfo

type CreateEntityAbilityResp struct {
	EntityID  github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=EntityID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"EntityID"`
	AbilityID github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,opt,name=AbilityID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"AbilityID"`
}

func (m *CreateEntityAbilityResp) Reset()      { *m = CreateEntityAbilityResp{} }
func (*CreateEntityAbilityResp) ProtoMessage() {}
func (*CreateEntityAbilityResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0596b60aed7bf7a, []int{3}
}
func (m *CreateEntityAbilityResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateEntityAbilityResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateEntityAbilityResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateEntityAbilityResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntityAbilityResp.Merge(m, src)
}
func (m *CreateEntityAbilityResp) XXX_Size() int {
	return m.Size()
}
func (m *CreateEntityAbilityResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntityAbilityResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntityAbilityResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListEntityReq)(nil), "dto.ListEntityReq")
	golang_proto.RegisterType((*ListEntityReq)(nil), "dto.ListEntityReq")
	proto.RegisterType((*ListEntityResp)(nil), "dto.ListEntityResp")
	golang_proto.RegisterType((*ListEntityResp)(nil), "dto.ListEntityResp")
	proto.RegisterType((*CreateEntityAbilityReq)(nil), "dto.CreateEntityAbilityReq")
	golang_proto.RegisterType((*CreateEntityAbilityReq)(nil), "dto.CreateEntityAbilityReq")
	proto.RegisterType((*CreateEntityAbilityResp)(nil), "dto.CreateEntityAbilityResp")
	golang_proto.RegisterType((*CreateEntityAbilityResp)(nil), "dto.CreateEntityAbilityResp")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/entity/dto/entity.proto", fileDescriptor_f0596b60aed7bf7a)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/entity/dto/entity.proto", fileDescriptor_f0596b60aed7bf7a)
}

var fileDescriptor_f0596b60aed7bf7a = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbf, 0x8e, 0xd3, 0x40,
	0x10, 0x87, 0x77, 0xce, 0x81, 0xe3, 0xf6, 0x0e, 0x8a, 0x15, 0x02, 0xeb, 0x8a, 0xb9, 0x28, 0x12,
	0x52, 0x04, 0xc2, 0x86, 0x44, 0xd4, 0xe8, 0x12, 0x47, 0x60, 0x41, 0xe5, 0x3c, 0x00, 0x72, 0xc8,
	0xca, 0x31, 0x38, 0xac, 0x89, 0x37, 0x45, 0xa8, 0x78, 0x03, 0x78, 0x0c, 0x5a, 0x3a, 0xca, 0x48,
	0x34, 0x29, 0x53, 0x46, 0x14, 0x11, 0x5e, 0x37, 0x94, 0x29, 0x29, 0x91, 0xed, 0xcd, 0x9f, 0x02,
	0x45, 0x04, 0x5d, 0xe5, 0x59, 0x6b, 0xbe, 0x6f, 0x7f, 0x33, 0x36, 0x7d, 0x12, 0x84, 0x72, 0x30,
	0xee, 0x59, 0xaf, 0xc5, 0xd0, 0xe6, 0x91, 0x78, 0xe3, 0x0f, 0xec, 0xc0, 0x1f, 0xf2, 0x57, 0x8f,
	0x9a, 0x76, 0xfc, 0x36, 0xb0, 0xf9, 0x3b, 0x19, 0xca, 0x89, 0xdd, 0x97, 0x42, 0x97, 0x56, 0x3c,
	0x12, 0x52, 0x30, 0xa3, 0x2f, 0xc5, 0x79, 0x63, 0x3f, 0x1b, 0x88, 0x40, 0x14, 0xdd, 0x45, 0x55,
	0x82, 0xe7, 0xcd, 0xfd, 0x8c, 0xdf, 0x0b, 0xa3, 0xfc, 0x42, 0xfd, 0xd4, 0xd0, 0xe3, 0x7f, 0x0a,
	0xb9, 0x1b, 0xb0, 0xf6, 0x1d, 0xe8, 0xcd, 0x97, 0x61, 0x22, 0x3b, 0xc5, 0x4b, 0x8f, 0xbf, 0x67,
	0x4f, 0xa9, 0xe1, 0x3a, 0x89, 0x09, 0x55, 0xa3, 0x7e, 0xd6, 0x7a, 0x38, 0x5b, 0x5e, 0x90, 0x1f,
	0xcb, 0x8b, 0x7b, 0xfb, 0xcd, 0xe3, 0x28, 0xec, 0x5b, 0xae, 0xe3, 0xe5, 0x24, 0x7b, 0x46, 0x8f,
	0xdb, 0x3c, 0x8a, 0x72, 0xc9, 0xd1, 0xff, 0x48, 0xd6, 0x34, 0x63, 0xb4, 0xd2, 0x0d, 0x3f, 0x70,
	0xd3, 0xa8, 0x42, 0xdd, 0xf0, 0x8a, 0x9a, 0xdd, 0xa6, 0xd7, 0xba, 0xd2, 0x97, 0xdc, 0xac, 0x54,
	0xa1, 0x7e, 0xe6, 0x95, 0x87, 0x5a, 0x97, 0xde, 0xda, 0x1d, 0x22, 0x89, 0xd9, 0x03, 0x7a, 0xa3,
	0x38, 0x85, 0xbc, 0x1c, 0xe5, 0xb4, 0x71, 0x62, 0xe9, 0xc1, 0x3b, 0xad, 0x4a, 0x1e, 0xc8, 0xdb,
	0x34, 0x6c, 0xa5, 0x47, 0xbb, 0xd2, 0x4f, 0x40, 0xef, 0xb4, 0x47, 0xdc, 0x97, 0xbc, 0xf4, 0x5e,
	0x96, 0xbb, 0xce, 0x77, 0xe4, 0x6a, 0xfb, 0xc4, 0x75, 0x4c, 0xc8, 0x99, 0x43, 0x67, 0xdc, 0xe0,
	0xec, 0x3e, 0x3d, 0xd6, 0xe2, 0xe2, 0xf6, 0xd3, 0x06, 0xb5, 0xd6, 0x1f, 0xf5, 0x52, 0x07, 0x5d,
	0x37, 0xd4, 0xbe, 0x02, 0xbd, 0xfb, 0xd7, 0x44, 0x49, 0x7c, 0x95, 0x91, 0x5e, 0xd0, 0x13, 0x6d,
	0x76, 0x9d, 0x72, 0x25, 0x87, 0xba, 0xb6, 0x7c, 0xeb, 0xf9, 0x2c, 0x45, 0x32, 0x4f, 0x91, 0x2c,
	0x52, 0x24, 0xab, 0x14, 0xe1, 0x77, 0x8a, 0xf0, 0x51, 0x21, 0x7c, 0x51, 0x08, 0xdf, 0x14, 0xc2,
	0x54, 0x21, 0xcc, 0x14, 0xc2, 0x5c, 0x21, 0xfc, 0x54, 0x08, 0xbf, 0x14, 0x92, 0x95, 0x42, 0xf8,
	0x9c, 0x21, 0x99, 0x66, 0x08, 0xf3, 0x0c, 0xc9, 0x22, 0x43, 0xd2, 0xbb, 0x5e, 0xfc, 0xb1, 0xcd,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x59, 0xa3, 0xba, 0x34, 0x8b, 0x03, 0x00, 0x00,
}

func (this *ListEntityReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListEntityReq)
	if !ok {
		that2, ok := that.(ListEntityReq)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.IDs) != len(that1.IDs) {
		return false
	}
	for i := range this.IDs {
		if !this.IDs[i].Equal(that1.IDs[i]) {
			return false
		}
	}
	if len(this.CellIDs) != len(that1.CellIDs) {
		return false
	}
	for i := range this.CellIDs {
		if !this.CellIDs[i].Equal(that1.CellIDs[i]) {
			return false
		}
	}
	if this.Size_ != that1.Size_ {
		return false
	}
	if !bytes.Equal(this.State, that1.State) {
		return false
	}
	return true
}
func (this *ListEntityResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListEntityResp)
	if !ok {
		that2, ok := that.(ListEntityResp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Entities) != len(that1.Entities) {
		return false
	}
	for i := range this.Entities {
		if !this.Entities[i].Equal(&that1.Entities[i]) {
			return false
		}
	}
	if !bytes.Equal(this.State, that1.State) {
		return false
	}
	return true
}
func (this *CreateEntityAbilityReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateEntityAbilityReq)
	if !ok {
		that2, ok := that.(CreateEntityAbilityReq)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.EntityID.Equal(that1.EntityID) {
		return false
	}
	if !this.Ability.Equal(&that1.Ability) {
		return false
	}
	return true
}
func (this *CreateEntityAbilityResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateEntityAbilityResp)
	if !ok {
		that2, ok := that.(CreateEntityAbilityResp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.EntityID.Equal(that1.EntityID) {
		return false
	}
	if !this.AbilityID.Equal(that1.AbilityID) {
		return false
	}
	return true
}
func (this *ListEntityReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&dto.ListEntityReq{")
	s = append(s, "IDs: "+fmt.Sprintf("%#v", this.IDs)+",\n")
	s = append(s, "CellIDs: "+fmt.Sprintf("%#v", this.CellIDs)+",\n")
	s = append(s, "Size_: "+fmt.Sprintf("%#v", this.Size_)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ListEntityResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dto.ListEntityResp{")
	if this.Entities != nil {
		vs := make([]entity.E, len(this.Entities))
		for i := range vs {
			vs[i] = this.Entities[i]
		}
		s = append(s, "Entities: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateEntityAbilityReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dto.CreateEntityAbilityReq{")
	s = append(s, "EntityID: "+fmt.Sprintf("%#v", this.EntityID)+",\n")
	s = append(s, "Ability: "+strings.Replace(this.Ability.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateEntityAbilityResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dto.CreateEntityAbilityResp{")
	s = append(s, "EntityID: "+fmt.Sprintf("%#v", this.EntityID)+",\n")
	s = append(s, "AbilityID: "+fmt.Sprintf("%#v", this.AbilityID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringEntity(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ListEntityReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListEntityReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListEntityReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintEntity(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x22
	}
	if m.Size_ != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x18
	}
	if len(m.CellIDs) > 0 {
		for iNdEx := len(m.CellIDs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.CellIDs[iNdEx].Size()
				i -= size
				if _, err := m.CellIDs[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintEntity(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.IDs) > 0 {
		for iNdEx := len(m.IDs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.IDs[iNdEx].Size()
				i -= size
				if _, err := m.IDs[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintEntity(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListEntityResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListEntityResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListEntityResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintEntity(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Entities) > 0 {
		for iNdEx := len(m.Entities) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entities[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEntity(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CreateEntityAbilityReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateEntityAbilityReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateEntityAbilityReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Ability.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEntity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.EntityID.Size()
		i -= size
		if _, err := m.EntityID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEntity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CreateEntityAbilityResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateEntityAbilityResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateEntityAbilityResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AbilityID.Size()
		i -= size
		if _, err := m.AbilityID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEntity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.EntityID.Size()
		i -= size
		if _, err := m.EntityID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEntity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEntity(dAtA []byte, offset int, v uint64) int {
	offset -= sovEntity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedListEntityReq(r randyEntity, easy bool) *ListEntityReq {
	this := &ListEntityReq{}
	v1 := r.Intn(10)
	this.IDs = make([]github_com_elojah_game_03_pkg_ulid.ID, v1)
	for i := 0; i < v1; i++ {
		v2 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
		this.IDs[i] = *v2
	}
	v3 := r.Intn(10)
	this.CellIDs = make([]github_com_elojah_game_03_pkg_ulid.ID, v3)
	for i := 0; i < v3; i++ {
		v4 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
		this.CellIDs[i] = *v4
	}
	this.Size_ = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Size_ *= -1
	}
	v5 := r.Intn(100)
	this.State = make([]byte, v5)
	for i := 0; i < v5; i++ {
		this.State[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedListEntityResp(r randyEntity, easy bool) *ListEntityResp {
	this := &ListEntityResp{}
	if r.Intn(5) != 0 {
		v6 := r.Intn(5)
		this.Entities = make([]entity.E, v6)
		for i := 0; i < v6; i++ {
			v7 := entity.NewPopulatedE(r, easy)
			this.Entities[i] = *v7
		}
	}
	v8 := r.Intn(100)
	this.State = make([]byte, v8)
	for i := 0; i < v8; i++ {
		this.State[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedCreateEntityAbilityReq(r randyEntity, easy bool) *CreateEntityAbilityReq {
	this := &CreateEntityAbilityReq{}
	v9 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.EntityID = *v9
	v10 := ability.NewPopulatedA(r, easy)
	this.Ability = *v10
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedCreateEntityAbilityResp(r randyEntity, easy bool) *CreateEntityAbilityResp {
	this := &CreateEntityAbilityResp{}
	v11 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.EntityID = *v11
	v12 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.AbilityID = *v12
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyEntity interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneEntity(r randyEntity) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringEntity(r randyEntity) string {
	v13 := r.Intn(100)
	tmps := make([]rune, v13)
	for i := 0; i < v13; i++ {
		tmps[i] = randUTF8RuneEntity(r)
	}
	return string(tmps)
}
func randUnrecognizedEntity(r randyEntity, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldEntity(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldEntity(dAtA []byte, r randyEntity, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(key))
		v14 := r.Int63()
		if r.Intn(2) == 0 {
			v14 *= -1
		}
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(v14))
	case 1:
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateEntity(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ListEntityReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IDs) > 0 {
		for _, e := range m.IDs {
			l = e.Size()
			n += 1 + l + sovEntity(uint64(l))
		}
	}
	if len(m.CellIDs) > 0 {
		for _, e := range m.CellIDs {
			l = e.Size()
			n += 1 + l + sovEntity(uint64(l))
		}
	}
	if m.Size_ != 0 {
		n += 1 + sovEntity(uint64(m.Size_))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovEntity(uint64(l))
	}
	return n
}

func (m *ListEntityResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entities) > 0 {
		for _, e := range m.Entities {
			l = e.Size()
			n += 1 + l + sovEntity(uint64(l))
		}
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovEntity(uint64(l))
	}
	return n
}

func (m *CreateEntityAbilityReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EntityID.Size()
	n += 1 + l + sovEntity(uint64(l))
	l = m.Ability.Size()
	n += 1 + l + sovEntity(uint64(l))
	return n
}

func (m *CreateEntityAbilityResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EntityID.Size()
	n += 1 + l + sovEntity(uint64(l))
	l = m.AbilityID.Size()
	n += 1 + l + sovEntity(uint64(l))
	return n
}

func sovEntity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEntity(x uint64) (n int) {
	return sovEntity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ListEntityReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListEntityReq{`,
		`IDs:` + fmt.Sprintf("%v", this.IDs) + `,`,
		`CellIDs:` + fmt.Sprintf("%v", this.CellIDs) + `,`,
		`Size_:` + fmt.Sprintf("%v", this.Size_) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ListEntityResp) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForEntities := "[]E{"
	for _, f := range this.Entities {
		repeatedStringForEntities += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForEntities += "}"
	s := strings.Join([]string{`&ListEntityResp{`,
		`Entities:` + repeatedStringForEntities + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CreateEntityAbilityReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateEntityAbilityReq{`,
		`EntityID:` + fmt.Sprintf("%v", this.EntityID) + `,`,
		`Ability:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Ability), "A", "ability.A", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CreateEntityAbilityResp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateEntityAbilityResp{`,
		`EntityID:` + fmt.Sprintf("%v", this.EntityID) + `,`,
		`AbilityID:` + fmt.Sprintf("%v", this.AbilityID) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEntity(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ListEntityReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListEntityReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListEntityReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_elojah_game_03_pkg_ulid.ID
			m.IDs = append(m.IDs, v)
			if err := m.IDs[len(m.IDs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CellIDs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_elojah_game_03_pkg_ulid.ID
			m.CellIDs = append(m.CellIDs, v)
			if err := m.CellIDs[len(m.CellIDs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = append(m.State[:0], dAtA[iNdEx:postIndex]...)
			if m.State == nil {
				m.State = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEntity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListEntityResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListEntityResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListEntityResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entities = append(m.Entities, entity.E{})
			if err := m.Entities[len(m.Entities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = append(m.State[:0], dAtA[iNdEx:postIndex]...)
			if m.State == nil {
				m.State = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEntity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateEntityAbilityReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateEntityAbilityReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateEntityAbilityReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EntityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ability", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ability.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEntity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateEntityAbilityResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateEntityAbilityResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateEntityAbilityResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EntityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbilityID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AbilityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEntity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEntity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntity
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEntity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEntity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEntity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEntity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEntity = fmt.Errorf("proto: unexpected end of group")
)
