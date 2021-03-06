// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//agent --> server
type RawData struct {
	IntranetIPV4         []string  `protobuf:"bytes,1,rep,name=IntranetIPV4,proto3" json:"IntranetIPV4,omitempty"`
	ExtranetIPV4         []string  `protobuf:"bytes,2,rep,name=ExtranetIPV4,proto3" json:"ExtranetIPV4,omitempty"`
	IntranetIPV6         []string  `protobuf:"bytes,3,rep,name=IntranetIPV6,proto3" json:"IntranetIPV6,omitempty"`
	ExtranetIPV6         []string  `protobuf:"bytes,4,rep,name=ExtranetIPV6,proto3" json:"ExtranetIPV6,omitempty"`
	HostName             string    `protobuf:"bytes,5,opt,name=HostName,proto3" json:"HostName,omitempty"`
	AgentID              string    `protobuf:"bytes,6,opt,name=AgentID,proto3" json:"AgentID,omitempty"`
	Timestamp            int64     `protobuf:"varint,7,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Version              string    `protobuf:"bytes,8,opt,name=Version,proto3" json:"Version,omitempty"`
	Pkg                  []*Record `protobuf:"bytes,9,rep,name=Pkg,proto3" json:"Pkg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RawData) Reset()         { *m = RawData{} }
func (m *RawData) String() string { return proto.CompactTextString(m) }
func (*RawData) ProtoMessage()    {}
func (*RawData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *RawData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawData.Unmarshal(m, b)
}
func (m *RawData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawData.Marshal(b, m, deterministic)
}
func (m *RawData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawData.Merge(m, src)
}
func (m *RawData) XXX_Size() int {
	return xxx_messageInfo_RawData.Size(m)
}
func (m *RawData) XXX_DiscardUnknown() {
	xxx_messageInfo_RawData.DiscardUnknown(m)
}

var xxx_messageInfo_RawData proto.InternalMessageInfo

func (m *RawData) GetIntranetIPV4() []string {
	if m != nil {
		return m.IntranetIPV4
	}
	return nil
}

func (m *RawData) GetExtranetIPV4() []string {
	if m != nil {
		return m.ExtranetIPV4
	}
	return nil
}

func (m *RawData) GetIntranetIPV6() []string {
	if m != nil {
		return m.IntranetIPV6
	}
	return nil
}

func (m *RawData) GetExtranetIPV6() []string {
	if m != nil {
		return m.ExtranetIPV6
	}
	return nil
}

func (m *RawData) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *RawData) GetAgentID() string {
	if m != nil {
		return m.AgentID
	}
	return ""
}

func (m *RawData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *RawData) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RawData) GetPkg() []*Record {
	if m != nil {
		return m.Pkg
	}
	return nil
}

type Record struct {
	Message              map[string]string `protobuf:"bytes,1,rep,name=message,proto3" json:"message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetMessage() map[string]string {
	if m != nil {
		return m.Message
	}
	return nil
}

//server --> agent
type Command struct {
	AgentCtrl            int32         `protobuf:"varint,1,opt,name=AgentCtrl,proto3" json:"AgentCtrl,omitempty"`
	Task                 *PluginTask   `protobuf:"bytes,2,opt,name=Task,proto3" json:"Task,omitempty"`
	Config               []*ConfigItem `protobuf:"bytes,3,rep,name=Config,proto3" json:"Config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{2}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetAgentCtrl() int32 {
	if m != nil {
		return m.AgentCtrl
	}
	return 0
}

func (m *Command) GetTask() *PluginTask {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *Command) GetConfig() []*ConfigItem {
	if m != nil {
		return m.Config
	}
	return nil
}

type PluginTask struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginTask) Reset()         { *m = PluginTask{} }
func (m *PluginTask) String() string { return proto.CompactTextString(m) }
func (*PluginTask) ProtoMessage()    {}
func (*PluginTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{3}
}

func (m *PluginTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginTask.Unmarshal(m, b)
}
func (m *PluginTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginTask.Marshal(b, m, deterministic)
}
func (m *PluginTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginTask.Merge(m, src)
}
func (m *PluginTask) XXX_Size() int {
	return xxx_messageInfo_PluginTask.Size(m)
}
func (m *PluginTask) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginTask.DiscardUnknown(m)
}

var xxx_messageInfo_PluginTask proto.InternalMessageInfo

func (m *PluginTask) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginTask) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *PluginTask) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ConfigItem struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`
	SHA256               string   `protobuf:"bytes,3,opt,name=SHA256,proto3" json:"SHA256,omitempty"`
	DownloadURL          []string `protobuf:"bytes,4,rep,name=DownloadURL,proto3" json:"DownloadURL,omitempty"`
	Detail               string   `protobuf:"bytes,5,opt,name=Detail,proto3" json:"Detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigItem) Reset()         { *m = ConfigItem{} }
