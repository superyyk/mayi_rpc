// Code generated by goctl. DO NOT EDIT.
// Source: auth.proto

package server

import (
	"context"

	"mayi/auth/rpc/internal/logic"
	"mayi/auth/rpc/internal/svc"
	"mayi/auth/rpc/types/auth"
)

type TokenServer struct {
	svcCtx *svc.ServiceContext
	auth.UnimplementedTokenServer
}

func NewTokenServer(svcCtx *svc.ServiceContext) *TokenServer {
	return &TokenServer{
		svcCtx: svcCtx,
	}
}

func (s *TokenServer) GetToken(ctx context.Context, in *auth.GetTokenRequest) (*auth.GetResponse, error) {
	l := logic.NewGetTokenLogic(ctx, s.svcCtx)
	return l.GetToken(in)
}

func (s *TokenServer) VerifyToken(ctx context.Context, in *auth.TokenRequest) (*auth.TokenResponse, error) {
	l := logic.NewVerifyTokenLogic(ctx, s.svcCtx)
	return l.VerifyToken(in)
}
