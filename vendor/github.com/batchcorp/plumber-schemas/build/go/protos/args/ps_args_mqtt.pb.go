// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_args_mqtt.proto

package args

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

type MQTTQoSLevel int32

const (
	MQTTQoSLevel_MQTT_QOS_LEVEL_AT_MOST_ONCE  MQTTQoSLevel = 0
	MQTTQoSLevel_MQTT_QOS_LEVEL_AT_LEAST_ONCE MQTTQoSLevel = 1
	MQTTQoSLevel_MQTT_QOS_LEVEL_EXACTLY_ONCE  MQTTQoSLevel = 2
)

var MQTTQoSLevel_name = map[int32]string{
	0: "MQTT_QOS_LEVEL_AT_MOST_ONCE",
	1: "MQTT_QOS_LEVEL_AT_LEAST_ONCE",
	2: "MQTT_QOS_LEVEL_EXACTLY_ONCE",
}

var MQTTQoSLevel_value = map[string]int32{
	"MQTT_QOS_LEVEL_AT_MOST_ONCE":  0,
	"MQTT_QOS_LEVEL_AT_LEAST_ONCE": 1,
	"MQTT_QOS_LEVEL_EXACTLY_ONCE":  2,
}

func (x MQTTQoSLevel) String() string {
	return proto.EnumName(MQTTQoSLevel_name, int32(x))
}

func (MQTTQoSLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_11ff87d0ae52f3d2, []int{0}
}

