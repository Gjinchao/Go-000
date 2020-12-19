//+build wireinject
//go:generate wire

package main

import (
	"github.com/Gjinchao/Go-000/Week04/api"
	"github.com/Gjinchao/Go-000/Week04/internal/biz"
	"github.com/Gjinchao/Go-000/Week04/internal/data"
	"github.com/Gjinchao/Go-000/Week04/internal/service"
	"github.com/google/wire"
)

func InitializeAllInstance() api.AccountInterface {
	wire.Build(api.NewAccountInterface, service.NewAccountService, biz.NewAccountBiz, data.NewAccountRepo)
	return api.AccountInterface{}
}
