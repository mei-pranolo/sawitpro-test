package delivery

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	m "github.com/SawitProRecruitment/UserService/types"
	usecase "github.com/SawitProRecruitment/UserService/usecase/mock"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestServer_PostEstate_Positive(t *testing.T) {
	mockUC := usecase.NewMockUsecaseInterface(gomock.NewController(t))
	request := `{"length":5,"width":5}`
	response := `{"id":"aaa"}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate", strings.NewReader(request))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Server{Usecase: mockUC}

	mockUC.EXPECT().CreateEstate(gomock.Any(), 5, 5).Return("aaa", nil)

	// Assertions
	if assert.NoError(t, h.PostEstate(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, response, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestServer_PostEstate_JSONUnmarshal_error(t *testing.T) {
	mockUC := usecase.NewMockUsecaseInterface(gomock.NewController(t))
	request := `{"length":5,"width":5`
	response := `{"message":"unexpected end of JSON input"}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate", strings.NewReader(request))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Server{Usecase: mockUC}

	// Assertions
	if assert.NoError(t, h.PostEstate(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, response, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestServer_PostEstate_UsecaseError(t *testing.T) {
	mockUC := usecase.NewMockUsecaseInterface(gomock.NewController(t))
	request := `{"length":5,"width":5}`
	response := `{"message":"usecase"}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate", strings.NewReader(request))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Server{Usecase: mockUC}

	mockUC.EXPECT().CreateEstate(gomock.Any(), 5, 5).Return("", errors.New("usecase"))

	// Assertions
	if assert.NoError(t, h.PostEstate(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, response, strings.TrimSuffix(rec.Body.String(), "\n")) // the actual response have \n on it
	}
}

func TestServer_PostEstateIdTree_Positive(t *testing.T) {
	mockUC := usecase.NewMockUsecaseInterface(gomock.NewController(t))
	request := `{"x":5,"y":5, "height":5}`
	response := `{"id":"00000000-0000-0000-0000-000000000000"}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/00000000-0000-0000-0000-000000000000/tree", strings.NewReader(request))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Server{Usecase: mockUC}

	mockUC.EXPECT().CreateTree(gomock.Any(), "00000000-0000-0000-0000-000000000000", m.Tree{X: 5, Y: 5, Height: 5}).Return("00000000-0000-0000-0000-000000000000", nil)

	// Assertions
	if assert.NoError(t, h.PostEstateIdTree(c, openapi_types.UUID{})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, response, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestServer_PostEstateIdTree_JSONUnmarshal_error(t *testing.T) {
	mockUC := usecase.NewMockUsecaseInterface(gomock.NewController(t))
	request := `{"x":5,"y":5, "height":5`
	response := `{"message":"unexpected end of JSON input"}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/00000000-0000-0000-0000-000000000000/tree", strings.NewReader(request))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Server{Usecase: mockUC}

	// Assertions
	if assert.NoError(t, h.PostEstateIdTree(c, openapi_types.UUID{})) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, response, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestServer_PostEstateIdTree_Usecase_error(t *testing.T) {
	mockUC := usecase.NewMockUsecaseInterface(gomock.NewController(t))
	request := `{"x":5,"y":5, "height":5}`
	response := `{"message":"usecase"}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/00000000-0000-0000-0000-000000000000/tree", strings.NewReader(request))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Server{Usecase: mockUC}

	mockUC.EXPECT().CreateTree(gomock.Any(), "00000000-0000-0000-0000-000000000000", m.Tree{X: 5, Y: 5, Height: 5}).Return("", errors.New("usecase"))

	// Assertions
	if assert.NoError(t, h.PostEstateIdTree(c, openapi_types.UUID{})) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, response, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestServer_GetEstateIdStats_Positive(t *testing.T) {
	mockUC := usecase.NewMockUsecaseInterface(gomock.NewController(t))
	response := `{"count":3,"max":10,"median":5,"min":3}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/estate/00000000-0000-0000-0000-000000000000/stats", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Server{Usecase: mockUC}

	mockUC.EXPECT().GetEstateStats(gomock.Any(), "00000000-0000-0000-0000-000000000000").Return(m.Stats{Count: 3, Max: 10, Min: 3, Median: 5}, nil)

	// Assertions
	if assert.NoError(t, h.GetEstateIdStats(c, openapi_types.UUID{})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, response, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestServer_GetEstateIdStats_Usecase_error(t *testing.T) {
	mockUC := usecase.NewMockUsecaseInterface(gomock.NewController(t))
	response := `{"message":"usecase"}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/estate/00000000-0000-0000-0000-000000000000/stats", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Server{Usecase: mockUC}

	mockUC.EXPECT().GetEstateStats(gomock.Any(), "00000000-0000-0000-0000-000000000000").Return(m.Stats{}, errors.New("usecase"))

	// Assertions
	if assert.NoError(t, h.GetEstateIdStats(c, openapi_types.UUID{})) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, response, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestServer_GetEstateIdDronePlan_Positive(t *testing.T) {
	mockUC := usecase.NewMockUsecaseInterface(gomock.NewController(t))
	response := `{"distance":50}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/estate/00000000-0000-0000-0000-000000000000/stats", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Server{Usecase: mockUC}

	mockUC.EXPECT().GetDroneDistance(gomock.Any(), "00000000-0000-0000-0000-000000000000").Return(50, nil)

	// Assertions
	if assert.NoError(t, h.GetEstateIdDronePlan(c, openapi_types.UUID{})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, response, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestServer_GetEstateIdDronePlan_Usecase_error(t *testing.T) {
	mockUC := usecase.NewMockUsecaseInterface(gomock.NewController(t))
	response := `{"message":"usecase"}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/estate/00000000-0000-0000-0000-000000000000/stats", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Server{Usecase: mockUC}

	mockUC.EXPECT().GetDroneDistance(gomock.Any(), "00000000-0000-0000-0000-000000000000").Return(0, errors.New("usecase"))

	// Assertions
	if assert.NoError(t, h.GetEstateIdDronePlan(c, openapi_types.UUID{})) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, response, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}
