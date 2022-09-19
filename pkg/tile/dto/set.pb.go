// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/tile/dto/set.proto

package dto

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

type CreateTilesetReq struct {
	JSON string `protobuf:"bytes,1,opt,name=JSON,proto3" json:"JSON,omitempty"`
}

func (m *CreateTilesetReq) Reset()      { *m = CreateTilesetReq{} }
func (*CreateTilesetReq) ProtoMessage() {}
func (*CreateTilesetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3240545ed3d4d328, []int{0}
}
func (m *CreateTilesetReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateTilesetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateTilesetReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateTilesetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTilesetReq.Merge(m, src)
}
func (m *CreateTilesetReq) XXX_Size() int {
	return m.Size()
}
func (m *CreateTilesetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTilesetReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTilesetReq proto.InternalMessageInfo

func (m *CreateTilesetReq) GetJSON() string {
	if m != nil {
		return m.JSON
	}
	return ""
}

type CreateTilesetResp struct {
	ID github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"ID"`
}

func (m *CreateTilesetResp) Reset()      { *m = CreateTilesetResp{} }
func (*CreateTilesetResp) ProtoMessage() {}
func (*CreateTilesetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3240545ed3d4d328, []int{1}
}
func (m *CreateTilesetResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateTilesetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateTilesetResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateTilesetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTilesetResp.Merge(m, src)
}
func (m *CreateTilesetResp) XXX_Size() int {
	return m.Size()
}
func (m *CreateTilesetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTilesetResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTilesetResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateTilesetReq)(nil), "dto.CreateTilesetReq")
	golang_proto.RegisterType((*CreateTilesetReq)(nil), "dto.CreateTilesetReq")
	proto.RegisterType((*CreateTilesetResp)(nil), "dto.CreateTilesetResp")
	golang_proto.RegisterType((*CreateTilesetResp)(nil), "dto.CreateTilesetResp")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/tile/dto/set.proto", fileDescriptor_3240545ed3d4d328)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/tile/dto/set.proto", fileDescriptor_3240545ed3d4d328)
}

var fileDescriptor_3240545ed3d4d328 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcd, 0xc9, 0xcf, 0x4a, 0xcc, 0xd0, 0x4f, 0x4f,
	0xcc, 0x4d, 0x8d, 0x37, 0x30, 0xd6, 0x2f, 0xc8, 0x4e, 0xd7, 0x2f, 0xc9, 0xcc, 0x49, 0xd5, 0x4f,
	0x29, 0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x29,
	0xc9, 0x97, 0xd2, 0x45, 0xd2, 0x96, 0x9e, 0x9f, 0x9e, 0xaf, 0x0f, 0x96, 0x4b, 0x2a, 0x4d, 0x03,
	0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x47, 0x49, 0x8d, 0x4b, 0xc0, 0xb9, 0x28, 0x35, 0xb1, 0x24,
	0x35, 0x24, 0x33, 0x27, 0xb5, 0x38, 0xb5, 0x24, 0x28, 0xb5, 0x50, 0x48, 0x88, 0x8b, 0xc5, 0x2b,
	0xd8, 0xdf, 0x4f, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x56, 0x0a, 0xe2, 0x12, 0x44,
	0x53, 0x57, 0x5c, 0x20, 0x64, 0xcb, 0xc5, 0xe4, 0xe9, 0x02, 0x56, 0xc6, 0xe3, 0xa4, 0x7b, 0xe2,
	0x9e, 0x3c, 0xc3, 0xad, 0x7b, 0xf2, 0xaa, 0xf8, 0x9d, 0x5d, 0x9a, 0x93, 0x99, 0xa2, 0xe7, 0xe9,
	0x12, 0xc4, 0xe4, 0xe9, 0xe2, 0xe4, 0x72, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f,
	0x1e, 0xca, 0x31, 0xfe, 0x78, 0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6,
	0x1d, 0x8f, 0xe4, 0x18, 0x0f, 0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x5f, 0x3c, 0x92, 0x63, 0xf8, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x03, 0x8f, 0xe5, 0x18, 0x2f, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x89,
	0x0d, 0xec, 0x11, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x08, 0x1d, 0xe7, 0x30, 0x01,
	0x00, 0x00,
}

func (this *CreateTilesetReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateTilesetReq)
	if !ok {
		that2, ok := that.(CreateTilesetReq)
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
	if this.JSON != that1.JSON {
		return false
	}
	return true
}
func (this *CreateTilesetResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateTilesetResp)
	if !ok {
		that2, ok := that.(CreateTilesetResp)
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
	return true
}
func (this *CreateTilesetReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dto.CreateTilesetReq{")
	s = append(s, "JSON: "+fmt.Sprintf("%#v", this.JSON)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateTilesetResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dto.CreateTilesetResp{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSet(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *CreateTilesetReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateTilesetReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateTilesetReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.JSON) > 0 {
		i -= len(m.JSON)
		copy(dAtA[i:], m.JSON)
		i = encodeVarintSet(dAtA, i, uint64(len(m.JSON)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreateTilesetResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateTilesetResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateTilesetResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSet(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintSet(dAtA []byte, offset int, v uint64) int {
	offset -= sovSet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedCreateTilesetReq(r randySet, easy bool) *CreateTilesetReq {
	this := &CreateTilesetReq{}
	this.JSON = string(randStringSet(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedCreateTilesetResp(r randySet, easy bool) *CreateTilesetResp {
	this := &CreateTilesetResp{}
	v1 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.ID = *v1
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randySet interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSet(r randySet) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringSet(r randySet) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneSet(r)
	}
	return string(tmps)
}
func randUnrecognizedSet(r randySet, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldSet(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldSet(dAtA []byte, r randySet, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateSet(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateSet(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateSet(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateSet(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateSet(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateSet(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateSet(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *CreateTilesetReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JSON)
	if l > 0 {
		n += 1 + l + sovSet(uint64(l))
	}
	return n
}

func (m *CreateTilesetResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovSet(uint64(l))
	return n
}

func sovSet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSet(x uint64) (n int) {
	return sovSet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CreateTilesetReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateTilesetReq{`,
		`JSON:` + fmt.Sprintf("%v", this.JSON) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CreateTilesetResp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateTilesetResp{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSet(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CreateTilesetReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSet
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
			return fmt.Errorf("proto: CreateTilesetReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateTilesetReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JSON", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
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
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JSON = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSet
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
func (m *CreateTilesetResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSet
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
			return fmt.Errorf("proto: CreateTilesetResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateTilesetResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
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
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSet
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
func skipSet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSet
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
					return 0, ErrIntOverflowSet
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
					return 0, ErrIntOverflowSet
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
				return 0, ErrInvalidLengthSet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSet = fmt.Errorf("proto: unexpected end of group")
)