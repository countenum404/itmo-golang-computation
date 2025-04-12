package model

type Result struct {
	Items []*PrintResult
}

func NewResult() *Result {
	return &Result{Items: make([]*PrintResult, 0)}
}

type PrintResult struct {
	Var   string
	Value string
}

func NewPrintResult(variable string, value string) *PrintResult {
	return &PrintResult{Var: variable, Value: value}
}

type Operation struct {
	Type  string `json:"type"`
	Op    string `json:"op,omitempty"`
	Var   string `json:"var"`
	Left  string `json:"left"`
	Right string `json:"right"`
}

type Request struct {
	Operations []Operation `json:"operations"`
}
