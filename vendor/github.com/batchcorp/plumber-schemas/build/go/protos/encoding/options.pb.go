// Code generated by protoc-gen-go. DO NOT EDIT.
// source: options.proto

package encoding

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

type Type int32

const (
	Type_NONE        Type = 0
	Type_JSON        Type = 1
	Type_JSON_SCHEMA Type = 2
	Type_PROTOBUF    Type = 3
	Type_AVRO        Type = 4
)

var Type_name = map[int32]string{
	0: "NONE",
	1: "JSON",
	2: "JSON_SCHEMA",
	3: "PROTOBUF",
	4: "AVRO",
}

var Type_value = map[string]int32{
	"NONE":        0,
	"JSON":        1,
	"JSON_SCHEMA": 2,
	"PROTOBUF":    3,
	"AVRO":        4,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{0}
}

type Protobuf struct {
	RootType             string   `protobuf:"bytes,1,opt,name=root_type,json=rootType,proto3" json:"root_type,omitempty"`
	ZipArchive           []byte   `protobuf:"bytes,2,opt,name=zip_archive,json=zipArchive,proto3" json:"zip_archive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Protobuf) Reset()         { *m = Protobuf{} }
func (m *Protobuf) String() string { return proto.CompactTextString(m) }
func (*Protobuf) ProtoMessage()    {}
func (*Protobuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{0}
}

func (m *Protobuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Protobuf.Unmarshal(m, b)
}
func (m *Protobuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Protobuf.Marshal(b, m, deterministic)
}
func (m *Protobuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Protobuf.Merge(m, src)
}
func (m *Protobuf) XXX_Size() int {
	return xxx_messageInfo_Protobuf.Size(m)
}
func (m *Protobuf) XXX_DiscardUnknown() {
	xxx_messageInfo_Protobuf.DiscardUnknown(m)
}

var xxx_messageInfo_Protobuf proto.InternalMessageInfo

func (m *Protobuf) GetRootType() string {
	if m != nil {
		return m.RootType
	}
	return ""
}

func (m *Protobuf) GetZipArchive() []byte {
	if m != nil {
		return m.ZipArchive
	}
	return nil
}

type JSONSchema struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JSONSchema) Reset()         { *m = JSONSchema{} }
func (m *JSONSchema) String() string { return proto.CompactTextString(m) }
func (*JSONSchema) ProtoMessage()    {}
func (*JSONSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{1}
}

func (m *JSONSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JSONSchema.Unmarshal(m, b)
}
func (m *JSONSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JSONSchema.Marshal(b, m, deterministic)
}
func (m *JSONSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JSONSchema.Merge(m, src)
}
func (m *JSONSchema) XXX_Size() int {
	return xxx_messageInfo_JSONSchema.Size(m)
}
func (m *JSONSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_JSONSchema.DiscardUnknown(m)
}

var xxx_messageInfo_JSONSchema proto.InternalMessageInfo

type Avro struct {
	Schema               []byte   `protobuf:"bytes,1,opt,name=Schema,proto3" json:"Schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Avro) Reset()         { *m = Avro{} }
func (m *Avro) String() string { return proto.CompactTextString(m) }
func (*Avro) ProtoMessage()    {}
func (*Avro) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{2}
}

func (m *Avro) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Avro.Unmarshal(m, b)
}
func (m *Avro) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Avro.Marshal(b, m, deterministic)
}
func (m *Avro) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Avro.Merge(m, src)
}
func (m *Avro) XXX_Size() int {
	return xxx_messageInfo_Avro.Size(m)
}
func (m *Avro) XXX_DiscardUnknown() {
	xxx_messageInfo_Avro.DiscardUnknown(m)
}

var xxx_messageInfo_Avro proto.InternalMessageInfo

func (m *Avro) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

type Options struct {
	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=protos.encoding.Type" json:"type,omitempty"`
	// Specify an existing stored schema to use instead of specifying a Encoding payload
	SchemaId string `protobuf:"bytes,2,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	// Only filled out if "type" is not NONE or JSON
	//
	// Types that are valid to be assigned to Encoding:
	//	*Options_Protobuf
	//	*Options_Avro
	//	*Options_JsonSchema
	Encoding             isOptions_Encoding `protobuf_oneof:"Encoding"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{3}
}

func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_NONE
}

func (m *Options) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

type isOptions_Encoding interface {
	isOptions_Encoding()
}

type Options_Protobuf struct {
	Protobuf *Protobuf `protobuf:"bytes,100,opt,name=protobuf,proto3,oneof"`
}

type Options_Avro struct {
	Avro *Avro `protobuf:"bytes,102,opt,name=avro,proto3,oneof"`
}

type Options_JsonSchema struct {
	JsonSchema *JSONSchema `protobuf:"bytes,101,opt,name=json_schema,json=jsonSchema,proto3,oneof"`
}

func (*Options_Protobuf) isOptions_Encoding() {}

func (*Options_Avro) isOptions_Encoding() {}

func (*Options_JsonSchema) isOptions_Encoding() {}

func (m *Options) GetEncoding() isOptions_Encoding {
	if m != nil {
		return m.Encoding
	}
	return nil
}

func (m *Options) GetProtobuf() *Protobuf {
	if x, ok := m.GetEncoding().(*Options_Protobuf); ok {
		return x.Protobuf
	}
	return nil
}

func (m *Options) GetAvro() *Avro {
	if x, ok := m.GetEncoding().(*Options_Avro); ok {
		return x.Avro
	}
	return nil
}

func (m *Options) GetJsonSchema() *JSONSchema {
	if x, ok := m.GetEncoding().(*Options_JsonSchema); ok {
		return x.JsonSchema
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Options) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Options_Protobuf)(nil),
		(*Options_Avro)(nil),
		(*Options_JsonSchema)(nil),
	}
}

