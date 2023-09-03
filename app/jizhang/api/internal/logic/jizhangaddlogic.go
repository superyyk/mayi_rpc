package logic

import (
	"context"
	"time"

	"mayi/app/jizhang/api/internal/svc"
	"mayi/app/jizhang/api/internal/types"
	"mayi/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type JizhangAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJizhangAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JizhangAddLogic {
	return &JizhangAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JizhangAddLogic) JizhangAdd(req *types.JizhangAddReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	jizhang := &model.Jizhang{
		Leixing:   req.Leixing,
		Leixingid: req.Leixingid,
		Puber:     req.Puber,
		Ty:        req.Ty,
		Num:       req.Num,
		Time:      time.Now().Unix(),
		Tip:       req.Tip,
	}
	if err := l.svcCtx.Db.Table("jizhang").Create(&jizhang).Error; err != nil {
		return &types.Resp{
			Res: "400",
			Msg: "添加失败",
		}, nil
	}
	return &types.Resp{
		Res: "200",
		Msg: "添加成功",
	}, nil
}
