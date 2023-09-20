package logic

import (
	"context"

	"github.com/superyyk/mayi_rpc/app/product/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductLogic) UpdateProduct(req *types.UpdateReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	row := l.svcCtx.Db.Table("products").Where("uid=?", req.Uid).Updates(map[string]interface{}{
		"name":      req.Name,
		"unit":      req.Unit,
		"guige":     req.Guige,
		"suply":     req.Suply,
		"pertime":   req.PerTime,
		"perhaocai": req.PerHaocai,
		"yuanliao":  req.Yuanliao,
		"price":     req.Price,
		"tip":       req.Tip,
		"imgs":      req.Imgs,
		"status":    req.Status,
	})
	if row.RowsAffected == 1 {
		return &types.Resp{
			Res: "200",
			Msg: "操作成功！",
		}, nil
	}
	return &types.Resp{
		Res: "400",
		Msg: "操作失败！稍后再试",
	}, nil
}
