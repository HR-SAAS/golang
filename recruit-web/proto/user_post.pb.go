// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: user_post.proto

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

type GetUserPostDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserPostDetailRequest) Reset() {
	*x = GetUserPostDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPostDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPostDetailRequest) ProtoMessage() {}

func (x *GetUserPostDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPostDetailRequest.ProtoReflect.Descriptor instead.
func (*GetUserPostDetailRequest) Descriptor() ([]byte, []int) {
	return file_user_post_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserPostDetailRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BatchUpdateUserPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids      []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Remark   string  `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	ReviewId int64   `protobuf:"varint,3,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	Status   int32   `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BatchUpdateUserPostRequest) Reset() {
	*x = BatchUpdateUserPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchUpdateUserPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchUpdateUserPostRequest) ProtoMessage() {}

func (x *BatchUpdateUserPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchUpdateUserPostRequest.ProtoReflect.Descriptor instead.
func (*BatchUpdateUserPostRequest) Descriptor() ([]byte, []int) {
	return file_user_post_proto_rawDescGZIP(), []int{1}
}

func (x *BatchUpdateUserPostRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *BatchUpdateUserPostRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *BatchUpdateUserPostRequest) GetReviewId() int64 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *BatchUpdateUserPostRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UserPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PostId     int64                  `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId     int64                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResumeId   int64                  `protobuf:"varint,4,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	ResumeType string                 `protobuf:"bytes,5,opt,name=resume_type,json=resumeType,proto3" json:"resume_type,omitempty"`
	ResumeName string                 `protobuf:"bytes,6,opt,name=resume_name,json=resumeName,proto3" json:"resume_name,omitempty"`
	Resume     string                 `protobuf:"bytes,7,opt,name=resume,proto3" json:"resume,omitempty"`
	ReviewId   int64                  `protobuf:"varint,8,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	Status     int32                  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	CompanyId  int64                  `protobuf:"varint,10,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Remark     string                 `protobuf:"bytes,11,opt,name=remark,proto3" json:"remark,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserPostResponse) Reset() {
	*x = UserPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPostResponse) ProtoMessage() {}

func (x *UserPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPostResponse.ProtoReflect.Descriptor instead.
func (*UserPostResponse) Descriptor() ([]byte, []int) {
	return file_user_post_proto_rawDescGZIP(), []int{2}
}

func (x *UserPostResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserPostResponse) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *UserPostResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserPostResponse) GetResumeId() int64 {
	if x != nil {
		return x.ResumeId
	}
	return 0
}

func (x *UserPostResponse) GetResumeType() string {
	if x != nil {
		return x.ResumeType
	}
	return ""
}

func (x *UserPostResponse) GetResumeName() string {
	if x != nil {
		return x.ResumeName
	}
	return ""
}

func (x *UserPostResponse) GetResume() string {
	if x != nil {
		return x.Resume
	}
	return ""
}

func (x *UserPostResponse) GetReviewId() int64 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *UserPostResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserPostResponse) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *UserPostResponse) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *UserPostResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserPostResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetUserPostListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32             `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32             `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Sort   map[string]string `protobuf:"bytes,4,rep,name=sort,proto3" json:"sort,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Search map[string]string `protobuf:"bytes,5,rep,name=search,proto3" json:"search,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetUserPostListRequest) Reset() {
	*x = GetUserPostListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPostListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPostListRequest) ProtoMessage() {}

func (x *GetUserPostListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPostListRequest.ProtoReflect.Descriptor instead.
func (*GetUserPostListRequest) Descriptor() ([]byte, []int) {
	return file_user_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserPostListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetUserPostListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUserPostListRequest) GetSort() map[string]string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *GetUserPostListRequest) GetSearch() map[string]string {
	if x != nil {
		return x.Search
	}
	return nil
}

type UserPostListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*UserPostResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int32               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *UserPostListResponse) Reset() {
	*x = UserPostListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPostListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPostListResponse) ProtoMessage() {}

