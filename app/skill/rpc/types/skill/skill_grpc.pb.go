// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: rpc/skill.proto

package skill

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

// SkillClient is the client API for Skill service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SkillClient interface {
	GetSkill(ctx context.Context, in *SkillRequest, opts ...grpc.CallOption) (*Rsp, error)
	SetSkill(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetRsp, error)
}

type skillClient struct {
	cc grpc.ClientConnInterface
}

func NewSkillClient(cc grpc.ClientConnInterface) SkillClient {
	return &skillClient{cc}
}

func (c *skillClient) GetSkill(ctx context.Context, in *SkillRequest, opts ...grpc.CallOption) (*Rsp, error) {
	out := new(Rsp)
	err := c.cc.Invoke(ctx, "/skill.Skill/getSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillClient) SetSkill(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetRsp, error) {
	out := new(SetRsp)
	err := c.cc.Invoke(ctx, "/skill.Skill/setSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SkillServer is the server API for Skill service.
// All implementations must embed UnimplementedSkillServer
// for forward compatibility
type SkillServer interface {
	GetSkill(context.Context, *SkillRequest) (*Rsp, error)
	SetSkill(context.Context, *SetReq) (*SetRsp, error)
	mustEmbedUnimplementedSkillServer()
}

// UnimplementedSkillServer must be embedded to have forward compatible implementations.
type UnimplementedSkillServer struct {
}

func (UnimplementedSkillServer) GetSkill(context.Context, *SkillRequest) (*Rsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkill not implemented")
}
func (UnimplementedSkillServer) SetSkill(context.Context, *SetReq) (*SetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSkill not implemented")
}
func (UnimplementedSkillServer) mustEmbedUnimplementedSkillServer() {}

// UnsafeSkillServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SkillServer will
// result in compilation errors.
type UnsafeSkillServer interface {
	mustEmbedUnimplementedSkillServer()
}

func RegisterSkillServer(s grpc.ServiceRegistrar, srv SkillServer) {
	s.RegisterService(&Skill_ServiceDesc, srv)
}

func _Skill_GetSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillServer).GetSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skill.Skill/getSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillServer).GetSkill(ctx, req.(*SkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Skill_SetSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillServer).SetSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skill.Skill/setSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillServer).SetSkill(ctx, req.(*SetReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Skill_ServiceDesc is the grpc.ServiceDesc for Skill service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Skill_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "skill.Skill",
	HandlerType: (*SkillServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSkill",
			Handler:    _Skill_GetSkill_Handler,
		},
		{
			MethodName: "setSkill",
			Handler:    _Skill_SetSkill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/skill.proto",
}
