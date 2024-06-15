// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: src/chat.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ChatService_CreateChat_FullMethodName                = "/src.ChatService/CreateChat"
	ChatService_GetChatById_FullMethodName               = "/src.ChatService/GetChatById"
	ChatService_GetChatsForUser_FullMethodName           = "/src.ChatService/GetChatsForUser"
	ChatService_UpdateChatName_FullMethodName            = "/src.ChatService/UpdateChatName"
	ChatService_AddParticipantToChat_FullMethodName      = "/src.ChatService/AddParticipantToChat"
	ChatService_DeleteParticipantFromChat_FullMethodName = "/src.ChatService/DeleteParticipantFromChat"
	ChatService_DeleteChatById_FullMethodName            = "/src.ChatService/DeleteChatById"
	ChatService_CreateMessage_FullMethodName             = "/src.ChatService/CreateMessage"
	ChatService_GetMessagesForChat_FullMethodName        = "/src.ChatService/GetMessagesForChat"
	ChatService_UpdateMessageText_FullMethodName         = "/src.ChatService/UpdateMessageText"
	ChatService_UpdateMessageStatus_FullMethodName       = "/src.ChatService/UpdateMessageStatus"
	ChatService_DeleteMessageById_FullMethodName         = "/src.ChatService/DeleteMessageById"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	// chats
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetChatById(ctx context.Context, in *GetChatByIdRequest, opts ...grpc.CallOption) (*GetChatByIdResponse, error)
	GetChatsForUser(ctx context.Context, in *GetChatsForUserRequest, opts ...grpc.CallOption) (*GetChatsForUserResponse, error)
	UpdateChatName(ctx context.Context, in *UpdateChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddParticipantToChat(ctx context.Context, in *AddParticipantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteParticipantFromChat(ctx context.Context, in *DeleteParticipantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteChatById(ctx context.Context, in *DeleteChatByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// messages
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetMessagesForChat(ctx context.Context, in *GetMessagesForChatRequest, opts ...grpc.CallOption) (*GetMessagesForChatResponse, error)
	UpdateMessageText(ctx context.Context, in *UpdateMessageTextRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateMessageStatus(ctx context.Context, in *UpdateMessageStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteMessageById(ctx context.Context, in *DeleteMessageByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_CreateChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChatById(ctx context.Context, in *GetChatByIdRequest, opts ...grpc.CallOption) (*GetChatByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChatByIdResponse)
	err := c.cc.Invoke(ctx, ChatService_GetChatById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChatsForUser(ctx context.Context, in *GetChatsForUserRequest, opts ...grpc.CallOption) (*GetChatsForUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChatsForUserResponse)
	err := c.cc.Invoke(ctx, ChatService_GetChatsForUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateChatName(ctx context.Context, in *UpdateChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_UpdateChatName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) AddParticipantToChat(ctx context.Context, in *AddParticipantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_AddParticipantToChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteParticipantFromChat(ctx context.Context, in *DeleteParticipantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_DeleteParticipantFromChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteChatById(ctx context.Context, in *DeleteChatByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_DeleteChatById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_CreateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetMessagesForChat(ctx context.Context, in *GetMessagesForChatRequest, opts ...grpc.CallOption) (*GetMessagesForChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessagesForChatResponse)
	err := c.cc.Invoke(ctx, ChatService_GetMessagesForChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateMessageText(ctx context.Context, in *UpdateMessageTextRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_UpdateMessageText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateMessageStatus(ctx context.Context, in *UpdateMessageStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_UpdateMessageStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteMessageById(ctx context.Context, in *DeleteMessageByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChatService_DeleteMessageById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	// chats
	CreateChat(context.Context, *CreateChatRequest) (*emptypb.Empty, error)
	GetChatById(context.Context, *GetChatByIdRequest) (*GetChatByIdResponse, error)
	GetChatsForUser(context.Context, *GetChatsForUserRequest) (*GetChatsForUserResponse, error)
	UpdateChatName(context.Context, *UpdateChatRequest) (*emptypb.Empty, error)
	AddParticipantToChat(context.Context, *AddParticipantRequest) (*emptypb.Empty, error)
	DeleteParticipantFromChat(context.Context, *DeleteParticipantRequest) (*emptypb.Empty, error)
	DeleteChatById(context.Context, *DeleteChatByIdRequest) (*emptypb.Empty, error)
	// messages
	CreateMessage(context.Context, *CreateMessageRequest) (*emptypb.Empty, error)
	GetMessagesForChat(context.Context, *GetMessagesForChatRequest) (*GetMessagesForChatResponse, error)
	UpdateMessageText(context.Context, *UpdateMessageTextRequest) (*emptypb.Empty, error)
	UpdateMessageStatus(context.Context, *UpdateMessageStatusRequest) (*emptypb.Empty, error)
	DeleteMessageById(context.Context, *DeleteMessageByIdRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) CreateChat(context.Context, *CreateChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatServiceServer) GetChatById(context.Context, *GetChatByIdRequest) (*GetChatByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatById not implemented")
}
func (UnimplementedChatServiceServer) GetChatsForUser(context.Context, *GetChatsForUserRequest) (*GetChatsForUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatsForUser not implemented")
}
func (UnimplementedChatServiceServer) UpdateChatName(context.Context, *UpdateChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChatName not implemented")
}
func (UnimplementedChatServiceServer) AddParticipantToChat(context.Context, *AddParticipantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddParticipantToChat not implemented")
}
func (UnimplementedChatServiceServer) DeleteParticipantFromChat(context.Context, *DeleteParticipantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteParticipantFromChat not implemented")
}
func (UnimplementedChatServiceServer) DeleteChatById(context.Context, *DeleteChatByIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChatById not implemented")
}
func (UnimplementedChatServiceServer) CreateMessage(context.Context, *CreateMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedChatServiceServer) GetMessagesForChat(context.Context, *GetMessagesForChatRequest) (*GetMessagesForChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesForChat not implemented")
}
func (UnimplementedChatServiceServer) UpdateMessageText(context.Context, *UpdateMessageTextRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessageText not implemented")
}
func (UnimplementedChatServiceServer) UpdateMessageStatus(context.Context, *UpdateMessageStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessageStatus not implemented")
}
func (UnimplementedChatServiceServer) DeleteMessageById(context.Context, *DeleteMessageByIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessageById not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChatById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChatById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatById(ctx, req.(*GetChatByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChatsForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatsForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatsForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChatsForUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatsForUser(ctx, req.(*GetChatsForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateChatName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateChatName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_UpdateChatName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateChatName(ctx, req.(*UpdateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_AddParticipantToChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddParticipantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddParticipantToChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_AddParticipantToChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddParticipantToChat(ctx, req.(*AddParticipantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteParticipantFromChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteParticipantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteParticipantFromChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_DeleteParticipantFromChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteParticipantFromChat(ctx, req.(*DeleteParticipantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteChatById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChatByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteChatById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_DeleteChatById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteChatById(ctx, req.(*DeleteChatByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetMessagesForChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesForChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetMessagesForChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetMessagesForChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetMessagesForChat(ctx, req.(*GetMessagesForChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateMessageText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateMessageText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_UpdateMessageText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateMessageText(ctx, req.(*UpdateMessageTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateMessageStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateMessageStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_UpdateMessageStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateMessageStatus(ctx, req.(*UpdateMessageStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteMessageById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteMessageById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_DeleteMessageById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteMessageById(ctx, req.(*DeleteMessageByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "src.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _ChatService_CreateChat_Handler,
		},
		{
			MethodName: "GetChatById",
			Handler:    _ChatService_GetChatById_Handler,
		},
		{
			MethodName: "GetChatsForUser",
			Handler:    _ChatService_GetChatsForUser_Handler,
		},
		{
			MethodName: "UpdateChatName",
			Handler:    _ChatService_UpdateChatName_Handler,
		},
		{
			MethodName: "AddParticipantToChat",
			Handler:    _ChatService_AddParticipantToChat_Handler,
		},
		{
			MethodName: "DeleteParticipantFromChat",
			Handler:    _ChatService_DeleteParticipantFromChat_Handler,
		},
		{
			MethodName: "DeleteChatById",
			Handler:    _ChatService_DeleteChatById_Handler,
		},
		{
			MethodName: "CreateMessage",
			Handler:    _ChatService_CreateMessage_Handler,
		},
		{
			MethodName: "GetMessagesForChat",
			Handler:    _ChatService_GetMessagesForChat_Handler,
		},
		{
			MethodName: "UpdateMessageText",
			Handler:    _ChatService_UpdateMessageText_Handler,
		},
		{
			MethodName: "UpdateMessageStatus",
			Handler:    _ChatService_UpdateMessageStatus_Handler,
		},
		{
			MethodName: "DeleteMessageById",
			Handler:    _ChatService_DeleteMessageById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/chat.proto",
}