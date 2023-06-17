package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23s_01/backend/src/repository"
	"github.com/traP-jp/h23s_01/backend/src/traq"
)

type userHandler struct {
	tc traq.TraqClient
	ur repository.UsersRepository
}

func NewUserHandler(tc traq.TraqClient, ur repository.UsersRepository) *userHandler {
	return &userHandler{
		tc: tc,
		ur: ur,
	}
}

func (uh *userHandler) patchUserHandler(c echo.Context) error {
	token, err := getToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	users, err := uh.tc.GetAllUsers(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}

	if err := uh.ur.RemakeUsersTable(users); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}

	return c.String(http.StatusOK, "ok")
}
