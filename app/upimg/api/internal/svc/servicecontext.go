package svc

import (
	"github.com/superyyk/mayi_rpc/app/upimg/api/internal/config"
	"github.com/superyyk/mayi_rpc/db"

	"github.com/jinzhu/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Db:     db.Db,
	}
}