func init() {
	proto.RegisterEnum("protos.encoding.Type", Type_name, Type_value)
	proto.RegisterType((*Protobuf)(nil), "protos.encoding.Protobuf")
	proto.RegisterType((*JSONSchema)(nil), "protos.encoding.JSONSchema")
	proto.RegisterType((*Avro)(nil), "protos.encoding.Avro")
	proto.RegisterType((*Options)(nil), "protos.encoding.Options")
}

func init() { proto.RegisterFile("options.proto", fileDescriptor_110d40819f1994f9) }

var fileDescriptor_110d40819f1994f9 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xdd, 0x8b, 0xda, 0x40,
	0x14, 0xc5, 0x13, 0x1b, 0x6c, 0xbc, 0x49, 0x6b, 0x18, 0x68, 0x49, 0x11, 0x5a, 0xc9, 0x93, 0x6d,
	0x69, 0x02, 0xf6, 0xa1, 0x4f, 0x6d, 0x89, 0x62, 0x49, 0x17, 0xd6, 0xc8, 0xe8, 0xee, 0xc3, 0xbe,
	0x84, 0x7c, 0x69, 0x66, 0xd1, 0xcc, 0x90, 0x0f, 0x41, 0x5f, 0xf7, 0x1f, 0x5f, 0x66, 0x46, 0x5d,
	0x58, 0x9f, 0x66, 0xee, 0xe1, 0x9c, 0x7b, 0x7f, 0xf7, 0xc2, 0x3b, 0xca, 0x1a, 0x42, 0xcb, 0xda,
	0x65, 0x15, 0x6d, 0x28, 0xea, 0x8b, 0xa7, 0x76, 0xf3, 0x32, 0xa5, 0x19, 0x29, 0x37, 0x4e, 0x00,
	0xfa, 0x82, 0x4b, 0x49, 0xbb, 0x46, 0x03, 0xe8, 0x55, 0x94, 0x36, 0x51, 0x73, 0x60, 0xb9, 0xad,
	0x0e, 0xd5, 0x51, 0x0f, 0xeb, 0x5c, 0x58, 0x1d, 0x58, 0x8e, 0xbe, 0x80, 0x71, 0x24, 0x2c, 0x8a,
	0xab, 0xb4, 0x20, 0xfb, 0xdc, 0xee, 0x0c, 0xd5, 0x91, 0x89, 0xe1, 0x48, 0x98, 0x2f, 0x15, 0xc7,
	0x04, 0xb8, 0x59, 0x86, 0xf3, 0x65, 0x5a, 0xe4, 0xbb, 0xd8, 0xf9, 0x0c, 0x9a, 0xbf, 0xaf, 0x28,
	0xfa, 0x08, 0x5d, 0xa9, 0x88, 0x86, 0x26, 0x3e, 0x55, 0xce, 0x53, 0x07, 0xde, 0x86, 0x12, 0x0d,
	0x7d, 0x05, 0xed, 0x32, 0xf2, 0xfd, 0xf8, 0x83, 0xfb, 0x8a, 0xd1, 0xe5, 0xf3, 0xb1, 0xb0, 0x70,
	0xc4, 0x5a, 0x34, 0x88, 0x48, 0x26, 0x18, 0x7a, 0x58, 0x97, 0xc2, 0xff, 0x0c, 0xfd, 0x02, 0x9d,
	0x9d, 0x76, 0xb1, 0xb3, 0xa1, 0x3a, 0x32, 0xc6, 0x9f, 0xae, 0x7a, 0x9d, 0x97, 0x0d, 0x14, 0x7c,
	0x31, 0xa3, 0xef, 0xa0, 0xc5, 0xfb, 0x8a, 0xda, 0x6b, 0x11, 0xba, 0x06, 0xe0, 0x9b, 0x04, 0x0a,
	0x16, 0x26, 0xf4, 0x07, 0x8c, 0xc7, 0x9a, 0x96, 0x91, 0x1c, 0x6b, 0xe7, 0x22, 0x33, 0xb8, 0xca,
	0xbc, 0xdc, 0x22, 0x50, 0x30, 0xf0, 0x84, 0xac, 0x26, 0x00, 0xfa, 0xec, 0x64, 0xfa, 0x36, 0x05,
	0x4d, 0x1c, 0x57, 0x07, 0x6d, 0x1e, 0xce, 0x67, 0x96, 0xc2, 0x7f, 0x3c, 0x69, 0xa9, 0xa8, 0x0f,
	0x06, 0xff, 0x45, 0xcb, 0x69, 0x30, 0xbb, 0xf5, 0xad, 0x0e, 0x32, 0x41, 0x5f, 0xe0, 0x70, 0x15,
	0x4e, 0xee, 0xfe, 0x59, 0x6f, 0xb8, 0xd1, 0xbf, 0xc7, 0xa1, 0xa5, 0x4d, 0xfe, 0x3e, 0xfc, 0xde,
	0x90, 0xa6, 0x68, 0x13, 0x37, 0xa5, 0x3b, 0x2f, 0x89, 0x9b, 0xb4, 0x48, 0x69, 0xc5, 0x3c, 0xb6,
	0x6d, 0x77, 0x49, 0x5e, 0xfd, 0x90, 0xa0, 0xb5, 0x97, 0xb4, 0x64, 0x9b, 0x79, 0x1b, 0xea, 0x49,
	0x54, 0xef, 0x8c, 0x9a, 0x74, 0x85, 0xf0, 0xf3, 0x39, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xbf, 0x3f,
	0xda, 0x2c, 0x02, 0x00, 0x00,
}