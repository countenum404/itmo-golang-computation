package core

/*
Solver is a stack based command pipeline
*/
type Solver interface {
	Solve() *Stack[*Result]
}

type BasicSolver struct {
	Variables map[string]string
	Commands  Stack[Command]
}

func NewBasicSolver(commands *Stack[Command]) *BasicSolver {
	return &BasicSolver{Commands: *commands, Variables: make(map[string]string)}
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
