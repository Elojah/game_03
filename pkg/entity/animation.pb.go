// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/entity/animation.proto

package entity

import (
	fmt "fmt"
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

type Animation struct {
	ID           github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"ID"`
	EntityID     github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,opt,name=EntityID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"EntityID"`
	SheetID      github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,3,opt,name=SheetID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"SheetID"`
	Name         string                                `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Start        int64                                 `protobuf:"varint,5,opt,name=Start,proto3" json:"Start,omitempty"`
	End          int64                                 `protobuf:"varint,6,opt,name=End,proto3" json:"End,omitempty"`
	Rate         int32                                 `protobuf:"varint,7,opt,name=Rate,proto3" json:"Rate,omitempty"`
	FrameWidth   int64                                 `protobuf:"varint,8,opt,name=FrameWidth,proto3" json:"FrameWidth,omitempty"`
	FrameHeight  int64                                 `protobuf:"varint,9,opt,name=FrameHeight,proto3" json:"FrameHeight,omitempty"`
	FrameStart   int64                                 `protobuf:"varint,10,opt,name=FrameStart,proto3" json:"FrameStart,omitempty"`
	FrameEnd     int64                                 `protobuf:"varint,11,opt,name=FrameEnd,proto3" json:"FrameEnd,omitempty"`
	FrameMargin  int64                                 `protobuf:"varint,12,opt,name=FrameMargin,proto3" json:"FrameMargin,omitempty"`
	FrameSpacing int64                                 `protobuf:"varint,13,opt,name=FrameSpacing,proto3" json:"FrameSpacing,omitempty"`
}

func (m *Animation) Reset()      { *m = Animation{} }
func (*Animation) ProtoMessage() {}
func (*Animation) Descriptor() ([]byte, []int) {
	return fileDescriptor_dda9402c6c71e187, []int{0}
}
func (m *Animation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Animation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Animation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Animation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Animation.Merge(m, src)
}
func (m *Animation) XXX_Size() int {
	return m.Size()
}
func (m *Animation) XXX_DiscardUnknown() {
	xxx_messageInfo_Animation.DiscardUnknown(m)
}

var xxx_messageInfo_Animation proto.InternalMessageInfo

func (m *Animation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Animation) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Animation) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Animation) GetRate() int32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *Animation) GetFrameWidth() int64 {
	if m != nil {
		return m.FrameWidth
	}
	return 0
}

func (m *Animation) GetFrameHeight() int64 {
	if m != nil {
		return m.FrameHeight
	}
	return 0
}

func (m *Animation) GetFrameStart() int64 {
	if m != nil {
		return m.FrameStart
	}
	return 0
}

func (m *Animation) GetFrameEnd() int64 {
	if m != nil {
		return m.FrameEnd
	}
	return 0
}

func (m *Animation) GetFrameMargin() int64 {
	if m != nil {
		return m.FrameMargin
	}
	return 0
}

func (m *Animation) GetFrameSpacing() int64 {
	if m != nil {
		return m.FrameSpacing
	}
	return 0
}

func init() {
	proto.RegisterType((*Animation)(nil), "entity.Animation")
	golang_proto.RegisterType((*Animation)(nil), "entity.Animation")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/entity/animation.proto", fileDescriptor_dda9402c6c71e187)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/entity/animation.proto", fileDescriptor_dda9402c6c71e187)
}

