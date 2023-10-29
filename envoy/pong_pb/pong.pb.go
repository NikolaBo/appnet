// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: pong_pb/pong.proto

package pong_pb

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

type PongHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PongHelloRequest) Reset() {
	*x = PongHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_pb_pong_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongHelloRequest) ProtoMessage() {}

func (x *PongHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pong_pb_pong_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongHelloRequest.ProtoReflect.Descriptor instead.
func (*PongHelloRequest) Descriptor() ([]byte, []int) {
	return file_pong_pb_pong_proto_rawDescGZIP(), []int{0}
}

func (x *PongHelloRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type PongHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PongHelloResponse) Reset() {
	*x = PongHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_pb_pong_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongHelloResponse) ProtoMessage() {}

func (x *PongHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pong_pb_pong_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongHelloResponse.ProtoReflect.Descriptor instead.
func (*PongHelloResponse) Descriptor() ([]byte, []int) {
	return file_pong_pb_pong_proto_rawDescGZIP(), []int{1}
}

func (x *PongHelloResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type PongWorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PongWorldRequest) Reset() {
	*x = PongWorldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_pb_pong_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongWorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongWorldRequest) ProtoMessage() {}

func (x *PongWorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pong_pb_pong_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongWorldRequest.ProtoReflect.Descriptor instead.
func (*PongWorldRequest) Descriptor() ([]byte, []int) {
	return file_pong_pb_pong_proto_rawDescGZIP(), []int{2}
}

func (x *PongWorldRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type PongWorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PongWorldResponse) Reset() {
	*x = PongWorldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_pb_pong_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongWorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongWorldResponse) ProtoMessage() {}

func (x *PongWorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pong_pb_pong_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongWorldResponse.ProtoReflect.Descriptor instead.
func (*PongWorldResponse) Descriptor() ([]byte, []int) {
	return file_pong_pb_pong_proto_rawDescGZIP(), []int{3}
}

func (x *PongWorldResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_pong_pb_pong_proto protoreflect.FileDescriptor

var file_pong_pb_pong_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6f, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x22, 0x26, 0x0a,
	0x10, 0x50, 0x6f, 0x6e, 0x67, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x27, 0x0a, 0x11, 0x50, 0x6f, 0x6e, 0x67, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x26,
	0x0a, 0x10, 0x50, 0x6f, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x27, 0x0a, 0x11, 0x50, 0x6f, 0x6e, 0x67, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32,
	0x95, 0x01, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x42, 0x0a, 0x09, 0x50, 0x6f, 0x6e, 0x67, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x19, 0x2e, 0x70,
	0x6f, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x6e, 0x67, 0x5f, 0x70,
	0x62, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x50, 0x6f, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x12, 0x19, 0x2e, 0x70, 0x6f, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f,
	0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x6f, 0x6e,
	0x67, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pong_pb_pong_proto_rawDescOnce sync.Once
	file_pong_pb_pong_proto_rawDescData = file_pong_pb_pong_proto_rawDesc
)

func file_pong_pb_pong_proto_rawDescGZIP() []byte {
	file_pong_pb_pong_proto_rawDescOnce.Do(func() {
		file_pong_pb_pong_proto_rawDescData = protoimpl.X.CompressGZIP(file_pong_pb_pong_proto_rawDescData)
	})
	return file_pong_pb_pong_proto_rawDescData
}

var file_pong_pb_pong_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pong_pb_pong_proto_goTypes = []interface{}{
	(*PongHelloRequest)(nil),  // 0: pong_pb.PongHelloRequest
	(*PongHelloResponse)(nil), // 1: pong_pb.PongHelloResponse
	(*PongWorldRequest)(nil),  // 2: pong_pb.PongWorldRequest
	(*PongWorldResponse)(nil), // 3: pong_pb.PongWorldResponse
}
var file_pong_pb_pong_proto_depIdxs = []int32{
	0, // 0: pong_pb.EchoService.PongHello:input_type -> pong_pb.PongHelloRequest
	2, // 1: pong_pb.EchoService.PongWorld:input_type -> pong_pb.PongWorldRequest
	1, // 2: pong_pb.EchoService.PongHello:output_type -> pong_pb.PongHelloResponse
	3, // 3: pong_pb.EchoService.PongWorld:output_type -> pong_pb.PongWorldResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pong_pb_pong_proto_init() }
func file_pong_pb_pong_proto_init() {
	if File_pong_pb_pong_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pong_pb_pong_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongHelloRequest); i {
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
		file_pong_pb_pong_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongHelloResponse); i {
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
		file_pong_pb_pong_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongWorldRequest); i {
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
		file_pong_pb_pong_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongWorldResponse); i {
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
			RawDescriptor: file_pong_pb_pong_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pong_pb_pong_proto_goTypes,
		DependencyIndexes: file_pong_pb_pong_proto_depIdxs,
		MessageInfos:      file_pong_pb_pong_proto_msgTypes,
	}.Build()
	File_pong_pb_pong_proto = out.File
	file_pong_pb_pong_proto_rawDesc = nil
	file_pong_pb_pong_proto_goTypes = nil
	file_pong_pb_pong_proto_depIdxs = nil
}
