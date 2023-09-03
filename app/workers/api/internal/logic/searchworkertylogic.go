package logic

import (
	"context"

	"mayi/app/workers/api/internal/svc"
	"mayi/app/workers/api/internal/types"
	"mayi/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type SearchWorkertyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchWorkertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchWorkertyLogic {
	return &SearchWorkertyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchWorkertyLogic) SearchWorkerty(req *types.SearchTyReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	var myworkerty []model.Workerty
	var workerty []model.Workerty
	_ = mr.Finish(func() (err error) {
		l.svcCtx.Db.Table("worker_ty").Where("puber=? AND status=? AND ty=?", req.Uid, 0, 1).Order("id desc").Find(&myworkerty)

		return nil
	}, func() (err error) {
		l.svcCtx.Db.Table("worker_ty").Where("status=? AND ty=?", 0, 0).Order("sort asc").Find(&workerty)
		return nil
	})

	res := make(map[string]interface{})
	res["workerty"] = workerty
	res["myworkerty"] = myworkerty
	return &types.Resp{
		Res:  "200",
		Msg:  "success",
		Data: res,
	}, nil
}
