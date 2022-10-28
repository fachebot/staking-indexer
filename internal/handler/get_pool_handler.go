package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"staking-indexer/internal/logic"
	"staking-indexer/internal/svc"
	"staking-indexer/internal/types"
)

func GetPoolHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPoolRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetPoolLogic(r.Context(), svcCtx)
		resp, err := l.GetPool(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
