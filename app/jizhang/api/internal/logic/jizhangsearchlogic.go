package logic

import (
	"context"

	"mayi/app/jizhang/api/internal/svc"
	"mayi/app/jizhang/api/internal/types"
	"mayi/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type JizhangSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJizhangSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JizhangSearchLogic {
	return &JizhangSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JizhangSearchLogic) JizhangSearch(req *types.JizhangSearchReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	var jizhang []model.Jizhang

	l.svcCtx.Db.Table("jizhang").Where("puber=? AND status=?", req.Puber, 0).Order("id desc").Find(&jizhang)
	var jizhangleixing []model.JizhangTy
	l.svcCtx.Db.Table("jizhang_ty").Where("puber=? AND status=?", req.Puber, 0).Order("id desc").Find(&jizhangleixing)

	res := make(map[string]interface{})
	res["jizhang"] = jizhang
	res["jizhang_ty"] = jizhangleixing
	return &types.Resp{
		Res:  "200",
		Msg:  "success",
		Data: res,
	}, nil
}
