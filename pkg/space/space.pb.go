// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/space/space.proto

package space

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
	strconv "strconv"
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

type Cross int32

const (
	In  Cross = 0
	Out Cross = 1
)

var Cross_name = map[int32]string{
	0: "In",
	1: "Out",
}

var Cross_value = map[string]int32{
	"In":  0,
	"Out": 1,
}

func (Cross) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dfc7dda2ffab8cb5, []int{0}
}

type World struct {
	ID         github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"ID"`
	OwnerID    github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,opt,name=OwnerID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"OwnerID"`
	NSection   geometry.Vec2                         `protobuf:"bytes,3,opt,name=NSection,proto3" json:"NSection"`
	DimSection geometry.Vec2                         `protobuf:"bytes,4,opt,name=DimSection,proto3" json:"DimSection"`
}

func (m *World) Reset()      { *m = World{} }
func (*World) ProtoMessage() {}
func (*World) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfc7dda2ffab8cb5, []int{0}
}
func (m *World) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *World) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_World.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *World) XXX_Merge(src proto.Message) {
	xxx_messageInfo_World.Merge(m, src)
}
func (m *World) XXX_Size() int {
	return m.Size()
}
func (m *World) XXX_DiscardUnknown() {
	xxx_messageInfo_World.DiscardUnknown(m)
}

var xxx_messageInfo_World proto.InternalMessageInfo

func (m *World) GetNSection() geometry.Vec2 {
	if m != nil {
		return m.NSection
	}
	return geometry.Vec2{}
}

func (m *World) GetDimSection() geometry.Vec2 {
	if m != nil {
		return m.DimSection
	}
	return geometry.Vec2{}
}

