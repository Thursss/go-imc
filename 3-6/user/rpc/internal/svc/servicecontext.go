package svc

import (
	"user/rpc/internal/config"
	"user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
