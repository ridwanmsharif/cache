// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	app.proto

It has these top-level messages:
	StoreReq
	StoreResp
	GetReq
	GetResp
	GetByTokenReq
	GetByTokenResp
*/
package rpc

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

// This request message containing a complete store request
// Stores a key-value pair
type StoreReq struct {
	Key           string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Val           []byte `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	AccountsToken string `protobuf:"bytes,3,opt,name=accounts_token,json=accountsToken" json:"accounts_token,omitempty"`
}

func (m *StoreReq) Reset()                    { *m = StoreReq{} }
func (m *StoreReq) String() string            { return proto.CompactTextString(m) }
func (*StoreReq) ProtoMessage()               {}
func (*StoreReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StoreReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StoreReq) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *StoreReq) GetAccountsToken() string {
	if m != nil {
		return m.AccountsToken
	}
	return ""
}

// This response message responds to a store request
type StoreResp struct {
}

func (m *StoreResp) Reset()                    { *m = StoreResp{} }
func (m *StoreResp) String() string            { return proto.CompactTextString(m) }
func (*StoreResp) ProtoMessage()               {}
func (*StoreResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// This request message containing a complete get request
// Get's a value from a key-value pair after providing the key
type GetReq struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetReq) Reset()                    { *m = GetReq{} }
func (m *GetReq) String() string            { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()               {}
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// This response message containing a response to a get request
// Responds to a get request with the value
type GetResp struct {
	Val []byte `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (m *GetResp) Reset()                    { *m = GetResp{} }
func (m *GetResp) String() string            { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()               {}
func (*GetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetResp) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

// This request contains the acount user identifier
type GetByTokenReq struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetByTokenReq) Reset()                    { *m = GetByTokenReq{} }
func (m *GetByTokenReq) String() string            { return proto.CompactTextString(m) }
func (*GetByTokenReq) ProtoMessage()               {}
func (*GetByTokenReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetByTokenReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// This response message respongs to a GetByToken request with an account
type GetByTokenResp struct {
	Allowed int64 `protobuf:"varint,1,opt,name=allowed" json:"allowed,omitempty"`
}

func (m *GetByTokenResp) Reset()                    { *m = GetByTokenResp{} }
func (m *GetByTokenResp) String() string            { return proto.CompactTextString(m) }
func (*GetByTokenResp) ProtoMessage()               {}
func (*GetByTokenResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetByTokenResp) GetAllowed() int64 {
	if m != nil {
		return m.Allowed
	}
	return 0
}

func init() {
	proto.RegisterType((*StoreReq)(nil), "rpc.StoreReq")
	proto.RegisterType((*StoreResp)(nil), "rpc.StoreResp")
	proto.RegisterType((*GetReq)(nil), "rpc.GetReq")
	proto.RegisterType((*GetResp)(nil), "rpc.GetResp")
	proto.RegisterType((*GetByTokenReq)(nil), "rpc.GetByTokenReq")
	proto.RegisterType((*GetByTokenResp)(nil), "rpc.GetByTokenResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cache service

type CacheClient interface {
	// Creates and stores a KV pair
	Store(ctx context.Context, in *StoreReq, opts ...grpc.CallOption) (*StoreResp, error)
	// Fetches a KV pair
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
}

type cacheClient struct {
	cc *grpc.ClientConn
}

func NewCacheClient(cc *grpc.ClientConn) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) Store(ctx context.Context, in *StoreReq, opts ...grpc.CallOption) (*StoreResp, error) {
	out := new(StoreResp)
	err := grpc.Invoke(ctx, "/rpc.Cache/Store", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := grpc.Invoke(ctx, "/rpc.Cache/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cache service

type CacheServer interface {
	// Creates and stores a KV pair
	Store(context.Context, *StoreReq) (*StoreResp, error)
	// Fetches a KV pair
	Get(context.Context, *GetReq) (*GetResp, error)
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cache/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Store(ctx, req.(*StoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cache/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Cache_Store_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Cache_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}

// Client API for Accounts service

type AccountsClient interface {
	// Gets the Account info (max number of keys user is allowed)
	GetByToken(ctx context.Context, in *GetByTokenReq, opts ...grpc.CallOption) (*GetByTokenResp, error)
}

type accountsClient struct {
	cc *grpc.ClientConn
}

func NewAccountsClient(cc *grpc.ClientConn) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) GetByToken(ctx context.Context, in *GetByTokenReq, opts ...grpc.CallOption) (*GetByTokenResp, error) {
	out := new(GetByTokenResp)
	err := grpc.Invoke(ctx, "/rpc.Accounts/GetByToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Accounts service

type AccountsServer interface {
	// Gets the Account info (max number of keys user is allowed)
	GetByToken(context.Context, *GetByTokenReq) (*GetByTokenResp, error)
}

func RegisterAccountsServer(s *grpc.Server, srv AccountsServer) {
	s.RegisterService(&_Accounts_serviceDesc, srv)
}

func _Accounts_GetByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Accounts/GetByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetByToken(ctx, req.(*GetByTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByToken",
			Handler:    _Accounts_GetByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}

func init() { proto.RegisterFile("app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x25, 0x44, 0x69, 0x9b, 0x6b, 0x13, 0xa1, 0x83, 0x21, 0x0a, 0x4b, 0x65, 0xa9, 0x52, 0xc4,
	0x90, 0xa1, 0x0c, 0xcc, 0xd0, 0x21, 0x7b, 0xa0, 0x33, 0x32, 0xe6, 0x24, 0xa4, 0x46, 0xf5, 0x11,
	0x1b, 0x50, 0xff, 0x1e, 0xd9, 0xa9, 0x49, 0x25, 0xba, 0xdd, 0x7b, 0xf7, 0xfc, 0xde, 0xdd, 0x19,
	0x52, 0xc9, 0x5c, 0x73, 0xaf, 0xad, 0xc6, 0xb8, 0x67, 0x25, 0xb6, 0x30, 0x7b, 0xb6, 0xba, 0xa7,
	0x96, 0x3e, 0xf1, 0x0a, 0xe2, 0x1d, 0x1d, 0x8a, 0x68, 0x19, 0x55, 0x69, 0xeb, 0x4a, 0xc7, 0x7c,
	0xcb, 0xae, 0xb8, 0x5c, 0x46, 0xd5, 0xa2, 0x75, 0x25, 0xae, 0x20, 0x97, 0x4a, 0xe9, 0xaf, 0xbd,
	0x35, 0xaf, 0x56, 0xef, 0x68, 0x5f, 0xc4, 0x5e, 0x9e, 0x05, 0xf6, 0xc5, 0x91, 0x62, 0x0e, 0xe9,
	0xd1, 0xd6, 0xb0, 0x28, 0x61, 0xd2, 0x90, 0x3d, 0x9b, 0x20, 0x6e, 0x61, 0xea, 0x7b, 0x86, 0x43,
	0x58, 0xf4, 0x17, 0x26, 0x56, 0x90, 0x35, 0x64, 0x9f, 0x0e, 0xde, 0xd3, 0xbd, 0xbf, 0x81, 0x64,
	0x08, 0x1d, 0x1c, 0x06, 0x20, 0xee, 0x20, 0x3f, 0x95, 0x19, 0xc6, 0x02, 0xa6, 0xb2, 0xeb, 0xf4,
	0x0f, 0xbd, 0x7b, 0x65, 0xdc, 0x06, 0xb8, 0xde, 0x42, 0xb2, 0x91, 0xea, 0x83, 0xb0, 0x82, 0xc4,
	0x4f, 0x88, 0x59, 0xdd, 0xb3, 0xaa, 0xc3, 0x11, 0xca, 0xfc, 0x14, 0x1a, 0x16, 0x17, 0x28, 0x20,
	0x6e, 0xc8, 0xe2, 0xdc, 0x37, 0x86, 0x45, 0xca, 0xc5, 0x08, 0x9c, 0x66, 0xbd, 0x81, 0xd9, 0xe3,
	0xf1, 0x00, 0xf8, 0x00, 0x30, 0x8e, 0x83, 0x18, 0x94, 0xe3, 0x1a, 0xe5, 0xf5, 0x3f, 0xce, 0x99,
	0xbc, 0x4d, 0xfc, 0xbf, 0xdc, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x00, 0x80, 0x75, 0x8d, 0xa4,
	0x01, 0x00, 0x00,
}
