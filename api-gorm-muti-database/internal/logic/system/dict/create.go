package dict

import (
	"api-gorm-muti-database/internal/model/system_dict_key"
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
	model := system_dict_key.NewTSystemDictKeyModel(l.svcCtx.GormConn)

	err = model.Insert(l.ctx, nil, &system_dict_key.TSystemDictKey{
		Uuid:         "uuid",
		CategoryCode: req.CategoryCode,
		CategoryDesc: req.CategoryDesc,
		Sort:         int64(req.Sort),
	})

	return
}
