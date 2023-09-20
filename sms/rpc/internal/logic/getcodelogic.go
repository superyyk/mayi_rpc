package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/superyyk/mayi_rpc/common"
	"github.com/superyyk/mayi_rpc/model"
	"github.com/superyyk/mayi_rpc/sms/rpc/internal/svc"
	"github.com/superyyk/mayi_rpc/sms/rpc/types/sms"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeLogic {
	return &GetCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCodeLogic) GetCode(in *sms.GetCodeRequest) (*sms.GetResponse, error) {

	name := "小工单"
	tel := in.Tel
	send_url := "http://www.kissnet.cn:39005/dianliang/index/send_rpc_sms?tel=" + tel + "&sign=" + name
	re, err := common.HttpGet(send_url)
	if err != nil {
		return &sms.GetResponse{}, errors.New("http_send error")
	}
	var h *model.Hp
	err = json.Unmarshal([]byte(re), &h)
	if h.Res != 200 {
		return &sms.GetResponse{}, errors.New("res error")
	}
	codes := model.Codes{
		Uid:  h.Uid,
		Code: h.Code,
		Tel:  tel,
		Time: time.Now().Unix(),
	}

	err = l.svcCtx.Db.Table("codes").Create(&codes).Error
	if err != nil {
		return &sms.GetResponse{}, errors.New("存库失败")
	}

	return &sms.GetResponse{
		Uid:  h.Uid,
		Code: "",
	}, nil

}
