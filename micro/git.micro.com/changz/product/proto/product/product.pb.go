// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/product/product.proto

package git_micro_com_changz_product

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

type ProductInfo struct {
	Id                   int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductName          string          `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	ProductSku           string          `protobuf:"bytes,3,opt,name=product_sku,json=productSku,proto3" json:"product_sku,omitempty"`
	ProductPrice         float64         `protobuf:"fixed64,4,opt,name=product_price,json=productPrice,proto3" json:"product_price,omitempty"`
	ProductDescription   string          `protobuf:"bytes,5,opt,name=product_description,json=productDescription,proto3" json:"product_description,omitempty"`
	ProductCategoryId    int64           `protobuf:"varint,6,opt,name=product_category_id,json=productCategoryId,proto3" json:"product_category_id,omitempty"`
	ProductImage         []*ProductImage `protobuf:"bytes,7,rep,name=product_image,json=productImage,proto3" json:"product_image,omitempty"`
	ProductSize          []*ProductSize  `protobuf:"bytes,8,rep,name=product_size,json=productSize,proto3" json:"product_size,omitempty"`
	ProductSeo           *ProductSeo     `protobuf:"bytes,9,opt,name=product_seo,json=productSeo,proto3" json:"product_seo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProductInfo) Reset()         { *m = ProductInfo{} }
func (m *ProductInfo) String() string { return proto.CompactTextString(m) }
func (*ProductInfo) ProtoMessage()    {}
func (*ProductInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fbc13de7c38f78, []int{0}
}

func (m *ProductInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductInfo.Unmarshal(m, b)
}
func (m *ProductInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductInfo.Marshal(b, m, deterministic)
}
func (m *ProductInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductInfo.Merge(m, src)
}
func (m *ProductInfo) XXX_Size() int {
	return xxx_messageInfo_ProductInfo.Size(m)
}
func (m *ProductInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProductInfo proto.InternalMessageInfo

func (m *ProductInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductInfo) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *ProductInfo) GetProductSku() string {
	if m != nil {
		return m.ProductSku
	}
	return ""
}

func (m *ProductInfo) GetProductPrice() float64 {
	if m != nil {
		return m.ProductPrice
	}
	return 0
}

func (m *ProductInfo) GetProductDescription() string {
	if m != nil {
		return m.ProductDescription
	}
	return ""
}

func (m *ProductInfo) GetProductCategoryId() int64 {
	if m != nil {
		return m.ProductCategoryId
	}
	return 0
}

func (m *ProductInfo) GetProductImage() []*ProductImage {
	if m != nil {
		return m.ProductImage
	}
	return nil
}

func (m *ProductInfo) GetProductSize() []*ProductSize {
	if m != nil {
		return m.ProductSize
	}
	return nil
}

func (m *ProductInfo) GetProductSeo() *ProductSeo {
	if m != nil {
		return m.ProductSeo
	}
	return nil
}

type ProductImage struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ImageName            string   `protobuf:"bytes,2,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	ImageCode            string   `protobuf:"bytes,3,opt,name=image_code,json=imageCode,proto3" json:"image_code,omitempty"`
	ImageUrl             string   `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductImage) Reset()         { *m = ProductImage{} }
func (m *ProductImage) String() string { return proto.CompactTextString(m) }
func (*ProductImage) ProtoMessage()    {}
func (*ProductImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fbc13de7c38f78, []int{1}
}

func (m *ProductImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductImage.Unmarshal(m, b)
}
func (m *ProductImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductImage.Marshal(b, m, deterministic)
}
func (m *ProductImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductImage.Merge(m, src)
}
func (m *ProductImage) XXX_Size() int {
	return xxx_messageInfo_ProductImage.Size(m)
}
func (m *ProductImage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductImage.DiscardUnknown(m)
}

var xxx_messageInfo_ProductImage proto.InternalMessageInfo

func (m *ProductImage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductImage) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *ProductImage) GetImageCode() string {
	if m != nil {
		return m.ImageCode
	}
	return ""
}

