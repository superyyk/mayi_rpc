package svc

import (
	"github.com/superyyk/mayi_rpc/app/skill/rpc/internal/config"
	"github.com/superyyk/mayi_rpc/db"

	"github.com/jinzhu/gorm"
	"github.com/redis/go-redis/v9"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Rdb    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Db:     db.Db,
		Rdb:    db.Rdb,
	}
}
