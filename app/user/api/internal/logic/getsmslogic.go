package logic

import (
	"context"
	"errors"

	"github.com/superyyk/mayi_rpc/app/user/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/user/api/internal/types"
	"github.com/superyyk/mayi_rpc/sms/rpc/types/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsLogic {
	return &GetSmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSmsLogic) GetSms(req *types.SmsReq) (resp *types.SmsResp, err error) {
	// todo: add your logic here and delete this line
	code_res, err := l.svcCtx.SmsRpc.GetCode(l.ctx, &sms.GetCodeRequest{
		Tel: req.Tel,
	})
	if err != nil {
		return &types.SmsResp{
			Res:  "400",
			Msg:  "Rpc调用失败",
			Data: 400,
		}, errors.New("发送失败")
	}
	//l.svcCtx.Rdb.Set(l.ctx, req.Tel, time.Now().Unix(), 0)
	return &types.SmsResp{
		Res:  "200",
		Msg:  "发送成功",
		Data: code_res,
	}, nil
}
