package logic

import (
	"context"

	"mayi/app/user/api/internal/svc"
	"mayi/app/user/api/internal/types"
	"mayi/common"
	"mayi/model"
	"mayi/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChangePassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChangePassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChangePassLogic {
	return &GetChangePassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChangePassLogic) GetChangePass(req *types.ChangePassReq) (resp *types.ChangePassResp, err error) {
	// todo: add your logic here and delete this line
	pass := tool.Md5(req.Pass)
	var codes []model.Codes
	l.svcCtx.Db.Table("codes").Where("tel=? AND code=? AND uid=?", req.Tel, req.Code, req.Uid).First(&codes)
	if len(codes) > 0 {
		row := l.svcCtx.Db.Table("userdata").Where("tel=?", req.Tel).Update("pass", pass)
		if row.RowsAffected == 1 {
			return &types.ChangePassResp{
				Res: "200",
				Msg: common.GetErrMsg(200),
			}, nil
		} else {
			return &types.ChangePassResp{
				Res:  "400",
				Msg:  "修改密码失败,或该手机号码未注册用户！",
				Data: "400",
			}, nil
		}

	}

	return &types.ChangePassResp{
		Res:  "400",
		Msg:  common.GetErrMsg(409),
		Data: "400",
	}, nil
}
