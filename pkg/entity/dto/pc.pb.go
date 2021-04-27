// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/entity/dto/pc.proto

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

type CreatePCReq struct {
	RoomID github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=RoomID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"RoomID"`
}

func (m *CreatePCReq) Reset()      { *m = CreatePCReq{} }
func (*CreatePCReq) ProtoMessage() {}
func (*CreatePCReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3acf6b9c23e297b, []int{0}
}
func (m *CreatePCReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePCReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePCReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePCReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePCReq.Merge(m, src)
}
func (m *CreatePCReq) XXX_Size() int {
	return m.Size()
}
func (m *CreatePCReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePCReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePCReq proto.InternalMessageInfo

type ListPCReq struct {
	RoomID github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,opt,name=RoomID,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"RoomID"`
	Size_  int64                                 `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	State  []byte                                `protobuf:"bytes,3,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *ListPCReq) Reset()      { *m = ListPCReq{} }
func (*ListPCReq) ProtoMessage() {}
func (*ListPCReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3acf6b9c23e297b, []int{1}
}
func (m *ListPCReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListPCReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListPCReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListPCReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPCReq.Merge(m, src)
}
func (m *ListPCReq) XXX_Size() int {
	return m.Size()
}
func (m *ListPCReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPCReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListPCReq proto.InternalMessageInfo

func (m *ListPCReq) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *ListPCReq) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

type ListPCResp struct {
	PCs   []entity.PC `protobuf:"bytes,1,rep,name=PCs,proto3" json:"PCs"`
	State []byte      `protobuf:"bytes,2,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *ListPCResp) Reset()      { *m = ListPCResp{} }
func (*ListPCResp) ProtoMessage() {}
func (*ListPCResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3acf6b9c23e297b, []int{2}
}
func (m *ListPCResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListPCResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListPCResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListPCResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPCResp.Merge(m, src)
}
func (m *ListPCResp) XXX_Size() int {
	return m.Size()
}
func (m *ListPCResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPCResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListPCResp proto.InternalMessageInfo

func (m *ListPCResp) GetPCs() []entity.PC {
	if m != nil {
		return m.PCs
	}
	return nil
}

func (m *ListPCResp) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatePCReq)(nil), "dto.CreatePCReq")
	golang_proto.RegisterType((*CreatePCReq)(nil), "dto.CreatePCReq")
	proto.RegisterType((*ListPCReq)(nil), "dto.ListPCReq")
	golang_proto.RegisterType((*ListPCReq)(nil), "dto.ListPCReq")
	proto.RegisterType((*ListPCResp)(nil), "dto.ListPCResp")
	golang_proto.RegisterType((*ListPCResp)(nil), "dto.ListPCResp")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/entity/dto/pc.proto", fileDescriptor_b3acf6b9c23e297b)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/entity/dto/pc.proto", fileDescriptor_b3acf6b9c23e297b)
}

var fileDescriptor_b3acf6b9c23e297b = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xb1, 0x4e, 0x32, 0x41,
	0x14, 0x85, 0xe7, 0xb2, 0xfc, 0x24, 0xff, 0x60, 0xb5, 0xb1, 0xd8, 0x58, 0x5c, 0xc8, 0x26, 0x26,
	0x34, 0xec, 0xa8, 0xbc, 0x01, 0x8b, 0x26, 0x24, 0x16, 0x64, 0xb0, 0x37, 0x0b, 0x8c, 0xcb, 0x2a,
	0x38, 0x2b, 0x0c, 0x85, 0xc6, 0xc2, 0x47, 0xf0, 0x31, 0x7c, 0x04, 0x4b, 0x4a, 0x4a, 0x4a, 0x62,
	0x41, 0xdc, 0xd9, 0xc6, 0x92, 0xd2, 0xd2, 0x30, 0x2b, 0x11, 0x0b, 0xad, 0xec, 0xee, 0xc9, 0x99,
	0xf3, 0x9d, 0xc9, 0xa1, 0x87, 0x61, 0xa4, 0xfa, 0x93, 0x8e, 0xd7, 0x95, 0x43, 0x26, 0x06, 0xf2,
	0x32, 0xe8, 0xb3, 0x30, 0x18, 0x8a, 0xf3, 0x83, 0x1a, 0x8b, 0xaf, 0x42, 0x26, 0xae, 0x55, 0xa4,
	0x6e, 0x59, 0x4f, 0x49, 0x16, 0x77, 0xbd, 0x78, 0x24, 0x95, 0xb4, 0xad, 0x9e, 0x92, 0x7b, 0xd5,
	0xad, 0x5c, 0x28, 0x43, 0xc9, 0x8c, 0xd7, 0x99, 0x5c, 0x18, 0x65, 0x84, 0xb9, 0xb2, 0xcc, 0xb7,
	0xe7, 0x3f, 0xd7, 0x6c, 0x2a, 0xdc, 0x33, 0x5a, 0xf4, 0x47, 0x22, 0x50, 0xa2, 0xe5, 0x73, 0x71,
	0x63, 0x1f, 0xd3, 0x02, 0x97, 0x72, 0xd8, 0x6c, 0x38, 0x50, 0x86, 0xca, 0x4e, 0xbd, 0x3a, 0x5b,
	0x96, 0xc8, 0xcb, 0xb2, 0xb4, 0xff, 0x3b, 0x75, 0x32, 0x88, 0x7a, 0x5e, 0xb3, 0xc1, 0x3f, 0xc3,
	0xee, 0x3d, 0xfd, 0x7f, 0x1a, 0x8d, 0xd5, 0x5f, 0x32, 0x6d, 0x9b, 0xe6, 0xdb, 0xd1, 0x9d, 0x70,
	0x72, 0x65, 0xa8, 0x58, 0xdc, 0xdc, 0xf6, 0x2e, 0xfd, 0xd7, 0x56, 0x81, 0x12, 0x8e, 0xb5, 0x26,
	0xf3, 0x4c, 0xb8, 0x27, 0x94, 0x6e, 0xda, 0xc7, 0xb1, 0xed, 0x52, 0xab, 0xe5, 0x8f, 0x1d, 0x28,
	0x5b, 0x95, 0xe2, 0x11, 0xf5, 0xb2, 0x01, 0xbc, 0x96, 0x5f, 0xcf, 0xaf, 0xff, 0xc1, 0xd7, 0xe6,
	0x17, 0x27, 0xb7, 0xc5, 0xa9, 0x37, 0xe6, 0x09, 0x92, 0x45, 0x82, 0x64, 0x95, 0x20, 0xbc, 0x27,
	0x08, 0x0f, 0x1a, 0xe1, 0x49, 0x23, 0x3c, 0x6b, 0x84, 0xa9, 0x46, 0x98, 0x69, 0x84, 0xb9, 0x46,
	0x78, 0xd5, 0x08, 0x6f, 0x1a, 0xc9, 0x4a, 0x23, 0x3c, 0xa6, 0x48, 0xa6, 0x29, 0xc2, 0x3c, 0x45,
	0xb2, 0x48, 0x91, 0x74, 0x0a, 0x66, 0xe8, 0xda, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0xba,
	0x94, 0x47, 0x00, 0x02, 0x00, 0x00,
}

func (this *CreatePCReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreatePCReq)
	if !ok {
		that2, ok := that.(CreatePCReq)
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
	if !this.RoomID.Equal(that1.RoomID) {
		return false
	}
	return true
}
func (this *ListPCReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListPCReq)
	if !ok {
		that2, ok := that.(ListPCReq)
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
	if !this.RoomID.Equal(that1.RoomID) {
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
func (this *ListPCResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListPCResp)
	if !ok {
		that2, ok := that.(ListPCResp)
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
	if len(this.PCs) != len(that1.PCs) {
		return false
	}
	for i := range this.PCs {
		if !this.PCs[i].Equal(&that1.PCs[i]) {
			return false
		}
	}
	if !bytes.Equal(this.State, that1.State) {
		return false
	}
	return true
}
func (this *CreatePCReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dto.CreatePCReq{")
	s = append(s, "RoomID: "+fmt.Sprintf("%#v", this.RoomID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ListPCReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&dto.ListPCReq{")
	s = append(s, "RoomID: "+fmt.Sprintf("%#v", this.RoomID)+",\n")
	s = append(s, "Size_: "+fmt.Sprintf("%#v", this.Size_)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ListPCResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dto.ListPCResp{")
	if this.PCs != nil {
		vs := make([]entity.PC, len(this.PCs))
		for i := range vs {
			vs[i] = this.PCs[i]
		}
		s = append(s, "PCs: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
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
func (m *CreatePCReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePCReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePCReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RoomID.Size()
		i -= size
		if _, err := m.RoomID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ListPCReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListPCReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListPCReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintPc(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Size_ != 0 {
		i = encodeVarintPc(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.RoomID.Size()
		i -= size
		if _, err := m.RoomID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ListPCResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListPCResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListPCResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintPc(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PCs) > 0 {
		for iNdEx := len(m.PCs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PCs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
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
func NewPopulatedCreatePCReq(r randyPc, easy bool) *CreatePCReq {
	this := &CreatePCReq{}
	v1 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.RoomID = *v1
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedListPCReq(r randyPc, easy bool) *ListPCReq {
	this := &ListPCReq{}
	v2 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
	this.RoomID = *v2
	this.Size_ = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Size_ *= -1
	}
	v3 := r.Intn(100)
	this.State = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.State[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedListPCResp(r randyPc, easy bool) *ListPCResp {
	this := &ListPCResp{}
	if r.Intn(5) != 0 {
		v4 := r.Intn(5)
		this.PCs = make([]entity.PC, v4)
		for i := 0; i < v4; i++ {
			v5 := entity.NewPopulatedPC(r, easy)
			this.PCs[i] = *v5
		}
	}
	v6 := r.Intn(100)
	this.State = make([]byte, v6)
	for i := 0; i < v6; i++ {
		this.State[i] = byte(r.Intn(256))
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
	v7 := r.Intn(100)
	tmps := make([]rune, v7)
	for i := 0; i < v7; i++ {
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
		v8 := r.Int63()
		if r.Intn(2) == 0 {
			v8 *= -1
		}
		dAtA = encodeVarintPopulatePc(dAtA, uint64(v8))
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
func (m *CreatePCReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.RoomID.Size()
	n += 1 + l + sovPc(uint64(l))
	return n
}

func (m *ListPCReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.RoomID.Size()
	n += 1 + l + sovPc(uint64(l))
	if m.Size_ != 0 {
		n += 1 + sovPc(uint64(m.Size_))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovPc(uint64(l))
	}
	return n
}

func (m *ListPCResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PCs) > 0 {
		for _, e := range m.PCs {
			l = e.Size()
			n += 1 + l + sovPc(uint64(l))
		}
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovPc(uint64(l))
	}
	return n
}

func sovPc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPc(x uint64) (n int) {
	return sovPc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CreatePCReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreatePCReq{`,
		`RoomID:` + fmt.Sprintf("%v", this.RoomID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ListPCReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListPCReq{`,
		`RoomID:` + fmt.Sprintf("%v", this.RoomID) + `,`,
		`Size_:` + fmt.Sprintf("%v", this.Size_) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ListPCResp) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForPCs := "[]PC{"
	for _, f := range this.PCs {
		repeatedStringForPCs += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForPCs += "}"
	s := strings.Join([]string{`&ListPCResp{`,
		`PCs:` + repeatedStringForPCs + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
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
func (m *CreatePCReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CreatePCReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePCReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomID", wireType)
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
			if err := m.RoomID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *ListPCReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListPCReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListPCReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomID", wireType)
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
			if err := m.RoomID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowPc
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
			m.State = append(m.State[:0], dAtA[iNdEx:postIndex]...)
			if m.State == nil {
				m.State = []byte{}
			}
			iNdEx = postIndex
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
func (m *ListPCResp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListPCResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListPCResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PCs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPc
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
				return ErrInvalidLengthPc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PCs = append(m.PCs, entity.PC{})
			if err := m.PCs[len(m.PCs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
			m.State = append(m.State[:0], dAtA[iNdEx:postIndex]...)
			if m.State == nil {
				m.State = []byte{}
			}
			iNdEx = postIndex
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
