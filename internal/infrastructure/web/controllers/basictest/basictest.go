package basictest

import (
	"context"
	"log"
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
			Method:      http.MethodGet,
			Pattern:     "/",
			Handler:     c.Test,
			Middlewares: c.DefineContextMiddlerware,
		},
	}
}

func (c *Controller) Test(w http.ResponseWriter, r *http.Request) {
	log.Println("Test endpoint hit")
	user := r.Context().Value("user")
	log.Println("user from context: ", user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, this is a test!"))
}

func (c *Controller) DefineContextMiddlerware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "vitinho")
		log.Println("Context middleware hit")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
