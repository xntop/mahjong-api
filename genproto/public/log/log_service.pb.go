// Code generated by protoc-gen-go. DO NOT EDIT.
// source: public/log/log_service.proto

package log

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type GetLogRequest struct {
	LogId                string   `protobuf:"bytes,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLogRequest) Reset()         { *m = GetLogRequest{} }
func (m *GetLogRequest) String() string { return proto.CompactTextString(m) }
func (*GetLogRequest) ProtoMessage()    {}
func (*GetLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2007b277fc4d0ad8, []int{0}
}

func (m *GetLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLogRequest.Unmarshal(m, b)
}
func (m *GetLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLogRequest.Marshal(b, m, deterministic)
}
func (m *GetLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLogRequest.Merge(m, src)
}
func (m *GetLogRequest) XXX_Size() int {
	return xxx_messageInfo_GetLogRequest.Size(m)
}
func (m *GetLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLogRequest proto.InternalMessageInfo

func (m *GetLogRequest) GetLogId() string {
	if m != nil {
		return m.LogId
	}
	return ""
}

type GetLogResponse struct {
	Log                  *Log     `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLogResponse) Reset()         { *m = GetLogResponse{} }
func (m *GetLogResponse) String() string { return proto.CompactTextString(m) }
func (*GetLogResponse) ProtoMessage()    {}
func (*GetLogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2007b277fc4d0ad8, []int{1}
}

func (m *GetLogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLogResponse.Unmarshal(m, b)
}
func (m *GetLogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLogResponse.Marshal(b, m, deterministic)
}
func (m *GetLogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLogResponse.Merge(m, src)
}
func (m *GetLogResponse) XXX_Size() int {
	return xxx_messageInfo_GetLogResponse.Size(m)
}
func (m *GetLogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLogResponse proto.InternalMessageInfo

func (m *GetLogResponse) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func init() {
	proto.RegisterType((*GetLogRequest)(nil), "mahjong.log.GetLogRequest")
	proto.RegisterType((*GetLogResponse)(nil), "mahjong.log.GetLogResponse")
}

func init() { proto.RegisterFile("public/log/log_service.proto", fileDescriptor_2007b277fc4d0ad8) }

var fileDescriptor_2007b277fc4d0ad8 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x28, 0x4d, 0xca,
	0xc9, 0x4c, 0xd6, 0xcf, 0xc9, 0x4f, 0x07, 0xe1, 0xf8, 0xe2, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54,
	0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xee, 0xdc, 0xc4, 0x8c, 0xac, 0xfc, 0xbc, 0x74, 0xbd,
	0x9c, 0xfc, 0x74, 0x29, 0x11, 0x54, 0xa5, 0x10, 0x25, 0x4a, 0x6a, 0x5c, 0xbc, 0xee, 0xa9, 0x25,
	0x3e, 0xf9, 0xe9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xa2, 0x5c, 0x6c, 0x20, 0x83,
	0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x58, 0x73, 0xf2, 0xd3, 0x3d, 0x53, 0x94,
	0x4c, 0xb8, 0xf8, 0x60, 0xea, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x94, 0xb8, 0x98, 0x73,
	0xf2, 0xd3, 0xc1, 0xaa, 0xb8, 0x8d, 0x04, 0xf4, 0x90, 0xac, 0xd2, 0x03, 0x29, 0x03, 0x49, 0x1a,
	0xf9, 0x73, 0x71, 0xf9, 0xe4, 0xa7, 0x07, 0x43, 0x1c, 0x25, 0xe4, 0xc8, 0xc5, 0x06, 0x31, 0x43,
	0x48, 0x0a, 0x45, 0x39, 0x8a, 0x03, 0xa4, 0xa4, 0xb1, 0xca, 0x41, 0x2c, 0x75, 0x32, 0x8b, 0x32,
	0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xc9, 0xcb, 0x2f, 0xcb,
	0xcc, 0xce, 0x4f, 0x4b, 0xd3, 0x87, 0x6a, 0xd1, 0x4d, 0x2c, 0xc8, 0xd4, 0x4f, 0x4f, 0xcd, 0x03,
	0xfb, 0x4d, 0x1f, 0xe1, 0xe1, 0x24, 0x36, 0xb0, 0x88, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x39,
	0x27, 0xae, 0x97, 0x30, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogServiceClient interface {
	GetLog(ctx context.Context, in *GetLogRequest, opts ...grpc.CallOption) (*GetLogResponse, error)
}

type logServiceClient struct {
	cc *grpc.ClientConn
}

func NewLogServiceClient(cc *grpc.ClientConn) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) GetLog(ctx context.Context, in *GetLogRequest, opts ...grpc.CallOption) (*GetLogResponse, error) {
	out := new(GetLogResponse)
	err := c.cc.Invoke(ctx, "/mahjong.log.LogService/GetLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServiceServer is the server API for LogService service.
type LogServiceServer interface {
	GetLog(context.Context, *GetLogRequest) (*GetLogResponse, error)
}

func RegisterLogServiceServer(s *grpc.Server, srv LogServiceServer) {
	s.RegisterService(&_LogService_serviceDesc, srv)
}

func _LogService_GetLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).GetLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mahjong.log.LogService/GetLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).GetLog(ctx, req.(*GetLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mahjong.log.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLog",
			Handler:    _LogService_GetLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public/log/log_service.proto",
}
