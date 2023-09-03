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

type GetLoginSmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoginSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginSmsLogic {
	return &GetLoginSmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginSmsLogic) GetLoginSms(req *types.LoginSmsReq) (resp *types.LoginSmsResp, err error) {
	// todo: add your logic here and delete this line
	var codes []model.Codes
	tx := l.svcCtx.Db.Begin()
	tx.Table("codes").Where("uid=? AND code=? AND tel=?", req.Uid, req.Code, req.Tel).First(&codes)
	if len(codes) > 0 { //匹配
		//判断超时
		is_time_out := common.IsOutTime(codes[0].Time, time.Now().Unix(), 5)
		if is_time_out {
			return &types.LoginSmsResp{
				Res:  "400",
				Msg:  "验证码已过期",
				Data: "400",
			}, nil
		}
		//获取用户信息
		var user []model.UserInfo
		var workerty []model.Workerty
		tx.Table("userdata").Where("tel=?", codes[0].Tel).First(&user)
		if len(user) > 0 { //找到用户注册信息
			//获取Token
			token, _ := l.svcCtx.AuthRpc.GetToken(l.ctx, &auth.GetTokenRequest{
				Userid: user[0].Uid,
			})
			l.svcCtx.Db.Table("worker_ty").Where("puber=? AND status=?", user[0].Uid, 0).Find(&workerty)
			var res map[string]interface{}
			res = make(map[string]interface{})
			res["token"] = token.Token
			res["user_info"] = user
			res["workerty"] = workerty
			//存入redis 登陆信息
			l.svcCtx.Rdb.Set(l.ctx, req.Tel, time.Now().Unix(), -1)
			return &types.LoginSmsResp{
				Res:  "200",
				Msg:  "success",
				Data: res,
			}, nil

		} else { //没有找到用户信息，新注册
			uid := tool.RandString(12)
			user_info := model.UserInfo{
				Tel:  req.Tel,
				Time: tool.Int64_string(time.Now().Unix()),
				Name: req.Tel,
				Uid:  uid,
				Head: "https://www.kissnet.cn:39900/static/images/logo.png",
			}
			//fmt.Println(user_info)
			if err := tx.Table("userdata").Create(&user_info).Error; err != nil {
				tx.Rollback()
				return &types.LoginSmsResp{
					Res:  "400",
					Msg:  "系统繁忙",
					Data: "400",
				}, nil
			}
			tx.Table("userdata").Where("tel=?", req.Tel).First(&user)
			//获取token
			token, _ := l.svcCtx.AuthRpc.GetToken(l.ctx, &auth.GetTokenRequest{
				Userid: user[0].Uid,
			})

			var res map[string]interface{}
			res = make(map[string]interface{})
			res["token"] = token.Token
			res["user_info"] = user
			l.svcCtx.Rdb.Set(l.ctx, req.Tel, time.Now().Unix(), -1)
			tx.Commit()
			return &types.LoginSmsResp{
				Res:  "200",
				Msg:  "登陆成功",
				Data: res,
			}, nil
		}

		// return &types.LoginSmsResp{
		// 	Res:  "200",
		// 	Msg:  "登陆成功",
		// 	Data: "400",
		// }, errors.New("验证码已过期")
	}

	//codes验证
	return &types.LoginSmsResp{
		Res:  "400",
		Msg:  "验证码不符",
		Data: "400",
	}, nil
}
