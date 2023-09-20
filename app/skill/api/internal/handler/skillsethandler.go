package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/superyyk/mayi_rpc/app/skill/api/internal/logic"
	"github.com/superyyk/mayi_rpc/app/skill/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/skill/api/internal/types"
)

func SkillSetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSkillSetLogic(r.Context(), svcCtx)
		resp, err := l.SkillSet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
