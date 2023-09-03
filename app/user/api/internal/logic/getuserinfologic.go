package logic

import (
	"context"

	"mayi/app/user/api/internal/svc"
	"mayi/app/user/api/internal/types"
	"mayi/auth/rpc/types/auth"
	"mayi/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.LoadUserReq, useruuid string) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	var user []model.UserInfo
	var workerty []model.Workerty
	res := make(map[string]interface{})
	rpcInfo, err := l.svcCtx.AuthRpc.VerifyToken(l.ctx, &auth.TokenRequest{
		Token: useruuid,
	})
	if err != nil {
		return &types.Resp{
			Res:  "401",
			Msg:  "Token失效",
			Data: user,
		}, nil
	}
	if req.UserUid != rpcInfo.Userid {
		return &types.Resp{
			Res:  "401",
			Msg:  "Token无效",
			Data: user,
		}, nil
	}
	l.svcCtx.Db.Table("userdata").Where("uid=?", rpcInfo.Userid).First(&user)
	l.svcCtx.Db.Table("worker_ty").Where("puber=? AND status=?", user[0].Uid, 0).Find(&workerty)
	res["workerty"] = workerty
	res["user_info"] = user
	if len(user) > 0 {
		return &types.Resp{
			Res:  "200",
			Msg:  "success",
			Data: res,
		}, nil
	}
	return &types.Resp{
		Res:  "401",
		Msg:  "未找到该用户信息",
		Data: user,
	}, nil
}
