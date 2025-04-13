package model

import "countenum404/itmo-golang-computation/pkg/app"

type SolutionResponse struct {
	Items []*PrintResult
}

func NewResult() *SolutionResponse {
	return &SolutionResponse{Items: make([]*PrintResult, 0)}
}

type PrintResult struct {
	Var   string
	Value string
}

func NewPrintResult(variable string, value string) *PrintResult {
	return &PrintResult{Var: variable, Value: value}
}

type Operation struct {
	*app.Operation
	Type  string `json:"type"`
	Op    string `json:"op,omitempty"`
	Var   string `json:"var"`
	Left  string `json:"left"`
	Right string `json:"right"`
}

type CalcRequest struct {
	*app.CalcRequest
	Operations []*Operation `json:"operations"`
}
