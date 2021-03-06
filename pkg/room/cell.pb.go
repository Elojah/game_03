// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/room/cell.proto

package room

import (
	fmt "fmt"
	github_com_elojah_game_03_pkg_ulid "github.com/elojah/game_03/pkg/ulid"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
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

type WorldCell struct {
	WorldID github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=WorldID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"WorldID"`
	CellID  github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,opt,name=CellID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"CellID"`
	X       int64                                 `protobuf:"varint,3,opt,name=X,proto3" json:"X,omitempty"`
	Y       int64                                 `protobuf:"varint,4,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *WorldCell) Reset()      { *m = WorldCell{} }
func (*WorldCell) ProtoMessage() {}
func (*WorldCell) Descriptor() ([]byte, []int) {
	return fileDescriptor_8408f5360c24017c, []int{0}
}
func (m *WorldCell) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorldCell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorldCell.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorldCell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorldCell.Merge(m, src)
}
func (m *WorldCell) XXX_Size() int {
	return m.Size()
}
func (m *WorldCell) XXX_DiscardUnknown() {
	xxx_messageInfo_WorldCell.DiscardUnknown(m)
}

var xxx_messageInfo_WorldCell proto.InternalMessageInfo

func (m *WorldCell) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *WorldCell) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Cell struct {
	ID         github_com_elojah_game_03_pkg_ulid.ID           `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"ID"`
	Contiguous map[int32]github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,rep,name=Contiguous,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"Contiguous" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tilemap    github_com_elojah_game_03_pkg_ulid.ID           `protobuf:"bytes,3,opt,name=Tilemap,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"Tilemap"`
	Tileset    github_com_elojah_game_03_pkg_ulid.ID           `protobuf:"bytes,4,opt,name=Tileset,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"Tileset"`
}

func (m *Cell) Reset()      { *m = Cell{} }
func (*Cell) ProtoMessage() {}
func (*Cell) Descriptor() ([]byte, []int) {
	return fileDescriptor_8408f5360c24017c, []int{1}
}
func (m *Cell) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cell.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cell.Merge(m, src)
}
func (m *Cell) XXX_Size() int {
	return m.Size()
}
func (m *Cell) XXX_DiscardUnknown() {
	xxx_messageInfo_Cell.DiscardUnknown(m)
}

var xxx_messageInfo_Cell proto.InternalMessageInfo

func init() {
	proto.RegisterType((*WorldCell)(nil), "room.WorldCell")
	golang_proto.RegisterType((*WorldCell)(nil), "room.WorldCell")
	proto.RegisterType((*Cell)(nil), "room.Cell")
	golang_proto.RegisterType((*Cell)(nil), "room.Cell")
	proto.RegisterMapType((map[int32]github_com_elojah_game_03_pkg_ulid.ID)(nil), "room.Cell.ContiguousEntry")
	golang_proto.RegisterMapType((map[int32]github_com_elojah_game_03_pkg_ulid.ID)(nil), "room.Cell.ContiguousEntry")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/room/cell.proto", fileDescriptor_8408f5360c24017c)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/room/cell.proto", fileDescriptor_8408f5360c24017c)
}