func (x *UserPostListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPostListResponse.ProtoReflect.Descriptor instead.
func (*UserPostListResponse) Descriptor() ([]byte, []int) {
	return file_user_post_proto_rawDescGZIP(), []int{4}
}

func (x *UserPostListResponse) GetData() []*UserPostResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UserPostListResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateUserPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId     int64  `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId     int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResumeId   int64  `protobuf:"varint,4,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	ResumeType string `protobuf:"bytes,5,opt,name=resume_type,json=resumeType,proto3" json:"resume_type,omitempty"`
	ResumeName string `protobuf:"bytes,6,opt,name=resume_name,json=resumeName,proto3" json:"resume_name,omitempty"`
	Resume     string `protobuf:"bytes,7,opt,name=resume,proto3" json:"resume,omitempty"`
	ReviewId   int64  `protobuf:"varint,8,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	Status     int32  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	CompanyId  int64  `protobuf:"varint,10,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Remark     string `protobuf:"bytes,11,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *CreateUserPostRequest) Reset() {
	*x = CreateUserPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserPostRequest) ProtoMessage() {}

func (x *CreateUserPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserPostRequest.ProtoReflect.Descriptor instead.
func (*CreateUserPostRequest) Descriptor() ([]byte, []int) {
	return file_user_post_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserPostRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *CreateUserPostRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateUserPostRequest) GetResumeId() int64 {
	if x != nil {
		return x.ResumeId
	}
	return 0
}

func (x *CreateUserPostRequest) GetResumeType() string {
	if x != nil {
		return x.ResumeType
	}
	return ""
}

func (x *CreateUserPostRequest) GetResumeName() string {
	if x != nil {
		return x.ResumeName
	}
	return ""
}

func (x *CreateUserPostRequest) GetResume() string {
	if x != nil {
		return x.Resume
	}
	return ""
}

func (x *CreateUserPostRequest) GetReviewId() int64 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *CreateUserPostRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateUserPostRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *CreateUserPostRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type UpdateUserPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PostId     int64  `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId     int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResumeId   int64  `protobuf:"varint,4,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	ResumeType string `protobuf:"bytes,5,opt,name=resume_type,json=resumeType,proto3" json:"resume_type,omitempty"`
	ResumeName string `protobuf:"bytes,6,opt,name=resume_name,json=resumeName,proto3" json:"resume_name,omitempty"`
	Resume     string `protobuf:"bytes,7,opt,name=resume,proto3" json:"resume,omitempty"`
	ReviewId   int64  `protobuf:"varint,8,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	Status     int32  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	CompanyId  int64  `protobuf:"varint,10,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Remark     string `protobuf:"bytes,11,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *UpdateUserPostRequest) Reset() {
	*x = UpdateUserPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserPostRequest) ProtoMessage() {}

func (x *UpdateUserPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserPostRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserPostRequest) Descriptor() ([]byte, []int) {
	return file_user_post_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserPostRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserPostRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *UpdateUserPostRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateUserPostRequest) GetResumeId() int64 {
	if x != nil {
		return x.ResumeId
	}
	return 0
}

func (x *UpdateUserPostRequest) GetResumeType() string {
	if x != nil {
		return x.ResumeType
	}
	return ""
}

func (x *UpdateUserPostRequest) GetResumeName() string {
	if x != nil {
		return x.ResumeName
	}
	return ""
}

func (x *UpdateUserPostRequest) GetResume() string {
	if x != nil {
		return x.Resume
	}
	return ""
}

func (x *UpdateUserPostRequest) GetReviewId() int64 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *UpdateUserPostRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateUserPostRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *UpdateUserPostRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type DeleteUserPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserPostRequest) Reset() {
	*x = DeleteUserPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_post_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserPostRequest) ProtoMessage() {}

