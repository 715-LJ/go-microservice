package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ManuscriptsModel = (*customManuscriptsModel)(nil)

type (
	// ManuscriptsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManuscriptsModel.
	ManuscriptsModel interface {
		manuscriptsModel
	}

	customManuscriptsModel struct {
		*defaultManuscriptsModel
	}
)

// NewManuscriptsModel returns a model for the database table.
func NewManuscriptsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ManuscriptsModel {
	return &customManuscriptsModel{
		defaultManuscriptsModel: newManuscriptsModel(conn, c, opts...),
	}
}
