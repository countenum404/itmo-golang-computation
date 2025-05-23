package core

import (
	"sync"
)

/*
Solver is a stack based command pipeline
*/
type Solver interface {
	Solve() *Stack[*Result]
}

type BasicSolver struct {
	Variables map[string]string
	Commands  *Stack[Command]
}

func NewBasicSolver(commands *Stack[Command]) *BasicSolver {
	return &BasicSolver{Commands: commands, Variables: make(map[string]string)}
}

func (bs *BasicSolver) Solve() *Stack[*Result] {
	result := NewStack[*Result]()
	for c, err := bs.Commands.Pop(); err == nil; c, err = bs.Commands.Pop() {
		r, err := c.Execute(bs.Variables)
		if r != nil {
			result.Push(r)
		}
		if err != nil {
			panic(err)
		}
	}
	return result
}

type AsyncSolver struct {
	Variables *sync.Map
	Commands  *Stack[Command]
	printChan chan *Result
	calcChan  chan *Result
}

func NewAsyncSolver(commands *Stack[Command]) *AsyncSolver {
	return &AsyncSolver{Commands: commands, Variables: new(sync.Map)}
}

func (as *AsyncSolver) Solve() *Stack[*Result] {
	result := NewStack[*Result]()
	var wg sync.WaitGroup
	wg.Add(as.Commands.Size())
	for c, err := as.Commands.Pop(); err == nil; c, err = as.Commands.Pop() {
		go func(cmd Command) {
			defer wg.Done()
			r, _ := cmd.AsyncExecute(as.Variables)
			if r != nil {
				result.SafePush(r)
			}
		}(c)
	}
	wg.Wait()
	return result
}
