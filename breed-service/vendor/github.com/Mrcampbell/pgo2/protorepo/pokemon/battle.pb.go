// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pokemon/battle.proto

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

type BattleState int32

const (
	BattleState_BATTLE_STATE_UNKNOWN    BattleState = 0
	BattleState_WAITING_ON_BOTH_PLAYERS BattleState = 1
	BattleState_WAITING_ON_PLAYER_A     BattleState = 2
	BattleState_WAITING_ON_PLAYER_B     BattleState = 3
	BattleState_MOVE_RESULTS_QUEUED     BattleState = 4
	BattleState_BATTLE_OVER             BattleState = 5
)

var BattleState_name = map[int32]string{
	0: "BATTLE_STATE_UNKNOWN",
	1: "WAITING_ON_BOTH_PLAYERS",
	2: "WAITING_ON_PLAYER_A",
	3: "WAITING_ON_PLAYER_B",
	4: "MOVE_RESULTS_QUEUED",
	5: "BATTLE_OVER",
}

var BattleState_value = map[string]int32{
	"BATTLE_STATE_UNKNOWN":    0,
	"WAITING_ON_BOTH_PLAYERS": 1,
	"WAITING_ON_PLAYER_A":     2,
	"WAITING_ON_PLAYER_B":     3,
	"MOVE_RESULTS_QUEUED":     4,
	"BATTLE_OVER":             5,
}

func (x BattleState) String() string {
	return proto.EnumName(BattleState_name, int32(x))
}

func (BattleState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_30c8eb8dca680047, []int{0}
}

type Battle struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PokemonABattleMask   *PokemonBattleMask `protobuf:"bytes,2,opt,name=pokemon_a_battle_mask,json=pokemonABattleMask,proto3" json:"pokemon_a_battle_mask,omitempty"`
	PokemonBBattleMask   *PokemonBattleMask `protobuf:"bytes,3,opt,name=pokemon_b_battle_mask,json=pokemonBBattleMask,proto3" json:"pokemon_b_battle_mask,omitempty"`
	PlayerAId            string             `protobuf:"bytes,4,opt,name=player_a_id,json=playerAId,proto3" json:"player_a_id,omitempty"`
	PlayerBId            string             `protobuf:"bytes,5,opt,name=player_b_id,json=playerBId,proto3" json:"player_b_id,omitempty"`
	Duration             int32              `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	State                BattleState        `protobuf:"varint,7,opt,name=state,proto3,enum=pokemon.BattleState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Battle) Reset()         { *m = Battle{} }
func (m *Battle) String() string { return proto.CompactTextString(m) }
func (*Battle) ProtoMessage()    {}
func (*Battle) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c8eb8dca680047, []int{0}
}

func (m *Battle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Battle.Unmarshal(m, b)
}
func (m *Battle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Battle.Marshal(b, m, deterministic)
}
func (m *Battle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Battle.Merge(m, src)
}
func (m *Battle) XXX_Size() int {
	return xxx_messageInfo_Battle.Size(m)
}
func (m *Battle) XXX_DiscardUnknown() {
	xxx_messageInfo_Battle.DiscardUnknown(m)
}

var xxx_messageInfo_Battle proto.InternalMessageInfo

func (m *Battle) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Battle) GetPokemonABattleMask() *PokemonBattleMask {
	if m != nil {
		return m.PokemonABattleMask
	}
	return nil
}

func (m *Battle) GetPokemonBBattleMask() *PokemonBattleMask {
	if m != nil {
		return m.PokemonBBattleMask
	}
	return nil
}

func (m *Battle) GetPlayerAId() string {
	if m != nil {
		return m.PlayerAId
	}
	return ""
}

func (m *Battle) GetPlayerBId() string {
	if m != nil {
		return m.PlayerBId
	}
	return ""
}

func (m *Battle) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Battle) GetState() BattleState {
	if m != nil {
		return m.State
	}
	return BattleState_BATTLE_STATE_UNKNOWN
}

