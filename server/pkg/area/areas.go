package area

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	oapi "nomikuimatch/generated"
)

type Areaservice struct{}

func (s *Areaservice) GetAreas(ctx echo.Context) error {
	area := "area"
	return ctx.JSON(http.StatusOK, &oapi.Area{
		Areaname: area,
		Id:       uuid.New(),
	})
}

func (s *Areaservice) PostAreas(ctx echo.Context) error {
	return nil
}
