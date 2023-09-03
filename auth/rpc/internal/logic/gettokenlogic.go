package logic

import (
	"context"

	"mayi/auth/rpc/internal/svc"
	"mayi/auth/rpc/types/auth"
	"mayi/tool"

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
