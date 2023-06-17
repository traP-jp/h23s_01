package config

import "os"

func GetMode() string {
	return os.Getenv("MODE")
}
