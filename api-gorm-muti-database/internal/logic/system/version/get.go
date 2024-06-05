package version

import (
	"context"

	"api-gorm-muti-database/internal/svc"
	"api-gorm-muti-database/internal/types"

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

func (l *Get) Get(req *types.Empty) (resp *types.Version, err error) {
	// todo: add your logic here and delete this line

	return
}
