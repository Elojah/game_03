// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/twitch/token.proto

package twitch

import (
	fmt "fmt"
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

type Token struct {
	AccessToken  string   `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken string   `protobuf:"bytes,2,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	ExpiresIn    int64    `protobuf:"varint,3,opt,name=ExpiresIn,proto3" json:"ExpiresIn,omitempty"`
	Scope        []string `protobuf:"bytes,4,rep,name=Scope,proto3" json:"Scope,omitempty"`
	TokenType    string   `protobuf:"bytes,5,opt,name=TokenType,proto3" json:"TokenType,omitempty"`
}

func (m *Token) Reset()      { *m = Token{} }
func (*Token) ProtoMessage() {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_043f90278a46d79f, []int{0}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return m.Size()
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Token) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func (m *Token) GetScope() []string {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *Token) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func init() {
	proto.RegisterType((*Token)(nil), "twitch.Token")
	golang_proto.RegisterType((*Token)(nil), "twitch.Token")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/twitch/token.proto", fileDescriptor_043f90278a46d79f)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/twitch/token.proto", fileDescriptor_043f90278a46d79f)
}

var fileDescriptor_043f90278a46d79f = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x40, 0x7d, 0x84, 0x54, 0x8a, 0x61, 0x8a, 0x18, 0x22, 0x84, 0x4e, 0x51, 0xa7, 0x2c, 0xd4,
	0x95, 0xfa, 0x05, 0x20, 0x18, 0x58, 0x43, 0x77, 0xd4, 0x58, 0xae, 0x13, 0x4a, 0x6b, 0x2b, 0x71,
	0x05, 0x6c, 0x7c, 0x02, 0x7f, 0xc0, 0xca, 0x27, 0x30, 0x76, 0x64, 0xcc, 0xd8, 0x91, 0x38, 0x0b,
	0x63, 0x47, 0x46, 0x54, 0x67, 0x68, 0xbb, 0xf9, 0x3d, 0xdf, 0x3b, 0xe9, 0xe8, 0x50, 0x16, 0x26,
	0x5f, 0x66, 0x03, 0xae, 0xe6, 0x4c, 0x3c, 0xa9, 0xc7, 0x49, 0xce, 0xe4, 0x64, 0x2e, 0x1e, 0x86,
	0x23, 0xa6, 0x67, 0x92, 0x99, 0xe7, 0xc2, 0xf0, 0x9c, 0x19, 0x35, 0x13, 0x8b, 0x81, 0x2e, 0x95,
	0x51, 0x61, 0xaf, 0x73, 0xe7, 0x97, 0x7b, 0xa5, 0x54, 0x52, 0x31, 0xf7, 0x9d, 0x2d, 0xa7, 0x8e,
	0x1c, 0xb8, 0x57, 0x97, 0xf5, 0x3f, 0x80, 0xfa, 0xe3, 0xed, 0x9a, 0x30, 0xa6, 0x27, 0x57, 0x9c,
	0x8b, 0xaa, 0x72, 0x18, 0x41, 0x0c, 0x49, 0x90, 0xee, 0xab, 0xb0, 0x4f, 0x4f, 0x53, 0x31, 0x2d,
	0x45, 0x95, 0x77, 0x23, 0x47, 0x6e, 0xe4, 0xc0, 0x85, 0x17, 0x34, 0xb8, 0x7d, 0xd1, 0x45, 0x29,
	0xaa, 0xbb, 0x45, 0xe4, 0xc5, 0x90, 0x78, 0xe9, 0x4e, 0x84, 0x67, 0xd4, 0xbf, 0xe7, 0x4a, 0x8b,
	0xe8, 0x38, 0xf6, 0x92, 0x20, 0xed, 0x60, 0xdb, 0xb8, 0x78, 0xfc, 0xaa, 0x45, 0xe4, 0xbb, 0xa5,
	0x3b, 0x71, 0x7d, 0x53, 0x37, 0x48, 0xd6, 0x0d, 0x92, 0x4d, 0x83, 0xf0, 0xd7, 0x20, 0xbc, 0x59,
	0x84, 0x4f, 0x8b, 0xf0, 0x65, 0x11, 0x56, 0x16, 0xe1, 0xdb, 0x22, 0xd4, 0x16, 0xe1, 0xc7, 0x22,
	0xfc, 0x5a, 0x24, 0x1b, 0x8b, 0xf0, 0xde, 0x22, 0x59, 0xb5, 0x08, 0x75, 0x8b, 0x64, 0xdd, 0x22,
	0xc9, 0x7a, 0xee, 0xdc, 0xd1, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x34, 0x1f, 0x72, 0x59,
	0x01, 0x00, 0x00,
}

