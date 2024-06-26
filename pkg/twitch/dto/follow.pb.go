// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/twitch/dto/follow.proto

package dto

import (
	fmt "fmt"
	_ "github.com/elojah/game_03/pkg/gogoproto"
	twitch "github.com/elojah/game_03/pkg/twitch"
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

type ListFollowReq struct {
	After string `protobuf:"bytes,1,opt,name=After,proto3" json:"After,omitempty"`
	First string `protobuf:"bytes,2,opt,name=First,proto3" json:"First,omitempty"`
}

func (m *ListFollowReq) Reset()      { *m = ListFollowReq{} }
func (*ListFollowReq) ProtoMessage() {}
func (*ListFollowReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_77bd2470ae47b642, []int{0}
}
func (m *ListFollowReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListFollowReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListFollowReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListFollowReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFollowReq.Merge(m, src)
}
func (m *ListFollowReq) XXX_Size() int {
	return m.Size()
}
func (m *ListFollowReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFollowReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListFollowReq proto.InternalMessageInfo

type ListFollowResp struct {
	Follows []twitch.Follow `protobuf:"bytes,1,rep,name=Follows,proto3" json:"Follows"`
	Total   uint64          `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	Cursor  string          `protobuf:"bytes,3,opt,name=Cursor,proto3" json:"Cursor,omitempty"`
}

func (m *ListFollowResp) Reset()      { *m = ListFollowResp{} }
func (*ListFollowResp) ProtoMessage() {}
func (*ListFollowResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_77bd2470ae47b642, []int{1}
}
func (m *ListFollowResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListFollowResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListFollowResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListFollowResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFollowResp.Merge(m, src)
}
func (m *ListFollowResp) XXX_Size() int {
	return m.Size()
}
func (m *ListFollowResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFollowResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListFollowResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListFollowReq)(nil), "dto.ListFollowReq")
	golang_proto.RegisterType((*ListFollowReq)(nil), "dto.ListFollowReq")
	proto.RegisterType((*ListFollowResp)(nil), "dto.ListFollowResp")
	golang_proto.RegisterType((*ListFollowResp)(nil), "dto.ListFollowResp")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/twitch/dto/follow.proto", fileDescriptor_77bd2470ae47b642)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/twitch/dto/follow.proto", fileDescriptor_77bd2470ae47b642)
}

var fileDescriptor_77bd2470ae47b642 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0x31, 0x4f, 0xc2, 0x40,
	0x18, 0x86, 0xef, 0x13, 0xc4, 0x78, 0x46, 0x86, 0xc6, 0x98, 0x86, 0xe1, 0x93, 0x30, 0x31, 0xf5,
	0x14, 0xe2, 0xe4, 0x24, 0x26, 0xc4, 0xc1, 0xa9, 0x71, 0x37, 0x05, 0x4a, 0x5b, 0x2d, 0x7e, 0xb5,
	0x3d, 0xc2, 0xea, 0x4f, 0xf0, 0x67, 0xf8, 0x13, 0x1c, 0x19, 0x3b, 0x76, 0x64, 0x32, 0xf6, 0xba,
	0x38, 0x32, 0x3a, 0x1a, 0xee, 0x34, 0x61, 0x32, 0x6e, 0xf7, 0xbc, 0xf9, 0xde, 0xf7, 0xc9, 0xf1,
	0xf3, 0x20, 0x92, 0xe1, 0x7c, 0xe4, 0x8c, 0x69, 0x26, 0xfc, 0x98, 0xee, 0xbd, 0x50, 0x04, 0xde,
	0xcc, 0xbf, 0x3b, 0xed, 0x8b, 0xe4, 0x21, 0x10, 0x72, 0x11, 0xc9, 0x71, 0x28, 0x26, 0x92, 0xc4,
	0x94, 0xe2, 0x98, 0x16, 0x4e, 0x92, 0x92, 0x24, 0xab, 0x36, 0x91, 0xd4, 0xea, 0xfd, 0xdd, 0x0d,
	0x28, 0x20, 0x7d, 0xad, 0x5f, 0xa6, 0xd8, 0x3a, 0xfb, 0x97, 0x6f, 0xdb, 0xd5, 0xb9, 0xe0, 0x87,
	0x37, 0x51, 0x26, 0x87, 0x3a, 0x73, 0xfd, 0x27, 0xeb, 0x88, 0xef, 0x5e, 0x4e, 0xa5, 0x9f, 0xda,
	0xd0, 0x86, 0xee, 0xbe, 0x6b, 0x60, 0x93, 0x0e, 0xa3, 0x34, 0x93, 0xf6, 0x8e, 0x49, 0x35, 0x74,
	0x1e, 0x79, 0x73, 0xbb, 0x9c, 0x25, 0x96, 0xc3, 0xf7, 0x0c, 0x65, 0x36, 0xb4, 0x6b, 0xdd, 0x83,
	0x5e, 0xd3, 0x31, 0x56, 0xc7, 0xc4, 0x83, 0x7a, 0xfe, 0x7e, 0xc2, 0xdc, 0xdf, 0xa3, 0xcd, 0xee,
	0x2d, 0x49, 0x2f, 0xd6, 0xbb, 0x75, 0xd7, 0x80, 0x75, 0xcc, 0x1b, 0x57, 0xf3, 0x34, 0xa3, 0xd4,
	0xae, 0x69, 0xdd, 0x0f, 0x0d, 0xae, 0xf3, 0x12, 0x59, 0x51, 0x22, 0x5b, 0x95, 0xc8, 0xd6, 0x25,
	0xc2, 0x57, 0x89, 0xf0, 0xac, 0x10, 0x5e, 0x15, 0xc2, 0x9b, 0x42, 0x58, 0x2a, 0x84, 0x5c, 0x21,
	0x14, 0x0a, 0xe1, 0x43, 0x21, 0x7c, 0x2a, 0x64, 0x6b, 0x85, 0xf0, 0x52, 0x21, 0x5b, 0x56, 0x08,
	0x45, 0x85, 0x6c, 0x55, 0x21, 0x1b, 0x35, 0xf4, 0xef, 0xfb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf6, 0xf4, 0x19, 0x09, 0xa2, 0x01, 0x00, 0x00,
}

func (this *ListFollowReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListFollowReq)
	if !ok {
		that2, ok := that.(ListFollowReq)
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
	if this.After != that1.After {
		return false
	}
	if this.First != that1.First {
		return false
	}
	return true
}
func (this *ListFollowResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListFollowResp)
	if !ok {
		that2, ok := that.(ListFollowResp)
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
	if len(this.Follows) != len(that1.Follows) {
		return false
	}
	for i := range this.Follows {
		if !this.Follows[i].Equal(&that1.Follows[i]) {
			return false
		}
	}
	if this.Total != that1.Total {
		return false
	}
	if this.Cursor != that1.Cursor {
		return false
	}
	return true
}
func (this *ListFollowReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dto.ListFollowReq{")
	s = append(s, "After: "+fmt.Sprintf("%#v", this.After)+",\n")
	s = append(s, "First: "+fmt.Sprintf("%#v", this.First)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ListFollowResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&dto.ListFollowResp{")
	if this.Follows != nil {
		vs := make([]twitch.Follow, len(this.Follows))
		for i := range vs {
			vs[i] = this.Follows[i]
		}
		s = append(s, "Follows: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "Total: "+fmt.Sprintf("%#v", this.Total)+",\n")
	s = append(s, "Cursor: "+fmt.Sprintf("%#v", this.Cursor)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFollow(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ListFollowReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListFollowReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListFollowReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.First) > 0 {
		i -= len(m.First)
		copy(dAtA[i:], m.First)
		i = encodeVarintFollow(dAtA, i, uint64(len(m.First)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.After) > 0 {
		i -= len(m.After)
		copy(dAtA[i:], m.After)
		i = encodeVarintFollow(dAtA, i, uint64(len(m.After)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListFollowResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListFollowResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListFollowResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cursor) > 0 {
		i -= len(m.Cursor)
		copy(dAtA[i:], m.Cursor)
		i = encodeVarintFollow(dAtA, i, uint64(len(m.Cursor)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Total != 0 {
		i = encodeVarintFollow(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Follows) > 0 {
		for iNdEx := len(m.Follows) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Follows[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFollow(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintFollow(dAtA []byte, offset int, v uint64) int {
	offset -= sovFollow(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedListFollowReq(r randyFollow, easy bool) *ListFollowReq {
	this := &ListFollowReq{}
	this.After = string(randStringFollow(r))
	this.First = string(randStringFollow(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedListFollowResp(r randyFollow, easy bool) *ListFollowResp {
	this := &ListFollowResp{}
	if r.Intn(5) != 0 {
		v1 := r.Intn(5)
		this.Follows = make([]twitch.Follow, v1)
		for i := 0; i < v1; i++ {
			v2 := twitch.NewPopulatedFollow(r, easy)
			this.Follows[i] = *v2
		}
	}
	this.Total = uint64(uint64(r.Uint32()))
	this.Cursor = string(randStringFollow(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyFollow interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFollow(r randyFollow) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFollow(r randyFollow) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneFollow(r)
	}
	return string(tmps)
}
func randUnrecognizedFollow(r randyFollow, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldFollow(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldFollow(dAtA []byte, r randyFollow, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateFollow(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateFollow(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateFollow(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateFollow(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateFollow(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateFollow(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateFollow(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ListFollowReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.After)
	if l > 0 {
		n += 1 + l + sovFollow(uint64(l))
	}
	l = len(m.First)
	if l > 0 {
		n += 1 + l + sovFollow(uint64(l))
	}
	return n
}

func (m *ListFollowResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Follows) > 0 {
		for _, e := range m.Follows {
			l = e.Size()
			n += 1 + l + sovFollow(uint64(l))
		}
	}
	if m.Total != 0 {
		n += 1 + sovFollow(uint64(m.Total))
	}
	l = len(m.Cursor)
	if l > 0 {
		n += 1 + l + sovFollow(uint64(l))
	}
	return n
}

func sovFollow(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFollow(x uint64) (n int) {
	return sovFollow(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ListFollowReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListFollowReq{`,
		`After:` + fmt.Sprintf("%v", this.After) + `,`,
		`First:` + fmt.Sprintf("%v", this.First) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ListFollowResp) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForFollows := "[]Follow{"
	for _, f := range this.Follows {
		repeatedStringForFollows += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForFollows += "}"
	s := strings.Join([]string{`&ListFollowResp{`,
		`Follows:` + repeatedStringForFollows + `,`,
		`Total:` + fmt.Sprintf("%v", this.Total) + `,`,
		`Cursor:` + fmt.Sprintf("%v", this.Cursor) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFollow(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ListFollowReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFollow
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
			return fmt.Errorf("proto: ListFollowReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListFollowReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field After", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFollow
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
				return ErrInvalidLengthFollow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFollow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.After = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field First", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFollow
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
				return ErrInvalidLengthFollow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFollow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.First = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFollow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFollow
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFollow
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
func (m *ListFollowResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFollow
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
			return fmt.Errorf("proto: ListFollowResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListFollowResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Follows", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFollow
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
				return ErrInvalidLengthFollow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFollow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Follows = append(m.Follows, twitch.Follow{})
			if err := m.Follows[len(m.Follows)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFollow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cursor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFollow
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
				return ErrInvalidLengthFollow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFollow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cursor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFollow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFollow
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFollow
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
func skipFollow(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFollow
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
					return 0, ErrIntOverflowFollow
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
					return 0, ErrIntOverflowFollow
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
				return 0, ErrInvalidLengthFollow
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFollow
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFollow
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFollow        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFollow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFollow = fmt.Errorf("proto: unexpected end of group")
)
