package entrypoint

import (
	"net/http"
)

const (
	BasePath = "/api"
)

type Router interface {
	GetRoutes() []Route
	Path() string
}

type Route struct {
	Method      string
	Pattern     string
	Handler     EndpointFunc
	Middlewares func(http.Handler) http.Handler
}
