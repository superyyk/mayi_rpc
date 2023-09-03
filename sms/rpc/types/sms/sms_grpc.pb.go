// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: rpc/sms.proto

package sms

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

// SmsCodeClient is the client API for SmsCode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsCodeClient interface {
	GetCode(ctx context.Context, in *GetCodeRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type smsCodeClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsCodeClient(cc grpc.ClientConnInterface) SmsCodeClient {
	return &smsCodeClient{cc}
}

func (c *smsCodeClient) GetCode(ctx context.Context, in *GetCodeRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/sms.SmsCode/getCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsCodeServer is the server API for SmsCode service.
// All implementations must embed UnimplementedSmsCodeServer
// for forward compatibility
type SmsCodeServer interface {
	GetCode(context.Context, *GetCodeRequest) (*GetResponse, error)
	mustEmbedUnimplementedSmsCodeServer()
}

// UnimplementedSmsCodeServer must be embedded to have forward compatible implementations.
type UnimplementedSmsCodeServer struct {
}

func (UnimplementedSmsCodeServer) GetCode(context.Context, *GetCodeRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCode not implemented")
}
func (UnimplementedSmsCodeServer) mustEmbedUnimplementedSmsCodeServer() {}

// UnsafeSmsCodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsCodeServer will
// result in compilation errors.
type UnsafeSmsCodeServer interface {
	mustEmbedUnimplementedSmsCodeServer()
}

func RegisterSmsCodeServer(s grpc.ServiceRegistrar, srv SmsCodeServer) {
	s.RegisterService(&SmsCode_ServiceDesc, srv)
}

func _SmsCode_GetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsCodeServer).GetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.SmsCode/getCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsCodeServer).GetCode(ctx, req.(*GetCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmsCode_ServiceDesc is the grpc.ServiceDesc for SmsCode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmsCode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sms.SmsCode",
	HandlerType: (*SmsCodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCode",
			Handler:    _SmsCode_GetCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/sms.proto",
}
