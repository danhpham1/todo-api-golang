package controllers

import (
	"github.com/getsentry/sentry-go"
	"net/http"
)

func SentryCaptureErr(err error) {
	sentry.CaptureException(err)
}

func SentryEvent(request *http.Request, message string) {
	event := sentry.NewEvent()
	event.Level = sentry.LevelInfo
	event.Message = message
	event.Request = sentry.NewRequest(request)
	sentry.CaptureEvent(event)
}
