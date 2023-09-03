package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mayi/app/product/api/internal/logic"
	"mayi/app/product/api/internal/svc"
	"mayi/app/product/api/internal/types"
)

func AddYuanliaoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.YuanliaoAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddYuanliaoLogic(r.Context(), svcCtx)
		resp, err := l.AddYuanliao(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
