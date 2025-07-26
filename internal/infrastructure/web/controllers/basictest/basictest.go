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
			Method:  http.MethodGet,
			Pattern: "/",
			Handler: c.Test,
		},
	}
}

func (c *Controller) Middleware() func(http.Handler) http.Handler {
	return c.DefineContextMiddlerware
}

func (c *Controller) Test(w http.ResponseWriter, r *http.Request) (any, int, error) {
	log.Println("Test endpoint hit")
	user := r.Context().Value(userContextKey)
	log.Println("user from context: ", user)

	return map[string]string{"message": "Hello, this is a test!"}, http.StatusOK, nil
}

type contextKey string

const userContextKey contextKey = "user"

func (c *Controller) DefineContextMiddlerware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), userContextKey, "vitinho")
		log.Println("Context middleware hit")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
