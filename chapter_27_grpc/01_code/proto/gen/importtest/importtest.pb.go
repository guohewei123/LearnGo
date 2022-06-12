// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: importtest.proto

package importtestpb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	common "trip/proto/gen/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_importtest_proto protoreflect.FileDescriptor

var file_importtest_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x42, 0x0a, 0x0a, 0x54, 0x65, 0x73,
	0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x28, 0x5a,
	0x26, 0x74, 0x72, 0x69, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x74, 0x65, 0x73, 0x74, 0x3b, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_importtest_proto_goTypes = []interface{}{
	(*empty.Empty)(nil),     // 0: google.protobuf.Empty
	(*common.Location)(nil), // 1: common.Location
}
var file_importtest_proto_depIdxs = []int32{
	0, // 0: importtest.TestImport.SayHello:input_type -> google.protobuf.Empty
	1, // 1: importtest.TestImport.SayHello:output_type -> common.Location
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_importtest_proto_init() }
func file_importtest_proto_init() {
	if File_importtest_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_importtest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_importtest_proto_goTypes,
		DependencyIndexes: file_importtest_proto_depIdxs,
	}.Build()
	File_importtest_proto = out.File
	file_importtest_proto_rawDesc = nil
	file_importtest_proto_goTypes = nil
	file_importtest_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestImportClient is the client API for TestImport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestImportClient interface {
	SayHello(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.Location, error)
}

type testImportClient struct {
	cc grpc.ClientConnInterface
}

func NewTestImportClient(cc grpc.ClientConnInterface) TestImportClient {
	return &testImportClient{cc}
}

func (c *testImportClient) SayHello(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.Location, error) {
	out := new(common.Location)
	err := c.cc.Invoke(ctx, "/importtest.TestImport/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestImportServer is the server API for TestImport service.
type TestImportServer interface {
	SayHello(context.Context, *empty.Empty) (*common.Location, error)
}

// UnimplementedTestImportServer can be embedded to have forward compatible implementations.
type UnimplementedTestImportServer struct {
}

func (*UnimplementedTestImportServer) SayHello(context.Context, *empty.Empty) (*common.Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterTestImportServer(s *grpc.Server, srv TestImportServer) {
	s.RegisterService(&_TestImport_serviceDesc, srv)
}

func _TestImport_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestImportServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/importtest.TestImport/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestImportServer).SayHello(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestImport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "importtest.TestImport",
	HandlerType: (*TestImportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _TestImport_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "importtest.proto",
}