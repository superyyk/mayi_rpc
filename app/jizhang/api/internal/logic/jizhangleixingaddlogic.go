package logic

import (
	"context"

	"github.com/superyyk/mayi_rpc/app/jizhang/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/jizhang/api/internal/types"
	"github.com/superyyk/mayi_rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type JizhangLeixingAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJizhangLeixingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JizhangLeixingAddLogic {
	return &JizhangLeixingAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JizhangLeixingAddLogic) JizhangLeixingAdd(req *types.JizhangLeixingAddReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	jizhangleixing := &model.JizhangTy{
		Name:  req.Name,
		Puber: req.Puber,
	}
	if err := l.svcCtx.Db.Table("jizhang_ty").Create(&jizhangleixing).Error; err != nil {
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
