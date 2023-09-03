package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserdataModel = (*customUserdataModel)(nil)

type (
	// UserdataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserdataModel.
	UserdataModel interface {
		userdataModel
	}

	customUserdataModel struct {
		*defaultUserdataModel
	}
)

// NewUserdataModel returns a model for the database table.
func NewUserdataModel(conn sqlx.SqlConn, c cache.CacheConf) UserdataModel {
	return &customUserdataModel{
		defaultUserdataModel: newUserdataModel(conn, c),
	}
}
