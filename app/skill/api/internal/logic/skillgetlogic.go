package logic

import (
	"context"

	"github.com/superyyk/mayi_rpc/app/skill/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/skill/api/internal/types"
	"github.com/superyyk/mayi_rpc/app/skill/rpc/types/skill"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkillGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkillGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkillGetLogic {
	return &SkillGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkillGetLogic) SkillGet(req *types.GetReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SkillRpc.GetSkill(l.ctx, &skill.SkillRequest{
		Uid: req.Userid,
	})
	if err != nil {
		return &types.Resp{
			Res: "400",
			Msg: "rpc 调用失败",
		}, nil
	}

	return &types.Resp{
		Res:  "200",
		Msg:  "rpc success",
		Data: res.Orderid,
	}, nil
}
