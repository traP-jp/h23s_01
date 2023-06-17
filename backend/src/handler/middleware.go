package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23s_01/backend/src/config"
)

func (tc *traqClient) checkTraqLoginMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to get session: %v", err.Error()))
		}

		switch validToken(sess) {
		case expired:
			return echo.NewHTTPError(http.StatusUnauthorized, "access token is expired")
		case noToken:
			return echo.NewHTTPError(http.StatusUnauthorized, "no token")
		}

		return next(c)
	}
}

func (tc *traqClient) checkAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := getToken(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		user, err := tc.tc.GetMe(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		admins := config.GetAdmins()

		for i := range admins {
			if admins[i] == user.Name {
				return next(c)
			}
		}

		return echo.NewHTTPError(http.StatusForbidden, "you are not admin")
	}
}

func validToken(session *sessions.Session) tokenStatus {
	_, ok := session.Values["access_token"].(string)
	if !ok {
		return noToken
	}

	expiresAt := session.Values["expires_at"].(time.Time)
	if expiresAt.Before(time.Now()) {
		return expired
	}

	return valid
}
