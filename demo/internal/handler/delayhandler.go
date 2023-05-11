package handler

import (
	"net/http"

	"docker-go-zero/demo/internal/logic"
	"docker-go-zero/demo/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDelayLogic(r.Context(), svcCtx)
		resp, err := l.Delay()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
