package health

import (
	"context"
	"errors"
)

var (
	BuildTime string
	Commit    string
	Version   string
)

type healthService struct {
	reps []repo
}

type repo interface {
	IsReady() bool
}

// ErrServiceNotReady is returned when service doesn`t ready
var ErrServiceNotReady = errors.New("service not ready")

func NewHealthService(reps ...repo) Service {
	return &healthService{
		reps: reps,
	}
}

func (s *healthService) Liveness(ctx context.Context, req *LivenessRequest) (resp *LivenessResponse, err error) {
	return &LivenessResponse{
		Status: "ok",
	}, nil
}

func (s *healthService) Readiness(ctx context.Context, req *ReadinessRequest) (resp *ReadinessResponse, err error) {
	for _, r := range s.reps {
		if !r.IsReady() {
			return nil, ErrServiceNotReady
		}
	}
	return &ReadinessResponse{
		Status: "ok",
	}, nil
}

func (s *healthService) Version(ctx context.Context, req *VersionRequest) (resp *VersionResponse, err error) {
	return &VersionResponse{
		BuildTime: BuildTime,
		Commit:    Commit,
		Version:   Version,
	}, nil
}
