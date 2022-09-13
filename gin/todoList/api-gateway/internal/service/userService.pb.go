// Code generated by protoc-gen-go. DO NOT EDIT.
// source: taskService.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UserRequest struct {
	// @inject_tag: json:"nick_name" form:"nick_name" uri:"nick_name"
	NickName string `protobuf:"bytes,1,opt,name=NickName,proto3" json:"NickName,omitempty"`
	// @inject_tag: json:"user_name" form:"user_name" uri:"user_name"
	UserName string `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	// @inject_tag: json:"password" form:"password" uri:"password"
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	// @inject_tag: json:"password_confirm" form:"password_confirm" uri:"password_confirm"
	PasswordConfirm      string   `protobuf:"bytes,4,opt,name=PasswordConfirm,proto3" json:"PasswordConfirm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_874b4b11a4ddc7c4, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *UserRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserRequest) GetPasswordConfirm() string {
	if m != nil {
		return m.PasswordConfirm
	}
	return ""
}

type UserDetailResponse struct {
	UserDetail           *UserModel `protobuf:"bytes,1,opt,name=UserDetail,proto3" json:"UserDetail,omitempty"`
	Code                 uint32     `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UserDetailResponse) Reset()         { *m = UserDetailResponse{} }
func (m *UserDetailResponse) String() string { return proto.CompactTextString(m) }
func (*UserDetailResponse) ProtoMessage()    {}
func (*UserDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_874b4b11a4ddc7c4, []int{1}
}

func (m *UserDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDetailResponse.Unmarshal(m, b)
}
func (m *UserDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDetailResponse.Marshal(b, m, deterministic)
}
func (m *UserDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDetailResponse.Merge(m, src)
}
func (m *UserDetailResponse) XXX_Size() int {
	return xxx_messageInfo_UserDetailResponse.Size(m)
}
func (m *UserDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserDetailResponse proto.InternalMessageInfo

func (m *UserDetailResponse) GetUserDetail() *UserModel {
	if m != nil {
		return m.UserDetail
	}
	return nil
}

func (m *UserDetailResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "pb.UserRequest")
	proto.RegisterType((*UserDetailResponse)(nil), "pb.UserDetailResponse")
}

func init() { proto.RegisterFile("taskService.proto", fileDescriptor_874b4b11a4ddc7c4) }

var fileDescriptor_874b4b11a4ddc7c4 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0x66, 0x6b, 0x11, 0x3b, 0xb5, 0x54, 0xe7, 0x20, 0xeb, 0x7a, 0x91, 0x9e, 0x7a, 0x71, 0x0b,
	0x55, 0xf0, 0xe0, 0xcd, 0x7a, 0xd4, 0x22, 0x91, 0x22, 0x78, 0xdb, 0xed, 0x8e, 0x25, 0xb8, 0x4d,
	0xd6, 0x4c, 0xaa, 0x17, 0xdf, 0xc0, 0x97, 0x96, 0x24, 0x1b, 0x95, 0x9e, 0x7a, 0xca, 0x7c, 0x3f,
	0x19, 0xbe, 0x99, 0x81, 0xe3, 0x0d, 0x93, 0x79, 0x22, 0xf3, 0x21, 0x97, 0x94, 0x37, 0x46, 0x5b,
	0x8d, 0x9d, 0xa6, 0xcc, 0x8e, 0x1c, 0xfd, 0xa0, 0x2b, 0xaa, 0x39, 0xb0, 0xa3, 0xef, 0x04, 0xfa,
	0x0b, 0x26, 0x23, 0xe8, 0x7d, 0x43, 0x6c, 0x31, 0x83, 0x83, 0xb9, 0x5c, 0xbe, 0xcd, 0x8b, 0x35,
	0xa5, 0xc9, 0x79, 0x32, 0xee, 0x89, 0x5f, 0xec, 0x34, 0x67, 0xf5, 0x5a, 0x27, 0x68, 0x11, 0x3b,
	0xed, 0xb1, 0x60, 0xfe, 0xd4, 0xa6, 0x4a, 0xf7, 0x82, 0x16, 0x31, 0x8e, 0x61, 0x18, 0xeb, 0x99,
	0x56, 0xaf, 0xd2, 0xac, 0xd3, 0xae, 0xb7, 0x6c, 0xd3, 0xa3, 0x67, 0x40, 0xd7, 0xf1, 0x8e, 0x6c,
	0x21, 0x6b, 0x41, 0xdc, 0x68, 0xc5, 0x84, 0x17, 0x00, 0x7f, 0xac, 0x4f, 0xd5, 0x9f, 0x0e, 0xf2,
	0xa6, 0xcc, 0x17, 0x71, 0x1a, 0xf1, 0xcf, 0x80, 0x08, 0xdd, 0x99, 0xae, 0x42, 0xc4, 0x81, 0xf0,
	0xf5, 0xf4, 0x2b, 0x4c, 0xd9, 0x6e, 0x04, 0xaf, 0xa0, 0xe7, 0xe0, 0xbd, 0x5e, 0x49, 0x85, 0xc3,
	0xd8, 0xaa, 0xdd, 0x41, 0x76, 0x12, 0x89, 0xad, 0x1c, 0xd7, 0x70, 0x18, 0x6c, 0x2b, 0xc9, 0x96,
	0xcc, 0xce, 0x1f, 0x6f, 0xcf, 0x5e, 0x4e, 0x27, 0x52, 0x59, 0x32, 0xaa, 0xa8, 0x27, 0x1c, 0x32,
	0xdc, 0xb4, 0x6f, 0xb9, 0xef, 0x0f, 0x71, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x98, 0xe6, 0x9e,
	0xa8, 0xb3, 0x01, 0x00, 0x00,
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
	UserLogin(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserDetailResponse, error)
	UserRegister(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserDetailResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserDetailResponse, error) {
	out := new(UserDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserRegister(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserDetailResponse, error) {
	out := new(UserDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	UserLogin(context.Context, *UserRequest) (*UserDetailResponse, error)
	UserRegister(context.Context, *UserRequest) (*UserDetailResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) UserLogin(ctx context.Context, req *UserRequest) (*UserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (*UnimplementedUserServiceServer) UserRegister(ctx context.Context, req *UserRequest) (*UserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRegister(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _UserService_UserRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taskService.proto",
}