func (x *DeleteUserPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_post_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserPostRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserPostRequest) Descriptor() ([]byte, []int) {
	return file_user_post_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUserPostRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_user_post_proto protoreflect.FileDescriptor

var file_user_post_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7b, 0x0a, 0x1a, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xad, 0x03, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xaa, 0x02, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x35, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f,
	0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x1a, 0x37, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x53, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xac, 0x02, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0xbc, 0x02, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x32, 0x9d, 0x03, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x41,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x16, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x40, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x16, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x41, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x13, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_post_proto_rawDescOnce sync.Once
	file_user_post_proto_rawDescData = file_user_post_proto_rawDesc
)

func file_user_post_proto_rawDescGZIP() []byte {
	file_user_post_proto_rawDescOnce.Do(func() {
		file_user_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_post_proto_rawDescData)
	})
	return file_user_post_proto_rawDescData
}

var file_user_post_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_post_proto_goTypes = []interface{}{
	(*GetUserPostDetailRequest)(nil),   // 0: GetUserPostDetailRequest
	(*BatchUpdateUserPostRequest)(nil), // 1: BatchUpdateUserPostRequest
	(*UserPostResponse)(nil),           // 2: UserPostResponse
	(*GetUserPostListRequest)(nil),     // 3: GetUserPostListRequest
	(*UserPostListResponse)(nil),       // 4: UserPostListResponse
	(*CreateUserPostRequest)(nil),      // 5: CreateUserPostRequest
	(*UpdateUserPostRequest)(nil),      // 6: UpdateUserPostRequest
	(*DeleteUserPostRequest)(nil),      // 7: DeleteUserPostRequest
	nil,                                // 8: GetUserPostListRequest.SortEntry
	nil,                                // 9: GetUserPostListRequest.SearchEntry
	(*timestamppb.Timestamp)(nil),      // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),              // 11: google.protobuf.Empty
}
var file_user_post_proto_depIdxs = []int32{
	10, // 0: UserPostResponse.created_at:type_name -> google.protobuf.Timestamp
	10, // 1: UserPostResponse.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 2: GetUserPostListRequest.sort:type_name -> GetUserPostListRequest.SortEntry
	9,  // 3: GetUserPostListRequest.search:type_name -> GetUserPostListRequest.SearchEntry
	2,  // 4: UserPostListResponse.data:type_name -> UserPostResponse
	3,  // 5: UserPost.GetUserPostList:input_type -> GetUserPostListRequest
	5,  // 6: UserPost.CreateUserPost:input_type -> CreateUserPostRequest
	6,  // 7: UserPost.UpdateUserPost:input_type -> UpdateUserPostRequest
	7,  // 8: UserPost.DeleteUserPost:input_type -> DeleteUserPostRequest
	0,  // 9: UserPost.GetUserPostDetail:input_type -> GetUserPostDetailRequest
	1,  // 10: UserPost.BatchUpdateUserPost:input_type -> BatchUpdateUserPostRequest
	4,  // 11: UserPost.GetUserPostList:output_type -> UserPostListResponse
	2,  // 12: UserPost.CreateUserPost:output_type -> UserPostResponse
	11, // 13: UserPost.UpdateUserPost:output_type -> google.protobuf.Empty
	11, // 14: UserPost.DeleteUserPost:output_type -> google.protobuf.Empty
	2,  // 15: UserPost.GetUserPostDetail:output_type -> UserPostResponse
	11, // 16: UserPost.BatchUpdateUserPost:output_type -> google.protobuf.Empty
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_user_post_proto_init() }
func file_user_post_proto_init() {
	if File_user_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPostDetailRequest); i {
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
		file_user_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchUpdateUserPostRequest); i {
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
		file_user_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPostResponse); i {
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
		file_user_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPostListRequest); i {
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
		file_user_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPostListResponse); i {
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
		file_user_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserPostRequest); i {
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
		file_user_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserPostRequest); i {
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
		file_user_post_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserPostRequest); i {
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
			RawDescriptor: file_user_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_post_proto_goTypes,
		DependencyIndexes: file_user_post_proto_depIdxs,
		MessageInfos:      file_user_post_proto_msgTypes,
	}.Build()
	File_user_post_proto = out.File
	file_user_post_proto_rawDesc = nil
	file_user_post_proto_goTypes = nil
	file_user_post_proto_depIdxs = nil
}
