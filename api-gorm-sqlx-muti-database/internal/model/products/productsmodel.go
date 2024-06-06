package products

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProductsModel = (*customProductsModel)(nil)

type (
	// ProductsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductsModel.
	ProductsModel interface {
		productsModel
		withSession(session sqlx.Session) ProductsModel
	}

	customProductsModel struct {
		*defaultProductsModel
	}
)

// NewProductsModel returns a model for the database table.
func NewProductsModel(conn sqlx.SqlConn) ProductsModel {
	return &customProductsModel{
		defaultProductsModel: newProductsModel(conn),
	}
}

func (m *customProductsModel) withSession(session sqlx.Session) ProductsModel {
	return NewProductsModel(sqlx.NewSqlConnFromSession(session))
}
