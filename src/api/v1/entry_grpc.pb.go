// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: src/api/v1/entry.proto

package ezgRPC

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

// EntryServicesClient is the client API for EntryServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntryServicesClient interface {
	Insert(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*InsertResponse, error)
	Modify(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*ModifyResponse, error)
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*Entry, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Entries, error)
}

type entryServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewEntryServicesClient(cc grpc.ClientConnInterface) EntryServicesClient {
	return &entryServicesClient{cc}
}

func (c *entryServicesClient) Insert(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, "/api.v1.EntryServices/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServicesClient) Modify(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*ModifyResponse, error) {
	out := new(ModifyResponse)
	err := c.cc.Invoke(ctx, "/api.v1.EntryServices/Modify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServicesClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/api.v1.EntryServices/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryServicesClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Entries, error) {
	out := new(Entries)
	err := c.cc.Invoke(ctx, "/api.v1.EntryServices/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntryServicesServer is the server API for EntryServices service.
// All implementations must embed UnimplementedEntryServicesServer
// for forward compatibility
type EntryServicesServer interface {
	Insert(context.Context, *Entry) (*InsertResponse, error)
	Modify(context.Context, *Entry) (*ModifyResponse, error)
	Retrieve(context.Context, *RetrieveRequest) (*Entry, error)
	List(context.Context, *ListRequest) (*Entries, error)
	mustEmbedUnimplementedEntryServicesServer()
}

// UnimplementedEntryServicesServer must be embedded to have forward compatible implementations.
type UnimplementedEntryServicesServer struct {
}

func (UnimplementedEntryServicesServer) Insert(context.Context, *Entry) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedEntryServicesServer) Modify(context.Context, *Entry) (*ModifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modify not implemented")
}
func (UnimplementedEntryServicesServer) Retrieve(context.Context, *RetrieveRequest) (*Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedEntryServicesServer) List(context.Context, *ListRequest) (*Entries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedEntryServicesServer) mustEmbedUnimplementedEntryServicesServer() {}

// UnsafeEntryServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntryServicesServer will
// result in compilation errors.
type UnsafeEntryServicesServer interface {
	mustEmbedUnimplementedEntryServicesServer()
}

func RegisterEntryServicesServer(s grpc.ServiceRegistrar, srv EntryServicesServer) {
	s.RegisterService(&EntryServices_ServiceDesc, srv)
}

func _EntryServices_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServicesServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.EntryServices/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServicesServer).Insert(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryServices_Modify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServicesServer).Modify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.EntryServices/Modify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServicesServer).Modify(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryServices_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServicesServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.EntryServices/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServicesServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntryServices_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServicesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.EntryServices/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServicesServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EntryServices_ServiceDesc is the grpc.ServiceDesc for EntryServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntryServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.EntryServices",
	HandlerType: (*EntryServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _EntryServices_Insert_Handler,
		},
		{
			MethodName: "Modify",
			Handler:    _EntryServices_Modify_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _EntryServices_Retrieve_Handler,
		},
		{
			MethodName: "List",
			Handler:    _EntryServices_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/api/v1/entry.proto",
}
