package dict

import (
	"api-gorm-muti-database/internal/model/system_dict_key"
	"context"
	"fmt"

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

	var dictKey system_dict_key.TSystemDictKey
	l.svcCtx.GormConn.First(&dictKey, 1)

	fmt.Println(dictKey)
	return
}
