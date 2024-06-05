package dict

import (
	"net/http"

	"api-gorm-muti-database/internal/logic/system/dict"
	"api-gorm-muti-database/internal/svc"
	"api-gorm-muti-database/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteValue(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteDictValueRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dict.NewDeleteValue(r.Context(), svcCtx)
		resp, err := l.DeleteValue(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
