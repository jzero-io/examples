package product

import (
	"api-gorm-sqlx-muti-database/internal/model/products"
	"api-gorm-sqlx-muti-database/internal/svc"
	"api-gorm-sqlx-muti-database/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGet(ctx context.Context, svcCtx *svc.ServiceContext) *Get {
	return &Get{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get) Get(req *types.IntIdPathRequest) (resp *types.GetProductResponse, err error) {
	//var data products.GormProduct

	//err = l.svcCtx.GormConn.Where("id = ?", req.Id).First(&data).Error
	//if err != nil {
	//	return nil, err
	//}
	//
	//resp = &types.GetProductResponse{
	//	Id:     int(data.ID),
	//	Code:   data.Code,
	//	Price:  int(data.Price),
	//	Remark: data.Remark,
	//}

	// use logx
	model := products.NewProductsModel(l.svcCtx.SqlxConn)
	one, err := model.FindOne(l.ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}
	resp = &types.GetProductResponse{
		Id:     int(one.Id),
		Code:   one.Code.String,
		Price:  int(one.Price.Int64),
		Remark: one.Remark.String,
	}

	return
}
