// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.2
// source: proto/indrasaputra/shortener/v1/shortener.proto

package shortenerv1

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type URLShortenerErrorCode int32

const (
	// Default enum code according to
	// https://medium.com/@akhaku/protobuf-definition-best-practices-87f281576f31.
	URLShortenerErrorCode_UNKNOWN_URL_SHORTENER_ERROR_CODE URLShortenerErrorCode = 0
	// Unexpected behavior occured in system.
	URLShortenerErrorCode_INTERNAL URLShortenerErrorCode = 1
	// URL instance is empty or nil.
	URLShortenerErrorCode_EMPTY_URL URLShortenerErrorCode = 2
	// Short URL already exists.
	// The uniqueness of a short URL is represented by code or short URL.
	URLShortenerErrorCode_ALREADY_EXISTS URLShortenerErrorCode = 3
	// URL not found in system.
	URLShortenerErrorCode_NOT_FOUND URLShortenerErrorCode = 4
)

// Enum value maps for URLShortenerErrorCode.
var (
	URLShortenerErrorCode_name = map[int32]string{
		0: "UNKNOWN_URL_SHORTENER_ERROR_CODE",
		1: "INTERNAL",
		2: "EMPTY_URL",
		3: "ALREADY_EXISTS",
		4: "NOT_FOUND",
	}
	URLShortenerErrorCode_value = map[string]int32{
		"UNKNOWN_URL_SHORTENER_ERROR_CODE": 0,
		"INTERNAL":                         1,
		"EMPTY_URL":                        2,
		"ALREADY_EXISTS":                   3,
		"NOT_FOUND":                        4,
	}
)

func (x URLShortenerErrorCode) Enum() *URLShortenerErrorCode {
	p := new(URLShortenerErrorCode)
	*p = x
	return p
}

func (x URLShortenerErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (URLShortenerErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_indrasaputra_shortener_v1_shortener_proto_enumTypes[0].Descriptor()
}

func (URLShortenerErrorCode) Type() protoreflect.EnumType {
	return &file_proto_indrasaputra_shortener_v1_shortener_proto_enumTypes[0]
}

func (x URLShortenerErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use URLShortenerErrorCode.Descriptor instead.
func (URLShortenerErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescGZIP(), []int{0}
}

type CreateShortURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *CreateShortURLRequest) Reset() {
	*x = CreateShortURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortURLRequest) ProtoMessage() {}

func (x *CreateShortURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortURLRequest.ProtoReflect.Descriptor instead.
func (*CreateShortURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShortURLRequest) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type CreateShortURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url *URL `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateShortURLResponse) Reset() {
	*x = CreateShortURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortURLResponse) ProtoMessage() {}

func (x *CreateShortURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortURLResponse.ProtoReflect.Descriptor instead.
func (*CreateShortURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShortURLResponse) GetUrl() *URL {
	if x != nil {
		return x.Url
	}
	return nil
}

type GetAllURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllURLRequest) Reset() {
	*x = GetAllURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllURLRequest) ProtoMessage() {}

func (x *GetAllURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllURLRequest.ProtoReflect.Descriptor instead.
func (*GetAllURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescGZIP(), []int{2}
}

type GetAllURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*URL `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *GetAllURLResponse) Reset() {
	*x = GetAllURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllURLResponse) ProtoMessage() {}

func (x *GetAllURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllURLResponse.ProtoReflect.Descriptor instead.
func (*GetAllURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllURLResponse) GetUrls() []*URL {
	if x != nil {
		return x.Urls
	}
	return nil
}

type StreamAllURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamAllURLRequest) Reset() {
	*x = StreamAllURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAllURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAllURLRequest) ProtoMessage() {}

func (x *StreamAllURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAllURLRequest.ProtoReflect.Descriptor instead.
func (*StreamAllURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescGZIP(), []int{4}
}

type StreamAllURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url *URL `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *StreamAllURLResponse) Reset() {
	*x = StreamAllURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAllURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAllURLResponse) ProtoMessage() {}

func (x *StreamAllURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAllURLResponse.ProtoReflect.Descriptor instead.
func (*StreamAllURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescGZIP(), []int{5}
}

func (x *StreamAllURLResponse) GetUrl() *URL {
	if x != nil {
		return x.Url
	}
	return nil
}

type GetURLDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetURLDetailRequest) Reset() {
	*x = GetURLDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLDetailRequest) ProtoMessage() {}

func (x *GetURLDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLDetailRequest.ProtoReflect.Descriptor instead.
func (*GetURLDetailRequest) Descriptor() ([]byte, []int) {
	return file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *GetURLDetailRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetURLDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url *URL `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetURLDetailResponse) Reset() {
	*x = GetURLDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLDetailResponse) ProtoMessage() {}

func (x *GetURLDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLDetailResponse.ProtoReflect.Descriptor instead.
func (*GetURLDetailResponse) Descriptor() ([]byte, []int) {
	return file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *GetURLDetailResponse) GetUrl() *URL {
	if x != nil {
		return x.Url
	}
	return nil
}

type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	ShortUrl    string                 `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl string                 `protobuf:"bytes,3,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	ExpiredAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *URL) Reset() {
	*x = URL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URL.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *URL) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *URL) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *URL) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

