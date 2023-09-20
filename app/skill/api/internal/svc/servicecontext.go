package svc

import (
	"github.com/superyyk/mayi_rpc/app/skill/api/internal/config"
	"github.com/superyyk/mayi_rpc/app/skill/rpc/skillclient"
	"github.com/superyyk/mayi_rpc/db"

	"github.com/Shopify/sarama"
	"github.com/jinzhu/gorm"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	KafkaClient   sarama.SyncProducer
	KafkaConsumer sarama.Consumer
	Db            *gorm.DB
	Rdb           *redis.Client
	SkillRpc      skillclient.Skill
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		KafkaClient:   db.KClient,
		KafkaConsumer: db.Consumer,
		Db:            db.Db,
		Rdb:           db.Rdb,
		SkillRpc:      skillclient.NewSkill(zrpc.MustNewClient(c.SkillRpc)),
	}
}
