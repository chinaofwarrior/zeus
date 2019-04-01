// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apiauth.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type PermRequest struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Perm                 string   `protobuf:"bytes,2,opt,name=perm" json:"perm,omitempty"`
	Domain               string   `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PermRequest) Reset()         { *m = PermRequest{} }
func (m *PermRequest) String() string { return proto.CompactTextString(m) }
func (*PermRequest) ProtoMessage()    {}
func (*PermRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_apiauth_b96b8b5d5a80cfe9, []int{0}
}
func (m *PermRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermRequest.Unmarshal(m, b)
}
func (m *PermRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermRequest.Marshal(b, m, deterministic)
}
func (dst *PermRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermRequest.Merge(dst, src)
}
func (m *PermRequest) XXX_Size() int {
	return xxx_messageInfo_PermRequest.Size(m)
}
func (m *PermRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PermRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PermRequest proto.InternalMessageInfo

func (m *PermRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PermRequest) GetPerm() string {
	if m != nil {
		return m.Perm
	}
	return ""
}

func (m *PermRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type PermResponse struct {
	Pass                 bool     `protobuf:"varint,1,opt,name=pass" json:"pass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PermResponse) Reset()         { *m = PermResponse{} }
func (m *PermResponse) String() string { return proto.CompactTextString(m) }
func (*PermResponse) ProtoMessage()    {}
func (*PermResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_apiauth_b96b8b5d5a80cfe9, []int{1}
}
func (m *PermResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermResponse.Unmarshal(m, b)
}
func (m *PermResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermResponse.Marshal(b, m, deterministic)
}
func (dst *PermResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermResponse.Merge(dst, src)
}
func (m *PermResponse) XXX_Size() int {
	return xxx_messageInfo_PermResponse.Size(m)
}
func (m *PermResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PermResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PermResponse proto.InternalMessageInfo

func (m *PermResponse) GetPass() bool {
	if m != nil {
		return m.Pass
	}
	return false
}

func init() {
	proto.RegisterType((*PermRequest)(nil), "proto.PermRequest")
	proto.RegisterType((*PermResponse)(nil), "proto.PermResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Apiauth service

type ApiauthClient interface {
	CheckPerm(ctx context.Context, in *PermRequest, opts ...grpc.CallOption) (*PermResponse, error)
}

type apiauthClient struct {
	cc *grpc.ClientConn
}

func NewApiauthClient(cc *grpc.ClientConn) ApiauthClient {
	return &apiauthClient{cc}
}

func (c *apiauthClient) CheckPerm(ctx context.Context, in *PermRequest, opts ...grpc.CallOption) (*PermResponse, error) {
	out := new(PermResponse)
	err := grpc.Invoke(ctx, "/proto.Apiauth/CheckPerm", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Apiauth service

type ApiauthServer interface {
	CheckPerm(context.Context, *PermRequest) (*PermResponse, error)
}

func RegisterApiauthServer(s *grpc.Server, srv ApiauthServer) {
	s.RegisterService(&_Apiauth_serviceDesc, srv)
}

func _Apiauth_CheckPerm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiauthServer).CheckPerm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Apiauth/CheckPerm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiauthServer).CheckPerm(ctx, req.(*PermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Apiauth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Apiauth",
	HandlerType: (*ApiauthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckPerm",
			Handler:    _Apiauth_CheckPerm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apiauth.proto",
}

func init() { proto.RegisterFile("apiauth.proto", fileDescriptor_apiauth_b96b8b5d5a80cfe9) }

var fileDescriptor_apiauth_b96b8b5d5a80cfe9 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0x8d, 0xeb, 0xae, 0xee, 0xb8, 0x82, 0x8c, 0x50, 0x8a, 0xa7, 0x92, 0x53, 0x4f, 0x39,
	0x28, 0x78, 0x4f, 0x3d, 0x7a, 0x29, 0xf9, 0x06, 0xb1, 0x1d, 0x6c, 0xd0, 0x34, 0x31, 0x7f, 0xc0,
	0x8f, 0x2f, 0x4d, 0x3c, 0xb8, 0xa7, 0x79, 0xef, 0x31, 0xef, 0xf1, 0x83, 0x3b, 0xed, 0x8d, 0xce,
	0x69, 0x11, 0x3e, 0xb8, 0xe4, 0x70, 0x5f, 0x0e, 0x7f, 0x83, 0xdb, 0x91, 0x82, 0x55, 0xf4, 0x9d,
	0x29, 0x26, 0xbc, 0x87, 0x5d, 0x36, 0x73, 0xcb, 0x3a, 0xd6, 0xef, 0xd5, 0x26, 0x11, 0xe1, 0xca,
	0x53, 0xb0, 0xed, 0x65, 0xc7, 0xfa, 0xa3, 0x2a, 0x1a, 0x1b, 0x38, 0xcc, 0xce, 0x6a, 0xb3, 0xb6,
	0xbb, 0x92, 0xfe, 0x39, 0xce, 0xe1, 0x54, 0xc7, 0xa2, 0x77, 0x6b, 0xa4, 0xd2, 0xd5, 0x31, 0x96,
	0xb9, 0x1b, 0x55, 0xf4, 0x93, 0x84, 0x6b, 0x59, 0x41, 0xf0, 0x05, 0x8e, 0xaf, 0x0b, 0x4d, 0x9f,
	0x5b, 0x07, 0xb1, 0x72, 0x89, 0x7f, 0x34, 0x8f, 0x0f, 0x67, 0x59, 0x1d, 0xe5, 0x17, 0x43, 0x0f,
	0x8d, 0x71, 0xe2, 0x23, 0xf8, 0x49, 0xd0, 0x8f, 0xb6, 0xfe, 0x8b, 0x62, 0x7d, 0x1c, 0x4e, 0xd2,
	0x1b, 0x99, 0xd3, 0x32, 0x6e, 0x6e, 0x64, 0xef, 0x87, 0x12, 0x3f, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xb5, 0x87, 0x25, 0xa0, 0xfc, 0x00, 0x00, 0x00,
}
