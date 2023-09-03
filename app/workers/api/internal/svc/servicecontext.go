package svc

import (
	"mayi/app/workers/api/internal/config"
	"mayi/auth/rpc/token"
	"mayi/db"

	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Db      *gorm.DB
	AuthRpc token.Token
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Db:      db.Db,
		AuthRpc: token.NewToken(zrpc.MustNewClient(c.AuthRpc)),
	}
}
