// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: signup_v1.proto

package monitorApiv1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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
	CPULoad     *CPULoad     `protobuf:"bytes,2,opt,name=CPULoad,proto3" json:"CPULoad,omitempty"`
	Mem         *Mem         `protobuf:"bytes,3,opt,name=Mem,proto3" json:"Mem,omitempty"`
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

func (x *State) GetCPULoad() *CPULoad {
	if x != nil {
		return x.CPULoad
	}
	return nil
}

func (x *State) GetMem() *Mem {
	if x != nil {
		return x.Mem
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

type CPULoad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   float64 `protobuf:"fixed64,1,opt,name=User,proto3" json:"User,omitempty"`
	System float64 `protobuf:"fixed64,2,opt,name=System,proto3" json:"System,omitempty"`
	Idle   float64 `protobuf:"fixed64,3,opt,name=Idle,proto3" json:"Idle,omitempty"`
}

func (x *CPULoad) Reset() {
	*x = CPULoad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signup_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPULoad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPULoad) ProtoMessage() {}

func (x *CPULoad) ProtoReflect() protoreflect.Message {
	mi := &file_signup_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPULoad.ProtoReflect.Descriptor instead.
func (*CPULoad) Descriptor() ([]byte, []int) {
	return file_signup_v1_proto_rawDescGZIP(), []int{4}
}

func (x *CPULoad) GetUser() float64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *CPULoad) GetSystem() float64 {
	if x != nil {
		return x.System
	}
	return 0
}

func (x *CPULoad) GetIdle() float64 {
	if x != nil {
		return x.Idle
	}
	return 0
}

type Mem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total float64 `protobuf:"fixed64,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Free  float64 `protobuf:"fixed64,2,opt,name=Free,proto3" json:"Free,omitempty"`
	Used  float64 `protobuf:"fixed64,3,opt,name=Used,proto3" json:"Used,omitempty"`
}

func (x *Mem) Reset() {
	*x = Mem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signup_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mem) ProtoMessage() {}

func (x *Mem) ProtoReflect() protoreflect.Message {
	mi := &file_signup_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mem.ProtoReflect.Descriptor instead.
func (*Mem) Descriptor() ([]byte, []int) {
	return file_signup_v1_proto_rawDescGZIP(), []int{5}
}

func (x *Mem) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Mem) GetFree() float64 {
	if x != nil {
		return x.Free
	}
	return 0
}

func (x *Mem) GetUsed() float64 {
	if x != nil {
		return x.Used
	}
	return 0
}

var File_signup_v1_proto protoreflect.FileDescriptor

var file_signup_v1_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x41, 0x70, 0x69, 0x76, 0x31, 0x22,
	0x53, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x65, 0x61, 0x6e, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4d, 0x65, 0x61, 0x6e, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x22, 0x3b, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x41,
	0x70, 0x69, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x9a, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x4c,
	0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x41, 0x70, 0x69, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x4c, 0x6f, 0x61,
	0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x43, 0x50, 0x55, 0x4c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x41, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x43, 0x50, 0x55, 0x4c, 0x6f, 0x61, 0x64,
	0x52, 0x07, 0x43, 0x50, 0x55, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x03, 0x4d, 0x65, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x41, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x52, 0x03, 0x4d, 0x65, 0x6d, 0x22, 0x4d,
	0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x4f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x4f, 0x6e, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x46, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x46,
	0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x66, 0x74, 0x65, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x46, 0x69, 0x66, 0x74, 0x65, 0x65, 0x6e, 0x22, 0x49, 0x0a,
	0x07, 0x43, 0x50, 0x55, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x64, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x49, 0x64, 0x6c, 0x65, 0x22, 0x43, 0x0a, 0x03, 0x4d, 0x65, 0x6d, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x65, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x46, 0x72, 0x65, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x55, 0x73, 0x65, 0x64, 0x32, 0x56, 0x0a,
	0x0d, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x45,
	0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x1b, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x41, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x41,
	0x70, 0x69, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x41, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var (
	file_signup_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
	file_signup_v1_proto_goTypes  = []interface{}{
		(*SignUpRequest)(nil),  // 0: monitorApiv1.SignUpRequest
		(*SignUpResponse)(nil), // 1: monitorApiv1.SignUpResponse
		(*State)(nil),          // 2: monitorApiv1.State
		(*LoadAverage)(nil),    // 3: monitorApiv1.LoadAverage
		(*CPULoad)(nil),        // 4: monitorApiv1.CPULoad
		(*Mem)(nil),            // 5: monitorApiv1.Mem
	}
)

var file_signup_v1_proto_depIdxs = []int32{
	2, // 0: monitorApiv1.SignUpResponse.State:type_name -> monitorApiv1.State
	3, // 1: monitorApiv1.State.LoadAverage:type_name -> monitorApiv1.LoadAverage
	4, // 2: monitorApiv1.State.CPULoad:type_name -> monitorApiv1.CPULoad
	5, // 3: monitorApiv1.State.Mem:type_name -> monitorApiv1.Mem
	0, // 4: monitorApiv1.SignUpHandler.SignUp:input_type -> monitorApiv1.SignUpRequest
	1, // 5: monitorApiv1.SignUpHandler.SignUp:output_type -> monitorApiv1.SignUpResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
		file_signup_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CPULoad); i {
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
		file_signup_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mem); i {
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
			NumMessages:   6,
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
