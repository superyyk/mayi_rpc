package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/superyyk/mayi_rpc/app/skill/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/skill/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/skill/api/internal/types"
)

func SkillGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSkillGetLogic(r.Context(), svcCtx)
		resp, err := l.SkillGet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
