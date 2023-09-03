package logic

import (
	"context"
	"time"

	"mayi/app/skill/rpc/internal/svc"
	"mayi/app/skill/rpc/types/skill"
	"mayi/model"
	"mayi/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSkillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSkillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSkillLogic {
	return &GetSkillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSkillLogic) GetSkill(in *skill.SkillRequest) (*skill.Rsp, error) {
	// todo: add your logic here and delete this line
	total, _ := l.svcCtx.Rdb.LLen(l.ctx, "skill").Result()
	orderid := tool.GetOrderUuid(time.Now())
	if total > 0 {
		goods_id := l.svcCtx.Rdb.RPop(l.ctx, "skill").String()

		sk := &model.Skill{
			Name:    tool.Int64_string(total),
			Uid:     goods_id,
			UserId:  in.Uid,
			OrderId: orderid,
			Time:    time.Now().Unix(),
		}
		if err := l.svcCtx.Db.Table("skill").Create(&sk).Error; err != nil {
			return &skill.Rsp{
				Orderid: "",
			}, nil
		}

	}
	return &skill.Rsp{
		Orderid: orderid,
	}, nil
}
