// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/entity/spawn.proto

package entity

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

type Spawn struct {
	EntityID github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=EntityID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"EntityID"`
	SpawnID  github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,opt,name=SpawnID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"SpawnID"`
}

func (m *Spawn) Reset()      { *m = Spawn{} }
func (*Spawn) ProtoMessage() {}
func (*Spawn) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a23a91fb132b29b, []int{0}
}
func (m *Spawn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Spawn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Spawn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Spawn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spawn.Merge(m, src)
}
func (m *Spawn) XXX_Size() int {
	return m.Size()
}
func (m *Spawn) XXX_DiscardUnknown() {
	xxx_messageInfo_Spawn.DiscardUnknown(m)
}

var xxx_messageInfo_Spawn proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Spawn)(nil), "entity.Spawn")
	golang_proto.RegisterType((*Spawn)(nil), "entity.Spawn")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/entity/spawn.proto", fileDescriptor_6a23a91fb132b29b)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/entity/spawn.proto", fileDescriptor_6a23a91fb132b29b)
}

var fileDescriptor_6a23a91fb132b29b = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcd, 0xc9, 0xcf, 0x4a, 0xcc, 0xd0, 0x4f, 0x4f,
	0xcc, 0x4d, 0x8d, 0x37, 0x30, 0xd6, 0x2f, 0xc8, 0x4e, 0xd7, 0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9,
	0xd4, 0x2f, 0x2e, 0x48, 0x2c, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x88,
	0x49, 0x19, 0xe1, 0xd7, 0x99, 0x9e, 0x9f, 0x9e, 0x0f, 0xd6, 0x00, 0x66, 0x41, 0xf4, 0x2a, 0xcd,
	0x66, 0xe4, 0x62, 0x0d, 0x06, 0x99, 0x25, 0xe4, 0xc9, 0xc5, 0xe1, 0x0a, 0x36, 0xc7, 0xd3, 0x45,
	0x82, 0x51, 0x81, 0x51, 0x83, 0xc7, 0x49, 0xf7, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0x55,
	0xf1, 0x9b, 0x5b, 0x9a, 0x93, 0x99, 0xa2, 0xe7, 0xe9, 0x12, 0x04, 0xd7, 0x2e, 0xe4, 0xce, 0xc5,
	0x0e, 0x36, 0xd3, 0xd3, 0x45, 0x82, 0x89, 0x1c, 0x93, 0x60, 0xba, 0x9d, 0x3c, 0x4e, 0x3c, 0x94,
	0x63, 0xb8, 0xf0, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x0f, 0x0f, 0xe5, 0x18, 0x7f, 0x3c,
	0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3, 0x8e, 0x47, 0x72, 0x8c, 0x07,
	0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x2f, 0x1e, 0xc9, 0x31, 0x7c, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x81, 0xc7, 0x72,
	0x8c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0xf6, 0xae, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x8c, 0x30, 0xe3, 0xdd, 0x5e, 0x01, 0x00, 0x00,
}

func (this *Spawn) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Spawn)
	if !ok {
		that2, ok := that.(Spawn)
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
	if !this.SpawnID.Equal(that1.SpawnID) {
		return false
	}
	return true
}
func (this *Spawn) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&entity.Spawn{")
	s = append(s, "EntityID: "+fmt.Sprintf("%#v", this.EntityID)+",\n")
	s = append(s, "SpawnID: "+fmt.Sprintf("%#v", this.SpawnID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSpawn(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Spawn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Spawn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Spawn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SpawnID.Size()
		i -= size
		if _, err := m.SpawnID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSpawn(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.EntityID.Size()
		i -= size
		if _, err := m.EntityID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSpawn(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintSpawn(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpawn(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedSpawn(r randySpawn, easy bool) *Spawn {
	this := &Spawn{}
	v1 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.EntityID = *v1
	v2 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.SpawnID = *v2
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randySpawn interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSpawn(r randySpawn) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringSpawn(r randySpawn) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneSpawn(r)
	}
	return string(tmps)
}
func randUnrecognizedSpawn(r randySpawn, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldSpawn(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldSpawn(dAtA []byte, r randySpawn, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateSpawn(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateSpawn(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateSpawn(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateSpawn(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateSpawn(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateSpawn(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateSpawn(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Spawn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EntityID.Size()
	n += 1 + l + sovSpawn(uint64(l))
	l = m.SpawnID.Size()
	n += 1 + l + sovSpawn(uint64(l))
	return n
}

func sovSpawn(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpawn(x uint64) (n int) {
	return sovSpawn(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Spawn) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Spawn{`,
		`EntityID:` + fmt.Sprintf("%v", this.EntityID) + `,`,
		`SpawnID:` + fmt.Sprintf("%v", this.SpawnID) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSpawn(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Spawn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpawn
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
			return fmt.Errorf("proto: Spawn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Spawn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpawn
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
				return ErrInvalidLengthSpawn
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpawn
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
				return fmt.Errorf("proto: wrong wireType = %d for field SpawnID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpawn
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
				return ErrInvalidLengthSpawn
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpawn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SpawnID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpawn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpawn
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSpawn
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
func skipSpawn(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpawn
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
					return 0, ErrIntOverflowSpawn
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
					return 0, ErrIntOverflowSpawn
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
				return 0, ErrInvalidLengthSpawn
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpawn
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpawn
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpawn        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpawn          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpawn = fmt.Errorf("proto: unexpected end of group")
)
