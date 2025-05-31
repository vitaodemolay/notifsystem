package basictest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/web/entrypoint"
)

func TestController_Test(t *testing.T) {
	controller := NewController()

	req, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := entrypoint.EndpointFunc(controller.Test).HandleError()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code to be 200 OK")
	assert.JSONEq(t, `{"message":"Hello, this is a test!"}`, strings.TrimSpace(rr.Body.String()), "Expected response body to match")
}
