package setup

import (
	"net/http"
	oapi "nomikuimatch/generated"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *OapiService) GetAreas(ctx echo.Context) error {
	res, err := s.a.GetAreas()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *OapiService) PostAreas(ctx echo.Context) error {
	var req oapi.PostAreasJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	res, err := s.a.PostAreas(&req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (s *OapiService) GetAreasAreaid(ctx echo.Context, areaid uuid.UUID) error {
	return nil
}

func (s *OapiService) PatchAreasAreaid(ctx echo.Context, areaid uuid.UUID) error {
	return nil
}

func (s *OapiService) DeleteAreasAreaid(ctx echo.Context, areaid uuid.UUID) error {
	return nil
}
