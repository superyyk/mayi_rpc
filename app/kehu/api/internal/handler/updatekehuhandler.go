package handler

import (
	"net/http"

	"mayi/app/kehu/api/internal/logic"
	"mayi/app/kehu/api/internal/svc"
	"mayi/app/kehu/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateKehuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateKehuLogic(r.Context(), svcCtx)
		token := r.Header.Get("token")
		resp, err := l.UpdateKehu(&req, token)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
