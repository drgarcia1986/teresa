// Code generated by protoc-gen-go.
// source: pkg/protobuf/deploy/deploy.proto
// DO NOT EDIT!

/*
Package deploy is a generated protocol buffer package.

It is generated from these files:
	pkg/protobuf/deploy/deploy.proto

It has these top-level messages:
	DeployRequest
	DeployResponse
*/
package deploy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeployRequest struct {
	// Types that are valid to be assigned to Value:
	//	*DeployRequest_Info_
	//	*DeployRequest_File_
	Value isDeployRequest_Value `protobuf_oneof:"value"`
}

func (m *DeployRequest) Reset()                    { *m = DeployRequest{} }
func (m *DeployRequest) String() string            { return proto.CompactTextString(m) }
func (*DeployRequest) ProtoMessage()               {}
func (*DeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isDeployRequest_Value interface {
	isDeployRequest_Value()
}

type DeployRequest_Info_ struct {
	Info *DeployRequest_Info `protobuf:"bytes,1,opt,name=info,oneof"`
}
type DeployRequest_File_ struct {
	File *DeployRequest_File `protobuf:"bytes,2,opt,name=file,oneof"`
}

func (*DeployRequest_Info_) isDeployRequest_Value() {}
func (*DeployRequest_File_) isDeployRequest_Value() {}

func (m *DeployRequest) GetValue() isDeployRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *DeployRequest) GetInfo() *DeployRequest_Info {
	if x, ok := m.GetValue().(*DeployRequest_Info_); ok {
		return x.Info
	}
	return nil
}

func (m *DeployRequest) GetFile() *DeployRequest_File {
	if x, ok := m.GetValue().(*DeployRequest_File_); ok {
		return x.File
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DeployRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DeployRequest_OneofMarshaler, _DeployRequest_OneofUnmarshaler, _DeployRequest_OneofSizer, []interface{}{
		(*DeployRequest_Info_)(nil),
		(*DeployRequest_File_)(nil),
	}
}

func _DeployRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DeployRequest)
	// value
	switch x := m.Value.(type) {
	case *DeployRequest_Info_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Info); err != nil {
			return err
		}
	case *DeployRequest_File_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.File); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DeployRequest.Value has unexpected type %T", x)
	}
	return nil
}

