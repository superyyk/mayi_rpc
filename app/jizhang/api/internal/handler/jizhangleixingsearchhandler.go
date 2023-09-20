package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/superyyk/mayi_rpc/app/jizhang/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/jizhang/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/jizhang/api/internal/types"
)

func jizhangLeixingSearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JizhangLeixingSearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewJizhangLeixingSearchLogic(r.Context(), svcCtx)
		resp, err := l.JizhangLeixingSearch(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
