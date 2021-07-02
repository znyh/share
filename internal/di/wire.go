// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/znyh/share/internal/dao"
	"github.com/znyh/share/internal/server/grpc"
	"github.com/znyh/share/internal/server/http"
	"github.com/znyh/share/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
