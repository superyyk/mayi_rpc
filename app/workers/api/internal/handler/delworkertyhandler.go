package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/superyyk/mayi_rpc/app/workers/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/workers/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/workers/api/internal/types"
)

func DelWorkertyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelTyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDelWorkertyLogic(r.Context(), svcCtx)
		resp, err := l.DelWorkerty(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
