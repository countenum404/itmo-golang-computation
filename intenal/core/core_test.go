package core

import (
	"testing"
)

func TestStack(t *testing.T) {

	testValue := 256

	stack := NewStack[int]()
	for i := 1; i <= testValue; i++ {
		stack.Push(i)
	}

	for i := testValue; i > 0; i-- {
		res, err := stack.Pop()
		t.Log(res)
		if res != i {
			t.Errorf("Invalid element res=%d, i=%d", res, i)
		}
		if err != nil {
			t.Errorf("Error is not nil")
		}
	}

	val, err := stack.Pop()
	if val != 0 {
		t.Errorf("Zero value is not returned")
	}
	if err == nil {
		t.Errorf("There is no error on the empty stack pop operation")
	}
}

func TestSanitySolver(t *testing.T) {
	stack := NewStack[Command]()
	stack.Push(NewPrintCommand("c"))
	stack.Push(NewPrintCommand("a"))
	stack.Push(NewPrintCommand("b"))
	stack.Push(NewPrintCommand("x"))

	stack.Push(NewCalcCommand("-", "c", "a", "b"))
	stack.Push(NewCalcCommand("/", "b", "x", "5"))
	stack.Push(NewCalcCommand("*", "a", "x", "20"))
	stack.Push(NewCalcCommand("+", "x", "10", "20"))

	s := NewBasicSolver(stack)
	rs := s.Solve()

	expectedKeys := []string{"c", "a", "b", "x"}
	expectedValues := []string{"594", "600", "6", "30"}

	for i := 0; i < len(expectedKeys); i++ {
		item, err := rs.Pop()
		if err != nil {
			t.Errorf("Error in stack: %s", err)
		}
		expectedKey, expectedValue := expectedKeys[i], expectedValues[i]
		actualVar := item.Items["var"]
		actualValue := item.Items["value"]

		if actualValue != expectedValue || actualVar != expectedKey {
			t.Errorf("Expected variable %s to be %s; Got %s", expectedKey, expectedValue, actualValue)
		}
	}

}
