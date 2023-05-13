// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/ability/dto/ability.proto

package dto

import (
	bytes "bytes"
	fmt "fmt"
	ability "github.com/elojah/game_03/pkg/ability"
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

type ListAbilityReq struct {
	EntityID github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=EntityID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"EntityID"`
	Size_    int64                                 `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	State    []byte                                `protobuf:"bytes,3,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *ListAbilityReq) Reset()      { *m = ListAbilityReq{} }
func (*ListAbilityReq) ProtoMessage() {}
func (*ListAbilityReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ac18f47e846bf2a, []int{0}
}
func (m *ListAbilityReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAbilityReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAbilityReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAbilityReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAbilityReq.Merge(m, src)
}
func (m *ListAbilityReq) XXX_Size() int {
	return m.Size()
}
func (m *ListAbilityReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAbilityReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListAbilityReq proto.InternalMessageInfo

type ListAbilityResp struct {
	Abilities []ability.A `protobuf:"bytes,1,rep,name=Abilities,proto3" json:"Abilities"`
	State     []byte      `protobuf:"bytes,2,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *ListAbilityResp) Reset()      { *m = ListAbilityResp{} }
func (*ListAbilityResp) ProtoMessage() {}
func (*ListAbilityResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ac18f47e846bf2a, []int{1}
}
func (m *ListAbilityResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListAbilityResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListAbilityResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListAbilityResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAbilityResp.Merge(m, src)
}
func (m *ListAbilityResp) XXX_Size() int {
	return m.Size()
}
func (m *ListAbilityResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAbilityResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListAbilityResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListAbilityReq)(nil), "dto.ListAbilityReq")
	golang_proto.RegisterType((*ListAbilityReq)(nil), "dto.ListAbilityReq")
	proto.RegisterType((*ListAbilityResp)(nil), "dto.ListAbilityResp")
	golang_proto.RegisterType((*ListAbilityResp)(nil), "dto.ListAbilityResp")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/ability/dto/ability.proto", fileDescriptor_1ac18f47e846bf2a)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/ability/dto/ability.proto", fileDescriptor_1ac18f47e846bf2a)
}