func (this *Token) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Token)
	if !ok {
		that2, ok := that.(Token)
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
	if this.AccessToken != that1.AccessToken {
		return false
	}
	if this.RefreshToken != that1.RefreshToken {
		return false
	}
	if this.ExpiresIn != that1.ExpiresIn {
		return false
	}
	if len(this.Scope) != len(that1.Scope) {
		return false
	}
	for i := range this.Scope {
		if this.Scope[i] != that1.Scope[i] {
			return false
		}
	}
	if this.TokenType != that1.TokenType {
		return false
	}
	return true
}
func (this *Token) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&twitch.Token{")
	s = append(s, "AccessToken: "+fmt.Sprintf("%#v", this.AccessToken)+",\n")
	s = append(s, "RefreshToken: "+fmt.Sprintf("%#v", this.RefreshToken)+",\n")
	s = append(s, "ExpiresIn: "+fmt.Sprintf("%#v", this.ExpiresIn)+",\n")
	s = append(s, "Scope: "+fmt.Sprintf("%#v", this.Scope)+",\n")
	s = append(s, "TokenType: "+fmt.Sprintf("%#v", this.TokenType)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringToken(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenType) > 0 {
		i -= len(m.TokenType)
		copy(dAtA[i:], m.TokenType)
		i = encodeVarintToken(dAtA, i, uint64(len(m.TokenType)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Scope) > 0 {
		for iNdEx := len(m.Scope) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Scope[iNdEx])
			copy(dAtA[i:], m.Scope[iNdEx])
			i = encodeVarintToken(dAtA, i, uint64(len(m.Scope[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.ExpiresIn != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.ExpiresIn))
		i--
		dAtA[i] = 0x18
	}
	if len(m.RefreshToken) > 0 {
		i -= len(m.RefreshToken)
		copy(dAtA[i:], m.RefreshToken)
		i = encodeVarintToken(dAtA, i, uint64(len(m.RefreshToken)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AccessToken) > 0 {
		i -= len(m.AccessToken)
		copy(dAtA[i:], m.AccessToken)
		i = encodeVarintToken(dAtA, i, uint64(len(m.AccessToken)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedToken(r randyToken, easy bool) *Token {
	this := &Token{}
	this.AccessToken = string(randStringToken(r))
	this.RefreshToken = string(randStringToken(r))
	this.ExpiresIn = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.ExpiresIn *= -1
	}
	v1 := r.Intn(10)
	this.Scope = make([]string, v1)
	for i := 0; i < v1; i++ {
		this.Scope[i] = string(randStringToken(r))
	}
	this.TokenType = string(randStringToken(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyToken interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneToken(r randyToken) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringToken(r randyToken) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneToken(r)
	}
	return string(tmps)
}
func randUnrecognizedToken(r randyToken, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldToken(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldToken(dAtA []byte, r randyToken, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateToken(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateToken(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateToken(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateToken(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateToken(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateToken(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateToken(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccessToken)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.RefreshToken)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.ExpiresIn != 0 {
		n += 1 + sovToken(uint64(m.ExpiresIn))
	}
	if len(m.Scope) > 0 {
		for _, s := range m.Scope {
			l = len(s)
			n += 1 + l + sovToken(uint64(l))
		}
	}
	l = len(m.TokenType)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	return n
}

func sovToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Token) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Token{`,
		`AccessToken:` + fmt.Sprintf("%v", this.AccessToken) + `,`,
		`RefreshToken:` + fmt.Sprintf("%v", this.RefreshToken) + `,`,
		`ExpiresIn:` + fmt.Sprintf("%v", this.ExpiresIn) + `,`,
		`Scope:` + fmt.Sprintf("%v", this.Scope) + `,`,
		`TokenType:` + fmt.Sprintf("%v", this.TokenType) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringToken(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefreshToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefreshToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresIn", wireType)
			}
			m.ExpiresIn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiresIn |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scope", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scope = append(m.Scope, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthToken
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthToken
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
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToken = fmt.Errorf("proto: unexpected end of group")
)
