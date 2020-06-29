// Code generated by protoc-gen-go. DO NOT EDIT.
// source: conn.proto

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

type DeliverMessageReq struct {
	Item                 *MessageItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	SeqId                uint64       `protobuf:"varint,2,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeliverMessageReq) Reset()         { *m = DeliverMessageReq{} }
func (m *DeliverMessageReq) String() string { return proto.CompactTextString(m) }
func (*DeliverMessageReq) ProtoMessage()    {}
func (*DeliverMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{0}
}

func (m *DeliverMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliverMessageReq.Unmarshal(m, b)
}
func (m *DeliverMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliverMessageReq.Marshal(b, m, deterministic)
}
func (m *DeliverMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliverMessageReq.Merge(m, src)
}
func (m *DeliverMessageReq) XXX_Size() int {
	return xxx_messageInfo_DeliverMessageReq.Size(m)
}
func (m *DeliverMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliverMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeliverMessageReq proto.InternalMessageInfo

func (m *DeliverMessageReq) GetItem() *MessageItem {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *DeliverMessageReq) GetSeqId() uint64 {
	if m != nil {
		return m.SeqId
	}
	return 0
}

type DeliverMessageResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliverMessageResp) Reset()         { *m = DeliverMessageResp{} }
func (m *DeliverMessageResp) String() string { return proto.CompactTextString(m) }
func (*DeliverMessageResp) ProtoMessage()    {}
func (*DeliverMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{1}
}

func (m *DeliverMessageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliverMessageResp.Unmarshal(m, b)
}
func (m *DeliverMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliverMessageResp.Marshal(b, m, deterministic)
}
func (m *DeliverMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliverMessageResp.Merge(m, src)
}
func (m *DeliverMessageResp) XXX_Size() int {
	return xxx_messageInfo_DeliverMessageResp.Size(m)
}
func (m *DeliverMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliverMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeliverMessageResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DeliverMessageReq)(nil), "pb.DeliverMessageReq")
	proto.RegisterType((*DeliverMessageResp)(nil), "pb.DeliverMessageResp")
}

func init() {
	proto.RegisterFile("conn.proto", fileDescriptor_f401a58c1fc7ceef)
}

var fileDescriptor_f401a58c1fc7ceef = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0xcf, 0xcb,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x2b, 0x49, 0x2e, 0xd0,
	0x45, 0x88, 0x29, 0xf9, 0x73, 0x09, 0xba, 0xa4, 0xe6, 0x64, 0x96, 0xa5, 0x16, 0xf9, 0xa6, 0x16,
	0x17, 0x27, 0xa6, 0xa7, 0x06, 0xa5, 0x16, 0x0a, 0x29, 0x73, 0xb1, 0x64, 0x96, 0xa4, 0xe6, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xf1, 0xeb, 0x15, 0x24, 0xe9, 0x41, 0x65, 0x3d, 0x4b, 0x52,
	0x73, 0x83, 0xc0, 0x92, 0x42, 0xa2, 0x5c, 0x6c, 0xc5, 0xa9, 0x85, 0xf1, 0x99, 0x29, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0xac, 0xc5, 0xa9, 0x85, 0x9e, 0x29, 0x4a, 0x22, 0x5c, 0x42, 0xe8,
	0x06, 0x16, 0x17, 0x18, 0xf9, 0x71, 0x71, 0x3b, 0xe7, 0xe7, 0xe5, 0x05, 0xa7, 0x16, 0x95, 0x65,
	0x26, 0xa7, 0x0a, 0xd9, 0x73, 0xf1, 0xa1, 0x2a, 0x12, 0x12, 0x05, 0x59, 0x82, 0xe1, 0x12, 0x29,
	0x31, 0x6c, 0xc2, 0xc5, 0x05, 0x49, 0x6c, 0x60, 0xd7, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x8f, 0x3a, 0x35, 0x91, 0xdf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnServiceClient is the client API for ConnService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnServiceClient interface {
	DeliverMessage(ctx context.Context, in *DeliverMessageReq, opts ...grpc.CallOption) (*DeliverMessageResp, error)
}

type connServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnServiceClient(cc grpc.ClientConnInterface) ConnServiceClient {
	return &connServiceClient{cc}
}

func (c *connServiceClient) DeliverMessage(ctx context.Context, in *DeliverMessageReq, opts ...grpc.CallOption) (*DeliverMessageResp, error) {
	out := new(DeliverMessageResp)
	err := c.cc.Invoke(ctx, "/pb.ConnService/DeliverMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnServiceServer is the server API for ConnService service.
type ConnServiceServer interface {
	DeliverMessage(context.Context, *DeliverMessageReq) (*DeliverMessageResp, error)
}

// UnimplementedConnServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConnServiceServer struct {
}

func (*UnimplementedConnServiceServer) DeliverMessage(ctx context.Context, req *DeliverMessageReq) (*DeliverMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverMessage not implemented")
}

func RegisterConnServiceServer(s *grpc.Server, srv ConnServiceServer) {
	s.RegisterService(&_ConnService_serviceDesc, srv)
}

func _ConnService_DeliverMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServiceServer).DeliverMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ConnService/DeliverMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServiceServer).DeliverMessage(ctx, req.(*DeliverMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConnService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ConnService",
	HandlerType: (*ConnServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeliverMessage",
			Handler:    _ConnService_DeliverMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conn.proto",
}
