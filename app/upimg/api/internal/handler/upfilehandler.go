package handler

import (
	"net/http"

	"github.com/superyyk/mayi_rpc/app/upimg/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/upimg/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/upimg/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpFileLogic(r.Context(), svcCtx)
		resp, err := l.UpFile(&req, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