func _DeployRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DeployRequest)
	switch tag {
	case 1: // value.info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeployRequest_Info)
		err := b.DecodeMessage(msg)
		m.Value = &DeployRequest_Info_{msg}
		return true, err
	case 2: // value.file
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeployRequest_File)
		err := b.DecodeMessage(msg)
		m.Value = &DeployRequest_File_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DeployRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DeployRequest)
	// value
	switch x := m.Value.(type) {
	case *DeployRequest_Info_:
		s := proto.Size(x.Info)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DeployRequest_File_:
		s := proto.Size(x.File)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DeployRequest_Info struct {
	App         string `protobuf:"bytes,1,opt,name=app" json:"app,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *DeployRequest_Info) Reset()                    { *m = DeployRequest_Info{} }
func (m *DeployRequest_Info) String() string            { return proto.CompactTextString(m) }
func (*DeployRequest_Info) ProtoMessage()               {}
func (*DeployRequest_Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *DeployRequest_Info) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *DeployRequest_Info) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type DeployRequest_File struct {
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (m *DeployRequest_File) Reset()                    { *m = DeployRequest_File{} }
func (m *DeployRequest_File) String() string            { return proto.CompactTextString(m) }
func (*DeployRequest_File) ProtoMessage()               {}
func (*DeployRequest_File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *DeployRequest_File) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

type DeployResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *DeployResponse) Reset()                    { *m = DeployResponse{} }
func (m *DeployResponse) String() string            { return proto.CompactTextString(m) }
func (*DeployResponse) ProtoMessage()               {}
func (*DeployResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeployResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*DeployRequest)(nil), "deploy.DeployRequest")
	proto.RegisterType((*DeployRequest_Info)(nil), "deploy.DeployRequest.Info")
	proto.RegisterType((*DeployRequest_File)(nil), "deploy.DeployRequest.File")
	proto.RegisterType((*DeployResponse)(nil), "deploy.DeployResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Deploy service

type DeployClient interface {
	Make(ctx context.Context, opts ...grpc.CallOption) (Deploy_MakeClient, error)
}

type deployClient struct {
	cc *grpc.ClientConn
}

func NewDeployClient(cc *grpc.ClientConn) DeployClient {
	return &deployClient{cc}
}

func (c *deployClient) Make(ctx context.Context, opts ...grpc.CallOption) (Deploy_MakeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Deploy_serviceDesc.Streams[0], c.cc, "/deploy.Deploy/Make", opts...)
	if err != nil {
		return nil, err
	}
	x := &deployMakeClient{stream}
	return x, nil
}

type Deploy_MakeClient interface {
	Send(*DeployRequest) error
	Recv() (*DeployResponse, error)
	grpc.ClientStream
}

type deployMakeClient struct {
	grpc.ClientStream
}

func (x *deployMakeClient) Send(m *DeployRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *deployMakeClient) Recv() (*DeployResponse, error) {
	m := new(DeployResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Deploy service

type DeployServer interface {
	Make(Deploy_MakeServer) error
}

func RegisterDeployServer(s *grpc.Server, srv DeployServer) {
	s.RegisterService(&_Deploy_serviceDesc, srv)
}

func _Deploy_Make_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeployServer).Make(&deployMakeServer{stream})
}

type Deploy_MakeServer interface {
	Send(*DeployResponse) error
	Recv() (*DeployRequest, error)
	grpc.ServerStream
}

type deployMakeServer struct {
	grpc.ServerStream
}

func (x *deployMakeServer) Send(m *DeployResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *deployMakeServer) Recv() (*DeployRequest, error) {
	m := new(DeployRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Deploy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deploy.Deploy",
	HandlerType: (*DeployServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Make",
			Handler:       _Deploy_Make_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/protobuf/deploy/deploy.proto",
}

func init() { proto.RegisterFile("pkg/protobuf/deploy/deploy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x8d, 0xa6, 0x2b, 0x9d, 0xaa, 0xc8, 0xa0, 0x52, 0x16, 0x0f, 0x4b, 0xf1, 0xd0, 0xd3,
	0xb6, 0xd4, 0x93, 0x1e, 0x55, 0x44, 0x0f, 0x5e, 0xf2, 0x0f, 0xb6, 0xed, 0xac, 0x86, 0x5d, 0x92,
	0xd8, 0x4d, 0x44, 0xff, 0xa8, 0xbf, 0x47, 0x32, 0x59, 0x41, 0x45, 0x7a, 0xca, 0x9b, 0xc7, 0xfb,
	0xde, 0x0c, 0x81, 0xc2, 0x35, 0xcf, 0x33, 0xb7, 0xb1, 0xde, 0x2e, 0x43, 0x3d, 0x5b, 0x93, 0x6b,
	0xed, 0x47, 0xff, 0x94, 0x6c, 0x63, 0x96, 0xa6, 0xc9, 0xa7, 0x80, 0xc3, 0x3b, 0x96, 0x8a, 0x5e,
	0x03, 0x75, 0x1e, 0xe7, 0x20, 0xb5, 0xa9, 0xed, 0x58, 0x14, 0x62, 0x3a, 0x5a, 0xe4, 0x65, 0x8f,
	0xfd, 0x0a, 0x95, 0x8f, 0xa6, 0xb6, 0x0f, 0x3b, 0x8a, 0x93, 0x91, 0xa8, 0x75, 0x4b, 0xe3, 0xdd,
	0x6d, 0xc4, 0xbd, 0x6e, 0x29, 0x12, 0x31, 0x99, 0x5f, 0x83, 0x8c, 0x0d, 0x78, 0x0c, 0x7b, 0x95,
	0x73, 0xbc, 0x6a, 0xa8, 0xa2, 0xc4, 0x02, 0x46, 0x6b, 0xea, 0x56, 0x1b, 0xed, 0xbc, 0xb6, 0x86,
	0x2b, 0x87, 0xea, 0xa7, 0x95, 0x9f, 0x83, 0x8c, 0x5d, 0x78, 0x02, 0x83, 0xd5, 0x4b, 0x30, 0x0d,
	0xd3, 0x07, 0x2a, 0x0d, 0x37, 0xfb, 0x30, 0x78, 0xab, 0xda, 0x40, 0x93, 0x0b, 0x38, 0xfa, 0x3e,
	0xa0, 0x73, 0xd6, 0x74, 0x84, 0x08, 0xd2, 0xd3, 0xbb, 0xef, 0xb7, 0xb1, 0x5e, 0xdc, 0x42, 0x96,
	0x52, 0x78, 0x05, 0xf2, 0xa9, 0x6a, 0x08, 0x4f, 0xff, 0x3d, 0x3f, 0x3f, 0xfb, 0x6b, 0xa7, 0xd2,
	0xa9, 0x98, 0x8b, 0x65, 0xc6, 0x5f, 0x7a, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x10, 0x1d,
	0x55, 0x76, 0x01, 0x00, 0x00,
}