func (m *ConfigItem) String() string { return proto.CompactTextString(m) }
func (*ConfigItem) ProtoMessage()    {}
func (*ConfigItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{4}
}

func (m *ConfigItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigItem.Unmarshal(m, b)
}
func (m *ConfigItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigItem.Marshal(b, m, deterministic)
}
func (m *ConfigItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigItem.Merge(m, src)
}
func (m *ConfigItem) XXX_Size() int {
	return xxx_messageInfo_ConfigItem.Size(m)
}
func (m *ConfigItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigItem.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigItem proto.InternalMessageInfo

func (m *ConfigItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConfigItem) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ConfigItem) GetSHA256() string {
	if m != nil {
		return m.SHA256
	}
	return ""
}

func (m *ConfigItem) GetDownloadURL() []string {
	if m != nil {
		return m.DownloadURL
	}
	return nil
}

func (m *ConfigItem) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func init() {
	proto.RegisterType((*RawData)(nil), "proto.RawData")
	proto.RegisterType((*Record)(nil), "proto.Record")
	proto.RegisterMapType((map[string]string)(nil), "proto.Record.MessageEntry")
	proto.RegisterType((*Command)(nil), "proto.Command")
	proto.RegisterType((*PluginTask)(nil), "proto.PluginTask")
	proto.RegisterType((*ConfigItem)(nil), "proto.ConfigItem")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x26, 0x49, 0x9b, 0xb4, 0xa7, 0x63, 0x02, 0x0b, 0x21, 0xab, 0x42, 0x22, 0x8a, 0x84, 0x14,
	0x6e, 0xaa, 0x29, 0x8c, 0x08, 0x4d, 0xdc, 0x8c, 0x76, 0xd2, 0x8a, 0x00, 0x55, 0xa6, 0xec, 0x82,
	0x3b, 0x6f, 0xf5, 0xa2, 0xa8, 0x89, 0x5d, 0x39, 0x1e, 0xa3, 0x3c, 0x03, 0x0f, 0xc2, 0x63, 0x22,
	0xff, 0x74, 0x49, 0xb7, 0x5d, 0xd5, 0xdf, 0xcf, 0xf9, 0x7a, 0x8e, 0x8f, 0x03, 0x50, 0xc8, 0xcd,
	0xd5, 0x64, 0x23, 0x85, 0x12, 0xa8, 0x6f, 0x7e, 0x92, 0x7f, 0x3e, 0x44, 0x84, 0xde, 0xce, 0xa8,
	0xa2, 0x28, 0x81, 0x83, 0x39, 0x57, 0x92, 0x72, 0xa6, 0xe6, 0x8b, 0x8b, 0x63, 0xec, 0xc5, 0x41,
	0x3a, 0x24, 0x7b, 0x9c, 0xf6, 0x9c, 0xfd, 0xee, 0x78, 0x7c, 0xeb, 0xe9, 0x72, 0xf7, 0x72, 0x72,
	0x1c, 0x3c, 0xc8, 0xc9, 0xef, 0xe5, 0xe4, 0xb8, 0xf7, 0x20, 0x27, 0x47, 0x63, 0x18, 0x9c, 0x8b,
	0x46, 0x7d, 0xa3, 0x35, 0xc3, 0xfd, 0xd8, 0x4b, 0x87, 0xe4, 0x0e, 0x23, 0x0c, 0xd1, 0x69, 0xc1,
	0xb8, 0x9a, 0xcf, 0x70, 0x68, 0xa4, 0x1d, 0x44, 0xaf, 0x60, 0xb8, 0x2c, 0x6b, 0xd6, 0x28, 0x5a,
	0x6f, 0x70, 0x14, 0x7b, 0x69, 0x40, 0x5a, 0x42, 0xd7, 0x5d, 0x30, 0xd9, 0x94, 0x82, 0xe3, 0x81,
	0xad, 0x73, 0x10, 0xbd, 0x86, 0x60, 0xb1, 0x2e, 0xf0, 0x30, 0x0e, 0xd2, 0x51, 0xf6, 0xd4, 0xde,
	0xd2, 0x84, 0xb0, 0x2b, 0x21, 0x57, 0x44, 0x2b, 0xc9, 0x1f, 0x08, 0x2d, 0x44, 0xc7, 0x10, 0xd5,
	0xac, 0x69, 0x68, 0xc1, 0xcc, 0x1d, 0x8d, 0xb2, 0xf1, 0x9e, 0x7d, 0xf2, 0xd5, 0x8a, 0x67, 0x5c,
	0xc9, 0x2d, 0xd9, 0x59, 0xc7, 0x27, 0x70, 0xd0, 0x15, 0xd0, 0x33, 0x08, 0xd6, 0x6c, 0x8b, 0x3d,
	0xd3, 0x86, 0x3e, 0xa2, 0x17, 0xd0, 0xff, 0x45, 0xab, 0x1b, 0x86, 0x7d, 0xc3, 0x59, 0x70, 0xe2,
	0x7f, 0xf0, 0x92, 0x2d, 0x44, 0x53, 0x51, 0xd7, 0x94, 0xaf, 0xf4, 0x7c, 0x66, 0xd4, 0xa9, 0x92,
	0x95, 0x29, 0xee, 0x93, 0x96, 0x40, 0x6f, 0xa0, 0xb7, 0xa4, 0xcd, 0xda, 0x24, 0x8c, 0xb2, 0xe7,
	0xae, 0xaf, 0x45, 0x75, 0x53, 0x94, 0x5c, 0x0b, 0xc4, 0xc8, 0xe8, 0x2d, 0x84, 0x53, 0xc1, 0xaf,
	0xcb, 0xc2, 0x2c, 0xa7, 0x35, 0x5a, 0x72, 0xae, 0x58, 0x4d, 0x9c, 0x21, 0xf9, 0x0c, 0xd0, 0x96,
	0x23, 0x04, 0x3d, 0xb3, 0x0f, 0xdb, 0xb5, 0x39, 0x6b, 0x4e, 0xbf, 0x1f, 0xd7, 0xb5, 0x39, 0xeb,
	0x51, 0x96, 0x62, 0xcd, 0x38, 0x0e, 0xec, 0x28, 0x06, 0x24, 0x7f, 0x3d, 0x80, 0xf6, 0x2f, 0x1e,
	0x0d, 0xeb, 0x2c, 0xc8, 0xdf, 0x5f, 0xd0, 0x4b, 0x08, 0xbf, 0x9f, 0x9f, 0x66, 0xef, 0x73, 0x97,
	0xe9, 0x10, 0x8a, 0x61, 0x34, 0x13, 0xb7, 0xbc, 0x12, 0x74, 0xf5, 0x83, 0x7c, 0x71, 0x2f, 0xa9,
	0x4b, 0xe9, 0xca, 0x19, 0x53, 0xb4, 0xac, 0xdc, 0x33, 0x72, 0x28, 0xfb, 0x08, 0x83, 0xa5, 0xa4,
	0xbc, 0xb9, 0x66, 0x12, 0x1d, 0x75, 0xce, 0x87, 0xbb, 0x75, 0xda, 0x0f, 0x63, 0x7c, 0x78, 0x77,
	0x3b, 0x66, 0x05, 0xc9, 0x93, 0xd4, 0x3b, 0xf2, 0x3e, 0xf5, 0x7e, 0xfa, 0x9b, 0xcb, 0xcb, 0xd0,
	0x48, 0xef, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x32, 0xa3, 0x4e, 0x5c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransferClient is the client API for Transfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransferClient interface {
	Transfer(ctx context.Context, opts ...grpc.CallOption) (Transfer_TransferClient, error)
}

type transferClient struct {
	cc *grpc.ClientConn
}

func NewTransferClient(cc *grpc.ClientConn) TransferClient {
	return &transferClient{cc}
}

func (c *transferClient) Transfer(ctx context.Context, opts ...grpc.CallOption) (Transfer_TransferClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Transfer_serviceDesc.Streams[0], "/proto.Transfer/Transfer", opts...)
	if err != nil {
		return nil, err
	}
	x := &transferTransferClient{stream}
	return x, nil
}

