package service

import (
	"countenum404/itmo-golang-computation/internal/core"
	"countenum404/itmo-golang-computation/internal/model"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("service",
	fx.Provide(fx.Annotate(core.NewBasicSolver, fx.As(new(core.Solver)))),
	fx.Provide(fx.Annotate(NewSolverService, fx.As(new(SolverService)))),
)

type SolverService interface {
	Solve(req *model.CalcRequest) (*model.SolutionResponse, error)
}

type BaseSolverService struct {
	Solver core.Solver
	Logger *zap.Logger
}

func NewSolverService() BaseSolverService {
	return BaseSolverService{}
}

func (b BaseSolverService) Solve(req *model.CalcRequest) (*model.SolutionResponse, error) {
	stack := convertRequestToStack(req)

	b.Solver = core.NewBasicSolver(stack)
	solution := b.Solver.Solve()
	result := model.NewResult()

	for solution.Size() > 0 {
		v, err := solution.Pop()
		if err != nil {
			return nil, err
		}
		vr := v.Items["var"].(string)
		va := v.Items["value"].(string)
		result.Items = append(result.Items, model.NewPrintResult(vr, va))
	}
	return result, nil
}

func convertRequestToStack(req *model.CalcRequest) *core.Stack[core.Command] {

	stack := core.NewStack[core.Command]()

	for i := len(req.Operations) - 1; i >= 0; i-- {
		op := req.Operations[i]
		if op.Type == "calc" {
			stack.Push(core.NewCalcCommand(op.Op, op.Var, op.Left, op.Right))
		} else {
			stack.Push(core.NewPrintCommand(op.Var))
		}
	}

	return stack
}
