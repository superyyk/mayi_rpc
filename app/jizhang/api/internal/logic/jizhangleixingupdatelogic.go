package logic

import (
	"context"

	"github.com/superyyk/mayi_rpc/app/jizhang/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/jizhang/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JizhangLeixingUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJizhangLeixingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JizhangLeixingUpdateLogic {
	return &JizhangLeixingUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JizhangLeixingUpdateLogic) JizhangLeixingUpdate(req *types.JizhangLeixingUpdateReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	row := l.svcCtx.Db.Table("jizhang_ty").Where("id=?", req.Id).Updates(map[string]interface{}{"name": req.Name, "status": req.Status})
	if row.RowsAffected == 1 {
		return &types.Resp{
			Res: "200",
			Msg: "操作成功",
		}, nil
	}
	return &types.Resp{
		Res:  "400",
		Msg:  "操作失败",
		Data: "",
	}, nil
}
