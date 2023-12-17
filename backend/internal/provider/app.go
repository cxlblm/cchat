package provider

import (
	"cchart/internal/kernel"
	"github.com/google/wire"
)

var AppWireSet = wire.NewSet(wire.Struct(new(kernel.Application), "*"))
