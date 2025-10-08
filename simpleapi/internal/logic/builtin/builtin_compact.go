package builtin

import (
	"context"
	"net/http"
	"os"
	"runtime"

	"github.com/jzero-io/jzero/core/swaggerv2"
	"github.com/zeromicro/go-zero/core/logx"

	"simpleapi/internal/svc"
	types "simpleapi/internal/types/builtin"
)

type Swagger struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
	w      http.ResponseWriter
}

func NewSwagger(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request, w http.ResponseWriter) *Swagger {
	return &Swagger{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
		w:      w,
	}
}

func (l *Swagger) Swagger() error {
	opts := new(swaggerv2.Swaggerv2Opts).DefaultOptions()
	swaggerv2.SwaggerHandler(opts, l.w, l.r)
	return nil
}

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
	return &types.VersionResponse{
		Version:   os.Getenv("VERSION"),
		GoVersion: runtime.Version(),
		Commit:    os.Getenv("COMMIT"),
		Date:      os.Getenv("DATE"),
	}, nil
}
