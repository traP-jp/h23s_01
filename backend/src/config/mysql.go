package config

import (
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

func GetMySqlConf() mysql.Config {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	conf := mysql.Config{
		User:      getEnvOrDefault("DB_USER", "root"),
		Passwd:    getEnvOrDefault("DB_PASSWORD", "password"),
		Net:       "tcp",
		Addr:      getEnvOrDefault("DB_HOST", "db") + ":" + getEnvOrDefault("DB_PORT", "3306"),
		DBName:    getEnvOrDefault("DB_NAME", "app"),
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	return conf
}