type Section struct {
	WorldID github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=WorldID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"WorldID"`
	Coord   geometry.Vec2                         `protobuf:"bytes,2,opt,name=Coord,proto3" json:"Coord"`
}

func (m *Section) Reset()      { *m = Section{} }
func (*Section) ProtoMessage() {}
func (*Section) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfc7dda2ffab8cb5, []int{1}
}
func (m *Section) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Section) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Section.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Section) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Section.Merge(m, src)
}
func (m *Section) XXX_Size() int {
	return m.Size()
}
func (m *Section) XXX_DiscardUnknown() {
	xxx_messageInfo_Section.DiscardUnknown(m)
}

var xxx_messageInfo_Section proto.InternalMessageInfo

func (m *Section) GetCoord() geometry.Vec2 {
	if m != nil {
		return m.Coord
	}
	return geometry.Vec2{}
}

type Entity struct {
	ID      github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"ID"`
	Section Section                               `protobuf:"bytes,2,opt,name=Section,proto3" json:"Section"`
	Cross   Cross                                 `protobuf:"varint,3,opt,name=Cross,proto3,enum=space.Cross" json:"Cross,omitempty"`
	State   github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,4,opt,name=State,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"State"`
}

func (m *Entity) Reset()      { *m = Entity{} }
func (*Entity) ProtoMessage() {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfc7dda2ffab8cb5, []int{2}
}
func (m *Entity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return m.Size()
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetSection() Section {
	if m != nil {
		return m.Section
	}
	return Section{}
}

func (m *Entity) GetCross() Cross {
	if m != nil {
		return m.Cross
	}
	return In
}

func init() {
	proto.RegisterEnum("space.Cross", Cross_name, Cross_value)
	golang_proto.RegisterEnum("space.Cross", Cross_name, Cross_value)
	proto.RegisterType((*World)(nil), "space.World")
	golang_proto.RegisterType((*World)(nil), "space.World")
	proto.RegisterType((*Section)(nil), "space.Section")
	golang_proto.RegisterType((*Section)(nil), "space.Section")
	proto.RegisterType((*Entity)(nil), "space.Entity")
	golang_proto.RegisterType((*Entity)(nil), "space.Entity")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/space/space.proto", fileDescriptor_dfc7dda2ffab8cb5)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/space/space.proto", fileDescriptor_dfc7dda2ffab8cb5)
}

var fileDescriptor_dfc7dda2ffab8cb5 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0x6a, 0xdc, 0x40,
	0x10, 0xc6, 0x77, 0x64, 0xeb, 0x14, 0x36, 0xc6, 0x98, 0xad, 0x84, 0x8b, 0xb1, 0x11, 0x04, 0x8c,
	0xc1, 0x92, 0x39, 0x27, 0x65, 0x9a, 0x3b, 0x05, 0xa3, 0x26, 0x86, 0x33, 0x24, 0x65, 0xd0, 0xe9,
	0x36, 0xb2, 0x92, 0x93, 0xf6, 0xd0, 0xad, 0x08, 0x6e, 0x42, 0x1e, 0x21, 0x8f, 0x91, 0x47, 0x48,
	0xe9, 0xd2, 0xe5, 0x95, 0x26, 0x85, 0xb1, 0x56, 0x4d, 0xca, 0xeb, 0x92, 0x32, 0x64, 0xf5, 0x87,
	0x6b, 0x72, 0x90, 0x73, 0xb3, 0xec, 0xec, 0xce, 0x6f, 0xbe, 0x99, 0x8f, 0xa1, 0x5e, 0x9c, 0xc8,
	0xab, 0x62, 0xec, 0x46, 0x22, 0xf5, 0xf8, 0x54, 0x7c, 0x08, 0xaf, 0xbc, 0x38, 0x4c, 0xf9, 0xbb,
	0xd3, 0x33, 0x6f, 0xf6, 0x31, 0xf6, 0xe6, 0xb3, 0x30, 0xe2, 0xf5, 0xe9, 0xce, 0x72, 0x21, 0x05,
	0x33, 0x75, 0xb0, 0x7f, 0xb2, 0xc2, 0xc5, 0x22, 0x16, 0x9e, 0xfe, 0x1d, 0x17, 0xef, 0x75, 0xa4,
	0x03, 0x7d, 0xab, 0xa9, 0xfd, 0x17, 0xeb, 0x65, 0x62, 0x2e, 0x52, 0x2e, 0xf3, 0xeb, 0xee, 0x52,
	0x63, 0xce, 0x2f, 0xa0, 0xe6, 0x5b, 0x91, 0x4f, 0x27, 0xec, 0x25, 0x35, 0x02, 0xdf, 0x86, 0x43,
	0x38, 0xda, 0x19, 0x9c, 0xdc, 0xde, 0x1f, 0x90, 0x1f, 0xf7, 0x07, 0xcf, 0xd6, 0x17, 0x2d, 0xa6,
	0xc9, 0xc4, 0x0d, 0xfc, 0x91, 0x11, 0xf8, 0xec, 0x9c, 0x5a, 0x17, 0x9f, 0x32, 0x9e, 0x07, 0xbe,
	0x6d, 0x6c, 0x52, 0xa3, 0xa5, 0xd9, 0x29, 0x7d, 0xf2, 0xfa, 0x92, 0x47, 0x32, 0x11, 0x99, 0xbd,
	0x75, 0x08, 0x47, 0x4f, 0xfb, 0xbb, 0x6e, 0xd7, 0xf4, 0x1b, 0x1e, 0xf5, 0x07, 0xdb, 0x7f, 0x2b,
	0x8f, 0xba, 0x2c, 0xf6, 0x9c, 0x52, 0x3f, 0x49, 0x5b, 0x66, 0x7b, 0x0d, 0xb3, 0x92, 0xe7, 0x7c,
	0xa6, 0x56, 0x5b, 0xe0, 0x9c, 0x5a, 0xda, 0x83, 0x4d, 0xe7, 0x6f, 0x69, 0x76, 0x4c, 0xcd, 0xa1,
	0x10, 0xf9, 0x44, 0x5b, 0xf0, 0xaf, 0x26, 0xea, 0x14, 0xe7, 0x01, 0x68, 0xef, 0x55, 0x26, 0x13,
	0x79, 0xfd, 0x58, 0xeb, 0xdd, 0x6e, 0x92, 0x4e, 0xb7, 0xde, 0xa7, 0xe6, 0xb5, 0xd1, 0xed, 0xc6,
	0x75, 0xa8, 0x39, 0xcc, 0xc5, 0x7c, 0xae, 0xed, 0xdd, 0xed, 0xef, 0x34, 0xd9, 0xfa, 0x6d, 0x54,
	0x7f, 0xb1, 0x21, 0x35, 0x2f, 0x65, 0x28, 0xb9, 0xb6, 0xf3, 0xbf, 0xbb, 0xaa, 0xd9, 0x63, 0xbb,
	0x11, 0x62, 0x3d, 0x6a, 0x04, 0xd9, 0x1e, 0x61, 0x16, 0xdd, 0xba, 0x28, 0xe4, 0x1e, 0x0c, 0xfc,
	0x45, 0x89, 0xe4, 0xae, 0x44, 0xb2, 0x2c, 0x11, 0x7e, 0x97, 0x08, 0x5f, 0x14, 0xc2, 0x37, 0x85,
	0xf0, 0x5d, 0x21, 0xdc, 0x28, 0x84, 0x5b, 0x85, 0xb0, 0x50, 0x08, 0x0f, 0x0a, 0xe1, 0xa7, 0x42,
	0xb2, 0x54, 0x08, 0x5f, 0x2b, 0x24, 0x37, 0x15, 0xc2, 0xa2, 0x42, 0x72, 0x57, 0x21, 0x19, 0xf7,
	0xf4, 0x0e, 0x9f, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xb9, 0xb6, 0x00, 0x63, 0x03, 0x00,
	0x00,
}

func (x Cross) String() string {
	s, ok := Cross_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *World) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*World)
	if !ok {
		that2, ok := that.(World)
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
	if !this.OwnerID.Equal(that1.OwnerID) {
		return false
	}
	if !this.NSection.Equal(&that1.NSection) {
		return false
	}
	if !this.DimSection.Equal(&that1.DimSection) {
		return false
	}
	return true
}
func (this *Section) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Section)
	if !ok {
		that2, ok := that.(Section)
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
	if !this.WorldID.Equal(that1.WorldID) {
		return false
	}
	if !this.Coord.Equal(&that1.Coord) {
		return false
	}
	return true
}
func (this *Entity) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Entity)
	if !ok {
		that2, ok := that.(Entity)
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
	if !this.Section.Equal(&that1.Section) {
		return false
	}
	if this.Cross != that1.Cross {
		return false
	}
	if !this.State.Equal(that1.State) {
		return false
	}
	return true
}
func (this *World) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&space.World{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "OwnerID: "+fmt.Sprintf("%#v", this.OwnerID)+",\n")
	s = append(s, "NSection: "+strings.Replace(this.NSection.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "DimSection: "+strings.Replace(this.DimSection.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Section) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&space.Section{")
	s = append(s, "WorldID: "+fmt.Sprintf("%#v", this.WorldID)+",\n")
	s = append(s, "Coord: "+strings.Replace(this.Coord.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Entity) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&space.Entity{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Section: "+strings.Replace(this.Section.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "Cross: "+fmt.Sprintf("%#v", this.Cross)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSpace(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *World) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *World) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *World) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DimSection.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSpace(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.NSection.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSpace(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.OwnerID.Size()
		i -= size
		if _, err := m.OwnerID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSpace(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSpace(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Section) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Section) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Section) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coord.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSpace(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.WorldID.Size()
		i -= size
		if _, err := m.WorldID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSpace(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Entity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Entity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.State.Size()
		i -= size
		if _, err := m.State.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSpace(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Cross != 0 {
		i = encodeVarintSpace(dAtA, i, uint64(m.Cross))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Section.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSpace(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSpace(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintSpace(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedWorld(r randySpace, easy bool) *World {
	this := &World{}
	v1 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.ID = *v1
	v2 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.OwnerID = *v2
	v3 := geometry.NewPopulatedVec2(r, easy)
	this.NSection = *v3
	v4 := geometry.NewPopulatedVec2(r, easy)
	this.DimSection = *v4
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSection(r randySpace, easy bool) *Section {
	this := &Section{}
	v5 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.WorldID = *v5
	v6 := geometry.NewPopulatedVec2(r, easy)
	this.Coord = *v6
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedEntity(r randySpace, easy bool) *Entity {
	this := &Entity{}
	v7 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.ID = *v7
	v8 := NewPopulatedSection(r, easy)
	this.Section = *v8
	this.Cross = Cross([]int32{0, 1}[r.Intn(2)])
	v9 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.State = *v9
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randySpace interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSpace(r randySpace) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringSpace(r randySpace) string {
	v10 := r.Intn(100)
	tmps := make([]rune, v10)
	for i := 0; i < v10; i++ {
		tmps[i] = randUTF8RuneSpace(r)
	}
	return string(tmps)
}
func randUnrecognizedSpace(r randySpace, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldSpace(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldSpace(dAtA []byte, r randySpace, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateSpace(dAtA, uint64(key))
		v11 := r.Int63()
		if r.Intn(2) == 0 {
			v11 *= -1
		}
		dAtA = encodeVarintPopulateSpace(dAtA, uint64(v11))
	case 1:
		dAtA = encodeVarintPopulateSpace(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateSpace(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateSpace(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateSpace(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateSpace(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *World) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovSpace(uint64(l))
	l = m.OwnerID.Size()
	n += 1 + l + sovSpace(uint64(l))
	l = m.NSection.Size()
	n += 1 + l + sovSpace(uint64(l))
	l = m.DimSection.Size()
	n += 1 + l + sovSpace(uint64(l))
	return n
}

func (m *Section) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.WorldID.Size()
	n += 1 + l + sovSpace(uint64(l))
	l = m.Coord.Size()
	n += 1 + l + sovSpace(uint64(l))
	return n
}

func (m *Entity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovSpace(uint64(l))
	l = m.Section.Size()
	n += 1 + l + sovSpace(uint64(l))
	if m.Cross != 0 {
		n += 1 + sovSpace(uint64(m.Cross))
	}
	l = m.State.Size()
	n += 1 + l + sovSpace(uint64(l))
	return n
}

func sovSpace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpace(x uint64) (n int) {
	return sovSpace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *World) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&World{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`OwnerID:` + fmt.Sprintf("%v", this.OwnerID) + `,`,
		`NSection:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.NSection), "Vec2", "geometry.Vec2", 1), `&`, ``, 1) + `,`,
		`DimSection:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.DimSection), "Vec2", "geometry.Vec2", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Section) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Section{`,
		`WorldID:` + fmt.Sprintf("%v", this.WorldID) + `,`,
		`Coord:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Coord), "Vec2", "geometry.Vec2", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Entity) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Entity{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Section:` + strings.Replace(strings.Replace(this.Section.String(), "Section", "Section", 1), `&`, ``, 1) + `,`,
		`Cross:` + fmt.Sprintf("%v", this.Cross) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSpace(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *World) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpace
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
			return fmt.Errorf("proto: World: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: World: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
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
				return ErrInvalidLengthSpace
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpace
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
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
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
				return ErrInvalidLengthSpace
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OwnerID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NSection", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
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
				return ErrInvalidLengthSpace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NSection.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DimSection", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
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
				return ErrInvalidLengthSpace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DimSection.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpace
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSpace
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
func (m *Section) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpace
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
			return fmt.Errorf("proto: Section: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Section: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorldID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
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
				return ErrInvalidLengthSpace
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WorldID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coord", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
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
				return ErrInvalidLengthSpace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coord.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpace
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSpace
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
func (m *Entity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpace
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
			return fmt.Errorf("proto: Entity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
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
				return ErrInvalidLengthSpace
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpace
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
				return fmt.Errorf("proto: wrong wireType = %d for field Section", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
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
				return ErrInvalidLengthSpace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Section.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cross", wireType)
			}
			m.Cross = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cross |= Cross(b&0x7F) << shift
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
					return ErrIntOverflowSpace
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
				return ErrInvalidLengthSpace
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.State.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpace
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSpace
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
func skipSpace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpace
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
					return 0, ErrIntOverflowSpace
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
					return 0, ErrIntOverflowSpace
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
				return 0, ErrInvalidLengthSpace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpace = fmt.Errorf("proto: unexpected end of group")
)
