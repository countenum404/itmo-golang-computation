package app

import (
	"context"
	"countenum404/itmo-golang-computation/internal/model"
	"countenum404/itmo-golang-computation/internal/service"
	"countenum404/itmo-golang-computation/pkg/app"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"unsafe"
)

type GrpcServer struct {
	app.UnimplementedSolverSvcServer
	Logger        *zap.Logger
	SolverService service.SolverService
	Server        *grpc.Server
}

func NewGrpcServer(server *grpc.Server, logger *zap.Logger, solverService service.SolverService) *GrpcServer {
	return &GrpcServer{Server: server, Logger: logger, SolverService: solverService}
}

func (g *GrpcServer) Start() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	app.RegisterSolverSvcServer(grpcServer, g)
	return grpcServer.Serve(lis)
}

func (g *GrpcServer) Solve(ctx context.Context, cr *app.CalcRequest) (*app.SolutionResponse, error) {
	solution, err := g.SolverService.Solve(convertCalcRequest(cr))
	if err != nil {
		return nil, err
	}
	s := convertSolutionResponse(solution)
	return s, nil
}

func convertCalcRequest(cr *app.CalcRequest) *model.CalcRequest {
	return (*model.CalcRequest)(unsafe.Pointer(cr))
}

func convertSolutionResponse(sr *model.SolutionResponse) *app.SolutionResponse {
	items := sr.Items
	newItems := make([]*app.PrintResult, len(items))
	for i := 0; i < len(items); i++ {
		newItems[i] = &app.PrintResult{Var: items[i].Var, Value: items[i].Value}
	}
	return &app.SolutionResponse{Items: newItems}
}
