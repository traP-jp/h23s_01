package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/traP-jp/h23s_01/backend/src/config"
)

func main() {
	conf := config.GetMySqlConf()

	var db *sqlx.DB
	var err error

	for i := 0; i < 10; i++ {
		log.Println(i)

		db, err = sqlx.Connect("mysql", conf.FormatDSN())
		log.Println(conf.FormatDSN())
		if err != nil && i == 9 {
			log.Fatal(err)
		}
		if err == nil {
			break
		}

		time.Sleep(time.Second * time.Duration(i+1))
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		var i int64
		err := db.Get(&i, "SELECT COUNT(*) FROM users")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, i)
	})

	e.Start(":8080")
}
