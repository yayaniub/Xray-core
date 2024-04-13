// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: proxy/vless/encoding/addons.proto

package encoding

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

type SeedMode int32

const (
	SeedMode_Unknown              SeedMode = 0
	SeedMode_PaddingOnly          SeedMode = 1
	SeedMode_PaddingPlusDelay     SeedMode = 2
	SeedMode_IndependentScheduler SeedMode = 3
)

// Enum value maps for SeedMode.
var (
	SeedMode_name = map[int32]string{
		0: "Unknown",
		1: "PaddingOnly",
		2: "PaddingPlusDelay",
		3: "IndependentScheduler",
	}
	SeedMode_value = map[string]int32{
		"Unknown":              0,
		"PaddingOnly":          1,
		"PaddingPlusDelay":     2,
		"IndependentScheduler": 3,
	}
)

func (x SeedMode) Enum() *SeedMode {
	p := new(SeedMode)
	*p = x
	return p
}

func (x SeedMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SeedMode) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_vless_encoding_addons_proto_enumTypes[0].Descriptor()
}

func (SeedMode) Type() protoreflect.EnumType {
	return &file_proxy_vless_encoding_addons_proto_enumTypes[0]
}

func (x SeedMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SeedMode.Descriptor instead.
func (SeedMode) EnumDescriptor() ([]byte, []int) {
	return file_proxy_vless_encoding_addons_proto_rawDescGZIP(), []int{0}
}

type Addons struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flow            string           `protobuf:"bytes,1,opt,name=Flow,proto3" json:"Flow,omitempty"`
	Seed            []byte           `protobuf:"bytes,2,opt,name=Seed,proto3" json:"Seed,omitempty"`
	Mode            SeedMode         `protobuf:"varint,3,opt,name=Mode,proto3,enum=xray.proxy.vless.encoding.SeedMode" json:"Mode,omitempty"`
	Duration        string           `protobuf:"bytes,4,opt,name=Duration,proto3" json:"Duration,omitempty"` // "0-8" means apply to number of packets, "1kb-" means start applying once both side exchange 1kb data, counting two-ways
	Padding         *PaddingConfig   `protobuf:"bytes,5,opt,name=Padding,proto3" json:"Padding,omitempty"`
	Delay           *DelayConfig     `protobuf:"bytes,6,opt,name=Delay,proto3" json:"Delay,omitempty"`
	SchedulerConfig *SchedulerConfig `protobuf:"bytes,7,opt,name=SchedulerConfig,proto3" json:"SchedulerConfig,omitempty"`
}

func (x *Addons) Reset() {
	*x = Addons{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_vless_encoding_addons_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Addons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Addons) ProtoMessage() {}

func (x *Addons) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_vless_encoding_addons_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Addons.ProtoReflect.Descriptor instead.
func (*Addons) Descriptor() ([]byte, []int) {
	return file_proxy_vless_encoding_addons_proto_rawDescGZIP(), []int{0}
}

func (x *Addons) GetFlow() string {
	if x != nil {
		return x.Flow
	}
	return ""
}

func (x *Addons) GetSeed() []byte {
	if x != nil {
		return x.Seed
	}
	return nil
}

func (x *Addons) GetMode() SeedMode {
	if x != nil {
		return x.Mode
	}
	return SeedMode_Unknown
}

func (x *Addons) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Addons) GetPadding() *PaddingConfig {
	if x != nil {
		return x.Padding
	}
	return nil
}

func (x *Addons) GetDelay() *DelayConfig {
	if x != nil {
		return x.Delay
	}
	return nil
}

func (x *Addons) GetSchedulerConfig() *SchedulerConfig {
	if x != nil {
		return x.SchedulerConfig
	}
	return nil
}

type PaddingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegularMin uint32 `protobuf:"varint,1,opt,name=RegularMin,proto3" json:"RegularMin,omitempty"`
	RegularMax uint32 `protobuf:"varint,2,opt,name=RegularMax,proto3" json:"RegularMax,omitempty"`
	LongMin    uint32 `protobuf:"varint,3,opt,name=LongMin,proto3" json:"LongMin,omitempty"`
	LongMax    uint32 `protobuf:"varint,4,opt,name=LongMax,proto3" json:"LongMax,omitempty"`
}