func (m *ProductImage) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type ProductSize struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SizeName             string   `protobuf:"bytes,2,opt,name=size_name,json=sizeName,proto3" json:"size_name,omitempty"`
	SizeCode             string   `protobuf:"bytes,3,opt,name=size_code,json=sizeCode,proto3" json:"size_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductSize) Reset()         { *m = ProductSize{} }
func (m *ProductSize) String() string { return proto.CompactTextString(m) }
func (*ProductSize) ProtoMessage()    {}
func (*ProductSize) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fbc13de7c38f78, []int{2}
}

func (m *ProductSize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSize.Unmarshal(m, b)
}
func (m *ProductSize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSize.Marshal(b, m, deterministic)
}
func (m *ProductSize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSize.Merge(m, src)
}
func (m *ProductSize) XXX_Size() int {
	return xxx_messageInfo_ProductSize.Size(m)
}
func (m *ProductSize) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSize.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSize proto.InternalMessageInfo

func (m *ProductSize) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSize) GetSizeName() string {
	if m != nil {
		return m.SizeName
	}
	return ""
}

func (m *ProductSize) GetSizeCode() string {
	if m != nil {
		return m.SizeCode
	}
	return ""
}

type ProductSeo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SeoTittle            string   `protobuf:"bytes,2,opt,name=seo_tittle,json=seoTittle,proto3" json:"seo_tittle,omitempty"`
	SeoKeywords          string   `protobuf:"bytes,3,opt,name=seo_keywords,json=seoKeywords,proto3" json:"seo_keywords,omitempty"`
	SeoDescription       string   `protobuf:"bytes,4,opt,name=seo_description,json=seoDescription,proto3" json:"seo_description,omitempty"`
	SeoCode              string   `protobuf:"bytes,5,opt,name=seo_code,json=seoCode,proto3" json:"seo_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductSeo) Reset()         { *m = ProductSeo{} }
func (m *ProductSeo) String() string { return proto.CompactTextString(m) }
func (*ProductSeo) ProtoMessage()    {}
func (*ProductSeo) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fbc13de7c38f78, []int{3}
}

func (m *ProductSeo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSeo.Unmarshal(m, b)
}
func (m *ProductSeo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSeo.Marshal(b, m, deterministic)
}
func (m *ProductSeo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSeo.Merge(m, src)
}
func (m *ProductSeo) XXX_Size() int {
	return xxx_messageInfo_ProductSeo.Size(m)
}
func (m *ProductSeo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSeo.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSeo proto.InternalMessageInfo

func (m *ProductSeo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSeo) GetSeoTittle() string {
	if m != nil {
		return m.SeoTittle
	}
	return ""
}

func (m *ProductSeo) GetSeoKeywords() string {
	if m != nil {
		return m.SeoKeywords
	}
	return ""
}

func (m *ProductSeo) GetSeoDescription() string {
	if m != nil {
		return m.SeoDescription
	}
	return ""
}

func (m *ProductSeo) GetSeoCode() string {
	if m != nil {
		return m.SeoCode
	}
	return ""
}

type ProductResponse struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductResponse) Reset()         { *m = ProductResponse{} }
func (m *ProductResponse) String() string { return proto.CompactTextString(m) }
func (*ProductResponse) ProtoMessage()    {}
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fbc13de7c38f78, []int{4}
}

func (m *ProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductResponse.Unmarshal(m, b)
}
func (m *ProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductResponse.Marshal(b, m, deterministic)
}
func (m *ProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductResponse.Merge(m, src)
}
func (m *ProductResponse) XXX_Size() int {
	return xxx_messageInfo_ProductResponse.Size(m)
}
func (m *ProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductResponse proto.InternalMessageInfo

func (m *ProductResponse) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

type RequestProductId struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestProductId) Reset()         { *m = RequestProductId{} }
func (m *RequestProductId) String() string { return proto.CompactTextString(m) }
func (*RequestProductId) ProtoMessage()    {}
func (*RequestProductId) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fbc13de7c38f78, []int{5}
}

