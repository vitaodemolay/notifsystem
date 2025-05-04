package container

import (
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/web/controllers/basictest"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/web/entrypoint"
)

type EntryPoint struct {
	BasicTest *basictest.Controller
}

func NewEntryPointContainer() (*EntryPoint, error) {

	return &EntryPoint{
		BasicTest: basictest.NewController(),
	}, nil
}

func (e *EntryPoint) GetControllers() []entrypoint.Router {
	return []entrypoint.Router{
		e.BasicTest,
	}
}
