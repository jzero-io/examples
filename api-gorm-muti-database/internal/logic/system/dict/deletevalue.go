package dict

import (
	"context"

	"api-gorm-muti-database/internal/svc"
	"api-gorm-muti-database/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteValue struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteValue(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteValue {
	return &DeleteValue{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteValue) DeleteValue(req *types.DeleteDictValueRequest) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
