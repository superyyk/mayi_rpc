package main

import (
	"flag"
	"fmt"

	"github.com/superyyk/mayi_rpc/app/skill/api/internal/config"
	"github.com/superyyk/mayi_rpc/app/skill/api/internal/handler"
	"github.com/superyyk/mayi_rpc/app/skill/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/skill.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
