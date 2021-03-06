// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/twitch/user.proto

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

type User struct {
	ID              string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	Login           string `protobuf:"bytes,2,opt,name=Login,proto3" json:"login"`
	DisplayName     string `protobuf:"bytes,3,opt,name=DisplayName,proto3" json:"display_name"`
	BroadcasterType string `protobuf:"bytes,4,opt,name=BroadcasterType,proto3" json:"broadcaster_type"`
	Description     string `protobuf:"bytes,5,opt,name=Description,proto3" json:"description"`
	ProfileImageURL string `protobuf:"bytes,6,opt,name=ProfileImageURL,proto3" json:"profile_image_url"`
	OfflineImageURL string `protobuf:"bytes,7,opt,name=OfflineImageURL,proto3" json:"offline_image_url"`
	ViewCount       int64  `protobuf:"varint,8,opt,name=ViewCount,proto3" json:"view_count"`
	Email           string `protobuf:"bytes,9,opt,name=Email,proto3" json:"email"`
	CreatedAt       string `protobuf:"bytes,10,opt,name=CreatedAt,proto3" json:"created_at"`
}

func (m *User) Reset()      { *m = User{} }
func (*User) ProtoMessage() {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_527fb6d1649d30ed, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *User) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetBroadcasterType() string {
	if m != nil {
		return m.BroadcasterType
	}
	return ""
}

func (m *User) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *User) GetProfileImageURL() string {
	if m != nil {
		return m.ProfileImageURL
	}
	return ""
}

func (m *User) GetOfflineImageURL() string {
	if m != nil {
		return m.OfflineImageURL
	}
	return ""
}

func (m *User) GetViewCount() int64 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "twitch.User")
	golang_proto.RegisterType((*User)(nil), "twitch.User")
}

func init() {
	proto.RegisterFile("github.com/elojah/game_03/pkg/twitch/user.proto", fileDescriptor_527fb6d1649d30ed)
}
func init() {
	golang_proto.RegisterFile("github.com/elojah/game_03/pkg/twitch/user.proto", fileDescriptor_527fb6d1649d30ed)
}

