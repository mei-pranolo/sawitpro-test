package delivery

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	m "github.com/SawitProRecruitment/UserService/types"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (s *Server) PostEstate(ctx echo.Context) error {
	rawBody, _ := io.ReadAll(ctx.Request().Body)
	var body m.Estate
	if err := json.Unmarshal(rawBody, &body); err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	id, err := s.Usecase.CreateEstate(ctx.Request().Context(), body.Length, body.Width)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, generated.CreateResponse{Id: id})
}

func (s *Server) PostEstateIdTree(ctx echo.Context, id openapi_types.UUID) error {
	rawBody, _ := io.ReadAll(ctx.Request().Body)
	var body m.Tree
	if err := json.Unmarshal(rawBody, &body); err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	id_returned, err := s.Usecase.CreateTree(ctx.Request().Context(), id.String(), m.Tree{X: body.X, Y: body.Y, Height: body.Height})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, generated.CreateResponse{Id: id_returned})
}

func (s *Server) GetEstateIdStats(ctx echo.Context, id openapi_types.UUID) error {
	stats, err := s.Usecase.GetEstateStats(ctx.Request().Context(), id.String())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, generated.EstateStats{Count: stats.Count, Max: stats.Max, Min: stats.Min, Median: stats.Median})
}

func (s *Server) GetEstateIdDronePlan(ctx echo.Context, id openapi_types.UUID) error {
	distance, err := s.Usecase.GetDroneDistance(ctx.Request().Context(), id.String())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, generated.DroneDistance{Distance: distance})
}
