package app

import (
	"context"
	"countenum404/itmo-golang-computation/internal/model"
	"countenum404/itmo-golang-computation/internal/service"
	"countenum404/itmo-golang-computation/pkg/app"
	"go.uber.org/fx"
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

func NewGrpcServer(lc fx.Lifecycle, logger *zap.Logger, solverService service.SolverService) *GrpcServer {
	grpcServer := grpc.NewServer()
	s := &GrpcServer{Server: grpcServer, Logger: logger, SolverService: solverService}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := s.Start()
				if err != nil {
					s.Logger.Fatal("Failed to start server", zap.Error(err))
				}
			}()
			s.Logger.Info("GRPC server started")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.Logger.Info("GRPC server stopped")
			return nil
		},
	})
	return s
}

func (g *GrpcServer) Start() error {
	lis, err := net.Listen("tcp", ":8080")
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
