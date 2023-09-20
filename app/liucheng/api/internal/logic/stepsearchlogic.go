package logic

import (
	"context"

	"github.com/superyyk/mayi_rpc/app/liucheng/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/liucheng/api/internal/types"
	"github.com/superyyk/mayi_rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type StepSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStepSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StepSearchLogic {
	return &StepSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StepSearchLogic) StepSearch(req *types.StepSearchReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	var liucheng []model.Liucheng
	l.svcCtx.Db.Table("make_rule").Where("productid=? AND status=?", req.Productid, 0).Order("id asc").Find(&liucheng)
	res := make(map[string]interface{})
	res["liucheng"] = liucheng
	return &types.Resp{
		Res:  "200",
		Msg:  "success",
		Data: res,
	}, nil
}
