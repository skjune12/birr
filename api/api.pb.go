// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PingMessage struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (m *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(m, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

type AddFileMessage struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Filename             string   `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Content              []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFileMessage) Reset()         { *m = AddFileMessage{} }
func (m *AddFileMessage) String() string { return proto.CompactTextString(m) }
func (*AddFileMessage) ProtoMessage()    {}
func (*AddFileMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *AddFileMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFileMessage.Unmarshal(m, b)
}
func (m *AddFileMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFileMessage.Marshal(b, m, deterministic)
}
func (m *AddFileMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFileMessage.Merge(m, src)
}
func (m *AddFileMessage) XXX_Size() int {
	return xxx_messageInfo_AddFileMessage.Size(m)
}
func (m *AddFileMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFileMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddFileMessage proto.InternalMessageInfo

func (m *AddFileMessage) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AddFileMessage) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *AddFileMessage) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type HashValue struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashValue) Reset()         { *m = HashValue{} }
func (m *HashValue) String() string { return proto.CompactTextString(m) }
func (*HashValue) ProtoMessage()    {}
func (*HashValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *HashValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashValue.Unmarshal(m, b)
}
func (m *HashValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashValue.Marshal(b, m, deterministic)
}
func (m *HashValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashValue.Merge(m, src)
}
func (m *HashValue) XXX_Size() int {
	return xxx_messageInfo_HashValue.Size(m)
}
func (m *HashValue) XXX_DiscardUnknown() {
	xxx_messageInfo_HashValue.DiscardUnknown(m)
}

var xxx_messageInfo_HashValue proto.InternalMessageInfo

func (m *HashValue) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *HashValue) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetFileMessage struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileMessage) Reset()         { *m = GetFileMessage{} }
func (m *GetFileMessage) String() string { return proto.CompactTextString(m) }
func (*GetFileMessage) ProtoMessage()    {}
func (*GetFileMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *GetFileMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileMessage.Unmarshal(m, b)
}
func (m *GetFileMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileMessage.Marshal(b, m, deterministic)
}
func (m *GetFileMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileMessage.Merge(m, src)
}
func (m *GetFileMessage) XXX_Size() int {
	return xxx_messageInfo_GetFileMessage.Size(m)
}
func (m *GetFileMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileMessage proto.InternalMessageInfo

func (m *GetFileMessage) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type ContentMessage struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentMessage) Reset()         { *m = ContentMessage{} }
func (m *ContentMessage) String() string { return proto.CompactTextString(m) }
func (*ContentMessage) ProtoMessage()    {}
func (*ContentMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *ContentMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentMessage.Unmarshal(m, b)
}
func (m *ContentMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentMessage.Marshal(b, m, deterministic)
}
func (m *ContentMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentMessage.Merge(m, src)
}
func (m *ContentMessage) XXX_Size() int {
	return xxx_messageInfo_ContentMessage.Size(m)
}
func (m *ContentMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ContentMessage proto.InternalMessageInfo

func (m *ContentMessage) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *ContentMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*PingMessage)(nil), "api.PingMessage")
	proto.RegisterType((*AddFileMessage)(nil), "api.AddFileMessage")
	proto.RegisterType((*HashValue)(nil), "api.HashValue")
	proto.RegisterType((*GetFileMessage)(nil), "api.GetFileMessage")
	proto.RegisterType((*ContentMessage)(nil), "api.ContentMessage")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x52, 0xe4, 0xe2, 0x0e, 0xc8,
	0xcc, 0x4b, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe2, 0x62, 0x29, 0xc8, 0xcc,
	0x4b, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xa2, 0xb8, 0xf8, 0x1c, 0x53,
	0x52, 0xdc, 0x32, 0x73, 0x52, 0x91, 0x54, 0x95, 0x54, 0x16, 0xa4, 0xc2, 0x54, 0x81, 0xd8, 0x42,
	0x52, 0x5c, 0x1c, 0x69, 0x99, 0x39, 0xa9, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x71, 0x38,
	0x5f, 0x48, 0x82, 0x8b, 0x3d, 0x39, 0x3f, 0xaf, 0x24, 0x35, 0xaf, 0x44, 0x82, 0x59, 0x81, 0x51,
	0x83, 0x27, 0x08, 0xc6, 0x55, 0xb2, 0xe6, 0xe2, 0xf4, 0x48, 0x2c, 0xce, 0x08, 0x4b, 0xcc, 0x29,
	0x45, 0x35, 0x82, 0x11, 0xcd, 0x08, 0x21, 0x2e, 0x96, 0x8c, 0xc4, 0xe2, 0x0c, 0xa8, 0xd1, 0x60,
	0xb6, 0x92, 0x0a, 0x17, 0x9f, 0x7b, 0x6a, 0x09, 0x9a, 0xc3, 0xc0, 0xaa, 0x18, 0x91, 0x54, 0xd9,
	0x71, 0xf1, 0x39, 0x43, 0x6c, 0xc3, 0xa3, 0x0a, 0xd9, 0x89, 0x10, 0x2b, 0x60, 0x5c, 0xa3, 0xe9,
	0x8c, 0x5c, 0x2c, 0x4e, 0x99, 0x45, 0x45, 0x42, 0x5a, 0x5c, 0x2c, 0xa0, 0xa0, 0x12, 0x12, 0xd0,
	0x03, 0x85, 0x21, 0x52, 0xa8, 0x49, 0x61, 0x88, 0x08, 0xe9, 0x71, 0xb1, 0x43, 0xc3, 0x4c, 0x48,
	0x18, 0x2c, 0x89, 0x1a, 0x82, 0x52, 0x7c, 0x60, 0x41, 0x84, 0xd7, 0x8d, 0xb9, 0xd8, 0xa1, 0x5e,
	0x81, 0xaa, 0x47, 0xf5, 0x98, 0x14, 0x44, 0x10, 0xd5, 0x1f, 0x49, 0x6c, 0xe0, 0x78, 0x34, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xba, 0x6e, 0x38, 0xb4, 0xd4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BirrClient is the client API for Birr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BirrClient interface {
	Ping(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PingMessage, error)
	AddFile(ctx context.Context, in *AddFileMessage, opts ...grpc.CallOption) (*HashValue, error)
	GetFile(ctx context.Context, in *GetFileMessage, opts ...grpc.CallOption) (*ContentMessage, error)
}

type birrClient struct {
	cc *grpc.ClientConn
}

func NewBirrClient(cc *grpc.ClientConn) BirrClient {
	return &birrClient{cc}
}

func (c *birrClient) Ping(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PingMessage, error) {
	out := new(PingMessage)
	err := c.cc.Invoke(ctx, "/api.Birr/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birrClient) AddFile(ctx context.Context, in *AddFileMessage, opts ...grpc.CallOption) (*HashValue, error) {
	out := new(HashValue)
	err := c.cc.Invoke(ctx, "/api.Birr/AddFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birrClient) GetFile(ctx context.Context, in *GetFileMessage, opts ...grpc.CallOption) (*ContentMessage, error) {
	out := new(ContentMessage)
	err := c.cc.Invoke(ctx, "/api.Birr/GetFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BirrServer is the server API for Birr service.
type BirrServer interface {
	Ping(context.Context, *PingMessage) (*PingMessage, error)
	AddFile(context.Context, *AddFileMessage) (*HashValue, error)
	GetFile(context.Context, *GetFileMessage) (*ContentMessage, error)
}

func RegisterBirrServer(s *grpc.Server, srv BirrServer) {
	s.RegisterService(&_Birr_serviceDesc, srv)
}

func _Birr_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirrServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Birr/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirrServer).Ping(ctx, req.(*PingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Birr_AddFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFileMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirrServer).AddFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Birr/AddFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirrServer).AddFile(ctx, req.(*AddFileMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Birr_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirrServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Birr/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirrServer).GetFile(ctx, req.(*GetFileMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Birr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Birr",
	HandlerType: (*BirrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Birr_Ping_Handler,
		},
		{
			MethodName: "AddFile",
			Handler:    _Birr_AddFile_Handler,
		},
		{
			MethodName: "GetFile",
			Handler:    _Birr_GetFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
