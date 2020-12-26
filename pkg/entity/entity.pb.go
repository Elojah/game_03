// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/entity/entity.proto

package entity

import (
	fmt "fmt"
	geometry "github.com/elojah/game_03/pkg/geometry"
	space "github.com/elojah/game_03/pkg/space"
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
	ID      github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"ID"`
	Section space.Section                         `protobuf:"bytes,2,opt,name=Section,proto3" json:"Section"`
	Coord   geometry.Vec2                         `protobuf:"bytes,3,opt,name=Coord,proto3" json:"Coord"`
	MaxHP   uint64                                `protobuf:"varint,4,opt,name=MaxHP,proto3" json:"MaxHP,omitempty"`
	HP      uint64                                `protobuf:"varint,5,opt,name=HP,proto3" json:"HP,omitempty"`
	MaxMP   uint64                                `protobuf:"varint,6,opt,name=MaxMP,proto3" json:"MaxMP,omitempty"`
	MP      uint64                                `protobuf:"varint,7,opt,name=MP,proto3" json:"MP,omitempty"`
	State   github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,8,opt,name=State,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"State"`
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

func (m *E) GetSection() space.Section {
	if m != nil {
		return m.Section
	}
	return space.Section{}
}

func (m *E) GetCoord() geometry.Vec2 {
	if m != nil {
		return m.Coord
	}
	return geometry.Vec2{}
}

func (m *E) GetMaxHP() uint64 {
	if m != nil {
		return m.MaxHP
	}
	return 0
}

func (m *E) GetHP() uint64 {
	if m != nil {
		return m.HP
	}
	return 0
}

func (m *E) GetMaxMP() uint64 {
	if m != nil {
		return m.MaxMP
	}
	return 0
}

func (m *E) GetMP() uint64 {
	if m != nil {
		return m.MP
	}
	return 0
}

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
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x31, 0x6f, 0xda, 0x40,
	0x14, 0xc7, 0xfd, 0x5c, 0x0c, 0xd5, 0xb5, 0x62, 0xb0, 0x3a, 0x9c, 0x18, 0x1e, 0xa8, 0x52, 0x25,
	0x54, 0x09, 0xbb, 0x05, 0x75, 0xec, 0x02, 0xae, 0x04, 0x83, 0x25, 0x0b, 0xa4, 0xae, 0x91, 0x31,
	0x17, 0xe3, 0x04, 0x38, 0xe4, 0x1c, 0x4a, 0xd8, 0xf2, 0x11, 0xf2, 0x31, 0xb2, 0x65, 0xcd, 0xc8,
	0xc8, 0xc8, 0x88, 0x32, 0xa0, 0xf8, 0xbc, 0x64, 0x64, 0xcc, 0x18, 0xe5, 0xec, 0xa0, 0x4c, 0x0c,
	0x99, 0xfc, 0xfe, 0xcf, 0xbf, 0xdf, 0x7b, 0x7a, 0x47, 0x7e, 0x87, 0x91, 0x18, 0x2f, 0x86, 0x56,
	0xc0, 0xa7, 0x36, 0x9b, 0xf0, 0x33, 0x7f, 0x6c, 0x87, 0xfe, 0x94, 0x9d, 0xfc, 0x6a, 0xd9, 0xf3,
	0xf3, 0xd0, 0x66, 0x33, 0x11, 0x89, 0x65, 0xfe, 0xb1, 0xe6, 0x31, 0x17, 0xdc, 0x2c, 0x66, 0xa9,
	0xd2, 0x78, 0xa7, 0x86, 0x3c, 0xe4, 0xb6, 0xfa, 0x3d, 0x5c, 0x9c, 0xaa, 0xa4, 0x82, 0xaa, 0x32,
	0xad, 0xf2, 0xe7, 0xf8, 0xa6, 0x90, 0xf1, 0x29, 0x13, 0xf1, 0xf2, 0x50, 0xe4, 0x9a, 0x7d, 0x5c,
	0xbb, 0x98, 0xfb, 0x01, 0xb3, 0x2f, 0x79, 0x3c, 0x19, 0x65, 0xc2, 0xf7, 0x3b, 0x9d, 0xc0, 0x3f,
	0xf3, 0x2f, 0xd1, 0x7b, 0x0e, 0x85, 0x1a, 0xd4, 0xbf, 0xb6, 0x1b, 0xeb, 0x5d, 0x55, 0x7b, 0xd8,
	0x55, 0x7f, 0x1c, 0x1f, 0xb5, 0x98, 0x44, 0x23, 0xab, 0xe7, 0xf4, 0xf5, 0x9e, 0x63, 0x5a, 0xa4,
	0x34, 0x60, 0x81, 0x88, 0xf8, 0x8c, 0xea, 0x35, 0xa8, 0x7f, 0x69, 0x96, 0x2d, 0xb5, 0xc9, 0xca,
	0xbb, 0xed, 0xc2, 0xeb, 0xcc, 0xfe, 0x1b, 0x64, 0xfe, 0x24, 0x46, 0x87, 0xf3, 0x78, 0x44, 0x3f,
	0xe5, 0xf4, 0xe1, 0x8a, 0xff, 0x2c, 0x68, 0xe6, 0x74, 0x86, 0x98, 0xdf, 0x88, 0xe1, 0xfa, 0x57,
	0x5d, 0x8f, 0x16, 0x6a, 0x50, 0x2f, 0xf4, 0xb3, 0x60, 0x96, 0x89, 0xde, 0xf5, 0xa8, 0xa1, 0x5a,
	0x7a, 0xd7, 0xcb, 0x29, 0xd7, 0xa3, 0xc5, 0x03, 0xe5, 0x2a, 0xca, 0xf5, 0x68, 0x29, 0xa3, 0x5c,
	0xcf, 0xec, 0x10, 0x63, 0x20, 0x7c, 0xc1, 0xe8, 0xe7, 0x8f, 0x5c, 0x9a, 0xb9, 0x6d, 0x67, 0x93,
	0xa0, 0xb6, 0x4d, 0x50, 0xdb, 0x27, 0x08, 0xcf, 0x09, 0xc2, 0xb5, 0x44, 0xb8, 0x95, 0x08, 0xf7,
	0x12, 0x61, 0x25, 0x11, 0xd6, 0x12, 0x61, 0x23, 0x11, 0x1e, 0x25, 0xc2, 0x93, 0x44, 0x6d, 0x2f,
	0x11, 0x6e, 0x52, 0xd4, 0x56, 0x29, 0xc2, 0x26, 0x45, 0x6d, 0x9b, 0xa2, 0x36, 0x2c, 0xaa, 0xe7,
	0x6f, 0xbd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x14, 0x35, 0x46, 0xdb, 0x52, 0x02, 0x00, 0x00,
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
	if !this.Section.Equal(&that1.Section) {
		return false
	}
	if !this.Coord.Equal(&that1.Coord) {
		return false
	}
	if this.MaxHP != that1.MaxHP {
		return false
	}
	if this.HP != that1.HP {
		return false
	}
	if this.MaxMP != that1.MaxMP {
		return false
	}
	if this.MP != that1.MP {
		return false
	}
	if !this.State.Equal(that1.State) {
		return false
	}
	return true
}
func (this *E) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 12)
	s = append(s, "&entity.E{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Section: "+strings.Replace(this.Section.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "Coord: "+strings.Replace(this.Coord.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "MaxHP: "+fmt.Sprintf("%#v", this.MaxHP)+",\n")
	s = append(s, "HP: "+fmt.Sprintf("%#v", this.HP)+",\n")
	s = append(s, "MaxMP: "+fmt.Sprintf("%#v", this.MaxMP)+",\n")
	s = append(s, "MP: "+fmt.Sprintf("%#v", this.MP)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
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
	{
		size := m.State.Size()
		i -= size
		if _, err := m.State.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEntity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.MP != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.MP))
		i--
		dAtA[i] = 0x38
	}
	if m.MaxMP != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.MaxMP))
		i--
		dAtA[i] = 0x30
	}
	if m.HP != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.HP))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxHP != 0 {
		i = encodeVarintEntity(dAtA, i, uint64(m.MaxHP))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Coord.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEntity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Section.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
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
	v2 := space.NewPopulatedSection(r, easy)
	this.Section = *v2
	v3 := geometry.NewPopulatedVec2(r, easy)
	this.Coord = *v3
	this.MaxHP = uint64(uint64(r.Uint32()))
	this.HP = uint64(uint64(r.Uint32()))
	this.MaxMP = uint64(uint64(r.Uint32()))
	this.MP = uint64(uint64(r.Uint32()))
	v4 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.State = *v4
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
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
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
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateEntity(dAtA, uint64(v6))
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
	l = m.Section.Size()
	n += 1 + l + sovEntity(uint64(l))
	l = m.Coord.Size()
	n += 1 + l + sovEntity(uint64(l))
	if m.MaxHP != 0 {
		n += 1 + sovEntity(uint64(m.MaxHP))
	}
	if m.HP != 0 {
		n += 1 + sovEntity(uint64(m.HP))
	}
	if m.MaxMP != 0 {
		n += 1 + sovEntity(uint64(m.MaxMP))
	}
	if m.MP != 0 {
		n += 1 + sovEntity(uint64(m.MP))
	}
	l = m.State.Size()
	n += 1 + l + sovEntity(uint64(l))
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
	s := strings.Join([]string{`&E{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Section:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Section), "Section", "space.Section", 1), `&`, ``, 1) + `,`,
		`Coord:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Coord), "Vec2", "geometry.Vec2", 1), `&`, ``, 1) + `,`,
		`MaxHP:` + fmt.Sprintf("%v", this.MaxHP) + `,`,
		`HP:` + fmt.Sprintf("%v", this.HP) + `,`,
		`MaxMP:` + fmt.Sprintf("%v", this.MaxMP) + `,`,
		`MP:` + fmt.Sprintf("%v", this.MP) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
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
				return fmt.Errorf("proto: wrong wireType = %d for field Section", wireType)
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
			if err := m.Section.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coord", wireType)
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
			if err := m.Coord.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxHP", wireType)
			}
			m.MaxHP = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxHP |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HP", wireType)
			}
			m.HP = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HP |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMP", wireType)
			}
			m.MaxMP = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMP |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MP", wireType)
			}
			m.MP = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MP |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
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
			if err := m.State.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
