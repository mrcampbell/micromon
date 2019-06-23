// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pokemon.proto

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
	BreedId              int32    `protobuf:"varint,2,opt,name=breed_id,json=breedId,proto3" json:"breed_id,omitempty" sql:",notnull"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xc8, 0xcf, 0x4e,
	0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x64, 0xd2,
	0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12,
	0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xca, 0x94, 0x4c, 0xb8, 0xd8, 0x03, 0x20, 0x0a, 0x85, 0xf8,
	0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x84, 0x24,
	0xb9, 0x38, 0x92, 0x8a, 0x52, 0x53, 0x53, 0xe2, 0x33, 0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58,
	0x83, 0xd8, 0xc1, 0x7c, 0xcf, 0x14, 0x25, 0x65, 0x2e, 0x41, 0xf7, 0xd4, 0x12, 0xa8, 0xc6, 0xa0,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x74, 0xfd, 0x4a, 0x0e, 0x5c, 0x42, 0xc8, 0x8a, 0x8a, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xb4, 0xb8, 0x60, 0x2e, 0x03, 0x2b, 0xe5, 0x36, 0x12, 0xd0, 0x83,
	0x39, 0x1c, 0xa6, 0x14, 0xa6, 0x40, 0x49, 0x84, 0x4b, 0xc8, 0x27, 0xb3, 0x18, 0xcd, 0x1e, 0x25,
	0x47, 0x2e, 0x61, 0x14, 0x51, 0x6c, 0x06, 0x33, 0xe3, 0x35, 0xd8, 0xe8, 0x02, 0x23, 0x17, 0x1f,
	0x54, 0x30, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x9e, 0x8b, 0x0b, 0xe1, 0x5a, 0x21,
	0x29, 0xb8, 0x5e, 0x0c, 0x7f, 0x4a, 0x49, 0x63, 0x95, 0x83, 0xb8, 0x42, 0x49, 0xa2, 0xe9, 0xf2,
	0x93, 0xc9, 0x4c, 0x42, 0x42, 0x02, 0xfa, 0x65, 0x86, 0xfa, 0x50, 0x75, 0xfa, 0xd5, 0x99, 0x29,
	0xb5, 0x42, 0xb1, 0x5c, 0xdc, 0x48, 0xce, 0x16, 0x42, 0x98, 0x82, 0xe9, 0x45, 0x29, 0x19, 0xec,
	0x92, 0x50, 0x3b, 0x84, 0xc1, 0x76, 0xf0, 0x0a, 0x71, 0x23, 0xd9, 0xe1, 0xc4, 0x19, 0x05, 0xf3,
	0x5d, 0x12, 0x1b, 0x38, 0x6a, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xcc, 0x16, 0xf5,
	0x12, 0x02, 0x00, 0x00,
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
