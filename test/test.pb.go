// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: test.proto

package test

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	B string `protobuf:"bytes,1,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x22, 0x17, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a,
	0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x22, 0x18, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x62, 0x32, 0x33, 0x0a, 0x06, 0x54, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x29, 0x0a, 0x06, 0x4d, 0x79, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: test.Request
	(*Response)(nil), // 1: test.Response
}
var file_test_proto_depIdxs = []int32{
	0, // 0: test.Tester.MyTest:input_type -> test.Request
	1, // 1: test.Tester.MyTest:output_type -> test.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
			switch v := v.(*Response); i {
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
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TesterClient is the client API for Tester service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TesterClient interface {
	MyTest(ctx context.Context, in *Request, opts ...grpc.CallOption) (Tester_MyTestClient, error)
}

type testerClient struct {
	cc grpc.ClientConnInterface
}

func NewTesterClient(cc grpc.ClientConnInterface) TesterClient {
	return &testerClient{cc}
}

func (c *testerClient) MyTest(ctx context.Context, in *Request, opts ...grpc.CallOption) (Tester_MyTestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tester_serviceDesc.Streams[0], "/test.Tester/MyTest", opts...)
	if err != nil {
		return nil, err
	}
	x := &testerMyTestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tester_MyTestClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type testerMyTestClient struct {
	grpc.ClientStream
}

func (x *testerMyTestClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TesterServer is the server API for Tester service.
type TesterServer interface {
	MyTest(*Request, Tester_MyTestServer) error
}

// UnimplementedTesterServer can be embedded to have forward compatible implementations.
type UnimplementedTesterServer struct {
}

func (*UnimplementedTesterServer) MyTest(*Request, Tester_MyTestServer) error {
	return status.Errorf(codes.Unimplemented, "method MyTest not implemented")
}

func RegisterTesterServer(s *grpc.Server, srv TesterServer) {
	s.RegisterService(&_Tester_serviceDesc, srv)
}

func _Tester_MyTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TesterServer).MyTest(m, &testerMyTestServer{stream})
}

type Tester_MyTestServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type testerMyTestServer struct {
	grpc.ServerStream
}

func (x *testerMyTestServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Tester_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.Tester",
	HandlerType: (*TesterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MyTest",
			Handler:       _Tester_MyTest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "test.proto",
}
