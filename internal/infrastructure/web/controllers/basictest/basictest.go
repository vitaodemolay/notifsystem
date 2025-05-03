package basictest

import (
	"net/http"

	"github.com/vitaodemolay/notifsystem/internal/infrastructure/web/entrypoint"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Path() string {
	return "/v1/test"
}

func (c *Controller) GetRoutes() []entrypoint.Route {
	return []entrypoint.Route{
		{
			Method:  http.MethodGet,
			Pattern: "/",
			Handler: c.Test,
		},
	}
}

func (c *Controller) Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, this is a test!"))
}
