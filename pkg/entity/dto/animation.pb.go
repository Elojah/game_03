// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/entity/dto/animation.proto

package dto

import (
	bytes "bytes"
	fmt "fmt"
	entity "github.com/elojah/game_03/pkg/entity"
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

type ListAnimationReq struct {
	IDs       []github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,rep,name=IDs,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"IDs"`
	EntityIDs []github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,2,rep,name=EntityIDs,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"EntityIDs"`
	Size_     int64                                   `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	State     []byte                                  `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *ListAnimationReq) Reset()      { *m = ListAnimationReq{} }
func (*ListAnimationReq) ProtoMessage() {}
func (*ListAnimationReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_90611f0100162c90, []int{0}
}
func (m *ListAnimationReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAnimationReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAnimationReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAnimationReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAnimationReq.Merge(m, src)
}
func (m *ListAnimationReq) XXX_Size() int {
	return m.Size()
}
func (m *ListAnimationReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAnimationReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListAnimationReq proto.InternalMessageInfo

func (m *ListAnimationReq) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *ListAnimationReq) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

type ListAnimationResp struct {
	Animations []entity.Animation `protobuf:"bytes,1,rep,name=Animations,proto3" json:"Animations"`
	State      []byte             `protobuf:"bytes,2,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *ListAnimationResp) Reset()      { *m = ListAnimationResp{} }
func (*ListAnimationResp) ProtoMessage() {}
func (*ListAnimationResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_90611f0100162c90, []int{1}
}
func (m *ListAnimationResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAnimationResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAnimationResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAnimationResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAnimationResp.Merge(m, src)
}
func (m *ListAnimationResp) XXX_Size() int {
	return m.Size()
}
func (m *ListAnimationResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAnimationResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListAnimationResp proto.InternalMessageInfo

func (m *ListAnimationResp) GetAnimations() []entity.Animation {
	if m != nil {
		return m.Animations
	}
	return nil
}

func (m *ListAnimationResp) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func init() {
	proto.RegisterType((*ListAnimationReq)(nil), "dto.ListAnimationReq")
	golang_proto.RegisterType((*ListAnimationReq)(nil), "dto.ListAnimationReq")
	proto.RegisterType((*ListAnimationResp)(nil), "dto.ListAnimationResp")
	golang_proto.RegisterType((*ListAnimationResp)(nil), "dto.ListAnimationResp")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/entity/dto/animation.proto", fileDescriptor_90611f0100162c90)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/entity/dto/animation.proto", fileDescriptor_90611f0100162c90)
}

var fileDescriptor_90611f0100162c90 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcd, 0xc9, 0xcf, 0x4a, 0xcc, 0xd0, 0x4f, 0x4f,
	0xcc, 0x4d, 0x8d, 0x37, 0x30, 0xd6, 0x2f, 0xc8, 0x4e, 0xd7, 0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9,
	0xd4, 0x4f, 0x29, 0xc9, 0xd7, 0x4f, 0xcc, 0xcb, 0xcc, 0x4d, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x29, 0xc9, 0x97, 0xd2, 0x45, 0xd2, 0x9e, 0x9e, 0x9f,
	0x9e, 0xaf, 0x0f, 0x96, 0x4b, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x47, 0xca,
	0x84, 0x28, 0xdb, 0xd0, 0x6c, 0x52, 0x3a, 0xc5, 0xc8, 0x25, 0xe0, 0x93, 0x59, 0x5c, 0xe2, 0x08,
	0x13, 0x0f, 0x4a, 0x2d, 0x14, 0xb2, 0xe7, 0x62, 0xf6, 0x74, 0x29, 0x96, 0x60, 0x54, 0x60, 0xd6,
	0xe0, 0x71, 0xd2, 0x3d, 0x71, 0x4f, 0x9e, 0xe1, 0xd6, 0x3d, 0x79, 0x55, 0xfc, 0xe6, 0x97, 0xe6,
	0x64, 0xa6, 0xe8, 0x79, 0xba, 0x04, 0x81, 0x74, 0x0a, 0x79, 0x73, 0x71, 0xba, 0x82, 0xed, 0x03,
	0x19, 0xc3, 0x44, 0x8e, 0x31, 0x08, 0xfd, 0x42, 0x42, 0x5c, 0x2c, 0xc1, 0x99, 0x55, 0xa9, 0x12,
	0xcc, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x60, 0xb6, 0x90, 0x08, 0x17, 0x6b, 0x70, 0x49, 0x62, 0x49,
	0xaa, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x84, 0xa3, 0x94, 0xc4, 0x25, 0x88, 0xe6, 0x97,
	0xe2, 0x02, 0x21, 0x73, 0x2e, 0x2e, 0xb8, 0x00, 0xc4, 0x4f, 0xdc, 0x46, 0x82, 0x7a, 0x90, 0xe0,
	0xd0, 0x83, 0xcb, 0x38, 0xb1, 0x80, 0xdc, 0x17, 0x84, 0xa4, 0x14, 0x61, 0x07, 0x13, 0x92, 0x1d,
	0x4e, 0x2e, 0x17, 0x1e, 0xca, 0x31, 0xdc, 0x78, 0x28, 0xc7, 0xf0, 0xe1, 0xa1, 0x1c, 0xe3, 0x8f,
	0x87, 0x72, 0x8c, 0x0d, 0x8f, 0xe4, 0x18, 0x57, 0x3c, 0x92, 0x63, 0xdc, 0xf1, 0x48, 0x8e, 0xf1,
	0xc0, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0x38, 0xf0, 0x58,
	0x8e, 0xf1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0xa1, 0x6f, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xa8, 0x2f, 0xda, 0x23, 0x02, 0x00, 0x00,
}

func (this *ListAnimationReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListAnimationReq)
	if !ok {
		that2, ok := that.(ListAnimationReq)
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
	if len(this.IDs) != len(that1.IDs) {
		return false
	}
	for i := range this.IDs {
		if !this.IDs[i].Equal(that1.IDs[i]) {
			return false
		}
	}
	if len(this.EntityIDs) != len(that1.EntityIDs) {
		return false
	}
	for i := range this.EntityIDs {
		if !this.EntityIDs[i].Equal(that1.EntityIDs[i]) {
			return false
		}
	}
	if this.Size_ != that1.Size_ {
		return false
	}
	if !bytes.Equal(this.State, that1.State) {
		return false
	}
	return true
}
func (this *ListAnimationResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListAnimationResp)
	if !ok {
		that2, ok := that.(ListAnimationResp)
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
	if len(this.Animations) != len(that1.Animations) {
		return false
	}
	for i := range this.Animations {
		if !this.Animations[i].Equal(&that1.Animations[i]) {
			return false
		}
	}
	if !bytes.Equal(this.State, that1.State) {
		return false
	}
	return true
}
func (this *ListAnimationReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&dto.ListAnimationReq{")
	s = append(s, "IDs: "+fmt.Sprintf("%#v", this.IDs)+",\n")
	s = append(s, "EntityIDs: "+fmt.Sprintf("%#v", this.EntityIDs)+",\n")
	s = append(s, "Size_: "+fmt.Sprintf("%#v", this.Size_)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ListAnimationResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dto.ListAnimationResp{")
	if this.Animations != nil {
		vs := make([]entity.Animation, len(this.Animations))
		for i := range vs {
			vs[i] = this.Animations[i]
		}
		s = append(s, "Animations: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAnimation(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ListAnimationReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAnimationReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAnimationReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintAnimation(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x22
	}
	if m.Size_ != 0 {
		i = encodeVarintAnimation(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x18
	}
	if len(m.EntityIDs) > 0 {
		for iNdEx := len(m.EntityIDs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.EntityIDs[iNdEx].Size()
				i -= size
				if _, err := m.EntityIDs[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintAnimation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.IDs) > 0 {
		for iNdEx := len(m.IDs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.IDs[iNdEx].Size()
				i -= size
				if _, err := m.IDs[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintAnimation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListAnimationResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAnimationResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAnimationResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintAnimation(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Animations) > 0 {
		for iNdEx := len(m.Animations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Animations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAnimation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAnimation(dAtA []byte, offset int, v uint64) int {
	offset -= sovAnimation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedListAnimationReq(r randyAnimation, easy bool) *ListAnimationReq {
	this := &ListAnimationReq{}
	v1 := r.Intn(10)
	this.IDs = make([]github_com_elojah_game_03_pkg_ulid.ID, v1)
	for i := 0; i < v1; i++ {
		v2 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
		this.IDs[i] = *v2
	}
	v3 := r.Intn(10)
	this.EntityIDs = make([]github_com_elojah_game_03_pkg_ulid.ID, v3)
	for i := 0; i < v3; i++ {
		v4 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
		this.EntityIDs[i] = *v4
	}
	this.Size_ = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Size_ *= -1
	}
	v5 := r.Intn(100)
	this.State = make([]byte, v5)
	for i := 0; i < v5; i++ {
		this.State[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedListAnimationResp(r randyAnimation, easy bool) *ListAnimationResp {
	this := &ListAnimationResp{}
	if r.Intn(5) != 0 {
		v6 := r.Intn(5)
		this.Animations = make([]entity.Animation, v6)
		for i := 0; i < v6; i++ {
			v7 := entity.NewPopulatedAnimation(r, easy)
			this.Animations[i] = *v7
		}
	}
	v8 := r.Intn(100)
	this.State = make([]byte, v8)
	for i := 0; i < v8; i++ {
		this.State[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyAnimation interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAnimation(r randyAnimation) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAnimation(r randyAnimation) string {
	v9 := r.Intn(100)
	tmps := make([]rune, v9)
	for i := 0; i < v9; i++ {
		tmps[i] = randUTF8RuneAnimation(r)
	}
	return string(tmps)
}
func randUnrecognizedAnimation(r randyAnimation, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAnimation(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAnimation(dAtA []byte, r randyAnimation, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(key))
		v10 := r.Int63()
		if r.Intn(2) == 0 {
			v10 *= -1
		}
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(v10))
	case 1:
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAnimation(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAnimation(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ListAnimationReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IDs) > 0 {
		for _, e := range m.IDs {
			l = e.Size()
			n += 1 + l + sovAnimation(uint64(l))
		}
	}
	if len(m.EntityIDs) > 0 {
		for _, e := range m.EntityIDs {
			l = e.Size()
			n += 1 + l + sovAnimation(uint64(l))
		}
	}
	if m.Size_ != 0 {
		n += 1 + sovAnimation(uint64(m.Size_))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovAnimation(uint64(l))
	}
	return n
}

func (m *ListAnimationResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Animations) > 0 {
		for _, e := range m.Animations {
			l = e.Size()
			n += 1 + l + sovAnimation(uint64(l))
		}
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovAnimation(uint64(l))
	}
	return n
}

func sovAnimation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAnimation(x uint64) (n int) {
	return sovAnimation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ListAnimationReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListAnimationReq{`,
		`IDs:` + fmt.Sprintf("%v", this.IDs) + `,`,
		`EntityIDs:` + fmt.Sprintf("%v", this.EntityIDs) + `,`,
		`Size_:` + fmt.Sprintf("%v", this.Size_) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ListAnimationResp) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForAnimations := "[]Animation{"
	for _, f := range this.Animations {
		repeatedStringForAnimations += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForAnimations += "}"
	s := strings.Join([]string{`&ListAnimationResp{`,
		`Animations:` + repeatedStringForAnimations + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAnimation(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ListAnimationReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnimation
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
			return fmt.Errorf("proto: ListAnimationReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAnimationReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
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
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_elojah_game_03_pkg_ulid.ID
			m.IDs = append(m.IDs, v)
			if err := m.IDs[len(m.IDs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityIDs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
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
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_elojah_game_03_pkg_ulid.ID
			m.EntityIDs = append(m.EntityIDs, v)
			if err := m.EntityIDs[len(m.EntityIDs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
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
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = append(m.State[:0], dAtA[iNdEx:postIndex]...)
			if m.State == nil {
				m.State = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAnimation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAnimation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAnimation
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
func (m *ListAnimationResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnimation
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
			return fmt.Errorf("proto: ListAnimationResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAnimationResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Animations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
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
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Animations = append(m.Animations, entity.Animation{})
			if err := m.Animations[len(m.Animations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimation
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
				return ErrInvalidLengthAnimation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = append(m.State[:0], dAtA[iNdEx:postIndex]...)
			if m.State == nil {
				m.State = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAnimation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAnimation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAnimation
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
func skipAnimation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAnimation
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
					return 0, ErrIntOverflowAnimation
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
					return 0, ErrIntOverflowAnimation
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
				return 0, ErrInvalidLengthAnimation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAnimation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAnimation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAnimation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAnimation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAnimation = fmt.Errorf("proto: unexpected end of group")
)
