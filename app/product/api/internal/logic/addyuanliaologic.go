package logic

import (
	"context"

	"mayi/app/product/api/internal/svc"
	"mayi/app/product/api/internal/types"
	"mayi/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddYuanliaoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddYuanliaoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddYuanliaoLogic {
	return &AddYuanliaoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddYuanliaoLogic) AddYuanliao(req *types.YuanliaoAddReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	yuanliao := &model.Yuanliao{
		Name:  req.Name,
		Puber: req.Puder,
	}
	if err := l.svcCtx.Db.Table("yuanliao").Create(&yuanliao).Error; err != nil {
		return &types.Resp{
			Res: "400",
			Msg: "系统繁忙！",
		}, nil
	}

	return &types.Resp{
		Res: "200",
		Msg: "添加成功！",
	}, nil
}
