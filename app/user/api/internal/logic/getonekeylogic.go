package logic

import (
	"context"
	"time"

	"github.com/superyyk/mayi_rpc/app/user/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/user/api/internal/types"
	"github.com/superyyk/mayi_rpc/auth/rpc/types/auth"
	"github.com/superyyk/mayi_rpc/model"
	"github.com/superyyk/mayi_rpc/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOneKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneKeyLogic {
	return &GetOneKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOneKeyLogic) GetOneKey(req *types.OneKeyReq) (resp *types.OneKeyResp, err error) {
	// todo: add your logic here and delete this line
	uid := tool.RandString(12)
	tx := l.svcCtx.Db.Begin()
	user_info := &model.UserInfo{
		Tel:    req.Tel,
		Time:   tool.Int64_string(time.Now().Unix()),
		Name:   req.Tel,
		Uid:    uid,
		Head:   "https://www.kissnet.cn:39900/static/images/logo.png",
		Openid: req.Openid,
	}
	var user []model.UserInfo
	var workerty []model.Workerty
	tx.Table("userdata").Where("tel=?", req.Tel).First(&user)
	//获取token
	token, _ := l.svcCtx.AuthRpc.GetToken(l.ctx, &auth.GetTokenRequest{
		Userid: user[0].Uid,
	})
	l.svcCtx.Db.Table("worker_ty").Where("puber=? AND status=?", user[0].Uid, 0).Find(&workerty)
	var res map[string]interface{}
	res = make(map[string]interface{})
	if len(user) > 0 { //找到用户

		res["token"] = token.Token
		res["user_info"] = user
		res["workerty"] = workerty
		l.svcCtx.Rdb.Set(l.ctx, req.Tel, time.Now().Unix(), -1)
		return &types.OneKeyResp{
			Res:  "200",
			Msg:  "登陆成功",
			Data: res,
		}, nil

	} else {
		if err := tx.Table("userdata").Create(user_info).Error; err != nil {
			tx.Rollback()
			return &types.OneKeyResp{
				Res:  "400",
				Msg:  "登陆失败",
				Data: "",
			}, nil
		} else {
			tx.Table("userdata").Where("tel=?", req.Tel).First(&user)
		}

		res["token"] = token.Token
		res["user_info"] = user
		l.svcCtx.Rdb.Set(l.ctx, req.Tel, time.Now().Unix(), 0)
		tx.Commit()
		return &types.OneKeyResp{
			Res:  "200",
			Msg:  "登陆成功",
			Data: res,
		}, nil

		// return &types.OneKeyResp{
		// 	Res:  "400",
		// 	Msg:  "未找到该用户信息",
		// 	Data: user_info,
		// }, nil
	}

}
