package fsm

import (
	"fmt"
	"strings"
)

type State string

type Symbol string

type TransitionFunction func(State, Symbol) State

type FiniteAutomaton struct {
	States             []State
	Alphabet           []Symbol
	InitialState       State
	AcceptingStates    []State
	TransitionFunction TransitionFunction
}

func NewFiniteAutomaton(
	states []State,
	alphabet []Symbol,
	initialState State,
	acceptingStates []State,
	transitionFunction TransitionFunction,
) *FiniteAutomaton {
	return &FiniteAutomaton{
		States:             states,
		Alphabet:           alphabet,
		InitialState:       initialState,
		AcceptingStates:    acceptingStates,
		TransitionFunction: transitionFunction,
	}
}

func (fa *FiniteAutomaton) ProcessInput(input string) (State, error) {
	currentState := fa.InitialState

	for i, char := range input {
		symbol := Symbol(string(char))

		if !fa.isValidSymbol(symbol) {
			return "", fmt.Errorf("invalid symbol '%s' at position %d: not in alphabet %v", symbol, i, fa.Alphabet)
		}

		currentState = fa.TransitionFunction(currentState, symbol)
	}

	return currentState, nil
}

func (fa *FiniteAutomaton) isValidSymbol(symbol Symbol) bool {
	for _, s := range fa.Alphabet {
		if s == symbol {
			return true
		}
	}
	return false
}

func (fa *FiniteAutomaton) IsAcceptingState(state State) bool {
	for _, acceptingState := range fa.AcceptingStates {
		if acceptingState == state {
			return true
		}
	}
	return false
}

func (fa *FiniteAutomaton) GetStates() []State {
	return fa.States
}

func (fa *FiniteAutomaton) GetAlphabet() []Symbol {
	return fa.Alphabet
}

func (fa *FiniteAutomaton) GetInitialState() State {
	return fa.InitialState
}

func (fa *FiniteAutomaton) GetAcceptingStates() []State {
	return fa.AcceptingStates
}

func (fa *FiniteAutomaton) String() string {
	var sb strings.Builder
	sb.WriteString("Finite Automaton:\n")
	sb.WriteString(fmt.Sprintf("  States: %v\n", fa.States))
	sb.WriteString(fmt.Sprintf("  Alphabet: %v\n", fa.Alphabet))
	sb.WriteString(fmt.Sprintf("  Initial State: %s\n", fa.InitialState))
	sb.WriteString(fmt.Sprintf("  Accepting States: %v\n", fa.AcceptingStates))
	return sb.String()
}
