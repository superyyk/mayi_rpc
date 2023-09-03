package logic

import (
	"context"

	"mayi/app/kehu/api/internal/svc"
	"mayi/app/kehu/api/internal/types"
	"mayi/auth/rpc/types/auth"
	"mayi/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateKehuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateKehuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateKehuLogic {
	return &UpdateKehuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateKehuLogic) UpdateKehu(req *types.UpdateReq, token string) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	kehu := &model.Kehu{
		Name:   req.Name,
		Age:    req.Age,
		Level:  req.Level,
		Tel:    req.Tel,
		Status: req.Status,
		Tip:    req.Tip,
		Uid:    req.Uid,
	}
	claims, _ := l.svcCtx.AuthRpc.VerifyToken(l.ctx, &auth.TokenRequest{
		Token: token,
	})

	row := l.svcCtx.Db.Table("kehu").Where("uid=? AND puber=?", req.Uid, claims.Userid).Updates(map[string]interface{}{"name": req.Name, "age": req.Age, "status": req.Status, "tel": req.Tel, "tip": req.Tip, "level": req.Level})
	if row.RowsAffected == 1 {
		return &types.Resp{
			Res: "200",
			Msg: "操作成功",
		}, nil
	}
	return &types.Resp{
		Res:  "400",
		Msg:  "操作失败",
		Data: kehu,
	}, nil
}