type CreateBattleRequest struct {
	PokemonAId           string   `protobuf:"bytes,1,opt,name=pokemon_a_id,json=pokemonAId,proto3" json:"pokemon_a_id,omitempty"`
	PokemonBId           string   `protobuf:"bytes,2,opt,name=pokemon_b_id,json=pokemonBId,proto3" json:"pokemon_b_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBattleRequest) Reset()         { *m = CreateBattleRequest{} }
func (m *CreateBattleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBattleRequest) ProtoMessage()    {}
func (*CreateBattleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c8eb8dca680047, []int{1}
}

func (m *CreateBattleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBattleRequest.Unmarshal(m, b)
}
func (m *CreateBattleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBattleRequest.Marshal(b, m, deterministic)
}
func (m *CreateBattleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBattleRequest.Merge(m, src)
}
func (m *CreateBattleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBattleRequest.Size(m)
}
func (m *CreateBattleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBattleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBattleRequest proto.InternalMessageInfo

func (m *CreateBattleRequest) GetPokemonAId() string {
	if m != nil {
		return m.PokemonAId
	}
	return ""
}

func (m *CreateBattleRequest) GetPokemonBId() string {
	if m != nil {
		return m.PokemonBId
	}
	return ""
}

type CreateBattleResponse struct {
	Battle               *Battle  `protobuf:"bytes,1,opt,name=battle,proto3" json:"battle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBattleResponse) Reset()         { *m = CreateBattleResponse{} }
func (m *CreateBattleResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBattleResponse) ProtoMessage()    {}
func (*CreateBattleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c8eb8dca680047, []int{2}
}

func (m *CreateBattleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBattleResponse.Unmarshal(m, b)
}
func (m *CreateBattleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBattleResponse.Marshal(b, m, deterministic)
}
func (m *CreateBattleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBattleResponse.Merge(m, src)
}
func (m *CreateBattleResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBattleResponse.Size(m)
}
func (m *CreateBattleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBattleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBattleResponse proto.InternalMessageInfo

func (m *CreateBattleResponse) GetBattle() *Battle {
	if m != nil {
		return m.Battle
	}
	return nil
}

type GetBattleRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBattleRequest) Reset()         { *m = GetBattleRequest{} }
func (m *GetBattleRequest) String() string { return proto.CompactTextString(m) }
func (*GetBattleRequest) ProtoMessage()    {}
func (*GetBattleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c8eb8dca680047, []int{3}
}

func (m *GetBattleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBattleRequest.Unmarshal(m, b)
}
func (m *GetBattleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBattleRequest.Marshal(b, m, deterministic)
}
func (m *GetBattleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBattleRequest.Merge(m, src)
}
func (m *GetBattleRequest) XXX_Size() int {
	return xxx_messageInfo_GetBattleRequest.Size(m)
}
func (m *GetBattleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBattleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBattleRequest proto.InternalMessageInfo

func (m *GetBattleRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetBattleResponse struct {
	Battle               *Battle  `protobuf:"bytes,1,opt,name=battle,proto3" json:"battle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBattleResponse) Reset()         { *m = GetBattleResponse{} }
func (m *GetBattleResponse) String() string { return proto.CompactTextString(m) }
func (*GetBattleResponse) ProtoMessage()    {}
func (*GetBattleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c8eb8dca680047, []int{4}
}

func (m *GetBattleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBattleResponse.Unmarshal(m, b)
}
func (m *GetBattleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBattleResponse.Marshal(b, m, deterministic)
}
func (m *GetBattleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBattleResponse.Merge(m, src)
}
func (m *GetBattleResponse) XXX_Size() int {
	return xxx_messageInfo_GetBattleResponse.Size(m)
}
func (m *GetBattleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBattleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBattleResponse proto.InternalMessageInfo

func (m *GetBattleResponse) GetBattle() *Battle {
	if m != nil {
		return m.Battle
	}
	return nil
}

type SubmitTurnRequest struct {
	BattleId             string      `protobuf:"bytes,1,opt,name=battle_id,json=battleId,proto3" json:"battle_id,omitempty"`
	PlayerId             string      `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Turn                 *PlayerTurn `protobuf:"bytes,3,opt,name=turn,proto3" json:"turn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SubmitTurnRequest) Reset()         { *m = SubmitTurnRequest{} }
func (m *SubmitTurnRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitTurnRequest) ProtoMessage()    {}
func (*SubmitTurnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c8eb8dca680047, []int{5}
}

func (m *SubmitTurnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTurnRequest.Unmarshal(m, b)
}
func (m *SubmitTurnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTurnRequest.Marshal(b, m, deterministic)
}
func (m *SubmitTurnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTurnRequest.Merge(m, src)
}
func (m *SubmitTurnRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitTurnRequest.Size(m)
}
func (m *SubmitTurnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTurnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTurnRequest proto.InternalMessageInfo

func (m *SubmitTurnRequest) GetBattleId() string {
	if m != nil {
		return m.BattleId
	}
	return ""
}

func (m *SubmitTurnRequest) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *SubmitTurnRequest) GetTurn() *PlayerTurn {
	if m != nil {
		return m.Turn
	}
	return nil
}

type SubmitTurnResponse struct {
	BattleId             string      `protobuf:"bytes,1,opt,name=battle_id,json=battleId,proto3" json:"battle_id,omitempty"`
	PlayerId             string      `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Turn                 *PlayerTurn `protobuf:"bytes,3,opt,name=turn,proto3" json:"turn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SubmitTurnResponse) Reset()         { *m = SubmitTurnResponse{} }
func (m *SubmitTurnResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitTurnResponse) ProtoMessage()    {}
func (*SubmitTurnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c8eb8dca680047, []int{6}
}

func (m *SubmitTurnResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTurnResponse.Unmarshal(m, b)
}
func (m *SubmitTurnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTurnResponse.Marshal(b, m, deterministic)
}
func (m *SubmitTurnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTurnResponse.Merge(m, src)
}
func (m *SubmitTurnResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitTurnResponse.Size(m)
}
func (m *SubmitTurnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTurnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTurnResponse proto.InternalMessageInfo

func (m *SubmitTurnResponse) GetBattleId() string {
	if m != nil {
		return m.BattleId
	}
	return ""
}

func (m *SubmitTurnResponse) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *SubmitTurnResponse) GetTurn() *PlayerTurn {
	if m != nil {
		return m.Turn
	}
	return nil
}

type BattleRound struct {
	PlayerATurn          *PlayerTurn `protobuf:"bytes,1,opt,name=player_a_turn,json=playerATurn,proto3" json:"player_a_turn,omitempty"`
	PlayerBTurn          *PlayerTurn `protobuf:"bytes,2,opt,name=player_b_turn,json=playerBTurn,proto3" json:"player_b_turn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BattleRound) Reset()         { *m = BattleRound{} }
func (m *BattleRound) String() string { return proto.CompactTextString(m) }
func (*BattleRound) ProtoMessage()    {}
func (*BattleRound) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c8eb8dca680047, []int{7}
}

func (m *BattleRound) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleRound.Unmarshal(m, b)
}
func (m *BattleRound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleRound.Marshal(b, m, deterministic)
}
func (m *BattleRound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleRound.Merge(m, src)
}
func (m *BattleRound) XXX_Size() int {
	return xxx_messageInfo_BattleRound.Size(m)
}
func (m *BattleRound) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleRound.DiscardUnknown(m)
}

var xxx_messageInfo_BattleRound proto.InternalMessageInfo

func (m *BattleRound) GetPlayerATurn() *PlayerTurn {
	if m != nil {
		return m.PlayerATurn
	}
	return nil
}

func (m *BattleRound) GetPlayerBTurn() *PlayerTurn {
	if m != nil {
		return m.PlayerBTurn
	}
	return nil
}

type PlayerTurn struct {
	IsAttack             bool          `protobuf:"varint,1,opt,name=is_attack,json=isAttack,proto3" json:"is_attack,omitempty"`
	Attack               *PlayerAttack `protobuf:"bytes,2,opt,name=attack,proto3" json:"attack,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PlayerTurn) Reset()         { *m = PlayerTurn{} }
func (m *PlayerTurn) String() string { return proto.CompactTextString(m) }
func (*PlayerTurn) ProtoMessage()    {}
func (*PlayerTurn) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c8eb8dca680047, []int{8}
}

func (m *PlayerTurn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerTurn.Unmarshal(m, b)
}
func (m *PlayerTurn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerTurn.Marshal(b, m, deterministic)
}
func (m *PlayerTurn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerTurn.Merge(m, src)
}
func (m *PlayerTurn) XXX_Size() int {
	return xxx_messageInfo_PlayerTurn.Size(m)
}
func (m *PlayerTurn) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerTurn.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerTurn proto.InternalMessageInfo

func (m *PlayerTurn) GetIsAttack() bool {
	if m != nil {
		return m.IsAttack
	}
	return false
}

func (m *PlayerTurn) GetAttack() *PlayerAttack {
	if m != nil {
		return m.Attack
	}
	return nil
}

type PlayerAttack struct {
	PokemonMoveSlot      int32    `protobuf:"varint,1,opt,name=pokemon_move_slot,json=pokemonMoveSlot,proto3" json:"pokemon_move_slot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerAttack) Reset()         { *m = PlayerAttack{} }
func (m *PlayerAttack) String() string { return proto.CompactTextString(m) }
func (*PlayerAttack) ProtoMessage()    {}
func (*PlayerAttack) Descriptor() ([]byte, []int) {
	return fileDescriptor_30c8eb8dca680047, []int{9}
}

func (m *PlayerAttack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerAttack.Unmarshal(m, b)
}
func (m *PlayerAttack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerAttack.Marshal(b, m, deterministic)
}
func (m *PlayerAttack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerAttack.Merge(m, src)
}
func (m *PlayerAttack) XXX_Size() int {
	return xxx_messageInfo_PlayerAttack.Size(m)
}
func (m *PlayerAttack) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerAttack.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerAttack proto.InternalMessageInfo

func (m *PlayerAttack) GetPokemonMoveSlot() int32 {
	if m != nil {
		return m.PokemonMoveSlot
	}
	return 0
}

func init() {
	proto.RegisterEnum("pokemon.BattleState", BattleState_name, BattleState_value)
	proto.RegisterType((*Battle)(nil), "pokemon.Battle")
	proto.RegisterType((*CreateBattleRequest)(nil), "pokemon.CreateBattleRequest")
	proto.RegisterType((*CreateBattleResponse)(nil), "pokemon.CreateBattleResponse")
	proto.RegisterType((*GetBattleRequest)(nil), "pokemon.GetBattleRequest")
	proto.RegisterType((*GetBattleResponse)(nil), "pokemon.GetBattleResponse")
	proto.RegisterType((*SubmitTurnRequest)(nil), "pokemon.SubmitTurnRequest")
	proto.RegisterType((*SubmitTurnResponse)(nil), "pokemon.SubmitTurnResponse")
	proto.RegisterType((*BattleRound)(nil), "pokemon.BattleRound")
	proto.RegisterType((*PlayerTurn)(nil), "pokemon.PlayerTurn")
	proto.RegisterType((*PlayerAttack)(nil), "pokemon.PlayerAttack")
}

func init() { proto.RegisterFile("pokemon/battle.proto", fileDescriptor_30c8eb8dca680047) }

var fileDescriptor_30c8eb8dca680047 = []byte{
	// 733 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0x61, 0x6f, 0xd2, 0x5c,
	0x14, 0xc7, 0x9f, 0x76, 0x83, 0xc1, 0x61, 0x1b, 0x70, 0xc7, 0x32, 0x9e, 0x6e, 0x2a, 0xe9, 0x9b,
	0x11, 0x12, 0xd7, 0x88, 0x31, 0x26, 0x8b, 0x89, 0x69, 0xb5, 0x99, 0x8d, 0x03, 0x66, 0x29, 0x9b,
	0x53, 0x93, 0xa6, 0xac, 0x37, 0x4b, 0x03, 0xb4, 0x48, 0x2f, 0x18, 0x33, 0x17, 0x13, 0xbf, 0x82,
	0x5f, 0xc0, 0x4f, 0xe2, 0x6b, 0xdf, 0xfb, 0x15, 0xfc, 0x20, 0xa6, 0xf7, 0x5e, 0x4a, 0x41, 0x16,
	0xf5, 0x8d, 0xaf, 0x9a, 0x7b, 0xfe, 0xe7, 0xfc, 0xce, 0xe9, 0xb9, 0xe7, 0xb4, 0x50, 0x1a, 0x06,
	0x3d, 0x3c, 0x08, 0x7c, 0xa5, 0xeb, 0x10, 0xd2, 0xc7, 0x07, 0xc3, 0x51, 0x40, 0x02, 0xb4, 0xc6,
	0xad, 0xd2, 0xde, 0x65, 0x10, 0x5c, 0xf6, 0xb1, 0xe2, 0x0c, 0x3d, 0xc5, 0xf1, 0xfd, 0x80, 0x38,
	0xc4, 0x0b, 0xfc, 0x90, 0xb9, 0x49, 0x77, 0xa6, 0xc1, 0xfc, 0xc9, 0x18, 0x03, 0x27, 0xec, 0x31,
	0x07, 0xf9, 0xab, 0x08, 0x69, 0x8d, 0x1a, 0xd1, 0x26, 0x88, 0x9e, 0x5b, 0x16, 0x2a, 0x42, 0x35,
	0x6b, 0x8a, 0x9e, 0x8b, 0x1a, 0xb0, 0xcd, 0xa3, 0x6c, 0xc7, 0x66, 0x81, 0x76, 0x14, 0x59, 0x16,
	0x2b, 0x42, 0x35, 0x57, 0x97, 0x0e, 0xb8, 0x7a, 0x70, 0xc2, 0x9e, 0x0c, 0xd3, 0x70, 0xc2, 0x9e,
	0x89, 0xb8, 0xa4, 0xce, 0x6c, 0x49, 0x5c, 0x77, 0x0e, 0xb7, 0xf2, 0xc7, 0x38, 0x2d, 0x81, 0xbb,
	0x0d, 0xb9, 0x61, 0xdf, 0x79, 0x8f, 0x47, 0xb6, 0x63, 0x7b, 0x6e, 0x79, 0x95, 0x96, 0x9d, 0x65,
	0x26, 0xd5, 0x70, 0x13, 0x7a, 0x37, 0xd2, 0x53, 0x49, 0x5d, 0x33, 0x5c, 0x24, 0x41, 0xc6, 0x1d,
	0x8f, 0x68, 0xb3, 0xca, 0xe9, 0x8a, 0x50, 0x4d, 0x99, 0xf1, 0x19, 0xd5, 0x20, 0x15, 0x12, 0x87,
	0xe0, 0xf2, 0x5a, 0x45, 0xa8, 0x6e, 0xd6, 0x4b, 0x71, 0x69, 0x2c, 0x7f, 0x3b, 0xd2, 0x4c, 0xe6,
	0x22, 0x9f, 0xc3, 0xd6, 0x93, 0x11, 0x76, 0x08, 0x66, 0x9a, 0x89, 0xdf, 0x8e, 0x71, 0x48, 0x50,
	0x05, 0xd6, 0x67, 0xcd, 0x8b, 0xdb, 0x0a, 0xd3, 0xbe, 0x18, 0x6e, 0xd2, 0x83, 0x56, 0x28, 0xce,
	0x79, 0x68, 0x86, 0x2b, 0x3f, 0x86, 0xd2, 0x3c, 0x3a, 0x1c, 0x06, 0x7e, 0x88, 0xd1, 0x3e, 0xa4,
	0x59, 0xff, 0x28, 0x35, 0x57, 0xcf, 0x2f, 0xd4, 0x67, 0x72, 0x59, 0x96, 0xa1, 0x70, 0x84, 0xc9,
	0x7c, 0x61, 0x0b, 0xb7, 0x2c, 0x3f, 0x82, 0x62, 0xc2, 0xe7, 0x6f, 0x33, 0x4c, 0xa0, 0xd8, 0x1e,
	0x77, 0x07, 0x1e, 0xb1, 0xc6, 0x23, 0x7f, 0x9a, 0x62, 0x17, 0xb2, 0xfc, 0x7e, 0xe3, 0x4c, 0x19,
	0x66, 0x30, 0xdc, 0x48, 0xe4, 0xf7, 0x12, 0xbf, 0x73, 0x86, 0x19, 0x0c, 0x17, 0xed, 0xc3, 0x2a,
	0x19, 0x8f, 0x7c, 0x3e, 0x12, 0x5b, 0xb3, 0x91, 0xa0, 0x0e, 0x34, 0x07, 0x75, 0x90, 0xdf, 0x01,
	0x4a, 0xe6, 0xe5, 0x65, 0xff, 0x83, 0xc4, 0x1f, 0x21, 0xc7, 0x5b, 0x10, 0x8c, 0x7d, 0x17, 0x3d,
	0x84, 0x8d, 0x78, 0x0a, 0x29, 0x40, 0xb8, 0x19, 0xc0, 0xe7, 0x51, 0x8d, 0x0e, 0x89, 0xc0, 0x2e,
	0x0b, 0x14, 0x7f, 0x1b, 0xa8, 0x45, 0x07, 0xf9, 0x25, 0xc0, 0x4c, 0x8a, 0x5e, 0xca, 0x0b, 0x6d,
	0x87, 0x10, 0xe7, 0xa2, 0x47, 0x73, 0x67, 0xcc, 0x8c, 0x17, 0xaa, 0xf4, 0x8c, 0xee, 0x42, 0x9a,
	0x2b, 0x0c, 0xbe, 0xbd, 0x00, 0x67, 0x6e, 0x26, 0x77, 0x92, 0x0f, 0x61, 0x3d, 0x69, 0x47, 0x35,
	0x28, 0x4e, 0x07, 0x74, 0x10, 0x4c, 0xb0, 0x1d, 0xf6, 0x03, 0x42, 0x73, 0xa4, 0xcc, 0x3c, 0x17,
	0x1a, 0xc1, 0x04, 0xb7, 0xfb, 0x01, 0xa9, 0x7d, 0x11, 0xa6, 0x7d, 0xa1, 0xcb, 0x81, 0xca, 0x50,
	0xd2, 0x54, 0xcb, 0x3a, 0xd6, 0xed, 0xb6, 0xa5, 0x5a, 0xba, 0xdd, 0x69, 0x3e, 0x6f, 0xb6, 0xce,
	0x9a, 0x85, 0xff, 0xd0, 0x2e, 0xec, 0x9c, 0xa9, 0x86, 0x65, 0x34, 0x8f, 0xec, 0x56, 0xd3, 0xd6,
	0x5a, 0xd6, 0x33, 0xfb, 0xe4, 0x58, 0x3d, 0xd7, 0xcd, 0x76, 0x41, 0x40, 0x3b, 0xb0, 0x95, 0x10,
	0x99, 0xdd, 0x56, 0x0b, 0xe2, 0x72, 0x41, 0x2b, 0xac, 0x44, 0x42, 0xa3, 0x75, 0xaa, 0xdb, 0xa6,
	0xde, 0xee, 0x1c, 0x5b, 0x6d, 0xfb, 0x45, 0x47, 0xef, 0xe8, 0x4f, 0x0b, 0xab, 0x28, 0x0f, 0x39,
	0x5e, 0x41, 0xeb, 0x54, 0x37, 0x0b, 0xa9, 0xfa, 0x37, 0x11, 0x36, 0x78, 0x89, 0x78, 0x34, 0xf1,
	0x2e, 0x30, 0x7a, 0x03, 0x69, 0xb6, 0x5f, 0x68, 0x2f, 0xee, 0xcc, 0x92, 0x5d, 0x96, 0x6e, 0xdd,
	0xa0, 0xb2, 0xa9, 0x93, 0xb7, 0x3f, 0x7d, 0xff, 0xf1, 0x59, 0xcc, 0xcb, 0xa0, 0x4c, 0xee, 0xf1,
	0x8f, 0xf4, 0xa1, 0x50, 0x43, 0xaf, 0x21, 0x1b, 0x2f, 0x16, 0xfa, 0x3f, 0x46, 0x2c, 0x2e, 0xa4,
	0x24, 0x2d, 0x93, 0x38, 0x7a, 0x87, 0xa2, 0x8b, 0x28, 0x3f, 0x43, 0x2b, 0x57, 0x9e, 0x7b, 0x8d,
	0x3e, 0x00, 0xcc, 0xe6, 0x1f, 0xcd, 0x10, 0xbf, 0x2c, 0xa3, 0xb4, 0xbb, 0x54, 0xe3, 0xfc, 0x07,
	0x94, 0xaf, 0xc8, 0xfb, 0x49, 0x7e, 0xbc, 0x42, 0xd7, 0x4a, 0x34, 0x99, 0xca, 0x55, 0xbc, 0x36,
	0xd7, 0x87, 0x74, 0x09, 0xb4, 0xec, 0xab, 0xe9, 0xef, 0xa7, 0x9b, 0xa6, 0xbf, 0x91, 0xfb, 0x3f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x14, 0x75, 0x41, 0x9e, 0xa6, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BattleServiceClient is the client API for BattleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BattleServiceClient interface {
	Create(ctx context.Context, in *CreateBattleRequest, opts ...grpc.CallOption) (*CreateBattleResponse, error)
	GetBattle(ctx context.Context, in *GetBattleRequest, opts ...grpc.CallOption) (*GetBattleResponse, error)
	SubmitTurn(ctx context.Context, in *SubmitTurnRequest, opts ...grpc.CallOption) (*SubmitTurnResponse, error)
}

type battleServiceClient struct {
	cc *grpc.ClientConn
}

func NewBattleServiceClient(cc *grpc.ClientConn) BattleServiceClient {
	return &battleServiceClient{cc}
}

func (c *battleServiceClient) Create(ctx context.Context, in *CreateBattleRequest, opts ...grpc.CallOption) (*CreateBattleResponse, error) {
	out := new(CreateBattleResponse)
	err := c.cc.Invoke(ctx, "/pokemon.BattleService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *battleServiceClient) GetBattle(ctx context.Context, in *GetBattleRequest, opts ...grpc.CallOption) (*GetBattleResponse, error) {
	out := new(GetBattleResponse)
	err := c.cc.Invoke(ctx, "/pokemon.BattleService/GetBattle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *battleServiceClient) SubmitTurn(ctx context.Context, in *SubmitTurnRequest, opts ...grpc.CallOption) (*SubmitTurnResponse, error) {
	out := new(SubmitTurnResponse)
	err := c.cc.Invoke(ctx, "/pokemon.BattleService/SubmitTurn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BattleServiceServer is the server API for BattleService service.
type BattleServiceServer interface {
	Create(context.Context, *CreateBattleRequest) (*CreateBattleResponse, error)
	GetBattle(context.Context, *GetBattleRequest) (*GetBattleResponse, error)
	SubmitTurn(context.Context, *SubmitTurnRequest) (*SubmitTurnResponse, error)
}

// UnimplementedBattleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBattleServiceServer struct {
}

func (*UnimplementedBattleServiceServer) Create(ctx context.Context, req *CreateBattleRequest) (*CreateBattleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedBattleServiceServer) GetBattle(ctx context.Context, req *GetBattleRequest) (*GetBattleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBattle not implemented")
}
func (*UnimplementedBattleServiceServer) SubmitTurn(ctx context.Context, req *SubmitTurnRequest) (*SubmitTurnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTurn not implemented")
}

func RegisterBattleServiceServer(s *grpc.Server, srv BattleServiceServer) {
	s.RegisterService(&_BattleService_serviceDesc, srv)
}

func _BattleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBattleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BattleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.BattleService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BattleServiceServer).Create(ctx, req.(*CreateBattleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BattleService_GetBattle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBattleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BattleServiceServer).GetBattle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.BattleService/GetBattle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BattleServiceServer).GetBattle(ctx, req.(*GetBattleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BattleService_SubmitTurn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTurnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BattleServiceServer).SubmitTurn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.BattleService/SubmitTurn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BattleServiceServer).SubmitTurn(ctx, req.(*SubmitTurnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BattleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.BattleService",
	HandlerType: (*BattleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BattleService_Create_Handler,
		},
		{
			MethodName: "GetBattle",
			Handler:    _BattleService_GetBattle_Handler,
		},
		{
			MethodName: "SubmitTurn",
			Handler:    _BattleService_SubmitTurn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokemon/battle.proto",
}
