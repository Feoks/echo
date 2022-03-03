// Code generated by git.repo.services.lenvendo.ru/grade-factor/go-kit-service-generator  REMOVE THIS STRING ON EDIT OR DO NOT EDIT.
package echo

import (
	"context"
	"strconv"

	"github.com/getsentry/sentry-go"
)

func NewSentryService(s Service) Service {
	return &sentryService{s}
}

type sentryService struct {
	Service
}

type sentryLog interface {
	SentryLog() []interface{}
}

func (s *sentryService) getSentryLog(req interface{}, resp interface{}) (out map[string][]interface{}) {
	out = make(map[string][]interface{})
	if sentry, ok := interface{}(req).(sentryLog); ok {
		out["request"] = append(out["request"], sentry.SentryLog()...)
	}

	if sentry, ok := interface{}(resp).(sentryLog); ok {
		out["response"] = append(out["response"], sentry.SentryLog()...)
	}
	return
}

func (s *sentryService) GetEcho(ctx context.Context, req *GetEchoListRequest) (resp *GetEchoListResponse, err error) {
	defer func() {
		if err != nil {
			log := s.getSentryLog(req, resp)
			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetTag("code", strconv.Itoa(getHTTPStatusCode(err)))
				scope.SetTag("method", "GetEcho")
				scope.SetExtra("request", log["request"])
				scope.SetExtra("response", log["response"])
			})
			sentry.CaptureException(err)
		}
	}()
	return s.Service.GetEcho(ctx, req)
}
