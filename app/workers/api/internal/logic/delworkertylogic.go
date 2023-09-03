package logic

import (
	"context"

	"mayi/app/workers/api/internal/svc"
	"mayi/app/workers/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelWorkertyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelWorkertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelWorkertyLogic {
	return &DelWorkertyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelWorkertyLogic) DelWorkerty(req *types.DelTyReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	row := l.svcCtx.Db.Table("worker_ty").Where("uid=?", req.Uid).Update("status", 1)
	if row.RowsAffected == 1 {
		return &types.Resp{
			Res: "200",
			Msg: "修改成功",
		}, nil
	}
	return &types.Resp{
		Res: "400",
		Msg: "修改失败",
	}, nil
}
