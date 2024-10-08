// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: apis/llmgapi/v1/openai/service.proto

package openai

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

const (
	OpenAIService_CreateChatCompletion_FullMethodName       = "/apis.llmgapi.v1.openai.OpenAIService/CreateChatCompletion"
	OpenAIService_CreateChatCompletionStream_FullMethodName = "/apis.llmgapi.v1.openai.OpenAIService/CreateChatCompletionStream"
)

// OpenAIServiceClient is the client API for OpenAIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenAIServiceClient interface {
	// CreateChatCompletion generates a model response for the given chat conversation.
	CreateChatCompletion(ctx context.Context, in *CreateChatCompletionRequest, opts ...grpc.CallOption) (*CreateChatCompletionResponse, error)
	// CreateChatCompletionStream generates a streaming model response for the given chat conversation.
	// It returns a stream of ChatCompletionChunk messages.
	CreateChatCompletionStream(ctx context.Context, in *CreateChatCompletionStreamRequest, opts ...grpc.CallOption) (OpenAIService_CreateChatCompletionStreamClient, error)
}

type openAIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenAIServiceClient(cc grpc.ClientConnInterface) OpenAIServiceClient {
	return &openAIServiceClient{cc}
}

func (c *openAIServiceClient) CreateChatCompletion(ctx context.Context, in *CreateChatCompletionRequest, opts ...grpc.CallOption) (*CreateChatCompletionResponse, error) {
	out := new(CreateChatCompletionResponse)
	err := c.cc.Invoke(ctx, OpenAIService_CreateChatCompletion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openAIServiceClient) CreateChatCompletionStream(ctx context.Context, in *CreateChatCompletionStreamRequest, opts ...grpc.CallOption) (OpenAIService_CreateChatCompletionStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &OpenAIService_ServiceDesc.Streams[0], OpenAIService_CreateChatCompletionStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &openAIServiceCreateChatCompletionStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OpenAIService_CreateChatCompletionStreamClient interface {
	Recv() (*CreateChatCompletionStreamResponse, error)
	grpc.ClientStream
}

type openAIServiceCreateChatCompletionStreamClient struct {
	grpc.ClientStream
}

func (x *openAIServiceCreateChatCompletionStreamClient) Recv() (*CreateChatCompletionStreamResponse, error) {
	m := new(CreateChatCompletionStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OpenAIServiceServer is the server API for OpenAIService service.
// All implementations must embed UnimplementedOpenAIServiceServer
// for forward compatibility
type OpenAIServiceServer interface {
	// CreateChatCompletion generates a model response for the given chat conversation.
	CreateChatCompletion(context.Context, *CreateChatCompletionRequest) (*CreateChatCompletionResponse, error)
	// CreateChatCompletionStream generates a streaming model response for the given chat conversation.
	// It returns a stream of ChatCompletionChunk messages.
	CreateChatCompletionStream(*CreateChatCompletionStreamRequest, OpenAIService_CreateChatCompletionStreamServer) error
	mustEmbedUnimplementedOpenAIServiceServer()
}

// UnimplementedOpenAIServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOpenAIServiceServer struct {
}

func (UnimplementedOpenAIServiceServer) CreateChatCompletion(context.Context, *CreateChatCompletionRequest) (*CreateChatCompletionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChatCompletion not implemented")
}
func (UnimplementedOpenAIServiceServer) CreateChatCompletionStream(*CreateChatCompletionStreamRequest, OpenAIService_CreateChatCompletionStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateChatCompletionStream not implemented")
}
func (UnimplementedOpenAIServiceServer) mustEmbedUnimplementedOpenAIServiceServer() {}

// UnsafeOpenAIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenAIServiceServer will
// result in compilation errors.
type UnsafeOpenAIServiceServer interface {
	mustEmbedUnimplementedOpenAIServiceServer()
}

func RegisterOpenAIServiceServer(s grpc.ServiceRegistrar, srv OpenAIServiceServer) {
	s.RegisterService(&OpenAIService_ServiceDesc, srv)
}

func _OpenAIService_CreateChatCompletion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatCompletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenAIServiceServer).CreateChatCompletion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenAIService_CreateChatCompletion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenAIServiceServer).CreateChatCompletion(ctx, req.(*CreateChatCompletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenAIService_CreateChatCompletionStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateChatCompletionStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OpenAIServiceServer).CreateChatCompletionStream(m, &openAIServiceCreateChatCompletionStreamServer{stream})
}

type OpenAIService_CreateChatCompletionStreamServer interface {
	Send(*CreateChatCompletionStreamResponse) error
	grpc.ServerStream
}

type openAIServiceCreateChatCompletionStreamServer struct {
	grpc.ServerStream
}

func (x *openAIServiceCreateChatCompletionStreamServer) Send(m *CreateChatCompletionStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// OpenAIService_ServiceDesc is the grpc.ServiceDesc for OpenAIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenAIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apis.llmgapi.v1.openai.OpenAIService",
	HandlerType: (*OpenAIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChatCompletion",
			Handler:    _OpenAIService_CreateChatCompletion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateChatCompletionStream",
			Handler:       _OpenAIService_CreateChatCompletionStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "apis/llmgapi/v1/openai/service.proto",
}
