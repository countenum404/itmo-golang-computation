package service

import (
	"countenum404/itmo-golang-computation/internal/core"
	"countenum404/itmo-golang-computation/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func commandsEqual(a, b core.Command) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Equals(b)
}

func TestConvertRequestToStack(t *testing.T) {
	tests := []struct {
		name     string
		req      *model.CalcRequest
		expected []core.Command
	}{
		{
			name: "Single Calc Operation",
			req: &model.CalcRequest{
				Operations: []*model.Operation{
					{Type: "calc", Op: "+", Var: "result", Left: "a", Right: "b"},
				},
			},
			expected: []core.Command{
				core.NewCalcCommand("+", "result", "a", "b"),
			},
		},
		{
			name: "Multiple Operations",
			req: &model.CalcRequest{
				Operations: []*model.Operation{
					{Type: "calc", Op: "*", Var: "result1", Left: "x", Right: "y"},
					{Type: "print", Var: "result1"},
					{Type: "calc", Op: "-", Var: "result2", Left: "m", Right: "n"},
				},
			},
			expected: []core.Command{
				core.NewCalcCommand("*", "result1", "x", "y"),
				core.NewPrintCommand("result1"),
				core.NewCalcCommand("-", "result2", "m", "n"),
			},
		},
		{
			name: "Only Print Operations",
			req: &model.CalcRequest{
				Operations: []*model.Operation{
					{Type: "print", Var: "result2"},
					{Type: "print", Var: "result1"},
				},
			},
			expected: []core.Command{
				core.NewPrintCommand("result2"),
				core.NewPrintCommand("result1"),
			},
		},
		{
			name:     "Empty Request",
			req:      &model.CalcRequest{Operations: []*model.Operation{}},
			expected: []core.Command{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := convertRequestToStack(tt.req)

			if stack.Size() != len(tt.expected) {
				t.Fatalf("expected stack size %d, got %d", len(tt.expected), stack.Size())
			}

			for i := 0; i < len(tt.expected); i++ {
				cmd, err := stack.Pop()
				if err != nil {
					t.Fatalf("unexpected error popping stack: %v", err)
				}
				expectedCmd := tt.expected[i]
				if !commandsEqual(cmd, expectedCmd) {
					t.Errorf("at index %d, expected command %+v, got %+v", i, expectedCmd, cmd)
				}
			}
		})
	}
}

func TestBaseSolverService_Solve(t *testing.T) {
	request := &model.CalcRequest{
		Operations: []*model.Operation{
			{Type: "calc", Op: "+", Var: "sum", Left: "1", Right: "2"},
			{Type: "print", Var: "sum"},
		},
	}

	svc := BaseSolverService{}

	response, err := svc.Solve(request)

	assert.Nil(t, err)

	assert.NotNil(t, response)
	assert.Len(t, response.Items, 1)

	firstItem := response.Items[0]
	assert.Equal(t, firstItem.Var, "sum")
	assert.Equal(t, firstItem.Value, "3")
}
