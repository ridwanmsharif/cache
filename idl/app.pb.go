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
	Account
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
	Account *Account `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *GetByTokenResp) Reset()                    { *m = GetByTokenResp{} }
func (m *GetByTokenResp) String() string            { return proto.CompactTextString(m) }
func (*GetByTokenResp) ProtoMessage()               {}
func (*GetByTokenResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetByTokenResp) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

// This implements an Account object
type Account struct {
	MaxCacheKeys int64 `protobuf:"varint,1,opt,name=max_cache_keys,json=maxCacheKeys" json:"max_cache_keys,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Account) GetMaxCacheKeys() int64 {
	if m != nil {
		return m.MaxCacheKeys
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
	proto.RegisterType((*Account)(nil), "rpc.Account")
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

// Client API for Acconuts service

type AcconutsClient interface {
	// Gets the Account info (max number of keys user is allowed)
	GetByToken(ctx context.Context, in *GetByTokenReq, opts ...grpc.CallOption) (*GetByTokenResp, error)
}

type acconutsClient struct {
	cc *grpc.ClientConn
}

func NewAcconutsClient(cc *grpc.ClientConn) AcconutsClient {
	return &acconutsClient{cc}
}

func (c *acconutsClient) GetByToken(ctx context.Context, in *GetByTokenReq, opts ...grpc.CallOption) (*GetByTokenResp, error) {
	out := new(GetByTokenResp)
	err := grpc.Invoke(ctx, "/rpc.Acconuts/GetByToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Acconuts service

type AcconutsServer interface {
	// Gets the Account info (max number of keys user is allowed)
	GetByToken(context.Context, *GetByTokenReq) (*GetByTokenResp, error)
}

func RegisterAcconutsServer(s *grpc.Server, srv AcconutsServer) {
	s.RegisterService(&_Acconuts_serviceDesc, srv)
}

func _Acconuts_GetByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcconutsServer).GetByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Acconuts/GetByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcconutsServer).GetByToken(ctx, req.(*GetByTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Acconuts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Acconuts",
	HandlerType: (*AcconutsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByToken",
			Handler:    _Acconuts_GetByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}

func init() { proto.RegisterFile("app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4f, 0x84, 0x30,
	0x10, 0x15, 0xc9, 0x2e, 0xcb, 0xf0, 0x11, 0x33, 0x7a, 0x20, 0x78, 0xd9, 0x34, 0xae, 0xe1, 0x84,
	0x09, 0x1e, 0xf4, 0xaa, 0x7b, 0xe0, 0xe0, 0x0d, 0xdd, 0x33, 0xa9, 0x4d, 0x13, 0x13, 0x5c, 0xa8,
	0xb4, 0x6b, 0x96, 0x7f, 0x6f, 0x5a, 0xa8, 0x6c, 0xa2, 0xb7, 0xbe, 0x37, 0x6f, 0xde, 0x9b, 0x99,
	0x82, 0x4f, 0x85, 0xc8, 0x45, 0xdf, 0xa9, 0x0e, 0xdd, 0x5e, 0x30, 0xb2, 0x83, 0xd5, 0xab, 0xea,
	0x7a, 0x5e, 0xf1, 0x2f, 0xbc, 0x00, 0xb7, 0xe1, 0x43, 0xe2, 0xac, 0x9d, 0xcc, 0xaf, 0xf4, 0x53,
	0x33, 0xdf, 0xf4, 0x33, 0x39, 0x5f, 0x3b, 0x59, 0x58, 0xe9, 0x27, 0x6e, 0x20, 0xa6, 0x8c, 0x75,
	0x87, 0x56, 0xc9, 0x5a, 0x75, 0x0d, 0x6f, 0x13, 0xd7, 0xc8, 0x23, 0xcb, 0xbe, 0x69, 0x92, 0x04,
	0xe0, 0x4f, 0xb6, 0x52, 0x90, 0x14, 0x96, 0x25, 0x57, 0xff, 0x26, 0x90, 0x6b, 0xf0, 0x4c, 0x4d,
	0x0a, 0x1b, 0xe6, 0xfc, 0x86, 0x91, 0x0d, 0x44, 0x25, 0x57, 0xcf, 0x83, 0xf1, 0xd4, 0xfd, 0x57,
	0xb0, 0x18, 0x43, 0x47, 0x87, 0x11, 0x90, 0x47, 0x88, 0x4f, 0x65, 0x52, 0xe0, 0x2d, 0x78, 0xd3,
	0x3c, 0x46, 0x19, 0x14, 0x61, 0xde, 0x0b, 0x96, 0x3f, 0x8d, 0x5c, 0x65, 0x8b, 0xe4, 0x0e, 0xbc,
	0x89, 0xc3, 0x1b, 0x88, 0xf7, 0xf4, 0x58, 0x33, 0xca, 0x3e, 0x78, 0xdd, 0xf0, 0x41, 0x9a, 0x4e,
	0xb7, 0x0a, 0xf7, 0xf4, 0xb8, 0xd5, 0xe4, 0x0b, 0x1f, 0x64, 0xb1, 0x83, 0x85, 0x01, 0x98, 0xc1,
	0xc2, 0x2c, 0x88, 0x91, 0x71, 0xb6, 0x37, 0x4c, 0xe3, 0x53, 0x28, 0x05, 0x39, 0x43, 0x02, 0x6e,
	0xc9, 0x15, 0x06, 0xa6, 0x30, 0xde, 0x21, 0x0d, 0x67, 0xa0, 0x35, 0xc5, 0x16, 0x56, 0x7a, 0x8e,
	0xf6, 0xa0, 0x24, 0x3e, 0x00, 0xcc, 0xdb, 0x20, 0x5a, 0xe5, 0x7c, 0x85, 0xf4, 0xf2, 0x0f, 0xa7,
	0x4d, 0xde, 0x97, 0xe6, 0x5b, 0xef, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x83, 0xcb, 0x8b, 0xe8,
	0xe3, 0x01, 0x00, 0x00,
}
