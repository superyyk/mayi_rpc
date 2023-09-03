package logic

import (
	"context"
	"time"

	"mayi/app/user/api/internal/svc"
	"mayi/app/user/api/internal/types"
	"mayi/auth/rpc/types/auth"
	"mayi/common"
	"mayi/model"
	"mayi/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginLogic {
	return &GetLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginLogic) GetLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	//判断验证码过期
	var user []model.UserInfo
	var workerty []model.Workerty
	l.svcCtx.Db.Table("userdata").Where("tel=? AND pass=?", req.Tel, tool.Md5(req.Pass)).Find(&user)

	if len(user) > 0 {
		res := make(map[string]interface{})
		//获取token
		token, _ := l.svcCtx.AuthRpc.GetToken(l.ctx, &auth.GetTokenRequest{
			Userid: user[0].Uid,
		})
		l.svcCtx.Db.Table("worker_ty").Where("puber=? AND status=?", user[0].Uid, 0).Find(&workerty)
		res["token"] = token.Token
		res["user_info"] = user
		res["workerty"] = workerty
		l.svcCtx.Rdb.Set(l.ctx, req.Tel, time.Now().Unix(), -1)
		return &types.LoginResp{
			Res:  "200",
			Msg:  common.GetErrMsg(200),
			Data: res,
		}, nil
	}
	return &types.LoginResp{
		Res:  "400",
		Msg:  common.GetErrMsg(410),
		Data: user,
	}, nil
}
