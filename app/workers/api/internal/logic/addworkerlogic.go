package logic

import (
	"context"
	"time"

	"mayi/app/workers/api/internal/svc"
	"mayi/app/workers/api/internal/types"
	"mayi/model"
	"mayi/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddworkerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddworkerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddworkerLogic {
	return &AddworkerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddworkerLogic) Addworker(req *types.AddReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	uid, _ := tool.GetUuid(12)
	worker := new(model.Worker)
	worker.Name = req.Name
	worker.Age = req.Age
	worker.Tel = req.Tel
	worker.Puber = req.Puber
	worker.Uid = uid
	worker.Salary = req.Salary
	worker.Time = time.Now().Unix()
	worker.Tip = req.Tip
	worker.Salary = req.Salary
	worker.Workertyid = req.Workertyid
	worker.Workerty = req.Workerty
	//var worker worker_model.worker
	//l.svcCtx.workerModel.Insert(l.ctx, worker)
	if err := l.svcCtx.Db.Table("workers").Create(&worker).Error; err != nil {
		logx.Errorf("创建客户失败：", err)
		return &types.Resp{
			Res:  "400",
			Msg:  "添加失败",
			Data: worker,
		}, nil
	}
	return &types.Resp{
		Res: "200",
		Msg: "添加成功",
	}, nil
}
