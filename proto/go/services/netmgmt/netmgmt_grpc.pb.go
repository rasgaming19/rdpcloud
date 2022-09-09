// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: services/netmgmt/netmgmt.proto

package netmgmt

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

// NetmgmtClient is the client API for Netmgmt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetmgmtClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetUserLocalGroups(ctx context.Context, in *GetUserLocalGroupsRequest, opts ...grpc.CallOption) (*GetUserLocalGroupsResponse, error)
	ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordRequest, opts ...grpc.CallOption) (*ChangeUserPasswordResponse, error)
	AddUserToLocalGroup(ctx context.Context, in *AddUserToLocalGroupRequest, opts ...grpc.CallOption) (*AddUserToLocalGroupResponse, error)
	RemoveUserFromLocalGroup(ctx context.Context, in *RemoveUserFromLocalGroupRequest, opts ...grpc.CallOption) (*RemoveUserFromLocalGroupResponse, error)
	GetLocalGroups(ctx context.Context, in *GetLocalGroupsRequest, opts ...grpc.CallOption) (*GetLocalGroupsResponse, error)
	GetUsersInLocalGroup(ctx context.Context, in *GetUsersInLocalGroupRequest, opts ...grpc.CallOption) (*GetUsersInLocalGroupResponse, error)
}

type netmgmtClient struct {
	cc grpc.ClientConnInterface
}

func NewNetmgmtClient(cc grpc.ClientConnInterface) NetmgmtClient {
	return &netmgmtClient{cc}
}