func (x *URL) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

func (x *URL) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type URLShortenerError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode URLShortenerErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=proto.indrasaputra.shortener.v1.URLShortenerErrorCode" json:"error_code,omitempty"`
}

func (x *URLShortenerError) Reset() {
	*x = URLShortenerError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLShortenerError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLShortenerError) ProtoMessage() {}

func (x *URLShortenerError) ProtoReflect() protoreflect.Message {
	mi := &file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLShortenerError.ProtoReflect.Descriptor instead.
func (*URLShortenerError) Descriptor() ([]byte, []int) {
	return file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescGZIP(), []int{9}
}

func (x *URLShortenerError) GetErrorCode() URLShortenerErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return URLShortenerErrorCode_UNKNOWN_URL_SHORTENER_ERROR_CODE
}

var File_proto_indrasaputra_shortener_v1_shortener_proto protoreflect.FileDescriptor

var file_proto_indrasaputra_shortener_v1_shortener_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x61, 0x70,
	0x75, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x61,
	0x70, 0x75, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3a, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x50, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x64,
	0x72, 0x61, 0x73, 0x61, 0x70, 0x75, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22,
	0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69,
	0x6e, 0x64, 0x72, 0x61, 0x73, 0x61, 0x70, 0x75, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x04, 0x75, 0x72,
	0x6c, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x6c, 0x6c, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x14, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x41, 0x6c, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x61, 0x70, 0x75,
	0x74, 0x72, 0x61, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x52, 0x4c, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x55, 0x52, 0x4c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x4e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x69, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x61, 0x70, 0x75, 0x74, 0x72, 0x61, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x52, 0x4c, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0xcf, 0x01, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c,
	0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6a, 0x0a, 0x11, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x55, 0x0a, 0x0a, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x61, 0x70,
	0x75, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x2a, 0x7d, 0x0a, 0x15, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x53, 0x48, 0x4f, 0x52, 0x54,
	0x45, 0x4e, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53,
	0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x04, 0x32, 0xe5, 0x04, 0x0a, 0x13, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x36, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x61, 0x70, 0x75, 0x74, 0x72,
	0x61, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x64,
	0x72, 0x61, 0x73, 0x61, 0x70, 0x75, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x84, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x52, 0x4c,
	0x12, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x61,
	0x70, 0x75, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x64, 0x72,
	0x61, 0x73, 0x61, 0x70, 0x75, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12,
	0x08, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x96, 0x01, 0x0a, 0x0c, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x41, 0x6c, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x34, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x61, 0x70, 0x75, 0x74, 0x72, 0x61, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x41, 0x6c, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x61,
	0x70, 0x75, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x6c, 0x6c, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x75, 0x72, 0x6c, 0x73,
	0x30, 0x01, 0x12, 0x94, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x64, 0x72,
	0x61, 0x73, 0x61, 0x70, 0x75, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x69, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x61, 0x70, 0x75, 0x74, 0x72, 0x61, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x52, 0x4c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72,
	0x6c, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x64, 0x65, 0x7d, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x61, 0x70,
	0x75, 0x74, 0x72, 0x61, 0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x64, 0x72, 0x61, 0x73, 0x61,
	0x70, 0x75, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescOnce sync.Once
	file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescData = file_proto_indrasaputra_shortener_v1_shortener_proto_rawDesc
)

func file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescGZIP() []byte {
	file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescOnce.Do(func() {
		file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescData)
	})
	return file_proto_indrasaputra_shortener_v1_shortener_proto_rawDescData
}