func (m *RequestProductId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestProductId.Unmarshal(m, b)
}
func (m *RequestProductId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestProductId.Marshal(b, m, deterministic)
}
func (m *RequestProductId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestProductId.Merge(m, src)
}
func (m *RequestProductId) XXX_Size() int {
	return xxx_messageInfo_RequestProductId.Size(m)
}
func (m *RequestProductId) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestProductId.DiscardUnknown(m)
}

var xxx_messageInfo_RequestProductId proto.InternalMessageInfo

func (m *RequestProductId) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fbc13de7c38f78, []int{6}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RequestProductAll struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestProductAll) Reset()         { *m = RequestProductAll{} }
func (m *RequestProductAll) String() string { return proto.CompactTextString(m) }
func (*RequestProductAll) ProtoMessage()    {}
func (*RequestProductAll) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fbc13de7c38f78, []int{7}
}

func (m *RequestProductAll) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestProductAll.Unmarshal(m, b)
}
func (m *RequestProductAll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestProductAll.Marshal(b, m, deterministic)
}
func (m *RequestProductAll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestProductAll.Merge(m, src)
}
func (m *RequestProductAll) XXX_Size() int {
	return xxx_messageInfo_RequestProductAll.Size(m)
}
func (m *RequestProductAll) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestProductAll.DiscardUnknown(m)
}

var xxx_messageInfo_RequestProductAll proto.InternalMessageInfo

