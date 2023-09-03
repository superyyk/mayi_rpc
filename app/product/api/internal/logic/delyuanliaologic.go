package logic

import (
	"context"

	"mayi/app/product/api/internal/svc"
	"mayi/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelYuanliaoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelYuanliaoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelYuanliaoLogic {
	return &DelYuanliaoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelYuanliaoLogic) DelYuanliao(req *types.YuanliaoReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	row := l.svcCtx.Db.Table("yuanliao").Where("id=? AND puber=?", req.Id, req.Puder).Update("status", 1)
	if row.RowsAffected == 1 {
		return &types.Resp{
			Res: "200",
			Msg: "删除成功",
		}, nil
	}
	
	return &types.Resp{
		Res: "400",
		Msg: "系统繁忙",
	}, nil
}
