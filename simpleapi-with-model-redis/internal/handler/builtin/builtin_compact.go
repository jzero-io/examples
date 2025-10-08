package builtin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"simpleapi-with-model-redis/internal/logic/builtin"
	"simpleapi-with-model-redis/internal/svc"
	types "simpleapi-with-model-redis/internal/types/builtin"
)

func Swagger(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := builtin.NewSwagger(r.Context(), svcCtx, r, w)
		err := l.Swagger()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

func Version(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VersionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := builtin.NewVersion(r.Context(), svcCtx, r)
		resp, err := l.Version(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
