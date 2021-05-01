// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: signup_v1.proto

package monitorApiv1

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

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version      int32 `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	ReportPeriod int32 `protobuf:"varint,2,opt,name=ReportPeriod,proto3" json:"ReportPeriod,omitempty"`
	MeanPeriod   int32 `protobuf:"varint,3,opt,name=MeanPeriod,proto3" json:"MeanPeriod,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signup_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_signup_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_signup_v1_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpRequest) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *SignUpRequest) GetReportPeriod() int32 {
	if x != nil {
		return x.ReportPeriod
	}
	return 0
}

func (x *SignUpRequest) GetMeanPeriod() int32 {
	if x != nil {
		return x.MeanPeriod
	}
	return 0
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *State `protobuf:"bytes,1,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signup_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_signup_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_signup_v1_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpResponse) GetState() *State {
	if x != nil {
		return x.State
	}
	return nil
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoadAverage *LoadAverage `protobuf:"bytes,1,opt,name=LoadAverage,proto3" json:"LoadAverage,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signup_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_signup_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_signup_v1_proto_rawDescGZIP(), []int{2}
}

func (x *State) GetLoadAverage() *LoadAverage {
	if x != nil {
		return x.LoadAverage
	}
	return nil
}

type LoadAverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	One     float64 `protobuf:"fixed64,1,opt,name=One,proto3" json:"One,omitempty"`
	Five    float64 `protobuf:"fixed64,2,opt,name=Five,proto3" json:"Five,omitempty"`
	Fifteen float64 `protobuf:"fixed64,3,opt,name=Fifteen,proto3" json:"Fifteen,omitempty"`
}

func (x *LoadAverage) Reset() {
	*x = LoadAverage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signup_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadAverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadAverage) ProtoMessage() {}

func (x *LoadAverage) ProtoReflect() protoreflect.Message {
	mi := &file_signup_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadAverage.ProtoReflect.Descriptor instead.
func (*LoadAverage) Descriptor() ([]byte, []int) {
	return file_signup_v1_proto_rawDescGZIP(), []int{3}
}

func (x *LoadAverage) GetOne() float64 {
	if x != nil {
		return x.One
	}
	return 0
}

func (x *LoadAverage) GetFive() float64 {
	if x != nil {
		return x.Five
	}
	return 0
}

func (x *LoadAverage) GetFifteen() float64 {
	if x != nil {
		return x.Fifteen
	}
	return 0
}

var File_signup_v1_proto protoreflect.FileDescriptor

var file_signup_v1_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x41, 0x70, 0x69, 0x76, 0x31, 0x22,
	0x6d, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x4d, 0x65, 0x61, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x4d, 0x65, 0x61, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x3b,
	0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x41, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x44, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x41, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x22, 0x4d, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x4f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x4f,
	0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x04, 0x46, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x66, 0x74, 0x65, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x46, 0x69, 0x66, 0x74, 0x65, 0x65, 0x6e,
	0x32, 0x56, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x1b, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x41, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x41, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x41, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_signup_v1_proto_rawDescOnce sync.Once
	file_signup_v1_proto_rawDescData = file_signup_v1_proto_rawDesc
)

func file_signup_v1_proto_rawDescGZIP() []byte {
	file_signup_v1_proto_rawDescOnce.Do(func() {
		file_signup_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_signup_v1_proto_rawDescData)
	})
	return file_signup_v1_proto_rawDescData
}

var file_signup_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_signup_v1_proto_goTypes = []interface{}{
	(*SignUpRequest)(nil),  // 0: monitorApiv1.SignUpRequest
	(*SignUpResponse)(nil), // 1: monitorApiv1.SignUpResponse
	(*State)(nil),          // 2: monitorApiv1.State
	(*LoadAverage)(nil),    // 3: monitorApiv1.LoadAverage
}
var file_signup_v1_proto_depIdxs = []int32{
	2, // 0: monitorApiv1.SignUpResponse.State:type_name -> monitorApiv1.State
	3, // 1: monitorApiv1.State.LoadAverage:type_name -> monitorApiv1.LoadAverage
	0, // 2: monitorApiv1.SignUpHandler.SignUp:input_type -> monitorApiv1.SignUpRequest
	1, // 3: monitorApiv1.SignUpHandler.SignUp:output_type -> monitorApiv1.SignUpResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_signup_v1_proto_init() }
func file_signup_v1_proto_init() {
	if File_signup_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_signup_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRequest); i {
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
		file_signup_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse); i {
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
		file_signup_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_signup_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadAverage); i {
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
			RawDescriptor: file_signup_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_signup_v1_proto_goTypes,
		DependencyIndexes: file_signup_v1_proto_depIdxs,
		MessageInfos:      file_signup_v1_proto_msgTypes,
	}.Build()
	File_signup_v1_proto = out.File
	file_signup_v1_proto_rawDesc = nil
	file_signup_v1_proto_goTypes = nil
	file_signup_v1_proto_depIdxs = nil
}