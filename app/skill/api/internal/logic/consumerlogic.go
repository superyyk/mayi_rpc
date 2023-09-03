package logic

import (
	"context"
	"fmt"

	"mayi/app/skill/api/internal/svc"
	"mayi/app/skill/api/internal/types"

	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/logx"
)

type ConsumerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConsumerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConsumerLogic {
	return &ConsumerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConsumerLogic) Consumer(req *types.ConsumerReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	topic := req.Topic
	consumer := l.svcCtx.KafkaConsumer
	res := make(map[string]interface{})
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return nil, nil
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return nil, nil
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				logx.Infof("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
				res[string(msg.Key)] = msg.Value

			}
		}(pc)
	}

	return &types.Resp{
		Res:  "200",
		Msg:  "success",
		Data: res,
	}, nil
}
