// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/elojah/game_03/pkg/room/room.proto

package room

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	go_parser "go/parser"
	math "math"
	math_rand "math/rand"
	testing "testing"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestRProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedR(popr, false)
	dAtA, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &R{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_gogo_protobuf_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestRMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedR(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &R{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkRProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*R, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedR(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_gogo_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkRProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_gogo_protobuf_proto.Marshal(NewPopulatedR(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &R{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_gogo_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestPublicProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedPublic(popr, false)
	dAtA, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Public{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_gogo_protobuf_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestPublicMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedPublic(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Public{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkPublicProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*Public, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedPublic(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dAtA, err := github_com_gogo_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(dAtA)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkPublicProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		dAtA, err := github_com_gogo_protobuf_proto.Marshal(NewPopulatedPublic(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = dAtA
	}
	msg := &Public{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_gogo_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestRJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedR(popr, true)
	marshaler := github_com_gogo_protobuf_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &R{}
	err = github_com_gogo_protobuf_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestPublicJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedPublic(popr, true)
	marshaler := github_com_gogo_protobuf_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Public{}
	err = github_com_gogo_protobuf_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestRProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedR(popr, true)
	dAtA := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &R{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestRProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedR(popr, true)
	dAtA := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &R{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestPublicProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedPublic(popr, true)
	dAtA := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &Public{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestPublicProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedPublic(popr, true)
	dAtA := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &Public{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestRGoString(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedR(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestPublicGoString(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedPublic(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestRSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedR(popr, true)
	size2 := github_com_gogo_protobuf_proto.Size(p)
	dAtA, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_gogo_protobuf_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkRSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*R, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedR(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestPublicSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedPublic(popr, true)
	size2 := github_com_gogo_protobuf_proto.Size(p)
	dAtA, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_gogo_protobuf_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkPublicSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*Public, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedPublic(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestRStringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedR(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestPublicStringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedPublic(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen
