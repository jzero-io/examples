package product

import (
	"net/http"

	"api-gorm-sqlx-muti-database/internal/logic/product"
	"api-gorm-sqlx-muti-database/internal/svc"
	"api-gorm-sqlx-muti-database/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Get(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IntIdPathRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewGet(r.Context(), svcCtx)
		resp, err := l.Get(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
