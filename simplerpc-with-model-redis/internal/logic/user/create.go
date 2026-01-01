package userlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"simplerpc-with-model-redis/internal/svc"
	"simplerpc-with-model-redis/internal/types/user"
)

type Create struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreate(ctx context.Context, svcCtx *svc.ServiceContext) *Create {
	return &Create{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Create) Create(in *user.CreateRequest) (*user.CreateResponse, error) {
	// todo: add your logic here and delete this line

	return &user.CreateResponse{}, nil
}
