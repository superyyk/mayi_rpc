package logic

import (
	"context"

	"mayi/app/jizhang/api/internal/svc"
	"mayi/app/jizhang/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JizhangUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJizhangUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JizhangUpdateLogic {
	return &JizhangUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JizhangUpdateLogic) JizhangUpdate(req *types.JizhangUpdateReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	row := l.svcCtx.Db.Table("jizhang").Where("id=?", req.Id).Updates(map[string]interface{}{"leixing": req.Leixing, "leixing_id": req.Leixingid, "status": req.Status, "num": req.Num, "time": req.Time, "Tip": req.Tip})
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
