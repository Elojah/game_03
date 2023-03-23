// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/geometry/geometry.proto

package geometry

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

type Vec2 struct {
	X int64 `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y int64 `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *Vec2) Reset()      { *m = Vec2{} }
func (*Vec2) ProtoMessage() {}
func (*Vec2) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dc0e58235b9cb07, []int{0}
}
func (m *Vec2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vec2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vec2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vec2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vec2.Merge(m, src)
}
func (m *Vec2) XXX_Size() int {
	return m.Size()
}
func (m *Vec2) XXX_DiscardUnknown() {
	xxx_messageInfo_Vec2.DiscardUnknown(m)
}

var xxx_messageInfo_Vec2 proto.InternalMessageInfo

type Rect struct {
	X      int64  `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty" cql:"x"`
	Y      int64  `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty" cql:"y"`
	Height uint64 `protobuf:"varint,3,opt,name=Height,proto3" json:"Height,omitempty" cql:"height"`
	Width  uint64 `protobuf:"varint,4,opt,name=Width,proto3" json:"Width,omitempty" cql:"width"`
	// degree clockwise
	Rotation uint64 `protobuf:"varint,5,opt,name=Rotation,proto3" json:"Rotation,omitempty" cql:"rotation"`
}

func (m *Rect) Reset()      { *m = Rect{} }
func (*Rect) ProtoMessage() {}
func (*Rect) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dc0e58235b9cb07, []int{1}
}
func (m *Rect) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Rect.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Rect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rect.Merge(m, src)
}
func (m *Rect) XXX_Size() int {
	return m.Size()
}
func (m *Rect) XXX_DiscardUnknown() {
	xxx_messageInfo_Rect.DiscardUnknown(m)
}

var xxx_messageInfo_Rect proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Vec2)(nil), "geometry.Vec2")
	golang_proto.RegisterType((*Vec2)(nil), "geometry.Vec2")
	proto.RegisterType((*Rect)(nil), "geometry.Rect")
	golang_proto.RegisterType((*Rect)(nil), "geometry.Rect")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/geometry/geometry.proto", fileDescriptor_6dc0e58235b9cb07)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/geometry/geometry.proto", fileDescriptor_6dc0e58235b9cb07)
}

