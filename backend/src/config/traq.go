package config

import (
	"os"

	traqoauth2 "github.com/ras0q/traq-oauth2"
)

func GetTraqClientConf() *traqoauth2.Config {
	clientId, ok := os.LookupEnv("TRAQ_CLIENT_ID")
	if !ok {
		panic("TRAQ_CLIENT_ID is not set")
	}
	traqRedirectUrl, ok := os.LookupEnv("TRAQ_REDIRECT_URL")
	if !ok {
		panic("TRAQ_REDIRECT_URL is not set")
	}
	return traqoauth2.NewConfig(clientId, traqRedirectUrl)
}