var file_proto_indrasaputra_shortener_v1_shortener_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_indrasaputra_shortener_v1_shortener_proto_goTypes = []interface{}{
	(URLShortenerErrorCode)(0),     // 0: proto.indrasaputra.shortener.v1.URLShortenerErrorCode
	(*CreateShortURLRequest)(nil),  // 1: proto.indrasaputra.shortener.v1.CreateShortURLRequest
	(*CreateShortURLResponse)(nil), // 2: proto.indrasaputra.shortener.v1.CreateShortURLResponse
	(*GetAllURLRequest)(nil),       // 3: proto.indrasaputra.shortener.v1.GetAllURLRequest
	(*GetAllURLResponse)(nil),      // 4: proto.indrasaputra.shortener.v1.GetAllURLResponse
	(*StreamAllURLRequest)(nil),    // 5: proto.indrasaputra.shortener.v1.StreamAllURLRequest
	(*StreamAllURLResponse)(nil),   // 6: proto.indrasaputra.shortener.v1.StreamAllURLResponse
	(*GetURLDetailRequest)(nil),    // 7: proto.indrasaputra.shortener.v1.GetURLDetailRequest
	(*GetURLDetailResponse)(nil),   // 8: proto.indrasaputra.shortener.v1.GetURLDetailResponse
	(*URL)(nil),                    // 9: proto.indrasaputra.shortener.v1.URL
	(*URLShortenerError)(nil),      // 10: proto.indrasaputra.shortener.v1.URLShortenerError
	(*timestamppb.Timestamp)(nil),  // 11: google.protobuf.Timestamp
}
var file_proto_indrasaputra_shortener_v1_shortener_proto_depIdxs = []int32{
	9,  // 0: proto.indrasaputra.shortener.v1.CreateShortURLResponse.url:type_name -> proto.indrasaputra.shortener.v1.URL
	9,  // 1: proto.indrasaputra.shortener.v1.GetAllURLResponse.urls:type_name -> proto.indrasaputra.shortener.v1.URL
	9,  // 2: proto.indrasaputra.shortener.v1.StreamAllURLResponse.url:type_name -> proto.indrasaputra.shortener.v1.URL
	9,  // 3: proto.indrasaputra.shortener.v1.GetURLDetailResponse.url:type_name -> proto.indrasaputra.shortener.v1.URL
	11, // 4: proto.indrasaputra.shortener.v1.URL.expired_at:type_name -> google.protobuf.Timestamp
	11, // 5: proto.indrasaputra.shortener.v1.URL.created_at:type_name -> google.protobuf.Timestamp
	0,  // 6: proto.indrasaputra.shortener.v1.URLShortenerError.error_code:type_name -> proto.indrasaputra.shortener.v1.URLShortenerErrorCode
	1,  // 7: proto.indrasaputra.shortener.v1.URLShortenerService.CreateShortURL:input_type -> proto.indrasaputra.shortener.v1.CreateShortURLRequest
	3,  // 8: proto.indrasaputra.shortener.v1.URLShortenerService.GetAllURL:input_type -> proto.indrasaputra.shortener.v1.GetAllURLRequest
	5,  // 9: proto.indrasaputra.shortener.v1.URLShortenerService.StreamAllURL:input_type -> proto.indrasaputra.shortener.v1.StreamAllURLRequest
	7,  // 10: proto.indrasaputra.shortener.v1.URLShortenerService.GetURLDetail:input_type -> proto.indrasaputra.shortener.v1.GetURLDetailRequest
	2,  // 11: proto.indrasaputra.shortener.v1.URLShortenerService.CreateShortURL:output_type -> proto.indrasaputra.shortener.v1.CreateShortURLResponse
	4,  // 12: proto.indrasaputra.shortener.v1.URLShortenerService.GetAllURL:output_type -> proto.indrasaputra.shortener.v1.GetAllURLResponse
	6,  // 13: proto.indrasaputra.shortener.v1.URLShortenerService.StreamAllURL:output_type -> proto.indrasaputra.shortener.v1.StreamAllURLResponse
	8,  // 14: proto.indrasaputra.shortener.v1.URLShortenerService.GetURLDetail:output_type -> proto.indrasaputra.shortener.v1.GetURLDetailResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_indrasaputra_shortener_v1_shortener_proto_init() }
func file_proto_indrasaputra_shortener_v1_shortener_proto_init() {
	if File_proto_indrasaputra_shortener_v1_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortURLRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortURLResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllURLRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllURLResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAllURLRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAllURLResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLDetailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URL); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLShortenerError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_indrasaputra_shortener_v1_shortener_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_indrasaputra_shortener_v1_shortener_proto_goTypes,
		DependencyIndexes: file_proto_indrasaputra_shortener_v1_shortener_proto_depIdxs,
		EnumInfos:         file_proto_indrasaputra_shortener_v1_shortener_proto_enumTypes,
		MessageInfos:      file_proto_indrasaputra_shortener_v1_shortener_proto_msgTypes,
	}.Build()
	File_proto_indrasaputra_shortener_v1_shortener_proto = out.File
	file_proto_indrasaputra_shortener_v1_shortener_proto_rawDesc = nil
	file_proto_indrasaputra_shortener_v1_shortener_proto_goTypes = nil
	file_proto_indrasaputra_shortener_v1_shortener_proto_depIdxs = nil
}
