package logic

import (
	"context"

	"github.com/superyyk/mayi_rpc/app/workers/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/workers/api/internal/types"
	"github.com/superyyk/mayi_rpc/model"
	"github.com/superyyk/mayi_rpc/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddWorkertyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddWorkertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddWorkertyLogic {
	return &AddWorkertyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddWorkertyLogic) AddWorkerty(req *types.AddTyReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	uid, _ := tool.GetUuid(12)
	workerty := &model.Workerty{
		Name:  req.Name,
		Puber: req.Uid,
		Uid:   uid,
		Ty:    1,
	}
	if err := l.svcCtx.Db.Table("worker_ty").Create(&workerty).Error; err != nil {
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