var fileDescriptor_1ac18f47e846bf2a = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3f, 0x4f, 0x32, 0x31,
	0x1c, 0xc7, 0xfb, 0xe3, 0x78, 0x9e, 0x68, 0x35, 0x9a, 0x34, 0x0e, 0x17, 0x86, 0x1f, 0x84, 0xc4,
	0x84, 0x85, 0xab, 0x91, 0xc1, 0x19, 0x82, 0x89, 0x24, 0x4e, 0xc7, 0xe0, 0x68, 0xee, 0xa4, 0x1e,
	0x55, 0xb0, 0x28, 0x65, 0xc0, 0xc9, 0xc5, 0xdd, 0x97, 0xe1, 0x4b, 0x70, 0x64, 0x64, 0x64, 0x24,
	0x0e, 0xc4, 0xeb, 0x2d, 0x8e, 0x8c, 0x8e, 0x86, 0x5e, 0x10, 0x58, 0x8c, 0xdb, 0xf7, 0xd3, 0xf6,
	0xfb, 0x27, 0xa5, 0x27, 0x91, 0xd4, 0xed, 0x41, 0xe8, 0x5d, 0xa9, 0x2e, 0x17, 0x1d, 0x75, 0x13,
	0xb4, 0x79, 0x14, 0x74, 0xc5, 0xe5, 0x51, 0x85, 0xf7, 0x6e, 0x23, 0x1e, 0x84, 0xb2, 0x23, 0xf5,
	0x90, 0xb7, 0xb4, 0x5a, 0x6a, 0xaf, 0xf7, 0xa0, 0xb4, 0x62, 0x4e, 0x4b, 0xab, 0x5c, 0x79, 0xcd,
	0x1d, 0xa9, 0x48, 0x71, 0x7b, 0x17, 0x0e, 0xae, 0x2d, 0x59, 0xb0, 0x2a, 0xf5, 0xe4, 0x2a, 0x7f,
	0x2b, 0xdb, 0x28, 0x2a, 0x3e, 0x03, 0xdd, 0x3b, 0x97, 0x7d, 0x5d, 0x4d, 0x4f, 0x7d, 0x71, 0xcf,
	0x1a, 0x74, 0xeb, 0xf4, 0x4e, 0x4b, 0x3d, 0x6c, 0xd4, 0x5d, 0x28, 0x40, 0x69, 0xb7, 0x56, 0x1e,
	0xcf, 0xf2, 0xe4, 0x7d, 0x96, 0x3f, 0xfc, 0xbd, 0x61, 0xd0, 0x91, 0x2d, 0xaf, 0x51, 0xf7, 0x7f,
	0xec, 0x8c, 0xd1, 0x6c, 0x53, 0x3e, 0x0a, 0x37, 0x53, 0x80, 0x92, 0xe3, 0x5b, 0xcd, 0x0e, 0xe8,
	0xbf, 0xa6, 0x0e, 0xb4, 0x70, 0x9d, 0x45, 0xb6, 0x9f, 0x42, 0xf1, 0x82, 0xee, 0x6f, 0xcc, 0xe8,
	0xf7, 0x98, 0x47, 0xb7, 0x53, 0x94, 0xa2, 0xef, 0x42, 0xc1, 0x29, 0xed, 0x1c, 0x53, 0x6f, 0xb9,
	0xbe, 0x5a, 0xcb, 0x2e, 0x46, 0xf9, 0xab, 0x27, 0xab, 0xe0, 0xcc, 0x5a, 0x70, 0xed, 0x6c, 0x1c,
	0x23, 0x99, 0xc4, 0x48, 0xa6, 0x31, 0x92, 0x79, 0x8c, 0xf0, 0x15, 0x23, 0x3c, 0x19, 0x84, 0x57,
	0x83, 0xf0, 0x66, 0x10, 0x46, 0x06, 0x61, 0x6c, 0x10, 0x26, 0x06, 0xe1, 0xc3, 0x20, 0x7c, 0x1a,
	0x24, 0x73, 0x83, 0xf0, 0x92, 0x20, 0x19, 0x25, 0x08, 0x93, 0x04, 0xc9, 0x34, 0x41, 0x12, 0xfe,
	0xb7, 0x3f, 0x56, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x73, 0x5a, 0x86, 0xd5, 0x01, 0x00,
	0x00,
}

