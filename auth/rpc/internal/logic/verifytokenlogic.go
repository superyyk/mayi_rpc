package logic

import (
	"context"
	"errors"

	"github.com/superyyk/mayi_rpc/auth/rpc/internal/svc"
	"github.com/superyyk/mayi_rpc/auth/rpc/types/auth"
	"github.com/superyyk/mayi_rpc/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyTokenLogic) VerifyToken(in *auth.TokenRequest) (*auth.TokenResponse, error) {
	// todo: add your logic here and delete this line

	claims, err := tool.ParseToken(in.Token)
	if err != nil {
		return nil, errors.New("token err")
	}
	return &auth.TokenResponse{
		Userid: claims.Id,
	}, nil
}
