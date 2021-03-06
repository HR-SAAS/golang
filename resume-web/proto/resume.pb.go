// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: resume.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetResumeDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetResumeDetailRequest) Reset() {
	*x = GetResumeDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResumeDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResumeDetailRequest) ProtoMessage() {}

func (x *GetResumeDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResumeDetailRequest.ProtoReflect.Descriptor instead.
func (*GetResumeDetailRequest) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{0}
}

func (x *GetResumeDetailRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ResumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name      string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type      string                 `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Tag       []string               `protobuf:"bytes,5,rep,name=tag,proto3" json:"tag,omitempty"`
	Content   string                 `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Status    int32                  `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	PostCount int32                  `protobuf:"varint,8,opt,name=post_count,json=postCount,proto3" json:"post_count,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ResumeResponse) Reset() {
	*x = ResumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeResponse) ProtoMessage() {}

func (x *ResumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeResponse.ProtoReflect.Descriptor instead.
func (*ResumeResponse) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{1}
}

func (x *ResumeResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResumeResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ResumeResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResumeResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ResumeResponse) GetTag() []string {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *ResumeResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ResumeResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ResumeResponse) GetPostCount() int32 {
	if x != nil {
		return x.PostCount
	}
	return 0
}

func (x *ResumeResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ResumeResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetResumeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32             `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32             `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Sort   map[string]string `protobuf:"bytes,4,rep,name=sort,proto3" json:"sort,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Search map[string]string `protobuf:"bytes,5,rep,name=search,proto3" json:"search,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetResumeListRequest) Reset() {
	*x = GetResumeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResumeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResumeListRequest) ProtoMessage() {}

func (x *GetResumeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResumeListRequest.ProtoReflect.Descriptor instead.
func (*GetResumeListRequest) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{2}
}

func (x *GetResumeListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetResumeListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetResumeListRequest) GetSort() map[string]string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *GetResumeListRequest) GetSearch() map[string]string {
	if x != nil {
		return x.Search
	}
	return nil
}

type ResumeListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*ResumeResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int32             `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ResumeListResponse) Reset() {
	*x = ResumeListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeListResponse) ProtoMessage() {}

func (x *ResumeListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeListResponse.ProtoReflect.Descriptor instead.
func (*ResumeListResponse) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{3}
}

func (x *ResumeListResponse) GetData() []*ResumeResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ResumeListResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateResumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type    string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Tag     []string `protobuf:"bytes,4,rep,name=tag,proto3" json:"tag,omitempty"`
	Content string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Status  int32    `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateResumeRequest) Reset() {
	*x = CreateResumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResumeRequest) ProtoMessage() {}

func (x *CreateResumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResumeRequest.ProtoReflect.Descriptor instead.
func (*CreateResumeRequest) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{4}
}

func (x *CreateResumeRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateResumeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateResumeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateResumeRequest) GetTag() []string {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *CreateResumeRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateResumeRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UpdateResumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId  int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name    string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type    string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Tag     []string `protobuf:"bytes,5,rep,name=tag,proto3" json:"tag,omitempty"`
	Content string   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Status  int32    `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateResumeRequest) Reset() {
	*x = UpdateResumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResumeRequest) ProtoMessage() {}

func (x *UpdateResumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResumeRequest.ProtoReflect.Descriptor instead.
func (*UpdateResumeRequest) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateResumeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateResumeRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateResumeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateResumeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateResumeRequest) GetTag() []string {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *UpdateResumeRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateResumeRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DeleteResumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteResumeRequest) Reset() {
	*x = DeleteResumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResumeRequest) ProtoMessage() {}

func (x *DeleteResumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResumeRequest.ProtoReflect.Descriptor instead.
func (*DeleteResumeRequest) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteResumeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_resume_proto protoreflect.FileDescriptor

var file_resume_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xba, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xa4, 0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x37, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x39, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4f, 0x0a, 0x12, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x9a, 0x01, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xb5, 0x02, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x14, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resume_proto_rawDescOnce sync.Once
	file_resume_proto_rawDescData = file_resume_proto_rawDesc
)

func file_resume_proto_rawDescGZIP() []byte {
	file_resume_proto_rawDescOnce.Do(func() {
		file_resume_proto_rawDescData = protoimpl.X.CompressGZIP(file_resume_proto_rawDescData)
	})
	return file_resume_proto_rawDescData
}

var file_resume_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_resume_proto_goTypes = []interface{}{
	(*GetResumeDetailRequest)(nil), // 0: GetResumeDetailRequest
	(*ResumeResponse)(nil),         // 1: ResumeResponse
	(*GetResumeListRequest)(nil),   // 2: GetResumeListRequest
	(*ResumeListResponse)(nil),     // 3: ResumeListResponse
	(*CreateResumeRequest)(nil),    // 4: CreateResumeRequest
	(*UpdateResumeRequest)(nil),    // 5: UpdateResumeRequest
	(*DeleteResumeRequest)(nil),    // 6: DeleteResumeRequest
	nil,                            // 7: GetResumeListRequest.SortEntry
	nil,                            // 8: GetResumeListRequest.SearchEntry
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 10: google.protobuf.Empty
}
var file_resume_proto_depIdxs = []int32{
	9,  // 0: ResumeResponse.created_at:type_name -> google.protobuf.Timestamp
	9,  // 1: ResumeResponse.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 2: GetResumeListRequest.sort:type_name -> GetResumeListRequest.SortEntry
	8,  // 3: GetResumeListRequest.search:type_name -> GetResumeListRequest.SearchEntry
	1,  // 4: ResumeListResponse.data:type_name -> ResumeResponse
	2,  // 5: Resume.GetResumeList:input_type -> GetResumeListRequest
	4,  // 6: Resume.CreateResume:input_type -> CreateResumeRequest
	5,  // 7: Resume.UpdateResume:input_type -> UpdateResumeRequest
	6,  // 8: Resume.DeleteResume:input_type -> DeleteResumeRequest
	0,  // 9: Resume.GetResumeDetail:input_type -> GetResumeDetailRequest
	3,  // 10: Resume.GetResumeList:output_type -> ResumeListResponse
	1,  // 11: Resume.CreateResume:output_type -> ResumeResponse
	10, // 12: Resume.UpdateResume:output_type -> google.protobuf.Empty
	10, // 13: Resume.DeleteResume:output_type -> google.protobuf.Empty
	1,  // 14: Resume.GetResumeDetail:output_type -> ResumeResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_resume_proto_init() }
func file_resume_proto_init() {
	if File_resume_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResumeDetailRequest); i {
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
		file_resume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeResponse); i {
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
		file_resume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResumeListRequest); i {
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
		file_resume_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeListResponse); i {
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
		file_resume_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResumeRequest); i {
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
		file_resume_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResumeRequest); i {
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
		file_resume_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResumeRequest); i {
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
			RawDescriptor: file_resume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resume_proto_goTypes,
		DependencyIndexes: file_resume_proto_depIdxs,
		MessageInfos:      file_resume_proto_msgTypes,
	}.Build()
	File_resume_proto = out.File
	file_resume_proto_rawDesc = nil
	file_resume_proto_goTypes = nil
	file_resume_proto_depIdxs = nil
}
