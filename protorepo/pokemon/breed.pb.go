// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pokemon/breed.proto

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

type BreedSummary struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BreedSummary) Reset()         { *m = BreedSummary{} }
func (m *BreedSummary) String() string { return proto.CompactTextString(m) }
func (*BreedSummary) ProtoMessage()    {}
func (*BreedSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a06e1c6efe7453, []int{0}
}

func (m *BreedSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BreedSummary.Unmarshal(m, b)
}
func (m *BreedSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BreedSummary.Marshal(b, m, deterministic)
}
func (m *BreedSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BreedSummary.Merge(m, src)
}
func (m *BreedSummary) XXX_Size() int {
	return xxx_messageInfo_BreedSummary.Size(m)
}
func (m *BreedSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_BreedSummary.DiscardUnknown(m)
}

var xxx_messageInfo_BreedSummary proto.InternalMessageInfo

func (m *BreedSummary) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BreedSummary) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type BreedDetail struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Hp                   int32    `protobuf:"varint,3,opt,name=hp,proto3" json:"hp,omitempty"`
	Attack               int32    `protobuf:"varint,4,opt,name=attack,proto3" json:"attack,omitempty"`
	Defense              int32    `protobuf:"varint,5,opt,name=defense,proto3" json:"defense,omitempty"`
	SpecialAttack        int32    `protobuf:"varint,6,opt,name=special_attack,json=specialAttack,proto3" json:"special_attack,omitempty"`
	SpecialDefense       int32    `protobuf:"varint,7,opt,name=special_defense,json=specialDefense,proto3" json:"special_defense,omitempty"`
	Speed                int32    `protobuf:"varint,8,opt,name=speed,proto3" json:"speed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BreedDetail) Reset()         { *m = BreedDetail{} }
func (m *BreedDetail) String() string { return proto.CompactTextString(m) }
func (*BreedDetail) ProtoMessage()    {}
func (*BreedDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a06e1c6efe7453, []int{1}
}

func (m *BreedDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BreedDetail.Unmarshal(m, b)
}
func (m *BreedDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BreedDetail.Marshal(b, m, deterministic)
}
func (m *BreedDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BreedDetail.Merge(m, src)
}
func (m *BreedDetail) XXX_Size() int {
	return xxx_messageInfo_BreedDetail.Size(m)
}
func (m *BreedDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_BreedDetail.DiscardUnknown(m)
}

var xxx_messageInfo_BreedDetail proto.InternalMessageInfo

func (m *BreedDetail) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BreedDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BreedDetail) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *BreedDetail) GetAttack() int32 {
	if m != nil {
		return m.Attack
	}
	return 0
}

func (m *BreedDetail) GetDefense() int32 {
	if m != nil {
		return m.Defense
	}
	return 0
}

func (m *BreedDetail) GetSpecialAttack() int32 {
	if m != nil {
		return m.SpecialAttack
	}
	return 0
}

func (m *BreedDetail) GetSpecialDefense() int32 {
	if m != nil {
		return m.SpecialDefense
	}
	return 0
}

func (m *BreedDetail) GetSpeed() int32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

type GetBreedSummaryRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBreedSummaryRequest) Reset()         { *m = GetBreedSummaryRequest{} }
func (m *GetBreedSummaryRequest) String() string { return proto.CompactTextString(m) }
func (*GetBreedSummaryRequest) ProtoMessage()    {}
func (*GetBreedSummaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a06e1c6efe7453, []int{2}
}

func (m *GetBreedSummaryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreedSummaryRequest.Unmarshal(m, b)
}
func (m *GetBreedSummaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreedSummaryRequest.Marshal(b, m, deterministic)
}
func (m *GetBreedSummaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreedSummaryRequest.Merge(m, src)
}
func (m *GetBreedSummaryRequest) XXX_Size() int {
	return xxx_messageInfo_GetBreedSummaryRequest.Size(m)
}
func (m *GetBreedSummaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreedSummaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreedSummaryRequest proto.InternalMessageInfo

func (m *GetBreedSummaryRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetBreedSummaryResponse struct {
	Breed                *BreedSummary `protobuf:"bytes,1,opt,name=breed,proto3" json:"breed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetBreedSummaryResponse) Reset()         { *m = GetBreedSummaryResponse{} }
