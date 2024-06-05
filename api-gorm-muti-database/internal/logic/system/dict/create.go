package dict

import (
	"context"

	"api-gorm-muti-database/internal/svc"
	"api-gorm-muti-database/internal/types"

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

func (l *Create) Create(req *types.CreateDictRequest) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
