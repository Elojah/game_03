// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/entity/entity.proto

package entity

import (
	fmt "fmt"
	geometry "github.com/elojah/game_03/pkg/geometry"
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

type E struct {
	ID             github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"ID"`
	UserID         github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,opt,name=UserID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"UserID"`
	CellID         github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,3,opt,name=CellID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"CellID"`
	Name           string                                `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	X              int64                                 `protobuf:"varint,5,opt,name=X,proto3" json:"X,omitempty"`
	Y              int64                                 `protobuf:"varint,6,opt,name=Y,proto3" json:"Y,omitempty"`
	Rot            int32                                 `protobuf:"varint,7,opt,name=Rot,proto3" json:"Rot,omitempty"`
	Radius         int32                                 `protobuf:"varint,8,opt,name=Radius,proto3" json:"Radius,omitempty"`
	At             int64                                 `protobuf:"varint,9,opt,name=At,proto3" json:"At,omitempty"`
	AnimationID    github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,10,opt,name=AnimationID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"AnimationID"`
	AnimationAt    int64                                 `protobuf:"varint,11,opt,name=AnimationAt,proto3" json:"AnimationAt,omitempty"`
	StaticObjects  []geometry.Rect                       `protobuf:"bytes,12,rep,name=StaticObjects,proto3" json:"StaticObjects"`
	DynamicObjects []geometry.Rect                       `protobuf:"bytes,13,rep,name=DynamicObjects,proto3" json:"DynamicObjects"`
}

func (m *E) Reset()      { *m = E{} }
func (*E) ProtoMessage() {}
func (*E) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2701362edd4c296, []int{0}
}
func (m *E) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *E) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_E.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *E) XXX_Merge(src proto.Message) {
	xxx_messageInfo_E.Merge(m, src)
}
func (m *E) XXX_Size() int {
	return m.Size()
}
func (m *E) XXX_DiscardUnknown() {
	xxx_messageInfo_E.DiscardUnknown(m)
}

var xxx_messageInfo_E proto.InternalMessageInfo

func init() {
	proto.RegisterType((*E)(nil), "entity.E")
	golang_proto.RegisterType((*E)(nil), "entity.E")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/entity/entity.proto", fileDescriptor_e2701362edd4c296)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/entity/entity.proto", fileDescriptor_e2701362edd4c296)
}

var fileDescriptor_e2701362edd4c296 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xef, 0x39, 0xa9, 0xa1, 0x97, 0x34, 0x42, 0x37, 0xa0, 0x53, 0x87, 0x57, 0x0b, 0x09,
	0xc9, 0x4b, 0x63, 0xa0, 0x62, 0x41, 0x30, 0x24, 0xb8, 0x12, 0x5e, 0xa8, 0x74, 0x08, 0xa9, 0x9d,
	0x90, 0x93, 0x1e, 0xae, 0x4b, 0x9c, 0xab, 0x92, 0xcb, 0x90, 0x8d, 0x8f, 0xc0, 0xb7, 0x80, 0x8f,
	0xc0, 0xd8, 0x31, 0x63, 0xc6, 0x8a, 0xa1, 0xc2, 0xe7, 0x85, 0xb1, 0x23, 0x23, 0xf2, 0x39, 0x0a,
	0x86, 0x21, 0x43, 0xa6, 0xfb, 0xff, 0xdf, 0xbd, 0xdf, 0xff, 0x3d, 0xe9, 0x8e, 0x3e, 0x4d, 0x52,
	0x7d, 0x31, 0x1b, 0x74, 0x87, 0x2a, 0x0b, 0xe4, 0x48, 0x5d, 0xc6, 0x17, 0x41, 0x12, 0x67, 0xf2,
	0xc3, 0x93, 0xa3, 0xe0, 0xea, 0x53, 0x12, 0xc8, 0xb1, 0x4e, 0xf5, 0x7c, 0x75, 0x74, 0xaf, 0x26,
	0x4a, 0x2b, 0xe6, 0x56, 0x6e, 0xff, 0xb0, 0x86, 0x26, 0x2a, 0x51, 0x81, 0xbd, 0x1e, 0xcc, 0x3e,
	0x5a, 0x67, 0x8d, 0x55, 0x15, 0xb6, 0xff, 0x7c, 0xf3, 0xa4, 0x44, 0xaa, 0x4c, 0xea, 0xc9, 0x7c,
	0x2d, 0x2a, 0xec, 0xd1, 0xd7, 0x26, 0x85, 0x63, 0xf6, 0x8a, 0x3a, 0x51, 0xc8, 0xc1, 0x03, 0xbf,
	0xdd, 0x3f, 0x5c, 0xdc, 0x1e, 0x90, 0x1f, 0xb7, 0x07, 0x8f, 0x37, 0x07, 0xce, 0x46, 0xe9, 0x79,
	0x37, 0x0a, 0x85, 0x13, 0x85, 0xec, 0x98, 0xba, 0xef, 0xa7, 0x72, 0x12, 0x85, 0xdc, 0xd9, 0x26,
	0x62, 0x05, 0x97, 0x31, 0xaf, 0xe5, 0x68, 0x14, 0x85, 0xbc, 0xb1, 0x55, 0x4c, 0x05, 0x33, 0x46,
	0x9b, 0x6f, 0xe3, 0x4c, 0xf2, 0xa6, 0x07, 0xfe, 0xae, 0xb0, 0x9a, 0xb5, 0x29, 0x9c, 0xf2, 0x1d,
	0x0f, 0xfc, 0x86, 0x80, 0xd3, 0xd2, 0x9d, 0x71, 0xb7, 0x72, 0x67, 0xec, 0x01, 0x6d, 0x08, 0xa5,
	0xf9, 0x3d, 0x0f, 0xfc, 0x1d, 0x51, 0x4a, 0xf6, 0x90, 0xba, 0x22, 0x3e, 0x4f, 0x67, 0x53, 0x7e,
	0xdf, 0x16, 0x57, 0x8e, 0x75, 0xa8, 0xd3, 0xd3, 0x7c, 0xd7, 0x82, 0x4e, 0x4f, 0xb3, 0x13, 0xda,
	0xea, 0x8d, 0xd3, 0x2c, 0xd6, 0xa9, 0x1a, 0x47, 0x21, 0xa7, 0xdb, 0x6c, 0x5d, 0x4f, 0x60, 0x5e,
	0x2d, 0xb0, 0xa7, 0x79, 0xcb, 0x4e, 0xaa, 0x97, 0xd8, 0x0b, 0xba, 0xf7, 0x4e, 0xc7, 0x3a, 0x1d,
	0x9e, 0x0c, 0x2e, 0xe5, 0x50, 0x4f, 0x79, 0xdb, 0x6b, 0xf8, 0xad, 0x67, 0x9d, 0xee, 0xfa, 0x5d,
	0x85, 0x1c, 0xea, 0x7e, 0xb3, 0x5c, 0x42, 0xfc, 0xdb, 0xca, 0x5e, 0xd2, 0x4e, 0x38, 0x1f, 0xc7,
	0xd9, 0x5f, 0x78, 0x6f, 0x03, 0xfc, 0x5f, 0x6f, 0xff, 0xcd, 0x22, 0x47, 0xb2, 0xcc, 0x91, 0xdc,
	0xe4, 0x48, 0xee, 0x72, 0x84, 0xdf, 0x39, 0xc2, 0x67, 0x83, 0xf0, 0xcd, 0x20, 0x7c, 0x37, 0x08,
	0xd7, 0x06, 0x61, 0x61, 0x10, 0x96, 0x06, 0xe1, 0xa7, 0x41, 0xf8, 0x65, 0x90, 0xdc, 0x19, 0x84,
	0x2f, 0x05, 0x92, 0xeb, 0x02, 0x61, 0x59, 0x20, 0xb9, 0x29, 0x90, 0x0c, 0x5c, 0xfb, 0xf5, 0x8e,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x93, 0x29, 0x85, 0x1d, 0x03, 0x00, 0x00,
}

func (this *E) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*E)
	if !ok {
		that2, ok := that.(E)
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
	if !this.UserID.Equal(that1.UserID) {
		return false
	}
	if !this.CellID.Equal(that1.CellID) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.X != that1.X {
		return false
	}
	if this.Y != that1.Y {
		return false
	}
	if this.Rot != that1.Rot {
		return false
	}
	if this.Radius != that1.Radius {
		return false
	}
	if this.At != that1.At {
		return false
	}
	if !this.AnimationID.Equal(that1.AnimationID) {
		return false
	}
	if this.AnimationAt != that1.AnimationAt {
		return false
	}
	if len(this.StaticObjects) != len(that1.StaticObjects) {
		return false
	}
	for i := range this.StaticObjects {
		if !this.StaticObjects[i].Equal(&that1.StaticObjects[i]) {
			return false
		}
	}
	if len(this.DynamicObjects) != len(that1.DynamicObjects) {
		return false
	}
	for i := range this.DynamicObjects {
		if !this.DynamicObjects[i].Equal(&that1.DynamicObjects[i]) {
			return false
		}
	}
	return true
}
func (this *E) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 17)
	s = append(s, "&entity.E{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "UserID: "+fmt.Sprintf("%#v", this.UserID)+",\n")
	s = append(s, "CellID: "+fmt.Sprintf("%#v", this.CellID)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "X: "+fmt.Sprintf("%#v", this.X)+",\n")
	s = append(s, "Y: "+fmt.Sprintf("%#v", this.Y)+",\n")
	s = append(s, "Rot: "+fmt.Sprintf("%#v", this.Rot)+",\n")
	s = append(s, "Radius: "+fmt.Sprintf("%#v", this.Radius)+",\n")
	s = append(s, "At: "+fmt.Sprintf("%#v", this.At)+",\n")
	s = append(s, "AnimationID: "+fmt.Sprintf("%#v", this.AnimationID)+",\n")
	s = append(s, "AnimationAt: "+fmt.Sprintf("%#v", this.AnimationAt)+",\n")
	if this.StaticObjects != nil {
		vs := make([]geometry.Rect, len(this.StaticObjects))
		for i := range vs {
			vs[i] = this.StaticObjects[i]
		}
		s = append(s, "StaticObjects: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	if this.DynamicObjects != nil {
		vs := make([]geometry.Rect, len(this.DynamicObjects))
		for i := range vs {
			vs[i] = this.DynamicObjects[i]
		}
		s = append(s, "DynamicObjects: "+fmt.Sprintf("%#v", vs)+",\n")
	}
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
func (m *E) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *E) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *E) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DynamicObjects) > 0 {
		for iNdEx := len(m.DynamicObjects) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DynamicObjects[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEntity(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.StaticObjects) > 0 {
		for iNdEx := len(m.StaticObjects) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StaticObjects[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEntity(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if m.AnimationAt != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.AnimationAt))
		i--
		dAtA[i] = 0x58
	}
	{
		size := m.AnimationID.Size()
		i -= size
		if _, err := m.AnimationID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEntity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.At != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.At))
		i--
		dAtA[i] = 0x48
	}
	if m.Radius != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.Radius))
		i--
		dAtA[i] = 0x40
	}
	if m.Rot != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.Rot))
		i--
		dAtA[i] = 0x38
	}
	if m.Y != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.Y))
		i--
		dAtA[i] = 0x30
	}
	if m.X != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEntity(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.CellID.Size()
		i -= size
		if _, err := m.CellID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEntity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.UserID.Size()
		i -= size
		if _, err := m.UserID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEntity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
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
func NewPopulatedE(r randyEntity, easy bool) *E {
	this := &E{}
	v1 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.ID = *v1
	v2 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.UserID = *v2
	v3 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.CellID = *v3
	this.Name = string(randStringEntity(r))
	this.X = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.X *= -1
	}
	this.Y = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Y *= -1
	}
	this.Rot = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Rot *= -1
	}
	this.Radius = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Radius *= -1
	}
	this.At = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.At *= -1
	}
	v4 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.AnimationID = *v4
	this.AnimationAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.AnimationAt *= -1
	}
	if r.Intn(5) != 0 {
		v5 := r.Intn(5)
		this.StaticObjects = make([]geometry.Rect, v5)
		for i := 0; i < v5; i++ {
			v6 := geometry.NewPopulatedRect(r, easy)
			this.StaticObjects[i] = *v6
		}
	}
	if r.Intn(5) != 0 {
		v7 := r.Intn(5)
		this.DynamicObjects = make([]geometry.Rect, v7)
		for i := 0; i < v7; i++ {
			v8 := geometry.NewPopulatedRect(r, easy)
			this.DynamicObjects[i] = *v8
		}
	}
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
	v9 := r.Intn(100)
	tmps := make([]rune, v9)
	for i := 0; i < v9; i++ {
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
		v10 := r.Int63()
		if r.Intn(2) == 0 {
			v10 *= -1
		}
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(v10))
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
func (m *E) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovEntity(uint64(l))
	l = m.UserID.Size()
	n += 1 + l + sovEntity(uint64(l))
	l = m.CellID.Size()
	n += 1 + l + sovEntity(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEntity(uint64(l))
	}
	if m.X != 0 {
		n += 1 + sovEntity(uint64(m.X))
	}
	if m.Y != 0 {
		n += 1 + sovEntity(uint64(m.Y))
	}
	if m.Rot != 0 {
		n += 1 + sovEntity(uint64(m.Rot))
	}
	if m.Radius != 0 {
		n += 1 + sovEntity(uint64(m.Radius))
	}
	if m.At != 0 {
		n += 1 + sovEntity(uint64(m.At))
	}
	l = m.AnimationID.Size()
	n += 1 + l + sovEntity(uint64(l))
	if m.AnimationAt != 0 {
		n += 1 + sovEntity(uint64(m.AnimationAt))
	}
	if len(m.StaticObjects) > 0 {
		for _, e := range m.StaticObjects {
			l = e.Size()
			n += 1 + l + sovEntity(uint64(l))
		}
	}
	if len(m.DynamicObjects) > 0 {
		for _, e := range m.DynamicObjects {
			l = e.Size()
			n += 1 + l + sovEntity(uint64(l))
		}
	}
	return n
}

func sovEntity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEntity(x uint64) (n int) {
	return sovEntity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *E) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForStaticObjects := "[]Rect{"
	for _, f := range this.StaticObjects {
		repeatedStringForStaticObjects += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForStaticObjects += "}"
	repeatedStringForDynamicObjects := "[]Rect{"
	for _, f := range this.DynamicObjects {
		repeatedStringForDynamicObjects += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForDynamicObjects += "}"
	s := strings.Join([]string{`&E{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`UserID:` + fmt.Sprintf("%v", this.UserID) + `,`,
		`CellID:` + fmt.Sprintf("%v", this.CellID) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`X:` + fmt.Sprintf("%v", this.X) + `,`,
		`Y:` + fmt.Sprintf("%v", this.Y) + `,`,
		`Rot:` + fmt.Sprintf("%v", this.Rot) + `,`,
		`Radius:` + fmt.Sprintf("%v", this.Radius) + `,`,
		`At:` + fmt.Sprintf("%v", this.At) + `,`,
		`AnimationID:` + fmt.Sprintf("%v", this.AnimationID) + `,`,
		`AnimationAt:` + fmt.Sprintf("%v", this.AnimationAt) + `,`,
		`StaticObjects:` + repeatedStringForStaticObjects + `,`,
		`DynamicObjects:` + repeatedStringForDynamicObjects + `,`,
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
func (m *E) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: E: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: E: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
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
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
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
			if err := m.UserID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CellID", wireType)
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
			if err := m.CellID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowEntity
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
				return ErrInvalidLengthEntity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			m.Y = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Y |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rot", wireType)
			}
			m.Rot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rot |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Radius", wireType)
			}
			m.Radius = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Radius |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field At", wireType)
			}
			m.At = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.At |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnimationID", wireType)
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
			if err := m.AnimationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnimationAt", wireType)
			}
			m.AnimationAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnimationAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StaticObjects", wireType)
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
			m.StaticObjects = append(m.StaticObjects, geometry.Rect{})
			if err := m.StaticObjects[len(m.StaticObjects)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DynamicObjects", wireType)
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
			m.DynamicObjects = append(m.DynamicObjects, geometry.Rect{})
			if err := m.DynamicObjects[len(m.DynamicObjects)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
