// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package server

import (
	"context"

	"github.com/superyyk/mayi_rpc/sms/rpc/internal/logic"
	"github.com/superyyk/mayi_rpc/sms/rpc/internal/svc"
	"github.com/superyyk/mayi_rpc/sms/rpc/types/sms"
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
