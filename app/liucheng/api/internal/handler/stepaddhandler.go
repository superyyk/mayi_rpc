package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mayi/app/liucheng/api/internal/logic"
	"mayi/app/liucheng/api/internal/svc"
	"mayi/app/liucheng/api/internal/types"
)

func stepAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StepAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewStepAddLogic(r.Context(), svcCtx)
		resp, err := l.StepAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
