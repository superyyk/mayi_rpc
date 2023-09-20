package handler

import (
	"net/http"

	"github.com/superyyk/mayi_rpc/app/user/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/user/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoadUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		UserUuid := r.Header.Get("token")
		l := logic.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo(&req, UserUuid)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
