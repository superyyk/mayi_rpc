package logic

import (
	"context"

	"mayi/app/liucheng/api/internal/svc"
	"mayi/app/liucheng/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StepUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStepUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StepUpdateLogic {
	return &StepUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StepUpdateLogic) StepUpdate(req *types.StepUpdateReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	row := l.svcCtx.Db.Table("make_rule").Where("id=?", req.Id).Updates(map[string]interface{}{"name": req.Name, "maker": req.Maker, "status": req.Status, "num": req.Num, "maketime": req.MakeTime, "bad": req.Bad})
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
