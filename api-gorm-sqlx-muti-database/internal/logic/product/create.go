package product

import (
	"api-gorm-sqlx-muti-database/internal/model/products"
	"context"
	"database/sql"

	"api-gorm-sqlx-muti-database/internal/svc"
	"api-gorm-sqlx-muti-database/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Create struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreate(ctx context.Context, svcCtx *svc.ServiceContext) *Create {
	return &Create{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Create) Create(req *types.CreateProductRequest) (resp *types.Empty, err error) {
	data := products.GormProduct{Code: req.Code, Price: uint(req.Price), Remark: req.Remark}
	err = l.svcCtx.GormConn.Create(&data).Error
	if err != nil {
		return nil, err
	}

	// use sqlx
	model := products.NewProductsModel(l.svcCtx.SqlxConn)
	_, err = model.Insert(l.ctx, &products.Products{
		Code: sql.NullString{
			Valid:  true,
			String: req.Code,
		},
		Price: sql.NullInt64{
			Valid: true,
			Int64: int64(data.Price),
		},
		Remark: sql.NullString{
			Valid:  true,
			String: req.Remark,
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
