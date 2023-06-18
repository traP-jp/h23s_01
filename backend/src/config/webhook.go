package config

import (
	"os"

	traqwriter "github.com/ras0q/traq-writer"
)

func GetWebhookConfig() *traqwriter.TraqWebhookWriter {
	webhookId, ok1 := os.LookupEnv("WEBHOOK_ID")
	webhookSecret, ok2 := os.LookupEnv("WEBHOOK_SECRET")
	if !(ok1 && ok2) {
		return nil
	}
	return traqwriter.NewTraqWebhookWriter(webhookId, webhookSecret, traqwriter.DefaultHTTPOrigin)
}
