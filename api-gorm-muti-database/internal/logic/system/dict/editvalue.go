package dict

import (
	"context"

	"api-gorm-muti-database/internal/svc"
	"api-gorm-muti-database/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditValue struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditValue(ctx context.Context, svcCtx *svc.ServiceContext) *EditValue {
	return &EditValue{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditValue) EditValue(req *types.EditDictValueRequest) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