var fileDescriptor_6dc0e58235b9cb07 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xfa, 0x40,
	0x14, 0xc0, 0xfb, 0xfe, 0x14, 0xfe, 0xa4, 0x10, 0x35, 0x37, 0x15, 0x87, 0x57, 0x72, 0x89, 0x09,
	0x8b, 0xd4, 0x48, 0x5c, 0x1c, 0x99, 0x98, 0x6f, 0x50, 0x98, 0x0c, 0xd4, 0xf3, 0x5a, 0x05, 0x0f,
	0xc9, 0x11, 0x65, 0xf3, 0x23, 0xf8, 0x31, 0xfc, 0x08, 0x2e, 0x26, 0x8c, 0x8c, 0x8c, 0x4c, 0x8d,
	0xbd, 0x2e, 0x8e, 0xa4, 0x93, 0xa3, 0xe1, 0xc0, 0x86, 0xed, 0xfd, 0xde, 0xef, 0xf7, 0x6e, 0x38,
	0xe7, 0x42, 0x44, 0x2a, 0x9c, 0x0e, 0x9a, 0x81, 0x1c, 0xf9, 0x7c, 0x28, 0xef, 0xfb, 0xa1, 0x2f,
	0xfa, 0x23, 0x7e, 0x73, 0xd6, 0xf2, 0xc7, 0x0f, 0xc2, 0x17, 0x5c, 0x8e, 0xb8, 0x9a, 0xcc, 0xf2,
	0xa1, 0x39, 0x9e, 0x48, 0x25, 0x49, 0xf9, 0x8f, 0x8f, 0x4f, 0xf7, 0x1e, 0x10, 0x52, 0x48, 0xdf,
	0x04, 0x83, 0xe9, 0x9d, 0x21, 0x03, 0x66, 0xda, 0x1e, 0x52, 0xea, 0xd8, 0x57, 0x3c, 0x38, 0x27,
	0x55, 0x07, 0xba, 0x2e, 0xd4, 0xa1, 0x51, 0x60, 0xd0, 0xdd, 0x50, 0xcf, 0xfd, 0xb7, 0xa5, 0x1e,
	0xfd, 0x04, 0xc7, 0x66, 0x3c, 0x50, 0xa4, 0x96, 0x47, 0xed, 0x4a, 0x16, 0x7b, 0xff, 0x83, 0xa7,
	0xe1, 0x25, 0x7d, 0xa1, 0x9b, 0x8b, 0x5a, 0x7e, 0xb1, 0xa7, 0x66, 0x94, 0x41, 0x8f, 0x34, 0x9c,
	0x52, 0x87, 0x47, 0x22, 0x54, 0x6e, 0xa1, 0x0e, 0x0d, 0xbb, 0x7d, 0x94, 0xc5, 0x5e, 0xd5, 0xf8,
	0xd0, 0xac, 0x29, 0xdb, 0x79, 0x72, 0xe2, 0x14, 0xaf, 0xa3, 0x5b, 0x15, 0xba, 0xb6, 0x09, 0x0f,
	0xb3, 0xd8, 0xab, 0x98, 0xf0, 0x79, 0xb3, 0xa5, 0x6c, 0x6b, 0x49, 0xd3, 0x29, 0x33, 0xa9, 0xfa,
	0x2a, 0x92, 0x8f, 0x6e, 0xd1, 0x94, 0x24, 0x8b, 0xbd, 0x03, 0x53, 0x4e, 0x76, 0x82, 0xb2, 0xbc,
	0x69, 0x77, 0x16, 0x09, 0x5a, 0xcb, 0x04, 0xad, 0x55, 0x82, 0xd6, 0x3a, 0x41, 0xf8, 0x49, 0x10,
	0x5e, 0x35, 0xc2, 0xbb, 0x46, 0xf8, 0xd0, 0x08, 0x73, 0x8d, 0xb0, 0xd0, 0x08, 0x4b, 0x8d, 0xf0,
	0xa5, 0x11, 0xbe, 0x35, 0x5a, 0x6b, 0x8d, 0xf0, 0x96, 0xa2, 0x35, 0x4f, 0x11, 0x96, 0x29, 0x5a,
	0xab, 0x14, 0xad, 0x41, 0xc9, 0x7c, 0x5a, 0xeb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x73, 0x0d,
	0xce, 0xa6, 0x01, 0x00, 0x00,
}

