package version

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"simpleapi-with-model-redis/internal/svc"
	types "simpleapi-with-model-redis/internal/types/version"
)

type Version struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewVersion(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Version {
	return &Version{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *Version) Version(req *types.VersionRequest) (resp *types.VersionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
