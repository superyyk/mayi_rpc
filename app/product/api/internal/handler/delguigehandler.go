package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/superyyk/mayi_rpc/app/product/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/product/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/product/api/internal/types"
)

func DelGuigeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GuigeDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDelGuigeLogic(r.Context(), svcCtx)
		resp, err := l.DelGuige(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
