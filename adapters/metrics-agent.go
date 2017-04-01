package adapters

import (
	log "github.com/Sirupsen/logrus"
	"github.com/newrelic/go-agent"

	"github.com/chrisenytc/ullli/config"
)

var agent newrelic.Application

func LoadMetricsAgent() {
	config := newrelic.NewConfig("Ullli", config.Get().NewRelicKey)
	app, err := newrelic.NewApplication(config)

	if err != nil {
		log.Panicf("Fatal error on metrics connection: %s", err)
	}

	agent = app

	log.Info("Metrics connection established successfully.")
	log.Info("The application is now using the metrics agent.")
}
