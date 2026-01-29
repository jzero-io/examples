package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"simpleapi/internal/logic/user"
	"simpleapi/internal/svc"
	types "simpleapi/internal/types/user"
)

func Create(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewCreate(r.Context(), svcCtx, r)
		resp, err := l.Create(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
