// Code generated by git.repo.services.lenvendo.ru/grade-factor/go-kit-service-generator  REMOVE THIS STRING ON EDIT OR DO NOT EDIT.
//go:generate mockgen -destination service_mock.go -package health  "git.repo.services.lenvendo.ru/grade-factor/echo/pkg/health Service
package health

import (
	"context"

	_ "github.com/golang/mock/mockgen/model"
)

// Service is the interface that provides health methods.
type Service interface {

	// Liveness returns a error if service doesn`t live.
	// The kubelet uses liveness probes to know when to restart a Container.
	Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error)

	// Readiness returns a error if service doesn`t ready.
	// Service doesn`t ready by default.
	// The kubelet uses readiness probes to know when a Container is ready to start accepting traffic.
	Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error)

	// Version returns buid time, last commit and version app
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}
