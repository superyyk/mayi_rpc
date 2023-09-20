package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/superyyk/mayi_rpc/app/kehu/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/kehu/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/kehu/api/internal/types"
)

func SearchKehuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSearchKehuLogic(r.Context(), svcCtx)
		resp, err := l.SearchKehu(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
