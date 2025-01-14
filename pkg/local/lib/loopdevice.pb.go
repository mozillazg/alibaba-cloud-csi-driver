// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: pkg/local/lib/loopdevice.proto

package lib

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

type CreateLoopDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PvName    string `protobuf:"bytes,1,opt,name=pv_name,json=pvName,proto3" json:"pv_name,omitempty"`
	QuotaSize string `protobuf:"bytes,2,opt,name=quota_size,json=quotaSize,proto3" json:"quota_size,omitempty"`
	RootPath  string `protobuf:"bytes,3,opt,name=root_path,json=rootPath,proto3" json:"root_path,omitempty"`
}

func (x *CreateLoopDeviceRequest) Reset() {
	*x = CreateLoopDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_local_lib_loopdevice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoopDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoopDeviceRequest) ProtoMessage() {}

func (x *CreateLoopDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_local_lib_loopdevice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoopDeviceRequest.ProtoReflect.Descriptor instead.
func (*CreateLoopDeviceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_local_lib_loopdevice_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLoopDeviceRequest) GetPvName() string {
	if x != nil {
		return x.PvName
	}
	return ""
}

func (x *CreateLoopDeviceRequest) GetQuotaSize() string {
	if x != nil {
		return x.QuotaSize
	}
	return ""
}

func (x *CreateLoopDeviceRequest) GetRootPath() string {
	if x != nil {
		return x.RootPath
	}
	return ""
}

type CreateLoopDeviceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoopDevicePath string `protobuf:"bytes,1,opt,name=loop_device_path,json=loopDevicePath,proto3" json:"loop_device_path,omitempty"`
	CommandOutput  string `protobuf:"bytes,2,opt,name=command_output,json=commandOutput,proto3" json:"command_output,omitempty"`
}

func (x *CreateLoopDeviceReply) Reset() {
	*x = CreateLoopDeviceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_local_lib_loopdevice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoopDeviceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoopDeviceReply) ProtoMessage() {}

func (x *CreateLoopDeviceReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_local_lib_loopdevice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoopDeviceReply.ProtoReflect.Descriptor instead.
func (*CreateLoopDeviceReply) Descriptor() ([]byte, []int) {
	return file_pkg_local_lib_loopdevice_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLoopDeviceReply) GetLoopDevicePath() string {
	if x != nil {
		return x.LoopDevicePath
	}
	return ""
}

func (x *CreateLoopDeviceReply) GetCommandOutput() string {
	if x != nil {
		return x.CommandOutput
	}
	return ""
}

type DeleteLoopDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PvName string `protobuf:"bytes,1,opt,name=pv_name,json=pvName,proto3" json:"pv_name,omitempty"`
}

func (x *DeleteLoopDeviceRequest) Reset() {
	*x = DeleteLoopDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_local_lib_loopdevice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLoopDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLoopDeviceRequest) ProtoMessage() {}

func (x *DeleteLoopDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_local_lib_loopdevice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLoopDeviceRequest.ProtoReflect.Descriptor instead.
func (*DeleteLoopDeviceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_local_lib_loopdevice_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteLoopDeviceRequest) GetPvName() string {
	if x != nil {
		return x.PvName
	}
	return ""
}

type DeleteLoopDeviceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandOutput string `protobuf:"bytes,1,opt,name=command_output,json=commandOutput,proto3" json:"command_output,omitempty"`
}

func (x *DeleteLoopDeviceReply) Reset() {
	*x = DeleteLoopDeviceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_local_lib_loopdevice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLoopDeviceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLoopDeviceReply) ProtoMessage() {}

func (x *DeleteLoopDeviceReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_local_lib_loopdevice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLoopDeviceReply.ProtoReflect.Descriptor instead.
func (*DeleteLoopDeviceReply) Descriptor() ([]byte, []int) {
	return file_pkg_local_lib_loopdevice_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteLoopDeviceReply) GetCommandOutput() string {
	if x != nil {
		return x.CommandOutput
	}
	return ""
}

var File_pkg_local_lib_loopdevice_proto protoreflect.FileDescriptor

var file_pkg_local_lib_loopdevice_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2f, 0x6c, 0x69, 0x62, 0x2f,
	0x6c, 0x6f, 0x6f, 0x70, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x6f, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x76, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f,
	0x6f, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x6f, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x68, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x6f, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x28, 0x0a, 0x10, 0x6c, 0x6f, 0x6f, 0x70, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x6f, 0x70,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x32, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x70, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x76, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x76, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x6f, 0x6f, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0xb4, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x6f, 0x70, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x6f, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x70, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x6f, 0x70, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x3b, 0x6c, 0x69, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_local_lib_loopdevice_proto_rawDescOnce sync.Once
	file_pkg_local_lib_loopdevice_proto_rawDescData = file_pkg_local_lib_loopdevice_proto_rawDesc
)

func file_pkg_local_lib_loopdevice_proto_rawDescGZIP() []byte {
	file_pkg_local_lib_loopdevice_proto_rawDescOnce.Do(func() {
		file_pkg_local_lib_loopdevice_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_local_lib_loopdevice_proto_rawDescData)
	})
	return file_pkg_local_lib_loopdevice_proto_rawDescData
}

var file_pkg_local_lib_loopdevice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_local_lib_loopdevice_proto_goTypes = []interface{}{
	(*CreateLoopDeviceRequest)(nil), // 0: proto.CreateLoopDeviceRequest
	(*CreateLoopDeviceReply)(nil),   // 1: proto.CreateLoopDeviceReply
	(*DeleteLoopDeviceRequest)(nil), // 2: proto.DeleteLoopDeviceRequest
	(*DeleteLoopDeviceReply)(nil),   // 3: proto.DeleteLoopDeviceReply
}
var file_pkg_local_lib_loopdevice_proto_depIdxs = []int32{
	0, // 0: proto.LoopDevice.CreateLoopDevice:input_type -> proto.CreateLoopDeviceRequest
	2, // 1: proto.LoopDevice.DeleteLoopDevice:input_type -> proto.DeleteLoopDeviceRequest
	1, // 2: proto.LoopDevice.CreateLoopDevice:output_type -> proto.CreateLoopDeviceReply
	3, // 3: proto.LoopDevice.DeleteLoopDevice:output_type -> proto.DeleteLoopDeviceReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_local_lib_loopdevice_proto_init() }
func file_pkg_local_lib_loopdevice_proto_init() {
	if File_pkg_local_lib_loopdevice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_local_lib_loopdevice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoopDeviceRequest); i {
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
		file_pkg_local_lib_loopdevice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoopDeviceReply); i {
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
		file_pkg_local_lib_loopdevice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLoopDeviceRequest); i {
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
		file_pkg_local_lib_loopdevice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLoopDeviceReply); i {
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
			RawDescriptor: file_pkg_local_lib_loopdevice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_local_lib_loopdevice_proto_goTypes,
		DependencyIndexes: file_pkg_local_lib_loopdevice_proto_depIdxs,
		MessageInfos:      file_pkg_local_lib_loopdevice_proto_msgTypes,
	}.Build()
	File_pkg_local_lib_loopdevice_proto = out.File
	file_pkg_local_lib_loopdevice_proto_rawDesc = nil
	file_pkg_local_lib_loopdevice_proto_goTypes = nil
	file_pkg_local_lib_loopdevice_proto_depIdxs = nil
}
