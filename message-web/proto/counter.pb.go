// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: counter.proto

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

type CountUserMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64             `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Search map[string]string `protobuf:"bytes,2,rep,name=search,proto3" json:"search,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Status int32             `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CountUserMessageRequest) Reset() {
	*x = CountUserMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUserMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUserMessageRequest) ProtoMessage() {}

func (x *CountUserMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUserMessageRequest.ProtoReflect.Descriptor instead.
func (*CountUserMessageRequest) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{0}
}

func (x *CountUserMessageRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CountUserMessageRequest) GetSearch() map[string]string {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *CountUserMessageRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CountUserMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountUserMessageResponse) Reset() {
	*x = CountUserMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUserMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUserMessageResponse) ProtoMessage() {}

func (x *CountUserMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUserMessageResponse.ProtoReflect.Descriptor instead.
func (*CountUserMessageResponse) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{1}
}

func (x *CountUserMessageResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_counter_proto protoreflect.FileDescriptor

var file_counter_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc3, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x30, 0x0a, 0x18, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x59, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x10, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_counter_proto_rawDescOnce sync.Once
	file_counter_proto_rawDescData = file_counter_proto_rawDesc
)

func file_counter_proto_rawDescGZIP() []byte {
	file_counter_proto_rawDescOnce.Do(func() {
		file_counter_proto_rawDescData = protoimpl.X.CompressGZIP(file_counter_proto_rawDescData)
	})
	return file_counter_proto_rawDescData
}

var file_counter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_counter_proto_goTypes = []interface{}{
	(*CountUserMessageRequest)(nil),  // 0: CountUserMessageRequest
	(*CountUserMessageResponse)(nil), // 1: CountUserMessageResponse
	nil,                              // 2: CountUserMessageRequest.SearchEntry
}
var file_counter_proto_depIdxs = []int32{
	2, // 0: CountUserMessageRequest.search:type_name -> CountUserMessageRequest.SearchEntry
	0, // 1: MessageCounter.CountUserMessage:input_type -> CountUserMessageRequest
	1, // 2: MessageCounter.CountUserMessage:output_type -> CountUserMessageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_counter_proto_init() }
func file_counter_proto_init() {
	if File_counter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_counter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUserMessageRequest); i {
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
		file_counter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUserMessageResponse); i {
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
			RawDescriptor: file_counter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_counter_proto_goTypes,
		DependencyIndexes: file_counter_proto_depIdxs,
		MessageInfos:      file_counter_proto_msgTypes,
	}.Build()
	File_counter_proto = out.File
	file_counter_proto_rawDesc = nil
	file_counter_proto_goTypes = nil
	file_counter_proto_depIdxs = nil
}
