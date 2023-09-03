package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mayi/app/jizhang/api/internal/logic"
	"mayi/app/jizhang/api/internal/svc"
	"mayi/app/jizhang/api/internal/types"
)

func jizhangLeixingAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JizhangLeixingAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewJizhangLeixingAddLogic(r.Context(), svcCtx)
		resp, err := l.JizhangLeixingAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
