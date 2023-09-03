package logic

import (
	"context"
	"fmt"

	"mayi/app/skill/api/internal/svc"
	"mayi/app/skill/api/internal/types"

	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/logx"
)

type PubLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPubLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PubLogic {
	return &PubLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PubLogic) Pub(req *types.PubReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	// 构造一个消息
	topic := req.Topic
	uid := req.Uid
    
	msg := &sarama.ProducerMessage{}
	msg.Topic = uid + topic
	msg.Value = sarama.StringEncoder("this is a test log")
	//defer client.Close()
	// 发送消息
	pid, offset, err := l.svcCtx.KafkaClient.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	res := make(map[string]interface{})
	res["pid"] = pid
	res["offset"] = offset
	// fmt.Printf("pid:%v offset:%v\n", pid, offset)
	return &types.Resp{
		Res:  "200",
		Msg:  "success",
		Data: res,
	}, nil
}