func (c *netmgmtClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, "/services.netmgmt.Netmgmt/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netmgmtClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/services.netmgmt.Netmgmt/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netmgmtClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/services.netmgmt.Netmgmt/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netmgmtClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/services.netmgmt.Netmgmt/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netmgmtClient) GetUserLocalGroups(ctx context.Context, in *GetUserLocalGroupsRequest, opts ...grpc.CallOption) (*GetUserLocalGroupsResponse, error) {
	out := new(GetUserLocalGroupsResponse)
	err := c.cc.Invoke(ctx, "/services.netmgmt.Netmgmt/GetUserLocalGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netmgmtClient) ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordRequest, opts ...grpc.CallOption) (*ChangeUserPasswordResponse, error) {
	out := new(ChangeUserPasswordResponse)
	err := c.cc.Invoke(ctx, "/services.netmgmt.Netmgmt/ChangeUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netmgmtClient) AddUserToLocalGroup(ctx context.Context, in *AddUserToLocalGroupRequest, opts ...grpc.CallOption) (*AddUserToLocalGroupResponse, error) {
	out := new(AddUserToLocalGroupResponse)
	err := c.cc.Invoke(ctx, "/services.netmgmt.Netmgmt/AddUserToLocalGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netmgmtClient) RemoveUserFromLocalGroup(ctx context.Context, in *RemoveUserFromLocalGroupRequest, opts ...grpc.CallOption) (*RemoveUserFromLocalGroupResponse, error) {
	out := new(RemoveUserFromLocalGroupResponse)
	err := c.cc.Invoke(ctx, "/services.netmgmt.Netmgmt/RemoveUserFromLocalGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netmgmtClient) GetLocalGroups(ctx context.Context, in *GetLocalGroupsRequest, opts ...grpc.CallOption) (*GetLocalGroupsResponse, error) {
	out := new(GetLocalGroupsResponse)
	err := c.cc.Invoke(ctx, "/services.netmgmt.Netmgmt/GetLocalGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netmgmtClient) GetUsersInLocalGroup(ctx context.Context, in *GetUsersInLocalGroupRequest, opts ...grpc.CallOption) (*GetUsersInLocalGroupResponse, error) {
	out := new(GetUsersInLocalGroupResponse)
	err := c.cc.Invoke(ctx, "/services.netmgmt.Netmgmt/GetUsersInLocalGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetmgmtServer is the server API for Netmgmt service.
// All implementations must embed UnimplementedNetmgmtServer
// for forward compatibility
type NetmgmtServer interface {
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetUserLocalGroups(context.Context, *GetUserLocalGroupsRequest) (*GetUserLocalGroupsResponse, error)
	ChangeUserPassword(context.Context, *ChangeUserPasswordRequest) (*ChangeUserPasswordResponse, error)
	AddUserToLocalGroup(context.Context, *AddUserToLocalGroupRequest) (*AddUserToLocalGroupResponse, error)
	RemoveUserFromLocalGroup(context.Context, *RemoveUserFromLocalGroupRequest) (*RemoveUserFromLocalGroupResponse, error)
	GetLocalGroups(context.Context, *GetLocalGroupsRequest) (*GetLocalGroupsResponse, error)
	GetUsersInLocalGroup(context.Context, *GetUsersInLocalGroupRequest) (*GetUsersInLocalGroupResponse, error)
	mustEmbedUnimplementedNetmgmtServer()
}

// UnimplementedNetmgmtServer must be embedded to have forward compatible implementations.
type UnimplementedNetmgmtServer struct {
}

func (UnimplementedNetmgmtServer) AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedNetmgmtServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedNetmgmtServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedNetmgmtServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedNetmgmtServer) GetUserLocalGroups(context.Context, *GetUserLocalGroupsRequest) (*GetUserLocalGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLocalGroups not implemented")
}
func (UnimplementedNetmgmtServer) ChangeUserPassword(context.Context, *ChangeUserPasswordRequest) (*ChangeUserPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserPassword not implemented")
}
func (UnimplementedNetmgmtServer) AddUserToLocalGroup(context.Context, *AddUserToLocalGroupRequest) (*AddUserToLocalGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToLocalGroup not implemented")
}
func (UnimplementedNetmgmtServer) RemoveUserFromLocalGroup(context.Context, *RemoveUserFromLocalGroupRequest) (*RemoveUserFromLocalGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserFromLocalGroup not implemented")
}
func (UnimplementedNetmgmtServer) GetLocalGroups(context.Context, *GetLocalGroupsRequest) (*GetLocalGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocalGroups not implemented")
}
func (UnimplementedNetmgmtServer) GetUsersInLocalGroup(context.Context, *GetUsersInLocalGroupRequest) (*GetUsersInLocalGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersInLocalGroup not implemented")
}
func (UnimplementedNetmgmtServer) mustEmbedUnimplementedNetmgmtServer() {}

// UnsafeNetmgmtServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetmgmtServer will
// result in compilation errors.
type UnsafeNetmgmtServer interface {
	mustEmbedUnimplementedNetmgmtServer()
}

func RegisterNetmgmtServer(s grpc.ServiceRegistrar, srv NetmgmtServer) {
	s.RegisterService(&Netmgmt_ServiceDesc, srv)
}

func _Netmgmt_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmgmtServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.netmgmt.Netmgmt/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmgmtServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Netmgmt_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmgmtServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.netmgmt.Netmgmt/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmgmtServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Netmgmt_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmgmtServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.netmgmt.Netmgmt/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmgmtServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Netmgmt_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmgmtServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.netmgmt.Netmgmt/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmgmtServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Netmgmt_GetUserLocalGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserLocalGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmgmtServer).GetUserLocalGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.netmgmt.Netmgmt/GetUserLocalGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmgmtServer).GetUserLocalGroups(ctx, req.(*GetUserLocalGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Netmgmt_ChangeUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmgmtServer).ChangeUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.netmgmt.Netmgmt/ChangeUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmgmtServer).ChangeUserPassword(ctx, req.(*ChangeUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Netmgmt_AddUserToLocalGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToLocalGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmgmtServer).AddUserToLocalGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.netmgmt.Netmgmt/AddUserToLocalGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmgmtServer).AddUserToLocalGroup(ctx, req.(*AddUserToLocalGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Netmgmt_RemoveUserFromLocalGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserFromLocalGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmgmtServer).RemoveUserFromLocalGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.netmgmt.Netmgmt/RemoveUserFromLocalGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmgmtServer).RemoveUserFromLocalGroup(ctx, req.(*RemoveUserFromLocalGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Netmgmt_GetLocalGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocalGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmgmtServer).GetLocalGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.netmgmt.Netmgmt/GetLocalGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmgmtServer).GetLocalGroups(ctx, req.(*GetLocalGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Netmgmt_GetUsersInLocalGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersInLocalGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmgmtServer).GetUsersInLocalGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.netmgmt.Netmgmt/GetUsersInLocalGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmgmtServer).GetUsersInLocalGroup(ctx, req.(*GetUsersInLocalGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Netmgmt_ServiceDesc is the grpc.ServiceDesc for Netmgmt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Netmgmt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.netmgmt.Netmgmt",
	HandlerType: (*NetmgmtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _Netmgmt_AddUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Netmgmt_DeleteUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Netmgmt_GetUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Netmgmt_GetUser_Handler,
		},
		{
			MethodName: "GetUserLocalGroups",
			Handler:    _Netmgmt_GetUserLocalGroups_Handler,
		},
		{
			MethodName: "ChangeUserPassword",
			Handler:    _Netmgmt_ChangeUserPassword_Handler,
		},
		{
			MethodName: "AddUserToLocalGroup",
			Handler:    _Netmgmt_AddUserToLocalGroup_Handler,
		},
		{
			MethodName: "RemoveUserFromLocalGroup",
			Handler:    _Netmgmt_RemoveUserFromLocalGroup_Handler,
		},
		{
			MethodName: "GetLocalGroups",
			Handler:    _Netmgmt_GetLocalGroups_Handler,
		},
		{
			MethodName: "GetUsersInLocalGroup",
			Handler:    _Netmgmt_GetUsersInLocalGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/netmgmt/netmgmt.proto",
}
