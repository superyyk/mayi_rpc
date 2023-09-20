package logic

import (
	"context"
	"time"

	"github.com/superyyk/mayi_rpc/app/liucheng/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/liucheng/api/internal/types"
	"github.com/superyyk/mayi_rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type StepAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStepAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StepAddLogic {
	return &StepAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StepAddLogic) StepAdd(req *types.StepAddReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	liucheng := &model.Liucheng{
		Name:      req.Name,
		ProductId: req.Productid,
		Maker:     req.Maker,
		Num:       req.Num,
		Maketime:  req.MakeTime,
		Bad:       req.Bad,
		Time:      time.Now().Unix(),
		Status:    req.Status,
	}
	if err := l.svcCtx.Db.Table("make_rule").Create(&liucheng).Error; err != nil {
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
