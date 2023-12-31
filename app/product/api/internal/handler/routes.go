// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/superyyk/mayi_rpc/app/product/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/products/add",
				Handler: AddProductHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/products/update",
				Handler: UpdateProductHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/products/search",
				Handler: SearchProductHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/products/guige/add",
				Handler: AddGuigeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/products/guige/del",
				Handler: DelGuigeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/products/yuanliao/add",
				Handler: AddYuanliaoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/products/yuanliao/del",
				Handler: DelYuanliaoHandler(serverCtx),
			},
		},
	)
}
