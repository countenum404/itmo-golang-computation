package handlers

import (
	"countenum404/itmo-golang-computation/internal/model"
	"countenum404/itmo-golang-computation/internal/service"
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Module("handlers", fx.Provide(fx.Annotate(NewSolverHandlers, fx.As(new(Handlers)))))

type Handlers interface {
	HandleCalc() func(w http.ResponseWriter, r *http.Request)
}

type SolverHandlers struct {
	SolverService service.SolverService
}

func NewSolverHandlers(lc fx.Lifecycle, solverService service.SolverService) *SolverHandlers {
	return &SolverHandlers{SolverService: solverService}
}

// @Summary Calculate result based on input parameters
// @Description This endpoint takes a calculation request and returns the result of the calculation.
// @Tags calculation
// @Accept json
// @Produce json
// @Param request body model.CalcRequest true "Calculation request"
// @Success 200 {object} model.SolutionResponse "Successful calculation result"
// @Router /calculate [post]
func (s *SolverHandlers) HandleCalc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		o := new(model.CalcRequest)
		ReadJson(r, o)

		resultModel, _ := s.SolverService.Solve(o)
		WriteJson(w, 200, resultModel)
	}
}
