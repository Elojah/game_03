// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/entity/pc.proto

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

type PC struct {
	ID       github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"ID"`
	UserID   github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,opt,name=UserID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"UserID"`
	WorldID  github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,3,opt,name=WorldID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"WorldID"`
	EntityID github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,4,opt,name=EntityID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"EntityID"`
	At       int64                                 `protobuf:"varint,5,opt,name=At,proto3" json:"At,omitempty"`
}

func (m *PC) Reset()      { *m = PC{} }
func (*PC) ProtoMessage() {}
func (*PC) Descriptor() ([]byte, []int) {
	return fileDescriptor_43374da76e05b622, []int{0}
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
	proto.RegisterType((*PC)(nil), "entity.PC")
	golang_proto.RegisterType((*PC)(nil), "entity.PC")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/entity/pc.proto", fileDescriptor_43374da76e05b622)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/entity/pc.proto", fileDescriptor_43374da76e05b622)
}

var fileDescriptor_43374da76e05b622 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcd, 0xc9, 0xcf, 0x4a, 0xcc, 0xd0, 0x4f, 0x4f,
	0xcc, 0x4d, 0x8d, 0x37, 0x30, 0xd6, 0x2f, 0xc8, 0x4e, 0xd7, 0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9,
	0xd4, 0x2f, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x08, 0x48, 0x21, 0x6b,
	0x4b, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x27, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05,
	0xd1, 0xa6, 0xb4, 0x8e, 0x89, 0x8b, 0x29, 0xc0, 0x59, 0xc8, 0x96, 0x8b, 0xc9, 0xd3, 0x45, 0x82,
	0x51, 0x81, 0x51, 0x83, 0xc7, 0x49, 0xf7, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0x55, 0xf1,
	0x3b, 0xa0, 0x34, 0x27, 0x33, 0x45, 0xcf, 0xd3, 0x25, 0x88, 0xc9, 0xd3, 0x45, 0xc8, 0x95, 0x8b,
	0x2d, 0xb4, 0x38, 0xb5, 0xc8, 0xd3, 0x45, 0x82, 0x89, 0x1c, 0x23, 0xa0, 0x9a, 0x85, 0xdc, 0xb9,
	0xd8, 0xc3, 0xf3, 0x8b, 0x72, 0x52, 0x3c, 0x5d, 0x24, 0x98, 0xc9, 0x31, 0x07, 0xa6, 0x5b, 0xc8,
	0x93, 0x8b, 0xc3, 0x15, 0x1c, 0x1c, 0x9e, 0x2e, 0x12, 0x2c, 0xe4, 0x98, 0x04, 0xd7, 0x2e, 0xc4,
	0xc7, 0xc5, 0xe4, 0x58, 0x22, 0xc1, 0xaa, 0xc0, 0xa8, 0xc1, 0x1c, 0xc4, 0xe4, 0x58, 0xe2, 0xe4,
	0x71, 0xe2, 0xa1, 0x1c, 0xc3, 0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x7c, 0x78, 0x28,
	0xc7, 0xf8, 0xe3, 0xa1, 0x1c, 0x63, 0xc3, 0x23, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18, 0x77, 0x3c,
	0x92, 0x63, 0x3c, 0xf0, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18,
	0x0e, 0x3c, 0x96, 0x63, 0xbc, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x24, 0x36, 0x70,
	0x0c, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x76, 0x02, 0xf8, 0x53, 0xe9, 0x01, 0x00, 0x00,
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
	if !this.UserID.Equal(that1.UserID) {
		return false
	}
	if !this.WorldID.Equal(that1.WorldID) {
		return false
	}
	if !this.EntityID.Equal(that1.EntityID) {
		return false
	}
	if this.At != that1.At {
		return false
	}
	return true
}
func (this *PC) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&entity.PC{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "UserID: "+fmt.Sprintf("%#v", this.UserID)+",\n")
	s = append(s, "WorldID: "+fmt.Sprintf("%#v", this.WorldID)+",\n")
	s = append(s, "EntityID: "+fmt.Sprintf("%#v", this.EntityID)+",\n")
	s = append(s, "At: "+fmt.Sprintf("%#v", this.At)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPc(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
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
	if m.At != 0 {
		i = encodeVarintPc(dAtA, i, uint64(m.At))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.EntityID.Size()
		i -= size
		if _, err := m.EntityID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.WorldID.Size()
		i -= size
		if _, err := m.WorldID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.UserID.Size()
		i -= size
		if _, err := m.UserID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintPc(dAtA []byte, offset int, v uint64) int {
	offset -= sovPc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedPC(r randyPc, easy bool) *PC {
	this := &PC{}
	v1 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.ID = *v1
	v2 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.UserID = *v2
	v3 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.WorldID = *v3
	v4 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.EntityID = *v4
	this.At = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.At *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyPc interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RunePc(r randyPc) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringPc(r randyPc) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RunePc(r)
	}
	return string(tmps)
}
func randUnrecognizedPc(r randyPc, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldPc(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldPc(dAtA []byte, r randyPc, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulatePc(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulatePc(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulatePc(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulatePc(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulatePc(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulatePc(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulatePc(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *PC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovPc(uint64(l))
	l = m.UserID.Size()
	n += 1 + l + sovPc(uint64(l))
	l = m.WorldID.Size()
	n += 1 + l + sovPc(uint64(l))
	l = m.EntityID.Size()
	n += 1 + l + sovPc(uint64(l))
	if m.At != 0 {
		n += 1 + sovPc(uint64(m.At))
	}
	return n
}

func sovPc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPc(x uint64) (n int) {
	return sovPc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PC) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PC{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`UserID:` + fmt.Sprintf("%v", this.UserID) + `,`,
		`WorldID:` + fmt.Sprintf("%v", this.WorldID) + `,`,
		`EntityID:` + fmt.Sprintf("%v", this.EntityID) + `,`,
		`At:` + fmt.Sprintf("%v", this.At) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPc(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPc
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
					return ErrIntOverflowPc
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
				return ErrInvalidLengthPc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPc
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
					return ErrIntOverflowPc
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
				return ErrInvalidLengthPc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPc
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
				return fmt.Errorf("proto: wrong wireType = %d for field WorldID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPc
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
				return ErrInvalidLengthPc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WorldID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPc
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
				return ErrInvalidLengthPc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EntityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field At", wireType)
			}
			m.At = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPc
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
		default:
			iNdEx = preIndex
			skippy, err := skipPc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPc
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
func skipPc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPc
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
					return 0, ErrIntOverflowPc
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
					return 0, ErrIntOverflowPc
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
				return 0, ErrInvalidLengthPc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPc = fmt.Errorf("proto: unexpected end of group")
)
