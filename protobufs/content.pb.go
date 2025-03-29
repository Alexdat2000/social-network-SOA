// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: content.proto

package content_grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          string                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	PostId        uint32                 `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_content_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *GetRequest) GetPostId() uint32 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type PostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          string                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IsPrivate     bool                   `protobuf:"varint,4,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
	Tags          []string               `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostRequest) Reset() {
	*x = PostRequest{}
	mi := &file_content_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRequest) ProtoMessage() {}

func (x *PostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRequest.ProtoReflect.Descriptor instead.
func (*PostRequest) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{1}
}

func (x *PostRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *PostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PostRequest) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *PostRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type PutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          string                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	PostId        uint32                 `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	IsPrivate     bool                   `protobuf:"varint,5,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
	Tags          []string               `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	mi := &file_content_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{2}
}

func (x *PutRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *PutRequest) GetPostId() uint32 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *PutRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PutRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PutRequest) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *PutRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type PostInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        uint32                 `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Author        string                 `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsPrivate     bool                   `protobuf:"varint,7,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
	Tags          []string               `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostInfo) Reset() {
	*x = PostInfo{}
	mi := &file_content_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostInfo) ProtoMessage() {}

func (x *PostInfo) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostInfo.ProtoReflect.Descriptor instead.
func (*PostInfo) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{3}
}

func (x *PostInfo) GetPostId() uint32 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *PostInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PostInfo) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *PostInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PostInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PostInfo) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *PostInfo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_content_proto protoreflect.FileDescriptor

var file_content_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x0b, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x0a, 0x50, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22,
	0x9c, 0x02, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x32, 0xbc,
	0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x12, 0x5a,
	0x10, 0x73, 0x6f, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_content_proto_rawDescOnce sync.Once
	file_content_proto_rawDescData []byte
)

func file_content_proto_rawDescGZIP() []byte {
	file_content_proto_rawDescOnce.Do(func() {
		file_content_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_content_proto_rawDesc), len(file_content_proto_rawDesc)))
	})
	return file_content_proto_rawDescData
}

var file_content_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_content_proto_goTypes = []any{
	(*GetRequest)(nil),            // 0: content_grpc.GetRequest
	(*PostRequest)(nil),           // 1: content_grpc.PostRequest
	(*PutRequest)(nil),            // 2: content_grpc.PutRequest
	(*PostInfo)(nil),              // 3: content_grpc.PostInfo
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_content_proto_depIdxs = []int32{
	4, // 0: content_grpc.PostInfo.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: content_grpc.PostInfo.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: content_grpc.Content.Get:input_type -> content_grpc.GetRequest
	1, // 3: content_grpc.Content.Post:input_type -> content_grpc.PostRequest
	2, // 4: content_grpc.Content.Put:input_type -> content_grpc.PutRequest
	3, // 5: content_grpc.Content.Get:output_type -> content_grpc.PostInfo
	3, // 6: content_grpc.Content.Post:output_type -> content_grpc.PostInfo
	3, // 7: content_grpc.Content.Put:output_type -> content_grpc.PostInfo
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_content_proto_init() }
func file_content_proto_init() {
	if File_content_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_content_proto_rawDesc), len(file_content_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_content_proto_goTypes,
		DependencyIndexes: file_content_proto_depIdxs,
		MessageInfos:      file_content_proto_msgTypes,
	}.Build()
	File_content_proto = out.File
	file_content_proto_goTypes = nil
	file_content_proto_depIdxs = nil
}
