// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: opcua.proto

package opcuaProto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OpcuaMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NodeId    string `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Data      []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp uint32 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *OpcuaMessage) Reset() {
	*x = OpcuaMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opcua_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpcuaMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpcuaMessage) ProtoMessage() {}

func (x *OpcuaMessage) ProtoReflect() protoreflect.Message {
	mi := &file_opcua_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpcuaMessage.ProtoReflect.Descriptor instead.
func (*OpcuaMessage) Descriptor() ([]byte, []int) {
	return file_opcua_proto_rawDescGZIP(), []int{0}
}

func (x *OpcuaMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OpcuaMessage) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *OpcuaMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *OpcuaMessage) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type OpcuaResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OpcuaResult) Reset() {
	*x = OpcuaResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opcua_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpcuaResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpcuaResult) ProtoMessage() {}

func (x *OpcuaResult) ProtoReflect() protoreflect.Message {
	mi := &file_opcua_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpcuaResult.ProtoReflect.Descriptor instead.
func (*OpcuaResult) Descriptor() ([]byte, []int) {
	return file_opcua_proto_rawDescGZIP(), []int{1}
}

func (x *OpcuaResult) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type StreamScreen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream bool `protobuf:"varint,1,opt,name=stream,proto3" json:"stream,omitempty"`
}

func (x *StreamScreen) Reset() {
	*x = StreamScreen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opcua_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamScreen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamScreen) ProtoMessage() {}

func (x *StreamScreen) ProtoReflect() protoreflect.Message {
	mi := &file_opcua_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamScreen.ProtoReflect.Descriptor instead.
func (*StreamScreen) Descriptor() ([]byte, []int) {
	return file_opcua_proto_rawDescGZIP(), []int{2}
}

func (x *StreamScreen) GetStream() bool {
	if x != nil {
		return x.Stream
	}
	return false
}

type HelloScreen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc      string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	FrameRate uint32 `protobuf:"varint,3,opt,name=frame_rate,json=frameRate,proto3" json:"frame_rate,omitempty"`
}

func (x *HelloScreen) Reset() {
	*x = HelloScreen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opcua_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloScreen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloScreen) ProtoMessage() {}

func (x *HelloScreen) ProtoReflect() protoreflect.Message {
	mi := &file_opcua_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloScreen.ProtoReflect.Descriptor instead.
func (*HelloScreen) Descriptor() ([]byte, []int) {
	return file_opcua_proto_rawDescGZIP(), []int{3}
}

func (x *HelloScreen) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloScreen) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *HelloScreen) GetFrameRate() uint32 {
	if x != nil {
		return x.FrameRate
	}
	return 0
}

type Screen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*Screen_Data
	//	*Screen_Hello
	Payload isScreen_Payload `protobuf_oneof:"payload"`
}

func (x *Screen) Reset() {
	*x = Screen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opcua_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Screen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Screen) ProtoMessage() {}

func (x *Screen) ProtoReflect() protoreflect.Message {
	mi := &file_opcua_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Screen.ProtoReflect.Descriptor instead.
func (*Screen) Descriptor() ([]byte, []int) {
	return file_opcua_proto_rawDescGZIP(), []int{4}
}

func (m *Screen) GetPayload() isScreen_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Screen) GetData() []byte {
	if x, ok := x.GetPayload().(*Screen_Data); ok {
		return x.Data
	}
	return nil
}

func (x *Screen) GetHello() *HelloScreen {
	if x, ok := x.GetPayload().(*Screen_Hello); ok {
		return x.Hello
	}
	return nil
}

type isScreen_Payload interface {
	isScreen_Payload()
}

type Screen_Data struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type Screen_Hello struct {
	Hello *HelloScreen `protobuf:"bytes,2,opt,name=hello,proto3,oneof"`
}

func (*Screen_Data) isScreen_Payload() {}

func (*Screen_Hello) isScreen_Payload() {}

var File_opcua_proto protoreflect.FileDescriptor

var file_opcua_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x70, 0x63, 0x75, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61,
	0x70, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x0c, 0x4f, 0x70, 0x63, 0x75, 0x61, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x1d, 0x0a, 0x0b, 0x4f, 0x70, 0x63, 0x75, 0x61, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x26, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x54, 0x0a, 0x0b, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x22, 0x58, 0x0a, 0x06, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x2d, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x61, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x42,
	0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x87, 0x01, 0x0a, 0x05, 0x4f,
	0x70, 0x63, 0x75, 0x61, 0x12, 0x3e, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4f, 0x70,
	0x63, 0x75, 0x61, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x70, 0x63, 0x75, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x61, 0x70,
	0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x63, 0x75, 0x61, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x70, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x6f, 0x70, 0x63, 0x75, 0x61, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opcua_proto_rawDescOnce sync.Once
	file_opcua_proto_rawDescData = file_opcua_proto_rawDesc
)

func file_opcua_proto_rawDescGZIP() []byte {
	file_opcua_proto_rawDescOnce.Do(func() {
		file_opcua_proto_rawDescData = protoimpl.X.CompressGZIP(file_opcua_proto_rawDescData)
	})
	return file_opcua_proto_rawDescData
}

var file_opcua_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_opcua_proto_goTypes = []interface{}{
	(*OpcuaMessage)(nil), // 0: appProto.OpcuaMessage
	(*OpcuaResult)(nil),  // 1: appProto.OpcuaResult
	(*StreamScreen)(nil), // 2: appProto.StreamScreen
	(*HelloScreen)(nil),  // 3: appProto.HelloScreen
	(*Screen)(nil),       // 4: appProto.Screen
}
var file_opcua_proto_depIdxs = []int32{
	3, // 0: appProto.Screen.hello:type_name -> appProto.HelloScreen
	0, // 1: appProto.Opcua.ReportOpcua:input_type -> appProto.OpcuaMessage
	4, // 2: appProto.Opcua.RemoteScreen:input_type -> appProto.Screen
	1, // 3: appProto.Opcua.ReportOpcua:output_type -> appProto.OpcuaResult
	2, // 4: appProto.Opcua.RemoteScreen:output_type -> appProto.StreamScreen
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_opcua_proto_init() }
func file_opcua_proto_init() {
	if File_opcua_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opcua_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpcuaMessage); i {
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
		file_opcua_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpcuaResult); i {
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
		file_opcua_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamScreen); i {
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
		file_opcua_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloScreen); i {
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
		file_opcua_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Screen); i {
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
	file_opcua_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Screen_Data)(nil),
		(*Screen_Hello)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opcua_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_opcua_proto_goTypes,
		DependencyIndexes: file_opcua_proto_depIdxs,
		MessageInfos:      file_opcua_proto_msgTypes,
	}.Build()
	File_opcua_proto = out.File
	file_opcua_proto_rawDesc = nil
	file_opcua_proto_goTypes = nil
	file_opcua_proto_depIdxs = nil
}
