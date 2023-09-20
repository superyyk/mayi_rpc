package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/superyyk/mayi_rpc/app/jizhang/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/jizhang/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/jizhang/api/internal/types"
)

func jizhangAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JizhangAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewJizhangAddLogic(r.Context(), svcCtx)
		resp, err := l.JizhangAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
