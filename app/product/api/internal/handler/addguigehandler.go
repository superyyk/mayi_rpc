package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/superyyk/mayi_rpc/app/product/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/product/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/product/api/internal/types"
)

func AddGuigeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GuigeAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddGuigeLogic(r.Context(), svcCtx)
		resp, err := l.AddGuige(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
