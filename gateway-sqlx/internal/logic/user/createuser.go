package userlogic

import (
	"context"
	"gateway-sqlx/internal/model/users"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"gateway-sqlx/internal/pb/userpb"
	"gateway-sqlx/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUser struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUser(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUser {
	return &CreateUser{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUser) CreateUser(in *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	// transaction
	var id int64
	err := l.svcCtx.SqlxConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		model := l.svcCtx.Model.User.WithSession(session)
		result, err := model.Insert(l.ctx, &users.Users{
			Username:   in.Username,
			Password:   in.Password,
			CreateTime: time.Now().Format(time.DateTime),
		})
		if err != nil {
			return err
		}
		id, err = result.LastInsertId()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// direct insert
	//result, err := l.svcCtx.Model.User.Insert(l.ctx, &users.Users{
	//	Username: in.Username,
	//	Password: in.Password,
	//  CreateTime: time.Now().Format(time.DateTime),
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//id, err := result.LastInsertId()
	//if err != nil {
	//	return nil, err
	//}

	return &userpb.CreateUserResponse{
		Id: int32(id),
	}, nil
}