func (this *Vec2) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Vec2)
	if !ok {
		that2, ok := that.(Vec2)
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
	if this.X != that1.X {
		return false
	}
	if this.Y != that1.Y {
		return false
	}
	return true
}
func (this *Rect) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Rect)
	if !ok {
		that2, ok := that.(Rect)
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
	if this.X != that1.X {
		return false
	}
	if this.Y != that1.Y {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if this.Width != that1.Width {
		return false
	}
	if this.Rotation != that1.Rotation {
		return false
	}
	return true
}
func (this *Vec2) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&geometry.Vec2{")
	s = append(s, "X: "+fmt.Sprintf("%#v", this.X)+",\n")
	s = append(s, "Y: "+fmt.Sprintf("%#v", this.Y)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Rect) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&geometry.Rect{")
	s = append(s, "X: "+fmt.Sprintf("%#v", this.X)+",\n")
	s = append(s, "Y: "+fmt.Sprintf("%#v", this.Y)+",\n")
	s = append(s, "Height: "+fmt.Sprintf("%#v", this.Height)+",\n")
	s = append(s, "Width: "+fmt.Sprintf("%#v", this.Width)+",\n")
	s = append(s, "Rotation: "+fmt.Sprintf("%#v", this.Rotation)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGeometry(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Vec2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vec2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vec2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Y != 0 {
		i = encodeVarintGeometry(dAtA, i, uint64(m.Y))
		i--
		dAtA[i] = 0x10
	}
	if m.X != 0 {
		i = encodeVarintGeometry(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Rect) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rect) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Rect) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Rotation != 0 {
		i = encodeVarintGeometry(dAtA, i, uint64(m.Rotation))
		i--
		dAtA[i] = 0x28
	}
	if m.Width != 0 {
		i = encodeVarintGeometry(dAtA, i, uint64(m.Width))
		i--
		dAtA[i] = 0x20
	}
	if m.Height != 0 {
		i = encodeVarintGeometry(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if m.Y != 0 {
		i = encodeVarintGeometry(dAtA, i, uint64(m.Y))
		i--
		dAtA[i] = 0x10
	}
	if m.X != 0 {
		i = encodeVarintGeometry(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGeometry(dAtA []byte, offset int, v uint64) int {
	offset -= sovGeometry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedVec2(r randyGeometry, easy bool) *Vec2 {
	this := &Vec2{}
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

func NewPopulatedRect(r randyGeometry, easy bool) *Rect {
	this := &Rect{}
	this.X = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.X *= -1
	}
	this.Y = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Y *= -1
	}
	this.Height = uint64(uint64(r.Uint32()))
	this.Width = uint64(uint64(r.Uint32()))
	this.Rotation = uint64(uint64(r.Uint32()))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyGeometry interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneGeometry(r randyGeometry) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringGeometry(r randyGeometry) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneGeometry(r)
	}
	return string(tmps)
}
func randUnrecognizedGeometry(r randyGeometry, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldGeometry(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldGeometry(dAtA []byte, r randyGeometry, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateGeometry(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateGeometry(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateGeometry(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateGeometry(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateGeometry(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateGeometry(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateGeometry(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Vec2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 1 + sovGeometry(uint64(m.X))
	}
	if m.Y != 0 {
		n += 1 + sovGeometry(uint64(m.Y))
	}
	return n
}

func (m *Rect) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 1 + sovGeometry(uint64(m.X))
	}
	if m.Y != 0 {
		n += 1 + sovGeometry(uint64(m.Y))
	}
	if m.Height != 0 {
		n += 1 + sovGeometry(uint64(m.Height))
	}
	if m.Width != 0 {
		n += 1 + sovGeometry(uint64(m.Width))
	}
	if m.Rotation != 0 {
		n += 1 + sovGeometry(uint64(m.Rotation))
	}
	return n
}

func sovGeometry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGeometry(x uint64) (n int) {
	return sovGeometry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Vec2) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Vec2{`,
		`X:` + fmt.Sprintf("%v", this.X) + `,`,
		`Y:` + fmt.Sprintf("%v", this.Y) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Rect) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Rect{`,
		`X:` + fmt.Sprintf("%v", this.X) + `,`,
		`Y:` + fmt.Sprintf("%v", this.Y) + `,`,
		`Height:` + fmt.Sprintf("%v", this.Height) + `,`,
		`Width:` + fmt.Sprintf("%v", this.Width) + `,`,
		`Rotation:` + fmt.Sprintf("%v", this.Rotation) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGeometry(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Vec2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGeometry
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
			return fmt.Errorf("proto: Vec2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vec2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeometry
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			m.Y = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeometry
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
			skippy, err := skipGeometry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGeometry
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
func (m *Rect) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGeometry
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
			return fmt.Errorf("proto: Rect: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rect: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeometry
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			m.Y = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeometry
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeometry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Width", wireType)
			}
			m.Width = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeometry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Width |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rotation", wireType)
			}
			m.Rotation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeometry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rotation |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGeometry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGeometry
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
func skipGeometry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGeometry
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
					return 0, ErrIntOverflowGeometry
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
					return 0, ErrIntOverflowGeometry
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
				return 0, ErrInvalidLengthGeometry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGeometry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGeometry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGeometry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGeometry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGeometry = fmt.Errorf("proto: unexpected end of group")
)