var fileDescriptor_dda9402c6c71e187 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x1c, 0xc5, 0xf3, 0x6f, 0xae, 0xbd, 0xd6, 0x57, 0x24, 0x64, 0x31, 0x58, 0x37, 0xfc, 0x2f, 0x3a,
	0x09, 0x29, 0xcb, 0x35, 0x48, 0xc7, 0xca, 0xc0, 0x29, 0x07, 0x64, 0x80, 0x21, 0x1d, 0x18, 0x91,
	0xdb, 0x1a, 0xc7, 0xd0, 0x24, 0x55, 0xe4, 0x0e, 0x6c, 0x7c, 0x04, 0x3e, 0x06, 0x1f, 0x81, 0xb1,
	0x63, 0xc7, 0x8e, 0x15, 0x43, 0x45, 0x9c, 0x85, 0xb1, 0x62, 0x62, 0x44, 0x75, 0xd4, 0x28, 0x2c,
	0x37, 0x74, 0x7b, 0xef, 0xf9, 0xff, 0x7e, 0x7a, 0x83, 0xc9, 0x73, 0xa9, 0x74, 0xb2, 0x9c, 0x8c,
	0xa6, 0x79, 0x1a, 0x88, 0x79, 0xfe, 0x89, 0x27, 0x81, 0xe4, 0xa9, 0xf8, 0xf0, 0xec, 0x36, 0x58,
	0x7c, 0x96, 0x81, 0xc8, 0xb4, 0xd2, 0x5f, 0x02, 0x9e, 0xa9, 0x94, 0x6b, 0x95, 0x67, 0xa3, 0x45,
	0x91, 0xeb, 0x9c, 0xf6, 0xea, 0xfc, 0xf2, 0xa6, 0xd5, 0x96, 0xb9, 0xcc, 0x03, 0xfb, 0x3c, 0x59,
	0x7e, 0xb4, 0xce, 0x1a, 0xab, 0xea, 0xda, 0xf5, 0x1f, 0x97, 0x0c, 0x5e, 0x1e, 0x51, 0xf4, 0x05,
	0xe9, 0x44, 0x21, 0x03, 0x0f, 0xfc, 0xe1, 0xdd, 0xcd, 0x7a, 0x77, 0xe5, 0xfc, 0xdc, 0x5d, 0x3d,
	0x7d, 0x78, 0xce, 0x72, 0xae, 0x66, 0xa3, 0x28, 0x8c, 0x3b, 0x51, 0x48, 0x23, 0xd2, 0xbf, 0xb7,
	0x2b, 0xa2, 0x90, 0x75, 0x4e, 0x81, 0x34, 0x75, 0xfa, 0x9a, 0x9c, 0x8f, 0x13, 0x21, 0x74, 0x14,
	0x32, 0xf7, 0x14, 0xd2, 0xb1, 0x4d, 0x29, 0x39, 0x7b, 0xc7, 0x53, 0xc1, 0xce, 0x3c, 0xf0, 0x07,
	0xb1, 0xd5, 0xf4, 0x09, 0xe9, 0x8e, 0x35, 0x2f, 0x34, 0xeb, 0x7a, 0xe0, 0xbb, 0x71, 0x6d, 0xe8,
	0x63, 0xe2, 0xde, 0x67, 0x33, 0xd6, 0xb3, 0xd9, 0x41, 0x1e, 0xba, 0x31, 0xd7, 0x82, 0x9d, 0x7b,
	0xe0, 0x77, 0x63, 0xab, 0x29, 0x12, 0xf2, 0xaa, 0xe0, 0xa9, 0x78, 0xaf, 0x66, 0x3a, 0x61, 0x7d,
	0x7b, 0xdc, 0x4a, 0xa8, 0x47, 0x2e, 0xac, 0x7b, 0x23, 0x94, 0x4c, 0x34, 0x1b, 0xd8, 0x83, 0x76,
	0xd4, 0x10, 0xea, 0x09, 0xa4, 0x45, 0xa8, 0x77, 0x5c, 0x92, 0xbe, 0x75, 0x87, 0x31, 0x17, 0xf6,
	0xb5, 0xf1, 0x0d, 0xfd, 0x2d, 0x2f, 0xa4, 0xca, 0xd8, 0xb0, 0x45, 0xaf, 0x23, 0x7a, 0x4d, 0x86,
	0x35, 0x6b, 0xc1, 0xa7, 0x2a, 0x93, 0xec, 0x91, 0x3d, 0xf9, 0x2f, 0xbb, 0x0b, 0x37, 0x25, 0x3a,
	0xdb, 0x12, 0x9d, 0x7d, 0x89, 0xf0, 0xb7, 0x44, 0xf8, 0x6a, 0x10, 0xbe, 0x1b, 0x84, 0x1f, 0x06,
	0x61, 0x65, 0x10, 0xd6, 0x06, 0x61, 0x63, 0x10, 0x7e, 0x19, 0x84, 0xdf, 0x06, 0x9d, 0xbd, 0x41,
	0xf8, 0x56, 0xa1, 0xb3, 0xaa, 0x10, 0x36, 0x15, 0x3a, 0xdb, 0x0a, 0x9d, 0x49, 0xcf, 0xfe, 0xa0,
	0xdb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xa3, 0xc0, 0xbb, 0xb0, 0x02, 0x00, 0x00,
}