func (x *PaddingConfig) Reset() {
	*x = PaddingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_vless_encoding_addons_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaddingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaddingConfig) ProtoMessage() {}

func (x *PaddingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_vless_encoding_addons_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaddingConfig.ProtoReflect.Descriptor instead.
func (*PaddingConfig) Descriptor() ([]byte, []int) {
	return file_proxy_vless_encoding_addons_proto_rawDescGZIP(), []int{1}
}

func (x *PaddingConfig) GetRegularMin() uint32 {
	if x != nil {
		return x.RegularMin
	}
	return 0
}

func (x *PaddingConfig) GetRegularMax() uint32 {
	if x != nil {
		return x.RegularMax
	}
	return 0
}

func (x *PaddingConfig) GetLongMin() uint32 {
	if x != nil {
		return x.LongMin
	}
	return 0
}

func (x *PaddingConfig) GetLongMax() uint32 {
	if x != nil {
		return x.LongMax
	}
	return 0
}

type DelayConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsRandom  bool   `protobuf:"varint,1,opt,name=IsRandom,proto3" json:"IsRandom,omitempty"`
	MinMillis uint32 `protobuf:"varint,2,opt,name=MinMillis,proto3" json:"MinMillis,omitempty"`
	MaxMillis uint32 `protobuf:"varint,3,opt,name=MaxMillis,proto3" json:"MaxMillis,omitempty"`
}

func (x *DelayConfig) Reset() {
	*x = DelayConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_vless_encoding_addons_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelayConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelayConfig) ProtoMessage() {}

func (x *DelayConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_vless_encoding_addons_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelayConfig.ProtoReflect.Descriptor instead.
func (*DelayConfig) Descriptor() ([]byte, []int) {
	return file_proxy_vless_encoding_addons_proto_rawDescGZIP(), []int{2}
}

func (x *DelayConfig) GetIsRandom() bool {
	if x != nil {
		return x.IsRandom
	}
	return false
}

func (x *DelayConfig) GetMinMillis() uint32 {
	if x != nil {
		return x.MinMillis
	}
	return 0
}

func (x *DelayConfig) GetMaxMillis() uint32 {
	if x != nil {
		return x.MaxMillis
	}
	return 0
}

type SchedulerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeoutMillis uint32 `protobuf:"varint,1,opt,name=TimeoutMillis,proto3" json:"TimeoutMillis,omitempty"` // original traffic will not be sent right away but when scheduler want to send or pending buffer times out
}

func (x *SchedulerConfig) Reset() {
	*x = SchedulerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_vless_encoding_addons_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerConfig) ProtoMessage() {}

func (x *SchedulerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_vless_encoding_addons_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerConfig.ProtoReflect.Descriptor instead.
func (*SchedulerConfig) Descriptor() ([]byte, []int) {
	return file_proxy_vless_encoding_addons_proto_rawDescGZIP(), []int{3}
}

func (x *SchedulerConfig) GetTimeoutMillis() uint32 {
	if x != nil {
		return x.TimeoutMillis
	}
	return 0
}

var File_proxy_vless_encoding_addons_proto protoreflect.FileDescriptor

var file_proxy_vless_encoding_addons_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x19, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x22, 0xdd,
	0x02, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x6c, 0x6f,
	0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x12, 0x0a,
	0x04, 0x53, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x53, 0x65, 0x65,
	0x64, 0x12, 0x37, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x23, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x6c, 0x65,
	0x73, 0x73, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x65, 0x64,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x07, 0x50, 0x61, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x07, 0x50, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x3c, 0x0a, 0x05, 0x44, 0x65,
	0x6c, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x78, 0x72, 0x61, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x05, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x54, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76,
	0x6c, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x83,
	0x01, 0x0a, 0x0d, 0x50, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x69, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x61, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x61, 0x78,
	0x12, 0x18, 0x0a, 0x07, 0x4c, 0x6f, 0x6e, 0x67, 0x4d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x4c, 0x6f, 0x6e, 0x67, 0x4d, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x6f,
	0x6e, 0x67, 0x4d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x4c, 0x6f, 0x6e,
	0x67, 0x4d, 0x61, 0x78, 0x22, 0x65, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12,
	0x1c, 0x0a, 0x09, 0x4d, 0x69, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x4d, 0x69, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x4d, 0x61, 0x78, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x4d, 0x61, 0x78, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x22, 0x37, 0x0a, 0x0f, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x24,
	0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x69,
	0x6c, 0x6c, 0x69, 0x73, 0x2a, 0x58, 0x0a, 0x08, 0x53, 0x65, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x50, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x6c, 0x79, 0x10, 0x01, 0x12, 0x14,
	0x0a, 0x10, 0x50, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x75, 0x73, 0x44, 0x65, 0x6c,
	0x61, 0x79, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x6e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x10, 0x03, 0x42, 0x6d,
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x50,
	0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x74,
	0x6c, 0x73, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0xaa, 0x02, 0x19, 0x58, 0x72, 0x61, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x56,
	0x6c, 0x65, 0x73, 0x73, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_vless_encoding_addons_proto_rawDescOnce sync.Once
	file_proxy_vless_encoding_addons_proto_rawDescData = file_proxy_vless_encoding_addons_proto_rawDesc
)

