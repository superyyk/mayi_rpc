package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/superyyk/mayi_rpc/app/liucheng/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/liucheng/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/liucheng/api/internal/types"
)

func stepUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StepUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewStepUpdateLogic(r.Context(), svcCtx)
		resp, err := l.StepUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