var fileDescriptor_8408f5360c24017c = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0xcb, 0xda, 0x40,
	0x18, 0xc7, 0xf3, 0x24, 0x6a, 0xe9, 0x55, 0x68, 0x09, 0x1d, 0x82, 0xc3, 0xa3, 0x08, 0x05, 0x17,
	0x93, 0x52, 0x97, 0x52, 0x70, 0xd1, 0x48, 0xc9, 0x1a, 0x0a, 0xd5, 0xa9, 0x44, 0xbd, 0xc6, 0xd4,
	0x8b, 0x27, 0x31, 0x29, 0xb8, 0xb5, 0xdf, 0xa0, 0x1f, 0xa3, 0xdf, 0xa0, 0x1d, 0x1d, 0x1d, 0x1d,
	0xa5, 0x83, 0x34, 0x97, 0xa5, 0xa3, 0x63, 0xc7, 0x97, 0x5c, 0x94, 0x57, 0xde, 0xe1, 0x05, 0xdd,
	0x9e, 0xff, 0x3d, 0xfc, 0x7f, 0xf7, 0x3c, 0xff, 0x3b, 0xd2, 0xf6, 0x83, 0x78, 0x96, 0x8c, 0xcd,
	0x09, 0x0f, 0x2d, 0xca, 0xf8, 0x17, 0x6f, 0x66, 0xf9, 0x5e, 0x48, 0x3f, 0xbd, 0xee, 0x58, 0xcb,
	0xb9, 0x6f, 0x45, 0x9c, 0x87, 0xd6, 0x84, 0x32, 0x66, 0x2e, 0x23, 0x1e, 0x73, 0xbd, 0x94, 0x1f,
	0xd4, 0x2e, 0x4d, 0x3e, 0xf7, 0xb9, 0x25, 0x9b, 0xe3, 0xe4, 0xb3, 0x54, 0x52, 0xc8, 0xaa, 0x30,
	0x35, 0x7f, 0x01, 0x79, 0xfa, 0x91, 0x47, 0x6c, 0xda, 0xa7, 0x8c, 0xe9, 0xef, 0xc9, 0x13, 0x29,
	0x1c, 0xdb, 0x80, 0x06, 0xb4, 0xaa, 0xbd, 0xf6, 0xf6, 0x50, 0x57, 0xfe, 0x1c, 0xea, 0xaf, 0x1e,
	0x1f, 0x25, 0x61, 0xc1, 0xd4, 0x74, 0x6c, 0xf7, 0xec, 0xd6, 0x07, 0xa4, 0x92, 0x03, 0x1d, 0xdb,
	0x50, 0x6f, 0xe1, 0x9c, 0xcc, 0x7a, 0x95, 0xc0, 0xd0, 0xd0, 0x1a, 0xd0, 0xd2, 0x5c, 0x18, 0xe6,
	0x6a, 0x64, 0x94, 0x0a, 0x35, 0x6a, 0x7e, 0xd7, 0x48, 0x49, 0x0e, 0xdd, 0x25, 0xea, 0xad, 0xf3,
	0xaa, 0x8e, 0xad, 0x07, 0x84, 0xf4, 0xf9, 0x22, 0x0e, 0xfc, 0x84, 0x27, 0x2b, 0x43, 0x6d, 0x68,
	0xad, 0x67, 0x6f, 0x6a, 0x66, 0x9e, 0xa5, 0x99, 0xe3, 0xcd, 0xfb, 0xe6, 0x60, 0x11, 0x47, 0xeb,
	0x6b, 0xaf, 0xb8, 0x80, 0xe7, 0xf1, 0x7e, 0x08, 0x18, 0x0d, 0xbd, 0xa5, 0x5c, 0xea, 0xfa, 0x78,
	0x4f, 0xee, 0x33, 0x68, 0x45, 0x63, 0x99, 0xc7, 0x6d, 0xa0, 0x15, 0x8d, 0x6b, 0x5d, 0xf2, 0xfc,
	0xc1, 0x7e, 0xfa, 0x0b, 0xa2, 0xcd, 0xe9, 0x5a, 0xe6, 0x59, 0x76, 0xf3, 0x52, 0x7f, 0x49, 0xca,
	0x5f, 0x3d, 0x96, 0xd0, 0xe2, 0x2d, 0xdd, 0x42, 0xbc, 0x53, 0xdf, 0x42, 0xcf, 0xde, 0xa5, 0xa8,
	0xec, 0x53, 0x54, 0x8e, 0x29, 0xc2, 0xff, 0x14, 0xe1, 0x9b, 0x40, 0xf8, 0x29, 0x10, 0x7e, 0x0b,
	0x84, 0x8d, 0x40, 0xd8, 0x0a, 0x84, 0x9d, 0x40, 0xf8, 0x2b, 0x10, 0xfe, 0x09, 0x54, 0x8e, 0x02,
	0xe1, 0x47, 0x86, 0xca, 0x26, 0x43, 0xd8, 0x65, 0xa8, 0xec, 0x33, 0x54, 0xc6, 0x15, 0xf9, 0x15,
	0x3b, 0x77, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0x9d, 0xa6, 0xfb, 0xf0, 0x02, 0x00, 0x00,
}

