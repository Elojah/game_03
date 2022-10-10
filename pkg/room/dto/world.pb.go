// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/room/dto/world.proto

package dto

import (
	fmt "fmt"
	room "github.com/elojah/game_03/pkg/room"
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

type ListWorldReq struct {
	IDs []github_com_elojah_game_03_pkg_ulid.ID `protobuf:"bytes,1,rep,name=IDs,proto3,customtype=github.com/elojah/game_03/pkg/ulid.ID" json:"IDs"`
}

func (m *ListWorldReq) Reset()      { *m = ListWorldReq{} }
func (*ListWorldReq) ProtoMessage() {}
func (*ListWorldReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f262c284d8b209a5, []int{0}
}
func (m *ListWorldReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListWorldReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListWorldReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListWorldReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorldReq.Merge(m, src)
}
func (m *ListWorldReq) XXX_Size() int {
	return m.Size()
}
func (m *ListWorldReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorldReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorldReq proto.InternalMessageInfo

type ListWorldResp struct {
	Worlds []room.World `protobuf:"bytes,1,rep,name=Worlds,proto3" json:"Worlds"`
}

func (m *ListWorldResp) Reset()      { *m = ListWorldResp{} }
func (*ListWorldResp) ProtoMessage() {}
func (*ListWorldResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f262c284d8b209a5, []int{1}
}
func (m *ListWorldResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListWorldResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListWorldResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListWorldResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorldResp.Merge(m, src)
}
func (m *ListWorldResp) XXX_Size() int {
	return m.Size()
}
func (m *ListWorldResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorldResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorldResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListWorldReq)(nil), "dto.ListWorldReq")
	golang_proto.RegisterType((*ListWorldReq)(nil), "dto.ListWorldReq")
	proto.RegisterType((*ListWorldResp)(nil), "dto.ListWorldResp")
	golang_proto.RegisterType((*ListWorldResp)(nil), "dto.ListWorldResp")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/room/dto/world.proto", fileDescriptor_f262c284d8b209a5)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/room/dto/world.proto", fileDescriptor_f262c284d8b209a5)
}

var fileDescriptor_f262c284d8b209a5 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcd, 0xc9, 0xcf, 0x4a, 0xcc, 0xd0, 0x4f, 0x4f,
	0xcc, 0x4d, 0x8d, 0x37, 0x30, 0xd6, 0x2f, 0xc8, 0x4e, 0xd7, 0x2f, 0xca, 0xcf, 0xcf, 0xd5, 0x4f,
	0x29, 0xc9, 0xd7, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4e, 0x29, 0xc9, 0x97, 0xd2, 0x45, 0xd2, 0x98, 0x9e, 0x9f, 0x9e, 0xaf, 0x0f, 0x96, 0x4b, 0x2a,
	0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x47, 0x4a, 0x8f, 0x08, 0x7b, 0x90, 0xec, 0x50,
	0xf2, 0xe7, 0xe2, 0xf1, 0xc9, 0x2c, 0x2e, 0x09, 0x07, 0x09, 0x05, 0xa5, 0x16, 0x0a, 0xd9, 0x73,
	0x31, 0x7b, 0xba, 0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0xf0, 0x38, 0xe9, 0x9e, 0xb8, 0x27, 0xcf,
	0x70, 0xeb, 0x9e, 0xbc, 0x2a, 0x7e, 0x43, 0x4b, 0x73, 0x32, 0x53, 0xf4, 0x3c, 0x5d, 0x82, 0x40,
	0x3a, 0x95, 0xac, 0xb8, 0x78, 0x91, 0x0c, 0x2c, 0x2e, 0x10, 0xd2, 0xe4, 0x62, 0x03, 0x73, 0x20,
	0x86, 0x72, 0x1b, 0x71, 0xeb, 0x81, 0x1c, 0xa1, 0x07, 0x16, 0x73, 0x62, 0x01, 0xd9, 0x10, 0x04,
	0x55, 0xe0, 0xe4, 0x71, 0xe2, 0xa1, 0x1c, 0xc3, 0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31,
	0x7c, 0x78, 0x28, 0xc7, 0xf8, 0xe3, 0xa1, 0x1c, 0x63, 0xc3, 0x23, 0x39, 0xc6, 0x15, 0x8f, 0xe4,
	0x18, 0x77, 0x3c, 0x92, 0x63, 0x3c, 0xf0, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f,
	0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x0e, 0x3c, 0x96, 0x63, 0xbc, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x24, 0x36, 0xb0, 0xef, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0xcd, 0xf0, 0x11, 0x77,
	0x01, 0x00, 0x00,
}