func (this *ListAbilityReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListAbilityReq)
	if !ok {
		that2, ok := that.(ListAbilityReq)
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
	if this.Size_ != that1.Size_ {
		return false
	}
	if !bytes.Equal(this.State, that1.State) {
		return false
	}
	return true
}
func (this *ListAbilityResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListAbilityResp)
	if !ok {
		that2, ok := that.(ListAbilityResp)
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
	if len(this.Abilities) != len(that1.Abilities) {
		return false
	}
	for i := range this.Abilities {
		if !this.Abilities[i].Equal(&that1.Abilities[i]) {
			return false
		}
	}
	if !bytes.Equal(this.State, that1.State) {
		return false
	}
	return true
}
func (this *ListAbilityReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&dto.ListAbilityReq{")
	s = append(s, "EntityID: "+fmt.Sprintf("%#v", this.EntityID)+",\n")
	s = append(s, "Size_: "+fmt.Sprintf("%#v", this.Size_)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ListAbilityResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dto.ListAbilityResp{")
	if this.Abilities != nil {
		vs := make([]ability.A, len(this.Abilities))
		for i := range vs {
			vs[i] = this.Abilities[i]
		}
		s = append(s, "Abilities: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAbility(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ListAbilityReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAbilityReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAbilityReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintAbility(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Size_ != 0 {
		i = encodeVarintAbility(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.EntityID.Size()
		i -= size
		if _, err := m.EntityID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAbility(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ListAbilityResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListAbilityResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListAbilityResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintAbility(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Abilities) > 0 {
		for iNdEx := len(m.Abilities) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Abilities[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAbility(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAbility(dAtA []byte, offset int, v uint64) int {
	offset -= sovAbility(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedListAbilityReq(r randyAbility, easy bool) *ListAbilityReq {
	this := &ListAbilityReq{}
	v1 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.EntityID = *v1
	this.Size_ = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Size_ *= -1
	}
	v2 := r.Intn(100)
	this.State = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.State[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedListAbilityResp(r randyAbility, easy bool) *ListAbilityResp {
	this := &ListAbilityResp{}
	if r.Intn(5) != 0 {
		v3 := r.Intn(5)
		this.Abilities = make([]ability.A, v3)
		for i := 0; i < v3; i++ {
			v4 := ability.NewPopulatedA(r, easy)
			this.Abilities[i] = *v4
		}
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

type randyAbility interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAbility(r randyAbility) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAbility(r randyAbility) string {
	v6 := r.Intn(100)
	tmps := make([]rune, v6)
	for i := 0; i < v6; i++ {
		tmps[i] = randUTF8RuneAbility(r)
	}
	return string(tmps)
}
func randUnrecognizedAbility(r randyAbility, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAbility(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAbility(dAtA []byte, r randyAbility, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAbility(dAtA, uint64(key))
		v7 := r.Int63()
		if r.Intn(2) == 0 {
			v7 *= -1
		}
		dAtA = encodeVarintPopulateAbility(dAtA, uint64(v7))
	case 1:
		dAtA = encodeVarintPopulateAbility(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAbility(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAbility(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAbility(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAbility(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ListAbilityReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EntityID.Size()
	n += 1 + l + sovAbility(uint64(l))
	if m.Size_ != 0 {
		n += 1 + sovAbility(uint64(m.Size_))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovAbility(uint64(l))
	}
	return n
}

func (m *ListAbilityResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Abilities) > 0 {
		for _, e := range m.Abilities {
			l = e.Size()
			n += 1 + l + sovAbility(uint64(l))
		}
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovAbility(uint64(l))
	}
	return n
}

func sovAbility(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAbility(x uint64) (n int) {
	return sovAbility(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ListAbilityReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListAbilityReq{`,
		`EntityID:` + fmt.Sprintf("%v", this.EntityID) + `,`,
		`Size_:` + fmt.Sprintf("%v", this.Size_) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ListAbilityResp) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForAbilities := "[]A{"
	for _, f := range this.Abilities {
		repeatedStringForAbilities += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForAbilities += "}"
	s := strings.Join([]string{`&ListAbilityResp{`,
		`Abilities:` + repeatedStringForAbilities + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAbility(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ListAbilityReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbility
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
			return fmt.Errorf("proto: ListAbilityReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAbilityReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
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
				return ErrInvalidLengthAbility
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAbility
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EntityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
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
				return ErrInvalidLengthAbility
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAbility
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
			skippy, err := skipAbility(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAbility
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
func (m *ListAbilityResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAbility
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
			return fmt.Errorf("proto: ListAbilityResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListAbilityResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abilities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAbility
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
				return ErrInvalidLengthAbility
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAbility
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Abilities = append(m.Abilities, ability.A{})
			if err := m.Abilities[len(m.Abilities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowAbility
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
				return ErrInvalidLengthAbility
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAbility
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
			skippy, err := skipAbility(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAbility
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
func skipAbility(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAbility
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
					return 0, ErrIntOverflowAbility
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
					return 0, ErrIntOverflowAbility
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
				return 0, ErrInvalidLengthAbility
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAbility
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAbility
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAbility        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAbility          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAbility = fmt.Errorf("proto: unexpected end of group")
)
