package logic

import (
	"context"

	"mayi/app/product/api/internal/svc"
	"mayi/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelGuigeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelGuigeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelGuigeLogic {
	return &DelGuigeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelGuigeLogic) DelGuige(req *types.GuigeDelReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	row := l.svcCtx.Db.Table("guige").Where("id=? AND puber=?", req.Id, req.Puder).Update("status", 1)
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
