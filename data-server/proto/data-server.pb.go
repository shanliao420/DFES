// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0--rc3
// source: data-server/proto/data-server.proto

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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FragmentId string `protobuf:"bytes,1,opt,name=FragmentId,proto3" json:"FragmentId,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_server_proto_data_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_server_proto_data_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_data_server_proto_data_server_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetFragmentId() string {
	if x != nil {
		return x.FragmentId
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FragmentData []byte `protobuf:"bytes,1,opt,name=FragmentData,proto3" json:"FragmentData,omitempty"`
	GetMessage   string `protobuf:"bytes,2,opt,name=GetMessage,proto3" json:"GetMessage,omitempty"`
	GetResult    bool   `protobuf:"varint,3,opt,name=GetResult,proto3" json:"GetResult,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_server_proto_data_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_server_proto_data_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_data_server_proto_data_server_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetFragmentData() []byte {
	if x != nil {
		return x.FragmentData
	}
	return nil
}

func (x *GetResponse) GetGetMessage() string {
	if x != nil {
		return x.GetMessage
	}
	return ""
}

func (x *GetResponse) GetGetResult() bool {
	if x != nil {
		return x.GetResult
	}
	return false
}

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FragmentData []byte `protobuf:"bytes,1,opt,name=FragmentData,proto3" json:"FragmentData,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_server_proto_data_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_server_proto_data_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_data_server_proto_data_server_proto_rawDescGZIP(), []int{2}
}

func (x *PushRequest) GetFragmentData() []byte {
	if x != nil {
		return x.FragmentData
	}
	return nil
}

type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FragmentId  string `protobuf:"bytes,1,opt,name=FragmentId,proto3" json:"FragmentId,omitempty"`
	ServiceName string `protobuf:"bytes,2,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	PushResult  bool   `protobuf:"varint,3,opt,name=PushResult,proto3" json:"PushResult,omitempty"`
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_server_proto_data_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_server_proto_data_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_data_server_proto_data_server_proto_rawDescGZIP(), []int{3}
}

func (x *PushResponse) GetFragmentId() string {
	if x != nil {
		return x.FragmentId
	}
	return ""
}

func (x *PushResponse) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *PushResponse) GetPushResult() bool {
	if x != nil {
		return x.PushResult
	}
	return false
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FragmentId string `protobuf:"bytes,1,opt,name=FragmentId,proto3" json:"FragmentId,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_server_proto_data_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_server_proto_data_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_data_server_proto_data_server_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetFragmentId() string {
	if x != nil {
		return x.FragmentId
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteResult bool `protobuf:"varint,1,opt,name=DeleteResult,proto3" json:"DeleteResult,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_server_proto_data_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_server_proto_data_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_data_server_proto_data_server_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteResponse) GetDeleteResult() bool {
	if x != nil {
		return x.DeleteResult
	}
	return false
}

var File_data_server_proto_data_server_proto protoreflect.FileDescriptor

var file_data_server_proto_data_server_proto_rawDesc = []byte{
	0x0a, 0x23, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x46, 0x72, 0x61, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x46, 0x72, 0x61,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x31, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x46, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x70, 0x0a, 0x0c, 0x50, 0x75,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2f, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x34, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0xf1, 0x01, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x1e, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_server_proto_data_server_proto_rawDescOnce sync.Once
	file_data_server_proto_data_server_proto_rawDescData = file_data_server_proto_data_server_proto_rawDesc
)

func file_data_server_proto_data_server_proto_rawDescGZIP() []byte {
	file_data_server_proto_data_server_proto_rawDescOnce.Do(func() {
		file_data_server_proto_data_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_server_proto_data_server_proto_rawDescData)
	})
	return file_data_server_proto_data_server_proto_rawDescData
}

var file_data_server_proto_data_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_data_server_proto_data_server_proto_goTypes = []interface{}{
	(*GetRequest)(nil),     // 0: data.server.proto.GetRequest
	(*GetResponse)(nil),    // 1: data.server.proto.GetResponse
	(*PushRequest)(nil),    // 2: data.server.proto.PushRequest
	(*PushResponse)(nil),   // 3: data.server.proto.PushResponse
	(*DeleteRequest)(nil),  // 4: data.server.proto.DeleteRequest
	(*DeleteResponse)(nil), // 5: data.server.proto.DeleteResponse
}
var file_data_server_proto_data_server_proto_depIdxs = []int32{
	2, // 0: data.server.proto.DataService.Push:input_type -> data.server.proto.PushRequest
	0, // 1: data.server.proto.DataService.Get:input_type -> data.server.proto.GetRequest
	4, // 2: data.server.proto.DataService.Delete:input_type -> data.server.proto.DeleteRequest
	3, // 3: data.server.proto.DataService.Push:output_type -> data.server.proto.PushResponse
	1, // 4: data.server.proto.DataService.Get:output_type -> data.server.proto.GetResponse
	5, // 5: data.server.proto.DataService.Delete:output_type -> data.server.proto.DeleteResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_server_proto_data_server_proto_init() }
func file_data_server_proto_data_server_proto_init() {
	if File_data_server_proto_data_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_server_proto_data_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_data_server_proto_data_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_data_server_proto_data_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_data_server_proto_data_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResponse); i {
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
		file_data_server_proto_data_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_data_server_proto_data_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_data_server_proto_data_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_server_proto_data_server_proto_goTypes,
		DependencyIndexes: file_data_server_proto_data_server_proto_depIdxs,
		MessageInfos:      file_data_server_proto_data_server_proto_msgTypes,
	}.Build()
	File_data_server_proto_data_server_proto = out.File
	file_data_server_proto_data_server_proto_rawDesc = nil
	file_data_server_proto_data_server_proto_goTypes = nil
	file_data_server_proto_data_server_proto_depIdxs = nil
}
