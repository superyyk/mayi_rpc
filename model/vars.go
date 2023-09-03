package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

type TestCity struct {
	Name string
	Lat  string
	Lng  string
}
