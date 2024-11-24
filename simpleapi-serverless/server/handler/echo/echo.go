package echo

import (
	"net/http"
	"simpleapi-serverless/server/logic/echo"
	"simpleapi-serverless/server/svc"
	types "simpleapi-serverless/server/types/echo"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Echo(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EchoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := echo.NewEcho(r.Context(), svcCtx)
		resp, err := l.Echo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
