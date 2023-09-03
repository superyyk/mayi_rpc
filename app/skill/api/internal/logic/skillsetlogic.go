package logic

import (
	"context"

	"mayi/app/skill/api/internal/svc"
	"mayi/app/skill/api/internal/types"
	"mayi/app/skill/rpc/types/skill"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkillSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkillSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkillSetLogic {
	return &SkillSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkillSetLogic) SkillSet(req *types.SetReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SkillRpc.SetSkill(l.ctx, &skill.SetReq{
		Num:    int64(req.Num),
		Userid: req.Userid,
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
		Data: res.Ok,
	}, nil
}
