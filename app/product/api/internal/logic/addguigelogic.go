package logic

import (
	"context"

	"github.com/superyyk/mayi_rpc/app/product/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/product/api/internal/types"
	"github.com/superyyk/mayi_rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGuigeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGuigeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGuigeLogic {
	return &AddGuigeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGuigeLogic) AddGuige(req *types.GuigeAddReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	guige := &model.Guige{
		Name:  req.Name,
		Puber: req.Puder,
	}
	if err := l.svcCtx.Db.Table("guige").Create(&guige).Error; err != nil {
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
