package entrypoint

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	internalerrors "github.com/vitaodemolay/notifsystem/pkg/internal-errors"
)

type testSuite struct {
	handler  EndpointFunc
	request  *http.Request
	response *httptest.ResponseRecorder
}

func setup(expectedObject any, expectedCode int, expectedError error) *testSuite {
	endpoint := func(w http.ResponseWriter, r *http.Request) (any, int, error) {
		return expectedObject, expectedCode, expectedError
	}

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	return &testSuite{
		handler:  endpoint,
		request:  req,
		response: w,
	}
}

func Test_HandleError_when_endpoint_returns_internal_error(t *testing.T) {
	// Arrange
	suite := setup(nil, 0, internalerrors.ErrInternal)

	// Act
	suite.handler.HandleError().ServeHTTP(suite.response, suite.request)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, suite.response.Code)
}

func Test_HandleError_when_endpoint_returns_not_found_error(t *testing.T) {
	// Arrange
	suite := setup(nil, 0, internalerrors.ErrNotFound)

	// Act
	suite.handler.HandleError().ServeHTTP(suite.response, suite.request)

	// Assert
	assert.Equal(t, http.StatusNotFound, suite.response.Code)
}

func Test_HandleError_when_endpoint_returns_bad_request_error(t *testing.T) {
	// Arrange
	suite := setup(nil, 0, internalerrors.ErrBadRequest)

	// Act
	suite.handler.HandleError().ServeHTTP(suite.response, suite.request)

	// Assert
	assert.Equal(t, http.StatusBadRequest, suite.response.Code)
}

func Test_HandleError_when_endpoint_returns_generic_error(t *testing.T) {
	// Arrange
	expectedError := "generic error"
	suite := setup(nil, 0, errors.New(expectedError))

	// Act
	suite.handler.HandleError().ServeHTTP(suite.response, suite.request)

	// Assert
	assert.Equal(t, http.StatusBadRequest, suite.response.Code)
	assert.JSONEq(t, `{"error": "`+expectedError+`"}`, suite.response.Body.String())
}

func Test_HandleError_when_endpoint_returns_object_and_code(t *testing.T) {
	// Arrange
	expectedObject := map[string]string{"message": "success"}
	expectedCode := http.StatusCreated
	suite := setup(expectedObject, expectedCode, nil)

	// Act
	suite.handler.HandleError().ServeHTTP(suite.response, suite.request)

	// Assert
	assert.Equal(t, expectedCode, suite.response.Code)
	assert.JSONEq(t, `{"message": "success"}`, suite.response.Body.String())
}

func Test_HandleError_when_endpoint_returns_nil_object_and_code(t *testing.T) {
	// Arrange
	expectedCode := http.StatusNoContent
	suite := setup(nil, expectedCode, nil)

	// Act
	suite.handler.HandleError().ServeHTTP(suite.response, suite.request)

	// Assert
	assert.Equal(t, expectedCode, suite.response.Code)
	assert.Empty(t, suite.response.Body.String())
}
