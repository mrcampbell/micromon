// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pokemon.proto

package pokemon

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	BreedId int32 `protobuf:"varint,2,opt,name=breed_id,json=breedId,proto3" json:"breed_id,omitempty" sql:",notnull"`
	// @inject_tag: sql:"type:timestamptz,default:now()"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" sql:"type:timestamptz,default:now()"`
	// @inject_tag: sql:"type:timestamptz"
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" sql:"type:timestamptz"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pokemon) Reset()         { *m = Pokemon{} }
func (m *Pokemon) String() string { return proto.CompactTextString(m) }
func (*Pokemon) ProtoMessage()    {}
func (*Pokemon) Descriptor() ([]byte, []int) {
	return fileDescriptor_75142c9f6896f36f, []int{0}
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

func (m *Pokemon) GetBreedId() int32 {
	if m != nil {
		return m.BreedId
	}
	return 0
}

func (m *Pokemon) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Pokemon) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
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
	return fileDescriptor_75142c9f6896f36f, []int{1}
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
	return fileDescriptor_75142c9f6896f36f, []int{2}
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
	return fileDescriptor_75142c9f6896f36f, []int{3}
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
	return fileDescriptor_75142c9f6896f36f, []int{4}
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

func init() {
	proto.RegisterType((*Pokemon)(nil), "pokemon.Pokemon")
	proto.RegisterType((*GetPokemonRequest)(nil), "pokemon.GetPokemonRequest")
	proto.RegisterType((*GetPokemonResponse)(nil), "pokemon.GetPokemonResponse")
	proto.RegisterType((*ListPokemonRequest)(nil), "pokemon.ListPokemonRequest")
	proto.RegisterType((*ListPokemonResponse)(nil), "pokemon.ListPokemonResponse")
}

func init() { proto.RegisterFile("pokemon.proto", fileDescriptor_75142c9f6896f36f) }

var fileDescriptor_75142c9f6896f36f = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0xd3, 0x1a, 0x40, 0xa6, 0x81, 0xe0, 0xe0, 0xa1, 0x56, 0x12, 0x49, 0xbd, 0x10, 0x0f,
	0x6d, 0xc4, 0x93, 0x37, 0xf1, 0x62, 0x4c, 0x3c, 0x98, 0xea, 0xc9, 0xc4, 0x90, 0xc2, 0x8e, 0x64,
	0xa3, 0x74, 0x6b, 0x77, 0xe1, 0x62, 0xbc, 0xf8, 0x0a, 0xbe, 0x85, 0x4f, 0x63, 0xe2, 0x2b, 0xf8,
	0x20, 0xc6, 0x76, 0x57, 0x10, 0x1a, 0x3d, 0xce, 0xce, 0x37, 0xff, 0xff, 0xcf, 0x2c, 0x34, 0x52,
	0x71, 0x4f, 0x53, 0x91, 0x04, 0x69, 0x26, 0x94, 0xc0, 0x9a, 0x2e, 0xbd, 0xce, 0x44, 0x88, 0xc9,
	0x03, 0x85, 0x71, 0xca, 0xc3, 0x38, 0x49, 0x84, 0x8a, 0x15, 0x17, 0x89, 0x2c, 0x30, 0x6f, 0x4f,
	0x77, 0xf3, 0x6a, 0x34, 0xbb, 0x0b, 0x15, 0x9f, 0x92, 0x54, 0xf1, 0x34, 0x2d, 0x00, 0xff, 0xcd,
	0x82, 0xda, 0x65, 0x21, 0x85, 0x4d, 0xb0, 0x39, 0x73, 0xad, 0xae, 0xd5, 0xab, 0x47, 0x36, 0x67,
	0xb8, 0x03, 0x9b, 0xa3, 0x8c, 0x88, 0x0d, 0x39, 0x73, 0xed, 0xae, 0xd5, 0xab, 0x44, 0xb5, 0xbc,
	0x3e, 0x67, 0x78, 0x0c, 0x30, 0xce, 0x28, 0x56, 0xc4, 0x86, 0xb1, 0x72, 0x2b, 0x5d, 0xab, 0xe7,
	0xf4, 0xbd, 0xa0, 0x30, 0x0b, 0x8c, 0x59, 0x70, 0x6d, 0xcc, 0xa2, 0xba, 0xa6, 0x07, 0xea, 0x7b,
	0x74, 0x96, 0x32, 0x33, 0x5a, 0xfd, 0x7f, 0x54, 0xd3, 0x03, 0xe5, 0xef, 0xc3, 0xd6, 0x19, 0x29,
	0x1d, 0x37, 0xa2, 0xc7, 0x19, 0x49, 0xb5, 0x9a, 0xda, 0x3f, 0x01, 0x5c, 0x86, 0x64, 0x2a, 0x12,
	0x49, 0x78, 0x00, 0xe6, 0x62, 0x39, 0xea, 0xf4, 0x5b, 0x81, 0x39, 0xa8, 0x41, 0x0d, 0xe0, 0x6f,
	0x03, 0x5e, 0x70, 0xb9, 0xe2, 0xe3, 0x0f, 0xa0, 0xfd, 0xeb, 0xb5, 0x4c, 0x78, 0xe3, 0x4f, 0xe1,
	0xfe, 0xbb, 0x05, 0x4d, 0xfd, 0x78, 0x45, 0xd9, 0x9c, 0x8f, 0x09, 0x87, 0x00, 0x8b, 0xb4, 0xe8,
	0xfd, 0xcc, 0xae, 0xed, 0xe9, 0xed, 0x96, 0xf6, 0x8a, 0x14, 0xbe, 0xfb, 0xf2, 0xf1, 0xf9, 0x6a,
	0x23, 0xb6, 0xc2, 0xf9, 0x61, 0xa8, 0xb9, 0xf0, 0x89, 0xb3, 0x67, 0xbc, 0x05, 0x67, 0x29, 0x36,
	0x2e, 0x54, 0xd6, 0x57, 0xf4, 0x3a, 0xe5, 0x4d, 0xed, 0xd1, 0xce, 0x3d, 0x1a, 0xe8, 0x2c, 0x79,
	0x9c, 0xd6, 0x6f, 0xcc, 0x76, 0xa3, 0x6a, 0xfe, 0x79, 0x47, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xdc, 0x02, 0xa3, 0x06, 0xaa, 0x02, 0x00, 0x00,
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

// PokemonServiceServer is the server API for PokemonService service.
type PokemonServiceServer interface {
	GetPokemon(context.Context, *GetPokemonRequest) (*GetPokemonResponse, error)
	ListPokemon(context.Context, *ListPokemonRequest) (*ListPokemonResponse, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokemon.proto",
}