var fileDescriptor_527fb6d1649d30ed = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xb1, 0x8e, 0xd3, 0x30,
	0x1c, 0xc6, 0xe3, 0xf6, 0x5a, 0xa8, 0x8b, 0xe8, 0x11, 0x01, 0xb2, 0x18, 0x9c, 0x13, 0xd3, 0x0d,
	0x70, 0x01, 0x6e, 0x07, 0xd1, 0x0b, 0x43, 0xa5, 0x13, 0x20, 0x8b, 0x63, 0x8d, 0xdc, 0xc4, 0x4d,
	0x0d, 0x49, 0x1c, 0x39, 0x0e, 0x55, 0x37, 0x1e, 0x81, 0xc7, 0xe0, 0x11, 0x18, 0x6f, 0x64, 0xec,
	0x78, 0x53, 0x44, 0x5d, 0x21, 0xa1, 0x4c, 0x1d, 0x19, 0x4f, 0x71, 0x86, 0x44, 0xdd, 0xe2, 0xdf,
	0xf7, 0xfb, 0xbe, 0xe1, 0xaf, 0x40, 0x37, 0xe2, 0x6a, 0x59, 0xcc, 0xcf, 0x02, 0x91, 0xb8, 0x2c,
	0x16, 0x5f, 0xe8, 0xd2, 0x8d, 0x68, 0xc2, 0xfc, 0x17, 0xe7, 0x6e, 0xf6, 0x35, 0x72, 0xd5, 0x8a,
	0xab, 0x60, 0xe9, 0x16, 0x39, 0x93, 0x67, 0x99, 0x14, 0x4a, 0xd8, 0xc3, 0x06, 0x3d, 0x79, 0xde,
	0x29, 0x46, 0x22, 0x12, 0xae, 0x89, 0xe7, 0xc5, 0xc2, 0xbc, 0xcc, 0xc3, 0x7c, 0x35, 0xb5, 0xa7,
	0x7f, 0xfb, 0xf0, 0xe8, 0x2a, 0x67, 0xd2, 0x7e, 0x0c, 0x7b, 0x33, 0x0f, 0x81, 0x13, 0x70, 0x3a,
	0x9a, 0x0e, 0xab, 0xd2, 0xe9, 0xf1, 0x90, 0xf4, 0x66, 0x9e, 0xed, 0xc0, 0xc1, 0xa5, 0x88, 0x78,
	0x8a, 0x7a, 0x26, 0x1a, 0x55, 0xa5, 0x33, 0x88, 0x6b, 0x40, 0x1a, 0x6e, 0xbf, 0x82, 0x63, 0x8f,
	0xe7, 0x59, 0x4c, 0xd7, 0xef, 0x69, 0xc2, 0x50, 0xdf, 0x68, 0xc7, 0x55, 0xe9, 0xdc, 0x0b, 0x1b,
	0xec, 0xa7, 0x34, 0x61, 0xa4, 0x2b, 0xd9, 0xaf, 0xe1, 0x64, 0x2a, 0x05, 0x0d, 0x03, 0x9a, 0x2b,
	0x26, 0x3f, 0xad, 0x33, 0x86, 0x8e, 0x4c, 0xef, 0x61, 0x55, 0x3a, 0xc7, 0xf3, 0x36, 0xf2, 0xd5,
	0x3a, 0x63, 0xe4, 0x50, 0xb6, 0x5f, 0xc2, 0xb1, 0xc7, 0xf2, 0x40, 0xf2, 0x4c, 0x71, 0x91, 0xa2,
	0x81, 0xe9, 0x4e, 0xaa, 0xd2, 0x19, 0x87, 0x2d, 0x26, 0x5d, 0xc7, 0x7e, 0x03, 0x27, 0x1f, 0xa5,
	0x58, 0xf0, 0x98, 0xcd, 0x12, 0x1a, 0xb1, 0x2b, 0x72, 0x89, 0x86, 0xa6, 0xf6, 0xa8, 0x2a, 0x9d,
	0x07, 0x59, 0x13, 0xf9, 0xbc, 0xce, 0xfc, 0x42, 0xc6, 0xe4, 0xd0, 0xae, 0x07, 0x3e, 0x2c, 0x16,
	0x31, 0x4f, 0xdb, 0x81, 0x3b, 0xed, 0x80, 0x68, 0xa2, 0xee, 0xc0, 0x81, 0x6d, 0x3f, 0x83, 0xa3,
	0xcf, 0x9c, 0xad, 0x2e, 0x44, 0x91, 0x2a, 0x74, 0xf7, 0x04, 0x9c, 0xf6, 0xa7, 0xf7, 0xab, 0xd2,
	0x81, 0xdf, 0x38, 0x5b, 0xf9, 0x41, 0x4d, 0x49, 0x2b, 0xd4, 0x77, 0x7f, 0x97, 0x50, 0x1e, 0xa3,
	0x51, 0x7b, 0x77, 0x56, 0x03, 0xd2, 0xf0, 0x7a, 0xee, 0x42, 0x32, 0xaa, 0x58, 0xf8, 0x56, 0x21,
	0x68, 0x24, 0x33, 0x17, 0x34, 0xd0, 0xa7, 0x8a, 0xb4, 0xc2, 0xd4, 0xdb, 0x6c, 0xb1, 0x75, 0xb3,
	0xc5, 0xd6, 0x7e, 0x8b, 0xc1, 0xff, 0x2d, 0x06, 0xdf, 0x35, 0x06, 0x3f, 0x35, 0x06, 0xbf, 0x34,
	0x06, 0xd7, 0x1a, 0x83, 0xdf, 0x1a, 0x83, 0x8d, 0xc6, 0xe0, 0x8f, 0xc6, 0xe0, 0x9f, 0xc6, 0xd6,
	0x5e, 0x63, 0xf0, 0x63, 0x87, 0xad, 0xeb, 0x1d, 0x06, 0x9b, 0x1d, 0xb6, 0x6e, 0x76, 0xd8, 0x9a,
	0x0f, 0xcd, 0x4f, 0x73, 0x7e, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x89, 0xbb, 0x1c, 0x19, 0x9e, 0x02,
	0x00, 0x00,
}

