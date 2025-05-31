package campaign

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	contract "github.com/vitaodemolay/notifsystem/internal/application/contract/campaign"
	mapper "github.com/vitaodemolay/notifsystem/internal/application/service/campaign"
	"github.com/vitaodemolay/notifsystem/internal/application/service/campaign/mock"
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
	assert.Len(t, routes, 2, "Expected two routes to be defined")
	assert.Equal(t, http.MethodGet, routes[0].Method, "Expected first route to be GET")
	assert.Equal(t, http.MethodPost, routes[1].Method, "Expected second route to be POST")
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
