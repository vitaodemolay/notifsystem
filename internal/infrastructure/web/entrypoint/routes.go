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
	Middleware() func(http.Handler) http.Handler
}

type Route struct {
	Method  string
	Pattern string
	Handler EndpointFunc
}
