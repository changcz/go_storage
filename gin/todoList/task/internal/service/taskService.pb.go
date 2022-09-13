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

type TaskRequest struct {
	TaskID               uint32   `protobuf:"varint,1,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	UserID               uint32   `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Status               uint32   `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Title                string   `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	StartTime            uint32   `protobuf:"varint,6,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime              uint32   `protobuf:"varint,7,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskRequest) Reset()         { *m = TaskRequest{} }
func (m *TaskRequest) String() string { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()    {}
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73398b334343aecd, []int{0}
}

func (m *TaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRequest.Unmarshal(m, b)
}
func (m *TaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRequest.Marshal(b, m, deterministic)
}
func (m *TaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRequest.Merge(m, src)
}
func (m *TaskRequest) XXX_Size() int {
	return xxx_messageInfo_TaskRequest.Size(m)
}
func (m *TaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRequest proto.InternalMessageInfo

func (m *TaskRequest) GetTaskID() uint32 {
	if m != nil {
		return m.TaskID
	}
	return 0
}

func (m *TaskRequest) GetUserID() uint32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *TaskRequest) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TaskRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TaskRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TaskRequest) GetStartTime() uint32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TaskRequest) GetEndTime() uint32 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type TasksDetailResponse struct {
	TaskDetail           []*TaskModel `protobuf:"bytes,1,rep,name=TaskDetail,proto3" json:"TaskDetail,omitempty"`
	Code                 uint32       `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TasksDetailResponse) Reset()         { *m = TasksDetailResponse{} }
func (m *TasksDetailResponse) String() string { return proto.CompactTextString(m) }
func (*TasksDetailResponse) ProtoMessage()    {}
func (*TasksDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73398b334343aecd, []int{1}
}

func (m *TasksDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TasksDetailResponse.Unmarshal(m, b)
}
func (m *TasksDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TasksDetailResponse.Marshal(b, m, deterministic)
}
func (m *TasksDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TasksDetailResponse.Merge(m, src)
}
func (m *TasksDetailResponse) XXX_Size() int {
	return xxx_messageInfo_TasksDetailResponse.Size(m)
}
func (m *TasksDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TasksDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TasksDetailResponse proto.InternalMessageInfo

func (m *TasksDetailResponse) GetTaskDetail() []*TaskModel {
	if m != nil {
		return m.TaskDetail
	}
	return nil
}

func (m *TasksDetailResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type CommonResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73398b334343aecd, []int{2}
}

func (m *CommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResponse.Unmarshal(m, b)
}
func (m *CommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResponse.Marshal(b, m, deterministic)
}
func (m *CommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResponse.Merge(m, src)
}
func (m *CommonResponse) XXX_Size() int {
	return xxx_messageInfo_CommonResponse.Size(m)
}
func (m *CommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResponse proto.InternalMessageInfo

func (m *CommonResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CommonResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *CommonResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*TaskRequest)(nil), "pb.TaskRequest")
	proto.RegisterType((*TasksDetailResponse)(nil), "pb.TasksDetailResponse")
	proto.RegisterType((*CommonResponse)(nil), "pb.CommonResponse")
}

func init() { proto.RegisterFile("taskService.proto", fileDescriptor_73398b334343aecd) }

var fileDescriptor_73398b334343aecd = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0x53, 0x7e, 0xed, 0x10, 0x14, 0x47, 0xa3, 0x23, 0xba, 0x20, 0xac, 0xd8, 0x08, 0x01,
	0xdd, 0xb9, 0xb3, 0x75, 0x81, 0x09, 0x9b, 0x52, 0x12, 0xe3, 0x6e, 0xb0, 0x37, 0xda, 0x50, 0x3a,
	0xb5, 0x73, 0xd1, 0xa7, 0xf3, 0x4d, 0x7c, 0x18, 0x33, 0x7f, 0x62, 0x70, 0xc3, 0xaa, 0x73, 0xbe,
	0xb9, 0x67, 0xee, 0x9d, 0x33, 0x25, 0xc7, 0xc8, 0xe5, 0x6a, 0x0e, 0xe5, 0x47, 0xfa, 0x02, 0xc3,
	0xa2, 0x14, 0x28, 0x68, 0xa5, 0x58, 0x76, 0x3b, 0x0a, 0xcf, 0x44, 0x02, 0x99, 0x34, 0xb4, 0xff,
	0xe5, 0x91, 0x56, 0xcc, 0xe5, 0x2a, 0x82, 0xf7, 0x0d, 0x48, 0xa4, 0x67, 0xa4, 0xa1, 0xe4, 0x34,
	0x64, 0x5e, 0xcf, 0x1b, 0xb4, 0x23, 0xab, 0x14, 0x5f, 0x48, 0x28, 0xa7, 0x21, 0xab, 0x18, 0x6e,
	0x94, 0xe2, 0x73, 0xe4, 0xb8, 0x91, 0xac, 0x6a, 0xb8, 0x51, 0xf4, 0x94, 0xd4, 0xe3, 0x14, 0x33,
	0x60, 0xb5, 0x9e, 0x37, 0xf0, 0x23, 0x23, 0x28, 0x23, 0xcd, 0x40, 0xe4, 0x08, 0x39, 0xb2, 0xba,
	0xe6, 0x4e, 0xd2, 0x2b, 0xe2, 0xcf, 0x91, 0x97, 0x18, 0xa7, 0x6b, 0x60, 0x0d, 0x7d, 0xd4, 0x16,
	0x28, 0xdf, 0x43, 0x9e, 0xe8, 0xbd, 0xa6, 0xde, 0x73, 0xb2, 0xff, 0x44, 0x4e, 0xd4, 0x84, 0x32,
	0x04, 0xe4, 0x69, 0x16, 0x81, 0x2c, 0x44, 0x2e, 0x81, 0x5e, 0x13, 0xa2, 0xb0, 0xa1, 0xcc, 0xeb,
	0x55, 0x07, 0xad, 0x49, 0x7b, 0x58, 0x2c, 0x87, 0xb1, 0x0b, 0x20, 0xfa, 0x53, 0x40, 0x29, 0xa9,
	0x05, 0x22, 0x01, 0x7b, 0x37, 0xbd, 0xee, 0x3f, 0x92, 0xc3, 0x40, 0xac, 0xd7, 0x22, 0xff, 0x3d,
	0xd4, 0x55, 0x79, 0xdb, 0x2a, 0xda, 0x21, 0xd5, 0x99, 0x7c, 0xd5, 0x46, 0x3f, 0x52, 0x4b, 0x55,
	0x15, 0x72, 0xe4, 0x3a, 0x0f, 0x3f, 0xd2, 0xeb, 0xc9, 0xb7, 0x4d, 0xd9, 0xbe, 0x08, 0x1d, 0x9b,
	0xf1, 0x82, 0x12, 0x38, 0x02, 0x3d, 0x72, 0x83, 0xd9, 0x47, 0xe8, 0x52, 0x05, 0x76, 0x9a, 0x5b,
	0xcb, 0xa2, 0x48, 0xf6, 0xb6, 0xdc, 0x92, 0x03, 0xdd, 0xf4, 0x4d, 0x7c, 0xfe, 0x37, 0x9c, 0x3b,
	0xb0, 0x1b, 0xdd, 0xd8, 0x45, 0x97, 0xc1, 0x9e, 0x8d, 0xee, 0x2f, 0x9f, 0x2f, 0x46, 0x69, 0x8e,
	0x50, 0xe6, 0x3c, 0x1b, 0x49, 0x73, 0xc7, 0x3b, 0xfb, 0x5d, 0x36, 0xf4, 0x8f, 0x76, 0xf3, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x83, 0x01, 0xe9, 0xca, 0x93, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskServiceClient interface {
	TaskCreate(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	TaskUpdate(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	TaskShow(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TasksDetailResponse, error)
	TaskDelete(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) TaskCreate(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pb.TaskService/TaskCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskUpdate(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pb.TaskService/TaskUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskShow(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TasksDetailResponse, error) {
	out := new(TasksDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.TaskService/TaskShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskDelete(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pb.TaskService/TaskDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
type TaskServiceServer interface {
	TaskCreate(context.Context, *TaskRequest) (*CommonResponse, error)
	TaskUpdate(context.Context, *TaskRequest) (*CommonResponse, error)
	TaskShow(context.Context, *TaskRequest) (*TasksDetailResponse, error)
	TaskDelete(context.Context, *TaskRequest) (*CommonResponse, error)
}

// UnimplementedTaskServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (*UnimplementedTaskServiceServer) TaskCreate(ctx context.Context, req *TaskRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskCreate not implemented")
}
func (*UnimplementedTaskServiceServer) TaskUpdate(ctx context.Context, req *TaskRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskUpdate not implemented")
}
func (*UnimplementedTaskServiceServer) TaskShow(ctx context.Context, req *TaskRequest) (*TasksDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskShow not implemented")
}
func (*UnimplementedTaskServiceServer) TaskDelete(ctx context.Context, req *TaskRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskDelete not implemented")
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_TaskCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TaskService/TaskCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskCreate(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TaskService/TaskUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskUpdate(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TaskService/TaskShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskShow(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TaskService/TaskDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskDelete(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskCreate",
			Handler:    _TaskService_TaskCreate_Handler,
		},
		{
			MethodName: "TaskUpdate",
			Handler:    _TaskService_TaskUpdate_Handler,
		},
		{
			MethodName: "TaskShow",
			Handler:    _TaskService_TaskShow_Handler,
		},
		{
			MethodName: "TaskDelete",
			Handler:    _TaskService_TaskDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taskService.proto",
}