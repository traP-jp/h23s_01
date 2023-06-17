package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23s_01/backend/src/config"
	"github.com/traP-jp/h23s_01/backend/src/handler"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	conf := config.GetMySqlConf()

	var db *sqlx.DB

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
	handler.SetUpRoutes(e, db)

}