type Transfer_TransferClient interface {
	Send(*RawData) error
	Recv() (*Command, error)
	grpc.ClientStream
}

type transferTransferClient struct {
	grpc.ClientStream
}

func (x *transferTransferClient) Send(m *RawData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transferTransferClient) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransferServer is the server API for Transfer service.
type TransferServer interface {
	Transfer(Transfer_TransferServer) error
}

// UnimplementedTransferServer can be embedded to have forward compatible implementations.
type UnimplementedTransferServer struct {
}

func (*UnimplementedTransferServer) Transfer(srv Transfer_TransferServer) error {
	return status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}

func RegisterTransferServer(s *grpc.Server, srv TransferServer) {
	s.RegisterService(&_Transfer_serviceDesc, srv)
}

func _Transfer_Transfer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransferServer).Transfer(&transferTransferServer{stream})
}

type Transfer_TransferServer interface {
	Send(*Command) error
	Recv() (*RawData, error)
	grpc.ServerStream
}

type transferTransferServer struct {
	grpc.ServerStream
}

func (x *transferTransferServer) Send(m *Command) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transferTransferServer) Recv() (*RawData, error) {
	m := new(RawData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Transfer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Transfer",
	HandlerType: (*TransferServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Transfer",
			Handler:       _Transfer_Transfer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc.proto",
}
