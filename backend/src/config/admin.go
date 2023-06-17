package config

import (
	"os"
	"strings"
)

func GetAdmins() []string {
	admins, ok := os.LookupEnv("ADMINS")
	if !ok {
		panic("ADMINS is not set")
	}
	return strings.Split(admins, ",")
}
