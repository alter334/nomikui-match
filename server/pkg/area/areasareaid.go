package area

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (a *AreaService) GetAreasAreaid(param uuid.UUID) (res *Area, err error) {
	var result Area
	err = a.db.Get(&result, "SELECT * FROM `Area` WHERE `id` = ?", param)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf(":::DB Select Error:%+v:::", err))
	}
	res = &result
	return res, nil
}

func (a *AreaService) PatchAreasAreaid(param uuid.UUID) error {
	return nil
}

func (a *AreaService) DeleteAreasAreaid(param uuid.UUID) (res *Area, err error) {
	var result Area
	err = a.db.Get(&result, "SELECT * FROM `Area` WHERE `id` = ?", param)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf(":::DB Select Error:%+v:::", err))
	}
	res = &result
	_, err = a.db.Exec("DELETE FROM `Area` WHERE `id` = ?", param)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf(":::DB Delete Error:%+v:::", err))
	}
	return res, nil
}
