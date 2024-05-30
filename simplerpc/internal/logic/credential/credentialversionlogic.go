package credentiallogic

import (
	"context"

	"simplerpc/internal/pb/credentialpb"
	"simplerpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CredentialVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCredentialVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CredentialVersionLogic {
	return &CredentialVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CredentialVersionLogic) CredentialVersion(in *credentialpb.Empty) (*credentialpb.CredentialVersionResponse, error) {
	// todo: add your logic here and delete this line

	return &credentialpb.CredentialVersionResponse{}, nil
}
