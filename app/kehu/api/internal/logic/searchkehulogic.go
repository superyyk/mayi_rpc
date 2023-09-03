package logic

import (
	"context"

	"mayi/app/kehu/api/internal/svc"
	"mayi/app/kehu/api/internal/types"
	"mayi/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchKehuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchKehuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchKehuLogic {
	return &SearchKehuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchKehuLogic) SearchKehu(req *types.SearchReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	var kehu []*model.Kehu
	l.svcCtx.Db.Table("kehu").Where("puber=? AND status=?", req.Uid, 0).Order("id desc").Find(&kehu)

	return &types.Resp{
		Res:  "200",
		Msg:  "success",
		Data: kehu,
	}, nil
}
