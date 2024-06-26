// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/faction/faction.proto

package faction

import (
	fmt "fmt"
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

type F struct {
	ID      github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"ID"`
	WorldID github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,opt,name=WorldID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"WorldID"`
	Name    string                                `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Icon    string                                `protobuf:"bytes,4,opt,name=Icon,proto3" json:"Icon,omitempty"`
	Max     int64                                 `protobuf:"varint,5,opt,name=Max,proto3" json:"Max,omitempty"`
}

func (m *F) Reset()      { *m = F{} }
func (*F) ProtoMessage() {}
func (*F) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e78b3c9dcf99c, []int{0}
}
func (m *F) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *F) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_F.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *F) XXX_Merge(src proto.Message) {
	xxx_messageInfo_F.Merge(m, src)
}
func (m *F) XXX_Size() int {
	return m.Size()
}
func (m *F) XXX_DiscardUnknown() {
	xxx_messageInfo_F.DiscardUnknown(m)
}

var xxx_messageInfo_F proto.InternalMessageInfo

type PC struct {
	ID         github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"ID"`
	FactionID  github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,opt,name=FactionID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"FactionID"`
	Permission int64                                 `protobuf:"varint,3,opt,name=Permission,proto3" json:"Permission,omitempty"`
}

func (m *PC) Reset()      { *m = PC{} }
func (*PC) ProtoMessage() {}
func (*PC) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4e78b3c9dcf99c, []int{1}
}
func (m *PC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PC.Merge(m, src)
}
func (m *PC) XXX_Size() int {
	return m.Size()
}
func (m *PC) XXX_DiscardUnknown() {
	xxx_messageInfo_PC.DiscardUnknown(m)
}

var xxx_messageInfo_PC proto.InternalMessageInfo

func init() {
	proto.RegisterType((*F)(nil), "faction.F")
	golang_proto.RegisterType((*F)(nil), "faction.F")
	proto.RegisterType((*PC)(nil), "faction.PC")
	golang_proto.RegisterType((*PC)(nil), "faction.PC")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/faction/faction.proto", fileDescriptor_bd4e78b3c9dcf99c)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/faction/faction.proto", fileDescriptor_bd4e78b3c9dcf99c)
}

var fileDescriptor_bd4e78b3c9dcf99c = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcd, 0xc9, 0xcf, 0x4a, 0xcc, 0xd0, 0x4f, 0x4f,
	0xcc, 0x4d, 0x8d, 0x37, 0x30, 0xd6, 0x2f, 0xc8, 0x4e, 0xd7, 0x4f, 0x4b, 0x4c, 0x2e, 0xc9, 0xcc,
	0xcf, 0x83, 0xd1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae, 0x94, 0x11, 0x7e,
	0xdd, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x1d, 0x60, 0x16, 0x44, 0xb3, 0xd2, 0x51, 0x46, 0x2e, 0x46,
	0x37, 0x21, 0x5b, 0x2e, 0x26, 0x4f, 0x17, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x1e, 0x27, 0xdd, 0x13,
	0xf7, 0xe4, 0x19, 0x6e, 0xdd, 0x93, 0x57, 0xc5, 0x6f, 0x5a, 0x69, 0x4e, 0x66, 0x8a, 0x9e, 0xa7,
	0x4b, 0x10, 0x93, 0xa7, 0x8b, 0x90, 0x3b, 0x17, 0x7b, 0x78, 0x7e, 0x51, 0x4e, 0x8a, 0xa7, 0x8b,
	0x04, 0x13, 0x39, 0x66, 0xc0, 0x74, 0x0b, 0x09, 0x71, 0xb1, 0xf8, 0x25, 0xe6, 0xa6, 0x4a, 0x30,
	0x2b, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x20, 0x31, 0xcf, 0xe4, 0xfc, 0x3c, 0x09, 0x16, 0x88,
	0x18, 0x88, 0x2d, 0x24, 0xc0, 0xc5, 0xec, 0x9b, 0x58, 0x21, 0xc1, 0xaa, 0xc0, 0xa8, 0xc1, 0x1c,
	0x04, 0x62, 0x2a, 0x6d, 0x60, 0xe4, 0x62, 0x0a, 0x70, 0xa6, 0xd4, 0x23, 0xde, 0x5c, 0x9c, 0x6e,
	0x90, 0xc0, 0x24, 0xd7, 0x2b, 0x08, 0xfd, 0x42, 0x72, 0x5c, 0x5c, 0x01, 0xa9, 0x45, 0xb9, 0x99,
	0xc5, 0xc5, 0x99, 0xf9, 0x79, 0x60, 0x2f, 0x31, 0x07, 0x21, 0x89, 0x38, 0x79, 0x9c, 0x78, 0x28,
	0xc7, 0x70, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0xfe, 0x78,
	0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x1d, 0x8f, 0xe4, 0x18, 0x0f,
	0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x5f, 0x3c, 0x92, 0x63, 0xf8, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x03, 0x8f, 0xe5,
	0x18, 0x2f, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0x1c, 0x97, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x50, 0x90, 0x5f, 0x3f, 0x02, 0x00, 0x00,
}

func (this *F) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*F)
	if !ok {
		that2, ok := that.(F)
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
	if !this.WorldID.Equal(that1.WorldID) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Icon != that1.Icon {
		return false
	}
	if this.Max != that1.Max {
		return false
	}
	return true
}
func (this *PC) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PC)
	if !ok {
		that2, ok := that.(PC)
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
	if !this.FactionID.Equal(that1.FactionID) {
		return false
	}
	if this.Permission != that1.Permission {
		return false
	}
	return true
}
func (this *F) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&faction.F{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "WorldID: "+fmt.Sprintf("%#v", this.WorldID)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Icon: "+fmt.Sprintf("%#v", this.Icon)+",\n")
	s = append(s, "Max: "+fmt.Sprintf("%#v", this.Max)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PC) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&faction.PC{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "FactionID: "+fmt.Sprintf("%#v", this.FactionID)+",\n")
	s = append(s, "Permission: "+fmt.Sprintf("%#v", this.Permission)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFaction(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *F) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *F) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *F) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Max != 0 {
		i = encodeVarintFaction(dAtA, i, uint64(m.Max))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Icon) > 0 {
		i -= len(m.Icon)
		copy(dAtA[i:], m.Icon)
		i = encodeVarintFaction(dAtA, i, uint64(len(m.Icon)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintFaction(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.WorldID.Size()
		i -= size
		if _, err := m.WorldID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFaction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFaction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Permission != 0 {
		i = encodeVarintFaction(dAtA, i, uint64(m.Permission))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.FactionID.Size()
		i -= size
		if _, err := m.FactionID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFaction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFaction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintFaction(dAtA []byte, offset int, v uint64) int {
	offset -= sovFaction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedF(r randyFaction, easy bool) *F {
	this := &F{}
	v1 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.ID = *v1
	v2 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.WorldID = *v2
	this.Name = string(randStringFaction(r))
	this.Icon = string(randStringFaction(r))
	this.Max = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Max *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedPC(r randyFaction, easy bool) *PC {
	this := &PC{}
	v3 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.ID = *v3
	v4 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.FactionID = *v4
	this.Permission = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Permission *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyFaction interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFaction(r randyFaction) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFaction(r randyFaction) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneFaction(r)
	}
	return string(tmps)
}
func randUnrecognizedFaction(r randyFaction, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldFaction(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldFaction(dAtA []byte, r randyFaction, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateFaction(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateFaction(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateFaction(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateFaction(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateFaction(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateFaction(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateFaction(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *F) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovFaction(uint64(l))
	l = m.WorldID.Size()
	n += 1 + l + sovFaction(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFaction(uint64(l))
	}
	l = len(m.Icon)
	if l > 0 {
		n += 1 + l + sovFaction(uint64(l))
	}
	if m.Max != 0 {
		n += 1 + sovFaction(uint64(m.Max))
	}
	return n
}

func (m *PC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovFaction(uint64(l))
	l = m.FactionID.Size()
	n += 1 + l + sovFaction(uint64(l))
	if m.Permission != 0 {
		n += 1 + sovFaction(uint64(m.Permission))
	}
	return n
}

func sovFaction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFaction(x uint64) (n int) {
	return sovFaction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *F) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&F{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`WorldID:` + fmt.Sprintf("%v", this.WorldID) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Icon:` + fmt.Sprintf("%v", this.Icon) + `,`,
		`Max:` + fmt.Sprintf("%v", this.Max) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PC) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PC{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`FactionID:` + fmt.Sprintf("%v", this.FactionID) + `,`,
		`Permission:` + fmt.Sprintf("%v", this.Permission) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFaction(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *F) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFaction
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
			return fmt.Errorf("proto: F: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: F: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaction
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
				return ErrInvalidLengthFaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFaction
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
				return fmt.Errorf("proto: wrong wireType = %d for field WorldID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaction
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
				return ErrInvalidLengthFaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WorldID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaction
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
				return ErrInvalidLengthFaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Icon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaction
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
				return ErrInvalidLengthFaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Icon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Max", wireType)
			}
			m.Max = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Max |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFaction
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFaction
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
func (m *PC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFaction
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
			return fmt.Errorf("proto: PC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaction
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
				return ErrInvalidLengthFaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFaction
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
				return fmt.Errorf("proto: wrong wireType = %d for field FactionID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaction
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
				return ErrInvalidLengthFaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FactionID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			m.Permission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Permission |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFaction
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFaction
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
func skipFaction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFaction
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
					return 0, ErrIntOverflowFaction
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
					return 0, ErrIntOverflowFaction
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
				return 0, ErrInvalidLengthFaction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFaction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFaction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFaction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFaction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFaction = fmt.Errorf("proto: unexpected end of group")
)
