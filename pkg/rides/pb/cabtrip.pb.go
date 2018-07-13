// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/cabtrip.proto

package pb // import "github.com/govinda-attal/cabride-api/pkg/rides/pb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FetchRq struct {
	Pickup               string   `protobuf:"bytes,1,opt,name=pickup,proto3" json:"pickup,omitempty"`
	Medallions           []string `protobuf:"bytes,2,rep,name=medallions,proto3" json:"medallions,omitempty"`
	NoCache              bool     `protobuf:"varint,3,opt,name=noCache,proto3" json:"noCache,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchRq) Reset()         { *m = FetchRq{} }
func (m *FetchRq) String() string { return proto.CompactTextString(m) }
func (*FetchRq) ProtoMessage()    {}
func (*FetchRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cabtrip_27cd66ad09de008a, []int{0}
}
func (m *FetchRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchRq.Unmarshal(m, b)
}
func (m *FetchRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchRq.Marshal(b, m, deterministic)
}
func (dst *FetchRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRq.Merge(dst, src)
}
func (m *FetchRq) XXX_Size() int {
	return xxx_messageInfo_FetchRq.Size(m)
}
func (m *FetchRq) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRq.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRq proto.InternalMessageInfo

func (m *FetchRq) GetPickup() string {
	if m != nil {
		return m.Pickup
	}
	return ""
}

func (m *FetchRq) GetMedallions() []string {
	if m != nil {
		return m.Medallions
	}
	return nil
}

func (m *FetchRq) GetNoCache() bool {
	if m != nil {
		return m.NoCache
	}
	return false
}

type FetchRs struct {
	Pickup               string      `protobuf:"bytes,1,opt,name=pickup,proto3" json:"pickup,omitempty"`
	Trips                []*TripData `protobuf:"bytes,2,rep,name=trips,proto3" json:"trips,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FetchRs) Reset()         { *m = FetchRs{} }
func (m *FetchRs) String() string { return proto.CompactTextString(m) }
func (*FetchRs) ProtoMessage()    {}
func (*FetchRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_cabtrip_27cd66ad09de008a, []int{1}
}
func (m *FetchRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchRs.Unmarshal(m, b)
}
func (m *FetchRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchRs.Marshal(b, m, deterministic)
}
func (dst *FetchRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRs.Merge(dst, src)
}
func (m *FetchRs) XXX_Size() int {
	return xxx_messageInfo_FetchRs.Size(m)
}
func (m *FetchRs) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRs.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRs proto.InternalMessageInfo

func (m *FetchRs) GetPickup() string {
	if m != nil {
		return m.Pickup
	}
	return ""
}

func (m *FetchRs) GetTrips() []*TripData {
	if m != nil {
		return m.Trips
	}
	return nil
}

