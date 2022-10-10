// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/entity/dto/animation.proto

package dto

import (
	bytes "bytes"
	fmt "fmt"
	entity "github.com/elojah/game_03/pkg/entity"
	github_com_elojah_game_03_pkg_ulid "github.com/elojah/game_03/pkg/ulid"
	_ "github.com/gogo/protobuf/gogoproto"
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

type ListAnimationReq struct {
	IDs       []github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,rep,name=IDs,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"IDs"`
	EntityIDs []github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,rep,name=EntityIDs,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"EntityIDs"`
	Size_     int64                                   `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	State     []byte                                  `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *ListAnimationReq) Reset()      { *m = ListAnimationReq{} }
func (*ListAnimationReq) ProtoMessage() {}
func (*ListAnimationReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_90611f0100162c90, []int{0}
}
func (m *ListAnimationReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAnimationReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAnimationReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAnimationReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAnimationReq.Merge(m, src)
}
func (m *ListAnimationReq) XXX_Size() int {
	return m.Size()
}
func (m *ListAnimationReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAnimationReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListAnimationReq proto.InternalMessageInfo

type ListAnimationResp struct {
	Animations []entity.Animation `protobuf:"bytes,1,rep,name=Animations,proto3" json:"Animations"`
	State      []byte             `protobuf:"bytes,2,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *ListAnimationResp) Reset()      { *m = ListAnimationResp{} }
func (*ListAnimationResp) ProtoMessage() {}
func (*ListAnimationResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_90611f0100162c90, []int{1}
}
func (m *ListAnimationResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAnimationResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAnimationResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAnimationResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAnimationResp.Merge(m, src)
}
func (m *ListAnimationResp) XXX_Size() int {
	return m.Size()
}
func (m *ListAnimationResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAnimationResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListAnimationResp proto.InternalMessageInfo

type CreateAnimationReq struct {
	Animation      entity.Animation `protobuf:"bytes,1,opt,name=Animation,proto3" json:"Animation"`
	ID             string           `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	SheetID        string           `protobuf:"bytes,3,opt,name=SheetID,proto3" json:"SheetID,omitempty"`
	DuplicateID    string           `protobuf:"bytes,4,opt,name=DuplicateID,proto3" json:"DuplicateID,omitempty"`
	EntityID       string           `protobuf:"bytes,5,opt,name=EntityID,proto3" json:"EntityID,omitempty"`
	EntityTemplate string           `protobuf:"bytes,6,opt,name=EntityTemplate,proto3" json:"EntityTemplate,omitempty"`
}

func (m *CreateAnimationReq) Reset()      { *m = CreateAnimationReq{} }
func (*CreateAnimationReq) ProtoMessage() {}
func (*CreateAnimationReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_90611f0100162c90, []int{2}
}
func (m *CreateAnimationReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateAnimationReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateAnimationReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateAnimationReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAnimationReq.Merge(m, src)
}
func (m *CreateAnimationReq) XXX_Size() int {
	return m.Size()
}
func (m *CreateAnimationReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAnimationReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAnimationReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListAnimationReq)(nil), "dto.ListAnimationReq")
	golang_proto.RegisterType((*ListAnimationReq)(nil), "dto.ListAnimationReq")
	proto.RegisterType((*ListAnimationResp)(nil), "dto.ListAnimationResp")
	golang_proto.RegisterType((*ListAnimationResp)(nil), "dto.ListAnimationResp")
	proto.RegisterType((*CreateAnimationReq)(nil), "dto.CreateAnimationReq")
	golang_proto.RegisterType((*CreateAnimationReq)(nil), "dto.CreateAnimationReq")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/entity/dto/animation.proto", fileDescriptor_90611f0100162c90)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/entity/dto/animation.proto", fileDescriptor_90611f0100162c90)
}

