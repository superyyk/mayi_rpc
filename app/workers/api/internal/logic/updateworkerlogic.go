package logic

import (
	"context"

	"github.com/superyyk/mayi_rpc/app/workers/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/workers/api/internal/types"
	"github.com/superyyk/mayi_rpc/auth/rpc/types/auth"
	"github.com/superyyk/mayi_rpc/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateworkerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateworkerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateworkerLogic {
	return &UpdateworkerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateworkerLogic) Updateworker(req *types.UpdateReq, token string) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	worker := &model.Worker{
		Name:       req.Name,
		Age:        req.Age,
		Salary:     req.Salary,
		Tel:        req.Tel,
		Ty:         req.Ty,
		Status:     req.Status,
		Tip:        req.Tip,
		Uid:        req.Uid,
		Workerty:   req.Workerty,
		Workertyid: req.Workertyid,
	}
	claims, _ := l.svcCtx.AuthRpc.VerifyToken(l.ctx, &auth.TokenRequest{
		Token: token,
	})

	row := l.svcCtx.Db.Table("workers").Where("uid=? AND puber=?", req.Uid, claims.Userid).Updates(map[string]interface{}{"name": req.Name, "age": req.Age, "status": req.Status, "tel": req.Tel, "tip": req.Tip, "salary": req.Salary, "ty": req.Ty, "workerty": req.Workerty, "workerty_id": req.Workertyid})
	if row.RowsAffected == 1 {
		return &types.Resp{
			Res: "200",
			Msg: "操作成功",
		}, nil
	}
	return &types.Resp{
		Res:  "400",
		Msg:  "操作失败",
		Data: worker,
	}, nil
}
