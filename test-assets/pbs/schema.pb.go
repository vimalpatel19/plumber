// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

package pbs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Schema_Type int32

const (
	Schema_UNKNOWN  Schema_Type = 0
	Schema_PLAIN    Schema_Type = 1
	Schema_JSON     Schema_Type = 2
	Schema_PROTOBUF Schema_Type = 3
)

var Schema_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "PLAIN",
	2: "JSON",
	3: "PROTOBUF",
}

var Schema_Type_value = map[string]int32{
	"UNKNOWN":  0,
	"PLAIN":    1,
	"JSON":     2,
	"PROTOBUF": 3,
}

func (x Schema_Type) String() string {
	return proto.EnumName(Schema_Type_name, int32(x))
}

func (Schema_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0, 0}
}

type Schema struct {
	// The collector will ONLY fill out the 'id' for incoming messages - it is
	// the responsibility of downstream consumers to lookup the corresponding
	// schema configuration by 'id'.
	Id   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type Schema_Type       `protobuf:"varint,2,opt,name=type,proto3,enum=events.Schema_Type" json:"type,omitempty"`
	Raw  map[string][]byte `protobuf:"bytes,3,rep,name=raw,proto3" json:"raw,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Only used when Type == PROTOBUF
	ProtobufMessageName       string   `protobuf:"bytes,4,opt,name=protobuf_message_name,json=protobufMessageName,proto3" json:"protobuf_message_name,omitempty"`
	ProtobufFileDescriptorSet []byte   `protobuf:"bytes,5,opt,name=protobuf_file_descriptor_set,json=protobufFileDescriptorSet,proto3" json:"protobuf_file_descriptor_set,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *Schema) Reset()         { *m = Schema{} }
func (m *Schema) String() string { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()    {}
func (*Schema) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

func (m *Schema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schema.Unmarshal(m, b)
}
func (m *Schema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schema.Marshal(b, m, deterministic)
}
func (m *Schema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schema.Merge(m, src)
}
func (m *Schema) XXX_Size() int {
	return xxx_messageInfo_Schema.Size(m)
}
func (m *Schema) XXX_DiscardUnknown() {
	xxx_messageInfo_Schema.DiscardUnknown(m)
}

var xxx_messageInfo_Schema proto.InternalMessageInfo

func (m *Schema) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Schema) GetType() Schema_Type {
	if m != nil {
		return m.Type
	}
	return Schema_UNKNOWN
}

func (m *Schema) GetRaw() map[string][]byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Schema) GetProtobufMessageName() string {
	if m != nil {
		return m.ProtobufMessageName
	}
	return ""
}

func (m *Schema) GetProtobufFileDescriptorSet() []byte {
	if m != nil {
		return m.ProtobufFileDescriptorSet
	}
	return nil
}

func init() {
	proto.RegisterEnum("events.Schema_Type", Schema_Type_name, Schema_Type_value)
	proto.RegisterType((*Schema)(nil), "events.Schema")
	proto.RegisterMapType((map[string][]byte)(nil), "events.Schema.RawEntry")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4f, 0x6b, 0xe2, 0x50,
	0x14, 0xc5, 0x27, 0x7f, 0x74, 0xf4, 0x1a, 0x24, 0x3c, 0x67, 0x98, 0x4c, 0xe9, 0x22, 0xb8, 0x69,
	0x0a, 0x6d, 0x02, 0x16, 0xa4, 0x74, 0x53, 0x2a, 0xad, 0xd0, 0x7f, 0x89, 0x44, 0xa5, 0xd0, 0x4d,
	0x78, 0x89, 0x57, 0x0d, 0x4d, 0x4c, 0x78, 0xef, 0x45, 0xc9, 0xc7, 0xec, 0x37, 0x2a, 0x26, 0xd5,
	0x45, 0x77, 0xf7, 0x72, 0x7e, 0x87, 0x73, 0x0e, 0x68, 0x3c, 0x5a, 0x63, 0x4a, 0xed, 0x9c, 0x65,
	0x22, 0x23, 0x4d, 0xdc, 0xe2, 0x46, 0xf0, 0xfe, 0xa7, 0x0c, 0xcd, 0x69, 0x25, 0x90, 0x2e, 0xc8,
	0xf1, 0xc2, 0x90, 0x4c, 0xc9, 0x6a, 0xfb, 0x72, 0xbc, 0x20, 0x67, 0xa0, 0x8a, 0x32, 0x47, 0x43,
	0x36, 0x25, 0xab, 0x3b, 0xe8, 0xd9, 0xb5, 0xc3, 0xae, 0x69, 0x7b, 0x56, 0xe6, 0xe8, 0x57, 0x00,
	0x39, 0x07, 0x85, 0xd1, 0x9d, 0xa1, 0x98, 0x8a, 0xd5, 0x19, 0xfc, 0xfb, 0xc1, 0xf9, 0x74, 0xf7,
	0xb0, 0x11, 0xac, 0xf4, 0xf7, 0x0c, 0x19, 0xc0, 0xdf, 0x2a, 0x3f, 0x2c, 0x96, 0x41, 0x8a, 0x9c,
	0xd3, 0x15, 0x06, 0x1b, 0x9a, 0xa2, 0xa1, 0x56, 0xb1, 0xbd, 0x83, 0xf8, 0x5a, 0x6b, 0x2e, 0x4d,
	0x91, 0xdc, 0xc2, 0xe9, 0xd1, 0xb3, 0x8c, 0x13, 0x0c, 0x16, 0xc8, 0x23, 0x16, 0xe7, 0x22, 0x63,
	0x01, 0x47, 0x61, 0x34, 0x4c, 0xc9, 0xd2, 0xfc, 0xff, 0x07, 0x66, 0x1c, 0x27, 0x78, 0x7f, 0x24,
	0xa6, 0x28, 0x4e, 0x86, 0xd0, 0x3a, 0xb4, 0x20, 0x3a, 0x28, 0x1f, 0x58, 0x7e, 0xaf, 0xdc, 0x9f,
	0xe4, 0x0f, 0x34, 0xb6, 0x34, 0x29, 0xea, 0x9d, 0x9a, 0x5f, 0x3f, 0x37, 0xf2, 0xb5, 0xd4, 0x1f,
	0x82, 0xba, 0x5f, 0x49, 0x3a, 0xf0, 0x7b, 0xee, 0x3e, 0xbb, 0xde, 0x9b, 0xab, 0xff, 0x22, 0x6d,
	0x68, 0x4c, 0x5e, 0xee, 0x1e, 0x5d, 0x5d, 0x22, 0x2d, 0x50, 0x9f, 0xa6, 0x9e, 0xab, 0xcb, 0x44,
	0x83, 0xd6, 0xc4, 0xf7, 0x66, 0xde, 0x68, 0x3e, 0xd6, 0x95, 0x91, 0xfd, 0x7e, 0xb1, 0x8a, 0xc5,
	0xba, 0x08, 0xed, 0x28, 0x4b, 0x9d, 0x90, 0x8a, 0x68, 0x1d, 0x65, 0x2c, 0x77, 0xf2, 0xa4, 0x48,
	0x43, 0x64, 0x8e, 0x40, 0x2e, 0x2e, 0x29, 0xe7, 0x28, 0xb8, 0x93, 0x87, 0x3c, 0x6c, 0x56, 0xd5,
	0xaf, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xb3, 0x1f, 0xb0, 0xa2, 0x01, 0x00, 0x00,
}