var fileDescriptor_90611f0100162c90 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x3d, 0x6f, 0x13, 0x41,
	0x10, 0xdd, 0xf1, 0x39, 0x81, 0x9b, 0x44, 0x11, 0x59, 0x51, 0x9c, 0x5c, 0x4c, 0x4e, 0x91, 0x40,
	0x6e, 0x72, 0x87, 0x08, 0x08, 0x3a, 0x84, 0x39, 0x24, 0x4e, 0x50, 0xad, 0xe9, 0xd1, 0x39, 0x5e,
	0xce, 0x0b, 0xb6, 0xf7, 0x88, 0xd7, 0x05, 0x54, 0xfc, 0x04, 0x7e, 0x06, 0x3f, 0x81, 0x32, 0xa5,
	0x45, 0xe5, 0x32, 0x02, 0x29, 0xe2, 0xd6, 0x0d, 0x65, 0x4a, 0x4a, 0xe4, 0x3d, 0xf9, 0x03, 0x17,
	0x11, 0x4a, 0x37, 0xef, 0xcd, 0xbc, 0x79, 0xfb, 0x46, 0x8b, 0x8f, 0x73, 0x65, 0x7a, 0xe3, 0x4e,
	0x74, 0xa2, 0x07, 0xb1, 0xec, 0xeb, 0x77, 0x59, 0x2f, 0xce, 0xb3, 0x81, 0x7c, 0x73, 0xef, 0x38,
	0x2e, 0xde, 0xe7, 0xb1, 0x1c, 0x1a, 0x65, 0x3e, 0xc6, 0x5d, 0xa3, 0xe3, 0x6c, 0xa8, 0x06, 0x99,
	0x51, 0x7a, 0x18, 0x15, 0xa7, 0xda, 0x68, 0xee, 0x75, 0x8d, 0x6e, 0x1c, 0xad, 0xc9, 0x73, 0x9d,
	0xeb, 0xd8, 0xf5, 0x3a, 0xe3, 0xb7, 0x0e, 0x39, 0xe0, 0xaa, 0x4a, 0xd3, 0x78, 0xf0, 0x5f, 0x6e,
	0x1b, 0x4e, 0x87, 0xdf, 0x01, 0x6f, 0xbd, 0x52, 0x23, 0xf3, 0x74, 0xc1, 0x0b, 0xf9, 0x81, 0x3f,
	0x41, 0x2f, 0x4d, 0x46, 0x01, 0x84, 0x5e, 0x73, 0xb7, 0x75, 0x34, 0xb9, 0x38, 0x60, 0x3f, 0x2e,
	0x0e, 0xee, 0x5c, 0xbd, 0x7f, 0xdc, 0x57, 0xdd, 0x28, 0x4d, 0xc4, 0x5c, 0xc9, 0x5f, 0xa2, 0xff,
	0xdc, 0xf9, 0xcd, 0xd7, 0xd4, 0xae, 0xb3, 0x66, 0xa5, 0xe7, 0x1c, 0xeb, 0x6d, 0xf5, 0x49, 0x06,
	0x5e, 0x08, 0x4d, 0x4f, 0xb8, 0x9a, 0xdf, 0xc6, 0xad, 0xb6, 0xc9, 0x8c, 0x0c, 0xea, 0x21, 0x34,
	0x77, 0x45, 0x05, 0x0e, 0x3b, 0xb8, 0xbf, 0x91, 0x65, 0x54, 0xf0, 0x47, 0x88, 0x4b, 0xa2, 0xca,
	0xb4, 0x73, 0x7f, 0x3f, 0xaa, 0xce, 0x11, 0x2d, 0x3b, 0xad, 0xfa, 0xfc, 0x7d, 0x62, 0x6d, 0x74,
	0xe5, 0x51, 0x5b, 0xf7, 0xf8, 0x09, 0xc8, 0x9f, 0x9d, 0xca, 0xcc, 0xc8, 0x7f, 0x4e, 0xf6, 0x10,
	0xfd, 0x25, 0x0e, 0x20, 0x84, 0xab, 0x4c, 0x56, 0x93, 0x7c, 0x0f, 0x6b, 0x69, 0xe2, 0x0c, 0x7c,
	0x51, 0x4b, 0x13, 0x1e, 0xe0, 0x8d, 0x76, 0x4f, 0x4a, 0x93, 0x26, 0x2e, 0xae, 0x2f, 0x16, 0x90,
	0x87, 0xb8, 0x93, 0x8c, 0x8b, 0xbe, 0x3a, 0xc9, 0x8c, 0x4c, 0x13, 0x97, 0xdb, 0x17, 0xeb, 0x14,
	0x6f, 0xe0, 0xcd, 0xc5, 0xd1, 0x82, 0x2d, 0xd7, 0x5e, 0x62, 0x7e, 0x17, 0xf7, 0xaa, 0xfa, 0xb5,
	0x1c, 0x14, 0xfd, 0x79, 0xa8, 0x6d, 0x37, 0xb1, 0xc1, 0xb6, 0x5e, 0x4c, 0x4a, 0x62, 0xd3, 0x92,
	0xd8, 0x79, 0x49, 0xec, 0xb2, 0x24, 0xf8, 0x53, 0x12, 0x7c, 0xb6, 0x04, 0x5f, 0x2d, 0xc1, 0x37,
	0x4b, 0x70, 0x66, 0x09, 0x26, 0x96, 0x60, 0x6a, 0x09, 0x7e, 0x59, 0x82, 0xdf, 0x96, 0xd8, 0xa5,
	0x25, 0xf8, 0x32, 0x23, 0x76, 0x36, 0x23, 0x98, 0xce, 0x88, 0x9d, 0xcf, 0x88, 0x75, 0xb6, 0xdd,
	0xff, 0x3a, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xce, 0xf8, 0xd2, 0x05, 0x03, 0x00, 0x00,
}

