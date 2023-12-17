//go:build wireinject

package main

import (
	"cchart/internal/http"
	"cchart/internal/kernel"
	"cchart/internal/provider"
	"github.com/google/wire"
)

func Inject() *kernel.Application {
	wire.Build(
		provider.NewConfig,
		provider.NewEngine,
		provider.NewSrv,
		provider.NewCache,
		provider.NewDB,
		provider.NewLogger,
		http.WireSet,
		provider.AppWireSet,
	)
	return &kernel.Application{}
}
