package area

import (
	"fmt"
	"net/http"
	oapi "nomikuimatch/generated"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (a *AreaService) GetAreas() (res []Area, err error) {
	err = a.db.Select(&res, "SELECT * FROM `Area`")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return res, nil
}

func (a *AreaService) PostAreas(req *oapi.PostAreasJSONRequestBody) (res *Area, err error) {

	uuid := uuid.New()
	req.Id = &uuid

	_, err = a.db.Exec("INSERT INTO `Area` VALUES(?,?)", *req.Id, req.Areaname)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}

	var tres = Area{Id: uuid, Areaname: req.Areaname}

	res = &tres

	return res, nil
}
