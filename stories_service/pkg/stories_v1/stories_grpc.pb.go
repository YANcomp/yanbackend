// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: stories.proto

package stories_v1

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

// StoriesV1Client is the client API for StoriesV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoriesV1Client interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type storiesV1Client struct {
	cc grpc.ClientConnInterface
}

func NewStoriesV1Client(cc grpc.ClientConnInterface) StoriesV1Client {
	return &storiesV1Client{cc}
}

func (c *storiesV1Client) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/stories_v1.StoriesV1/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storiesV1Client) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/stories_v1.StoriesV1/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoriesV1Server is the server API for StoriesV1 service.
// All implementations must embed UnimplementedStoriesV1Server
// for forward compatibility
type StoriesV1Server interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedStoriesV1Server()
}

// UnimplementedStoriesV1Server must be embedded to have forward compatible implementations.
type UnimplementedStoriesV1Server struct {
}

func (UnimplementedStoriesV1Server) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStoriesV1Server) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStoriesV1Server) mustEmbedUnimplementedStoriesV1Server() {}

// UnsafeStoriesV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoriesV1Server will
// result in compilation errors.
type UnsafeStoriesV1Server interface {
	mustEmbedUnimplementedStoriesV1Server()
}

func RegisterStoriesV1Server(s grpc.ServiceRegistrar, srv StoriesV1Server) {
	s.RegisterService(&StoriesV1_ServiceDesc, srv)
}

func _StoriesV1_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesV1Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories_v1.StoriesV1/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesV1Server).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoriesV1_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesV1Server).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories_v1.StoriesV1/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesV1Server).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoriesV1_ServiceDesc is the grpc.ServiceDesc for StoriesV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoriesV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stories_v1.StoriesV1",
	HandlerType: (*StoriesV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _StoriesV1_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StoriesV1_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stories.proto",
}