type ResponseProductAll struct {
	ProductInfo          []*ProductInfo `protobuf:"bytes,1,rep,name=product_info,json=productInfo,proto3" json:"product_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResponseProductAll) Reset()         { *m = ResponseProductAll{} }
func (m *ResponseProductAll) String() string { return proto.CompactTextString(m) }
func (*ResponseProductAll) ProtoMessage()    {}
func (*ResponseProductAll) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fbc13de7c38f78, []int{8}
}

func (m *ResponseProductAll) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseProductAll.Unmarshal(m, b)
}
func (m *ResponseProductAll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseProductAll.Marshal(b, m, deterministic)
}
func (m *ResponseProductAll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseProductAll.Merge(m, src)
}
func (m *ResponseProductAll) XXX_Size() int {
	return xxx_messageInfo_ResponseProductAll.Size(m)
}
func (m *ResponseProductAll) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseProductAll.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseProductAll proto.InternalMessageInfo

func (m *ResponseProductAll) GetProductInfo() []*ProductInfo {
	if m != nil {
		return m.ProductInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductInfo)(nil), "git.micro.com.changz.product.ProductInfo")
	proto.RegisterType((*ProductImage)(nil), "git.micro.com.changz.product.ProductImage")
	proto.RegisterType((*ProductSize)(nil), "git.micro.com.changz.product.ProductSize")
	proto.RegisterType((*ProductSeo)(nil), "git.micro.com.changz.product.ProductSeo")
	proto.RegisterType((*ProductResponse)(nil), "git.micro.com.changz.product.ProductResponse")
	proto.RegisterType((*RequestProductId)(nil), "git.micro.com.changz.product.RequestProductId")
	proto.RegisterType((*Response)(nil), "git.micro.com.changz.product.Response")
	proto.RegisterType((*RequestProductAll)(nil), "git.micro.com.changz.product.RequestProductAll")
	proto.RegisterType((*ResponseProductAll)(nil), "git.micro.com.changz.product.ResponseProductAll")
}

func init() { proto.RegisterFile("proto/product/product.proto", fileDescriptor_14fbc13de7c38f78) }

var fileDescriptor_14fbc13de7c38f78 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xed, 0x4e, 0xd4, 0x40,
	0x14, 0xa5, 0x2c, 0xc2, 0xf6, 0x2e, 0x9f, 0xc3, 0x9f, 0x0a, 0x18, 0xd7, 0x9a, 0x68, 0x35, 0xb1,
	0x8b, 0xeb, 0x13, 0xac, 0x10, 0x93, 0x8d, 0x46, 0x49, 0x91, 0xf8, 0x73, 0x53, 0x3a, 0x97, 0x65,
	0xb2, 0x6d, 0xa7, 0x76, 0xda, 0x98, 0xdd, 0xc7, 0xf1, 0x05, 0x7c, 0x16, 0xdf, 0xc8, 0xcc, 0xec,
	0x4c, 0x5b, 0xd1, 0x40, 0xe1, 0x17, 0xbd, 0x1f, 0x73, 0xef, 0x39, 0xe7, 0x1e, 0x00, 0x0e, 0xb3,
	0x9c, 0x17, 0x7c, 0x90, 0xe5, 0x9c, 0x96, 0x51, 0x61, 0x7e, 0xfa, 0x2a, 0x4b, 0x8e, 0xa6, 0xac,
	0xf0, 0x13, 0x16, 0xe5, 0xdc, 0x8f, 0x78, 0xe2, 0x47, 0xd7, 0x61, 0x3a, 0x5d, 0xf8, 0xba, 0xc7,
	0xfd, 0xdd, 0x81, 0xde, 0xd9, 0xf2, 0x7b, 0x9c, 0x5e, 0x71, 0xb2, 0x0d, 0xab, 0x8c, 0x3a, 0x56,
	0xdf, 0xf2, 0x3a, 0xc1, 0x2a, 0xa3, 0xe4, 0x19, 0x6c, 0xea, 0xd6, 0x49, 0x1a, 0x26, 0xe8, 0xac,
	0xf6, 0x2d, 0xcf, 0x0e, 0x7a, 0x3a, 0xf7, 0x39, 0x4c, 0x90, 0x3c, 0x05, 0x13, 0x4e, 0xc4, 0xac,
	0x74, 0x3a, 0xaa, 0x03, 0x74, 0xea, 0x7c, 0x56, 0x92, 0xe7, 0xb0, 0x65, 0x1a, 0xb2, 0x9c, 0x45,
	0xe8, 0xac, 0xf5, 0x2d, 0xcf, 0x0a, 0xcc, 0xe0, 0x33, 0x99, 0x23, 0x03, 0xd8, 0x37, 0x4d, 0x14,
	0x45, 0x94, 0xb3, 0xac, 0x60, 0x3c, 0x75, 0x1e, 0xa9, 0x69, 0x44, 0x97, 0x4e, 0xeb, 0x0a, 0xf1,
	0xeb, 0x07, 0x51, 0x58, 0xe0, 0x94, 0xe7, 0xf3, 0x09, 0xa3, 0xce, 0xba, 0x82, 0xbe, 0xa7, 0x4b,
	0x27, 0xba, 0x32, 0xa6, 0xe4, 0x4b, 0x8d, 0x82, 0x25, 0xe1, 0x14, 0x9d, 0x8d, 0x7e, 0xc7, 0xeb,
	0x0d, 0x5f, 0xfb, 0xb7, 0xe9, 0xe3, 0x1b, 0x6d, 0xe4, 0x8b, 0x0a, 0xb1, 0x8a, 0xc8, 0xa7, 0x5a,
	0x1a, 0xc1, 0x16, 0xe8, 0x74, 0xd5, 0xbc, 0x57, 0xad, 0xe6, 0x9d, 0xb3, 0x05, 0x56, 0x2a, 0xca,
	0x80, 0x8c, 0x1b, 0x2a, 0x22, 0x77, 0xec, 0xbe, 0xe5, 0xf5, 0x86, 0x5e, 0xbb, 0x61, 0xc8, 0x6b,
	0xbd, 0x91, 0xbb, 0x73, 0xd8, 0x6c, 0xc2, 0xfe, 0xe7, 0xa6, 0x4f, 0x00, 0x94, 0x02, 0xcd, 0x8b,
	0xda, 0x2a, 0xa3, 0xee, 0x59, 0x95, 0x23, 0x4e, 0x51, 0x9f, 0x73, 0x59, 0x3e, 0xe1, 0x14, 0xc9,
	0x21, 0x2c, 0x83, 0x49, 0x99, 0xc7, 0xea, 0x92, 0x76, 0xd0, 0x55, 0x89, 0x8b, 0x3c, 0x76, 0xbf,
	0x55, 0x6e, 0x52, 0xa4, 0x6e, 0x6e, 0x3e, 0x04, 0x5b, 0x4a, 0xd5, 0x5c, 0xdc, 0x95, 0x09, 0xb5,
	0xd7, 0x14, 0x1b, 0x6b, 0x55, 0x51, 0x6e, 0x75, 0x7f, 0x5a, 0x00, 0x35, 0xdd, 0xff, 0x51, 0x12,
	0xc8, 0x27, 0x05, 0x2b, 0x8a, 0xb8, 0xa2, 0x24, 0x90, 0x7f, 0x55, 0x09, 0xe9, 0x62, 0x59, 0x9e,
	0xe1, 0xfc, 0x07, 0xcf, 0xa9, 0xd0, 0xd3, 0x7b, 0x02, 0xf9, 0x47, 0x9d, 0x22, 0x2f, 0x61, 0x47,
	0xb6, 0x34, 0xbd, 0xb7, 0x24, 0xb7, 0x2d, 0x90, 0x37, 0x7d, 0xf7, 0x18, 0xba, 0xb2, 0x51, 0xa1,
	0x5c, 0xba, 0x73, 0x43, 0x20, 0x57, 0x20, 0x8f, 0x61, 0x47, 0x63, 0x0c, 0x50, 0x64, 0x3c, 0x15,
	0x4a, 0xcc, 0xca, 0x75, 0x06, 0xb0, 0x6d, 0x6c, 0x44, 0xdd, 0xb7, 0xb0, 0x1b, 0xe0, 0xf7, 0x12,
	0x45, 0x61, 0x2e, 0x46, 0xef, 0x7a, 0x72, 0x04, 0xdd, 0x6a, 0xfa, 0x2e, 0x74, 0x12, 0x31, 0x55,
	0x3d, 0x76, 0x20, 0x3f, 0xdd, 0x7d, 0xd8, 0xfb, 0x7b, 0xe0, 0x28, 0x8e, 0xdd, 0x4b, 0x20, 0xe6,
	0x49, 0x9d, 0x6d, 0xfa, 0x97, 0xa5, 0x57, 0xdc, 0xb1, 0xee, 0xe1, 0x5f, 0xf9, 0xb7, 0xa2, 0xf2,
	0xaf, 0x0c, 0x86, 0xbf, 0xd6, 0x60, 0x43, 0x17, 0xc9, 0x35, 0xc0, 0x88, 0x52, 0x13, 0xb5, 0x9f,
	0x78, 0xf0, 0xa6, 0x55, 0xab, 0xe1, 0xe2, 0xae, 0x90, 0x14, 0x76, 0x3e, 0xb0, 0xd4, 0xac, 0x7a,
	0x2f, 0x7f, 0xcf, 0xfd, 0xdb, 0x67, 0xdc, 0x94, 0xfb, 0xa0, 0x3d, 0x3c, 0x77, 0x85, 0x50, 0xd8,
	0xba, 0xc8, 0x68, 0x58, 0xe0, 0x03, 0xc8, 0xbd, 0xb8, 0x0b, 0x58, 0xc5, 0x2a, 0x81, 0xbd, 0x53,
	0x8c, 0xb1, 0xda, 0xf2, 0x20, 0x5e, 0xed, 0xd7, 0x95, 0xb0, 0x2d, 0x45, 0x1c, 0xc5, 0xb1, 0x61,
	0x35, 0xb8, 0xcf, 0xae, 0x51, 0x1c, 0x1f, 0x1c, 0xb7, 0x5b, 0xd6, 0xf0, 0xe4, 0xca, 0xe5, 0xba,
	0xfa, 0xff, 0xf4, 0xee, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0xe2, 0x26, 0x03, 0xbe, 0x06,
	0x00, 0x00,
}
