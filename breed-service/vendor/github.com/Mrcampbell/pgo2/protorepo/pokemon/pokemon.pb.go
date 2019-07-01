// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pokemon/pokemon.proto

package pokemon

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Pokemon struct {
	// @inject_tag: sql:",pk"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:",pk"`
	// @inject_tag: sql:",notnull"
	BreedId string `protobuf:"bytes,2,opt,name=breed_id,json=breedId,proto3" json:"breed_id,omitempty" sql:",notnull"`
	// @inject_tag: sql:"-"
	BreedSummary *BreedSummary `protobuf:"bytes,3,opt,name=breed_summary,json=breedSummary,proto3" json:"breed_summary,omitempty" sql:"-"`
	Iv           *StatGroup    `protobuf:"bytes,4,opt,name=iv,proto3" json:"iv,omitempty"`
	Ev           *StatGroup    `protobuf:"bytes,5,opt,name=ev,proto3" json:"ev,omitempty"`
	Stat         *StatGroup    `protobuf:"bytes,6,opt,name=stat,proto3" json:"stat,omitempty"`
	// @inject_tag: sql:"-"
	MoveOne *MoveSummary `protobuf:"bytes,7,opt,name=move_one,json=moveOne,proto3" json:"move_one,omitempty" sql:"-"`
	// @inject_tag: sql:"-"
	MoveTwo *MoveSummary `protobuf:"bytes,8,opt,name=move_two,json=moveTwo,proto3" json:"move_two,omitempty" sql:"-"`
	// @inject_tag: sql:"-"
	MoveThree *MoveSummary `protobuf:"bytes,9,opt,name=move_three,json=moveThree,proto3" json:"move_three,omitempty" sql:"-"`
	// @inject_tag: sql:"-"
	MoveFour             *MoveSummary `protobuf:"bytes,10,opt,name=move_four,json=moveFour,proto3" json:"move_four,omitempty" sql:"-"`
	MoveOneId            string       `protobuf:"bytes,11,opt,name=move_one_id,json=moveOneId,proto3" json:"move_one_id,omitempty"`
	MoveTwoId            string       `protobuf:"bytes,12,opt,name=move_two_id,json=moveTwoId,proto3" json:"move_two_id,omitempty"`
	MoveThreeId          string       `protobuf:"bytes,13,opt,name=move_three_id,json=moveThreeId,proto3" json:"move_three_id,omitempty"`
	MoveFourId           string       `protobuf:"bytes,14,opt,name=move_four_id,json=moveFourId,proto3" json:"move_four_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Pokemon) Reset()         { *m = Pokemon{} }
func (m *Pokemon) String() string { return proto.CompactTextString(m) }
func (*Pokemon) ProtoMessage()    {}
func (*Pokemon) Descriptor() ([]byte, []int) {
	return fileDescriptor_27fc31544a7745c3, []int{0}
}

func (m *Pokemon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pokemon.Unmarshal(m, b)
}
func (m *Pokemon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pokemon.Marshal(b, m, deterministic)
}
func (m *Pokemon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pokemon.Merge(m, src)
}
func (m *Pokemon) XXX_Size() int {
	return xxx_messageInfo_Pokemon.Size(m)
}
func (m *Pokemon) XXX_DiscardUnknown() {
	xxx_messageInfo_Pokemon.DiscardUnknown(m)
}

var xxx_messageInfo_Pokemon proto.InternalMessageInfo

func (m *Pokemon) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Pokemon) GetBreedId() string {
	if m != nil {
		return m.BreedId
	}
	return ""
}

func (m *Pokemon) GetBreedSummary() *BreedSummary {
	if m != nil {
		return m.BreedSummary
	}
	return nil
}

func (m *Pokemon) GetIv() *StatGroup {
	if m != nil {
		return m.Iv
	}
	return nil
}

func (m *Pokemon) GetEv() *StatGroup {
	if m != nil {
		return m.Ev
	}
	return nil
}

func (m *Pokemon) GetStat() *StatGroup {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *Pokemon) GetMoveOne() *MoveSummary {
	if m != nil {
		return m.MoveOne
	}
	return nil
}

func (m *Pokemon) GetMoveTwo() *MoveSummary {
	if m != nil {
		return m.MoveTwo
	}
	return nil
}

func (m *Pokemon) GetMoveThree() *MoveSummary {
	if m != nil {
		return m.MoveThree
	}
	return nil
}

func (m *Pokemon) GetMoveFour() *MoveSummary {
	if m != nil {
		return m.MoveFour
	}
	return nil
}

func (m *Pokemon) GetMoveOneId() string {
	if m != nil {
		return m.MoveOneId
	}
	return ""
}

func (m *Pokemon) GetMoveTwoId() string {
	if m != nil {
		return m.MoveTwoId
	}
	return ""
}

func (m *Pokemon) GetMoveThreeId() string {
	if m != nil {
		return m.MoveThreeId
	}
	return ""
}

func (m *Pokemon) GetMoveFourId() string {
	if m != nil {
		return m.MoveFourId
	}
	return ""
}

type GetPokemonRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPokemonRequest) Reset()         { *m = GetPokemonRequest{} }
func (m *GetPokemonRequest) String() string { return proto.CompactTextString(m) }
func (*GetPokemonRequest) ProtoMessage()    {}
func (*GetPokemonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27fc31544a7745c3, []int{1}
}

func (m *GetPokemonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPokemonRequest.Unmarshal(m, b)
}
func (m *GetPokemonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPokemonRequest.Marshal(b, m, deterministic)
}
func (m *GetPokemonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPokemonRequest.Merge(m, src)
}
func (m *GetPokemonRequest) XXX_Size() int {
	return xxx_messageInfo_GetPokemonRequest.Size(m)
}
func (m *GetPokemonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPokemonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPokemonRequest proto.InternalMessageInfo

func (m *GetPokemonRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetPokemonResponse struct {
	Pokemon              *Pokemon `protobuf:"bytes,1,opt,name=pokemon,proto3" json:"pokemon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPokemonResponse) Reset()         { *m = GetPokemonResponse{} }
func (m *GetPokemonResponse) String() string { return proto.CompactTextString(m) }
func (*GetPokemonResponse) ProtoMessage()    {}
func (*GetPokemonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27fc31544a7745c3, []int{2}
}

func (m *GetPokemonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPokemonResponse.Unmarshal(m, b)
}
func (m *GetPokemonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPokemonResponse.Marshal(b, m, deterministic)
}
func (m *GetPokemonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPokemonResponse.Merge(m, src)
}
func (m *GetPokemonResponse) XXX_Size() int {
	return xxx_messageInfo_GetPokemonResponse.Size(m)
}
func (m *GetPokemonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPokemonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPokemonResponse proto.InternalMessageInfo

func (m *GetPokemonResponse) GetPokemon() *Pokemon {
	if m != nil {
		return m.Pokemon
	}
	return nil
}

type ListPokemonRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPokemonRequest) Reset()         { *m = ListPokemonRequest{} }
func (m *ListPokemonRequest) String() string { return proto.CompactTextString(m) }
func (*ListPokemonRequest) ProtoMessage()    {}
func (*ListPokemonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27fc31544a7745c3, []int{3}
}

func (m *ListPokemonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPokemonRequest.Unmarshal(m, b)
}
func (m *ListPokemonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPokemonRequest.Marshal(b, m, deterministic)
}
func (m *ListPokemonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPokemonRequest.Merge(m, src)
}
func (m *ListPokemonRequest) XXX_Size() int {
	return xxx_messageInfo_ListPokemonRequest.Size(m)
}
func (m *ListPokemonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPokemonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPokemonRequest proto.InternalMessageInfo

type ListPokemonResponse struct {
	Pokemon              []*Pokemon `protobuf:"bytes,1,rep,name=pokemon,proto3" json:"pokemon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListPokemonResponse) Reset()         { *m = ListPokemonResponse{} }
func (m *ListPokemonResponse) String() string { return proto.CompactTextString(m) }
func (*ListPokemonResponse) ProtoMessage()    {}
func (*ListPokemonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27fc31544a7745c3, []int{4}
}

func (m *ListPokemonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPokemonResponse.Unmarshal(m, b)
}
func (m *ListPokemonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPokemonResponse.Marshal(b, m, deterministic)
}
func (m *ListPokemonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPokemonResponse.Merge(m, src)
}
func (m *ListPokemonResponse) XXX_Size() int {
	return xxx_messageInfo_ListPokemonResponse.Size(m)
}
func (m *ListPokemonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPokemonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPokemonResponse proto.InternalMessageInfo

func (m *ListPokemonResponse) GetPokemon() []*Pokemon {
	if m != nil {
		return m.Pokemon
	}
	return nil
}

type InternalCreatePokemonRequest struct {
	BreedId              string       `protobuf:"bytes,1,opt,name=breed_id,json=breedId,proto3" json:"breed_id,omitempty"`
	Level                int32        `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	OverrideGenderIsMale bool         `protobuf:"varint,3,opt,name=override_gender_is_male,json=overrideGenderIsMale,proto3" json:"override_gender_is_male,omitempty"`
	OverrideStatGroup    *StatGroup   `protobuf:"bytes,4,opt,name=override_stat_group,json=overrideStatGroup,proto3" json:"override_stat_group,omitempty"`
	VersionGroupId       VersionGroup `protobuf:"varint,5,opt,name=version_group_id,json=versionGroupId,proto3,enum=pokemon.VersionGroup" json:"version_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *InternalCreatePokemonRequest) Reset()         { *m = InternalCreatePokemonRequest{} }
func (m *InternalCreatePokemonRequest) String() string { return proto.CompactTextString(m) }
func (*InternalCreatePokemonRequest) ProtoMessage()    {}
func (*InternalCreatePokemonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27fc31544a7745c3, []int{5}
}

func (m *InternalCreatePokemonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalCreatePokemonRequest.Unmarshal(m, b)
}
func (m *InternalCreatePokemonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalCreatePokemonRequest.Marshal(b, m, deterministic)
}
func (m *InternalCreatePokemonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalCreatePokemonRequest.Merge(m, src)
}
func (m *InternalCreatePokemonRequest) XXX_Size() int {
	return xxx_messageInfo_InternalCreatePokemonRequest.Size(m)
}
func (m *InternalCreatePokemonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalCreatePokemonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InternalCreatePokemonRequest proto.InternalMessageInfo

func (m *InternalCreatePokemonRequest) GetBreedId() string {
	if m != nil {
		return m.BreedId
	}
	return ""
}

func (m *InternalCreatePokemonRequest) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *InternalCreatePokemonRequest) GetOverrideGenderIsMale() bool {
	if m != nil {
		return m.OverrideGenderIsMale
	}
	return false
}

func (m *InternalCreatePokemonRequest) GetOverrideStatGroup() *StatGroup {
	if m != nil {
		return m.OverrideStatGroup
	}
	return nil
}

func (m *InternalCreatePokemonRequest) GetVersionGroupId() VersionGroup {
	if m != nil {
		return m.VersionGroupId
	}
	return VersionGroup_UNKNOWN_VERSION_GROUP
}

type InternalCreatePokemonResponse struct {
	Pokemon              *Pokemon `protobuf:"bytes,1,opt,name=pokemon,proto3" json:"pokemon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalCreatePokemonResponse) Reset()         { *m = InternalCreatePokemonResponse{} }
func (m *InternalCreatePokemonResponse) String() string { return proto.CompactTextString(m) }
func (*InternalCreatePokemonResponse) ProtoMessage()    {}
func (*InternalCreatePokemonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27fc31544a7745c3, []int{6}
}

func (m *InternalCreatePokemonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalCreatePokemonResponse.Unmarshal(m, b)
}
func (m *InternalCreatePokemonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalCreatePokemonResponse.Marshal(b, m, deterministic)
}
func (m *InternalCreatePokemonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalCreatePokemonResponse.Merge(m, src)
}
func (m *InternalCreatePokemonResponse) XXX_Size() int {
	return xxx_messageInfo_InternalCreatePokemonResponse.Size(m)
}
func (m *InternalCreatePokemonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalCreatePokemonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InternalCreatePokemonResponse proto.InternalMessageInfo

func (m *InternalCreatePokemonResponse) GetPokemon() *Pokemon {
	if m != nil {
		return m.Pokemon
	}
	return nil
}

func init() {
	proto.RegisterType((*Pokemon)(nil), "pokemon.Pokemon")
	proto.RegisterType((*GetPokemonRequest)(nil), "pokemon.GetPokemonRequest")
	proto.RegisterType((*GetPokemonResponse)(nil), "pokemon.GetPokemonResponse")
	proto.RegisterType((*ListPokemonRequest)(nil), "pokemon.ListPokemonRequest")
	proto.RegisterType((*ListPokemonResponse)(nil), "pokemon.ListPokemonResponse")
	proto.RegisterType((*InternalCreatePokemonRequest)(nil), "pokemon.InternalCreatePokemonRequest")
	proto.RegisterType((*InternalCreatePokemonResponse)(nil), "pokemon.InternalCreatePokemonResponse")
}

func init() { proto.RegisterFile("pokemon/pokemon.proto", fileDescriptor_27fc31544a7745c3) }

var fileDescriptor_27fc31544a7745c3 = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x55, 0xdc, 0x8f, 0x24, 0xd7, 0x4d, 0xd4, 0xde, 0xa4, 0xaa, 0xeb, 0xd7, 0xf7, 0x5e, 0xe5,
	0xa7, 0x57, 0x55, 0x5d, 0xd4, 0x6a, 0x2b, 0x36, 0xdd, 0x00, 0x45, 0xa2, 0xb2, 0xa0, 0x02, 0xa5,
	0x15, 0x0b, 0x24, 0x64, 0xb9, 0xf8, 0x12, 0x46, 0x24, 0x9e, 0x30, 0x9e, 0x38, 0x42, 0x88, 0x0d,
	0x1b, 0x16, 0x2c, 0xf9, 0x23, 0xfc, 0x17, 0x76, 0xac, 0xf9, 0x21, 0x68, 0xc6, 0x1e, 0xc7, 0x6d,
	0xd3, 0x54, 0xac, 0xa2, 0x7b, 0xcf, 0x39, 0x73, 0xcf, 0x8c, 0xcf, 0x0d, 0xac, 0x8f, 0xf8, 0x3b,
	0x1a, 0xf2, 0xc4, 0x2f, 0x7e, 0xf7, 0x47, 0x82, 0x4b, 0x8e, 0xf5, 0xa2, 0x74, 0xb7, 0xfa, 0x9c,
	0xf7, 0x07, 0xe4, 0x47, 0x23, 0xe6, 0x47, 0x49, 0xc2, 0x65, 0x24, 0x19, 0x4f, 0xd2, 0x9c, 0xe6,
	0x76, 0x8c, 0xfa, 0x52, 0x10, 0xc5, 0x45, 0x73, 0xc3, 0x34, 0x53, 0x19, 0xc9, 0xbe, 0xe0, 0xe3,
	0x51, 0x01, 0xa0, 0x01, 0x86, 0x3c, 0xa3, 0xa2, 0xe7, 0x9a, 0x5e, 0x46, 0x22, 0x65, 0x3c, 0xa9,
	0xf0, 0xbd, 0xef, 0x8b, 0x50, 0x7f, 0x9e, 0xc3, 0xd8, 0x06, 0x8b, 0xc5, 0x4e, 0x6d, 0xbb, 0xb6,
	0xdb, 0xec, 0x59, 0x2c, 0xc6, 0x4d, 0x68, 0xe8, 0x99, 0x21, 0x8b, 0x1d, 0x4b, 0x77, 0xeb, 0xba,
	0x0e, 0x62, 0x3c, 0x86, 0x56, 0x0e, 0xa5, 0xe3, 0xe1, 0x30, 0x12, 0x1f, 0x9c, 0x85, 0xed, 0xda,
	0xae, 0x7d, 0xb8, 0xbe, 0x6f, 0xae, 0x78, 0xa2, 0xd0, 0xf3, 0x1c, 0xec, 0xad, 0x5c, 0x56, 0x2a,
	0xf4, 0xc0, 0x62, 0x99, 0xb3, 0xa8, 0x05, 0x58, 0x0a, 0xce, 0x65, 0x24, 0x4f, 0x95, 0xb1, 0x9e,
	0xc5, 0x32, 0xc5, 0xa1, 0xcc, 0x59, 0xba, 0x9d, 0x43, 0x19, 0xee, 0xc0, 0xa2, 0xba, 0xbd, 0xb3,
	0x7c, 0x2b, 0x4b, 0xe3, 0xe8, 0x43, 0x43, 0x3d, 0x46, 0xc8, 0x13, 0x72, 0xea, 0x9a, 0xdb, 0x2d,
	0xb9, 0x67, 0x3c, 0x23, 0xe3, 0xb2, 0xae, 0x58, 0xcf, 0x12, 0x2a, 0x05, 0x72, 0xc2, 0x9d, 0xc6,
	0x5d, 0x82, 0x8b, 0x09, 0xc7, 0x23, 0x80, 0x5c, 0xf0, 0x56, 0x10, 0x39, 0xcd, 0x39, 0x92, 0xa6,
	0x96, 0x28, 0x1a, 0x1e, 0x80, 0x2e, 0xc2, 0x37, 0x7c, 0x2c, 0x1c, 0x98, 0xa3, 0xd1, 0x66, 0x1e,
	0xf3, 0xb1, 0xc0, 0x7f, 0xc0, 0x36, 0x37, 0x51, 0xdf, 0xc4, 0xd6, 0xdf, 0xa4, 0x59, 0xd8, 0x0e,
	0xe2, 0x12, 0x97, 0x13, 0xae, 0xf0, 0x95, 0x29, 0x7e, 0x31, 0xe1, 0x41, 0x8c, 0x1e, 0xb4, 0xa6,
	0x3e, 0x15, 0xa3, 0xa5, 0x19, 0x76, 0x69, 0x2a, 0x88, 0x71, 0x1b, 0x56, 0x4a, 0x5b, 0x8a, 0xd2,
	0xd6, 0x14, 0x30, 0x1e, 0x82, 0xd8, 0xfb, 0x0f, 0xd6, 0x4e, 0x49, 0x16, 0xa1, 0xe9, 0xd1, 0xfb,
	0x31, 0xa5, 0xf2, 0x7a, 0x76, 0xbc, 0x07, 0x80, 0x55, 0x52, 0x3a, 0xe2, 0x49, 0x4a, 0xb8, 0x07,
	0x26, 0xf4, 0x9a, 0x6a, 0x1f, 0xae, 0x96, 0x37, 0x36, 0x54, 0x43, 0xf0, 0xba, 0x80, 0x4f, 0x59,
	0x7a, 0x6d, 0x8e, 0xf7, 0x10, 0x3a, 0x57, 0xba, 0xb3, 0x0e, 0x5e, 0x98, 0x7f, 0xf0, 0x57, 0x0b,
	0xb6, 0x82, 0x44, 0x92, 0x48, 0xa2, 0xc1, 0x23, 0x41, 0x91, 0xa4, 0x6b, 0x77, 0xa9, 0xe6, 0xbe,
	0x76, 0x35, 0xf7, 0x5d, 0x58, 0x1a, 0x50, 0x46, 0x03, 0xbd, 0x0f, 0x4b, 0xbd, 0xbc, 0xc0, 0x7b,
	0xb0, 0xc1, 0x33, 0x12, 0x82, 0xc5, 0x14, 0xf6, 0x29, 0x89, 0x49, 0x84, 0x2c, 0x0d, 0x87, 0xd1,
	0x80, 0xf4, 0x5e, 0x34, 0x7a, 0x5d, 0x03, 0x9f, 0x6a, 0x34, 0x48, 0xcf, 0xa2, 0x01, 0xe1, 0x09,
	0x74, 0x4a, 0x99, 0x4a, 0x6a, 0xa8, 0x17, 0x73, 0xce, 0x66, 0xac, 0x19, 0x7a, 0xd9, 0xc2, 0xfb,
	0xb0, 0x5a, 0x6c, 0x75, 0xae, 0x56, 0x9e, 0xd5, 0xda, 0xb4, 0x2b, 0xbb, 0xf8, 0x22, 0x27, 0xe4,
	0x67, 0xb4, 0xb3, 0x4a, 0x15, 0xc4, 0xde, 0x13, 0xf8, 0xfb, 0x96, 0xc7, 0xf8, 0xf3, 0x6f, 0x76,
	0xf8, 0xd3, 0x82, 0x76, 0xd1, 0x3c, 0x27, 0x91, 0xb1, 0xd7, 0x84, 0x21, 0xc0, 0x34, 0x08, 0xe8,
	0x96, 0xda, 0x1b, 0x11, 0x72, 0xff, 0x9a, 0x89, 0xe5, 0x2e, 0x3c, 0xe7, 0xf3, 0x8f, 0x5f, 0xdf,
	0x2c, 0xc4, 0x55, 0x3f, 0x3b, 0x30, 0xff, 0xa3, 0xfe, 0x47, 0x16, 0x7f, 0xc2, 0x57, 0x60, 0x57,
	0x12, 0x81, 0xd3, 0x53, 0x6e, 0xa6, 0xc7, 0xdd, 0x9a, 0x0d, 0x16, 0x33, 0x3a, 0x7a, 0x46, 0x0b,
	0xed, 0xca, 0x0c, 0xfc, 0x52, 0x83, 0xf5, 0x99, 0x0f, 0x84, 0xff, 0x97, 0x87, 0xcd, 0x4b, 0x93,
	0xbb, 0x73, 0x17, 0xad, 0x98, 0xfe, 0xaf, 0x9e, 0xbe, 0xe9, 0x75, 0x7d, 0x56, 0xf0, 0x2a, 0x36,
	0x8e, 0x6b, 0x7b, 0x27, 0xcd, 0x97, 0xe6, 0x9d, 0x2f, 0x97, 0xf5, 0x9f, 0xf7, 0xd1, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xdd, 0x8e, 0xa6, 0x15, 0x5a, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PokemonServiceClient is the client API for PokemonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PokemonServiceClient interface {
	GetPokemon(ctx context.Context, in *GetPokemonRequest, opts ...grpc.CallOption) (*GetPokemonResponse, error)
	ListPokemon(ctx context.Context, in *ListPokemonRequest, opts ...grpc.CallOption) (*ListPokemonResponse, error)
	InternalCreatePokemon(ctx context.Context, in *InternalCreatePokemonRequest, opts ...grpc.CallOption) (*InternalCreatePokemonResponse, error)
}

type pokemonServiceClient struct {
	cc *grpc.ClientConn
}

func NewPokemonServiceClient(cc *grpc.ClientConn) PokemonServiceClient {
	return &pokemonServiceClient{cc}
}

func (c *pokemonServiceClient) GetPokemon(ctx context.Context, in *GetPokemonRequest, opts ...grpc.CallOption) (*GetPokemonResponse, error) {
	out := new(GetPokemonResponse)
	err := c.cc.Invoke(ctx, "/pokemon.PokemonService/GetPokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) ListPokemon(ctx context.Context, in *ListPokemonRequest, opts ...grpc.CallOption) (*ListPokemonResponse, error) {
	out := new(ListPokemonResponse)
	err := c.cc.Invoke(ctx, "/pokemon.PokemonService/ListPokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) InternalCreatePokemon(ctx context.Context, in *InternalCreatePokemonRequest, opts ...grpc.CallOption) (*InternalCreatePokemonResponse, error) {
	out := new(InternalCreatePokemonResponse)
	err := c.cc.Invoke(ctx, "/pokemon.PokemonService/InternalCreatePokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokemonServiceServer is the server API for PokemonService service.
type PokemonServiceServer interface {
	GetPokemon(context.Context, *GetPokemonRequest) (*GetPokemonResponse, error)
	ListPokemon(context.Context, *ListPokemonRequest) (*ListPokemonResponse, error)
	InternalCreatePokemon(context.Context, *InternalCreatePokemonRequest) (*InternalCreatePokemonResponse, error)
}

// UnimplementedPokemonServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPokemonServiceServer struct {
}

func (*UnimplementedPokemonServiceServer) GetPokemon(ctx context.Context, req *GetPokemonRequest) (*GetPokemonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemon not implemented")
}
func (*UnimplementedPokemonServiceServer) ListPokemon(ctx context.Context, req *ListPokemonRequest) (*ListPokemonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPokemon not implemented")
}
func (*UnimplementedPokemonServiceServer) InternalCreatePokemon(ctx context.Context, req *InternalCreatePokemonRequest) (*InternalCreatePokemonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InternalCreatePokemon not implemented")
}

func RegisterPokemonServiceServer(s *grpc.Server, srv PokemonServiceServer) {
	s.RegisterService(&_PokemonService_serviceDesc, srv)
}

func _PokemonService_GetPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).GetPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.PokemonService/GetPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).GetPokemon(ctx, req.(*GetPokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_ListPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).ListPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.PokemonService/ListPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).ListPokemon(ctx, req.(*ListPokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_InternalCreatePokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalCreatePokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).InternalCreatePokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.PokemonService/InternalCreatePokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).InternalCreatePokemon(ctx, req.(*InternalCreatePokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PokemonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.PokemonService",
	HandlerType: (*PokemonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPokemon",
			Handler:    _PokemonService_GetPokemon_Handler,
		},
		{
			MethodName: "ListPokemon",
			Handler:    _PokemonService_ListPokemon_Handler,
		},
		{
			MethodName: "InternalCreatePokemon",
			Handler:    _PokemonService_InternalCreatePokemon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokemon/pokemon.proto",
}
