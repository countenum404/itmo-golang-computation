// Package core provides the implementation of the computation/*
package core

import (
	"fmt"
	"strconv"
	"sync"
)

var operations = map[string]func(left, right int) int{
	"+": func(left, right int) int { return left + right },
	"-": func(left, right int) int { return left - right },
	"*": func(left, right int) int { return left * right },
	"/": func(left, right int) int { return left / right },
}

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
	AsyncExecute(vars *sync.Map) (*Result, error)
	Equals(command Command) bool
}

// CalcCommand Operation with Left and Right operands/*
type CalcCommand struct {
	Op    string
	Var   string
	Left  string
	Right string
}

func (oc *CalcCommand) Equals(command Command) bool {
	if other, ok := command.(*CalcCommand); ok {
		return oc.Op == other.Op && oc.Var == other.Var && oc.Left == other.Left && oc.Right == other.Right
	}
	return false
}

func readWithRetry(key string, vars *sync.Map, valueChan chan string) {
	var v any = nil
	var ok = false
	for !ok {
		v, ok = vars.Load(key)
	}
	valueChan <- v.(string)
}

func getOperand(key string, vars *sync.Map) int {
	value, err := strconv.Atoi(key)
	if err != nil {
		valueChan := make(chan string)
		go readWithRetry(key, vars, valueChan)
		valueMemory := <-valueChan
		value, _ = strconv.Atoi(valueMemory)
	}
	return value
}

func (oc *CalcCommand) AsyncExecute(vars *sync.Map) (*Result, error) {
	_, ok := vars.Load(oc.Var)
	if ok {
		return nil, fmt.Errorf("cannot overwrite variable: %s", oc.Var)
	}
	left := getOperand(oc.Left, vars)
	right := getOperand(oc.Right, vars)
	vars.Store(oc.Var, strconv.Itoa(operations[oc.Op](left, right)))
	return nil, nil
}

func NewCalcCommand(op, v, l, r string) *CalcCommand {
	return &CalcCommand{Op: op, Var: v, Left: l, Right: r}
}

func (oc *CalcCommand) Execute(vars map[string]string) (*Result, error) {
	_, ok := vars[oc.Var]
	if ok {
		return nil, fmt.Errorf("cannot overwrite variable: %s", oc.Var)
	}

	left, err := strconv.Atoi(oc.Left)

	if err != nil {
		leftVar := vars[oc.Left]
		left, _ = strconv.Atoi(leftVar)
	}

	right, err := strconv.Atoi(oc.Right)

	if err != nil {
		rightVar := vars[oc.Right]
		right, _ = strconv.Atoi(rightVar)
	}

	vars[oc.Var] = strconv.Itoa(operations[oc.Op](left, right))

	return nil, nil
}

// PrintCommand /*
type PrintCommand struct {
	Var string
}

func (pc *PrintCommand) Equals(command Command) bool {
	if other, ok := command.(*PrintCommand); ok {
		return pc.Var == other.Var
	}
	return false
}

func (pc *PrintCommand) AsyncExecute(vars *sync.Map) (*Result, error) {
	valueChan := make(chan string)
	go readWithRetry(pc.Var, vars, valueChan)
	value := <-valueChan
	return NewResult(map[any]any{"var": pc.Var, "value": value}), nil
}

func NewPrintCommand(v string) *PrintCommand { return &PrintCommand{Var: v} }

func (pc *PrintCommand) Execute(vars map[string]string) (*Result, error) {
	return NewResult(map[any]any{"var": pc.Var, "value": vars[pc.Var]}), nil
}
