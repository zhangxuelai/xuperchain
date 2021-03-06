// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package xuperp2p

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

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xa8, 0x28, 0x2d, 0x48, 0x2d, 0x2a,
	0x30, 0x2a, 0x90, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x85, 0x48, 0x18, 0x85, 0x70,
	0x71, 0x15, 0x18, 0x15, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xb9, 0x71, 0xf1, 0x05,
	0xa7, 0xe6, 0xa5, 0x04, 0x18, 0x15, 0xf8, 0x42, 0x54, 0x09, 0x89, 0xe9, 0xc1, 0x74, 0xea, 0x45,
	0x80, 0x18, 0x50, 0x71, 0x29, 0x1c, 0xe2, 0x4a, 0x0c, 0x1a, 0x8c, 0x06, 0x8c, 0x49, 0x6c, 0x60,
	0xc3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x42, 0x01, 0xb2, 0xd1, 0x85, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// P2PServiceClient is the client API for P2PService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type P2PServiceClient interface {
	SendP2PMessage(ctx context.Context, opts ...grpc.CallOption) (P2PService_SendP2PMessageClient, error)
}

type p2PServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewP2PServiceClient(cc grpc.ClientConnInterface) P2PServiceClient {
	return &p2PServiceClient{cc}
}

func (c *p2PServiceClient) SendP2PMessage(ctx context.Context, opts ...grpc.CallOption) (P2PService_SendP2PMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_P2PService_serviceDesc.Streams[0], "/xuperp2p.p2pService/SendP2pMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &p2PServiceSendP2PMessageClient{stream}
	return x, nil
}

type P2PService_SendP2PMessageClient interface {
	Send(*XuperMessage) error
	Recv() (*XuperMessage, error)
	grpc.ClientStream
}

type p2PServiceSendP2PMessageClient struct {
	grpc.ClientStream
}

func (x *p2PServiceSendP2PMessageClient) Send(m *XuperMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *p2PServiceSendP2PMessageClient) Recv() (*XuperMessage, error) {
	m := new(XuperMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// P2PServiceServer is the server API for P2PService service.
type P2PServiceServer interface {
	SendP2PMessage(P2PService_SendP2PMessageServer) error
}

// UnimplementedP2PServiceServer can be embedded to have forward compatible implementations.
type UnimplementedP2PServiceServer struct {
}

func (*UnimplementedP2PServiceServer) SendP2PMessage(srv P2PService_SendP2PMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendP2PMessage not implemented")
}

func RegisterP2PServiceServer(s *grpc.Server, srv P2PServiceServer) {
	s.RegisterService(&_P2PService_serviceDesc, srv)
}

func _P2PService_SendP2PMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(P2PServiceServer).SendP2PMessage(&p2PServiceSendP2PMessageServer{stream})
}

type P2PService_SendP2PMessageServer interface {
	Send(*XuperMessage) error
	Recv() (*XuperMessage, error)
	grpc.ServerStream
}

type p2PServiceSendP2PMessageServer struct {
	grpc.ServerStream
}

func (x *p2PServiceSendP2PMessageServer) Send(m *XuperMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *p2PServiceSendP2PMessageServer) Recv() (*XuperMessage, error) {
	m := new(XuperMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _P2PService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xuperp2p.p2pService",
	HandlerType: (*P2PServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendP2pMessage",
			Handler:       _P2PService_SendP2PMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server.proto",
}
