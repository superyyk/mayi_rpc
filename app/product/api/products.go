package main

import (
	"flag"
	"fmt"

	"mayi/app/product/api/internal/config"
	"mayi/app/product/api/internal/handler"
	"mayi/app/product/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/products.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	c.KeyFile = "../../../ssl/ssl.key"
	c.CertFile = "../../../ssl/ssl.pem"
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
