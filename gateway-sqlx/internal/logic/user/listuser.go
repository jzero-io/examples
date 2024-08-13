package userlogic

import (
	"context"
	"github.com/jzero-io/jzero-contrib/condition"

	"gateway-sqlx/internal/pb/userpb"
	"gateway-sqlx/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUser struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUser(ctx context.Context, svcCtx *svc.ServiceContext) *ListUser {
	return &ListUser{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUser) ListUser(in *userpb.ListUserRequest) (*userpb.ListUserResponse, error) {
	page, total, err := l.svcCtx.Model.User.Page(l.ctx, condition.New(condition.Condition{
		Operator: "LIMIT",
		Value:    in.Size,
	}, condition.Condition{
		Operator: "OFFSET",
		Value:    (in.Page - 1) * in.Size,
	}, condition.Condition{
		Operator: "ORDER BY",
		Value:    "create_time DESC",
	}, condition.Condition{
		Skip:     in.GetUsername() == "",
		Field:    "username",
		Operator: "LIKE",
		Value:    "%" + in.GetUsername() + "%",
	})...)
	if err != nil {
		return nil, err
	}

	var list []*userpb.UserItem

	for _, item := range page {
		list = append(list, &userpb.UserItem{
			Username:   item.Username,
			CreateTime: item.CreateTime,
		})
	}

	return &userpb.ListUserResponse{
		Total: int32(total),
		List:  list,
	}, nil
}
