package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	SmsRpc  zrpc.RpcClientConf
	UserRpc zrpc.RpcClientConf
	AuthRpc zrpc.RpcClientConf
}
