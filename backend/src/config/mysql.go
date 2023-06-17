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
		User:      getEnvOrDefault("NS_MARIADB_USER", "root"),
		Passwd:    getEnvOrDefault("NS_MARIADB_PASSWORD", "password"),
		Net:       "tcp",
		Addr:      getEnvOrDefault("NS_MARIADB_HOSTNAME", "db") + ":" + getEnvOrDefault("NS_MARIADB_PORT", "3306"),
		DBName:    getEnvOrDefault("NS_MARIADB_DATABASE", "app"),
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
		AllowNativePasswords: true,
	}

	return conf
}
