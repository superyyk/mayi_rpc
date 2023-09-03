// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package server

import (
	"context"

	"mayi/sms/rpc/internal/logic"
	"mayi/sms/rpc/internal/svc"
	"mayi/sms/rpc/types/sms"
)

type SmsCodeServer struct {
	svcCtx *svc.ServiceContext
	sms.UnimplementedSmsCodeServer
}

func NewSmsCodeServer(svcCtx *svc.ServiceContext) *SmsCodeServer {
	return &SmsCodeServer{
		svcCtx: svcCtx,
	}
}

func (s *SmsCodeServer) GetCode(ctx context.Context, in *sms.GetCodeRequest) (*sms.GetResponse, error) {
	l := logic.NewGetCodeLogic(ctx, s.svcCtx)
	return l.GetCode(in)
}