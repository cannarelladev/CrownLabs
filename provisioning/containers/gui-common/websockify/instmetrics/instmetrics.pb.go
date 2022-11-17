// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: instmetrics.proto

package instmetrics

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

type ContainerMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuPerc   float32 `protobuf:"fixed32,1,opt,name=cpu_perc,json=cpuPerc,proto3" json:"cpu_perc,omitempty"`
	MemBytes  uint64  `protobuf:"varint,2,opt,name=mem_bytes,json=memBytes,proto3" json:"mem_bytes,omitempty"`
	DiskBytes uint64  `protobuf:"varint,3,opt,name=disk_bytes,json=diskBytes,proto3" json:"disk_bytes,omitempty"`
}

func (x *ContainerMetricsResponse) Reset() {
	*x = ContainerMetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instmetrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerMetricsResponse) ProtoMessage() {}

func (x *ContainerMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_instmetrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerMetricsResponse.ProtoReflect.Descriptor instead.
func (*ContainerMetricsResponse) Descriptor() ([]byte, []int) {
	return file_instmetrics_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerMetricsResponse) GetCpuPerc() float32 {
	if x != nil {
		return x.CpuPerc
	}
	return 0
}

func (x *ContainerMetricsResponse) GetMemBytes() uint64 {
	if x != nil {
		return x.MemBytes
	}
	return 0
}

func (x *ContainerMetricsResponse) GetDiskBytes() uint64 {
	if x != nil {
		return x.DiskBytes
	}
	return 0
}

type ContainerMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Filter needed to find target "application container"
	PodName string `protobuf:"bytes,1,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
}

func (x *ContainerMetricsRequest) Reset() {
	*x = ContainerMetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instmetrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerMetricsRequest) ProtoMessage() {}

func (x *ContainerMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_instmetrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerMetricsRequest.ProtoReflect.Descriptor instead.
func (*ContainerMetricsRequest) Descriptor() ([]byte, []int) {
	return file_instmetrics_proto_rawDescGZIP(), []int{1}
}

func (x *ContainerMetricsRequest) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

var File_instmetrics_proto protoreflect.FileDescriptor

var file_instmetrics_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x22, 0x71, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x70, 0x75, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x63, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x74, 0x0a, 0x0f, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x61, 0x0a, 0x10,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x12, 0x24, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_instmetrics_proto_rawDescOnce sync.Once
	file_instmetrics_proto_rawDescData = file_instmetrics_proto_rawDesc
)

func file_instmetrics_proto_rawDescGZIP() []byte {
	file_instmetrics_proto_rawDescOnce.Do(func() {
		file_instmetrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_instmetrics_proto_rawDescData)
	})
	return file_instmetrics_proto_rawDescData
}

var file_instmetrics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_instmetrics_proto_goTypes = []interface{}{
	(*ContainerMetricsResponse)(nil), // 0: instmetrics.ContainerMetricsResponse
	(*ContainerMetricsRequest)(nil),  // 1: instmetrics.ContainerMetricsRequest
}
var file_instmetrics_proto_depIdxs = []int32{
	1, // 0: instmetrics.InstanceMetrics.ContainerMetrics:input_type -> instmetrics.ContainerMetricsRequest
	0, // 1: instmetrics.InstanceMetrics.ContainerMetrics:output_type -> instmetrics.ContainerMetricsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_instmetrics_proto_init() }
func file_instmetrics_proto_init() {
	if File_instmetrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_instmetrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerMetricsResponse); i {
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
		file_instmetrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerMetricsRequest); i {
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
			RawDescriptor: file_instmetrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_instmetrics_proto_goTypes,
		DependencyIndexes: file_instmetrics_proto_depIdxs,
		MessageInfos:      file_instmetrics_proto_msgTypes,
	}.Build()
	File_instmetrics_proto = out.File
	file_instmetrics_proto_rawDesc = nil
	file_instmetrics_proto_goTypes = nil
	file_instmetrics_proto_depIdxs = nil
}