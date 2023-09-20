package logic

import (
	"context"

	"github.com/superyyk/mayi_rpc/auth/rpc/internal/svc"
	"github.com/superyyk/mayi_rpc/auth/rpc/types/auth"
	"github.com/superyyk/mayi_rpc/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenLogic) GetToken(in *auth.GetTokenRequest) (*auth.GetResponse, error) {
	// todo: add your logic here and delete this line

	token := tool.GetJWT(24, in.Userid, "yyk*2012")
	return &auth.GetResponse{
		Token: token,
	}, nil
}