func file_proxy_vless_encoding_addons_proto_rawDescGZIP() []byte {
	file_proxy_vless_encoding_addons_proto_rawDescOnce.Do(func() {
		file_proxy_vless_encoding_addons_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_vless_encoding_addons_proto_rawDescData)
	})
	return file_proxy_vless_encoding_addons_proto_rawDescData
}

var file_proxy_vless_encoding_addons_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proxy_vless_encoding_addons_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proxy_vless_encoding_addons_proto_goTypes = []interface{}{
	(SeedMode)(0),           // 0: xray.proxy.vless.encoding.SeedMode
	(*Addons)(nil),          // 1: xray.proxy.vless.encoding.Addons
	(*PaddingConfig)(nil),   // 2: xray.proxy.vless.encoding.PaddingConfig
	(*DelayConfig)(nil),     // 3: xray.proxy.vless.encoding.DelayConfig
	(*SchedulerConfig)(nil), // 4: xray.proxy.vless.encoding.SchedulerConfig
}
var file_proxy_vless_encoding_addons_proto_depIdxs = []int32{
	0, // 0: xray.proxy.vless.encoding.Addons.Mode:type_name -> xray.proxy.vless.encoding.SeedMode
	2, // 1: xray.proxy.vless.encoding.Addons.Padding:type_name -> xray.proxy.vless.encoding.PaddingConfig
	3, // 2: xray.proxy.vless.encoding.Addons.Delay:type_name -> xray.proxy.vless.encoding.DelayConfig
	4, // 3: xray.proxy.vless.encoding.Addons.SchedulerConfig:type_name -> xray.proxy.vless.encoding.SchedulerConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proxy_vless_encoding_addons_proto_init() }
func file_proxy_vless_encoding_addons_proto_init() {
	if File_proxy_vless_encoding_addons_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_vless_encoding_addons_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Addons); i {
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
		file_proxy_vless_encoding_addons_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaddingConfig); i {
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
		file_proxy_vless_encoding_addons_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelayConfig); i {
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
		file_proxy_vless_encoding_addons_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerConfig); i {
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
			RawDescriptor: file_proxy_vless_encoding_addons_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_vless_encoding_addons_proto_goTypes,
		DependencyIndexes: file_proxy_vless_encoding_addons_proto_depIdxs,
		EnumInfos:         file_proxy_vless_encoding_addons_proto_enumTypes,
		MessageInfos:      file_proxy_vless_encoding_addons_proto_msgTypes,
	}.Build()
	File_proxy_vless_encoding_addons_proto = out.File
	file_proxy_vless_encoding_addons_proto_rawDesc = nil
	file_proxy_vless_encoding_addons_proto_goTypes = nil
	file_proxy_vless_encoding_addons_proto_depIdxs = nil
}