func (this *WorldCell) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorldCell)
	if !ok {
		that2, ok := that.(WorldCell)
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
	if !this.CellID.Equal(that1.CellID) {
		return false
	}
	if this.X != that1.X {
		return false
	}
	if this.Y != that1.Y {
		return false
	}
	return true
}
func (this *Cell) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Cell)
	if !ok {
		that2, ok := that.(Cell)
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
	if len(this.Contiguous) != len(that1.Contiguous) {
		return false
	}
	for i := range this.Contiguous {
		if !this.Contiguous[i].Equal(that1.Contiguous[i]) { //not nullable
			return false
		}
	}
	if !this.Tilemap.Equal(that1.Tilemap) {
		return false
	}
	if !this.Tileset.Equal(that1.Tileset) {
		return false
	}
	return true
}
func (this *WorldCell) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&room.WorldCell{")
	s = append(s, "WorldID: "+fmt.Sprintf("%#v", this.WorldID)+",\n")
	s = append(s, "CellID: "+fmt.Sprintf("%#v", this.CellID)+",\n")
	s = append(s, "X: "+fmt.Sprintf("%#v", this.X)+",\n")
	s = append(s, "Y: "+fmt.Sprintf("%#v", this.Y)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Cell) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&room.Cell{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	keysForContiguous := make([]int32, 0, len(this.Contiguous))
	for k, _ := range this.Contiguous {
		keysForContiguous = append(keysForContiguous, k)
	}
	github_com_gogo_protobuf_sortkeys.Int32s(keysForContiguous)
	mapStringForContiguous := "map[int32]github_com_elojah_game_03_pkg_ulid.ID{"
	for _, k := range keysForContiguous {
		mapStringForContiguous += fmt.Sprintf("%#v: %#v,", k, this.Contiguous[k])
	}
	mapStringForContiguous += "}"
	if this.Contiguous != nil {
		s = append(s, "Contiguous: "+mapStringForContiguous+",\n")
	}
	s = append(s, "Tilemap: "+fmt.Sprintf("%#v", this.Tilemap)+",\n")
	s = append(s, "Tileset: "+fmt.Sprintf("%#v", this.Tileset)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCell(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *WorldCell) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorldCell) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorldCell) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Y != 0 {
		i = encodeVarintCell(dAtA, i, uint64(m.Y))
		i--
		dAtA[i] = 0x20
	}
	if m.X != 0 {
		i = encodeVarintCell(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.CellID.Size()
		i -= size
		if _, err := m.CellID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCell(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.WorldID.Size()
		i -= size
		if _, err := m.WorldID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCell(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Cell) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cell) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cell) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Tileset.Size()
		i -= size
		if _, err := m.Tileset.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCell(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Tilemap.Size()
		i -= size
		if _, err := m.Tilemap.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCell(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Contiguous) > 0 {
		for k := range m.Contiguous {
			v := m.Contiguous[k]
			baseI := i
			{
				size := v.Size()
				i -= size
				if _, err := v.MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintCell(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i = encodeVarintCell(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintCell(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCell(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintCell(dAtA []byte, offset int, v uint64) int {
	offset -= sovCell(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedWorldCell(r randyCell, easy bool) *WorldCell {
	this := &WorldCell{}
	v1 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.WorldID = *v1
	v2 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.CellID = *v2
	this.X = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.X *= -1
	}
	this.Y = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Y *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedCell(r randyCell, easy bool) *Cell {
	this := &Cell{}
	v3 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.ID = *v3
	if r.Intn(5) != 0 {
		v4 := r.Intn(10)
		this.Contiguous = make(map[int32]github_com_elojah_game_03_pkg_ulid.ID)
		for i := 0; i < v4; i++ {
			this.Contiguous[int32(r.Int31())] = (github_com_elojah_game_03_pkg_ulid.ID)(*github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r))
		}
	}
	v5 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.Tilemap = *v5
	v6 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.Tileset = *v6
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyCell interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneCell(r randyCell) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringCell(r randyCell) string {
	v7 := r.Intn(100)
	tmps := make([]rune, v7)
	for i := 0; i < v7; i++ {
		tmps[i] = randUTF8RuneCell(r)
	}
	return string(tmps)
}
func randUnrecognizedCell(r randyCell, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldCell(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldCell(dAtA []byte, r randyCell, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateCell(dAtA, uint64(key))
		v8 := r.Int63()
		if r.Intn(2) == 0 {
			v8 *= -1
		}
		dAtA = encodeVarintPopulateCell(dAtA, uint64(v8))
	case 1:
		dAtA = encodeVarintPopulateCell(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateCell(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateCell(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateCell(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateCell(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *WorldCell) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.WorldID.Size()
	n += 1 + l + sovCell(uint64(l))
	l = m.CellID.Size()
	n += 1 + l + sovCell(uint64(l))
	if m.X != 0 {
		n += 1 + sovCell(uint64(m.X))
	}
	if m.Y != 0 {
		n += 1 + sovCell(uint64(m.Y))
	}
	return n
}

func (m *Cell) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovCell(uint64(l))
	if len(m.Contiguous) > 0 {
		for k, v := range m.Contiguous {
			_ = k
			_ = v
			l = 0
			l = v.Size()
			l += 1 + sovCell(uint64(l))
			mapEntrySize := 1 + sovCell(uint64(k)) + l
			n += mapEntrySize + 1 + sovCell(uint64(mapEntrySize))
		}
	}
	l = m.Tilemap.Size()
	n += 1 + l + sovCell(uint64(l))
	l = m.Tileset.Size()
	n += 1 + l + sovCell(uint64(l))
	return n
}

func sovCell(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCell(x uint64) (n int) {
	return sovCell(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *WorldCell) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WorldCell{`,
		`WorldID:` + fmt.Sprintf("%v", this.WorldID) + `,`,
		`CellID:` + fmt.Sprintf("%v", this.CellID) + `,`,
		`X:` + fmt.Sprintf("%v", this.X) + `,`,
		`Y:` + fmt.Sprintf("%v", this.Y) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Cell) String() string {
	if this == nil {
		return "nil"
	}
	keysForContiguous := make([]int32, 0, len(this.Contiguous))
	for k, _ := range this.Contiguous {
		keysForContiguous = append(keysForContiguous, k)
	}
	github_com_gogo_protobuf_sortkeys.Int32s(keysForContiguous)
	mapStringForContiguous := "map[int32]github_com_elojah_game_03_pkg_ulid.ID{"
	for _, k := range keysForContiguous {
		mapStringForContiguous += fmt.Sprintf("%v: %v,", k, this.Contiguous[k])
	}
	mapStringForContiguous += "}"
	s := strings.Join([]string{`&Cell{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Contiguous:` + mapStringForContiguous + `,`,
		`Tilemap:` + fmt.Sprintf("%v", this.Tilemap) + `,`,
		`Tileset:` + fmt.Sprintf("%v", this.Tileset) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCell(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *WorldCell) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCell
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
			return fmt.Errorf("proto: WorldCell: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorldCell: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorldID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCell
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
				return ErrInvalidLengthCell
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCell
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
				return fmt.Errorf("proto: wrong wireType = %d for field CellID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCell
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
				return ErrInvalidLengthCell
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCell
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CellID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCell
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			m.Y = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCell
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
		default:
			iNdEx = preIndex
			skippy, err := skipCell(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCell
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
func (m *Cell) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCell
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
			return fmt.Errorf("proto: Cell: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cell: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCell
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
				return ErrInvalidLengthCell
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCell
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
				return fmt.Errorf("proto: wrong wireType = %d for field Contiguous", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCell
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
				return ErrInvalidLengthCell
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCell
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Contiguous == nil {
				m.Contiguous = make(map[int32]github_com_elojah_game_03_pkg_ulid.ID)
			}
			var mapkey int32
			var mapvalue1 github_com_elojah_game_03_pkg_ulid.ID
			var mapvalue = &mapvalue1
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCell
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCell
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCell
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthCell
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex < 0 {
						return ErrInvalidLengthCell
					}
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postbytesIndex]); err != nil {
						return err
					}
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCell(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthCell
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Contiguous[mapkey] = ((github_com_elojah_game_03_pkg_ulid.ID)(*mapvalue))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tilemap", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCell
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
				return ErrInvalidLengthCell
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCell
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tilemap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tileset", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCell
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
				return ErrInvalidLengthCell
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCell
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tileset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCell(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCell
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
func skipCell(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCell
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
					return 0, ErrIntOverflowCell
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
					return 0, ErrIntOverflowCell
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
				return 0, ErrInvalidLengthCell
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCell
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCell
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCell        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCell          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCell = fmt.Errorf("proto: unexpected end of group")
)
