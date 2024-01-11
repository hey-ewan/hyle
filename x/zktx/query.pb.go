// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hyle/hyle/zktx/v1/query.proto

package zktx

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryCounterRequest is the request type for the Query/Counter RPC
// method.
type QueryCounterRequest struct {
	// address defines the address to query for the counter.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryCounterRequest) Reset()         { *m = QueryCounterRequest{} }
func (m *QueryCounterRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCounterRequest) ProtoMessage()    {}
func (*QueryCounterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbc4c2692c495b87, []int{0}
}
func (m *QueryCounterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCounterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCounterRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCounterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCounterRequest.Merge(m, src)
}
func (m *QueryCounterRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCounterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCounterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCounterRequest proto.InternalMessageInfo

func (m *QueryCounterRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// QueryCounterResponse is the response type for the Query/Counter RPC
// method.
type QueryCounterResponse struct {
	// counter defines the current counter for the sender.
	Counter uint64 `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (m *QueryCounterResponse) Reset()         { *m = QueryCounterResponse{} }
func (m *QueryCounterResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCounterResponse) ProtoMessage()    {}
func (*QueryCounterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbc4c2692c495b87, []int{1}
}
func (m *QueryCounterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCounterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCounterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCounterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCounterResponse.Merge(m, src)
}
func (m *QueryCounterResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCounterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCounterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCounterResponse proto.InternalMessageInfo

func (m *QueryCounterResponse) GetCounter() uint64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

// QueryCountersResponse is the request type for the Query/Counters RPC
type QueryCountersRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCountersRequest) Reset()         { *m = QueryCountersRequest{} }
func (m *QueryCountersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCountersRequest) ProtoMessage()    {}
func (*QueryCountersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbc4c2692c495b87, []int{2}
}
func (m *QueryCountersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCountersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCountersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCountersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCountersRequest.Merge(m, src)
}
func (m *QueryCountersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCountersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCountersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCountersRequest proto.InternalMessageInfo

func (m *QueryCountersRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryCountersResponse is the response type for the Query/Counters RPC
// method.
type QueryCountersResponse struct {
	// counters defines all the counters in the store.
	Counters []*Counter `protobuf:"bytes,1,rep,name=counters,proto3" json:"counters,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCountersResponse) Reset()         { *m = QueryCountersResponse{} }
func (m *QueryCountersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCountersResponse) ProtoMessage()    {}
func (*QueryCountersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbc4c2692c495b87, []int{3}
}
func (m *QueryCountersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCountersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCountersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCountersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCountersResponse.Merge(m, src)
}
func (m *QueryCountersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCountersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCountersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCountersResponse proto.InternalMessageInfo

func (m *QueryCountersResponse) GetCounters() []*Counter {
	if m != nil {
		return m.Counters
	}
	return nil
}

func (m *QueryCountersResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbc4c2692c495b87, []int{4}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbc4c2692c495b87, []int{5}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryCounterRequest)(nil), "hyle.hyle.zktx.v1.QueryCounterRequest")
	proto.RegisterType((*QueryCounterResponse)(nil), "hyle.hyle.zktx.v1.QueryCounterResponse")
	proto.RegisterType((*QueryCountersRequest)(nil), "hyle.hyle.zktx.v1.QueryCountersRequest")
	proto.RegisterType((*QueryCountersResponse)(nil), "hyle.hyle.zktx.v1.QueryCountersResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "hyle.hyle.zktx.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "hyle.hyle.zktx.v1.QueryParamsResponse")
}

func init() { proto.RegisterFile("hyle/hyle/zktx/v1/query.proto", fileDescriptor_bbc4c2692c495b87) }

var fileDescriptor_bbc4c2692c495b87 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0x8e, 0x5b, 0x48, 0xda, 0xeb, 0xd4, 0x23, 0x48, 0x89, 0x43, 0x0d, 0x58, 0x90, 0x46, 0x95,
	0xb8, 0x23, 0xa9, 0xc4, 0xc4, 0x42, 0x90, 0x40, 0x6c, 0xc5, 0xdd, 0x18, 0x40, 0x97, 0xf4, 0xe4,
	0x5a, 0x34, 0x3e, 0xd7, 0x77, 0x89, 0x48, 0x11, 0x0b, 0x0b, 0x2c, 0x48, 0x48, 0x2c, 0x2c, 0xec,
	0x8c, 0x0c, 0xfc, 0x88, 0x8e, 0x15, 0x2c, 0x4c, 0x08, 0x25, 0x48, 0xfc, 0x0d, 0xe4, 0xbb, 0x77,
	0x6d, 0xdd, 0xba, 0xf2, 0x72, 0xf2, 0xbd, 0xf7, 0xbe, 0xf7, 0x7d, 0xef, 0x7b, 0x67, 0xb4, 0xb6,
	0x3b, 0xdd, 0xe3, 0x54, 0x1f, 0x07, 0x2f, 0xd5, 0x2b, 0x3a, 0xe9, 0xd2, 0xfd, 0x31, 0x4f, 0xa7,
	0x24, 0x49, 0x85, 0x12, 0x78, 0x35, 0xcb, 0x10, 0x7d, 0x64, 0x69, 0x32, 0xe9, 0xba, 0x1b, 0x43,
	0x21, 0x47, 0x42, 0xd2, 0x01, 0x93, 0xdc, 0xd4, 0xd2, 0x49, 0x77, 0xc0, 0x15, 0xeb, 0xd2, 0x84,
	0x85, 0x51, 0xcc, 0x54, 0x24, 0x62, 0x03, 0x77, 0x0b, 0xba, 0xab, 0x69, 0xc2, 0x25, 0xa4, 0xaf,
	0x85, 0x42, 0x84, 0x7b, 0x9c, 0xb2, 0x24, 0xa2, 0x2c, 0x8e, 0x85, 0xd2, 0x58, 0x9b, 0x6d, 0x01,
	0x91, 0xe5, 0x38, 0x2d, 0xcc, 0x5d, 0x65, 0xa3, 0x28, 0x16, 0x54, 0x9f, 0x10, 0xaa, 0x87, 0x22,
	0x14, 0xfa, 0x93, 0x66, 0x5f, 0x10, 0x6d, 0x9a, 0x2e, 0x2f, 0x4c, 0xc2, 0x5c, 0x4c, 0xca, 0x7f,
	0x82, 0xae, 0x3c, 0xcd, 0x5a, 0x3e, 0x14, 0xe3, 0x58, 0xf1, 0x34, 0xe0, 0xfb, 0x63, 0x2e, 0x15,
	0xee, 0xa1, 0x1a, 0xdb, 0xd9, 0x49, 0xb9, 0x94, 0x0d, 0xe7, 0x86, 0xd3, 0x59, 0xee, 0x37, 0x7e,
	0x7c, 0xbf, 0x53, 0x07, 0xe4, 0x03, 0x93, 0xd9, 0x56, 0x69, 0x14, 0x87, 0x81, 0x2d, 0xf4, 0xef,
	0xa2, 0x7a, 0xbe, 0x95, 0x4c, 0x44, 0x2c, 0x39, 0x6e, 0xa0, 0xda, 0xd0, 0x84, 0x74, 0xaf, 0x4b,
	0x81, 0xbd, 0xfa, 0xcf, 0xf3, 0x08, 0x69, 0xd9, 0x1f, 0x21, 0x74, 0x62, 0xa3, 0x06, 0xad, 0xf4,
	0xda, 0x04, 0xd8, 0x33, 0xcf, 0x89, 0xb1, 0x01, 0x3c, 0x27, 0x5b, 0x2c, 0xe4, 0x80, 0x0d, 0x4e,
	0x21, 0xfd, 0xcf, 0x0e, 0xba, 0x7a, 0x86, 0x00, 0x34, 0xdd, 0x43, 0x4b, 0x20, 0x22, 0x1b, 0x70,
	0xb1, 0xb3, 0xd2, 0x73, 0xc9, 0xb9, 0x35, 0x13, 0x3b, 0xc9, 0x71, 0x2d, 0x7e, 0x9c, 0x53, 0xb6,
	0xa0, 0x95, 0xad, 0x97, 0x2a, 0x33, 0xa4, 0x39, 0x69, 0x75, 0x84, 0xb5, 0xb2, 0x2d, 0x96, 0xb2,
	0x91, 0x1d, 0xdc, 0xdf, 0x86, 0x6d, 0xd8, 0x28, 0xa8, 0xbd, 0x8f, 0xaa, 0x89, 0x8e, 0x80, 0x17,
	0xcd, 0x02, 0xad, 0x06, 0xd2, 0x5f, 0x3e, 0xfc, 0x7d, 0xbd, 0xf2, 0xf5, 0xdf, 0xb7, 0x0d, 0x27,
	0x00, 0x4c, 0xef, 0xcb, 0x22, 0xba, 0xac, 0xbb, 0xe2, 0x0f, 0x0e, 0xaa, 0xc1, 0x4c, 0xb8, 0x5d,
	0xd0, 0xa3, 0xe0, 0x25, 0xb8, 0xeb, 0xa5, 0x75, 0x46, 0xa4, 0xdf, 0x7d, 0x9f, 0xb1, 0xbe, 0xfd,
	0xf9, 0xf7, 0xd3, 0x42, 0x1b, 0xdf, 0xa2, 0xe7, 0x5f, 0x3d, 0x98, 0x48, 0x5f, 0xc3, 0x83, 0x79,
	0x83, 0xdf, 0x39, 0x68, 0xc9, 0xae, 0x06, 0x97, 0x11, 0x59, 0x93, 0xdc, 0x4e, 0x79, 0x21, 0x48,
	0xea, 0x9c, 0x48, 0x5a, 0xc3, 0xad, 0x8b, 0x25, 0x49, 0x7c, 0x80, 0xaa, 0xc6, 0x40, 0x7c, 0xfb,
	0xa2, 0xee, 0xb9, 0x4d, 0xb9, 0xed, 0xb2, 0x32, 0x90, 0x70, 0x53, 0xb3, 0xb7, 0x70, 0xb3, 0x80,
	0xdd, 0xec, 0xa7, 0xbf, 0x79, 0x38, 0xf3, 0x9c, 0xa3, 0x99, 0xe7, 0xfc, 0x99, 0x79, 0xce, 0xc7,
	0xb9, 0x57, 0x39, 0x9a, 0x7b, 0x95, 0x5f, 0x73, 0xaf, 0xf2, 0xac, 0x19, 0x46, 0x6a, 0x77, 0x3c,
	0x20, 0x43, 0x31, 0x3a, 0x03, 0x1f, 0x54, 0xf5, 0xef, 0xbb, 0xf9, 0x3f, 0x00, 0x00, 0xff, 0xff,
	0xcb, 0x90, 0x94, 0xa4, 0xbc, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Counter returns the current counter value.
	Counter(ctx context.Context, in *QueryCounterRequest, opts ...grpc.CallOption) (*QueryCounterResponse, error)
	// Counters returns all the counter values.
	Counters(ctx context.Context, in *QueryCountersRequest, opts ...grpc.CallOption) (*QueryCountersResponse, error)
	// Params returns the module parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Counter(ctx context.Context, in *QueryCounterRequest, opts ...grpc.CallOption) (*QueryCounterResponse, error) {
	out := new(QueryCounterResponse)
	err := c.cc.Invoke(ctx, "/hyle.hyle.zktx.v1.Query/Counter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Counters(ctx context.Context, in *QueryCountersRequest, opts ...grpc.CallOption) (*QueryCountersResponse, error) {
	out := new(QueryCountersResponse)
	err := c.cc.Invoke(ctx, "/hyle.hyle.zktx.v1.Query/Counters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/hyle.hyle.zktx.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Counter returns the current counter value.
	Counter(context.Context, *QueryCounterRequest) (*QueryCounterResponse, error)
	// Counters returns all the counter values.
	Counters(context.Context, *QueryCountersRequest) (*QueryCountersResponse, error)
	// Params returns the module parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Counter(ctx context.Context, req *QueryCounterRequest) (*QueryCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Counter not implemented")
}
func (*UnimplementedQueryServer) Counters(ctx context.Context, req *QueryCountersRequest) (*QueryCountersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Counters not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Counter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Counter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyle.hyle.zktx.v1.Query/Counter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Counter(ctx, req.(*QueryCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Counters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCountersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Counters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyle.hyle.zktx.v1.Query/Counters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Counters(ctx, req.(*QueryCountersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyle.hyle.zktx.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hyle.hyle.zktx.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Counter",
			Handler:    _Query_Counter_Handler,
		},
		{
			MethodName: "Counters",
			Handler:    _Query_Counters_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hyle/hyle/zktx/v1/query.proto",
}

func (m *QueryCounterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCounterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCounterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCounterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCounterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCounterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Counter != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Counter))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCountersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCountersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCountersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCountersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCountersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCountersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Counters) > 0 {
		for iNdEx := len(m.Counters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Counters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCounterRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCounterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Counter != 0 {
		n += 1 + sovQuery(uint64(m.Counter))
	}
	return n
}

func (m *QueryCountersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCountersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Counters) > 0 {
		for _, e := range m.Counters {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCounterRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCounterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCounterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCounterResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCounterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCounterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCountersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCountersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCountersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCountersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCountersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCountersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Counters = append(m.Counters, &Counter{})
			if err := m.Counters[len(m.Counters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
