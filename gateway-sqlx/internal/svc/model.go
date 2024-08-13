package svc

import (
	"gateway-sqlx/internal/model/users"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Model struct {
	User users.UsersModel
}

func NewModel(conn sqlx.SqlConn) Model {
	return Model{
		User: users.NewUsersModel(conn),
	}
}