type MQTTTLSOptions struct {
	// @gotags: kong:"help='CA file (only needed if addr is ssl://)',env=PLUMBER_RELAY_MQTT_TLS_CA_FILE"
	CaFile []byte `protobuf:"bytes,1,opt,name=ca_file,json=caFile,proto3" json:"ca_file,omitempty" kong:"help='CA file (only needed if addr is ssl://)',env=PLUMBER_RELAY_MQTT_TLS_CA_FILE"`
	// @gotags: kong:"help='Client cert file (only needed if addr is ssl://)',env=PLUMBER_RELAY_MQTT_TLS_CLIENT_CERT"
	ClientCert []byte `protobuf:"bytes,2,opt,name=client_cert,json=clientCert,proto3" json:"client_cert,omitempty" kong:"help='Client cert file (only needed if addr is ssl://)',env=PLUMBER_RELAY_MQTT_TLS_CLIENT_CERT"`
	// @gotags: kong:"help='Client key file (only needed if addr is ssl://)',env=PLUMBER_RELAY_MQTT_TLS_CLIENT_KEY"
	ClientKey []byte `protobuf:"bytes,3,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty" kong:"help='Client key file (only needed if addr is ssl://)',env=PLUMBER_RELAY_MQTT_TLS_CLIENT_KEY"`
	// @gotags: kong:"help='Whether to verify server certificate',env=PLUMBER_RELAY_MQTT_SKIP_VERIFY_TLS"
	SkipVerify           bool     `protobuf:"varint,4,opt,name=skip_verify,json=skipVerify,proto3" json:"skip_verify,omitempty" kong:"help='Whether to verify server certificate',env=PLUMBER_RELAY_MQTT_SKIP_VERIFY_TLS"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MQTTTLSOptions) Reset()         { *m = MQTTTLSOptions{} }
func (m *MQTTTLSOptions) String() string { return proto.CompactTextString(m) }
func (*MQTTTLSOptions) ProtoMessage()    {}
func (*MQTTTLSOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_11ff87d0ae52f3d2, []int{0}
}

func (m *MQTTTLSOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MQTTTLSOptions.Unmarshal(m, b)
}
func (m *MQTTTLSOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MQTTTLSOptions.Marshal(b, m, deterministic)
}
func (m *MQTTTLSOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MQTTTLSOptions.Merge(m, src)
}
func (m *MQTTTLSOptions) XXX_Size() int {
	return xxx_messageInfo_MQTTTLSOptions.Size(m)
}
func (m *MQTTTLSOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MQTTTLSOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MQTTTLSOptions proto.InternalMessageInfo

func (m *MQTTTLSOptions) GetCaFile() []byte {
	if m != nil {
		return m.CaFile
	}
	return nil
}

func (m *MQTTTLSOptions) GetClientCert() []byte {
	if m != nil {
		return m.ClientCert
	}
	return nil
}

func (m *MQTTTLSOptions) GetClientKey() []byte {
	if m != nil {
		return m.ClientKey
	}
	return nil
}

func (m *MQTTTLSOptions) GetSkipVerify() bool {
	if m != nil {
		return m.SkipVerify
	}
	return false
}

type MQTTConn struct {
	// @gotags: kong:"help='MQTT address',default='tcp://localhost:1883',env='PLUMBER_RELAY_MQTT_ADDRESS',required"
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" kong:"help='MQTT address',default='tcp://localhost:1883',env='PLUMBER_RELAY_MQTT_ADDRESS',required"`
	// @gotags: kong:"help='How long to attempt to connect for',env='PLUMBER_RELAY_MQTT_CONNECT_TIMEOUT',default=5"
	ConnTimeoutSeconds uint32 `protobuf:"varint,3,opt,name=conn_timeout_seconds,json=connTimeoutSeconds,proto3" json:"conn_timeout_seconds,omitempty" kong:"help='How long to attempt to connect for',env='PLUMBER_RELAY_MQTT_CONNECT_TIMEOUT',default=5"`
	// @gotags: kong:"help='Client id presented to MQTT broker',env='PLUMBER_RELAY_MQTT_CLIENT_ID',default=plumber"
	ClientId string `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" kong:"help='Client id presented to MQTT broker',env='PLUMBER_RELAY_MQTT_CLIENT_ID',default=plumber"`
	// @gotags: kong:"help='QoS level to use for pub/sub (options: at_most_once, at_least_once, exactly_once)',env=PLUMBER_RELAY_MQTT_QOS,type=pbenum,pbenum_strip_prefix=MQTT_QOS_LEVEL_,pbenum_lowercase,default=at_most_once"
	QosLevel MQTTQoSLevel `protobuf:"varint,5,opt,name=qos_level,json=qosLevel,proto3,enum=protos.args.MQTTQoSLevel" json:"qos_level,omitempty" kong:"help='QoS level to use for pub/sub (options: at_most_once, at_least_once, exactly_once)',env=PLUMBER_RELAY_MQTT_QOS,type=pbenum,pbenum_strip_prefix=MQTT_QOS_LEVEL_,pbenum_lowercase,default=at_most_once"`
	// @gotags: kong:"embed"
	TlsOptions           *MQTTTLSOptions `protobuf:"bytes,6,opt,name=tls_options,json=tlsOptions,proto3" json:"tls_options,omitempty" kong:"embed"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MQTTConn) Reset()         { *m = MQTTConn{} }
func (m *MQTTConn) String() string { return proto.CompactTextString(m) }
func (*MQTTConn) ProtoMessage()    {}
func (*MQTTConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_11ff87d0ae52f3d2, []int{1}
}

func (m *MQTTConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MQTTConn.Unmarshal(m, b)
}
func (m *MQTTConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MQTTConn.Marshal(b, m, deterministic)
}
func (m *MQTTConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MQTTConn.Merge(m, src)
}
func (m *MQTTConn) XXX_Size() int {
	return xxx_messageInfo_MQTTConn.Size(m)
}
func (m *MQTTConn) XXX_DiscardUnknown() {
	xxx_messageInfo_MQTTConn.DiscardUnknown(m)
}

var xxx_messageInfo_MQTTConn proto.InternalMessageInfo

func (m *MQTTConn) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MQTTConn) GetConnTimeoutSeconds() uint32 {
	if m != nil {
		return m.ConnTimeoutSeconds
	}
	return 0
}

func (m *MQTTConn) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *MQTTConn) GetQosLevel() MQTTQoSLevel {
	if m != nil {
		return m.QosLevel
	}
	return MQTTQoSLevel_MQTT_QOS_LEVEL_AT_MOST_ONCE
}

func (m *MQTTConn) GetTlsOptions() *MQTTTLSOptions {
	if m != nil {
		return m.TlsOptions
	}
	return nil
}

type MQTTReadArgs struct {
	// @gotags: kong:"help='Topic to read message(s) from',env='PLUMBER_RELAY_MQTT_TOPIC',required"
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty" kong:"help='Topic to read message(s) from',env='PLUMBER_RELAY_MQTT_TOPIC',required"`
	// @gotags: kong:"help='How long to attempt to read message(s)',default=0,env='PLUMBER_RELAY_MQTT_READ_TIMEOUT_SECONDS'"
	ReadTimeoutSeconds   uint32   `protobuf:"varint,2,opt,name=read_timeout_seconds,json=readTimeoutSeconds,proto3" json:"read_timeout_seconds,omitempty" kong:"help='How long to attempt to read message(s)',default=0,env='PLUMBER_RELAY_MQTT_READ_TIMEOUT_SECONDS'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MQTTReadArgs) Reset()         { *m = MQTTReadArgs{} }
func (m *MQTTReadArgs) String() string { return proto.CompactTextString(m) }
func (*MQTTReadArgs) ProtoMessage()    {}
func (*MQTTReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_11ff87d0ae52f3d2, []int{2}
}

func (m *MQTTReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MQTTReadArgs.Unmarshal(m, b)
}
func (m *MQTTReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MQTTReadArgs.Marshal(b, m, deterministic)
}
func (m *MQTTReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MQTTReadArgs.Merge(m, src)
}
func (m *MQTTReadArgs) XXX_Size() int {
	return xxx_messageInfo_MQTTReadArgs.Size(m)
}
func (m *MQTTReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_MQTTReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_MQTTReadArgs proto.InternalMessageInfo

func (m *MQTTReadArgs) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *MQTTReadArgs) GetReadTimeoutSeconds() uint32 {
	if m != nil {
		return m.ReadTimeoutSeconds
	}
	return 0
}

type MQTTWriteArgs struct {
	// @gotags: kong:"help='Topic to write message(s) to',required"
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty" kong:"help='Topic to write message(s) to',required"`
	// @gotags: kong:"help='How long to attempt to publish message(s)',default=5"
	WriteTimeoutSeconds  uint32   `protobuf:"varint,2,opt,name=write_timeout_seconds,json=writeTimeoutSeconds,proto3" json:"write_timeout_seconds,omitempty" kong:"help='How long to attempt to publish message(s)',default=5"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MQTTWriteArgs) Reset()         { *m = MQTTWriteArgs{} }
func (m *MQTTWriteArgs) String() string { return proto.CompactTextString(m) }
func (*MQTTWriteArgs) ProtoMessage()    {}
func (*MQTTWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_11ff87d0ae52f3d2, []int{3}
}

func (m *MQTTWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MQTTWriteArgs.Unmarshal(m, b)
}
func (m *MQTTWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MQTTWriteArgs.Marshal(b, m, deterministic)
}
func (m *MQTTWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MQTTWriteArgs.Merge(m, src)
}
func (m *MQTTWriteArgs) XXX_Size() int {
	return xxx_messageInfo_MQTTWriteArgs.Size(m)
}
func (m *MQTTWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_MQTTWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_MQTTWriteArgs proto.InternalMessageInfo

func (m *MQTTWriteArgs) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *MQTTWriteArgs) GetWriteTimeoutSeconds() uint32 {
	if m != nil {
		return m.WriteTimeoutSeconds
	}
	return 0
}

func init() {
	proto.RegisterEnum("protos.args.MQTTQoSLevel", MQTTQoSLevel_name, MQTTQoSLevel_value)
	proto.RegisterType((*MQTTTLSOptions)(nil), "protos.args.MQTTTLSOptions")
	proto.RegisterType((*MQTTConn)(nil), "protos.args.MQTTConn")
	proto.RegisterType((*MQTTReadArgs)(nil), "protos.args.MQTTReadArgs")
	proto.RegisterType((*MQTTWriteArgs)(nil), "protos.args.MQTTWriteArgs")
}

func init() { proto.RegisterFile("ps_args_mqtt.proto", fileDescriptor_11ff87d0ae52f3d2) }

var fileDescriptor_11ff87d0ae52f3d2 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x51, 0x6f, 0xd3, 0x30,
	0x18, 0x24, 0x85, 0x75, 0xed, 0xd7, 0x6d, 0x42, 0x66, 0x88, 0xa0, 0x82, 0x16, 0xf5, 0xa9, 0x42,
	0x22, 0x41, 0x45, 0x42, 0x42, 0xf0, 0x52, 0xaa, 0x20, 0x21, 0x32, 0xaa, 0x26, 0x51, 0x61, 0xbc,
	0x58, 0x89, 0xe3, 0xb5, 0xd6, 0x92, 0x38, 0xb5, 0xdd, 0xa1, 0xfe, 0x05, 0xfe, 0x29, 0xff, 0x02,
	0xd9, 0x6e, 0x05, 0x55, 0x61, 0x4f, 0xf6, 0x77, 0xf7, 0xf9, 0x7c, 0x77, 0x80, 0x1a, 0x89, 0x33,
	0xb1, 0x90, 0xb8, 0x5a, 0x29, 0xe5, 0x37, 0x82, 0x2b, 0x8e, 0x7a, 0xe6, 0x90, 0xbe, 0xc6, 0x07,
	0x3f, 0x1d, 0x38, 0xbb, 0x9c, 0xa5, 0x69, 0x1a, 0x25, 0xd3, 0x46, 0x31, 0x5e, 0x4b, 0xf4, 0x04,
	0x8e, 0x49, 0x86, 0xaf, 0x59, 0x49, 0x5d, 0xc7, 0x73, 0x86, 0x27, 0x71, 0x9b, 0x64, 0x1f, 0x59,
	0x49, 0xd1, 0x05, 0xf4, 0x48, 0xc9, 0x68, 0xad, 0x30, 0xa1, 0x42, 0xb9, 0x2d, 0x43, 0x82, 0x85,
	0x26, 0x54, 0x28, 0xf4, 0x1c, 0xb6, 0x13, 0xbe, 0xa1, 0x1b, 0xf7, 0xbe, 0xe1, 0xbb, 0x16, 0xf9,
	0x4c, 0x37, 0xfa, 0xbd, 0xbc, 0x61, 0x0d, 0xbe, 0xa5, 0x82, 0x5d, 0x6f, 0xdc, 0x07, 0x9e, 0x33,
	0xec, 0xc4, 0xa0, 0xa1, 0xb9, 0x41, 0x06, 0xbf, 0x1c, 0xe8, 0x68, 0x33, 0x13, 0x5e, 0xd7, 0xc8,
	0x85, 0xe3, 0xac, 0x28, 0x04, 0x95, 0xd2, 0xd8, 0xe8, 0xc6, 0xbb, 0x11, 0xbd, 0x82, 0x73, 0xc2,
	0xeb, 0x1a, 0x2b, 0x56, 0x51, 0xbe, 0x56, 0x58, 0x52, 0xc2, 0xeb, 0x42, 0x9a, 0x0f, 0x4f, 0x63,
	0xa4, 0xb9, 0xd4, 0x52, 0x89, 0x65, 0x50, 0x1f, 0xb6, 0x36, 0x30, 0x2b, 0xcc, 0xbf, 0xdd, 0xb8,
	0x63, 0x81, 0x4f, 0x05, 0x7a, 0x03, 0xdd, 0x15, 0x97, 0xb8, 0xa4, 0xb7, 0xb4, 0x74, 0x8f, 0x3c,
	0x67, 0x78, 0x36, 0x7a, 0xea, 0xff, 0xd5, 0x91, 0xaf, 0x2d, 0xcd, 0x78, 0x12, 0xe9, 0x85, 0xb8,
	0xb3, 0xe2, 0xd2, 0xdc, 0xd0, 0x7b, 0xe8, 0xa9, 0x52, 0x62, 0x6e, 0x6b, 0x73, 0xdb, 0x9e, 0x33,
	0xec, 0x8d, 0xfa, 0x07, 0x2f, 0xff, 0x34, 0x1b, 0x83, 0x2a, 0xe5, 0xf6, 0x3e, 0x98, 0xc3, 0x89,
	0x66, 0x63, 0x9a, 0x15, 0x63, 0xb1, 0x90, 0xe8, 0x1c, 0x8e, 0x14, 0x6f, 0x18, 0xd9, 0x86, 0xb5,
	0x83, 0x8e, 0x2a, 0x68, 0x56, 0x1c, 0x44, 0x6d, 0xd9, 0xa8, 0x9a, 0xdb, 0x8f, 0x3a, 0xb8, 0x82,
	0x53, 0xad, 0xfb, 0x55, 0x30, 0x45, 0xef, 0x10, 0x1e, 0xc1, 0xe3, 0x1f, 0x7a, 0xe5, 0x3f, 0xca,
	0x8f, 0x0c, 0xb9, 0x2f, 0xfd, 0x42, 0x58, 0xcb, 0xbb, 0x2a, 0xd0, 0x05, 0xf4, 0xf5, 0x8c, 0x67,
	0xd3, 0x04, 0x47, 0xe1, 0x3c, 0x8c, 0xf0, 0x38, 0xc5, 0x97, 0xd3, 0x24, 0xc5, 0xd3, 0x2f, 0x93,
	0xf0, 0xe1, 0x3d, 0xe4, 0xc1, 0xb3, 0xc3, 0x85, 0x28, 0x1c, 0xef, 0x36, 0x9c, 0x7f, 0x48, 0x84,
	0xdf, 0xc6, 0x93, 0x34, 0xba, 0xb2, 0x0b, 0xad, 0x0f, 0xef, 0xbe, 0xbf, 0x5d, 0x30, 0xb5, 0x5c,
	0xe7, 0x3e, 0xe1, 0x55, 0x90, 0x67, 0x8a, 0x2c, 0x09, 0x17, 0x4d, 0xd0, 0x94, 0xeb, 0x2a, 0xa7,
	0xe2, 0xa5, 0x24, 0x4b, 0x5a, 0x65, 0x32, 0xc8, 0xd7, 0xac, 0x2c, 0x82, 0x05, 0x0f, 0x6c, 0xfd,
	0x81, 0xae, 0x3f, 0x6f, 0x9b, 0xe1, 0xf5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x21, 0x8b,
	0x30, 0x06, 0x03, 0x00, 0x00,
}
