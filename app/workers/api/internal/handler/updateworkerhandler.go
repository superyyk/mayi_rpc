package handler

import (
	"net/http"

	"github.com/superyyk/mayi_rpc/app/workers/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/workers/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/workers/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateworkerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateworkerLogic(r.Context(), svcCtx)
		token := r.Header.Get("token")

		resp, err := l.Updateworker(&req, token)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
