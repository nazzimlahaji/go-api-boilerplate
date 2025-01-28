package config

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
)

func SentryConfig(dsn string) (*sentryhttp.Handler, error) {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,

		EnableTracing: true,

		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		return nil, fmt.Errorf("Sentry initialization failed: %v\n", err)
	}

	// Create an instance of sentryhttp
	sentryHandler := sentryhttp.New(sentryhttp.Options{
		// Repanic: true,
		// WaitForDelivery: true,
	})

	return sentryHandler, nil
}
