package helloworld

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HelloworldModel = (*customHelloworldModel)(nil)

type (
	// HelloworldModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHelloworldModel.
	HelloworldModel interface {
		helloworldModel
	}

	customHelloworldModel struct {
		*defaultHelloworldModel
	}
)

// NewHelloworldModel returns a model for the database table.
func NewHelloworldModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) HelloworldModel {
	return &customHelloworldModel{
		defaultHelloworldModel: newHelloworldModel(conn, op...),
	}
}
