package config

import "os"

func GetAccessControlAllowOrigin() string {
	origin, ok := os.LookupEnv("ACCESS_CONTROL_ALLOW_ORIGIN")
	if !ok {
		return "*"
	}
	return origin
}
