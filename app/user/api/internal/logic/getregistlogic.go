package logic

import (
	"context"

	"mayi/app/user/api/internal/svc"
	"mayi/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRegistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRegistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegistLogic {
	return &GetRegistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRegistLogic) GetRegist(req *types.RegistReq) (resp *types.RegistResp, err error) {
	// todo: add your logic here and delete this line

	return
}
