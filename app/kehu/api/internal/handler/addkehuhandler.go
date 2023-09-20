package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/superyyk/mayi_rpc/app/kehu/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/kehu/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/kehu/api/internal/types"
)

func AddKehuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddKehuLogic(r.Context(), svcCtx)
		resp, err := l.AddKehu(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
