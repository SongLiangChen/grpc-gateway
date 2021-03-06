// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user-service/user-service.proto

package user_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	base "grpc-gateway/proto/base"
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

type UserInfoReq struct {
	CommonParam          *base.CommonParam `protobuf:"bytes,1,opt,name=commonParam,proto3" json:"commonParam,omitempty"`
	Session              *base.Session     `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UserInfoReq) Reset()         { *m = UserInfoReq{} }
func (m *UserInfoReq) String() string { return proto.CompactTextString(m) }
func (*UserInfoReq) ProtoMessage()    {}
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e882b7622885a4db, []int{0}
}

func (m *UserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoReq.Unmarshal(m, b)
}
func (m *UserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoReq.Marshal(b, m, deterministic)
}
func (m *UserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoReq.Merge(m, src)
}
func (m *UserInfoReq) XXX_Size() int {
	return xxx_messageInfo_UserInfoReq.Size(m)
}
func (m *UserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoReq proto.InternalMessageInfo

func (m *UserInfoReq) GetCommonParam() *base.CommonParam {
	if m != nil {
		return m.CommonParam
	}
	return nil
}

func (m *UserInfoReq) GetSession() *base.Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type UserInfoResp struct {
	Id                   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64         `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	NickName             string        `protobuf:"bytes,3,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Session              *base.Session `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserInfoResp) Reset()         { *m = UserInfoResp{} }
func (m *UserInfoResp) String() string { return proto.CompactTextString(m) }
func (*UserInfoResp) ProtoMessage()    {}
func (*UserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e882b7622885a4db, []int{1}
}

func (m *UserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResp.Unmarshal(m, b)
}
func (m *UserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResp.Marshal(b, m, deterministic)
}
func (m *UserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResp.Merge(m, src)
}
func (m *UserInfoResp) XXX_Size() int {
	return xxx_messageInfo_UserInfoResp.Size(m)
}
func (m *UserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResp proto.InternalMessageInfo

func (m *UserInfoResp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoResp) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfoResp) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *UserInfoResp) GetSession() *base.Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInfoReq)(nil), "user.service.UserInfoReq")
	proto.RegisterType((*UserInfoResp)(nil), "user.service.UserInfoResp")
}

func init() { proto.RegisterFile("user-service/user-service.proto", fileDescriptor_e882b7622885a4db) }

var fileDescriptor_e882b7622885a4db = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x55, 0xd2, 0x4f, 0xfd, 0xc0, 0x29, 0x15, 0x58, 0x02, 0x85, 0x80, 0x44, 0x95, 0x85, 0x0a,
	0xa9, 0x31, 0xb4, 0x1b, 0x23, 0x0c, 0xa8, 0x0b, 0x42, 0xae, 0x58, 0xd8, 0xdc, 0xe4, 0x12, 0x59,
	0x25, 0xb6, 0xb1, 0x0d, 0x08, 0xb1, 0xf1, 0x0a, 0x3c, 0x1a, 0xaf, 0xc0, 0x83, 0xa0, 0xd8, 0x2d,
	0xf5, 0x02, 0x4b, 0xa4, 0x7b, 0x7e, 0x72, 0x7c, 0xcf, 0x45, 0x47, 0x4f, 0x06, 0xf4, 0xc8, 0x80,
	0x7e, 0xe6, 0x25, 0x90, 0x70, 0x28, 0x94, 0x96, 0x56, 0xe2, 0x5e, 0x8b, 0x15, 0x4b, 0x2c, 0x3b,
	0xac, 0xa5, 0xac, 0x1f, 0x80, 0x30, 0xc5, 0x09, 0x13, 0x42, 0x5a, 0x66, 0xb9, 0x14, 0xc6, 0x6b,
	0xb3, 0xbc, 0xd6, 0xaa, 0x1c, 0xd5, 0xcc, 0xc2, 0x0b, 0x7b, 0x25, 0x0e, 0x23, 0x73, 0x66, 0xc0,
	0x7d, 0xbc, 0x26, 0x5f, 0xa0, 0xe4, 0xd6, 0x80, 0x9e, 0x8a, 0x7b, 0x49, 0xe1, 0x11, 0x4f, 0x50,
	0x52, 0xca, 0xa6, 0x91, 0xe2, 0x86, 0x69, 0xd6, 0xa4, 0xd1, 0x20, 0x1a, 0x26, 0xe3, 0x9d, 0xc2,
	0x19, 0x2e, 0xd7, 0x04, 0x0d, 0x55, 0xf8, 0x18, 0xfd, 0x37, 0x60, 0x0c, 0x97, 0x22, 0x8d, 0x9d,
	0x61, 0xcb, 0x1b, 0x66, 0x1e, 0xa4, 0x2b, 0x36, 0x7f, 0x43, 0xbd, 0x75, 0x98, 0x51, 0xb8, 0x8f,
	0x62, 0x5e, 0xb9, 0x90, 0x0e, 0x8d, 0x79, 0x85, 0xf7, 0x50, 0xb7, 0x5d, 0x6f, 0x5a, 0xb9, 0xff,
	0x74, 0xe8, 0x72, 0xc2, 0x19, 0xda, 0x10, 0xbc, 0x5c, 0x5c, 0xb3, 0x06, 0xd2, 0xce, 0x20, 0x1a,
	0x6e, 0xd2, 0x9f, 0x39, 0x0c, 0xff, 0xf7, 0x57, 0xf8, 0x58, 0xf9, 0x4d, 0x67, 0xbe, 0x3a, 0xcc,
	0x50, 0x72, 0x05, 0x76, 0xf5, 0x1c, 0xbc, 0x5f, 0x84, 0xc5, 0x16, 0x41, 0x27, 0x59, 0xf6, 0x1b,
	0x65, 0x54, 0x7e, 0xf0, 0xfe, 0xf9, 0xf5, 0x11, 0xef, 0xe6, 0xdb, 0xe4, 0xf4, 0xcc, 0x9d, 0x8b,
	0xd4, 0x60, 0x5b, 0xf6, 0x3c, 0x3a, 0xb9, 0xe8, 0xdf, 0xf5, 0xc2, 0x0b, 0xce, 0xbb, 0xae, 0xf2,
	0xc9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0xdc, 0x41, 0x7c, 0xe5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, "/user.service.UserService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUserInfo(ctx context.Context, req *UserInfoReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.UserService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.service.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _UserService_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user-service/user-service.proto",
}
