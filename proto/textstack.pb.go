// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/textstack.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListTagsByUUIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TextUuid string `protobuf:"bytes,2,opt,name=text_uuid,json=textUuid,proto3" json:"text_uuid,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListTagsByUUIDResponse) Reset() {
	*x = ListTagsByUUIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_textstack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsByUUIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsByUUIDResponse) ProtoMessage() {}

func (x *ListTagsByUUIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_textstack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsByUUIDResponse.ProtoReflect.Descriptor instead.
func (*ListTagsByUUIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_textstack_proto_rawDescGZIP(), []int{0}
}

func (x *ListTagsByUUIDResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListTagsByUUIDResponse) GetTextUuid() string {
	if x != nil {
		return x.TextUuid
	}
	return ""
}

func (x *ListTagsByUUIDResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListTagsByUUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TextUuid string `protobuf:"bytes,1,opt,name=text_uuid,json=textUuid,proto3" json:"text_uuid,omitempty"`
}

func (x *ListTagsByUUIDRequest) Reset() {
	*x = ListTagsByUUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_textstack_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsByUUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsByUUIDRequest) ProtoMessage() {}

func (x *ListTagsByUUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_textstack_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsByUUIDRequest.ProtoReflect.Descriptor instead.
func (*ListTagsByUUIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_textstack_proto_rawDescGZIP(), []int{1}
}

func (x *ListTagsByUUIDRequest) GetTextUuid() string {
	if x != nil {
		return x.TextUuid
	}
	return ""
}

type ListTagsByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TextUuid string `protobuf:"bytes,2,opt,name=text_uuid,json=textUuid,proto3" json:"text_uuid,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListTagsByNameResponse) Reset() {
	*x = ListTagsByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_textstack_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsByNameResponse) ProtoMessage() {}

func (x *ListTagsByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_textstack_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsByNameResponse.ProtoReflect.Descriptor instead.
func (*ListTagsByNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_textstack_proto_rawDescGZIP(), []int{2}
}

func (x *ListTagsByNameResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListTagsByNameResponse) GetTextUuid() string {
	if x != nil {
		return x.TextUuid
	}
	return ""
}

func (x *ListTagsByNameResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListTagsByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListTagsByNameRequest) Reset() {
	*x = ListTagsByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_textstack_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsByNameRequest) ProtoMessage() {}

func (x *ListTagsByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_textstack_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsByNameRequest.ProtoReflect.Descriptor instead.
func (*ListTagsByNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_textstack_proto_rawDescGZIP(), []int{3}
}

func (x *ListTagsByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTagResponse) Reset() {
	*x = CreateTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_textstack_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagResponse) ProtoMessage() {}

func (x *CreateTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_textstack_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagResponse.ProtoReflect.Descriptor instead.
func (*CreateTagResponse) Descriptor() ([]byte, []int) {
	return file_proto_textstack_proto_rawDescGZIP(), []int{4}
}

type CreateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TextUuid string `protobuf:"bytes,1,opt,name=text_uuid,json=textUuid,proto3" json:"text_uuid,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateTagRequest) Reset() {
	*x = CreateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_textstack_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagRequest) ProtoMessage() {}

func (x *CreateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_textstack_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagRequest.ProtoReflect.Descriptor instead.
func (*CreateTagRequest) Descriptor() ([]byte, []int) {
	return file_proto_textstack_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTagRequest) GetTextUuid() string {
	if x != nil {
		return x.TextUuid
	}
	return ""
}

func (x *CreateTagRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid  string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body  string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *GetTextResponse) Reset() {
	*x = GetTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_textstack_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextResponse) ProtoMessage() {}

func (x *GetTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_textstack_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextResponse.ProtoReflect.Descriptor instead.
func (*GetTextResponse) Descriptor() ([]byte, []int) {
	return file_proto_textstack_proto_rawDescGZIP(), []int{6}
}

func (x *GetTextResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetTextResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetTextResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetTextResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type GetTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetTextRequest) Reset() {
	*x = GetTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_textstack_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextRequest) ProtoMessage() {}

func (x *GetTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_textstack_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextRequest.ProtoReflect.Descriptor instead.
func (*GetTextRequest) Descriptor() ([]byte, []int) {
	return file_proto_textstack_proto_rawDescGZIP(), []int{7}
}

func (x *GetTextRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CreateTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *CreateTextResponse) Reset() {
	*x = CreateTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_textstack_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTextResponse) ProtoMessage() {}

func (x *CreateTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_textstack_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTextResponse.ProtoReflect.Descriptor instead.
func (*CreateTextResponse) Descriptor() ([]byte, []int) {
	return file_proto_textstack_proto_rawDescGZIP(), []int{8}
}

func (x *CreateTextResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CreateTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CreateTextRequest) Reset() {
	*x = CreateTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_textstack_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTextRequest) ProtoMessage() {}

func (x *CreateTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_textstack_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTextRequest.ProtoReflect.Descriptor instead.
func (*CreateTextRequest) Descriptor() ([]byte, []int) {
	return file_proto_textstack_proto_rawDescGZIP(), []int{9}
}

func (x *CreateTextRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTextRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_proto_textstack_proto protoreflect.FileDescriptor

var file_proto_textstack_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x61, 0x67, 0x73, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x78, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x34, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x42, 0x79,
	0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x65, 0x78, 0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x61, 0x67, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x78, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x78, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5f, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x24, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x22, 0x28, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0xb4, 0x02, 0x0a, 0x09, 0x54,
	0x65, 0x78, 0x74, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x43, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x61, 0x67, 0x73, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x12, 0x16, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x67, 0x73, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x42, 0x79, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x67, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x12,
	0x11, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x30, 0x63, 0x63, 0x68, 0x69, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_textstack_proto_rawDescOnce sync.Once
	file_proto_textstack_proto_rawDescData = file_proto_textstack_proto_rawDesc
)

func file_proto_textstack_proto_rawDescGZIP() []byte {
	file_proto_textstack_proto_rawDescOnce.Do(func() {
		file_proto_textstack_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_textstack_proto_rawDescData)
	})
	return file_proto_textstack_proto_rawDescData
}

var file_proto_textstack_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_textstack_proto_goTypes = []interface{}{
	(*ListTagsByUUIDResponse)(nil), // 0: ListTagsByUUIDResponse
	(*ListTagsByUUIDRequest)(nil),  // 1: ListTagsByUUIDRequest
	(*ListTagsByNameResponse)(nil), // 2: ListTagsByNameResponse
	(*ListTagsByNameRequest)(nil),  // 3: ListTagsByNameRequest
	(*CreateTagResponse)(nil),      // 4: CreateTagResponse
	(*CreateTagRequest)(nil),       // 5: CreateTagRequest
	(*GetTextResponse)(nil),        // 6: GetTextResponse
	(*GetTextRequest)(nil),         // 7: GetTextRequest
	(*CreateTextResponse)(nil),     // 8: CreateTextResponse
	(*CreateTextRequest)(nil),      // 9: CreateTextRequest
}
var file_proto_textstack_proto_depIdxs = []int32{
	1, // 0: Textstack.ListTagsByUUID:input_type -> ListTagsByUUIDRequest
	3, // 1: Textstack.ListTagsByName:input_type -> ListTagsByNameRequest
	5, // 2: Textstack.CreateTag:input_type -> CreateTagRequest
	7, // 3: Textstack.GetText:input_type -> GetTextRequest
	9, // 4: Textstack.CreateText:input_type -> CreateTextRequest
	0, // 5: Textstack.ListTagsByUUID:output_type -> ListTagsByUUIDResponse
	2, // 6: Textstack.ListTagsByName:output_type -> ListTagsByNameResponse
	4, // 7: Textstack.CreateTag:output_type -> CreateTagResponse
	6, // 8: Textstack.GetText:output_type -> GetTextResponse
	8, // 9: Textstack.CreateText:output_type -> CreateTextResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_textstack_proto_init() }
func file_proto_textstack_proto_init() {
	if File_proto_textstack_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_textstack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsByUUIDResponse); i {
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
		file_proto_textstack_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsByUUIDRequest); i {
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
		file_proto_textstack_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsByNameResponse); i {
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
		file_proto_textstack_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsByNameRequest); i {
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
		file_proto_textstack_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagResponse); i {
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
		file_proto_textstack_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagRequest); i {
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
		file_proto_textstack_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextResponse); i {
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
		file_proto_textstack_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextRequest); i {
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
		file_proto_textstack_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTextResponse); i {
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
		file_proto_textstack_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTextRequest); i {
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
			RawDescriptor: file_proto_textstack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_textstack_proto_goTypes,
		DependencyIndexes: file_proto_textstack_proto_depIdxs,
		MessageInfos:      file_proto_textstack_proto_msgTypes,
	}.Build()
	File_proto_textstack_proto = out.File
	file_proto_textstack_proto_rawDesc = nil
	file_proto_textstack_proto_goTypes = nil
	file_proto_textstack_proto_depIdxs = nil
}
