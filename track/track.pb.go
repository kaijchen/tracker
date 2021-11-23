// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: track/track.proto

package tracker

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Source   int64  `protobuf:"varint,3,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_track_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_track_track_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_track_track_proto_rawDescGZIP(), []int{0}
}

func (x *ReportRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReportRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ReportRequest) GetSource() int64 {
	if x != nil {
		return x.Source
	}
	return 0
}

type ReportReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ReportReply) Reset() {
	*x = ReportReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_track_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportReply) ProtoMessage() {}

func (x *ReportReply) ProtoReflect() protoreflect.Message {
	mi := &file_track_track_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportReply.ProtoReflect.Descriptor instead.
func (*ReportReply) Descriptor() ([]byte, []int) {
	return file_track_track_proto_rawDescGZIP(), []int{1}
}

func (x *ReportReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_track_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_track_track_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_track_track_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type QueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Source   int64  `protobuf:"varint,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *QueryReply) Reset() {
	*x = QueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_track_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryReply) ProtoMessage() {}

func (x *QueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_track_track_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryReply.ProtoReflect.Descriptor instead.
func (*QueryReply) Descriptor() ([]byte, []int) {
	return file_track_track_proto_rawDescGZIP(), []int{3}
}

func (x *QueryReply) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *QueryReply) GetSource() int64 {
	if x != nil {
		return x.Source
	}
	return 0
}

var File_track_track_proto protoreflect.FileDescriptor

var file_track_track_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x22, 0x55, 0x0a, 0x0d, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b,
	0x22, 0x20, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x40, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x32, 0x72, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12,
	0x34, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x13,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x69, 0x6a, 0x63, 0x68, 0x65, 0x6e, 0x2f,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_track_track_proto_rawDescOnce sync.Once
	file_track_track_proto_rawDescData = file_track_track_proto_rawDesc
)

func file_track_track_proto_rawDescGZIP() []byte {
	file_track_track_proto_rawDescOnce.Do(func() {
		file_track_track_proto_rawDescData = protoimpl.X.CompressGZIP(file_track_track_proto_rawDescData)
	})
	return file_track_track_proto_rawDescData
}

var file_track_track_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_track_track_proto_goTypes = []interface{}{
	(*ReportRequest)(nil), // 0: track.ReportRequest
	(*ReportReply)(nil),   // 1: track.ReportReply
	(*QueryRequest)(nil),  // 2: track.QueryRequest
	(*QueryReply)(nil),    // 3: track.QueryReply
}
var file_track_track_proto_depIdxs = []int32{
	0, // 0: track.Tracker.Report:input_type -> track.ReportRequest
	2, // 1: track.Tracker.Query:input_type -> track.QueryRequest
	1, // 2: track.Tracker.Report:output_type -> track.ReportReply
	3, // 3: track.Tracker.Query:output_type -> track.QueryReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_track_track_proto_init() }
func file_track_track_proto_init() {
	if File_track_track_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_track_track_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
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
		file_track_track_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportReply); i {
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
		file_track_track_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_track_track_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryReply); i {
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
			RawDescriptor: file_track_track_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_track_track_proto_goTypes,
		DependencyIndexes: file_track_track_proto_depIdxs,
		MessageInfos:      file_track_track_proto_msgTypes,
	}.Build()
	File_track_track_proto = out.File
	file_track_track_proto_rawDesc = nil
	file_track_track_proto_goTypes = nil
	file_track_track_proto_depIdxs = nil
}
