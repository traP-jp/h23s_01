package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srinathgs/mysqlstore"
	"github.com/traP-jp/h23s_01/backend/src/config"
	"github.com/traP-jp/h23s_01/backend/src/reading"
	"github.com/traP-jp/h23s_01/backend/src/repository/implement"
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
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} |${time_rfc3339} |${method} |${host}${uri}\t|error:\"${error}\"\t|latency:${latency_human}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{config.GetAccessControlAllowOrigin()},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderAccessControlAllowOrigin, echo.HeaderOrigin, echo.HeaderXHTTPMethodOverride},
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

	tc := traq.NewTraqClient(gotraq.NewAPIClient(gotraq.NewConfiguration()))
	ur := implement.NewUsers(db)

	client := NewTraqClient(tc)
	channelHandler := NewChannelHandler(tc, implement.NewChannels(db))
	userHandler := NewUserHandler(tc, ur)
	messagesHandler := NewMessageHandler(tc, implement.NewChannels(db), ur, reading.NewTokenizer())
	scoreHandler := NewScoreHandler(tc, implement.NewScore(db))

	oauth := api.Group("/oauth2")
	oauth.GET("/authorize", authorizeHandler)
	oauth.GET("/callback", callbackHandler)

	admin := api.Group("/admin")
	if config.GetMode() == "production" {
		admin.Use(client.checkAdminMiddleware)
	}
	admin.PATCH("/channel", channelHandler.patchChennelsHandler)
	admin.PATCH("/user", userHandler.patchUserHandler)

	api.POST("/post", client.postScoreHandler, client.checkTraqLoginMiddleware)
	api.POST("/score", scoreHandler.registerScoreHandler, client.checkTraqLoginMiddleware)
	api.GET("/score/highest", scoreHandler.getHighestScoreHandler, client.checkTraqLoginMiddleware)
	api.GET("/me", client.getMeHandler, client.checkTraqLoginMiddleware)
	api.GET("/message", messagesHandler.getMessagesHandler, client.checkTraqLoginMiddleware)

	e.Start(":8080")
}