func (this *User) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*User)
	if !ok {
		that2, ok := that.(User)
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
	if this.ID != that1.ID {
		return false
	}
	if this.Login != that1.Login {
		return false
	}
	if this.DisplayName != that1.DisplayName {
		return false
	}
	if this.BroadcasterType != that1.BroadcasterType {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.ProfileImageURL != that1.ProfileImageURL {
		return false
	}
	if this.OfflineImageURL != that1.OfflineImageURL {
		return false
	}
	if this.ViewCount != that1.ViewCount {
		return false
	}
	if this.Email != that1.Email {
		return false
	}
	if this.CreatedAt != that1.CreatedAt {
		return false
	}
	return true
}
func (this *User) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 14)
	s = append(s, "&twitch.User{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Login: "+fmt.Sprintf("%#v", this.Login)+",\n")
	s = append(s, "DisplayName: "+fmt.Sprintf("%#v", this.DisplayName)+",\n")
	s = append(s, "BroadcasterType: "+fmt.Sprintf("%#v", this.BroadcasterType)+",\n")
	s = append(s, "Description: "+fmt.Sprintf("%#v", this.Description)+",\n")
	s = append(s, "ProfileImageURL: "+fmt.Sprintf("%#v", this.ProfileImageURL)+",\n")
	s = append(s, "OfflineImageURL: "+fmt.Sprintf("%#v", this.OfflineImageURL)+",\n")
	s = append(s, "ViewCount: "+fmt.Sprintf("%#v", this.ViewCount)+",\n")
	s = append(s, "Email: "+fmt.Sprintf("%#v", this.Email)+",\n")
	s = append(s, "CreatedAt: "+fmt.Sprintf("%#v", this.CreatedAt)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringUser(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CreatedAt) > 0 {
		i -= len(m.CreatedAt)
		copy(dAtA[i:], m.CreatedAt)
		i = encodeVarintUser(dAtA, i, uint64(len(m.CreatedAt)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x4a
	}
	if m.ViewCount != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.ViewCount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.OfflineImageURL) > 0 {
		i -= len(m.OfflineImageURL)
		copy(dAtA[i:], m.OfflineImageURL)
		i = encodeVarintUser(dAtA, i, uint64(len(m.OfflineImageURL)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ProfileImageURL) > 0 {
		i -= len(m.ProfileImageURL)
		copy(dAtA[i:], m.ProfileImageURL)
		i = encodeVarintUser(dAtA, i, uint64(len(m.ProfileImageURL)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BroadcasterType) > 0 {
		i -= len(m.BroadcasterType)
		copy(dAtA[i:], m.BroadcasterType)
		i = encodeVarintUser(dAtA, i, uint64(len(m.BroadcasterType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DisplayName) > 0 {
		i -= len(m.DisplayName)
		copy(dAtA[i:], m.DisplayName)
		i = encodeVarintUser(dAtA, i, uint64(len(m.DisplayName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Login) > 0 {
		i -= len(m.Login)
		copy(dAtA[i:], m.Login)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Login)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintUser(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedUser(r randyUser, easy bool) *User {
	this := &User{}
	this.ID = string(randStringUser(r))
	this.Login = string(randStringUser(r))
	this.DisplayName = string(randStringUser(r))
	this.BroadcasterType = string(randStringUser(r))
	this.Description = string(randStringUser(r))
	this.ProfileImageURL = string(randStringUser(r))
	this.OfflineImageURL = string(randStringUser(r))
	this.ViewCount = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.ViewCount *= -1
	}
	this.Email = string(randStringUser(r))
	this.CreatedAt = string(randStringUser(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyUser interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneUser(r randyUser) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringUser(r randyUser) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneUser(r)
	}
	return string(tmps)
}
func randUnrecognizedUser(r randyUser, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldUser(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldUser(dAtA []byte, r randyUser, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateUser(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateUser(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateUser(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateUser(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Login)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.BroadcasterType)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.ProfileImageURL)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.OfflineImageURL)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.ViewCount != 0 {
		n += 1 + sovUser(uint64(m.ViewCount))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *User) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&User{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Login:` + fmt.Sprintf("%v", this.Login) + `,`,
		`DisplayName:` + fmt.Sprintf("%v", this.DisplayName) + `,`,
		`BroadcasterType:` + fmt.Sprintf("%v", this.BroadcasterType) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`ProfileImageURL:` + fmt.Sprintf("%v", this.ProfileImageURL) + `,`,
		`OfflineImageURL:` + fmt.Sprintf("%v", this.OfflineImageURL) + `,`,
		`ViewCount:` + fmt.Sprintf("%v", this.ViewCount) + `,`,
		`Email:` + fmt.Sprintf("%v", this.Email) + `,`,
		`CreatedAt:` + fmt.Sprintf("%v", this.CreatedAt) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringUser(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Login", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Login = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BroadcasterType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BroadcasterType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileImageURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfileImageURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfflineImageURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OfflineImageURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ViewCount", wireType)
			}
			m.ViewCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ViewCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
