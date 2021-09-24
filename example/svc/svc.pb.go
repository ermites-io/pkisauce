// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: svc.proto

package svc

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  int64  `protobuf:"varint,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Args string `protobuf:"bytes,2,opt,name=args,proto3" json:"args,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_svc_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetCmd() int64 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *Request) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret int64  `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_svc_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetRet() int64 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_svc_proto protoreflect.FileDescriptor

var file_svc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0x2f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72,
	0x67, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x65, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x32, 0x5c, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x29, 0x0a, 0x06,
	0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x65, 0x12, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x32, 0x67, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x12, 0x2b,
	0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x65, 0x72, 0x65, 0x12, 0x0d, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x09, 0x43,
	0x61, 0x6c, 0x6c, 0x54, 0x68, 0x65, 0x72, 0x65, 0x12, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x67, 0x0a, 0x09, 0x44, 0x69, 0x66,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x69,
	0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x05, 0x53, 0x68, 0x61, 0x7a,
	0x7a, 0x12, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x73, 0x76, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_proto_rawDescOnce sync.Once
	file_svc_proto_rawDescData = file_svc_proto_rawDesc
)

func file_svc_proto_rawDescGZIP() []byte {
	file_svc_proto_rawDescOnce.Do(func() {
		file_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_proto_rawDescData)
	})
	return file_svc_proto_rawDescData
}

var file_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_svc_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: main.Request
	(*Response)(nil), // 1: main.Response
}
var file_svc_proto_depIdxs = []int32{
	0, // 0: main.Public.CallMe:input_type -> main.Request
	0, // 1: main.Public.Ping:input_type -> main.Request
	0, // 2: main.Restricted.CallHere:input_type -> main.Request
	0, // 3: main.Restricted.CallThere:input_type -> main.Request
	0, // 4: main.Different.CallDifferent:input_type -> main.Request
	0, // 5: main.Different.Shazz:input_type -> main.Request
	1, // 6: main.Public.CallMe:output_type -> main.Response
	1, // 7: main.Public.Ping:output_type -> main.Response
	1, // 8: main.Restricted.CallHere:output_type -> main.Response
	1, // 9: main.Restricted.CallThere:output_type -> main.Response
	1, // 10: main.Different.CallDifferent:output_type -> main.Response
	1, // 11: main.Different.Shazz:output_type -> main.Response
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svc_proto_init() }
func file_svc_proto_init() {
	if File_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_svc_proto_goTypes,
		DependencyIndexes: file_svc_proto_depIdxs,
		MessageInfos:      file_svc_proto_msgTypes,
	}.Build()
	File_svc_proto = out.File
	file_svc_proto_rawDesc = nil
	file_svc_proto_goTypes = nil
	file_svc_proto_depIdxs = nil
}
