package dict

import (
	"context"

	"api-gorm-muti-database/internal/svc"
	"api-gorm-muti-database/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListKey struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListKey(ctx context.Context, svcCtx *svc.ServiceContext) *ListKey {
	return &ListKey{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListKey) ListKey(req *types.ListDictKeyRequest) (resp *types.ListDictKeyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
