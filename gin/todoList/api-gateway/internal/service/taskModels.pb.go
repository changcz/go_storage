// Code generated by protoc-gen-go. DO NOT EDIT.
// source: taskModels.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type TaskModel struct {
	// @inject_tag: json:"favorite_id"
	TaskID uint32 `protobuf:"varint,1,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	// @inject_tag: json:"user_id"
	UserID uint32 `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	// @inject_tag: json:"status"
	Status uint32 `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
	// @inject_tag: json:"title"
	Title string `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	// @inject_tag: json:"content"
	Content string `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	// @inject_tag: json:"start_time"
	StartTime uint32 `protobuf:"varint,6,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	// @inject_tag: json:"end_time"
	EndTime              uint32   `protobuf:"varint,7,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskModel) Reset()         { *m = TaskModel{} }
func (m *TaskModel) String() string { return proto.CompactTextString(m) }
func (*TaskModel) ProtoMessage()    {}
func (*TaskModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a262ed70f83c393, []int{0}
}

func (m *TaskModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskModel.Unmarshal(m, b)
}
func (m *TaskModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskModel.Marshal(b, m, deterministic)
}
func (m *TaskModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskModel.Merge(m, src)
}
func (m *TaskModel) XXX_Size() int {
	return xxx_messageInfo_TaskModel.Size(m)
}
func (m *TaskModel) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskModel.DiscardUnknown(m)
}

var xxx_messageInfo_TaskModel proto.InternalMessageInfo

func (m *TaskModel) GetTaskID() uint32 {
	if m != nil {
		return m.TaskID
	}
	return 0
}

func (m *TaskModel) GetUserID() uint32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *TaskModel) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TaskModel) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TaskModel) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TaskModel) GetStartTime() uint32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TaskModel) GetEndTime() uint32 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func init() {
	proto.RegisterType((*TaskModel)(nil), "pb.TaskModel")
}

func init() { proto.RegisterFile("taskModels.proto", fileDescriptor_5a262ed70f83c393) }

var fileDescriptor_5a262ed70f83c393 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x49, 0x2c, 0xce,
	0xf6, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0xda, 0xcd, 0xc8, 0xc5, 0x19, 0x02, 0x93, 0x10, 0x12, 0xe3, 0x62, 0x03, 0x71, 0x3c, 0x5d,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0xa0, 0x3c, 0x90, 0x78, 0x68, 0x71, 0x6a, 0x91, 0xa7,
	0x8b, 0x04, 0x13, 0x44, 0x1c, 0xc2, 0x03, 0x89, 0x07, 0x97, 0x24, 0x96, 0x94, 0x16, 0x4b, 0x30,
	0x43, 0xc4, 0x21, 0x3c, 0x21, 0x11, 0x2e, 0xd6, 0x90, 0xcc, 0x92, 0x9c, 0x54, 0x09, 0x16, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x82, 0x8b, 0xdd, 0x39, 0x3f, 0xaf, 0x24, 0x35, 0xaf,
	0x44, 0x82, 0x15, 0x2c, 0x0e, 0xe3, 0x0a, 0xc9, 0x70, 0x71, 0x06, 0x97, 0x24, 0x16, 0x95, 0x84,
	0x64, 0xe6, 0xa6, 0x4a, 0xb0, 0x81, 0x8d, 0x42, 0x08, 0x80, 0xf4, 0xb9, 0xe6, 0xa5, 0x80, 0xe5,
	0xd8, 0xc1, 0x72, 0x30, 0xae, 0x93, 0x74, 0x94, 0xa4, 0x7e, 0x66, 0x5e, 0x49, 0x6a, 0x51, 0x5e,
	0x62, 0x8e, 0x7e, 0x71, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x35, 0x94, 0x4e, 0x62, 0x03, 0xfb,
	0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xf3, 0x18, 0xa8, 0xf9, 0x00, 0x00, 0x00,
}
