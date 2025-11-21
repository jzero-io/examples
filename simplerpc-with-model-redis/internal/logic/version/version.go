package versionlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"simplerpc-with-model-redis/internal/svc"
	"simplerpc-with-model-redis/internal/types/version"
)

type Version struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVersion(ctx context.Context, svcCtx *svc.ServiceContext) *Version {
	return &Version{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Version) Version(in *version.VersionRequest) (*version.VersionResponse, error) {
	// todo: add your logic here and delete this line

	return &version.VersionResponse{}, nil
}
