// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: test.proto

package testing

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Enum int32

const (
	Enum_V1 Enum = 0
	Enum_V2 Enum = 1
	Enum_V3 Enum = 2
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "V1",
		1: "V2",
		2: "V3",
	}
	Enum_value = map[string]int32{
		"V1": 0,
		"V2": 1,
		"V3": 2,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *Object) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Object) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Dummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enum            Enum                 `protobuf:"varint,1,opt,name=enum,proto3,enum=authorization.test.Enum" json:"enum,omitempty"`
	RepeatedStr     []string             `protobuf:"bytes,2,rep,name=repeated_str,json=repeatedStr,proto3" json:"repeated_str,omitempty"`
	MapStrObject    map[string]*Object   `protobuf:"bytes,3,rep,name=map_str_object,json=mapStrObject,proto3" json:"map_str_object,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapIntObject    map[int64]*Object    `protobuf:"bytes,4,rep,name=map_int_object,json=mapIntObject,proto3" json:"map_int_object,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapBoolObject   map[bool]*Object     `protobuf:"bytes,5,rep,name=map_bool_object,json=mapBoolObject,proto3" json:"map_bool_object,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Timestamp       *timestamp.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	RepeatedObjects []*Object            `protobuf:"bytes,7,rep,name=repeated_objects,json=repeatedObjects,proto3" json:"repeated_objects,omitempty"`
	SingleObject    *Object              `protobuf:"bytes,8,opt,name=single_object,json=singleObject,proto3" json:"single_object,omitempty"`
	BooleanTrue     bool                 `protobuf:"varint,9,opt,name=boolean_true,json=booleanTrue,proto3" json:"boolean_true,omitempty"`
	BooleanFalse    bool                 `protobuf:"varint,10,opt,name=boolean_false,json=booleanFalse,proto3" json:"boolean_false,omitempty"`
	Uint32          uint32               `protobuf:"varint,11,opt,name=uint32,proto3" json:"uint32,omitempty"`
	Uint64          uint64               `protobuf:"varint,12,opt,name=uint64,proto3" json:"uint64,omitempty"`
	Bytes           []byte               `protobuf:"bytes,13,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Sint32          int32                `protobuf:"zigzag32,14,opt,name=sint32,proto3" json:"sint32,omitempty"`
	// unsupported fields bellow:
	Double   float64 `protobuf:"fixed64,15,opt,name=double,proto3" json:"double,omitempty"`
	Float    float32 `protobuf:"fixed32,16,opt,name=float,proto3" json:"float,omitempty"`
	Overflow uint64  `protobuf:"varint,17,opt,name=overflow,proto3" json:"overflow,omitempty"`
}

func (x *Dummy) Reset() {
	*x = Dummy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dummy) ProtoMessage() {}

func (x *Dummy) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dummy.ProtoReflect.Descriptor instead.
func (*Dummy) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *Dummy) GetEnum() Enum {
	if x != nil {
		return x.Enum
	}
	return Enum_V1
}

func (x *Dummy) GetRepeatedStr() []string {
	if x != nil {
		return x.RepeatedStr
	}
	return nil
}

func (x *Dummy) GetMapStrObject() map[string]*Object {
	if x != nil {
		return x.MapStrObject
	}
	return nil
}

func (x *Dummy) GetMapIntObject() map[int64]*Object {
	if x != nil {
		return x.MapIntObject
	}
	return nil
}

func (x *Dummy) GetMapBoolObject() map[bool]*Object {
	if x != nil {
		return x.MapBoolObject
	}
	return nil
}

func (x *Dummy) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Dummy) GetRepeatedObjects() []*Object {
	if x != nil {
		return x.RepeatedObjects
	}
	return nil
}

func (x *Dummy) GetSingleObject() *Object {
	if x != nil {
		return x.SingleObject
	}
	return nil
}

func (x *Dummy) GetBooleanTrue() bool {
	if x != nil {
		return x.BooleanTrue
	}
	return false
}

func (x *Dummy) GetBooleanFalse() bool {
	if x != nil {
		return x.BooleanFalse
	}
	return false
}

func (x *Dummy) GetUint32() uint32 {
	if x != nil {
		return x.Uint32
	}
	return 0
}

func (x *Dummy) GetUint64() uint64 {
	if x != nil {
		return x.Uint64
	}
	return 0
}

func (x *Dummy) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Dummy) GetSint32() int32 {
	if x != nil {
		return x.Sint32
	}
	return 0
}

func (x *Dummy) GetDouble() float64 {
	if x != nil {
		return x.Double
	}
	return 0
}

func (x *Dummy) GetFloat() float32 {
	if x != nil {
		return x.Float
	}
	return 0
}

func (x *Dummy) GetOverflow() uint64 {
	if x != nil {
		return x.Overflow
	}
	return 0
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x32, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x9e, 0x08, 0x0a, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x12,
	0x2c, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72,
	0x12, 0x51, 0x0a, 0x0e, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x5f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x75,
	0x6d, 0x6d, 0x79, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x51, 0x0a, 0x0e, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6d, 0x61, 0x70, 0x49, 0x6e, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x54, 0x0a, 0x0f, 0x6d, 0x61, 0x70, 0x5f, 0x62, 0x6f,
	0x6f, 0x6c, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x4d, 0x61, 0x70, 0x42, 0x6f,
	0x6f, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6d,
	0x61, 0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x45, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0f, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x3f, 0x0a,
	0x0d, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x74, 0x72, 0x75, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x54, 0x72, 0x75,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x66, 0x61, 0x6c,
	0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61,
	0x6e, 0x46, 0x61, 0x6c, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x5b,
	0x0a, 0x11, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5b, 0x0a, 0x11, 0x4d,
	0x61, 0x70, 0x49, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5c, 0x0a, 0x12, 0x4d, 0x61, 0x70, 0x42,
	0x6f, 0x6f, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x1e, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x06,
	0x0a, 0x02, 0x56, 0x31, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x32, 0x10, 0x01, 0x12, 0x06,
	0x0a, 0x02, 0x56, 0x33, 0x10, 0x02, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_test_proto_goTypes = []interface{}{
	(Enum)(0),                   // 0: authorization.test.Enum
	(*Object)(nil),              // 1: authorization.test.Object
	(*Dummy)(nil),               // 2: authorization.test.Dummy
	nil,                         // 3: authorization.test.Dummy.MapStrObjectEntry
	nil,                         // 4: authorization.test.Dummy.MapIntObjectEntry
	nil,                         // 5: authorization.test.Dummy.MapBoolObjectEntry
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_test_proto_depIdxs = []int32{
	0,  // 0: authorization.test.Dummy.enum:type_name -> authorization.test.Enum
	3,  // 1: authorization.test.Dummy.map_str_object:type_name -> authorization.test.Dummy.MapStrObjectEntry
	4,  // 2: authorization.test.Dummy.map_int_object:type_name -> authorization.test.Dummy.MapIntObjectEntry
	5,  // 3: authorization.test.Dummy.map_bool_object:type_name -> authorization.test.Dummy.MapBoolObjectEntry
	6,  // 4: authorization.test.Dummy.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 5: authorization.test.Dummy.repeated_objects:type_name -> authorization.test.Object
	1,  // 6: authorization.test.Dummy.single_object:type_name -> authorization.test.Object
	1,  // 7: authorization.test.Dummy.MapStrObjectEntry.value:type_name -> authorization.test.Object
	1,  // 8: authorization.test.Dummy.MapIntObjectEntry.value:type_name -> authorization.test.Object
	1,  // 9: authorization.test.Dummy.MapBoolObjectEntry.value:type_name -> authorization.test.Object
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dummy); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		EnumInfos:         file_test_proto_enumTypes,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
