// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/dataflow/v1beta3/messages.proto

package dataflowpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessagesV1Beta3Client is the client API for MessagesV1Beta3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagesV1Beta3Client interface {
	// Request the job status.
	//
	// To request the status of a job, we recommend using
	// `projects.locations.jobs.messages.list` with a [regional endpoint]
	// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints). Using
	// `projects.jobs.messages.list` is not recommended, as you can only request
	// the status of jobs that are running in `us-central1`.
	ListJobMessages(ctx context.Context, in *ListJobMessagesRequest, opts ...grpc.CallOption) (*ListJobMessagesResponse, error)
}

type messagesV1Beta3Client struct {
	cc grpc.ClientConnInterface
}

func NewMessagesV1Beta3Client(cc grpc.ClientConnInterface) MessagesV1Beta3Client {
	return &messagesV1Beta3Client{cc}
}

func (c *messagesV1Beta3Client) ListJobMessages(ctx context.Context, in *ListJobMessagesRequest, opts ...grpc.CallOption) (*ListJobMessagesResponse, error) {
	out := new(ListJobMessagesResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.dataflow.v1beta3.MessagesV1Beta3/ListJobMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagesV1Beta3Server is the server API for MessagesV1Beta3 service.
// All implementations must embed UnimplementedMessagesV1Beta3Server
// for forward compatibility
type MessagesV1Beta3Server interface {
	// Request the job status.
	//
	// To request the status of a job, we recommend using
	// `projects.locations.jobs.messages.list` with a [regional endpoint]
	// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints). Using
	// `projects.jobs.messages.list` is not recommended, as you can only request
	// the status of jobs that are running in `us-central1`.
	ListJobMessages(context.Context, *ListJobMessagesRequest) (*ListJobMessagesResponse, error)
	mustEmbedUnimplementedMessagesV1Beta3Server()
}

// UnimplementedMessagesV1Beta3Server must be embedded to have forward compatible implementations.
type UnimplementedMessagesV1Beta3Server struct {
}

func (UnimplementedMessagesV1Beta3Server) ListJobMessages(context.Context, *ListJobMessagesRequest) (*ListJobMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJobMessages not implemented")
}
func (UnimplementedMessagesV1Beta3Server) mustEmbedUnimplementedMessagesV1Beta3Server() {}

// UnsafeMessagesV1Beta3Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagesV1Beta3Server will
// result in compilation errors.
type UnsafeMessagesV1Beta3Server interface {
	mustEmbedUnimplementedMessagesV1Beta3Server()
}

func RegisterMessagesV1Beta3Server(s grpc.ServiceRegistrar, srv MessagesV1Beta3Server) {
	s.RegisterService(&MessagesV1Beta3_ServiceDesc, srv)
}

func _MessagesV1Beta3_ListJobMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesV1Beta3Server).ListJobMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.dataflow.v1beta3.MessagesV1Beta3/ListJobMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesV1Beta3Server).ListJobMessages(ctx, req.(*ListJobMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessagesV1Beta3_ServiceDesc is the grpc.ServiceDesc for MessagesV1Beta3 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagesV1Beta3_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.dataflow.v1beta3.MessagesV1Beta3",
	HandlerType: (*MessagesV1Beta3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListJobMessages",
			Handler:    _MessagesV1Beta3_ListJobMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/dataflow/v1beta3/messages.proto",
}
