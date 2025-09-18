package versionlogic

import (
	"context"
	"os"
	"runtime"

	"github.com/zeromicro/go-zero/core/logx"

	"simplerpc/internal/pb/versionpb"
	"simplerpc/internal/svc"
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

func (l *Version) Version(in *versionpb.VersionRequest) (*versionpb.VersionResponse, error) {
	return &versionpb.VersionResponse{
		Version:   os.Getenv("VERSION"),
		GoVersion: runtime.Version(),
		Commit:    os.Getenv("COMMIT"),
		Date:      os.Getenv("DATE"),
	}, nil
}