func (this *Animation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Animation)
	if !ok {
		that2, ok := that.(Animation)
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
	if !this.ID.Equal(that1.ID) {
		return false
	}
	if !this.EntityID.Equal(that1.EntityID) {
		return false
	}
	if !this.SheetID.Equal(that1.SheetID) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	if this.Rate != that1.Rate {
		return false
	}
	if this.FrameWidth != that1.FrameWidth {
		return false
	}
	if this.FrameHeight != that1.FrameHeight {
		return false
	}
	if this.FrameStart != that1.FrameStart {
		return false
	}
	if this.FrameEnd != that1.FrameEnd {
		return false
	}
	if this.FrameMargin != that1.FrameMargin {
		return false
	}
	if this.FrameSpacing != that1.FrameSpacing {
		return false
	}
	return true
}
func (this *Animation) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 17)
	s = append(s, "&entity.Animation{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "EntityID: "+fmt.Sprintf("%#v", this.EntityID)+",\n")
	s = append(s, "SheetID: "+fmt.Sprintf("%#v", this.SheetID)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Start: "+fmt.Sprintf("%#v", this.Start)+",\n")
	s = append(s, "End: "+fmt.Sprintf("%#v", this.End)+",\n")
	s = append(s, "Rate: "+fmt.Sprintf("%#v", this.Rate)+",\n")
	s = append(s, "FrameWidth: "+fmt.Sprintf("%#v", this.FrameWidth)+",\n")
	s = append(s, "FrameHeight: "+fmt.Sprintf("%#v", this.FrameHeight)+",\n")
	s = append(s, "FrameStart: "+fmt.Sprintf("%#v", this.FrameStart)+",\n")
	s = append(s, "FrameEnd: "+fmt.Sprintf("%#v", this.FrameEnd)+",\n")
	s = append(s, "FrameMargin: "+fmt.Sprintf("%#v", this.FrameMargin)+",\n")
	s = append(s, "FrameSpacing: "+fmt.Sprintf("%#v", this.FrameSpacing)+",\n")
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
func (m *Animation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Animation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Animation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FrameSpacing != 0 {
		i = encodeVarintAnimation(dAtA, i, uint64(m.FrameSpacing))
		i--
		dAtA[i] = 0x68
	}
	if m.FrameMargin != 0 {
		i = encodeVarintAnimation(dAtA, i, uint64(m.FrameMargin))
		i--
		dAtA[i] = 0x60
	}
	if m.FrameEnd != 0 {
		i = encodeVarintAnimation(dAtA, i, uint64(m.FrameEnd))
		i--
		dAtA[i] = 0x58
	}
	if m.FrameStart != 0 {
		i = encodeVarintAnimation(dAtA, i, uint64(m.FrameStart))
		i--
		dAtA[i] = 0x50
	}
	if m.FrameHeight != 0 {
		i = encodeVarintAnimation(dAtA, i, uint64(m.FrameHeight))
		i--
		dAtA[i] = 0x48
	}
	if m.FrameWidth != 0 {
		i = encodeVarintAnimation(dAtA, i, uint64(m.FrameWidth))
		i--
		dAtA[i] = 0x40
	}
	if m.Rate != 0 {
		i = encodeVarintAnimation(dAtA, i, uint64(m.Rate))
		i--
		dAtA[i] = 0x38
	}
	if m.End != 0 {
		i = encodeVarintAnimation(dAtA, i, uint64(m.End))
		i--
		dAtA[i] = 0x30
	}
	if m.Start != 0 {
		i = encodeVarintAnimation(dAtA, i, uint64(m.Start))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAnimation(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.SheetID.Size()
		i -= size
		if _, err := m.SheetID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAnimation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.EntityID.Size()
		i -= size
		if _, err := m.EntityID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAnimation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
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
func NewPopulatedAnimation(r randyAnimation, easy bool) *Animation {
	this := &Animation{}
	v1 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.ID = *v1
	v2 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.EntityID = *v2
	v3 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.SheetID = *v3
	this.Name = string(randStringAnimation(r))
	this.Start = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Start *= -1
	}
	this.End = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.End *= -1
	}
	this.Rate = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Rate *= -1
	}
	this.FrameWidth = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.FrameWidth *= -1
	}
	this.FrameHeight = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.FrameHeight *= -1
	}
	this.FrameStart = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.FrameStart *= -1
	}
	this.FrameEnd = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.FrameEnd *= -1
	}
	this.FrameMargin = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.FrameMargin *= -1
	}
	this.FrameSpacing = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.FrameSpacing *= -1
	}
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
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
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
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(v5))
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
func (m *Animation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovAnimation(uint64(l))
	l = m.EntityID.Size()
	n += 1 + l + sovAnimation(uint64(l))
	l = m.SheetID.Size()
	n += 1 + l + sovAnimation(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAnimation(uint64(l))
	}
	if m.Start != 0 {
		n += 1 + sovAnimation(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovAnimation(uint64(m.End))
	}
	if m.Rate != 0 {
		n += 1 + sovAnimation(uint64(m.Rate))
	}
	if m.FrameWidth != 0 {
		n += 1 + sovAnimation(uint64(m.FrameWidth))
	}
	if m.FrameHeight != 0 {
		n += 1 + sovAnimation(uint64(m.FrameHeight))
	}
	if m.FrameStart != 0 {
		n += 1 + sovAnimation(uint64(m.FrameStart))
	}
	if m.FrameEnd != 0 {
		n += 1 + sovAnimation(uint64(m.FrameEnd))
	}
	if m.FrameMargin != 0 {
		n += 1 + sovAnimation(uint64(m.FrameMargin))
	}
	if m.FrameSpacing != 0 {
		n += 1 + sovAnimation(uint64(m.FrameSpacing))
	}
	return n
}

func sovAnimation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAnimation(x uint64) (n int) {
	return sovAnimation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Animation) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Animation{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`EntityID:` + fmt.Sprintf("%v", this.EntityID) + `,`,
		`SheetID:` + fmt.Sprintf("%v", this.SheetID) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Start:` + fmt.Sprintf("%v", this.Start) + `,`,
		`End:` + fmt.Sprintf("%v", this.End) + `,`,
		`Rate:` + fmt.Sprintf("%v", this.Rate) + `,`,
		`FrameWidth:` + fmt.Sprintf("%v", this.FrameWidth) + `,`,
		`FrameHeight:` + fmt.Sprintf("%v", this.FrameHeight) + `,`,
		`FrameStart:` + fmt.Sprintf("%v", this.FrameStart) + `,`,
		`FrameEnd:` + fmt.Sprintf("%v", this.FrameEnd) + `,`,
		`FrameMargin:` + fmt.Sprintf("%v", this.FrameMargin) + `,`,
		`FrameSpacing:` + fmt.Sprintf("%v", this.FrameSpacing) + `,`,
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
func (m *Animation) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Animation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Animation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
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
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
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
			if err := m.EntityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SheetID", wireType)
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
			if err := m.SheetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			m.Rate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rate |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrameWidth", wireType)
			}
			m.FrameWidth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FrameWidth |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrameHeight", wireType)
			}
			m.FrameHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FrameHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrameStart", wireType)
			}
			m.FrameStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FrameStart |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrameEnd", wireType)
			}
			m.FrameEnd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FrameEnd |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrameMargin", wireType)
			}
			m.FrameMargin = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FrameMargin |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrameSpacing", wireType)
			}
			m.FrameSpacing = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FrameSpacing |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAnimation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAnimation
			}
			if (iNdEx + skippy) < 0 {
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
