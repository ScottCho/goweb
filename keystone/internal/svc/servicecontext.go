package svc

import (
	"github.com/ScottCho/goweb/keystone/internal/config"
	"github.com/ScottCho/goweb/keystone/models/users"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel users.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.Datasource)

	return &ServiceContext{
		Config:    c,
		UserModel: users.NewUsersModel(sqlConn),
	}
}
