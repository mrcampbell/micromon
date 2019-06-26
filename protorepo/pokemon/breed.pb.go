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
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type_1               Type     `protobuf:"varint,3,opt,name=type_1,json=type1,proto3,enum=pokemon.Type" json:"type_1,omitempty"`
	Type_2               Type     `protobuf:"varint,4,opt,name=type_2,json=type2,proto3,enum=pokemon.Type" json:"type_2,omitempty"`
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

func (m *BreedSummary) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BreedSummary) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BreedSummary) GetType_1() Type {
	if m != nil {
		return m.Type_1
	}
	return Type_UNKNOWN
}

func (m *BreedSummary) GetType_2() Type {
	if m != nil {
		return m.Type_2
	}
	return Type_UNKNOWN
}

type BreedDetail struct {
	Summary              *BreedSummary `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	Stats                *StatGroup    `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
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

func (m *BreedDetail) GetSummary() *BreedSummary {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *BreedDetail) GetStats() *StatGroup {
	if m != nil {
		return m.Stats
	}
	return nil
}

type GetBreedSummaryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
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

func (m *GetBreedSummaryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetBreedSummaryResponse struct {
	Summary              *BreedSummary `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
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

func (m *GetBreedSummaryResponse) GetSummary() *BreedSummary {
	if m != nil {
		return m.Summary
	}
	return nil
}

type GetBreedDetailRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
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

func (m *GetBreedDetailRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetBreedDetailResponse struct {
	Detail               *BreedDetail `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
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

func (m *GetBreedDetailResponse) GetDetail() *BreedDetail {
	if m != nil {
		return m.Detail
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
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0xc6, 0x49, 0xec, 0x1f, 0x3a, 0xd5, 0x28, 0x4f, 0xdb, 0x86, 0x20, 0xb5, 0x04, 0xc1, 0x2c,
	0xa4, 0xa1, 0xf1, 0x06, 0x45, 0x2c, 0xb8, 0x4c, 0x5d, 0xb9, 0x91, 0xa9, 0x19, 0xea, 0x60, 0x93,
	0x89, 0xc9, 0xb4, 0x50, 0x44, 0x17, 0x5e, 0xc1, 0xa3, 0x79, 0x05, 0xf7, 0x5e, 0x41, 0x32, 0x99,
	0x89, 0xa6, 0x35, 0x0b, 0x77, 0xe1, 0xbd, 0xdf, 0x7c, 0xdf, 0xf7, 0x5e, 0x66, 0xd0, 0x61, 0xcc,
	0x1e, 0x49, 0xc8, 0x22, 0x77, 0x96, 0x10, 0x12, 0x0c, 0xe3, 0x84, 0x71, 0x06, 0x4d, 0x59, 0xb4,
	0x8e, 0xe7, 0x8c, 0xcd, 0x17, 0xc4, 0xc5, 0x31, 0x75, 0x71, 0x14, 0x31, 0x8e, 0x39, 0x65, 0x51,
	0x9a, 0x63, 0x16, 0xa8, 0xb3, 0x7c, 0x1d, 0x13, 0x59, 0xeb, 0xa9, 0x5a, 0xca, 0x31, 0x9f, 0x27,
	0x6c, 0x19, 0xe7, 0x0d, 0xfb, 0x15, 0xed, 0x8e, 0x33, 0x8b, 0xe9, 0x32, 0x0c, 0x71, 0xb2, 0x06,
	0x03, 0xe9, 0x34, 0x30, 0xb5, 0x81, 0xe6, 0xb4, 0x7c, 0x9d, 0x06, 0x00, 0xa8, 0x16, 0xe1, 0x90,
	0x98, 0xba, 0xa8, 0x88, 0x6f, 0x38, 0x45, 0x8d, 0x4c, 0xfa, 0x6e, 0x64, 0xee, 0x0c, 0x34, 0xc7,
	0xf0, 0xf6, 0x86, 0x52, 0x7d, 0x78, 0xb3, 0x8e, 0x89, 0x5f, 0xcf, 0x9a, 0xa3, 0x82, 0xf2, 0xcc,
	0x5a, 0x25, 0xe5, 0xd9, 0x0f, 0xa8, 0x2d, 0xfc, 0x2f, 0x09, 0xc7, 0x74, 0x01, 0x2e, 0x6a, 0xa6,
	0x79, 0x12, 0x91, 0xa1, 0xed, 0x75, 0x8a, 0x53, 0xbf, 0x63, 0xfa, 0x8a, 0x02, 0x07, 0xd5, 0xb3,
	0x91, 0x52, 0x11, 0xb0, 0xed, 0x41, 0x81, 0x4f, 0x39, 0xe6, 0x93, 0x6c, 0x50, 0x3f, 0x07, 0x6c,
	0x07, 0x75, 0x27, 0x84, 0x97, 0x54, 0xc8, 0xd3, 0x92, 0xa4, 0x7c, 0x73, 0x66, 0xfb, 0x1a, 0xf5,
	0xb6, 0xc8, 0x34, 0x66, 0x51, 0x4a, 0xfe, 0x9d, 0xcf, 0x3e, 0x43, 0x1d, 0xa5, 0x95, 0x8f, 0x58,
	0x65, 0x7a, 0xf5, 0x13, 0x4f, 0x81, 0xd2, 0xf3, 0x1c, 0x35, 0x02, 0x51, 0x91, 0x96, 0x47, 0x65,
	0x4b, 0x49, 0x4b, 0xc6, 0xfb, 0xd2, 0xd4, 0x1f, 0x25, 0xc9, 0x8a, 0xde, 0x13, 0xa0, 0xc8, 0x28,
	0x0b, 0x43, 0xbf, 0x10, 0xf8, 0x33, 0x9a, 0x75, 0x52, 0xd9, 0xcf, 0x13, 0xd9, 0xdd, 0xb7, 0x8f,
	0xcf, 0x77, 0xfd, 0x00, 0x0c, 0x77, 0x35, 0xca, 0x6f, 0xa8, 0xfb, 0x4c, 0x83, 0x17, 0xe0, 0x68,
	0x7f, 0x63, 0x71, 0xb0, 0xad, 0x55, 0x5e, 0xbe, 0x35, 0xa8, 0x06, 0xa4, 0x5b, 0x5f, 0xb8, 0x99,
	0xd0, 0x2d, 0xbb, 0xb9, 0x72, 0xc5, 0xe3, 0xd6, 0xad, 0x7a, 0x18, 0xb3, 0x86, 0xb8, 0xd4, 0x17,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xef, 0x8b, 0x5d, 0x3f, 0x03, 0x00, 0x00,
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
