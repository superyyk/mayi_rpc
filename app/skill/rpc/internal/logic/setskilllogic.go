package logic

import (
	"context"

	"mayi/app/skill/rpc/internal/svc"
	"mayi/app/skill/rpc/types/skill"
	"mayi/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSkillLogic {
	return &SetSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetSkillLogic) SetSkill(in *skill.SetReq) (*skill.SetRsp, error) {
	// todo: add your logic here and delete this line
	for i := 0; i < int(in.Num); i++ {
		uid, _ := tool.GetUuid(12)
		l.svcCtx.Rdb.LPush(l.ctx, "skill", uid)
	}

	return &skill.SetRsp{
		Ok: true,
	}, nil
}
