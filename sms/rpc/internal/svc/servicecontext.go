package svc

import (
	"github.com/superyyk/mayi_rpc/db"
	"github.com/superyyk/mayi_rpc/sms/rpc/internal/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
