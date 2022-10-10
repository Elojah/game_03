// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/entity/entity.proto

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

type E struct {
	ID          github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"ID"`
	UserID      github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,opt,name=UserID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"UserID"`
	CellID      github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,3,opt,name=CellID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"CellID"`
	Name        string                                `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	X           int64                                 `protobuf:"varint,5,opt,name=X,proto3" json:"X,omitempty"`
	Y           int64                                 `protobuf:"varint,6,opt,name=Y,proto3" json:"Y,omitempty"`
	Rot         int32                                 `protobuf:"varint,7,opt,name=Rot,proto3" json:"Rot,omitempty"`
	Radius      int32                                 `protobuf:"varint,8,opt,name=Radius,proto3" json:"Radius,omitempty"`
	At          int64                                 `protobuf:"varint,9,opt,name=At,proto3" json:"At,omitempty"`
	AnimationID github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,10,opt,name=AnimationID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"AnimationID"`
	AnimationAt int64                                 `protobuf:"varint,11,opt,name=AnimationAt,proto3" json:"AnimationAt,omitempty"`
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
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0xcf, 0xd2, 0x40,
	0x18, 0xc7, 0xef, 0x69, 0xa1, 0xca, 0x41, 0x8c, 0xb9, 0xc1, 0x5c, 0x1c, 0x1e, 0x1a, 0x13, 0x93,
	0x2e, 0x50, 0x0d, 0xb3, 0x43, 0xb1, 0x24, 0x76, 0xd1, 0xa4, 0x89, 0x09, 0x4c, 0xa6, 0x48, 0x2d,
	0xd5, 0x96, 0x23, 0x70, 0x1d, 0xdc, 0xfc, 0x08, 0x7e, 0x0c, 0x17, 0x77, 0x47, 0x46, 0x46, 0x46,
	0xe2, 0x40, 0xec, 0x75, 0x71, 0x64, 0x7c, 0xc7, 0x37, 0x5c, 0x19, 0x3a, 0xbd, 0x03, 0xd3, 0xfd,
	0x7f, 0xcf, 0xe5, 0xff, 0xbb, 0x67, 0x38, 0xfa, 0x3a, 0x49, 0xe5, 0xb2, 0x98, 0x0f, 0x3f, 0x8b,
	0xdc, 0x8d, 0x33, 0xf1, 0x35, 0x5a, 0xba, 0x49, 0x94, 0xc7, 0x9f, 0x5e, 0x8d, 0xdc, 0xf5, 0xb7,
	0xc4, 0x8d, 0x57, 0x32, 0x95, 0xdf, 0xaf, 0xc7, 0x70, 0xbd, 0x11, 0x52, 0x30, 0xab, 0xa6, 0xe7,
	0x83, 0x46, 0x35, 0x11, 0x89, 0x70, 0xf5, 0xf5, 0xbc, 0xf8, 0xa2, 0x49, 0x83, 0x4e, 0x75, 0xed,
	0xc5, 0x6f, 0x93, 0xc2, 0x84, 0xbd, 0xa1, 0x46, 0xe0, 0x73, 0xb0, 0xc1, 0xe9, 0x8d, 0x07, 0xfb,
	0x53, 0x9f, 0xfc, 0x3d, 0xf5, 0x5f, 0x3e, 0xbc, 0x43, 0x91, 0xa5, 0x8b, 0x61, 0xe0, 0x87, 0x46,
	0xe0, 0xb3, 0x09, 0xb5, 0x3e, 0x6e, 0xe3, 0x4d, 0xe0, 0x73, 0xe3, 0x16, 0xc5, 0xb5, 0x7c, 0xd1,
	0xbc, 0x8d, 0xb3, 0x2c, 0xf0, 0xb9, 0x79, 0x93, 0xa6, 0x2e, 0x33, 0x46, 0x5b, 0xef, 0xa3, 0x3c,
	0xe6, 0x2d, 0x1b, 0x9c, 0x4e, 0xa8, 0x33, 0xeb, 0x51, 0x98, 0xf2, 0xb6, 0x0d, 0x8e, 0x19, 0xc2,
	0xf4, 0x42, 0x33, 0x6e, 0xd5, 0x34, 0x63, 0x4f, 0xa9, 0x19, 0x0a, 0xc9, 0x1f, 0xd9, 0xe0, 0xb4,
	0xc3, 0x4b, 0x64, 0xcf, 0xa8, 0x15, 0x46, 0x8b, 0xb4, 0xd8, 0xf2, 0xc7, 0x7a, 0x78, 0x25, 0xf6,
	0x84, 0x1a, 0x9e, 0xe4, 0x1d, 0x5d, 0x34, 0x3c, 0xc9, 0x3e, 0xd0, 0xae, 0xb7, 0x4a, 0xf3, 0x48,
	0xa6, 0x62, 0x15, 0xf8, 0x9c, 0xde, 0xb2, 0x75, 0xd3, 0xc0, 0xec, 0x86, 0xd0, 0x93, 0xbc, 0xab,
	0x5f, 0x6a, 0x8e, 0xc6, 0xef, 0xf6, 0x25, 0x92, 0x43, 0x89, 0xe4, 0x58, 0x22, 0x39, 0x97, 0x08,
	0x77, 0x25, 0xc2, 0x0f, 0x85, 0xf0, 0x4b, 0x21, 0xfc, 0x51, 0x08, 0x3b, 0x85, 0xb0, 0x57, 0x08,
	0x07, 0x85, 0xf0, 0x4f, 0x21, 0xfc, 0x57, 0x48, 0xce, 0x0a, 0xe1, 0x67, 0x85, 0x64, 0x57, 0x21,
	0x1c, 0x2a, 0x24, 0xc7, 0x0a, 0xc9, 0xdc, 0xd2, 0x1f, 0x60, 0x74, 0x1f, 0x00, 0x00, 0xff, 0xff,
	0x41, 0x08, 0xa7, 0x24, 0x6c, 0x02, 0x00, 0x00,
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
	return true
}
func (this *E) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 15)
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
