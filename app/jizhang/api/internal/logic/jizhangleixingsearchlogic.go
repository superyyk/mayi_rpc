package logic

import (
	"context"

	"mayi/app/jizhang/api/internal/svc"
	"mayi/app/jizhang/api/internal/types"
	"mayi/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type JizhangLeixingSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJizhangLeixingSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JizhangLeixingSearchLogic {
	return &JizhangLeixingSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JizhangLeixingSearchLogic) JizhangLeixingSearch(req *types.JizhangLeixingSearchReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	var jizhangleixing []model.JizhangTy
	l.svcCtx.Db.Table("jizhang_ty").Where("puber=? AND status=?", req.Puber, 0).Order("id desc").Find(&jizhangleixing)

	return &types.Resp{
		Res:  "200",
		Msg:  "success23433",
		Data: jizhangleixing,
	}, nil
}
