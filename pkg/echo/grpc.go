// Code generated by git.repo.services.lenvendo.ru/grade-factor/go-kit-service-generator  REMOVE THIS STRING ON EDIT OR DO NOT EDIT.
package echo

import (
	"context"
	"errors"

	pb "git.repo.services.lenvendo.ru/grade-factor/echo/internal/echopb"
	"git.repo.services.lenvendo.ru/grade-factor/echo/tools/logging"
	"git.repo.services.lenvendo.ru/grade-factor/echo/tools/tracing"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	googlegrpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type grpcServer struct {
	getEcho grpctransport.Handler
}

type ContextGRPCKey struct{}

type GRPCInfo struct{}

// NewGRPCServer makes a set of endpoints available as a gRPC echoServer.
func NewGRPCServer(ctx context.Context, s Service) pb.EchoServiceServer {
	logger := logging.FromContext(ctx)
	logger = log.With(logger, "grpc handler", "echo")
	tracer := tracing.FromContext(ctx)

	options := []grpctransport.ServerOption{
		// grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(grpcToContext()),
		grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "grpc server", logger)),
		grpctransport.ServerFinalizer(closeGRPCTracer()),
	}

	return &grpcServer{
		getEcho: grpctransport.NewServer(
			makeGetEchoEndpoint(s),
			decodeGRPCGetEchoRequest,
			encodeGRPCGetEchoResponse,
			options...,
		),
	}
}

func JoinGRPC(ctx context.Context, s Service) func(*googlegrpc.Server) {
	return func(g *googlegrpc.Server) {
		pb.RegisterEchoServiceServer(g, NewGRPCServer(ctx, s))
	}
}

func grpcToContext() grpc.ServerRequestFunc {
	return func(ctx context.Context, md metadata.MD) context.Context {
		return context.WithValue(ctx, ContextGRPCKey{}, GRPCInfo{})
	}
}

func closeGRPCTracer() grpc.ServerFinalizerFunc {
	return func(ctx context.Context, err error) {
		span := stdopentracing.SpanFromContext(ctx)
		span.Finish()
	}
}

func (s *grpcServer) GetEcho(ctx context.Context, req *pb.GetEchoListRequest) (*pb.GetEchoListResponse, error) {
	_, rep, err := s.getEcho.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetEchoListResponse), nil
}

func decodeGRPCGetEchoRequest(_ context.Context, request interface{}) (interface{}, error) {
	inReq, ok := request.(*pb.GetEchoListRequest)
	if !ok {
		return nil, errors.New("decodeGRPCGetEchoRequest wrong request")
	}

	req := PBToGetEchoListRequest(inReq)
	if err := validate(req); err != nil {
		return nil, err
	}
	return *req, nil
}

func encodeGRPCGetEchoResponse(_ context.Context, response interface{}) (interface{}, error) {
	inResp, ok := response.(*GetEchoListResponse)
	if !ok {
		return nil, errors.New("encodeGRPCGetEchoResponse wrong response")
	}

	return GetEchoListResponseToPB(inResp), nil
}

func EchoToPB(d *Echo) *pb.Echo {
	if d == nil {
		return nil
	}

	resp := pb.Echo{
		Id:       d.Id,
		Title:    d.Title,
		Reminder: d.Reminder,
	}

	return &resp
}

func PBToEcho(d *pb.Echo) *Echo {
	if d == nil {
		return nil
	}

	resp := Echo{
		Id:       d.Id,
		Title:    d.Title,
		Reminder: d.Reminder,
	}

	return &resp
}

func GetEchoListRequestToPB(d *GetEchoListRequest) *pb.GetEchoListRequest {
	if d == nil {
		return nil
	}

	resp := pb.GetEchoListRequest{}

	return &resp
}

func PBToGetEchoListRequest(d *pb.GetEchoListRequest) *GetEchoListRequest {
	if d == nil {
		return nil
	}

	resp := GetEchoListRequest{}

	return &resp
}

func GetEchoListResponseToPB(d *GetEchoListResponse) *pb.GetEchoListResponse {
	if d == nil {
		return nil
	}

	resp := pb.GetEchoListResponse{
		Err: d.Err,
	}

	for _, v := range d.Echos {
		resp.Echos = append(resp.Echos, EchoToPB(&v))
	}

	return &resp
}

func PBToGetEchoListResponse(d *pb.GetEchoListResponse) *GetEchoListResponse {
	if d == nil {
		return nil
	}

	resp := GetEchoListResponse{
		Err: d.Err,
	}

	for _, v := range d.Echos {
		resp.Echos = append(resp.Echos, *PBToEcho(v))
	}

	return &resp
}
