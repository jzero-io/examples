package dict

import (
	"context"

	"api-gorm-muti-database/internal/svc"
	"api-gorm-muti-database/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListValue struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListValue(ctx context.Context, svcCtx *svc.ServiceContext) *ListValue {
	return &ListValue{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListValue) ListValue(req *types.ListDictValueRequest) (resp *types.ListDictValueResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
