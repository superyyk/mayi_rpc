package logic

import (
	"context"

	"mayi/app/workers/api/internal/svc"
	"mayi/app/workers/api/internal/types"
	"mayi/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchworkerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchworkerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchworkerLogic {
	return &SearchworkerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchworkerLogic) Searchworker(req *types.SearchReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	var worker []*model.Worker

	l.svcCtx.Db.Table("workers").Where("puber=? AND status=?", req.Uid, 0).Order("id desc").Find(&worker)
	// for key, val := range worker {
	// 	var workerty []*model.Workerty
	// 	l.svcCtx.Db.Table("worker_ty").Where("status=? OR uid=? ", 0, val.Workertyid).Find(&workerty)
	// 	worker[key].Workerty = workerty
	// }
	var workerty []*model.Workerty
	var myworkerty []*model.Workerty
	l.svcCtx.Db.Table("worker_ty").Where("status=? AND ty=?", 0, 0).Order("sort asc").Find(&workerty)
	for key, v := range worker {
		var wy []*model.Workerty
		l.svcCtx.Db.Table("worker_ty").Where("uid=?", v.Workertyid).First(&wy)
		if len(wy) > 0 {
			worker[key].Workerty = wy[0].Name
		}

	}
	l.svcCtx.Db.Table("worker_ty").Where("status=? AND puber=? AND ty=?", 0, req.Uid, 1).Order("id desc").Find(&myworkerty)
	res := make(map[string]interface{})
	res["workers"] = worker
	res["workerty"] = workerty
	res["myworkerty"] = myworkerty
	return &types.Resp{
		Res:  "200",
		Msg:  "success",
		Data: res,
	}, nil
}
