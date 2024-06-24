package products

import (
	"context"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductsModel = (*customProductsModel)(nil)

type (
	// ProductsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductsModel.
	ProductsModel interface {
		productsModel
		withSession(session sqlx.Session) ProductsModel

		FindById(ctx context.Context, id int64) (*Products, error)
	}

	customProductsModel struct {
		*defaultProductsModel
	}
)

func (m *customProductsModel) FindById(ctx context.Context, id int64) (*Products, error) {
	sb := sqlbuilder.NewSelectBuilder()

	sb = sb.Select("*").From(m.table)
	sb.Where(sb.E("id", id))

	sql, args := sb.Build()
	fmt.Println(sql, args)
	return nil, nil
}

// NewProductsModel returns a model for the database table.
func NewProductsModel(conn sqlx.SqlConn) ProductsModel {
	return &customProductsModel{
		defaultProductsModel: newProductsModel(conn),
	}
}

func (m *customProductsModel) withSession(session sqlx.Session) ProductsModel {
	return NewProductsModel(sqlx.NewSqlConnFromSession(session))
}
