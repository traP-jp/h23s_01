package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srinathgs/mysqlstore"
	"github.com/traP-jp/h23s_01/backend/src/traq"
	gotraq "github.com/traPtitech/go-traq"
)

func SetUpRoutes(e *echo.Echo, db *sqlx.DB) {
	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "session", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
		panic(err)
	}

	e.Use(session.Middleware(store))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error {
		var i int64
		err := db.Get(&i, "SELECT COUNT(*) FROM users")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, i)
	})

	api := e.Group("/api")

	client := NewTraqClient(traq.NewTraqClient(gotraq.NewAPIClient(gotraq.NewConfiguration())))

	oauth := api.Group("/oauth2")
	oauth.GET("/authorize", authorizeHandler)
	oauth.GET("/callback", callbackHandler)

	api.GET("/me", client.getMeHandler, client.checkTraqLoginMiddleware)

	e.Start(":8080")
}
