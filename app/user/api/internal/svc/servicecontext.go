package svc

import (
	"github.com/superyyk/mayi_rpc/app/user/api/internal/config"
	"github.com/superyyk/mayi_rpc/app/user/rpc/userclient"
	"github.com/superyyk/mayi_rpc/auth/rpc/token"
	"github.com/superyyk/mayi_rpc/db"
	"github.com/superyyk/mayi_rpc/sms/rpc/smscode"

	"github.com/redis/go-redis/v9"
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Db      *gorm.DB      //mysql
	Rdb     *redis.Client //redis
	SmsRpc  smscode.SmsCode
	UserRpc userclient.User
	AuthRpc token.Token
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Db:      db.Db,
		Rdb:     db.Rdb,
		SmsRpc:  smscode.NewSmsCode(zrpc.MustNewClient(c.SmsRpc)),
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AuthRpc: token.NewToken(zrpc.MustNewClient(c.AuthRpc)),
	}
}
