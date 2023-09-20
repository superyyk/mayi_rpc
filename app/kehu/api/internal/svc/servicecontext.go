package svc

import (
	"github.com/superyyk/mayi_rpc/app/kehu/api/internal/config"
	"github.com/superyyk/mayi_rpc/auth/rpc/token"
	"github.com/superyyk/mayi_rpc/db"

	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	//KehuModel kehu_model.KehuModel
	Db      *gorm.DB
	AuthRpc token.Token
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//KehuModel: kehu_model.NewKehuModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
		Db:      db.Db,
		AuthRpc: token.NewToken(zrpc.MustNewClient(c.AuthRpc)),
	}
}