func (this *ListAnimationReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListAnimationReq)
	if !ok {
		that2, ok := that.(ListAnimationReq)
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
	if len(this.EntityIDs) != len(that1.EntityIDs) {
		return false
	}
	for i := range this.EntityIDs {
		if !this.EntityIDs[i].Equal(that1.EntityIDs[i]) {
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
func (this *ListAnimationResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListAnimationResp)
	if !ok {
		that2, ok := that.(ListAnimationResp)
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
	if len(this.Animations) != len(that1.Animations) {
		return false
	}
	for i := range this.Animations {
		if !this.Animations[i].Equal(&that1.Animations[i]) {
			return false
		}
	}
	if !bytes.Equal(this.State, that1.State) {
		return false
	}
	return true
}
func (this *CreateAnimationReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateAnimationReq)
	if !ok {
		that2, ok := that.(CreateAnimationReq)
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
	if !this.Animation.Equal(&that1.Animation) {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.SheetID != that1.SheetID {
		return false
	}
	if this.DuplicateID != that1.DuplicateID {
		return false
	}
	if this.EntityID != that1.EntityID {
		return false
	}
	if this.EntityTemplate != that1.EntityTemplate {
		return false
	}
	return true
}
func (this *ListAnimationReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&dto.ListAnimationReq{")
	s = append(s, "IDs: "+fmt.Sprintf("%#v", this.IDs)+",\n")
	s = append(s, "EntityIDs: "+fmt.Sprintf("%#v", this.EntityIDs)+",\n")
	s = append(s, "Size_: "+fmt.Sprintf("%#v", this.Size_)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ListAnimationResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dto.ListAnimationResp{")
	if this.Animations != nil {
		vs := make([]entity.Animation, len(this.Animations))
		for i := range vs {
			vs[i] = this.Animations[i]
		}
		s = append(s, "Animations: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateAnimationReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&dto.CreateAnimationReq{")
	s = append(s, "Animation: "+strings.Replace(this.Animation.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "SheetID: "+fmt.Sprintf("%#v", this.SheetID)+",\n")
	s = append(s, "DuplicateID: "+fmt.Sprintf("%#v", this.DuplicateID)+",\n")
	s = append(s, "EntityID: "+fmt.Sprintf("%#v", this.EntityID)+",\n")
	s = append(s, "EntityTemplate: "+fmt.Sprintf("%#v", this.EntityTemplate)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAnimation(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ListAnimationReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAnimationReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAnimationReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintAnimation(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x22
	}
	if m.Size_ != 0 {
		i = encodeVarintAnimation(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x18
	}
	if len(m.EntityIDs) > 0 {
		for iNdEx := len(m.EntityIDs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.EntityIDs[iNdEx].Size()
				i -= size
				if _, err := m.EntityIDs[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintAnimation(dAtA, i, uint64(size))
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
				i = encodeVarintAnimation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListAnimationResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAnimationResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAnimationResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintAnimation(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Animations) > 0 {
		for iNdEx := len(m.Animations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Animations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAnimation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CreateAnimationReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateAnimationReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateAnimationReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EntityTemplate) > 0 {
		i -= len(m.EntityTemplate)
		copy(dAtA[i:], m.EntityTemplate)
		i = encodeVarintAnimation(dAtA, i, uint64(len(m.EntityTemplate)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.EntityID) > 0 {
		i -= len(m.EntityID)
		copy(dAtA[i:], m.EntityID)
		i = encodeVarintAnimation(dAtA, i, uint64(len(m.EntityID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DuplicateID) > 0 {
		i -= len(m.DuplicateID)
		copy(dAtA[i:], m.DuplicateID)
		i = encodeVarintAnimation(dAtA, i, uint64(len(m.DuplicateID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SheetID) > 0 {
		i -= len(m.SheetID)
		copy(dAtA[i:], m.SheetID)
		i = encodeVarintAnimation(dAtA, i, uint64(len(m.SheetID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintAnimation(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Animation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAnimation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAnimation(dAtA []byte, offset int, v uint64) int {
	offset -= sovAnimation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedListAnimationReq(r randyAnimation, easy bool) *ListAnimationReq {
	this := &ListAnimationReq{}
	v1 := r.Intn(10)
	this.IDs = make([]github_com_elojah_game_03_pkg_ulid.ID, v1)
	for i := 0; i < v1; i++ {
		v2 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
		this.IDs[i] = *v2
	}
	v3 := r.Intn(10)
	this.EntityIDs = make([]github_com_elojah_game_03_pkg_ulid.ID, v3)
	for i := 0; i < v3; i++ {
		v4 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
		this.EntityIDs[i] = *v4
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

func NewPopulatedListAnimationResp(r randyAnimation, easy bool) *ListAnimationResp {
	this := &ListAnimationResp{}
	if r.Intn(5) != 0 {
		v6 := r.Intn(5)
		this.Animations = make([]entity.Animation, v6)
		for i := 0; i < v6; i++ {
			v7 := entity.NewPopulatedAnimation(r, easy)
			this.Animations[i] = *v7
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

func NewPopulatedCreateAnimationReq(r randyAnimation, easy bool) *CreateAnimationReq {
	this := &CreateAnimationReq{}
	v9 := entity.NewPopulatedAnimation(r, easy)
	this.Animation = *v9
	this.ID = string(randStringAnimation(r))
	this.SheetID = string(randStringAnimation(r))
	this.DuplicateID = string(randStringAnimation(r))
	this.EntityID = string(randStringAnimation(r))
	this.EntityTemplate = string(randStringAnimation(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyAnimation interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAnimation(r randyAnimation) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAnimation(r randyAnimation) string {
	v10 := r.Intn(100)
	tmps := make([]rune, v10)
	for i := 0; i < v10; i++ {
		tmps[i] = randUTF8RuneAnimation(r)
	}
	return string(tmps)
}
func randUnrecognizedAnimation(r randyAnimation, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAnimation(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAnimation(dAtA []byte, r randyAnimation, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(key))
		v11 := r.Int63()
		if r.Intn(2) == 0 {
			v11 *= -1
		}
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(v11))
	case 1:
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAnimation(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ListAnimationReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IDs) > 0 {
		for _, e := range m.IDs {
			l = e.Size()
			n += 1 + l + sovAnimation(uint64(l))
		}
	}
	if len(m.EntityIDs) > 0 {
		for _, e := range m.EntityIDs {
			l = e.Size()
			n += 1 + l + sovAnimation(uint64(l))
		}
	}
	if m.Size_ != 0 {
		n += 1 + sovAnimation(uint64(m.Size_))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovAnimation(uint64(l))
	}
	return n
}

func (m *ListAnimationResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Animations) > 0 {
		for _, e := range m.Animations {
			l = e.Size()
			n += 1 + l + sovAnimation(uint64(l))
		}
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovAnimation(uint64(l))
	}
	return n
}

func (m *CreateAnimationReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Animation.Size()
	n += 1 + l + sovAnimation(uint64(l))
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovAnimation(uint64(l))
	}
	l = len(m.SheetID)
	if l > 0 {
		n += 1 + l + sovAnimation(uint64(l))
	}
	l = len(m.DuplicateID)
	if l > 0 {
		n += 1 + l + sovAnimation(uint64(l))
	}
	l = len(m.EntityID)
	if l > 0 {
		n += 1 + l + sovAnimation(uint64(l))
	}
	l = len(m.EntityTemplate)
	if l > 0 {
		n += 1 + l + sovAnimation(uint64(l))
	}
	return n
}

func sovAnimation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAnimation(x uint64) (n int) {
	return sovAnimation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ListAnimationReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListAnimationReq{`,
		`IDs:` + fmt.Sprintf("%v", this.IDs) + `,`,
		`EntityIDs:` + fmt.Sprintf("%v", this.EntityIDs) + `,`,
		`Size_:` + fmt.Sprintf("%v", this.Size_) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ListAnimationResp) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForAnimations := "[]Animation{"
	for _, f := range this.Animations {
		repeatedStringForAnimations += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForAnimations += "}"
	s := strings.Join([]string{`&ListAnimationResp{`,
		`Animations:` + repeatedStringForAnimations + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CreateAnimationReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateAnimationReq{`,
		`Animation:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Animation), "Animation", "entity.Animation", 1), `&`, ``, 1) + `,`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`SheetID:` + fmt.Sprintf("%v", this.SheetID) + `,`,
		`DuplicateID:` + fmt.Sprintf("%v", this.DuplicateID) + `,`,
		`EntityID:` + fmt.Sprintf("%v", this.EntityID) + `,`,
		`EntityTemplate:` + fmt.Sprintf("%v", this.EntityTemplate) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAnimation(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ListAnimationReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnimation
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
			return fmt.Errorf("proto: ListAnimationReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAnimationReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
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
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
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
				return fmt.Errorf("proto: wrong wireType = %d for field EntityIDs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
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
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_elojah_game_03_pkg_ulid.ID
			m.EntityIDs = append(m.EntityIDs, v)
			if err := m.EntityIDs[len(m.EntityIDs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowAnimation
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
					return ErrIntOverflowAnimation
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
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
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
			skippy, err := skipAnimation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAnimation
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
func (m *ListAnimationResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnimation
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
			return fmt.Errorf("proto: ListAnimationResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAnimationResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Animations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
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
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Animations = append(m.Animations, entity.Animation{})
			if err := m.Animations[len(m.Animations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowAnimation
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
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
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
			skippy, err := skipAnimation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAnimation
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
func (m *CreateAnimationReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnimation
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
			return fmt.Errorf("proto: CreateAnimationReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateAnimationReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Animation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
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
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Animation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SheetID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SheetID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DuplicateID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DuplicateID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntityID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityTemplate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntityTemplate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAnimation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAnimation
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
func skipAnimation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAnimation
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
					return 0, ErrIntOverflowAnimation
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
					return 0, ErrIntOverflowAnimation
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
				return 0, ErrInvalidLengthAnimation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAnimation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAnimation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAnimation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAnimation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAnimation = fmt.Errorf("proto: unexpected end of group")
)
