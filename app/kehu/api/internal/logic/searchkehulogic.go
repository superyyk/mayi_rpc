package logic

import (
	"context"

	"github.com/superyyk/mayi_rpc/app/kehu/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/kehu/api/internal/types"
	"github.com/superyyk/mayi_rpc/model"

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
