package setup

import (
	oapi "nomikuimatch/generated"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *OapiService) GetUsers(ctx echo.Context) error {

	return nil
}

func (s *OapiService) PostUsers(ctx echo.Context) error {

	return nil
}

func (s *OapiService) GetUsersUserid(ctx echo.Context, userid uuid.UUID) error {
	return nil
}

func (s *OapiService) PatchUsersUserid(ctx echo.Context, userid uuid.UUID) error {
	return nil
}

func (s *OapiService) DeleteUsersUserid(ctx echo.Context, userid uuid.UUID) error {
	return nil
}

func (s *OapiService) GetUsersUseridFavorite(ctx echo.Context, userid uuid.UUID) error {
	return nil
}

func (s *OapiService) PostUsersUseridFavorite(ctx echo.Context, userid uuid.UUID, restaurantid oapi.PostUsersUseridFavoriteParams) error {
	return nil
}

func (s *OapiService) DeleteUsersUseridFavorite(ctx echo.Context, userid uuid.UUID, restaurantid oapi.DeleteUsersUseridFavoriteParams) error {
	return nil
}
