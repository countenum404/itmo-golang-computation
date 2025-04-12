/*
Core package provides the implementation of the computation
*/
package core

import (
	"fmt"
	"strconv"
)

type Result struct {
	Items map[any]any
}

func NewResult(items map[any]any) *Result {
	return &Result{Items: items}
}

/*
Command pattern allows easily to add new commands
*/
type Command interface {
	Execute(vars map[string]string) (*Result, error)
}

/*
Calc command: Operation with Left and Right operands
*/
type CalcCommand struct {
	Op    string
	Var   string
	Left  string
	Right string
}

func NewCalcCommand(op, v, l, r string) *CalcCommand {
	return &CalcCommand{Op: op, Var: v, Left: l, Right: r}
}

func (oc *CalcCommand) Execute(vars map[string]string) (*Result, error) {
	_, ok := vars[oc.Var]
	if ok {
		return nil, fmt.Errorf("cannot overwrite variable: %s", oc.Var)
	}

	operations := map[string]func(left, right int) int{
		"+": func(left, right int) int { return left + right },
		"-": func(left, right int) int { return left - right },
		"*": func(left, right int) int { return left * right },
		"/": func(left, right int) int { return left / right },
	}

	left, err := strconv.Atoi(oc.Left)

	if err != nil {
		lftval := vars[oc.Left]
		left, _ = strconv.Atoi(lftval)
	}

	right, err := strconv.Atoi(oc.Right)

	if err != nil {
		rgtval := vars[oc.Right]
		right, _ = strconv.Atoi(rgtval)
	}

	vars[oc.Var] = strconv.Itoa(operations[oc.Op](left, right))

	return nil, nil
}

/*
Print command
*/
type PrintCommand struct {
	Var string
}

func NewPrintCommand(v string) *PrintCommand { return &PrintCommand{Var: v} }

func (pc *PrintCommand) Execute(vars map[string]string) (*Result, error) {
	return NewResult(map[any]any{"var": pc.Var, "value": vars[pc.Var]}), nil
}