func (this *ListWorldReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListWorldReq)
	if !ok {
		that2, ok := that.(ListWorldReq)
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
	return true
}
func (this *ListWorldResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListWorldResp)
	if !ok {
		that2, ok := that.(ListWorldResp)
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
	if len(this.Worlds) != len(that1.Worlds) {
		return false
	}
	for i := range this.Worlds {
		if !this.Worlds[i].Equal(&that1.Worlds[i]) {
			return false
		}
	}
	return true
}
func (this *ListWorldReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dto.ListWorldReq{")
	s = append(s, "IDs: "+fmt.Sprintf("%#v", this.IDs)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ListWorldResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dto.ListWorldResp{")
	if this.Worlds != nil {
		vs := make([]room.World, len(this.Worlds))
		for i := range vs {
			vs[i] = this.Worlds[i]
		}
		s = append(s, "Worlds: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringWorld(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ListWorldReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListWorldReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListWorldReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IDs) > 0 {
		for iNdEx := len(m.IDs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.IDs[iNdEx].Size()
				i -= size
				if _, err := m.IDs[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintWorld(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListWorldResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListWorldResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListWorldResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Worlds) > 0 {
		for iNdEx := len(m.Worlds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Worlds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWorld(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorld(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorld(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedListWorldReq(r randyWorld, easy bool) *ListWorldReq {
	this := &ListWorldReq{}
	v1 := r.Intn(10)
	this.IDs = make([]github_com_elojah_game_03_pkg_ulid.ID, v1)
	for i := 0; i < v1; i++ {
		v2 := github_com_elojah_game_03_pkg_ulid.NewPopulatedID(r)
		this.IDs[i] = *v2
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedListWorldResp(r randyWorld, easy bool) *ListWorldResp {
	this := &ListWorldResp{}
	if r.Intn(5) != 0 {
		v3 := r.Intn(5)
		this.Worlds = make([]room.World, v3)
		for i := 0; i < v3; i++ {
			v4 := room.NewPopulatedWorld(r, easy)
			this.Worlds[i] = *v4
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyWorld interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneWorld(r randyWorld) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringWorld(r randyWorld) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneWorld(r)
	}
	return string(tmps)
}
func randUnrecognizedWorld(r randyWorld, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldWorld(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldWorld(dAtA []byte, r randyWorld, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateWorld(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateWorld(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateWorld(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateWorld(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateWorld(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateWorld(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateWorld(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ListWorldReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IDs) > 0 {
		for _, e := range m.IDs {
			l = e.Size()
			n += 1 + l + sovWorld(uint64(l))
		}
	}
	return n
}

func (m *ListWorldResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Worlds) > 0 {
		for _, e := range m.Worlds {
			l = e.Size()
			n += 1 + l + sovWorld(uint64(l))
		}
	}
	return n
}

func sovWorld(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorld(x uint64) (n int) {
	return sovWorld(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ListWorldReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListWorldReq{`,
		`IDs:` + fmt.Sprintf("%v", this.IDs) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ListWorldResp) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForWorlds := "[]World{"
	for _, f := range this.Worlds {
		repeatedStringForWorlds += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForWorlds += "}"
	s := strings.Join([]string{`&ListWorldResp{`,
		`Worlds:` + repeatedStringForWorlds + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringWorld(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ListWorldReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorld
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
			return fmt.Errorf("proto: ListWorldReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListWorldReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorld
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
				return ErrInvalidLengthWorld
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWorld
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
		default:
			iNdEx = preIndex
			skippy, err := skipWorld(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWorld
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
func (m *ListWorldResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorld
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
			return fmt.Errorf("proto: ListWorldResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListWorldResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Worlds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorld
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
				return ErrInvalidLengthWorld
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorld
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Worlds = append(m.Worlds, room.World{})
			if err := m.Worlds[len(m.Worlds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorld(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWorld
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
func skipWorld(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorld
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
					return 0, ErrIntOverflowWorld
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
					return 0, ErrIntOverflowWorld
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
				return 0, ErrInvalidLengthWorld
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWorld
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWorld
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWorld        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorld          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWorld = fmt.Errorf("proto: unexpected end of group")
)
