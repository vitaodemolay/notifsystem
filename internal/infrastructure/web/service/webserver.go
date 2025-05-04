package service

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/web/entrypoint"
)

type WebServer struct {
	Port              string
	Router            *chi.Mux
	custonLoggerIsSet bool
}

func CreateWebServer(port string) (*WebServer, error) {
	if port == "" {
		return nil, errors.New("port cannot be empty")
	}

	router := chi.NewRouter()
	return &WebServer{
		Port:              port,
		Router:            router,
		custonLoggerIsSet: false,
	}, nil
}

func (ws *WebServer) Start() error {
	return http.ListenAndServe(ws.Port, ws.Router)
}

func (ws *WebServer) SetLogger(logger *httplog.Logger) {
	ws.Router.Use(httplog.RequestLogger(logger))
	ws.custonLoggerIsSet = true
}

func (ws *WebServer) InitalizeRoutes(routers ...entrypoint.Router) {
	ws.Router.Use(middleware.RequestID)
	ws.Router.Use(middleware.RealIP)
	ws.Router.Use(middleware.Recoverer)

	if !ws.custonLoggerIsSet {
		ws.Router.Use(middleware.Logger)
	}

	ws.Router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	ws.Router.Route(entrypoint.BasePath, func(r chi.Router) {
		for _, router := range routers {
			r.Route(router.Path(), func(r chi.Router) {
				for _, route := range router.GetRoutes() {
					if route.Middlewares != nil {
						r.Use(route.Middlewares)
					}
					r.Method(route.Method, route.Pattern, route.Handler)
				}
			})
		}
	})
}