func (m *GetBreedSummaryResponse) String() string { return proto.CompactTextString(m) }
func (*GetBreedSummaryResponse) ProtoMessage()    {}
func (*GetBreedSummaryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a06e1c6efe7453, []int{3}
}

func (m *GetBreedSummaryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreedSummaryResponse.Unmarshal(m, b)
}
func (m *GetBreedSummaryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreedSummaryResponse.Marshal(b, m, deterministic)
}
func (m *GetBreedSummaryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreedSummaryResponse.Merge(m, src)
}
func (m *GetBreedSummaryResponse) XXX_Size() int {
	return xxx_messageInfo_GetBreedSummaryResponse.Size(m)
}
func (m *GetBreedSummaryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreedSummaryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreedSummaryResponse proto.InternalMessageInfo

func (m *GetBreedSummaryResponse) GetBreed() *BreedSummary {
	if m != nil {
		return m.Breed
	}
	return nil
}

type GetBreedDetailRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBreedDetailRequest) Reset()         { *m = GetBreedDetailRequest{} }
func (m *GetBreedDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetBreedDetailRequest) ProtoMessage()    {}
func (*GetBreedDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a06e1c6efe7453, []int{4}
}

func (m *GetBreedDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreedDetailRequest.Unmarshal(m, b)
}
func (m *GetBreedDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreedDetailRequest.Marshal(b, m, deterministic)
}
func (m *GetBreedDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreedDetailRequest.Merge(m, src)
}
func (m *GetBreedDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetBreedDetailRequest.Size(m)
}
func (m *GetBreedDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreedDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreedDetailRequest proto.InternalMessageInfo

func (m *GetBreedDetailRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetBreedDetailResponse struct {
	Breed                *BreedDetail `protobuf:"bytes,1,opt,name=breed,proto3" json:"breed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetBreedDetailResponse) Reset()         { *m = GetBreedDetailResponse{} }
func (m *GetBreedDetailResponse) String() string { return proto.CompactTextString(m) }
func (*GetBreedDetailResponse) ProtoMessage()    {}
func (*GetBreedDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a06e1c6efe7453, []int{5}
}

func (m *GetBreedDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreedDetailResponse.Unmarshal(m, b)
}
func (m *GetBreedDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreedDetailResponse.Marshal(b, m, deterministic)
}
func (m *GetBreedDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreedDetailResponse.Merge(m, src)
}
func (m *GetBreedDetailResponse) XXX_Size() int {
	return xxx_messageInfo_GetBreedDetailResponse.Size(m)
}
func (m *GetBreedDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreedDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreedDetailResponse proto.InternalMessageInfo

func (m *GetBreedDetailResponse) GetBreed() *BreedDetail {
	if m != nil {
		return m.Breed
	}
	return nil
}

func init() {
	proto.RegisterType((*BreedSummary)(nil), "pokemon.BreedSummary")
	proto.RegisterType((*BreedDetail)(nil), "pokemon.BreedDetail")
	proto.RegisterType((*GetBreedSummaryRequest)(nil), "pokemon.GetBreedSummaryRequest")
	proto.RegisterType((*GetBreedSummaryResponse)(nil), "pokemon.GetBreedSummaryResponse")
	proto.RegisterType((*GetBreedDetailRequest)(nil), "pokemon.GetBreedDetailRequest")
	proto.RegisterType((*GetBreedDetailResponse)(nil), "pokemon.GetBreedDetailResponse")
}

func init() { proto.RegisterFile("pokemon/breed.proto", fileDescriptor_b9a06e1c6efe7453) }

var fileDescriptor_b9a06e1c6efe7453 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x8e, 0xd2, 0x40,
	0x14, 0xc6, 0xd3, 0x4a, 0x41, 0x0e, 0x5a, 0xcc, 0x08, 0x75, 0x42, 0x0c, 0x92, 0x26, 0x06, 0xa2,
	0x09, 0x8d, 0xf8, 0x04, 0x12, 0xa2, 0xf7, 0x78, 0xe7, 0x8d, 0x19, 0xe8, 0x11, 0x26, 0xd0, 0xce,
	0xd8, 0x0e, 0x24, 0xc6, 0x78, 0xe3, 0x2b, 0xec, 0xa3, 0xed, 0xd5, 0xde, 0xef, 0xfd, 0xbe, 0xc2,
	0x86, 0xf9, 0xb3, 0xbb, 0x85, 0x25, 0xd9, 0x3b, 0x66, 0xce, 0xef, 0xfb, 0xbe, 0xc3, 0x39, 0x53,
	0x78, 0x2d, 0xc5, 0x06, 0x33, 0x91, 0x27, 0x8b, 0x02, 0x31, 0x1d, 0xcb, 0x42, 0x28, 0x41, 0x1a,
	0xf6, 0xb2, 0xf7, 0x76, 0x25, 0xc4, 0x6a, 0x8b, 0x09, 0x93, 0x3c, 0x61, 0x79, 0x2e, 0x14, 0x53,
	0x5c, 0xe4, 0xa5, 0xc1, 0xe2, 0x09, 0xbc, 0x98, 0x1e, 0x54, 0xdf, 0x77, 0x59, 0xc6, 0x8a, 0x3f,
	0x24, 0x04, 0x9f, 0xa7, 0xd4, 0x1b, 0x78, 0xa3, 0x60, 0xee, 0xf3, 0x94, 0x10, 0xa8, 0xe5, 0x2c,
	0x43, 0xea, 0x0f, 0xbc, 0x51, 0x73, 0xae, 0x7f, 0xc7, 0x57, 0x1e, 0xb4, 0xb4, 0x68, 0x86, 0x8a,
	0xf1, 0xed, 0x53, 0x34, 0x07, 0x66, 0x2d, 0xe9, 0x33, 0xc3, 0xac, 0x25, 0x89, 0xa0, 0xce, 0x94,
	0x62, 0xcb, 0x0d, 0xad, 0xe9, 0x3b, 0x7b, 0x22, 0x14, 0x1a, 0x29, 0xfe, 0xc2, 0xbc, 0x44, 0x1a,
	0xe8, 0x82, 0x3b, 0x92, 0xf7, 0x10, 0x96, 0x12, 0x97, 0x9c, 0x6d, 0x7f, 0x5a, 0x65, 0x5d, 0x03,
	0x2f, 0xed, 0xed, 0x17, 0x63, 0x30, 0x84, 0xb6, 0xc3, 0x9c, 0x51, 0x43, 0x73, 0x4e, 0x3d, 0xb3,
	0x7e, 0x1d, 0x08, 0x4a, 0x89, 0x98, 0xd2, 0xe7, 0xba, 0x6c, 0x0e, 0xf1, 0x08, 0xa2, 0x6f, 0xa8,
	0x1e, 0x8e, 0x64, 0x8e, 0xbf, 0x77, 0x58, 0xaa, 0xe3, 0x7f, 0x19, 0x7f, 0x85, 0x37, 0x27, 0x64,
	0x29, 0xc5, 0xc1, 0xfa, 0x23, 0x04, 0x7a, 0x15, 0x9a, 0x6e, 0x4d, 0xba, 0x63, 0xbb, 0x8b, 0x71,
	0x85, 0x36, 0x4c, 0x3c, 0x84, 0xae, 0xf3, 0x31, 0xf3, 0x3c, 0x17, 0x38, 0xbb, 0x6f, 0xcd, 0x81,
	0x36, 0xef, 0x43, 0x35, 0xaf, 0x53, 0xcd, 0xb3, 0xb0, 0x41, 0x26, 0x37, 0x9e, 0xdb, 0x38, 0x16,
	0x7b, 0xbe, 0x44, 0xc2, 0x21, 0xac, 0xda, 0x92, 0xfe, 0x9d, 0xfe, 0xd1, 0xc6, 0x7a, 0xef, 0xce,
	0xd6, 0x4d, 0x3f, 0x71, 0xf4, 0xff, 0xf2, 0xfa, 0xc2, 0x7f, 0x45, 0xc2, 0x64, 0xff, 0xc9, 0x3c,
	0xca, 0xe4, 0x2f, 0x4f, 0xff, 0x11, 0x05, 0xed, 0xa3, 0x91, 0x91, 0x53, 0xaf, 0xea, 0xd8, 0x7b,
	0x83, 0xf3, 0x80, 0x4d, 0xeb, 0xeb, 0x34, 0x4a, 0xa2, 0x6a, 0x5a, 0x52, 0x1a, 0x6e, 0xda, 0xfc,
	0xe1, 0xbe, 0x85, 0x45, 0x5d, 0x3f, 0xfa, 0xcf, 0xb7, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85, 0xb7,
	0x9c, 0x2b, 0x32, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BreedServiceClient is the client API for BreedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BreedServiceClient interface {
	GetBreedDetail(ctx context.Context, in *GetBreedDetailRequest, opts ...grpc.CallOption) (*GetBreedDetailResponse, error)
	GetBreedSummary(ctx context.Context, in *GetBreedSummaryRequest, opts ...grpc.CallOption) (*GetBreedSummaryResponse, error)
}

type breedServiceClient struct {
	cc *grpc.ClientConn
}

func NewBreedServiceClient(cc *grpc.ClientConn) BreedServiceClient {
	return &breedServiceClient{cc}
}

func (c *breedServiceClient) GetBreedDetail(ctx context.Context, in *GetBreedDetailRequest, opts ...grpc.CallOption) (*GetBreedDetailResponse, error) {
	out := new(GetBreedDetailResponse)
	err := c.cc.Invoke(ctx, "/pokemon.BreedService/GetBreedDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *breedServiceClient) GetBreedSummary(ctx context.Context, in *GetBreedSummaryRequest, opts ...grpc.CallOption) (*GetBreedSummaryResponse, error) {
	out := new(GetBreedSummaryResponse)
	err := c.cc.Invoke(ctx, "/pokemon.BreedService/GetBreedSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BreedServiceServer is the server API for BreedService service.
type BreedServiceServer interface {
	GetBreedDetail(context.Context, *GetBreedDetailRequest) (*GetBreedDetailResponse, error)
	GetBreedSummary(context.Context, *GetBreedSummaryRequest) (*GetBreedSummaryResponse, error)
}

// UnimplementedBreedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBreedServiceServer struct {
}

func (*UnimplementedBreedServiceServer) GetBreedDetail(ctx context.Context, req *GetBreedDetailRequest) (*GetBreedDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBreedDetail not implemented")
}
func (*UnimplementedBreedServiceServer) GetBreedSummary(ctx context.Context, req *GetBreedSummaryRequest) (*GetBreedSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBreedSummary not implemented")
}

func RegisterBreedServiceServer(s *grpc.Server, srv BreedServiceServer) {
	s.RegisterService(&_BreedService_serviceDesc, srv)
}

func _BreedService_GetBreedDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBreedDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreedServiceServer).GetBreedDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.BreedService/GetBreedDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreedServiceServer).GetBreedDetail(ctx, req.(*GetBreedDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BreedService_GetBreedSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBreedSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreedServiceServer).GetBreedSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.BreedService/GetBreedSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreedServiceServer).GetBreedSummary(ctx, req.(*GetBreedSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BreedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.BreedService",
	HandlerType: (*BreedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBreedDetail",
			Handler:    _BreedService_GetBreedDetail_Handler,
		},
		{
			MethodName: "GetBreedSummary",
			Handler:    _BreedService_GetBreedSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokemon/breed.proto",
}
