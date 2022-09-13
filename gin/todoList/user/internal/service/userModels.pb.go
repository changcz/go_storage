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

type UserModel struct {
	// @inject_tag: json:"user_id"
	UserID uint32 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	// @inject_tag: json:"user_name"
	UserName string `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	// @inject_tag: json:"nick_name"
	NickName             string   `protobuf:"bytes,3,opt,name=NickName,proto3" json:"NickName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserModel) Reset()         { *m = UserModel{} }
func (m *UserModel) String() string { return proto.CompactTextString(m) }
func (*UserModel) ProtoMessage()    {}
func (*UserModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab4681ca824f7c7a, []int{0}
}

func (m *UserModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserModel.Unmarshal(m, b)
}
func (m *UserModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserModel.Marshal(b, m, deterministic)
}
func (m *UserModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserModel.Merge(m, src)
}
func (m *UserModel) XXX_Size() int {
	return xxx_messageInfo_UserModel.Size(m)
}
func (m *UserModel) XXX_DiscardUnknown() {
	xxx_messageInfo_UserModel.DiscardUnknown(m)
}

var xxx_messageInfo_UserModel proto.InternalMessageInfo

func (m *UserModel) GetUserID() uint32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *UserModel) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserModel) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func init() {
	proto.RegisterType((*UserModel)(nil), "pb.UserModel")
}

func init() { proto.RegisterFile("taskModels.proto", fileDescriptor_ab4681ca824f7c7a) }

var fileDescriptor_ab4681ca824f7c7a = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2d, 0x4e, 0x2d,
	0xf2, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0x8a, 0xe6, 0xe2, 0x0c, 0x85, 0x89, 0x0b, 0x89, 0x71, 0xb1, 0x81, 0x38, 0x9e, 0x2e, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x50, 0x9e, 0x90, 0x14, 0x17, 0x07, 0x88, 0xe5, 0x97, 0x98,
	0x9b, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x83, 0xe4, 0xfc, 0x32, 0x93, 0xb3,
	0xc1, 0x72, 0xcc, 0x10, 0x39, 0x18, 0xdf, 0x49, 0x3a, 0x4a, 0x52, 0x3f, 0x33, 0xaf, 0x24, 0xb5,
	0x28, 0x2f, 0x31, 0x47, 0xbf, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0xd5, 0x1a, 0x4a, 0x27, 0xb1,
	0x81, 0x1d, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x7e, 0x81, 0x35, 0x98, 0x00, 0x00,
	0x00,
}
