package svc

import (
	"mayi/app/kehu/api/internal/config"
	"mayi/auth/rpc/token"
	"mayi/db"

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
