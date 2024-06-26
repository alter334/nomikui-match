package area

import (
	"fmt"
	"net/http"
	oapi "nomikuimatch/generated"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (a *AreaService) GetAreas() (res *oapi.Area, err error) {
	err = a.db.Select(res, "SELECT * FROM `area`")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return res, nil
}

func (a *AreaService) PostAreas(req *oapi.PostAreasJSONRequestBody) (res *oapi.Area, err error) {
	req.Id = uuid.New()
	_, err = a.db.Exec("INSERT INTO `area` VALUES(?,?)", req.Id, req.Areaname)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return req, nil
}
