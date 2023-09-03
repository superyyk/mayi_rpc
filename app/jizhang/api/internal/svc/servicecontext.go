package svc

import (
	"mayi/app/jizhang/api/internal/config"
	"mayi/db"

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
