package svc

import (
	"mayi/app/skill/rpc/internal/config"
	"mayi/db"

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
