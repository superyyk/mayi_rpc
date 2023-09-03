package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mayi/app/workers/api/internal/logic"
	"mayi/app/workers/api/internal/svc"
	"mayi/app/workers/api/internal/types"
)

func SearchWorkertyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchTyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSearchWorkertyLogic(r.Context(), svcCtx)
		resp, err := l.SearchWorkerty(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
