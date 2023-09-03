package svc

import (
	"mayi/app/product/api/internal/config"
	"mayi/db"

	"mayi/model/model"

	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	Db        *gorm.DB
	UserModel model.UserdataModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Db:        db.Db,
		UserModel: model.NewUserdataModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
