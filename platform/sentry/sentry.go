package sentry

import (
	"errors"
	"time"

	"example.com/refcode/v1/pkg/configs"
	"github.com/getsentry/sentry-go"
)

func init() {
	sentry.Init(sentry.ClientOptions{
		Dsn: configs.SentryDNS,
	})
	sentry.CaptureException(errors.New("my error"))

}

func CaptureException(err error) {
	sentry.Init(sentry.ClientOptions{
		Dsn: configs.SentryDNS,
	})
	sentry.CaptureException(err)
	sentry.Flush(time.Second * 5)
}

func CaptureMessage(message string) {
	sentry.Init(sentry.ClientOptions{
		Dsn: configs.SentryDNS,
	})
	sentry.CaptureMessage(message)
	sentry.Flush(time.Second * 5)
}