type ClearCacheRq struct {
	Pickup               string   `protobuf:"bytes,1,opt,name=pickup,proto3" json:"pickup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearCacheRq) Reset()         { *m = ClearCacheRq{} }
func (m *ClearCacheRq) String() string { return proto.CompactTextString(m) }
func (*ClearCacheRq) ProtoMessage()    {}
func (*ClearCacheRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cabtrip_27cd66ad09de008a, []int{2}
}
func (m *ClearCacheRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearCacheRq.Unmarshal(m, b)
}
func (m *ClearCacheRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearCacheRq.Marshal(b, m, deterministic)
}
func (dst *ClearCacheRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearCacheRq.Merge(dst, src)
}
func (m *ClearCacheRq) XXX_Size() int {
	return xxx_messageInfo_ClearCacheRq.Size(m)
}
func (m *ClearCacheRq) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearCacheRq.DiscardUnknown(m)
}

var xxx_messageInfo_ClearCacheRq proto.InternalMessageInfo

func (m *ClearCacheRq) GetPickup() string {
	if m != nil {
		return m.Pickup
	}
	return ""
}

type ClearCacheRs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearCacheRs) Reset()         { *m = ClearCacheRs{} }
func (m *ClearCacheRs) String() string { return proto.CompactTextString(m) }
func (*ClearCacheRs) ProtoMessage()    {}
func (*ClearCacheRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_cabtrip_27cd66ad09de008a, []int{3}
}
func (m *ClearCacheRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearCacheRs.Unmarshal(m, b)
}
func (m *ClearCacheRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearCacheRs.Marshal(b, m, deterministic)
}
func (dst *ClearCacheRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearCacheRs.Merge(dst, src)
}
func (m *ClearCacheRs) XXX_Size() int {
	return xxx_messageInfo_ClearCacheRs.Size(m)
}
func (m *ClearCacheRs) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearCacheRs.DiscardUnknown(m)
}

var xxx_messageInfo_ClearCacheRs proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FetchRq)(nil), "proto.FetchRq")
	proto.RegisterType((*FetchRs)(nil), "proto.FetchRs")
	proto.RegisterType((*ClearCacheRq)(nil), "proto.ClearCacheRq")
	proto.RegisterType((*ClearCacheRs)(nil), "proto.ClearCacheRs")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CabTripsClient is the client API for CabTrips service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CabTripsClient interface {
	Fetch(ctx context.Context, in *FetchRq, opts ...grpc.CallOption) (*FetchRs, error)
	ClearCache(ctx context.Context, in *ClearCacheRq, opts ...grpc.CallOption) (*ClearCacheRs, error)
}

type cabTripsClient struct {
	cc *grpc.ClientConn
}

func NewCabTripsClient(cc *grpc.ClientConn) CabTripsClient {
	return &cabTripsClient{cc}
}

func (c *cabTripsClient) Fetch(ctx context.Context, in *FetchRq, opts ...grpc.CallOption) (*FetchRs, error) {
	out := new(FetchRs)
	err := c.cc.Invoke(ctx, "/proto.CabTrips/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cabTripsClient) ClearCache(ctx context.Context, in *ClearCacheRq, opts ...grpc.CallOption) (*ClearCacheRs, error) {
	out := new(ClearCacheRs)
	err := c.cc.Invoke(ctx, "/proto.CabTrips/ClearCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CabTripsServer is the server API for CabTrips service.
type CabTripsServer interface {
	Fetch(context.Context, *FetchRq) (*FetchRs, error)
	ClearCache(context.Context, *ClearCacheRq) (*ClearCacheRs, error)
}

func RegisterCabTripsServer(s *grpc.Server, srv CabTripsServer) {
	s.RegisterService(&_CabTrips_serviceDesc, srv)
}

func _CabTrips_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CabTripsServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CabTrips/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CabTripsServer).Fetch(ctx, req.(*FetchRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CabTrips_ClearCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearCacheRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CabTripsServer).ClearCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CabTrips/ClearCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CabTripsServer).ClearCache(ctx, req.(*ClearCacheRq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CabTrips_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CabTrips",
	HandlerType: (*CabTripsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _CabTrips_Fetch_Handler,
		},
		{
			MethodName: "ClearCache",
			Handler:    _CabTrips_ClearCache_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cabtrip.proto",
}

func init() { proto.RegisterFile("proto/cabtrip.proto", fileDescriptor_cabtrip_27cd66ad09de008a) }

var fileDescriptor_cabtrip_27cd66ad09de008a = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0x2b, 0x31,
	0x18, 0xc5, 0x99, 0x96, 0xfe, 0xcb, 0xbd, 0xb4, 0xdc, 0xf4, 0xa2, 0xc3, 0xa0, 0x52, 0x22, 0x4a,
	0x29, 0xb4, 0xa1, 0xed, 0xce, 0xa5, 0x15, 0x71, 0x21, 0x08, 0x83, 0x2b, 0x5d, 0x7d, 0x93, 0x19,
	0xa7, 0xa1, 0xd3, 0x49, 0x9c, 0xa4, 0xdd, 0x88, 0x1b, 0x5f, 0xc1, 0xa7, 0xf0, 0x79, 0x7c, 0x05,
	0x1f, 0x44, 0x92, 0x4c, 0xb1, 0x05, 0xbb, 0x1a, 0xce, 0xf9, 0x92, 0xdf, 0x39, 0x93, 0x0f, 0x75,
	0x65, 0x21, 0xb4, 0xa0, 0x0c, 0x22, 0x5d, 0x70, 0x39, 0xb2, 0x0a, 0xd7, 0xec, 0x27, 0x38, 0x4a,
	0x85, 0x48, 0xb3, 0x84, 0x82, 0xe4, 0x14, 0xf2, 0x5c, 0x68, 0xd0, 0x5c, 0xe4, 0xca, 0x1d, 0x0a,
	0xfe, 0xbb, 0x9b, 0xe6, 0x5a, 0x0c, 0x1a, 0x9c, 0x4b, 0x1e, 0x51, 0xe3, 0x3a, 0xd1, 0x6c, 0x1e,
	0x3e, 0xe3, 0x03, 0x54, 0x97, 0x9c, 0x2d, 0x56, 0xd2, 0xf7, 0x7a, 0x5e, 0xbf, 0x15, 0x96, 0x0a,
	0x9f, 0x20, 0xb4, 0x4c, 0x62, 0xc8, 0x32, 0x03, 0xf3, 0x2b, 0xbd, 0x6a, 0xbf, 0x15, 0x6e, 0x39,
	0xd8, 0x47, 0x8d, 0x5c, 0xcc, 0x80, 0xcd, 0x13, 0xbf, 0xda, 0xf3, 0xfa, 0xcd, 0x70, 0x23, 0xc9,
	0xcd, 0x06, 0xae, 0xf6, 0xc2, 0xcf, 0x50, 0xcd, 0x34, 0x72, 0xdc, 0x3f, 0x93, 0x8e, 0xab, 0x35,
	0xba, 0x2f, 0xb8, 0xbc, 0x02, 0x0d, 0xa1, 0x9b, 0x92, 0x73, 0xf4, 0x77, 0x96, 0x25, 0x50, 0x58,
	0xee, 0xfe, 0xae, 0xa4, 0xbd, 0x73, 0x4e, 0x4d, 0x3e, 0x3c, 0xd4, 0x9c, 0x41, 0x64, 0x70, 0x0a,
	0xdf, 0xa2, 0x9a, 0xad, 0x83, 0xdb, 0x65, 0x4a, 0xf9, 0xe7, 0xc1, 0xae, 0x56, 0xe4, 0xf4, 0xed,
	0xf3, 0xeb, 0xbd, 0x72, 0x4c, 0x7c, 0xba, 0x1e, 0xdb, 0x07, 0x53, 0xf4, 0xc9, 0x8c, 0xe8, 0x8b,
	0x8b, 0x79, 0xbd, 0xf0, 0x06, 0xf8, 0x0e, 0xa1, 0x9f, 0x28, 0xdc, 0x2d, 0x11, 0xdb, 0x2d, 0x83,
	0x5f, 0x4c, 0x45, 0x0e, 0x2d, 0xfc, 0xdf, 0xa0, 0x63, 0xe0, 0xcc, 0x98, 0x94, 0x99, 0xf9, 0xe5,
	0xf4, 0x61, 0x9c, 0x72, 0x3d, 0x5f, 0x45, 0x23, 0x26, 0x96, 0x34, 0x15, 0x6b, 0x9e, 0xc7, 0x30,
	0x04, 0xad, 0x21, 0x33, 0xfb, 0x2e, 0x78, 0x9c, 0x0c, 0xcd, 0x6a, 0xe5, 0x22, 0xa5, 0x46, 0x28,
	0x2a, 0xa3, 0xa8, 0x6e, 0x13, 0xa6, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xde, 0x45, 0x77,
	0x18, 0x02, 0x00, 0x00,
}
