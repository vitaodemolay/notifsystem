package campaign

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	contract "github.com/vitaodemolay/notifsystem/internal/application/contract/campaign"
	mapper "github.com/vitaodemolay/notifsystem/internal/application/service/campaign"
	"github.com/vitaodemolay/notifsystem/internal/application/service/campaign/mock"
	"github.com/vitaodemolay/notifsystem/internal/infrastructure/web/entrypoint"
	internalerrors "github.com/vitaodemolay/notifsystem/pkg/internal-errors"
	"go.uber.org/mock/gomock"
)

type testSuite struct {
	Request    *contract.CreateCampaign
	Control    *gomock.Controller
	Service    *mock.MockCampaignService
	Controller *Controller
	request    *http.Request
	response   *httptest.ResponseRecorder
}

func fakeCreateCampaing() *contract.CreateCampaign {
	return &contract.CreateCampaign{
		Title:   "Campaign X of Test",
		Content: "Body of Campaign X of Test",
		Emails:  []string{"email1@test.com", "email2@test.com"},
	}
}

func createRequest(method string, path string, body any) *http.Request {
	if body != nil {
		jsonBody, _ := json.Marshal(body)
		return httptest.NewRequest(method, path, bytes.NewReader(jsonBody))
	}
	return httptest.NewRequest(method, path, nil)
}

func setup(t *testing.T, requestBody *contract.CreateCampaign, method string) *testSuite {
	ctrl := gomock.NewController(t)
	service := mock.NewMockCampaignService(ctrl)

	req := createRequest(method, "/v1/campaign", requestBody)
	w := httptest.NewRecorder()

	return &testSuite{
		Request: requestBody,
		Control: ctrl,
		Service: service,
		Controller: &Controller{
			service: service,
		},
		request:  req,
		response: w,
	}
}

func Test_NewController(t *testing.T) {
	// Arrange
	suite := setup(t, nil, http.MethodGet)

	// Act
	controller := NewController(suite.Service)

	// Assert
	assert.NotNil(t, controller, "Expected controller to be created")
	assert.Equal(t, suite.Service, controller.service, "Expected controller service to match")
}

func Test_Controller_Path(t *testing.T) {
	// Arrange
	suite := setup(t, nil, http.MethodGet)

	// Act
	path := suite.Controller.Path()

	// Assert
	assert.Equal(t, "/v1/campaign", path, "Expected path to be /v1/campaign")
}

func Test_GetRoutes(t *testing.T) {
	// Arrange
	suite := setup(t, nil, http.MethodGet)

	// Act
	routes := suite.Controller.GetRoutes()

	// Assert
	assert.NotEmpty(t, routes, "Expected routes to be returned")
	assert.Len(t, routes, 3, "Expected two routes to be defined")
	assert.Equal(t, http.MethodGet, routes[0].Method, "Expected first route to be GET")
	assert.Equal(t, "/{id}", routes[0].Pattern, "Expected first route pattern to be /{id}")
	assert.Equal(t, http.MethodGet, routes[1].Method, "Expected second route to be GET")
	assert.Equal(t, "/", routes[1].Pattern, "Expected second route pattern to be /")
	assert.Equal(t, http.MethodPost, routes[2].Method, "Expected second route to be POST")
	assert.Equal(t, "/", routes[2].Pattern, "Expected second route pattern to be /")
}

func Test_CreateCampaing_when_body_is_nil(t *testing.T) {
	// Arrange
	suite := setup(t, nil, http.MethodPost)
	expectedError := errors.New("request body is nil")
	suite.Service.EXPECT().CreateCampaign(gomock.Any()).Return("", expectedError).Times(1)

	// Act
	object, statusCode, err := suite.Controller.CreateCampaign(suite.response, suite.request)

	// Assert
	assert.Nil(t, object, "Expected object to be nil")
	assert.Equal(t, 0, statusCode, "Expected status code to be 0")
	assert.ErrorIs(t, err, expectedError, "Expected error to match")
}

func Test_CreateCampaing_when_body_Is_Valid(t *testing.T) {
	// Arrange
	suite := setup(t, fakeCreateCampaing(), http.MethodPost)
	expectedID := "12345"
	suite.Service.EXPECT().CreateCampaign(gomock.Any()).
		DoAndReturn(func(req *contract.CreateCampaign) (string, error) {
			_, err := mapper.MapToDomain(req)
			assert.NoError(t, err, "Expected no error when mapping to domain")
			return expectedID, nil
		}).Times(1)

	// Act
	object, statusCode, err := suite.Controller.CreateCampaign(suite.response, suite.request)

	// Assert
	assert.NotNil(t, object, "Expected object to not be nil")
	assert.Equal(t, http.StatusCreated, statusCode, "Expected status code to be 201")
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, map[string]string{"campaign_id": expectedID}, object, "Expected response to contain campaign ID")
}

func Test_GetCampaignByID_when_id_is_empty(t *testing.T) {
	// Arrange
	suite := setup(t, nil, http.MethodGet)
	suite.request = createRequest(http.MethodGet, "/v1/campaign/", nil)
	suite.Service.EXPECT().GetCampaignByID(gomock.Any()).Return(nil, nil).Times(0)

	// Act
	object, statusCode, err := suite.Controller.GetCampaignByID(suite.response, suite.request)

	// Assert
	assert.Nil(t, object, "Expected object to be nil")
	assert.Equal(t, 0, statusCode, "Expected status code to be 0")
	assert.ErrorIs(t, err, internalerrors.ErrBadRequest, "Expected error to match")
}

func Test_GetCampaignByID_when_id_is_valid(t *testing.T) {
	// Arrange
	suite := setup(t, nil, http.MethodGet)
	campaignID := "12345"
	routepath := "/v1/campaign/"
	suite.request = createRequest(http.MethodGet, routepath+campaignID, nil)
	expectedCampaign := &contract.Campaign{ID: campaignID, Title: "Test Campaign"}
	suite.Service.EXPECT().GetCampaignByID(campaignID).Return(expectedCampaign, nil).Times(1)
	json, _ := json.Marshal(expectedCampaign)
	expectedCampaignJson := string(json)

	var handler entrypoint.EndpointFunc = suite.Controller.GetCampaignByID
	r := chi.NewRouter()
	r.Get(routepath+"{id}", handler.HandleError())

	// Act
	r.ServeHTTP(suite.response, suite.request)

	// Assert
	assert.Equal(t, http.StatusOK, suite.response.Code, "Expected status code to be 200")
	assert.JSONEq(t, expectedCampaignJson, suite.response.Body.String(), "Expected response body to match campaign")
}
