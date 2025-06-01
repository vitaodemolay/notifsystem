package entrypoint

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	internalerrors "github.com/vitaodemolay/notifsystem/pkg/internal-errors"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (any, int, error)

func (endpointFunc EndpointFunc) HandleError() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, code, err := endpointFunc(w, r)
		if err != nil {
			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(r, http.StatusInternalServerError)
			} else if errors.Is(err, internalerrors.ErrNotFound) {
				render.Status(r, http.StatusNotFound)
			} else {
				render.Status(r, http.StatusBadRequest)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, code)
		if obj != nil {
			render.JSON(w, r, obj)
			return
		}
		render.NoContent(w, r)
	})
}
